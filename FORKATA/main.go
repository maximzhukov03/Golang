package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type Config struct {
    Server Server `yaml:"server"`
    Db     Db     `yaml:"db"`
}

type Server struct {
    Port string `yaml:"port"`
}

type Db struct {
    Host     string `yaml:"host"`
    Port     string `yaml:"port"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
}

func getYAML(conf []Config) (string, error){
	encoder, err := yaml.Marshal(conf)
	if err != nil{
		return "", err
	}
	return string(encoder), nil
}

func main(){
    conf := []Config{
        {
            Server: Server{Port: "8080"},
            Db: Db{
                Host:     "localhost",
                Port:     "5432",
                User:     "admin",
                Password: "password123",
            },
        },
    }
	yamlData, err := getYAML(conf)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(yamlData)
}