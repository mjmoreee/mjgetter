package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

func main() {
	var banner = `
	███╗░░░███╗░░░░░██╗░██████╗░███████╗████████╗████████╗███████╗██████╗░
	████╗░████║░░░░░██║██╔════╝░██╔════╝╚══██╔══╝╚══██╔══╝██╔════╝██╔══██╗
	██╔████╔██║░░░░░██║██║░░██╗░█████╗░░░░░██║░░░░░░██║░░░█████╗░░██████╔╝
	██║╚██╔╝██║██╗░░██║██║░░╚██╗██╔══╝░░░░░██║░░░░░░██║░░░██╔══╝░░██╔══██╗
	██║░╚═╝░██║╚█████╔╝╚██████╔╝███████╗░░░██║░░░░░░██║░░░███████╗██║░░██║
	╚═╝░░░░░╚═╝░╚════╝░░╚═════╝░╚══════╝░░░╚═╝░░░░░░╚═╝░░░╚══════╝╚═╝░░╚═╝
	created with love by mj 💜
	`
	color.Magenta(banner)

	// url target
	url := flag.String("url", "unknown", "URL")

	// dns queries recon
	dnsLookup := flag.Bool("dl", false, "DNS Lookup")
	reverseDNS := flag.Bool("rd", false, "Reverse DNS")
	dnsHostRecord := flag.Bool("fd", false, "Find DNS Host Records")
	sharedDNSServers := flag.String("fs", "unknown", "Find Shared DNS Servers e.g: ns1.example.com")
	zoneTransfer := flag.Bool("zt", false, "Zone Transfer Online Test")

	// ip address recon
	ipGeoLocation := flag.String("gl", "unknown", "IP Geolocation Lookup")
	ipReverseIpLookup := flag.Bool("ir", false, "IP Reverse Lookup")
	asnLookup := flag.String("as", "unknown", "ASN Lookup")

	// web tools
	extractPageLink := flag.Bool("ep", false, "Extract Page Link")
	reverseAnalyticSearch := flag.Bool("rs", false, "Reverse Analytics Search")
	httpHeader := flag.Bool("ht", false, "HTTP Headers")

	flag.Parse()

	color.Green("Target Setup [ %s ]", *url)

	// dns queries recon
	if *dnsLookup {
		hackerTargetRecon("https://api.hackertarget.com/dnslookup/?q=", *url, "DNS Lookup")
	}
	if *reverseDNS {
		hackerTargetRecon("https://api.hackertarget.com/reversedns/?q=", *url, "Reverse DNS")
	}
	if *dnsHostRecord {
		hackerTargetRecon("https://api.hackertarget.com/hostsearch/?q=", *url, "Find DNS Host Record")
	}
	if *sharedDNSServers != "unknown" {
		hackerTargetRecon("https://api.hackertarget.com/findshareddns/?q=", *sharedDNSServers, "Find Shared DNS Servers")
	}
	if *zoneTransfer {
		hackerTargetRecon("https://api.hackertarget.com/zonetransfer/?q=", *url, "Zone Transfer")
	}

	// ip address recon
	if *ipGeoLocation != "unknown" {
		hackerTargetRecon("https://api.hackertarget.com/ipgeo/?q=", *ipGeoLocation, "IP Geolocation Lookup")
	}

	if *ipReverseIpLookup {
		hackerTargetRecon("https://api.hackertarget.com/reverseiplookup/?q=", *url, "IP Reverse Lookup")
	}

	if *asnLookup != "unknown" {
		hackerTargetRecon("https://api.hackertarget.com/aslookup/?q=", *asnLookup, "ASN Lookup")
	}

	// web tools
	if *extractPageLink {
		hackerTargetRecon("https://api.hackertarget.com/pagelinks/?q=", *url, "Extract Page Link")
	}

	if *reverseAnalyticSearch {
		hackerTargetRecon("https://api.hackertarget.com/analyticslookup/?q=", *url, "Reverse Analytics Search")
	}

	if *httpHeader {
		hackerTargetRecon("https://api.hackertarget.com/httpheaders/?q=", *url, "HTTP Headers")
	}
}

func hackerTargetRecon(url, victim, msg string) {
	color.Yellow("%s [ %s ]", msg, victim)

	urls := url + victim
	req, err := http.Get(urls)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	s := strings.Split(string(res), "\n")
	for i := 0; i < len(s); i++ {
		color.Cyan("🔎  %s", s[i])
	}
}
