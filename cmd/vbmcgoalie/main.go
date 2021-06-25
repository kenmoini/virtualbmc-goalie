package vbmcgoalie

// Func main should be as small as possible and do as little as possible by convention
func main() {
	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := ParseFlags()
	checkAndFail(err)

	// Setup server config
	cfg, err := NewConfig(cfgPath)
	checkAndFail(err)

	// Run preflight
	PreflightSetup()

	// Run file mode application via CLI
	cfg.StartApplication()

}

// Bootstrap just allows for exporting of the package
func Bootstrap() {
	main()
}
