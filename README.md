# FinnEcommerce(Developing)

## Description

字节跳动青训营项目，抖音电商项目。

## TechStack

`Golang`, `Kitex`, `Hertz`, `Consul`, `MySQL`, `Redis`, `RabbitMQ`, `ElasticSearch`

## Start

1. 环境准备
   - Docker 拉取运行 `RabbitMQ` 服务
     ```bash
     cd ./docker && sudo docker-compose run rabbitmq
     ```
     
   - 下载 `RabbitMQ` 延时队列插件
     ```bash 
     curl -o ~/rabbitmq/plugins/rabbitmq_delayed_message_exchange-3.13.0.ez https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/v3.13.0/rabbitmq_delayed_message_exchange-3.13.0.ez
     ```
     
   - Copy 到容器插件目录下
     ```bash
     sudo docker cp ~/rabbitmq/plugins/rabbitmq_delayed_message_exchange-3.13.0.ez rabbitmq-finn:/plugins
     ```
   
   - 启动插件
     ```bash
     sudo docker exec -it rabbitmq-finn rabbitmq-plugins enable rabbitmq_delayed_message_exchange
     ```
     
   - restart
     ```bash
     sudo docker restart rabbitmq-finn
     ```
     
   - 启动其他服务
     ```bash
     sudo docker-compose up -d
     ```
     
2. 编译运行
   (待补充)