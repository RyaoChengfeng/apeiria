package config

var (
	// C 全局配置文件，在Init调用前为nil
	C *Config
)

// Config 配置
type Config struct {
	Bot     bot     `yaml:"bot"`
	LogConf logConf `yaml:"logConf"`
	Setu    setu    `yaml:"setu"`
	QQ      qq      `yaml:"qq"`
	Proxy   string  `yaml:"proxy"`
	Debug   bool    `yaml:"debug"`
}

type bot struct {
	Name    string `yaml:"name"`
	Addr    string `yaml:"addr"`
	BotPort string `yaml:"bot_port"`
	WsPort  string `yaml:"ws_port"`
}

type setu struct {
	API       string `yaml:"api"`
	R18       string `yaml:"r18"`
	Size      string `yaml:"size"`
	ImagePath string `yaml:"image_path"`
}

type qq struct {
	GroupList     []string `yaml:"group"`
	UserList      []string `yaml:"user"`
	SuperUserList []string `yaml:"super_user"`
}

type logConf struct {
	LogPath string `yaml:"log_path"`
	LogFile string `yaml:"log_file"`
}
