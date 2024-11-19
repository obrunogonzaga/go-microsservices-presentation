package setup

func SetupQueues(ch *amqp.Channel) error {
	// Declara Exchange para DLQ
	err := ch.ExchangeDeclare(
			"dlx",      // nome
			"direct",   // tipo
			true,       // durable
			false,      // auto-delete
			false,      // internal
			false,      // no-wait
			nil,        // arguments
	)

	// Configuração da fila principal
	args := amqp.Table{
			"x-dead-letter-exchange": "dlx",
			"x-dead-letter-routing-key": "notifications.failed",
			"x-message-ttl": 30000,  // 30 segundos
	}

	// Declara fila principal
	_, err = ch.QueueDeclare(
			"notifications",  // nome
			true,            // durable
			false,           // delete when unused
			false,           // exclusive
			false,           // no-wait
			args,            // arguments
	)

	// Declara DLQ
	_, err = ch.QueueDeclare(
			"notifications.dlq",  // nome
			true,                // durable
			false,               // delete when unused
			false,               // exclusive
			false,               // no-wait
			nil,                // arguments
	)

	// Bind DLQ ao exchange
	err = ch.QueueBind(
			"notifications.dlq",    // queue name
			"notifications.failed", // routing key
			"dlx",                 // exchange
			false,                 // no-wait
			nil,                  // arguments
	)

	return err
}

// Consumidor da DLQ
func ProcessDLQ(ch *amqp.Channel) {
	msgs, _ := ch.Consume(
			"notifications.dlq",
			"",
			false,  // auto-ack
			false,
			false,
			false,
			nil,
	)

	for msg := range msgs {
			// Log da mensagem falha
			log.Error("mensagem falha",
					"body", string(msg.Body),
					"retries", msg.Headers["x-retry-count"],
			)

			// Tenta reprocessar após delay
			if retries < 3 {
					// Republica na fila original
					ch.Publish(
							"",               // exchange
							"notifications",  // routing key
							false,
							false,
							amqp.Publishing{
									Headers: amqp.Table{
											"x-retry-count": retries + 1,
									},
									Body: msg.Body,
							},
					)
			}

			msg.Ack(false)
	}
}