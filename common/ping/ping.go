package ping

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

func Ping(ip net.IP, port int, timeout uint) error {
	// Ping the server with UDP

	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: ip, Port: port})
	if err != nil {
		log.Fatal(err)
	}
	defer socket.Close()

	socket.Write([]byte("ping"))

	ch := make(chan error)
	//defer close(ch)

	go func() {
		buffer := make([]byte, 1024)
		n, _, err := socket.ReadFromUDP(buffer)
		if err != nil {
			//log.Fatal(err)
			ch <- err
		}
		if bytes.Equal(buffer[0:n], []byte("pong")) {
			log.Default().Printf("Received pong from %s:%d", ip, port)
		}
		ch <- nil
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(time.Duration(timeout) * time.Second):
		return fmt.Errorf("no response")
	}

}

func Pong(port int) {
	// Listen for pings and respond with
	socket, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: port})
	if err != nil {
		log.Fatal(err)
	}
	defer socket.Close()
	for {
		buffer := make([]byte, 1024)
		n, addr, err := socket.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		if bytes.Equal(buffer[0:n], []byte("ping")) {
			socket.WriteToUDP([]byte("pong"), addr)
			log.Default().Printf("Sent pong to %s", addr)
		}
		time.Sleep(1 * time.Microsecond)
	}
}
