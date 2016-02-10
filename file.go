package snitchin

import (
	"os"
)

func File(name string) *os.File {
	file, error := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if error != nil {
		ERROR("Unable to open the file: " + name)
	}

	/* Return the file writer */
	return file
}
