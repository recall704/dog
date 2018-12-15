package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/recall704/dog/src/log"
	"github.com/recall704/go-ipset/ipset"
	"github.com/sirupsen/logrus"
)

var (
	ip       = flag.String("ip", "0.0.0.0", "server ip.")
	port     = flag.Int("port", 5051, "server port.")
	name     = flag.String("name", "gost", "ipset name")
	logLevel = flag.String("log_level", "info", "log level")
)

func main() {
	r := gin.Default()
	r.GET("/add/:ip", HandleIPsetAdd)
	r.GET("/del/:ip", HandleIPsetDel)
	r.GET("/list", HandleGetIPSetList)

	r.Run(fmt.Sprintf("%s:%d", *ip, *port))
}

var gost *ipset.IPSet

func init() {
	flag.Parse()
	log.Init(*logLevel)

	var err error
	gost, err = ipset.New(*name, "hash:net", &ipset.Params{
		Exist: true,
	})
	if err != nil {
		logrus.Error("get ipset err", err)
		os.Exit(1)
	}

}

// HandleIPsetAdd ...
func HandleIPsetAdd(c *gin.Context) {
	ip := strings.TrimSpace(c.Param("ip"))

	if ip == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid ip",
		})
		return
	}

	if err := gost.Add(ip, 0); err != nil {
		logrus.Error("add to ipset error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "add ip to ipset OK",
	})
	return
}

// HandleIPsetDel ...
func HandleIPsetDel(c *gin.Context) {
	ip := strings.TrimSpace(c.Param("ip"))

	if ip == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid ip",
		})
		return
	}

	if err := gost.Del(ip); err != nil {
		logrus.Error("delete from ipset error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "del ip from ipset OK",
	})
	return
}

// HandleGetIPSetList ...
func HandleGetIPSetList(c *gin.Context) {

	ipList, err := gost.List()
	if err != nil {
		logrus.Error("get ipset list error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list": ipList,
	})
	return
}
