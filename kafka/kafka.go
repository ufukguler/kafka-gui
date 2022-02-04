package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
	"strconv"
)

func Dial(addr string) (*kafka.Conn, error) {
	return kafka.Dial("tcp", addr)
}

func DialLeader(topic, partition, addr string) (*kafka.Conn, error) {
	parseInt, err := strconv.Atoi(partition)
	if err != nil {
		return nil, err
	}
	return kafka.DialLeader(context.Background(), "tcp", addr, topic, parseInt)
}

func GetTopics(addr string) ([]string, error) {
	arr := make([]string, 0)
	conn, err := kafka.Dial("tcp", addr)
	if err != nil {
		return arr, err
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return arr, err
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k, _ := range m {
		arr = append(arr, k)
	}
	return arr, err
}

func CreateTopic(topic, partition, addr string) (*kafka.Conn, error) {
	return DialLeader(topic, partition, addr)
}

func Produce(topic, partition, addr, data string, header protocol.Header) error {
	partitionInt, err := strconv.Atoi(partition)
	if err != nil {
		return err
	}
	w := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	message := kafka.Message{
		Partition: partitionInt,
		Headers:   []kafka.Header{header},
		Value:     []byte(data),
	}
	defer w.Close()
	return w.WriteMessages(context.Background(), message)
}
