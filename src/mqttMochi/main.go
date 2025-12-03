package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func main() {
	// 创建信号用于等待服务端关闭信号
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	// 创建新的 MQTT 服务器。
	server := mqtt.New(nil)

	// 允许所有连接(权限)。
	_ = server.AddHook(new(auth.AllowHook), nil)

	// 在标1883端口上创建一个 TCP 服务端。
	tcp := listeners.NewTCP(listeners.Config{ID: "t1", Address: ":1883"})
	err := server.AddListener(tcp)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// 服务端等待关闭信号
	<-done

	// 关闭服务端时需要做的一些清理工作
}
