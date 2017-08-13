package model


type Config struct {
	Monitor struct {
		PoolSize	int			`yaml:"pool-size"`
	}

	DB struct {
		Id   		string 		`yaml:"id"`
		Password   	string 		`yaml:"password"`
		Name 		string  	`yaml:"name"`
		PoolSize	int			`yaml:"pool-size"`
	}
}
