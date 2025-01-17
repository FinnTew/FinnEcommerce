package dal

import "github.com/FinnTew/FinnEcommerce/src/internal/service/email/biz/dal/rabbitmq"

func Init() {
	//redis.Init()
	//mysql.Init()
	rabbitmq.Init()
}
