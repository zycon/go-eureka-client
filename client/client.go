package main

import (
	"flag"
	"github.com/ArthurHlt/go-eureka-client/eureka"
)

func main() {

	//hostName, app, ip string, port int, ttl uint, isSsl bool

	host := flag.String("hostname", "localhost", "HostName")
	app := flag.String("app", "nashwan-app", "App Name")
	ip := flag.String("ip", "127.0.0.1", "IP Address")
	port := flag.Int("port", 8761, "Port Number")
	ttl := flag.Uint("ttl", 10000, "TTL ")
	ssl := flag.Bool("ssl", false, "SSL Connection")
	eurekaUrl := flag.String("sd", "http://127.0.0.1:8761/eureka", "Eureka Server URL")
	flag.Parse()

	client := eureka.NewClient([]string{
		*eurekaUrl,
	})
	instance := eureka.NewInstanceInfo(*host, *app, *ip, *port, *ttl, *ssl)
	instance.Metadata = &eureka.MetaData{
		Map: make(map[string]string),
	}
	instance.Metadata.Map["foo"] = "bar"
	client.RegisterInstance("myapp", instance)
	client.GetApplication(instance.App)
	client.GetInstance(instance.App, instance.HostName)
	client.SendHeartbeat(instance.App, instance.HostName)

}
