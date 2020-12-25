package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {

	p, err := rocketmq.NewProducer(
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{"192.168.50.18:9876"})),
		producer.WithRetry(2),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: "RocketMQ",
			SecretKey: "12345678",
		}),
	)

	if err != nil {
		fmt.Println("init producer error: " + err.Error())
		os.Exit(0)
	}

	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	for i := 0; i < 100000; i++ {
		res, err := p.SendSync(context.Background(), primitive.NewMessage("test",
			[]byte("Hello RocketMQ Go Client!")))

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
