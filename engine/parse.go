package main

import (
	"os"
	"encoding/json"
	// "path/filepath"
	"bytes"
	"fmt"
	"io"
	"flag"
	"log"
	"io/ioutil"
)

// Config

type Config struct {
	Name string
	Version string
	Description string
}

var config *Config

func main() {
	var (
		err error
		configPath string
	)

	flag.StringVar(&configPath, "config", "", "Path to the configuration file")
	flag.Parse()
	config, err = load_config(configPath)
	if err != nil {
		panic(err.Error())
	}

//	fmt.Println(config.Name)
	servicename := config.Name

	// create a new directory for the service.
	filenames := []string{"service.conf", "Dockerfile", "README.md", "serviceREST.scala", "service.go"}
	createDir(servicename)
	for i:=0; i<=4; i++ {
		dest := createFile(servicename, filenames[i])
		parseAndReplace(servicename, dest)
	}
}

func load_config(configPath string) (*Config, error) {
	config := &Config{}

	// if no path is supplied look in the current dir
	// if configPath == "" {
	// 	configPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	// 	configPath += "/parse.conf"
	// }

	cfile, _ := os.Open(configPath)
	if err := json.NewDecoder(cfile).Decode(&config); err != nil {
		return config, err
	}
	fmt.Println(config.Name)

	return config, nil
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createDir(service_name string) {
	name := service_name
	os.Mkdir(name,0700)
}

func createFile(service_name, filename string) string {
	src := "template" + "/" + filename + ".tpl"
	dest := service_name + "/" + filename
	sFile,err := os.Open(src)
	Check(err)

	eFile, err := os.Create(dest)
	Check(err)

	_, err = io.Copy(eFile, sFile)
	if err != nil {
		log.Fatal(err)
	}
	err = eFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

	return dest
}

func readFile(dest string) (input []byte, err error){
	input, err = ioutil.ReadFile(dest) // or anyfile
	if err != nil {
	  fmt.Println(err)
	  os.Exit(1)
   }
   return input, err
}

func parseAndReplace(servicename, dest string) {
// replace a word other name

	version := config.Version
	description := config.Description

	input, err := readFile(dest)
   output := bytes.Replace(input, []byte("{$name}"), []byte(servicename), -1)
   if err = ioutil.WriteFile(dest, output, 0666); err != nil {
			Check(err)
			 os.Exit(1)
	  }

	input,err = readFile(dest)
	output = bytes.Replace(input, []byte("{$version}"), []byte(version), -1)
   if err = ioutil.WriteFile(dest, output, 0666); err != nil {
			Check(err)
			 os.Exit(1)
	  }

	input,err = readFile(dest)
	output = bytes.Replace(input, []byte("{$description}"), []byte(description), -1)
   if err = ioutil.WriteFile(dest, output, 0666); err != nil {
			Check(err)
			 os.Exit(1)
	  }
 }
