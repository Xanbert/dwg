package ping

import (
	"log"
	"net"
	"testing"
	"time"
)

func TestPingPong(t *testing.T) {
	go Pong(12345)
	time.Sleep(1 * time.Second)
	err := Ping(net.IPv4(127, 0, 0, 1), 12345, 5)
	if err != nil {
		t.Errorf("Ping failed: %v", err)
		t.Fail()
	}
}

func TestPingNoPong(t *testing.T) {
	err := Ping(net.IPv4(127, 0, 0, 1), 12345, 5)
	if err == nil {
		t.Errorf("Ping should have failed")
		t.Fail()
	}
}

func FakePong(port int) {
	socket, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: port})
	if err != nil {
		panic(err)
	}
	defer socket.Close()
	for {
		time.Sleep(1 * time.Second)
	}
}

func TestPingTimeout(t *testing.T) {
	go FakePong(12345)
	time.Sleep(1 * time.Second)
	err := Ping(net.IPv4(127, 0, 0, 1), 12345, 5)
	log.Default().Printf("Ping returned %v", err)
	if err == nil || err.Error() != "no response" {
		t.Errorf("Ping should have timed out")
		t.Fail()
	}
}
