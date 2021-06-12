/*
Copyright © 2021 Richard Halford <richard@xhalford.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	DBUsername  string `yaml:"db_username" env:"GOAPI_DB_USERNAME"`
	DBPassword  string `yaml:"db_password" env:"GOAPI_DB_PASSWORD"`
	Name        string `yaml:"name" env:"GOAPI_DB_NAME"`
	Port        string `yaml:"port" env:"GOAPI_DB_PORT"`
	APIUsername string `yaml:"api_username" env:"GOAPI_API_USERNAME"`
	APIPassword string `yaml:"api_password" env:"GOAPI_API_PASSWORD"`
}

var (
	cfg Configuration
	//cfgPath = "${XDG_CONFIG_HOME:-$HOME/.config}/godo/config.yaml"
	cfgPath = "./config.yaml"
)

func GetString(key string) string {
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		fmt.Println(err)
	}

	switch key {
	case "name":
		value := cfg.Name
		return value
	case "db_username":
		value := cfg.DBUsername
		return value
	case "db_password":
		value := cfg.DBPassword
		return value
	case "port":
		value := cfg.Port
		return value
	case "api_username":
		value := cfg.APIUsername
		return value
	case "api_password":
		value := cfg.APIPassword
		return value
	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}
