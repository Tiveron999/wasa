package api

import (
	"fmt"
	"strconv"
	"strings"
)

func idFromAuthorizationHeader(header string) (uint64, bool) {
	fmt.Println("Authorization header over here:", header)

	// Check if the Authorization header is not empty and starts with "Bearer "
	if header != "" && strings.HasPrefix(header, "Bearer ") {
		// Extract the token (user ID) from the Authorization header
		token := strings.TrimPrefix(header, "Bearer ")

		// Convert the token (string) to a uint64
		userID, err := strconv.ParseUint(token, 10, 64)
		if err != nil {
			// Return an error if the token cannot be parsed as a uint64
			return 0, false
		}

		// Return the user ID as a uint64
		return userID, true
	}

	// If the Authorization header is invalid or not present, return an empty string and false
	return 0, false
}
