package config

type Mysql struct {
	// 最大链接数，该值设置的越大，可以并发执行的数据库查询就越多
	MaxOpenConns int `yaml:"max_open_conns"`
	// 最大空闲链接数
	MaxIdleConns int `yaml:"max_idle_conns"`
	// 链接关闭前可以保持打开的最长时间
	MaxLifetime int `yaml:"max_life_time"`
	// 链接在关闭之前可用空闲的最长时间
	MaxIdleTime int    `yaml:"max_idle_time"`
	Charset     string `yaml:"charset"`
	Sources     []*Db  `yaml:"sources"`
	Replicas    []*Db  `yaml:"replicas"`
}

type Db struct {
	// Host     string `yaml:"host"`
	// Port     int    `yaml:"port"`
	Addr     string `yaml:"addr"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}
