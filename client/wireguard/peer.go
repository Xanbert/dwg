package wireguard

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/xanbert/dwg/common/record"
)

type Peer struct {
	pubkey   string
	endpoint struct {
		ip   net.IP
		port int
	}
}

func GetPeers(iface string) map[string]Peer {
	var ret = make(map[string]Peer)
	list := wgGetPeers(iface)
	for _, peer := range list {
		ret[peer.pubkey] = peer
		log.Default().Output(2, fmt.Sprintf("peer: %v", peer.pubkey))
	}
	return ret
}

func (p *Peer) SetAddr(addr record.Addr) {
	p.endpoint.ip = addr.GetIP()
	p.endpoint.port = addr.GetPort()
	wgSetPeerEndpoint(p.pubkey, net.JoinHostPort(p.endpoint.ip.String(), strconv.Itoa(p.endpoint.port)))
}
