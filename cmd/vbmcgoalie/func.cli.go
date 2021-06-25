package vbmcgoalie

import (
	"flag"
	"path/filepath"
)

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

// FileModeApplication will run the File Mode application logic bootstrap
func (config Config) StartApplication() {

	absoluteTargetDirectory, err := filepath.Abs(readConfig.VBMC.TargetDirectory)
	check(err)

	logStdOut(absoluteTargetDirectory)

	// Read in Zones file
	// zones, err := NewZones(config)
	// check(err)

}
