package wireguard

import (
	"log"
	"strings"

	"github.com/spf13/viper"
	"github.com/xanbert/dwg/common/cmd"
)

func wgSetPeerEndpoint(pubkey string, endpoint string) {
	msg, err := cmd.Run("wg", "set", "peer", pubkey, "endpoint", endpoint)
	if err != nil {
		log.Fatal(err)
	}
	log.Default().Output(2, msg)
}

func parsePeers(conf string) []Peer {
	var ret []Peer
	v := viper.New()
	v.SetConfigType("ini")
	v.ReadConfig(strings.NewReader(conf))

	return ret
}

func wgGetPeers(iface string) []Peer {
	conf, err := cmd.Run("wg", "showconf", iface)
	if err != nil {
		log.Fatalf("wg showconf: %v", err)
		return nil
	}
	return parsePeers(conf)
}
