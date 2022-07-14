package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

var (
	fixTopic           = flag.String("creative_fix_kafka_topic", "ads_midas_creative-fix", "kafka topic for limit budget log topic")
	fixBrokers         = flag.String("creative_fix_kafka_broker", "test.mw.cbs.sg1.kafka:9092", "kafka address for limit budget log change broker")
	fixGroup           = flag.String("creative_fix_kafka_group", "group1", "kafka group")
	adSetKafkaInitMode = flag.Int("creative_adsetinfo_kafka_init_mode", 2, "consume mode 1 for oldest 2 for newest")
)

type FixedAdCreativeInfo struct {
	AdId uint64 `json:"ad_id"`
}
type FixedAdCreativeKafkaMsg struct {
	Content *FixedAdCreativeInfo `json:"after"`
	Source  *SourceKafka         `json:"source"`
	Op      string               `json:"op"`
	TsMs    uint64               `json:"ts_ms"`
}

// SourceKafka kafka公共source参数
type SourceKafka struct {
	Version   string `json:"version"`
	Connector string `json:"connector"`
	Name      string `json:"name"`
	TsMs      uint64 `json:"ts_ms"`
	Db        string `json:"db"`
	Table     string `json:"table"`
}

func doneData() {
	data := FixedAdCreativeKafkaMsg{
		Content: &FixedAdCreativeInfo{AdId: 444},
		Source:  &SourceKafka{},
	}
	jsonData, _ := json.Marshal(data)
	//fmt.Println(jsonData)
	//config := sarama.NewConfig()
	//config.Producer.Flush.Messages = 10
	//config.Producer.Return.Successes = true
	//broker := []string{"test.mw.cbs.sg1.kafka:9092"}
	//
	//asyncProducer, err := sarama.NewAsyncProducer(broker, config)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//asyncProducer.Input() <- &sarama.ProducerMessage{
	//	Topic: "ads_midas_creative-fix",
	//	Value: sarama.StringEncoder(jsonData),
	//}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	client, err := sarama.NewClient([]string{"test.mw.cbs.sg1.kafka:9092"}, config)
	if err != nil {
		log.Fatalf("unable to create kafka client: %q", err)
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Fatalf("unable to create kafka producer: %q", err)
	}
	defer producer.Close()
	topic := "ads_midas_creative-fix"
	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(jsonData)})
	if err != nil {
		log.Fatalf("unable to produce message: %q", err)
	}
	fmt.Println(partition, offset)
}
func main() {
	doneData()
}
