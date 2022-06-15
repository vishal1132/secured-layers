package main

import (
	"github.com/vishal1132/secured-layers/log"
	"github.com/vishal1132/secured-layers/security"
)

type User struct {
	Name  string
	Email string
}

func main() {
	u := User{
		Name:  "abcd",
		Email: "efgh",
	}
	s := security.New()
	s.RegisterFields(u, []string{"Email"})
	l := log.New(log.WithSecurity(s))
	l.Println(&u)
}
