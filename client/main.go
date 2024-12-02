package client

import (
	"flag"
	"time"

	"github.com/xanbert/dwg/client/wireguard"
	"github.com/xanbert/dwg/common/ping"
	"github.com/xanbert/dwg/common/record"
)

type conf struct {
	server      string
	port        int
	wgInterface string
	testPort    int
	timeout     uint
}

var (
	config  conf
	peers   map[string]wireguard.Peer
	records []record.Record
)

func Init() {
	flag.StringVar(&config.server, "server", "", "server address")
	flag.StringVar(&config.server, "s", "", "server address")
	flag.IntVar(&config.port, "port", 0, "server port")
	flag.IntVar(&config.port, "p", 0, "server port")
	flag.StringVar(&config.wgInterface, "interface", "wg0", "wireguard interface")
	flag.StringVar(&config.wgInterface, "i", "wg0", "wireguard interface")
	flag.IntVar(&config.testPort, "test", 12345, "port to test connectivity")
	flag.IntVar(&config.testPort, "t", 12345, "port to test connectivity")
	flag.UintVar(&config.timeout, "timeout", 1, "timeout for pings")
	flag.Parse()
}

func Main() {
	peers = wireguard.GetPeers(config.wgInterface)
	for {
		pushRecord()
		pullRecord()
		for _, record := range records {
			addrs := record.GetAddrs()
			pubkey := record.GetPubkey()
			_, ok := peers[pubkey]
			if !ok {
				break
			}
			for _, addr := range addrs {
				err := ping.Ping(addr.GetIP(), addr.GetTestPort(), config.timeout)
				if err == nil {
					peer := peers[pubkey]
					peer.SetAddr(addr)
					break
				}
			}
		}
		time.Sleep(30 * time.Second)
	}
}
