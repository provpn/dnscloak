package main

import (
	dnscryptproxy "github.com/jedisct1/dnscrypt-proxy/dnscrypt-proxy/ios"
	"log"
)

func main() {
	app := dnscryptproxy.Main("dnscrypt-proxy.toml")
	cb := &cloakCallback{}

	app.Run(cb)
}

type cloakCallback struct {
}

func (rcv cloakCallback) ProxyReady() {
	log.Println("cloak ready")
}
