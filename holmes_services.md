#Holmes Services:
## Overview
Holmes-totem-Serivces is the part of HolmesProcessing that does analysis and creates a JSON file.

All the Planners of HolmesProcessing are Microservices, where Holmes Totem is a central orchestrator for all of these services. For performing analysis HolmesProcessing
uses Holmes-Totem-Services which are built on the ideas of microservices. These are pluggable services that we one can use or remove easily according to the needs.

Holmes Totem-Services are kind of webservers that does analysis using a particular library and returns a JSON file. The analyser library can be anything We can either use libary or directly parse the commandline output (as we need to do process call and parse the ouput. which is not recomended and time consuming).

Totem communicates with the services with are running a particular port via HTTP Protocol.

Each services runs in an isolated environment with the illusion that they are the only process running in the system. These are fine grained services and uses a light weight HTTP protocol for communication.

Each Totem-Services are atomic and does analysis individually. They don't depend on other Totem-Sevice to create the final output. Also there will not be any interprocess commnication between Totem-Services. The reason for this design choice is that this imporoves fault-tolerance and provides great flexibility and enforces modular structure.

Totem uses containers to isolate each Service.

we need to build a webserver, with an endpoints.

To communicate with Totem, Holmes Services uses external message queues
To communicate with services we use. external message queues.

##Communication Protocol
Totem Communicates with its services via the http protocol. The request for the service is of type GET with an empty body.

.. code-block:: html

    http://address:port/servicename/sampleid

The result returned by the service is an appropriate status code, paired with a
valid JSON reply body on success or a http error message reflecting the error
that occured. (plus optionally a JSON reply body)

.. code-block:: json

    {
      "key": "value",
      "key": {
        "subkey": "value",
        "subkey": "value"
      }
    }

This means that your service must act like a webserver. If the service isn't
written in a language that offers simple webserver capabilities (like golang),
then using a wsgi compatible python wrapper is adviced, like `Tornado`_.

.. _Tornado: http://www.tornadoweb.org/en/stable/

## Structure
To make satisfy all the required properties of the Services, Holmes Totem Services follow a structure so the each file has its own functionality.
1. Dockerfile -> to cantainersize and isolate the entire service
2. Scala File -> to communitate with totem.
3. Configuration File -> Configuration settings which can be changed as needed. These will be used by service Logic
4. Service Logic -> Webserver and Logic of the entire service.

### Containerisation


#### Why do we need containerisation

1. easier to boothstrap
2. When something gets messed up, we have have minimal penality.
3. 

Choosing the docker image for the microservices of your programming language. How to build docker images and what distributions to use. Choosing lightweight containers will make easier 
 As explained above, Holmes-Services are simple webservers that does analysis. We need to pack this entire system so that Services get an illusion they are the only Process runing. For our purpose, conceptually we need a VM. What kind of virtualisation do we need? why? We dont need an entire Operating system virtualisation. We only need a process isolation or better yet, we need protable process isolation. In docker you are starting a linux container. In case of Vagrant we are literally staring a new VM segmented from the host. We want this Docker because, if we have messed up something, and locked yourself out of the system, its no big deal. You can restart the container. If something Because we are dealing with malware samples, if something gets messed up, the penality should as mininal as possible. Also with docker we always have the same box.

we must be able to distribute the images to the applications. Easier to use and quick startup. If we want to make a change, we have to reprovision it. Where as with docker, you just can destroy and pull up a new container. 

Docker caching. That is: if you just chanage the last line, first lines are cached already. You dont have to reprovision everything. That feature is pretty powerfull We docker you need to understand some concepts of containers upfront: linked containers, volumes, Union file system.
Deyployment challenges. 

Choosing the right container for your lanaguage.
for Go: from alpine:go
for Python from alpine:python
For Rust FROM buildpack-deps:jessie
etc..
To isolate the Service, we will have to use virutalization to provide an isolated and constrained environment so that you dont have to worry about affecting other applications. and Each instance will be independent of each other. For this purpose we dont need a complete VM. We only want to isolate or analysis. 
For microservice heavy environments, Docker can be attractive because you can easily start a single Docker VM and start many containers above that very quickly

We choose docker because, 
Docker can be integrated into various infrastructure tools, including Amazon Web Services, Google Cloud Platform, Kubernetes, Microsoft Azure, OpenStack Nova, Vagrant,[35] and VMware vSphere Integrated Containers.

#### Things that you need to setup before staring with holmes Totem Services.
1. Docker
2. Go or Python or the language in which we are going to implement 
3. Suitable containers for the programming language. 

#### Base image.
`FROM` directive in the Dockerfile. To make images lean, we use [Apine linux](https://www.alpinelinux.org/about/), which is a security oriented lightweight Linux Distribution.
#### loading files
#### loading configuration files from the Storage URI

### Communication with TOTEM


### Configuration file.


### Service Logic

#### API Endpoints


### Output

#### HTTP Error Codes.
Holmes Services attempts to return appropriate HTTP Error Codes when a requested for analysis. This codes will be used by watchdog to manage results.

The Error codes table follow this pattern:

200 (OK) - Service returns analysis results in JSON format.
4XX (client errors) from General errors when processing 
5XX (server errors) from Service logic section

If the service returns an error, it returns so in the JSON format.

{
“Error”: {
“Message”: <Text>
“Code”: <number>
}
}
	
List of Errors codes 

Code
Text
Description
400
Bad request
Trying use new endpoint other than `/`and `/analyse`
401
Unauthorised
Invalid authentication
403
Forbidden
All the services by default should run on port 8080 so that docker compose can port forward it to a specified port.
404
Not Found
Invalid argument (the object provided is not found)
406
Not acceptable 
When there is a missing argument obj
420
Docker Error
Errors in Docker
421
Bad JSON file
Cannot Find or Parse JSON file 
500
Internal Server Error
When there is problem in the service logic like 
503
Cannot Create JSON
Due to some problem, the service cannot create output JSON file. 
505
HTTP Version Not
Supported. 
When a user doesn’t use HTTP 1.1 version 
(will change when (gRPC + HTTP 2) is introduced)

	
This error codes are will notify watchdog about the behaviour of the service and based on these error codes, watchdog can manage results more intelligently.

Like (sample proposal)( taken from issue #121)

4XX -> move to misbehaved queue.
5XX -> retask, restart the services

#### JSON output.



