package config

func (c *Config) Validate() error {
	if c.Server.Port == 0 {
			return errors.New("porta do servidor é obrigatória")
	}
	
	if c.Database.Host == "" {
			return errors.New("host do banco é obrigatório")
	}
	
	if c.JWT.Secret == "" {
			return errors.New("JWT secret é obrigatório")
	}
	
	return nil
}