package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"Trpcx/service"
	"github.com/kataras/iris/v12"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := service.NewService()
	errChan := make(chan error)

	//按 CTRL+C 时停止服务器
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// 映射端点
	endpoints := service.Endpoints{
		GetEndpoint:      service.MakeGetEndpoint(srv),
		StatusEndpoint:   service.MakeStatusEndpoint(srv),
		ValidateEndpoint: service.MakeValidateEndpoint(srv),
	}

	// HTTP 传输
	go func() {
		log.Println("service is listening on port:", *httpAddr)
		app := service.NewHTTPServer(ctx, endpoints)
		errChan <- app.Run(iris.Addr(*httpAddr))
	}()

	log.Fatalln(<-errChan)

}
