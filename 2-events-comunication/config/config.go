package config

type Config struct {
	Server struct {
			Port    int    
			Timeout time.Duration 
	}
	
	Database struct {
			Host     string
			Port     int
			Name     string
			User     string    // from secret
			Password string    // from secret
	}
	
	RabbitMQ struct {
			URI      string    // from secret
			Queue    string
			Exchange string
	}
	
	JWT struct {
			Secret     string  // from secret
			ExpiryTime int    
	}
}