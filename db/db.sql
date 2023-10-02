
use fat_bear;
DROP TABLE IF EXISTS `fat_bear.t_rooms`;
DROP TABLE IF EXISTS `fat_bear.t_users`;

CREATE TABLE `fat_bear.t_rooms`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `room_name`   varchar(255)    NOT NULL DEFAULT '' COMMENT 'room_name',
    `room_status` int             NOT NULL DEFAULT '0' COMMENT 'room_status',
    `is_deleted`  TINYINT         NOT NULL DEFAULT 0 COMMENT 'is_deleted',
    `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 4
  DEFAULT CHARSET = UTF8MB4;

CREATE TABLE `fat_bear.t_users`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_name`   varchar(255)    NOT NULL DEFAULT '' COMMENT 'room_name',
    `user_status` int             NOT NULL DEFAULT '0' COMMENT 'room_status',
    `is_deleted`  TINYINT         NOT NULL DEFAULT 0 COMMENT 'is_deleted',
    `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 4
  DEFAULT CHARSET = UTF8MB4;

CREATE TABLE `fat_bear.t_devs`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `dev_name`    varchar(255)    NOT NULL DEFAULT '' COMMENT 'dev_name',
    `cmd`         varchar(64)             NOT NULL DEFAULT '' COMMENT 'cmd',
    `is_deleted`  TINYINT         NOT NULL DEFAULT 0 COMMENT 'is_deleted',
    `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 4
  DEFAULT CHARSET = UTF8MB4;

