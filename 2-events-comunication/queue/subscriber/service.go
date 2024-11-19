package subscriber

type Subscriber interface {
	Subscribe() error
	ProcessEvent(event RequestEvent) error
}

type service struct {
	rabbitmq  *amqp.Channel
	queue     string
	mailer    MailService
}

func (s *service) Subscribe() error {
	// Consome mensagens da fila
	msgs, err := s.rabbitmq.Consume(
			s.queue,     // queue name
			"",         // consumer
			false,      // auto-ack
			false,      // exclusive
			false,      // no-local
			false,      // no-wait
			nil,        // args
	)
	if err != nil {
			return fmt.Errorf("erro ao consumir fila: %w", err)
	}

	// Processa mensagens em goroutine
	go func() {
			for msg := range msgs {
					// Deserializa evento
					var event RequestEvent
					if err := json.Unmarshal(msg.Body, &event) {
							msg.Nack(false, true)  // Rejeita e requeua
							continue
					}

					// Processa evento
					if err := s.ProcessEvent(event); err != nil {
							msg.Nack(false, true)
							continue
					}

					msg.Ack(false)  // Confirma processamento
			}
	}()

	return nil
}

func (s *service) ProcessEvent(event RequestEvent) error {
	switch event.Type {
	case "REQUEST_CREATED":
			return s.mailer.SendNewRequestEmail(event)
	case "REQUEST_APPROVED":
			return s.mailer.SendApprovalEmail(event)
	case "REQUEST_REJECTED":
			return s.mailer.SendRejectionEmail(event)
	}
	return nil
}