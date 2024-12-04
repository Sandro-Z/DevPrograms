package engine

type Key struct {
	MTU                      int    `yaml:"mtu"`
	Mark                     int    `yaml:"fwmark"`
	Proxy                    string `yaml:"proxy"`
	Device                   string `yaml:"device"`
	LogLevel                 string `yaml:"loglevel"`
	Interface                string `yaml:"interface"`
	TCPModerateReceiveBuffer bool   `yaml:"tcp-moderate-receive-buffer"`
	TCPSendBufferSize        string `yaml:"tcp-send-buffer-size"`
	TCPReceiveBufferSize     string `yaml:"tcp-receive-buffer-size"`
	EnableSocks5             bool   `yaml:"enable-socks5"`
	Socks5Addr               string `yaml:"socks5-addr"`
	Socks5Username           string `yaml:"socks5-username"`
	Socks5Password           string `yaml:"socks5-password"`
	HttpProxyAddr            string `yaml:"http-proxy-addr"`
	EnableHttpProxy          bool   `yaml:"enable-http"`
	NoGUI                    bool   `yaml:"no-gui"`
	Token                    string `yaml:"token"`
}
