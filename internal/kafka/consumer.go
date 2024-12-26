package kafka

import (
	"encoding/json"
	"go-rest-api-kafka/internal/models"
	"strings"

	"github.com/IBM/sarama"
	"gorm.io/gorm"
)

type Consumer struct {
    consumer sarama.Consumer
    db       *gorm.DB
    topic    string
}

func NewConsumer(brokers string, topic string, db *gorm.DB) (*Consumer, error) {
    config := sarama.NewConfig()
    consumer, err := sarama.NewConsumer(strings.Split(brokers, ","), config)
    if err != nil {
        return nil, err
    }

    return &Consumer{
        consumer: consumer,
        db:       db,
        topic:    topic,
    }, nil
}

func (c *Consumer) Start() error {
    partitionConsumer, err := c.consumer.ConsumePartition(c.topic, 0, sarama.OffsetNewest)
    if err != nil {
        return err
    }

    go func() {
        for msg := range partitionConsumer.Messages() {
            var plan models.Plan
            if err := json.Unmarshal(msg.Value, &plan); err != nil {
                continue
            }

            c.db.Create(&plan)
        }
    }()

    return nil
}