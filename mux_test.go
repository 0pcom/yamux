package yamux

import (
	"net"
	"testing"

	"golang.org/x/net/nettest"
)

func TestConn(t *testing.T) {
	nettest.TestConn(t, func() (c1, c2 net.Conn, stop func(), err error) {
		connC, connS := net.Pipe()

		c, err := Client(connC, DefaultConfig())
		if err != nil {
			return nil, nil, nil, err
		}

		s, err := Server(connS, DefaultConfig())
		if err != nil {
			return nil, nil, nil, err
		}

		if c1, err = c.OpenStream(); err != nil {
			return
		}

		if c2, err = s.AcceptStream(); err != nil {
			return
		}

		stop = func() {
			_ = c1.Close()
			_ = c2.Close()
			_ = c.Close()
			_ = s.Close()
			_ = connC.Close()
			_ = connS.Close()
		}

		return
	})
}
