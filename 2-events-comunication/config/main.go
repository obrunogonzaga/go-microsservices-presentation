package main

func main() {
	// Carrega configurações
	cfg, err := config.LoadConfig()
	if err != nil {
			log.Fatalf("erro ao carregar config: %v", err)
	}

	// Inicializa componentes
	db := database.New(cfg.Database)
	broker := rabbitmq.New(cfg.RabbitMQ)
	auth := auth.New(cfg.JWT)
	
	// Inicia servidor
	server := server.New(cfg.Server)
	server.Run()
}