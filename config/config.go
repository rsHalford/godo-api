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
	DBUsername  string `env:"DB_USERNAME"`
	DBPassword  string `env:"DB_PASSWORD"`
	DBName      string `env:"DB_NAME"`
	APIUsername string `env:"API_USERNAME"`
	APIPassword string `env:"API_PASSWORD"`
}

var cfg Configuration

func GetString(key string) string {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		fmt.Println(err)
	}

	switch key {
	case "DB_NAME":
		value := cfg.DBName
		return value
	case "DB_USERNAME":
		value := cfg.DBUsername
		return value
	case "DB_PASSWORD":
		value := cfg.DBPassword
		return value
	case "API_USERNAME":
		value := cfg.APIUsername
		return value
	case "API_PASSWORD":
		value := cfg.APIPassword
		return value
	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}
