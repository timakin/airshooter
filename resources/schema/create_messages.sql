CREATE TABLE IF NOT EXISTS messages (
    `id`           int(10) unsigned not null auto_increment,
    `text`         text not null,
    `created_at`   int(10) unsigned not null,
    `updated_at`   int(10) unsigned not null,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
