# Template Generator for Holmes Totem Services

## OverView

This program generates the boiler plate code required to create a Holmes Totem Service. Currently this is only for creating Static Services.

## Configuration 

Specify the required options for creating service in the configuration file`parse.conf`.

```json
{
	servicename : "helloworld"
	version : 1.0
}
```

## Installation
You need to have Go installed. Configure the required settings in `parse.conf` and just run `parse.go`file.

```
$ go run parse.go
```

## Implimentation

The parser takes the template (in the template folder) and configuration file as input and create a directory with the servicename and creates all the boilerplate code required. After this you can directly jump into servicelogic section of the {servicename}.go and add additional configuration settings and finally will add dependencies to the Dockerfile.

The created folder structure will be:
```
ServiceName
		|-----Service.{go,py}
		|-----Dockerfile
		|-----ServiceREST.Scala
		|-----Service.conf
```		

