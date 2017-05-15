##e Reading configuration file

Each Totem Service runs independently in an isolated docker container. The service.conf file in each service is needed for the configuration of the service. The configuration settings feeded into this file is used by the service logic for the working of the service

Reading configuration is different for different languages. I will show how to read JSON and how to use this in Services in Golang and Python.
##### In Golang


the general format of the service.conf for Golang is just like any other JSON configuration file.

```json
{
	"HTTPBinding": ":8080", 
	"MaxNumberOfObjects": 10000 ,
}
```
Here is the sample configuration file for pdfparse service.

With the json package it's a snap to read JSON data into your Go programs. The json package provides Decoder and Encoder types to support the common operation of reading and writing streams of JSON data. We read the JSON file and then we fit the output to Config struct

```go
import (
"encoding/json" "flag" "os"
)

// ....

var config *Config var configPath string

// ....

type Config struct {
HTTPBinding string MaxNumberOfObjects int
}

// ....

flag.StringVar(&configPath, "config", "", "Path to the configuration file") flag.Parse()

config := &Config{}

cfile, _ := os.Open(configPath) dec := json.NewDecoder(cfile) // reading from json data

if err := dec.Decode(&config); err != nil {
// handle error

}

```
