package service

import "fmt"

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service

type People interface {
	Say() string
}

func SayAnything(p People) {
	fmt.Printf("say: %s\n", p.Say())
}

type DB interface {
	GetName() string
}

type Server interface {
	Get() DB
}

func GetDB(srv Server) DB {
	return srv.Get()
}
