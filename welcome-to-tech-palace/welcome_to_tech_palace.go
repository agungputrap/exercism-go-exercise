package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var welcomeMessage string = ""
	// Upper border
	welcomeMessage += strings.Repeat("*", numStarsPerLine) + "\n"
	welcomeMessage += welcomeMsg + "\n"
	welcomeMessage += strings.Repeat("*", numStarsPerLine)

	return welcomeMessage
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var cleanUpMessage string = strings.Replace(oldMsg, "*", "", -1)
	cleanUpMessage = strings.TrimSpace(cleanUpMessage)

	return cleanUpMessage
}
