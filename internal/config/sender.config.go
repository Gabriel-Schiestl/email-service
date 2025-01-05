package config

type SenderConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewSenderConfig(host, username, password string, port int) *SenderConfig {
	return &SenderConfig{
		Host:     host,
		Username: username,
		Password: password,
		Port:     port,
	}
}