package config

type Config struct {
	Port         int    `env:"RESTSERVER_PORT" envDefault:"8000"`
	Env          string `env:"RESTSERVER_ENV" envDefault:"TEST"`
	Hosts        string `env:"RESTSERVER_HOSTS" envDefault:"127.0.0.1:"`
	WriteTimeout int    `env:"RESTSERVER_WRITETIMEOUT" envDefault:"15"`
	ReadTimeout  int    `env:"RESTSERVER_WRITETIMEOUT" envDefault:"15"`
}
