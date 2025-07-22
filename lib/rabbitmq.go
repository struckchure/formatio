package lib

import (
	"context"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	connection *amqp.Connection
}

type PublishArgs struct {
	Queue   string
	Content string
}

type SubscribeArgs struct {
	Queue    string
	Callback func(content string) error
}

func (r *RabbitMQ) Publish(args PublishArgs) error {
	ch, err := r.connection.Channel()
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		args.Queue, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(args.Content),
		})
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	return nil
}

func (r *RabbitMQ) Subscribe(args SubscribeArgs) error {
	ch, err := r.connection.Channel()
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		args.Queue, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Printf("[x][%s] %s", args.Queue, err)

		return err
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			err := args.Callback(string(d.Body))
			if err != nil {
				log.Printf("[*][%s] Error %s", args.Queue, err)
				break
			}

			d.Ack(false)
		}
	}()

	log.Printf("[*][%s] Waiting for messages. To exit press CTRL+C", args.Queue)
	<-forever

	return nil
}

func (r *RabbitMQ) SubscribeWithWorkers(workers int, args SubscribeArgs) {
	var wg sync.WaitGroup

	log.Printf("[x][%s] Spawning %d workers", args.Queue, workers)

	for i := 1; i <= workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			r.Subscribe(args)
		}()

		log.Printf("[x][%s] Worker %d just joined the party", args.Queue, i)
	}

	wg.Wait()
}

func NewRabbitMQ(connection *amqp.Connection) RabbitMQ {
	return RabbitMQ{connection: connection}
}

func NewRabbitMQConnection(env Env) *amqp.Connection {
	connection, err := amqp.Dial(env.RABBITMQ_URL)

	if err != nil {
		log.Printf("[x] %s: while connecting", err)
	}

	log.Println("[x] connection established")

	return connection
}
