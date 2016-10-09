// inspired by: https://play.golang.org/p/m8TNTtygK0
// which was written by Russ Cox for golang-nuts google group:
// https://groups.google.com/forum/#!topic/golang-nuts/zlcYA4qk-94

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <cidr range> ...", os.Args[0])
		os.Exit(1)
	}
	for _, cidr := range os.Args[1:] {
		ip, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse CIDR %s: %s", cidr, err)
			os.Exit(1)
		}
		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
			fmt.Println(ip)
		}
	}
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
