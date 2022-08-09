package configuration

type ApplicationConfig struct {
	Server  ServerConfig  `yaml:"server"`
	MongoDb MongoDbConfig `yaml:"mongodb"`
	Log     LogConfig     `yaml:"log"`
}

type MongoDbConfig struct {
	ConnectionString string `yaml:"connectionString"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type LogConfig struct {
	Level         string `yaml:"level"`
	Address       string `yaml:"address"`
	Facility      string `yaml:"facility"`
	AppName       string `yaml:"appName"`
	Domain        string `yaml:"domain"`
	ActiveProfile string `yaml:"activeProfile"`
}
