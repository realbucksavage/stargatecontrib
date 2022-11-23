# Eureka Service Lister

A `stargate.ServiceLister` based on [hudl/fargo](https://github.com/hudl/fargo) that queries the specified Eureka instances for registered applications.

```go
ls, err := eureka.NewLister("http://my-eureka-1/eureka", "http://my-eureka-2/eureka")
```

[Example](./_examples/main.go).
