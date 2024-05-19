package domain

type Config struct {
	ServerAddr   string `env:"SERVER_ADDR,required"`
	LogFile      string `env:"LOG_FILE,required"`
	DatabaseAddr string `env:"MYSQL_ADDR,required"`
	DatabaseName string `env:"MYSQL_DATABASE,required"`
	DatabaseUser string `env:"MYSQL_USER,required"`
	DatabasePass string `env:"MYSQL_PASSWORD,required"`
}
