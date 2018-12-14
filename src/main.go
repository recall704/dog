package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/recall704/go-ipset/ipset"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	r.GET("/add/:ip", HandleIPsetAdd)
	r.GET("/del/:ip", HandleIPsetDel)
	r.GET("/list", HandleGetIPSetList)

	r.Run(":8080")
}

var gost *ipset.IPSet

func init() {
	var err error
	gost, err = ipset.New("gost", "hash:net", &ipset.Params{
		Exist: true,
	})
	if err != nil {
		log.Error("get ipset err", err)
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
		log.Error("add to ipset error", err)
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
		log.Error("delete from ipset error", err)
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
		log.Error("get ipset list error", err)
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
