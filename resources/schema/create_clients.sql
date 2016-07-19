CREATE TABLE IF NOT EXISTS clients (
    `uid`         varchar(128) not null default "",
    `secret`      varchar(128) not null default "",
    `created_at`  int(10) unsigned not null,
    `updated_at`  int(10) unsigned not null,
    PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
