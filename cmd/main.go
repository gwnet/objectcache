package cmd

import (
	"fmt"
	"path/filepath"
)

// Main main for minio server.
func Main(args []string) {
	// Set the minio app name.
	appName := filepath.Base(args[0])
	fmt.Println(appName)
}
