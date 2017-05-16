### Should we keep the libary of Not?

My initial thoughts about the Totem-library were that: providing libraries would make the service code look lot cleaner. That’s the only reason that why I liked the library. But the problem with this is that we are forcing the user to use the API or Library. Also reading the code with the library included will be difficult because a user has to refer to docs everytime he has to deal an new method.

For example,

```
resultset = ServiceResultSet()
```

Every time a user wants to write a Service, the user has to first refer to the documentation of the library and has to remember this function( and many such functions ).Considering the above senarios, I am thinking not to provide with libraries for the create microservices. The user has freedom follow any style.


My idea is that to provide a **Generic Template** which is common for any programming language. The language uses its own style while sticking to template. And as an example, We will provide immplimentation of services in Golang and Python.


