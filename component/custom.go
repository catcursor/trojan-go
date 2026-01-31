//go:build custom || full
// +build custom full

package build

import (
	_ "github.com/catcursor/trojan-go/proxy/custom"
	_ "github.com/catcursor/trojan-go/tunnel/adapter"
	_ "github.com/catcursor/trojan-go/tunnel/dokodemo"
	_ "github.com/catcursor/trojan-go/tunnel/freedom"
	_ "github.com/catcursor/trojan-go/tunnel/http"
	_ "github.com/catcursor/trojan-go/tunnel/mux"
	_ "github.com/catcursor/trojan-go/tunnel/router"
	_ "github.com/catcursor/trojan-go/tunnel/shadowsocks"
	_ "github.com/catcursor/trojan-go/tunnel/simplesocks"
	_ "github.com/catcursor/trojan-go/tunnel/socks"
	_ "github.com/catcursor/trojan-go/tunnel/tls"
	_ "github.com/catcursor/trojan-go/tunnel/tproxy"
	_ "github.com/catcursor/trojan-go/tunnel/transport"
	_ "github.com/catcursor/trojan-go/tunnel/trojan"
	_ "github.com/catcursor/trojan-go/tunnel/websocket"
)
