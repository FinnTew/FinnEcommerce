package dal

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/payment/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
