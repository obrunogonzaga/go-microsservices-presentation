package domain

type RequestEvent struct {
	ID        string      // Identificador único do evento
	Type      string      // Tipo do evento
	RequestID string      // ID da solicitação
	Status    string      // Status atual
	Metadata  Metadata    // Dados adicionais
	Timestamp time.Time   // Momento do evento
}

type Metadata struct {
	UserID    string
	Action    string
	Details   map[string]interface{}
}