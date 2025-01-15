package dal

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/order/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
