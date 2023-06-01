package config

type Grpc struct {
	Name                  string      `yaml:"name"`
	ServerID              int         `yaml:"server_id"`
	Port                  int         `yaml:"port"`
	MaxConnectionIdle     int         `yaml:"max_connection_idle"`
	MaxConnectionAge      int         `yaml:"max_connection_age"`
	MaxConnectionAgeGrace int         `yaml:"max_connection_age_grace"`
	Time                  int         `yaml:"time"`
	Timeout               int         `yaml:"timeout"`
	ConnectionLimit       int         `yaml:"connection_limit"`
	StreamsLimit          uint32      `yaml:"streams_limit"`
	MaxRecvMsgSize        int         `yaml:"max_recv_msg_size"`
	Credential            *Credential `yaml:"credential"`
}

type GrpcDialOption struct {
	ServiceName string
	Etcd        *Etcd
}

type GrpcServer struct {
	Name string `yaml:"name"`
	//Cert *Cert  TODO 证书先不加了
}
