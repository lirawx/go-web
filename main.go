package main

import (
	"fmt"
	"strconv"

	"lirawx.cn/go-web/config"
	"lirawx.cn/go-web/routers"
)

// @title go-web API 文档
// @version 1.0
// @description API 文档
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	r := routers.SetupRouter()
	conf := config.ReadConfig()
	if err := r.Run(":" + strconv.Itoa(conf.Port)); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
