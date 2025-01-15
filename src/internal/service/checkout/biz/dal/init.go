package dal

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/checkout/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
