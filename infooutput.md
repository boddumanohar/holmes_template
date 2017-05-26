### How can I  create a SERVICE-INFO page?

Hopefully if you are aware of generating webpages, in Server side, this section becomes easier.
You can use Service-info page to view the information about the Service. Basically this page should state every aspect of Service the you are creating.

The INFO page should contain

1. Author name.
2. Service name and version. ( or any metadata about the service. )
3. Brief Description about the Service.
4. Licence 
5. General info about how to use the Service and expected JSON output.


#### Generating in Go.

```go
func info_output(f_response http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(f_response, `<p>%s - %s</p>
        <hr>
        <p>%s</p>
        <hr>
        <p>%s</p>
        `,
        metadata.Name,
        metadata.Version,
        metadata.Description,
        metadata.License)
}
```
We've just shown you how to parse and render templates. 

#### Info-Output Generation in Python

```python
class InfoHandler(tornado.web.RequestHandler):
        # Emits a string which describes the purpose of the analytics
        def get(self):
            info = """
	<p>{name:s} - {version:s}</p>
	<hr>
	<p>{description:s}</p>
	<hr>
	<p>{license:s}</p>
	<hr>
	<p>{copyright:s}</p>
            """.strip().format(
                name        = name,
                version     = version,
                description = description,
                license     = license,
                copyright   = copyright
            )
            self.write(info)
    return InfoHandler
```
