document.getElementById("verifyButton").addEventListener("click", async () => {
    try {
        const email = document.getElementById("emailInput").value;

        if (!email) {
            alert("Please enter an email address.");
            return;
        }

        const response = await fetch(`http://localhost:8080/verify?email=${encodeURIComponent(email)}`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const result = await response.json();

        const resultDiv = document.getElementById("result");
        resultDiv.innerHTML = `
            <p>Email: ${result.email}</p>
            <p>Valid Syntax: ${result.validSyntax}</p>
            <p>Has MX Records: ${result.hasMX}</p>
            <p>Has SPF: ${result.hasSPF}</p>
            <p>SPF Record: ${result.spfRecord}</p>
            <p>Has DMARC: ${result.hasDMARC}</p>
            <p>DMARC Record: ${result.dmarcRecord}</p>
            <p>Valid SMTP: ${result.validSMTP}</p>
        `;
    } catch (error) {
        console.error("Error during fetch:", error);
        alert("Failed to verify email. Please check the console for details.");
    }
});

// Listen for Enter key press on the email input field
document.getElementById("emailInput").addEventListener("keydown", (event) => {
    if (event.key === "Enter") { // Check if the pressed key is Enter
        document.getElementById("verifyButton").click(); // Trigger the button click
    }
});