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

type Log struct {
	Level string `mapstructure:"level"`
	File  struct {
		Enable bool   `mapstructure:"enable"`
		Path   string `mapstructure:"path"`
	} `mapstructure:"file"`
}

type Data struct {
	Pgsql struct {
		Enable   bool   `mapstructure:"enable"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Db       string `mapstructure:"db"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"pgsql"`

	Redis struct {
		Enable   bool   `mapstructure:"enable"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Db       int    `mapstructure:"db"`
	} `mapstructure:"redis"`

	Mqtt struct {
		Enable   bool   `mapstructure:"enable"`
		Broker   string `mapstructure:"broker"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"mqtt"`

	Rabbitmq struct {
		Enable   bool   `mapstructure:"enable"`
		Broker   string `mapstructure:"broker"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"rabbitmq"`
}

type Server struct {
	Gin struct {
		Enable bool   `mapstructure:"enable"`
		Host   string `mapstructure:"host"`
		Port   int    `mapstructure:"port"`
	} `mapstructure:"gin"`
}
