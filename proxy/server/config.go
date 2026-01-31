package server

import (
	"github.com/catcursor/trojan-go/config"
	"github.com/catcursor/trojan-go/proxy/client"
)

func init() {
	config.RegisterConfigCreator(Name, func() interface{} {
		return new(client.Config)
	})
}
