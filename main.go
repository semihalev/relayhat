package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"

	log "github.com/inconshreveable/log15"
	rpio "github.com/stianeikeland/go-rpio"
)

const (
	portCount = 4

	dbFile  = "/etc/relayhat.json"
	version = "0.1"
)

// RelayHat structure
type RelayHat struct {
	sync.RWMutex

	names [portCount]string
	pins  [portCount]*rpio.Pin
}

type port struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func (rh *RelayHat) status() (data [portCount]port) {
	for i, pin := range rh.pins {
		data[i].Status = pin.Read() == rpio.High
		data[i].Name = rh.names[i]
	}

	return
}

func (rh *RelayHat) portsStatus(c *gin.Context) {
	rh.RLock()
	defer rh.RUnlock()

	c.JSON(http.StatusOK, rh.status())
}

func (rh *RelayHat) saveNames() error {
	f, err := os.OpenFile(dbFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	err = f.Truncate(0)
	if err != nil {
		return err
	}

	buf, err := json.Marshal(rh.names)
	if err != nil {
		return err
	}

	_, err = f.Write(buf)

	return err
}

func (rh *RelayHat) readNames() error {
	f, err := os.OpenFile(dbFile, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 1024)

	n, err := f.Read(buf)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf[:n], &rh.names)

	return err
}

func (rh *RelayHat) togglePort(c *gin.Context) {
	rh.Lock()
	defer rh.Unlock()

	port, err := strconv.Atoi(c.Param("port"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if port < 0 || port >= portCount {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "port invalid"})
		return
	}

	rh.pins[port].Toggle()

	c.JSON(http.StatusOK, rh.status())
}

func (rh *RelayHat) setName(c *gin.Context) {
	rh.Lock()
	defer rh.Unlock()

	port, err := strconv.Atoi(c.Param("port"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if port < 0 || port >= portCount {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "port invalid"})
		return
	}

	rh.names[port] = c.Param("name")
	err = rh.saveNames()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rh.status())
}

func showDash(c *gin.Context) {
	//dashboard, err := template.New("dashboard.html").Delims("[[", "]]").ParseFiles("templates/dashboard.html")

	dashboardTemplate, err := templatesDashboardHtml()
	if err != nil {
		c.String(http.StatusOK, "template execution: %s", err)
		return
	}
	dashboard, _ := template.New("dashboard.html").Delims("[[", "]]").Parse(string(dashboardTemplate.bytes))

	type htmlData struct {
		Version string
	}

	html := htmlData{Version: "v" + version}
	var doc bytes.Buffer

	err = dashboard.Execute(&doc, html)
	if err != nil {
		c.String(http.StatusOK, "template execution: %s", err)
		return
	}

	c.Writer.Header().Set("Content-Type", "text/html")
	c.String(http.StatusOK, doc.String())
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	err := rpio.Open()
	if err != nil {
		log.Crit("GPIO open failed", "error", err.Error())
		os.Exit(1)
	}
	defer rpio.Close()

	rh := new(RelayHat)

	err = rh.readNames()
	if err != nil {
		log.Warn("Port names read failed", "error", err.Error())

		if strings.Contains(err.Error(), "no such file or directory") {
			for i := range rh.names {
				rh.names[i] = fmt.Sprintf("Port %d", i+1)
			}

			err := rh.saveNames()
			if err != nil {
				log.Error("Port names write failed", "error", err.Error())
			}
		}
	}

	for i := range rh.pins {
		p := rpio.Pin(uint8(i) + 21)
		p.Output()
		p.Low()

		rh.pins[i] = &p
	}

	r := gin.Default()
	r.Use(stats.RequestStats())

	r.GET("/", showDash)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/status", rh.portsStatus)
		v1.GET("/toggle/:port", rh.togglePort)
		v1.GET("/setName/:port/:name", rh.setName)
		v1.GET("/stats", func(c *gin.Context) { c.JSON(200, stats.Report()) })
	}

	go func() {
		srv := &http.Server{
			Addr:         ":10080",
			Handler:      r,
			ReadTimeout:  60 * time.Second,
			WriteTimeout: 60 * time.Second,
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Crit("HTTP server listen fault", "error", err.Error())
			os.Exit(1)
		}
	}()

	log.Info("RelayHAT service started", "version", "v"+version)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	<-c

	log.Info("RelayHAT service stopping")
}
