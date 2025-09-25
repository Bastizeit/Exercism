package techpalace

import "strings"

// WelcomePrefix is the welcome message prefix.
const WelcomePrefix = "Welcome to the Tech Palace, "

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return WelcomePrefix + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	lines := strings.Split(oldMsg, "\n")
	var contentLines []string

	for _, line := range lines {
		// Remove leading and trailing stars and whitespace
		cleaned := strings.Trim(line, "* ")
		// Only keep lines that have content after cleaning
		if cleaned != "" {
			contentLines = append(contentLines, cleaned)
		}
	}

	// Join the content lines and trim any remaining whitespace
	return strings.TrimSpace(strings.Join(contentLines, " "))
}
