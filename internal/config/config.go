package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)


type User struct {
    Username string
    Password string
    Role     string
} 

type HTTPServer struct {
	Addr string `yaml:"address"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	//StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer   `yaml:"http_server"`
	DbHost string `yaml:"db_host" env-required:"true"`
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password"`
	Dbname string `yaml:"dbname" env-required:"true"`
	Users      []User
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath =="" {
		flagPath := flag.String("config", "", "Path to the configuration file")
		flag.Parse()
		configPath = *flagPath
		fmt.Println("configPath: ", configPath)
		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}
	if _,err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Can not read config file: %s", err.Error())
	}

	//users := []User{{"readonly", "readonly", "readonly"}, {"admin", "admin", "admin"}}
    // users, err := GetUser()
	// if err != nil {
	// 	log.Fatalf("Can not read config file: %s", err.Error())
	// }
    //cfg.Users = users
	fmt.Println("Users: ",cfg.Users)
	return &cfg
}