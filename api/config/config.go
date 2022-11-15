package config

import "os"

type Config struct {
	Rabbitmq Rabbitmq
}

type Rabbitmq struct {
	Exchange   string
	RoutingKey string
	User       string
	Password   string
	Host       string
	Port       string
}

func LoadConfig() Config {
	return Config{
		Rabbitmq: Rabbitmq{
			Exchange:   os.Getenv("RABBITMQ_EXCHANGE"),
			RoutingKey: os.Getenv("RABBITMQ_ROUTING_KEY"),
			User:       os.Getenv("RABBITMQ_USER"),
			Password:   os.Getenv("RABBITMQ_PASSWORD"),
			Host:       os.Getenv("RABBITMQ_HOST"),
			Port:       os.Getenv("RABBITMQ_PORT"),
		},
	}
}
