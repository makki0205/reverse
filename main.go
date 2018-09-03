package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(cros)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "alive")
	})
	r.POST("/", func(c *gin.Context) {
		req := struct {
			Method  string            `json:"method"`
			URL     string            `json:"url"`
			Body    string            `json:"body"`
			Headers map[string]string `json:"header"`
		}{}
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		r, _ := http.NewRequest(
			req.Method,
			req.URL,
			strings.NewReader(req.Body),
		)
		for key, value := range req.Headers {
			r.Header.Set(key, value)
		}

		// Content-Type 設定
		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		var body []byte
		res.Body.Read(body)

		start, ok := c.Get("start_time")
		if ok {
			c.Header("X-Server-Latency", time.Since(start.(time.Time)).String())
		}
		c.JSON(http.StatusOK, gin.H{
			"status": res.StatusCode,
			"body":   string(body),
			"header": res.Header,
		})
	})
	r.Run(":3000")
}

func cros(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", headers)
	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	}
	c.Set("start_time", time.Now())
	c.Next()
}
