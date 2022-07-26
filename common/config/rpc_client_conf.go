package config

type RpcClientConfig struct {
	Endpoints []string  `json:"endpoints,optional" yaml:"endpoints"`
	Target    string    `json:"target,optional" yaml:"target"`
	App       string    `json:"app,optional" yaml:"app"`
	Token     string    `json:"token,optional" yaml:"token"`
	NonBlock  bool      `json:"non_block,optional" yaml:"non_block"`
	Timeout   int64     `json:"timeout,default=2000" yaml:"timeout" default:"2000"`
	Etcd      RpcClient `json:"etc,optional" yaml:"etcd"`
}

type RpcClient struct {
	Key                string   `json:"key" yaml:"key"`
	Driver             string   `mapstructure:"driver" json:"driver" yaml:"driver"`
	User               string   `json:"user" yaml:"user"`
	Pass               string   `json:"pass" yaml:"pass"`
	CertFile           string   `json:"cert_file,optional" yaml:"cert_file"`
	CertKeyFile        string   `json:"cert_key_file" yaml:"cert_key_file"`
	CACertFile         string   `json:"ca_cert_file" yaml:"ca_cert_file"`
	InsecureSkipVerify bool     `json:"insecure_skip_verify " yaml:"insecure_skip_verify"`
	Hosts              []string `json:"hosts" yaml:"hosts"`
}

type AppConfig struct {
	Port int `json:"port" yaml:"port"`
}
