### Generating Info-output

This report explains what the directory  "/" should contain and how to generate it. At the end I will show an example on how to do it in Golang and Python.

The "/" should contain
1. General info about how to use, expected JSON output.
2. Author
3. Service name and version
4. Small Description about the service.

Basically this directory docuements very apect of the Service.


#### Info-Output Generation in Golang

```
import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func check(err error) {
    if err != nil {
        panic(err.Error())
    }
}

type Metadata struct {
    Name        string
    Version     string
    Description string
    Copyright   string
    License     string
}

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

func main() {
    
metadata := &Metadata{
        Name:        "service-name",
        Version:     "0.1",
        Description: "service-description",
        Copyright:   "service-copyright",
        License:     "service-license",
    }

    router := httprouter.New()
    router.GET("/", info_output)

    err := http.ListenAndServe("127.0.0.1:8080", router)
    check(err)
}

```

#### Info-Output Generation in Golang

```
import tornado
from tornado import web, httpserver, ioloop


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


class Application(tornado.web.Application):
  def __init__(self, infoHandler):
    handlers = [
      (r"/", infoHandler),
    ]
    settings = dict(
        template_path=os.path.join(os.path.dirname(__file__), 'templates'),
        static_path=os.path.join(os.path.dirname(__file__), 'static'),
    )
    tornado.web.Application.__init__(self, handlers, **settings)
    self.engine = None


class Metadata(object):
    def __init__(self, name, version, description, copyright, license):
        self.name = name
        self.version = version
        self.description = description
        self.copyright = copyright
        self.license = license


m = Metadata(
  name="test-service",
  version="1.0",
  description="some fancy description",
  copyright="you can copy as much as you like",
  license="provided without any license"
)

infoHandler = InfoHandler(metadata=m)
server = tornado.httpserver.HTTPServer(Application())
server.listen(8080)
tornado.ioloop.IOLoop.instance().start()

```
