##e Reading configuration file

Each Totem Service runs independently in an isolated docker container. The configuration settings for the Serivice has to be provided in `service.conf`. The configuration settings feeded into this file is used by the service logic for the working of the Service.

HolesProcessing uses JSON format for all the configuration files. This is because its easy for machines to parse. 

```json
{
	"HTTPBinding": ":8080", 
	"MaxNumberOfObjects": 10000 ,
}
```

Reading configuration is different for different languages. I will show how to read JSON and how to use this in Services in Golang and Python.
##### In Golang




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


Try opening the path, reading it all in and parsing it as json. If an error occures, throw a tornado.web.HTTPError (well define behaviour by tornado for these)
 If parsing succeeds, update provided config dictionary.

 ```python 

 def ParseConfig(config, path="service.conf", data=None):

    if not isinstance(config, dict):
        raise ValueError("Invalid parameter supplied to ParseConfig(config), given {}, but expects a dict".format(type(config)))

    if data is None:
        try:
            with open(path, "r") as file:
                try:
                    loaded_config = json.loads(file.read())
                except Exception as e:
                    raise HTTPError(500, "Error parsing config file: {}".format(e), reason="Bad Service Configuration")
        except Exception as e:
            raise HTTPError(500, "Error opening config file: {}".format(e), reason="Bad Service Configuration")
    else:
        try:
            loaded_config = json.loads(data)
        except Exception as e:
            raise HTTPError(500, "Error parsing config input: {}".format(e), reason="Bad Service Configuration")

    __updateDict(config, loaded_config)
    return StructDict(config)

def __toLower(key):
    if isinstance(key, str):
        return key.lower()
    return key

def __updateDict(old, new):
    keymap = {}
    for key in old:
        keymap[__toLower(key)] = key

    for key in new:
        _key = __toLower(key)
        if _key in keymap:
            ofrag = old[keymap[_key]]
            nfrag = new[key]

            if isinstance(ofrag, dict):
                if not isinstance(nfrag, dict):
                    raise ValueError("Mismatching config, expected dict, got: {}".format(type(nfrag)))
                __updateDict(ofrag, nfrag)

            else:
                # other entries are replaced if their types match
                if type(ofrag) != type(nfrag):
                    raise ValueError("Mismatching config, expected {}, got: {}".format(type(ofrag),type(nfrag)))
                old[keymap[_key]] = nfrag

class ZipMetaApp(tornado.web.Application):
	
	# ...
	# Application logic 
	# ...

Config = ServiceConfig("./service.conf")

def main():
    server = tornado.httpserver.HTTPServer(App())
    server.listen(Config.settings.port) # we use the configuration settings 
    tornado.ioloop.IOLoop.instance().start()

if __name__ == '__main__':
    main()
```
    