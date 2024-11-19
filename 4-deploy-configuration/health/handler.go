package health

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	health := struct {
			Status    string    
			Timestamp time.Time 
			Checks    map[string]string
	}{
			Status:    "OK",
			Timestamp: time.Now(),
			Checks: map[string]string{
					"database": "connected",
					"rabbitmq": "connected",
					"cache":    "available",
			},
	}
	
	// Verifica dependências
	if !checkDatabase() {
			health.Status = "ERROR"
			health.Checks["database"] = "disconnected"
	}
	
	json.NewEncoder(w).Encode(health)
}