package vbmcgoalie

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// CLIOpts defines the CLI Arguements
type CLIOpts struct {
	Config string
}

// Config struct for start up config
type Config struct {
	VBMC ConfigYAML `yaml:"vbmc"`
}

// ConfigYAML is what is defined for this VBMC Goalie parser
type ConfigYAML struct {
	TargetDirectory string `yaml:"target_dir"`
	Hosts           []Host `yaml:"hosts"`
}

// Host is a host structure related to Libvirt
type Host struct {
	LibvirtDomainName string `yaml:"libvirt_domain_name"`
	LibvirtURI        string `yaml:"libvirt_uri"`
	IPMIUsername      string `yaml:"ipmi_username"`
	IPMIPassword      string `yaml:"ipmi_password"`
	IPMIPort          int    `yaml:"ipmi_port"`
}
