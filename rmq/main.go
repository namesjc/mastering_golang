package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"

	"github.com/streadway/amqp"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(99))
}

func main() {
	// Load client cert
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("Error loading client cert: %v", err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile("cacert.pem")
	if err != nil {
		log.Fatalf("Error loading CA cert: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// TLS config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	// AMQP URI
	amqpURI := "amqps://user:ocA7o50VekMzIRd5@rabbitmq:5671"

	// Dial
	conn, err := amqp.DialTLS(amqpURI, tlsConfig)
	if err != nil {
		log.Fatalf("Error dialing RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Open channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error opening channel: %v", err)
	}
	defer ch.Close()

	// Declare queue
	queueName := "myqueue"
	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error declaring queue: %v", err)
	}

	// Publish message
	for i := 0; i < 100; i++ {
		message := "Hello, " + RandomString(6)
		err = ch.Publish("", queueName, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}
	}

	// Consume message
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error consuming message: %v", err)
	}

	// Process message
	for msg := range msgs {
		log.Printf("Received message: %s", msg.Body)
	}

}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
