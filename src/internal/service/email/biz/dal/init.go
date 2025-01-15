package dal

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/email/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
