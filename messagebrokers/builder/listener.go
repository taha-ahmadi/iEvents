package builder

import (
	"errors"
	"log"
	"os"

	"github.com/taha-ahmadi/iEvents/messagebrokers"
	"github.com/taha-ahmadi/iEvents/messagebrokers/amqp"
)

func NewEventListenerFromEnvironment() (messagebrokers.EventListener, error) {
	var listener messagebrokers.EventListener
	var err error

	if url := os.Getenv("AMQP_URL"); url != "" {
		log.Printf("connecting to AMQP broker at %s", url)

		listener, err = amqp.NewAMQPEventListenerFromEnvironment()
		if err != nil {
			return nil, err
		}
	} else if brokers := os.Getenv("KAFKA_BROKERS"); brokers != "" {
		log.Printf("connecting to Kafka brokers at %s", brokers)

		listener, err = kafka.NewKafkaEventListenerFromEnvironment()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Neither AMQP_URL nor KAFKA_BROKERS specified")
	}

	return listener, nil
}
