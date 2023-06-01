package config

type WsServer struct {
	ServerID             int    `yaml:"server_id" json:"server_id"`
	Name                 string `yaml:"name" json:"name"`
	Port                 int    `yaml:"port" json:"port"`
	WriteWait            int    `yaml:"write_wait" json:"write_wait"`
	PongWait             int    `yaml:"pong_wait" json:"pong_wait"`
	PingPeriod           int    `yaml:"ping_period" json:"ping_period"`
	MaxMsgSize           int    `yaml:"max_msg_size" json:"max_msg_size"`
	ReadBufferSize       int    `yaml:"read_buffer_size" json:"read_buffer_size"`
	WriteBufferSize      int    `yaml:"write_buffer_size" json:"write_buffer_size"`
	HeaderLength         int    `yaml:"header_length" json:"header_length"`
	ChanClientSendMsg    int    `yaml:"chan_client_send_msg" json:"chan_client_send_msg"`
	ChanServerReadMsg    int    `yaml:"chan_server_read_msg" json:"chan_server_read_msg"`
	ChanServerRegister   int    `yaml:"chan_server_register" json:"chan_server_register"`
	ChanServerUnRegister int    `yaml:"chan_server_unregister" json:"chan_server_unregister"`
	MaxConnections       int    `yaml:"max_connections" json:"max_connections"`
	MinTimeInterval      int    `yaml:"min_time_interval" json:"min_time_interval"`
}
