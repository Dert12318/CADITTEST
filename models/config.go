package models

type ServerConfig struct {
	Name           string `env:"NAME_SERVER"`
	Port           string `env:"PORT_SERVER,required"`
	Host           string `env:"HOST_SERVER,required"`
	JSONPathFile   string `env:"JSON_PATHFILE,required"`
	ElasticConfig  ElasticConfig
}


type ElasticConfig struct {
	Host     string `env:"HOST_ELASTICSEARCH,required"`
	Port     string `env:"PORT_ELASTICSEARCH,required"`
	User     string `env:"USER_ELASTICSEARCH"`
	Password string `env:"PASS_ELASTICSEARCH"`
	Index    string `env:"INDEX_ELASTICSEARCH,required"`
}
