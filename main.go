package main

import (
	"flag"
	"fmt"
	"github.com/miekg/dns"
	"os"
)

var dnsTypes = []uint16{
	dns.TypeA,
	dns.TypeAAAA,
	dns.TypeCNAME,
	dns.TypeMX,
	dns.TypeNS,
	dns.TypeSOA,
	dns.TypeTXT,
}

// List of common subdomains to query
var subdomains = []string{
	"www",
	"mail",
	"ftp",
	// add more if needed
}

func main() {
	domainPtr := flag.String("domain", "", "Domain name to create the zone file for")
	flag.Parse()

	if *domainPtr == "" {
		fmt.Println("Please provide a domain name with the --domain flag")
		os.Exit(1)
	}

	domain := *domainPtr
	f, _ := os.Create(domain + ".zone")
	defer f.Close()

	// Query for domain
	for _, dnsType := range dnsTypes {
		res, err := dnsQuery(domain, dnsType)
		if err != nil {
			continue
		}

		for _, rr := range res.Answer {
			f.WriteString(rr.String() + "\n")
		}
	}

	// Query for subdomains
	for _, subdomain := range subdomains {
		for _, dnsType := range dnsTypes {
			res, err := dnsQuery(fmt.Sprintf("%s.%s", subdomain, domain), dnsType)
			if err != nil {
				continue
			}

			for _, rr := range res.Answer {
				f.WriteString(rr.String() + "\n")
			}
		}
	}

	fmt.Printf("Zone file has been created: %s.zone\n", domain)
}

func dnsQuery(domain string, dnsType uint16) (*dns.Msg, error) {
	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(dns.Fqdn(domain), dnsType)
	r, _, err := c.Exchange(&m, "8.8.8.8:53")
	if err != nil {
		return nil, err
	}

	return r, nil
}

