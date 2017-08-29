package main

import (
	"github.com/docker/libcontainer/netlink"
	"github.com/miekg/dns"
	"fmt"
	"flag"
	"log"
	"os"
)

func dump() {
	if rts, err := netlink.NetworkGetRoutes(); err == nil {
		for _, r := range rts {
			if r.IPNet != nil {
				fmt.Println(r.Default, r.IP, r.Iface.Name)
			}
		}
	}
}

func main() {

	dnsServer := flag.String("dns", "8.8.8.8", "DNS server to use")

	ifacePtr := flag.String("iface", "", "interface to add the route to")

	flag.Parse()

	hosts := flag.Args()

	if ifacePtr == nil || *ifacePtr == "" || len(hosts) == 0 {
		fmt.Println("Usage: -iface <interface> domain1 domain2 domain3 domain4 ...")
		flag.PrintDefaults()
		os.Exit(1)
	}

	c := new(dns.Client)
	c.Net = "udp"

	m := &dns.Msg{
		MsgHdr: dns.MsgHdr{
			RecursionDesired: true,
			Opcode:           dns.OpcodeQuery,
		},
	}

	for _, host := range hosts {
		m.SetQuestion(dns.Fqdn(host), dns.TypeA)
		if r, _, err := c.Exchange(m, *dnsServer+":53"); err == nil {
			for _, resp := range r.Answer {
				switch aa := resp.(type) {
				case *dns.A:
					if err := netlink.AddRoute(aa.A.String()+"/32", "", "", *ifacePtr); err != nil {
						log.Fatal(err)
					}
				}

			}
		} else {
			log.Fatal(err)
		}
	}

}
