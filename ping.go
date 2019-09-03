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
	opt := &Options{
		Address: addr,
		Timeout: DefaultTimeOut,
	}
	result, err := AdvancedPing(opt)
	if err != nil {
		return 0, err
	}
	return result.RoundTimeDelay, nil
}

// Options for advanced ping
type Options struct {
	Address string
	Timeout time.Duration
}

// Result is the result of advanced ping
type Result struct {
	RoundTimeDelay time.Duration
}

// AdvancedPing pings with options
func AdvancedPing(opt *Options) (*Result, error) {
	c, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		return nil, err
	}
	defer c.Close()

	addrs, err := net.LookupHost(opt.Address)
	if err != nil {
		return nil, err
	}
	if len(addrs) == 0 {
		return nil, fmt.Errorf("no ip address found for %s", opt.Address)
	}
	addr := addrs[0]

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("Hi"),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		return nil, err
	}

	startTime := time.Now()
	deadline := startTime.Add(DefaultTimeOut)
	c.SetDeadline(deadline)
	if _, err := c.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP(addr)}); err != nil {
		return nil, err
	}

	rb := make([]byte, 1500)
	n, _, err := c.ReadFrom(rb)
	if err != nil {
		return nil, err
	}
	rm, err := icmp.ParseMessage(ProtoalICMP, rb[:n])
	if err != nil {
		return nil, err
	}

	duration := time.Since(startTime)
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		return &Result{RoundTimeDelay: duration}, nil
	default:
		return nil, errors.New("received package type err")
	}
}

// ProtoalICMP see golang.org/x/net/internal/iana.ProtoalICMP
const ProtoalICMP = 1

// DefaultTimeOut 1s
const DefaultTimeOut = time.Second
