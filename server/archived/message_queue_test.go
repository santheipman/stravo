package main

//
//import (
//	"fmt"
//	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
//	"log"
//	"strings"
//	"testing"
//)
//
//func TestPublish(t *testing.T) {
//	if err := publish(); err != nil {
//		t.Error(err)
//	}
//}
//
//func publish() error {
//	p, err := kafka.NewProducer(&kafka.ConfigMap{
//		"bootstrap.servers": "localhost:9094",
//		//"client.id": socket.gethostname(),
//		"acks": "all"})
//
//	if err != nil {
//		//fmt.Printf("Failed to create producer: %s\n", err)
//		//os.Exit(1)
//		return err
//	}
//
//	//users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
//	//items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}
//	////topic := "purchases"
//	topic := "top1"
//	//
//	//for n := 0; n < 10; n++ {
//	//	key := users[0]
//	//	data := items[0]
//	//	p.Produce(&kafka.Message{
//	//		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
//	//		Key:            []byte(key),
//	//		Value:          []byte(data),
//	//	}, nil)
//	//}
//	//
//	//time.Sleep(10 * time.Second)
//
//	deliv := make(chan kafka.Event)
//	err = p.Produce(&kafka.Message{
//		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
//		Value:          []byte("val12"),
//		Key:            []byte("key1"),
//	},
//		deliv, // delivery channel
//	)
//
//	if err != nil {
//		fmt.Printf("failed to produce msg: %s\n", err)
//		return err
//	}
//
//	out := <-deliv
//	fmt.Println(out)
//
//	for e := range p.Events() {
//		log.Println("event", e)
//	}
//
//	return nil
//}
//
////func consumer() error {
////
////}
//
//func replaceWords(dictionary []string, sentence string) string {
//	dictSet := make(map[string]struct{})
//
//	for _, word := range dictionary {
//		dictSet[word] = struct{}{}
//	}
//
//	var result strings.Builder
//	result.String()
//}
