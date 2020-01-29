package myerror

import "log"

// Check is a quick function designed to check if an error exists,
// and if so please unwrap while being printed to log.Fatalf()
func Check(e error) {
	if e != nil {
		log.Fatalf("Error Check: %w", e)
		// See about getting logging to a separate file by creating new logger.
		// Which could be useful in the future for the logging manager.
	}
}
