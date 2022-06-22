package types

var (
	HttpConfigName = "config/http.yml"
)

type Authorization struct {
    Login string `yaml:"login"`
    Password string `yaml:"password"`
}

type HttpCheck struct {
    Name string `yaml:"name"`
    Path string `yaml:"path"`
    Method string `yaml:"method"`
    Response_code int `yaml:"response_code"`
    Body string `yaml:"body"`
    Response_body string `yaml:"response_body"`
    Token string `yaml:"token"`
	Auth Authorization `yaml:"auth"`
	Headers map[string]string `yaml:"headers"`
    Playload string `yaml:"playload"`
    Base_url string `yaml:"base_url"`
}

type HttpConfig struct {
    Base_url string `yaml:"base_url"`
    Checks []*HttpCheck
}