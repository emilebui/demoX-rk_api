package repositories

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaRepo struct {
	p     *kafka.Producer
	topic string
}

func NewKafkaRepo(p *kafka.Producer, topic string) Repo {
	return &KafkaRepo{
		p:     p,
		topic: topic,
	}
}

func (r *KafkaRepo) PushMessage(msg string) error {

	err := r.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &r.topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(msg),
	}, nil)

	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
		return err
	}

	return nil
}
