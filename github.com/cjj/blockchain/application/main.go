package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cjj/blockchain/application/blockchain"
	_ "github.com/cjj/blockchain/application/docs"
	"github.com/cjj/blockchain/application/pkg/setting"
	"github.com/cjj/blockchain/application/routers"
	"github.com/cjj/blockchain/application/service"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		log.Printf("时区设置失败 %s", err)
	}
	time.Local = timeLocal
	blockchain.Init()
	go service.Init()
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	if err := server.ListenAndServe(); err != nil {
		log.Printf("start http server failed %s", err)
	}
}
