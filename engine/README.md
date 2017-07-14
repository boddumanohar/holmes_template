# Template Generator

*Warning : Under the state of development*

## OverView

This project generates the boiler plate code required to create a Holmes Totem Service. This not a IDE which generates all the boilerplate code when a project with particular requirements are given.  


## Configuration 


## Installation


## Implimentation

This project creates a template engine which will have a set of templates. When a user specifies his requirements through commandline, this requirements are stored in a JSON configuration file. This configuration file is fed into template engine, Then the engine outputs the service boilerplate so the user only can concentrate on Service logic without having to worry about other things configurations (for now only service configuration, but later, gRPC settings etcc will come along). This will save time for the user.


This Engine is specially focussed towards microservices architechture. 

A typical template engine will be something like, it will take data and templates as input and return an HTML file. But here, we generate code. So more colloquially, this is a code generation script that is written in GO. 
