package utils

import (
	"log"
	"os"
)

// CheckAndCreateDataDirectory creates the data directory if it doesn't exist
func CheckAndCreateDataDirectory() {
	_, err := os.Stat("./data")
	if os.IsNotExist(err) {
		log.Println("Data directory does not exist, creating it...")

		os.MkdirAll("./data/uploads", 0755)
		return
	}

	_, err = os.Stat("./data/uploads")
	if os.IsNotExist(err) {
		log.Println("Uploads directory does not exist, creating it...")

		os.Mkdir("./data/uploads", 0755)
	}
}
