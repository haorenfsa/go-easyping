package easyping

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// Ping send icmp to target and measures delay
func Ping(addr string) (time.Duration, error) {
	c, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		return 0, err
	}
	defer c.Close()

	addrs, err := net.LookupHost(addr)
	if err != nil {
		return 0, err
	}
	if len(addrs) == 0 {
		return 0, fmt.Errorf("no ip address found for %s", addr)
	}
	addr = addrs[0]

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("Hi"),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		return 0, err
	}
	startTime := time.Now()
	if _, err := c.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP(addr)}); err != nil {
		return 0, err
	}

	rb := make([]byte, 1500)
	n, _, err := c.ReadFrom(rb)
	if err != nil {
		return 0, err
	}
	rm, err := icmp.ParseMessage(ProtoalICMP, rb[:n])
	if err != nil {
		return 0, err
	}
	duration := time.Since(startTime)
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		return duration, nil
	default:
		return 0, errors.New("received package type err")
	}
}

// ProtoalICMP see golang.org/x/net/internal/iana.ProtoalICMP
const ProtoalICMP = 1
