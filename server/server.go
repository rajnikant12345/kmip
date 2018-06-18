package server

import (
	"fmt"
	"kmipserver/kmip"
	"net"
)

func ReadBuffer(c net.Conn) {

	for {
		b1 := make([]byte, 16)
		expected := 16
		tmp := b1
		for {
			n, e := c.Read(b1)
			if e != nil {
				fmt.Println(e.Error())
				c.Close()
				return
			}
			expected -= n
			if expected < 0 {
				c.Close()
				return
			}
			if expected != 0 {
				tmp = tmp[n:]
				continue
			}
			expected = kmip.StringToInt(string(b1[8:16])) * 2
			if expected > 30000 || expected <= 0 {
				c.Close()
				return
			}
			break
		}
		b2 := make([]byte, expected)
		tmp = b2
		for {
			n, e := c.Read(tmp)
			if e != nil {
				fmt.Println(e.Error())
				c.Close()
				return
			}
			expected -= n
			if expected != 0 {
				tmp = tmp[n:]
				continue
			}
			break
		}
		fmt.Println(string(append(b1,b2...)))

		ProcessPacket(c,append(b1,b2...))

	}
}

func ListenToTCPServer(ip, port string) {
	l, e := net.Listen("tcp", ip+":"+port)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	//signal handling code for closing connection

	for {
		c, e := l.Accept()
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		go ReadBuffer(c)
	}
}
