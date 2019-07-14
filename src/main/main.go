package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"util"
)

const versionName = "v1"

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Println("useage : serverMonitor [ip]:[port] (eg: :8080| 0.0.0.0:8080)")
		return
	}

	kernelInfo := util.GetKernelInfo()
	log.Println("kernel Info:", kernelInfo)

	addr := args[1]
	log.Println("start monitor at ", addr)

	// 发布模式
	gin.SetMode(gin.ReleaseMode)
	log.Println("start device data...")
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
	router := gin.Default()

	g := router.Group(versionName)
	{
		g.GET("/GetServerInfo", handler.GetServerInfo)
	}

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("start server error,", err)
		}
	}()

	// gracefully shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("shuntdown server error", err)
	}
	log.Println("Server exiting")
}
