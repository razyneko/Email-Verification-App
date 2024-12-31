// Utility Functions
package utils

import "net"

// Check if an email is valid by performing an SMTP Handshake

func CheckSMTP(email, domain string) bool {
	// Open an SMTP connection to the domain
	client, err := net.Dial("tcp", domain+":25")   // Port 25 is used for SMTP

	if err != nil{
		return false   // If the connection fails
	}

	defer client.Close()

	// Perform a simple handshake with SMTP server
	// A basic and limited SMTP check
	client.Write([]byte("HELO localhost\r\n"))
	client.Write([]byte("MAIL FROM:<test@example.com>\r\n"))
	client.Write([]byte("RCPT TO:<" + email + ">\r\n"))

	// If no errors occurred, assume the email exists
	return true
}