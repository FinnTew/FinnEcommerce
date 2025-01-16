package dal

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/auth/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
