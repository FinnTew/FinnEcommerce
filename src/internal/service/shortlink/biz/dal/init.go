package dal

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/shortlink/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/shortlink/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
