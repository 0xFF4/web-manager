package functions

import (
	"math/rand"
	"time"
)

// MakeRandom generates a random string of a specified length.
func MakeRandom(length int) string {

	// Initialize a new random number generator using a seed derived from the current time.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define the character set to use for the random string.
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a slice to store the individual characters of the password.
	s := make([]byte, length)

	// Loop to generate the required number of characters.
	for i := 0; i < length; i++ {
		// Pick a random character from the charset and add it to the password slice.
		s[i] = charset[r.Intn(len(charset))]
	}

	// Return the final password as a string.
	return string(s)
}
