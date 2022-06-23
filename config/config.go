package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type configuration struct {
	Port 		string `toml:"port"`
	Log			string `toml:"log_path"`
	ApiKey   	string `toml:"api_key"`
	Ethereum	Ethereum	`toml:"ethereum"`
	Database    Database `toml:"database"`
	RedisNetwork string `toml:"redis_network"`
	RedisPassword string `toml:"redis_password"`
	KeyFile string `toml:"keyFile"`
	CertFile string `toml:"certFile"`
	Power float32 `toml:"power"`
}
type Ethereum struct {
	Network string 	`toml:"network"`
	SptAddress string	`toml:"sptAddress"`
	ChainId	  int	`toml:"chainId"`
	KeyStore 	string	`toml:"KeyStore"`
	KeyStorePath string		`toml:"KeyStorePath"`
	Bnblp  string `toml:"bnblp"`
	//ContractAddress	string	`toml:"contractAddress"`
	KeyStorePathLP string `toml:"keyStorePathLP"`
}
type Database struct {
	DBHost       string  `toml:"db_host"`
	DBPort       string  `toml:"db_port"`
	DBSchema     string  `toml:"db_schema"`
	DBUserName 	 string  `toml:"db_username"`
	DBPassword   string  `toml:"db_password"`
	DBArgs       string  `toml:"db_args"`
}
var Cfg  *configuration


func Config() *configuration {
	if _, err := toml.DecodeFile("./config/config.toml", &Cfg); err != nil {
		log.Println(err)
	}
	return Cfg
}
