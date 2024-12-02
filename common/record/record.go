package record

import (
	"log"
	"net"
)

type Record struct {
	ips    []Addr
	addrs  []Addr
	punch  []Addr
	pubkey string
}
type Addr struct {
	addr     string
	port     int
	testPort int
}

func nslookup(addr Addr) []string {
	ips, err := net.LookupIP(addr.addr)
	var ret []string
	if err != nil {
		log.Fatal(err)
		return ret
	}
	for _, ip := range ips {
		ret = append(ret, ip.String())
	}
	return ret
}

func (r *Record) GetAddrs() []Addr {
	for _, addr := range r.addrs {
		ips := nslookup(addr)
		for _, ip := range ips {
			r.ips = append(
				r.ips,
				Addr{addr: ip, port: addr.port, testPort: addr.testPort},
			)
		}
	}
	r.ips = append(r.ips, r.punch...)
	return r.addrs
}
func (r Record) GetPubkey() string {
	return r.pubkey
}
func (a Addr) GetIP() net.IP {
	return net.ParseIP(a.addr)
}
func (a Addr) GetTestPort() int {
	return a.testPort
}
func (a Addr) GetPort() int {
	return a.port
}
