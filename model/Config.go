package model

type Config struct {
	Http struct {
		Port string `yaml:"port"`
	}
	Monitor struct {
		PoolSize int `yaml:"pool-size"`
	}

	DB struct {
		Ip       string `yaml:"ip"`
		Port     string `yaml:"port"`
		Id       string `yaml:"id"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		PoolSize int    `yaml:"pool-size"`
	}
}
