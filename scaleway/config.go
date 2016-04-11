package scaleway

type Config struct {
	Organization string `yaml:"organization",json:"organization"`
	Token        string `yaml:"token",json:"token"`
	UserAgent    string `yaml:"user_agent",json:"user_agent"`
}
