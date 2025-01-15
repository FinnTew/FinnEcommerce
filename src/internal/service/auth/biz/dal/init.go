package dal

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/auth/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
