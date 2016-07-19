CREATE TABLE IF NOT EXISTS access_tokens (
    `id`         int(10) unsigned not null auto_increment,
    `token`      varchar(256) not null default "",
    `created_at`      int(10) unsigned not null,
    `updated_at`      int(10) unsigned not null,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
