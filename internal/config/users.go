package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetUser()([]User, error) {
	viper.SetConfigName("application")
	viper.SetConfigType("properties")
	viper.AddConfigPath("../cluster-portal")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Load users
	users := []User{}
	for i := 0; ; i++ {
		prefix := fmt.Sprintf("users.%d.", i)
		username := viper.GetString(prefix + "username")
		if username == "" {
			break
		}
		users = append(users, User{
			Username: username,
			Password: viper.GetString(prefix + "password"),
			Role:     viper.GetString(prefix + "role"),
		})
	}
	return users, nil
}