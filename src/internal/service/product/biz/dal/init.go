package dal

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/product/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
