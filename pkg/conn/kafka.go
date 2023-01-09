package conn

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
)

func CreateProducer(conf *rkentry.ConfigEntry) *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": conf.GetString("kafka_addr")})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		panic(err)
	}
	return p
}
