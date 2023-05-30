package config

type Credential struct {
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
	Enabled  bool   `yaml:"enabled"`
}
