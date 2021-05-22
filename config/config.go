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
	Database string `yaml:"database"`
	API      string `yaml:"api"`
}

type Database struct {
	Username string `yaml:"username" env:"GOAPI_DATABASE_USERNAME"`
	Password string `yaml:"password" env:"GOAPI_DATABASE_PASSWORD"`
	Name     string `yaml:"api" env:"GOAPI_DATABASE_NAME"`
	Port     string `yaml:"editor" env:"GOAPI_DATABASE_PORT"`
}

type API struct {
	Username string `yaml:"username" env:"GOAPI_API_USERNAME"`
	Password string `yaml:"password" env:"GOAPI_API_PASSWORD"`
}

var (
	dataCfg Database
	apiCfg  API
	//cfgPath = "${XDG_CONFIG_HOME:-$HOME/.config}/godo/config.yaml"
	cfgPath = "./config.yaml"
)

func GetDatabaseString(key string) string {
	if err := cleanenv.ReadConfig(cfgPath, &dataCfg); err != nil {
		fmt.Println(err)
	}

	switch key {
	case "name":
		value := dataCfg.Name
		return value
	case "username":
		value := dataCfg.Username
		return value
	case "password":
		value := dataCfg.Password
		return value
	case "port":
		value := dataCfg.Port
		return value
	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}

func GetAPIString(key string) string {
	if err := cleanenv.ReadConfig(cfgPath, &apiCfg); err != nil {
		fmt.Println(err)
	}

	switch key {
	case "username":
		value := apiCfg.Username
		return value
	case "password":
		value := apiCfg.Password
		return value
	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}
