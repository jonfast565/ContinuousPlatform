package main

import "fmt"

type Configuration struct {
	Host      string
	Port      int
	Username  string
	DbName    string
	Password  string
	EnableSsl bool
}

func (c Configuration) GetPostgresConnectionString() string {
	enableSslResult := "enable"
	if c.EnableSsl == false {
		enableSslResult = "disable"
	}
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		c.Host,
		c.Port,
		c.Username,
		c.DbName,
		c.Password,
		enableSslResult)
}
