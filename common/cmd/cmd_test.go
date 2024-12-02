package cmd

import (
	"log"
	"testing"
)

func TestWgShowconf(t *testing.T) {
	strings, err := Run("wg", "showconf", "wg2")
	log.Default().Output(2, strings)
	if err != nil {
		log.Fatal(err)
		t.Error(err)
	}
}
func TestIpAddr(t *testing.T) {
	strings, err := Run("ip", "addr")
	log.Default().Output(2, strings)
	if err != nil {
		log.Fatal(err)
		t.Error(err)
	}
}
