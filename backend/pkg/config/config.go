package config

type Config struct {
	App    App    `mapstructure:"app"`
	Log    Log    `mapstructure:"log"`
	Data   Data   `mapstructure:"data"`
	Server Server `mapstructure:"server"`
}

type App struct {
	Version string   `mapstructure:"version"`
	Env     []string `mapstructure:"company"`
}

type File struct {
	Enable bool   `mapstructure:"enable"`
	Path   string `mapstructure:"path"`
}

type Log struct {
	Level string `mapstructure:"level"`
	File  File   `mapstructure:"file"`
}

type Pgsql struct {
	Enable   bool   `mapstructure:"enable"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       string `mapstructure:"db"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Redis struct {
	Enable   bool   `mapstructure:"enable"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type Mqtt struct {
	Enable   bool   `mapstructure:"enable"`
	Broker   string `mapstructure:"broker"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Rabbitmq struct {
	Enable   bool   `mapstructure:"enable"`
	Broker   string `mapstructure:"broker"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Data struct {
	Pgsql    Pgsql    `mapstructure:"pgsql"`
	Redis    Redis    `mapstructure:"redis"`
	Mqtt     Mqtt     `mapstructure:"mqtt"`
	Rabbitmq Rabbitmq `mapstructure:"rabbitmq"`
}

type Gin struct {
	Enable bool   `mapstructure:"enable"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
}

type Server struct {
	Gin Gin `mapstructure:"gin"`
}
