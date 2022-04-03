package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("domain, hasMX, hasSPF, sprRecord,hasDMARC,DmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, val := range txtRecords {
		if strings.HasPrefix(val, "v=spf1") {
			hasSPF = true
			spfRecord = val
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error : %v \n", err)
	}

	for _, val := range dmarcRecords {
		if strings.HasPrefix(val, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = val
			break
		}
	}

	fmt.Printf("%[1]v %[2]v %[3]v %[4]v %[5]v %[6]v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
