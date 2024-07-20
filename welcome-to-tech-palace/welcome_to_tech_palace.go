package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var welcomeMsg string = "Welcome to the Tech Palace, "+ strings.ToUpper(customer)
	return welcomeMsg
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var asterisk string = strings.Repeat("*", numStarsPerLine)
	var message string = asterisk + "\n" + welcomeMsg + "\n" + asterisk
	return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var cleaned string = strings.TrimSpace(strings.Trim(oldMsg, "*\n\t"))
	return cleaned
}
