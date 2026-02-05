package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"
	"strconv"
)

type Config struct {
	Server int `yaml:"service"`
}

func main(){
	filename,_ := filepath.Abs("config.yaml")
	yamlfile,err := ioutil.ReadFile(filename)

	if err!=nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlfile, &config)

	// create server for each host
	router := gin.Default()

	for i:=0;i<config.Server;i++ {
		fmt.Printf("create routing %s : v%s ", strconv.Itoa(i))

		
		route := router.Group("v"+strconv.Itoa(i))

		createRouter(route);
	}

	router.Run(":5000")
}