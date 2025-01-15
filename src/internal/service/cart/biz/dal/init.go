package dal

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/cart/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
