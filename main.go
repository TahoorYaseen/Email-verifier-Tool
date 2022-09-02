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

	fmt.Printf("Domain, hasMX, hasSPF, SpfRecord, hasDMARC, DMARCRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error : Could not read from input %v\n", err)
	}

}

func checkDomain(Domain string) {
	var hasDMARC, hasMX, hasSPF bool
	var SpfRecord, DMARCRecord string

	mxrecords, err := net.LookupMX(Domain)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	if len(mxrecords) > 0 {
		hasMX = true
	}

	txtrecords, err := net.LookupTXT(Domain)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, record := range txtrecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SpfRecord = record
			break
		}
	}

	dmarcrecords, err := net.LookupTXT("_dmarc." + Domain)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, record := range dmarcrecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			DMARCRecord = record
			break
		}
	}

	fmt.Printf("%v %v %v %v %v %v", Domain, hasMX, hasSPF, SpfRecord, hasDMARC, DMARCRecord)

}
