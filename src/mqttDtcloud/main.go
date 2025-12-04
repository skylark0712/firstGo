package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// 创建 MQTT 客户端配置
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("go-mqtt-client")

	// 创建 MQTT 客户端实例
	client := MQTT.NewClient(opts)

	// 连接到 MQTT 服务器
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 在连接成功后进行订阅和发布操作
	go func() {
		// 订阅主题
		if token := client.Subscribe("my/topic", 0, nil); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}

		// 发布消息
		for i := 0; i < 5; i++ {
			text := fmt.Sprintf("Message %d", i)
			token := client.Publish("my/topic", 0, false, text)
			token.Wait()
			fmt.Println("Published:", text)
			time.Sleep(time.Second)
		}
	}()

	// 等待退出信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// 断开与 MQTT 服务器的连接
	client.Disconnect(250)
}
