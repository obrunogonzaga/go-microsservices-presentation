package publisher

type Publisher interface {
	PublishEvent(ctx context.Context, req *Request) error
}

type service struct {
	rabbitmq  *amqp.Channel
	exchange  string
}

func (s *service) PublishEvent(ctx context.Context, req *Request) error {
	// Cria o evento
	event := RequestEvent{
			ID:        uuid.New().String(),
			Type:      "REQUEST_CREATED",
			RequestID: req.ID,
			Status:    req.Status,
			UserID:    ctx.Value("userID").(string),
			Timestamp: time.Now(),
	}

	// Serializa para JSON
	payload, err := json.Marshal(event)
	if err != nil {
			return fmt.Errorf("erro ao serializar evento: %w", err)
	}

	// Publica no RabbitMQ
	err = s.rabbitmq.Publish(
			s.exchange,   // exchange name
			"requests",   // routing key
			false,       // mandatory
			false,       // immediate
			amqp.Publishing{
					ContentType: "application/json",
					Body:       payload,
			},
	)

	if err != nil {
			return fmt.Errorf("erro ao publicar evento: %w", err)
	}

	return nil
}