### Reading configuration file


Before we go about reading configuration file, Lets discuss how Holmes being a distributed system, uses configuration files. Configuration files configure the parameters and intial settings for the Services.The settings of each of these components will be stored in a configuration file.

#### Service Configuration  for HolmesTotem

Each Service in HolmesTotem can be called as a Microservice. The essential settings needed for starting the Service are location in configuraion file. The general file format of configuration files is `.conf`. The format of the text is JSON. A user can change these values to change the behaviour of the Service. The configuration file is read by the service logic.

For each service, The service author should create a file called `service.conf` where the essential configuration settings for the service like what ports to listen etc... For example the configuration file( __service.conf__ ) for [gogadget](https://github.com/HolmesProcessing/Holmes-Totem/blob/master/src/main/scala/org/holmesprocessing/totem/services/gogadget) service is given below

```json
{
    "HTTPBinding": ":8080", 
    "MaxNumberOfObjects": 10000,
    "SearchDepth": 10
}
```

#### Centralised Service configuration of HolmesTotem

Holmes system allows an admin to store the configurations for the Services of Holmes-Totem in a central location and make the system automatically load it from there upon upstart. This is useful when you startup multiple Totem at different locations all working with the same service configuration because copying of all the Services everywhere is quite tedious Instead you only have to modify the config file on one machine and upload it and then rebuilt the containers on all the machines. It makes distributed service configuration changes easier.

Holmes-Storage was extended to allow for storing configuration-files (uploaded over HTTP) into its database and query them (also over HTTP). The [upload_config.sh](https://github.com/HolmesProcessing/Holmes-Totem/blob/master/config/upload_configs.sh) script can be used for uploading configuration to Storage. This script creates environment variables `CONFSTORAGE` which storage the uri where the configuration files are stored in Storage and configuration files are uploaded to storage. The [compose_download_conf.sh](https://github.com/HolmesProcessing/Holmes-Totem/blob/master/config/compose_download_conf.sh) script creates the environment variable for each Service like for example `CONFSTORAGE_ASNMETA` which contains the uri of the configuration file for the Service.


The docker-compose.yml.example therefore takes a look at environment variables pointing to the running instance of Holmes-Storage and sets these arguments correctly for each Service. By doing this, whenever the containers are built, the configuration-files are pulled from the server. As you can see in the docker-compose.yml.example, the Services always have a line like this:
```
conf: ${CONFSTORAGE_ASNMETA}service.conf
```

Usually the environment variable `CONFSTORAGE_ASNMETA` (and all the others as well) are empty, so the docker-files just get the local config version

Also The following are the modifications done for Dockerfile to accept an argument specifying the location of the file `service.conf`
```
# add the configuration file (possibly from a storage uri)
ARG conf=service.conf
ADD $conf /service/service.conf
```
If docker-compose did not set the conf-argument, it defaults to service.conf, otherwise it is left as it was. Docker's ADD can also download the file via HTTP.


#### Reading Configuration files for Services.

Each Service runs independently in an isolated docker container.The configuration settings for the Serivice has to be provided in `service.conf`. The configuration settings feeded into this file by the Service logic.

Some computer programs only read their configuration files at startup.
This configuration settings will be used by this service only. 

Reading configuration is just as easy as reading reason file. I will show how to read JSON and how to use this in Services
1. [Golang](##### Reading configuration in Golang)
2. [Python](##### Reading configuration in Python)

##### Reading configuration in Golang

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

##### Reading configuration in Python

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
