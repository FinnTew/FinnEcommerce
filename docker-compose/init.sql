CREATE TABLE `users`
(
    `id`                BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `email`             VARCHAR(255) NOT NULL UNIQUE COMMENT '用户邮箱',
    `password`          VARCHAR(255) NOT NULL COMMENT '密码（hash后）',
    `is_email_verified` TINYINT      NOT NULL DEFAULT 0 COMMENT '邮箱是否已验证(0=未验证,1=已验证)',

    -- 用户角色（粗粒度权限判断，例如 admin,user ）
    `role`              VARCHAR(50)  NOT NULL DEFAULT 'user' COMMENT '用户角色(如: user, admin)',

    -- 细粒度权限，存储 JSON 格式，配合正则匹配
    `permissions`       JSON                  DEFAULT NULL COMMENT '用户权限信息(用于正则匹配)',

    `created_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
    `updated_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `products`
(
    `id`          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name`        VARCHAR(255)   NOT NULL COMMENT '商品名称',
    `description` TEXT COMMENT '商品描述',
    `picture`     VARCHAR(1024)           DEFAULT NULL COMMENT '商品图片地址',
    `price`       DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '价格',
    `created_at`  DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `categories`
(
    `id`   BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL COMMENT '分类名称'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `product_categories`
(
    `id`          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `product_id`  BIGINT UNSIGNED NOT NULL,
    `category_id` BIGINT UNSIGNED NOT NULL,
    CONSTRAINT `fk_pc_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    CONSTRAINT `fk_pc_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `carts`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `user_id`    BIGINT UNSIGNED NOT NULL COMMENT '所属用户',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fk_cart_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cart_items`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `cart_id`    BIGINT UNSIGNED NOT NULL COMMENT '购物车 ID',
    `product_id` BIGINT UNSIGNED NOT NULL COMMENT '商品 ID',
    `quantity`   INT NOT NULL DEFAULT 1 COMMENT '商品数量',
    CONSTRAINT `fk_cart_items_cart` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`),
    CONSTRAINT `fk_cart_items_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `user_addresses`
(
    `id`             BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `user_id`        BIGINT UNSIGNED NOT NULL,
    `street_address` VARCHAR(255) NOT NULL,
    `city`           VARCHAR(100) NOT NULL,
    `state`          VARCHAR(100) NOT NULL,
    `country`        VARCHAR(100) NOT NULL,
    `zip_code`       VARCHAR(50)  NOT NULL,
    `created_at`     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fk_user_address_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `orders`
(
    `id`             BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '自增主键',
    `order_id`       VARCHAR(64)  NOT NULL UNIQUE COMMENT '订单唯一编号',
    `user_id`        BIGINT UNSIGNED NOT NULL COMMENT '下单用户',
    `user_currency`  VARCHAR(10)           DEFAULT 'CNY' COMMENT '下单时的货币类型',
    `street_address` VARCHAR(255) NOT NULL COMMENT '收货地址(快照)',
    `city`           VARCHAR(100) NOT NULL,
    `state`          VARCHAR(100) NOT NULL,
    `country`        VARCHAR(100) NOT NULL,
    `zip_code`       VARCHAR(50)  NOT NULL,
    `email`          VARCHAR(255) NOT NULL COMMENT '订单联系人邮箱',
    `status`         TINYINT      NOT NULL DEFAULT 0 COMMENT '订单状态 0待支付/1已支付/2已取消',
    `created_at`     INT          NOT NULL COMMENT '创建时间(秒级时间戳)',
    `updated_at`     INT          NOT NULL COMMENT '更新时间(秒级时间戳)',
    CONSTRAINT `fk_order_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `order_items`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `order_id`   VARCHAR(64)    NOT NULL COMMENT '关联到 orders 表的 order_id',
    `product_id` BIGINT UNSIGNED NOT NULL COMMENT '商品 ID',
    `quantity`   INT            NOT NULL COMMENT '购买数量',
    `cost`       DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品单价(或小计)',
    CONSTRAINT `fk_order_item_order` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `payments`
(
    `id`             BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `transaction_id` VARCHAR(64)    NOT NULL UNIQUE COMMENT '第三方支付流水号或自定义的交易号',
    `order_id`       VARCHAR(64)    NOT NULL COMMENT '关联订单号',
    `user_id`        BIGINT UNSIGNED NOT NULL COMMENT '支付用户',
    `amount`         DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '支付金额',
    `status`         TINYINT        NOT NULL DEFAULT 0 COMMENT '支付状态 0:成功 1:失败',
    `created_at`     DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `fk_payment_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    CONSTRAINT `fk_payment_order` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `short_links`
(
    `id`              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `short_link`      VARCHAR(255) NOT NULL UNIQUE COMMENT '短链接标识或slug',
    `long_link`       TEXT         NOT NULL COMMENT '原始长链接',
    `custom_slug`     VARCHAR(255)          DEFAULT NULL COMMENT '用户自定义的 slug',
    `expiration_time` DATETIME              DEFAULT NULL COMMENT '过期时间',
    `created_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;