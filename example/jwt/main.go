package main

import (
	"fmt"

	"github.com/marsxingzhi/goim/pkg/common/xzjwt"
)

func main() {
	token, _, err := xzjwt.GenerateAccessToken(1, 1, 3600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(token)
}
