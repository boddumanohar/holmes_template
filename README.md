# Holmes_template

Now lets discuss about the Generic Template. This template demonstrates how a service is created. This template is language independed.

In general, there are 2 kinds of services. 
1. File types
	These type of services does analysis on files and produces appropriate result.
2. Non File Types
	These type of services does analysis on non-file types like IPv4/6, Domain.

The only difference while create services occurs in step 2.1 and step 2.3

### Generic Template

#### 1. Create a Hello.lang ( **For File Types** )

All the services are RESTful applications. So create a webserver with "/info" and "/analyze" directories. 

1. `/info` directory displays general info and discription of the Servie on how to use it.( refer [Info-Output](./infooutput.md) )
2. `/analyze` directory analyses the raw data provided.And has to provide following functionalities
    1. Scan the URL and get the `raw_file` to be analysed from the filesystem or FTP ( refer [URL API Scheme](./urlapischeme.md) )
	2. Read the configuration file ( Refer [Read configuration scheme](./readconfiguration.md) )
	3. Send the `raw_file` to `analyser library` and parse output the result.
	4. Raise Appropriate HTTP Error Code ( refer [HTTP error codes](./httperrorcodes.md) )
	5. Fetch the result and fit into a JSON file

#### 2. Create a HelloREST.Scala
1. This one is the access point for the totem to interact with Services. This connects totem with services. This takes the result produced by the service to the totem.

##### How Totem Interacts with the services?
WORK IN PROGRESS 

#### 3. Running the entire server in a Docker or Virtual Machine.

In the section we shall discuss role of Docker and virutal machines in Microservices architechture and why we have choosen Docker over everything else.

##### Containerisation vs virtualisation?

WORK IN PROGRESS