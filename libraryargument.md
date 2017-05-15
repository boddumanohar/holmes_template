### Should we keep the libary of Not?

My initial thoughts about the library were that, providing libraries would make the service code look lot cleaner. That’s the only reason that why I liked the library. But the problem with this is that we are forcing the user to use the API or Library. Also reading the code with the library included in will be difficult. 

```
resultset = ServiceResultSet()
```

The user will have to remember this function every time he wants to write a Service. Every time the user wanted to create a service, he wants to first refer to the documentation of the library.

My idea is that to provide a **Generic Template** and as an example, provide to 2 examples in Golang and Python. 
Now lets discuss about the Generic Template.

##### Generic Template

###### 1. Create a Hello.lang ( **For File Types** )

All the services are RESTful applications. So create a webserver with "/" and "/analyze" directories. 

1. `/` directory displays general info and discription of the Servie on how to use it.( refer Info-Output )
2. `/analyze` directory analyses the raw data provided.
    1. Get the  get the object from the URL file from the OS (refer URL API Scheme)
	2. Read configuration file ( Refer Read configuration scheme )
	3. Parse the result by sending the file to `analyser library`.
	4. Raise Appropriate HTTP Error Code ( refer HTTP error codes )
	5. Fetch the result and fit into a JSON fil3

###### 2. Create a HelloREST.Scala
1. This one takes the result to the totem


