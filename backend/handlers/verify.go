// Handlers for specific routes
package handlers

import (
	"net/http"
	"strings"
	"encoding/json"
	"github.com/email-verifier/utils"
)

func VerifyHandler(w http.ResponseWriter, r *http.Request) {

	
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK) // Return 200 OK for preflight
		return
	}

	// Get the email from the query string
	email := r.URL.Query().Get("email")

	// if email is not provided 
	if email == ""{
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Email Syntax Validation
	if !utils.IsValidEmail(email){
		http.Error(w, "Invalid Email Syntax", http.StatusBadRequest)
		return
	}

	// Split the email to extract the domain
	parts := strings.Split(email,"@")

	if len(parts) != 2 {
		http.Error(w, "Invalid Email Format", http.StatusBadRequest)
		return
	}

	// Domain is everything after "@"
	domain := parts[1]

	// Perform DNS Checks (for MX, SPF and DMARC records)
	hasMX, hasSPF, hasDMARC, spfRecord, dmarcRecord := utils.CheckDNSRecords(domain)

	// Perform SMTP handshake check to verify email existence
	validSMTP := utils.CheckSMTP(email, domain)

	// Create a response object with the results
	// Used a Map Literal with interface type values to store all types of data
	response := map[string]interface{}{
		"email":      email,
		"validSyntax": utils.IsValidEmail(email),
		"hasMX":      hasMX,
		"hasSPF":     hasSPF,
		"spfRecord":  spfRecord,
		"hasDMARC":   hasDMARC,
		"dmarcRecord": dmarcRecord,
		"validSMTP":  validSMTP,
	}

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// Send the response back as JSON
	json.NewEncoder(w).Encode(response)
}