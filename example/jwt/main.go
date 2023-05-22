package main

import (
	"fmt"

	"github.com/marsingzhi/goim/pkg/common/xzjwt"
)

func main() {
	token, err := xzjwt.GenerateToken(1, 1, 3600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(token)
}
