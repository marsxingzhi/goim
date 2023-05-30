package main

import (
	"github.com/marsxingzhi/goim/apps/interfaces/internal/config"
	"github.com/marsxingzhi/goim/apps/interfaces/internal/server"
)

func main() {
	fakeDig()

	s := server.NewServer()
	s.Run()
}

// TODO-xz 优化
// 依赖注入，由于还未考虑使用哪种依赖注入，因此这里暂时手动注入
func fakeDig() {
	_ = config.GetConfig()
}
