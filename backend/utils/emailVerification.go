// Utility Functions
package utils

import (
	"net"
	"regexp"
	"strings"
)

// Validate email syntax using a regular expression
func IsValidEmail(email string) bool {
	// Regular expression for a basic email pattern
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Check if the email matches the pattern
	return regexp.MustCompile(regex).MatchString(email)
}

// Check DNS Records

func CheckDNSRecords(domain string) (hasMX, hasSPF, hasDMARC bool, spfRecord, dmarcRecord string){
	// Look up MX(Mail Exchange) records
	mxRecords, _ := net.LookupMX(domain)

	if len(mxRecords) > 0{
		hasMX = true //  MX records are found
	}

	// Lookup TXT(text) records for SPF information
	txtRecords, _ := net.LookupTXT(domain)

	for _, record := range txtRecords{
		if strings.HasPrefix(record,"v=spf1"){
			hasSPF = true	// SPF records are found
			spfRecord = record
			break
		}
	}

	// Lookup TXT records for DMARC information
	dmarcRecords, _ := net.LookupTXT("_dmarc." + domain)

	for _, record := range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true     //DMARC records are found
			dmarcRecord = record
			break
		}
	}

	// in this function named return is used
	return

}