package shadowsocks

import (
	"net"

	"github.com/catcursor/trojan-go/tunnel"
)

type Conn struct {
	aeadConn net.Conn
	readBuf  []byte
	tunnel.Conn
}

func (c *Conn) Read(p []byte) (n int, err error) {
	if len(c.readBuf) > 0 {
		n = copy(p, c.readBuf)
		c.readBuf = c.readBuf[n:]
		return n, nil
	}
	return c.aeadConn.Read(p)
}

func (c *Conn) Write(p []byte) (n int, err error) {
	return c.aeadConn.Write(p)
}

func (c *Conn) Close() error {
	c.Conn.Close()
	return c.aeadConn.Close()
}

func (c *Conn) Metadata() *tunnel.Metadata {
	return c.Conn.Metadata()
}
