// kafkaconsumer project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

func StartConsumer(group string, zkaddrs []string, topics []string) error {
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ResetOffsets = true

	consumer, err := consumergroup.JoinConsumerGroup(
		group, topics, zkaddrs, config)
	if nil != err {
		return err
	}

	go func() {
		for event := range consumer.Messages() {
			fmt.Println(fmt.Sprintf("Receive Time:%s, Topic :%s, Partition:%d, Key:%s, Value:%s.",
				time.Now().Format("2006-01-02 15:04:05"), event.Topic, event.Partition,
				event.Key, event.Value))
		}
	}()

	return nil
}

func main() {
	err := StartConsumer("new_message_group", []string{"192.168.144.74:2181"},
		[]string{"pc-search", "pc-ware", "pc-cart", "pc-order"})
	if nil != err {
		fmt.Println(fmt.Sprintf("Start consumer failed, %s.", err.Error()))
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(0)
	}

	tick := time.Tick(time.Second * time.Duration(10))
	for {
		<-tick
		fmt.Println(fmt.Sprintf("I'm alive ,Time:%s.", time.Now().Format("2006-01-02 15:04:05")))
	}

}
