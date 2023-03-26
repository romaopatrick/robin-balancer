package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var port int64
	flag.Int64Var(&port, "port", 8080, "Port to listen")
	flag.Parse()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	serve(fmt.Sprintf(":%d", port), r)

}

func serve(addr string, r *gin.Engine) {
	server := http.Server{
		Addr:    addr,
		Handler: r.Handler(),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Printf("turning off server %s", addr)
		log.Fatalf("ListenAndServe: %v", err)

	}
}
