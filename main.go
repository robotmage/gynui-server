package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const version = "0.1"

type DATA struct {
	Time  string `json:"time"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

func main() {
	boolPtr := flag.Bool("version", false, "a bool")
	flag.Parse()
	if *boolPtr {
		fmt.Println(version)
		return
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/sleep/:sleeptime", func(c *gin.Context) {
		str := c.Param("sleeptime")
		value, _ := strconv.Atoi(str)
		time.Sleep(time.Duration(rand.Intn(value)) * time.Second)
		c.JSON(200, gin.H{
			"sleep": str,
		})
	})

	r.POST("/store/press", func(c *gin.Context) {
		var data DATA
		err := c.BindJSON(&data)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"a": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"time":  data.Time,
			"value": data.Value,
			"unit":  data.Unit,
		})
	})

	r.Run("0.0.0.0:1986")
	//r.Run()
}
