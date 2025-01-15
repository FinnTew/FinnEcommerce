package dal

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
