package configuration

type Server struct {
	Host    string `yaml:"host"`
	Port    int16	`yaml:"port"`
	Gateway string `yaml:"gateway"`
}

type Telegram struct {
	ChatID	string	`yaml:"chat_id"`
	TokenEnv	string	`yaml:"token_env"`
	APIBaseURL	string `yaml:"api_base_url"`
	APISendMessage	string	`yaml:"api_send_message"`
}

type Conf struct {
	ConfigureFile string
	Server	Server	`yaml:"server"`
	Telegram	Telegram	`yaml:"telegram"`
}
