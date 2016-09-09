# goserveit
Go Lang server template

This is a sample template to build a server in Go using `net/http` package and
`gorilla mux` as muxer

## Getting Started

### Installing

Use Git to clone the repo:
```sh
$ git clone https://github.com/mukeshm/goserveit.git
```

This will clone the goserveit into your local filesystem

### Running the server

```sh
# switch to your local repo
cd goserveit
go run *.go
```
This will run a server on default port `8080`, so we can open `http://localhost:8080` in browser

### Adding extra functionalities

#### Adding handlers
Handlers for the new routes can be added in `handlers.go` file like :
```go
func myHandler(w http.ResponseWriter, r *http.Request) {
  //Do what ever you want
}
```
#### Adding routes
Once defining the handlers in `handlers.go` file, we have to map those handlers to routes. We can just add the mapping in `routes.go` like:
```go
var routes = Routes{
  Route{
    "/",
    http.MethodGet,
    indexHandler,
    "Index",
  },
  Route{
    "/hello",         //defining a pattern to match path
    http.MethodGet,   //for which HTTP method
    myHandler,        //handler function to call for the path if matched
    "HelloHandler",   //name - not used, just for clarification
  },
}
```

#### Changing port number
Change the default port number `8080` to whatever valid port number in `main` function in `server.go` 
