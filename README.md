###Etcd

Is the vulcand implementation for herald.

[![GoDoc](https://godoc.org/github.com/duckbunny/vulcand?status.svg)](https://godoc.org/github.com/duckbunny/vulcand)


# vulcand
--
    import "github.com/duckbunny/vulcand"




## Usage

```go
var (
	KVpath string = "/vulcand/backends"
	TTL    int    = 15
	// Title for specifying herald in flags
	Title string = "vulcand"
)
```

#### func  BackendPath

```go
func BackendPath(s *service.Service) string
```

#### func  BasePath

```go
func BasePath(s *service.Service) string
```

#### func  Register

```go
func Register()
```
Register this herald with consul

#### func  ServerPath

```go
func ServerPath(s *service.Service) string
```

#### func  Url

```go
func Url(s *service.Service) string
```

#### type Backend

```go
type Backend struct {
	Type `json:"Type"`
}
```


#### type Server

```go
type Server struct {
	URL `json:"URL"`
}
```


#### type Vulcand

```go
type Vulcand struct {
	*etcd.Etcd
}
```


#### func (*Vulcand) Start

```go
func (v *Vulcand) Start(s *service.Service) error
```

#### func (*Vulcand) Stop

```go
func (v *Vulcand) Stop(s *service.Service) error
```
