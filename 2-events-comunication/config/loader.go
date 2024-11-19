package config

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	
	// Carrega configurações do arquivo
	viper.SetConfigName("config")
	viper.AddConfigPath("/app/config")
	if err := viper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("erro ao ler config: %w", err)
	}
	
	// Mapeia env vars
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	
	// Bind env vars para secrets
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("rabbitmq.uri", "RABBITMQ_URI")
	viper.BindEnv("jwt.secret", "JWT_SECRET")
	
	if err := viper.Unmarshal(cfg); err != nil {
			return nil, fmt.Errorf("erro unmarshal: %w", err)
	}
	
	return cfg, cfg.Validate()
}