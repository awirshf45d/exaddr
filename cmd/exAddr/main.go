package main

import (
	"flag"
	"fmt"
	"sort"

	"net"
	"os"
	"regexp"
	"strings"
)

// Validates if a domain name is valid
func isValidDomain(domain string) bool {
	re := regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)
	return re.MatchString(domain)
}

// Validates if a host is valid
func isValidHost(host string) bool {
	return net.ParseIP(host) != nil || isValidDomain(host)
}

func main() {
	// Define flags
	filePath := flag.String("file", "", "Path to the input file")
	domains := flag.String("d", "", "Comma separated list of domains")
	output := flag.String("o", "cli", "Output method: 'cli' or file path")
	exIPs := flag.Bool("ip", false, "Use this flag to extract IPs.")
	// Parse flags
	flag.Parse()

	// Validate input flags
	if *filePath == "" {
		fmt.Println("Error: -file flag is required")
		os.Exit(1)
	}
	if *domains == "" && !*exIPs {
		fmt.Println("Error: -d flag is required if you want to extract domains.")
		os.Exit(1)
	}

	// Read input file
	content, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	var (
		hostList []string
		IPList   []string
	)
	if *exIPs {
		IPList = extractIPs(string(content))
		// Output results
		if *output == "cli" {
			for _, ip := range IPList {
				fmt.Println(ip)
			}
		} else {
			err := os.WriteFile(*output, []byte(strings.Join(IPList, "\n")), 0644)
			if err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
				os.Exit(1)
			}
		}
	} else {

		// Validate domains
		domainList := strings.Split(*domains, ",")
		for _, domain := range domainList {
			domain = strings.TrimSpace(domain)
			if !isValidDomain(domain) {
				fmt.Printf("Error: Invalid domain provided: %s\n", domain)
				os.Exit(1)
			}
		}

		// Extract and validate unique hosts
		hostList = extractHosts(string(content), domainList)

		// Output results
		if *output == "cli" {
			for _, host := range hostList {
				fmt.Println(host)
			}
		} else {
			err := os.WriteFile(*output, []byte(strings.Join(hostList, "\n")), 0644)
			if err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
				os.Exit(1)
			}
		}
	}

}

func extractHosts(content string, domains []string) []string {
	uniqueHosts := make(map[string]struct{})
	for _, domain := range domains {
		domain = strings.TrimSpace(domain)
		re := regexp.MustCompile(fmt.Sprintf(`[a-zA-Z0-9._-]+\.%s`, regexp.QuoteMeta(domain)))
		matches := re.FindAllString(content, -1)
		for _, match := range matches {
			if isValidHost(match) {
				uniqueHosts[match] = struct{}{}
			}
		}
	}
	var hosts []string
	for host := range uniqueHosts {
		hosts = append(hosts, host)
	}
	sort.Strings(hosts)
	return hosts
}

func extractIPs(conent string) []string {
	uniqueIPs := make(map[string]struct{})
	re := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	matches := re.FindAllString(conent, -1)
	for _, match := range matches {
		if _, exists := uniqueIPs[match]; exists == false {
			uniqueIPs[match] = struct{}{}
		}
	}
	var IPs []string
	for ip := range uniqueIPs {
		IPs = append(IPs, ip)
	}
	sort.Strings(IPs)
	return IPs
}
