CREATE TABLE IF NOT EXISTS notifications (
    `id`           int(10) unsigned not null auto_increment,
    `title`        varchar(255)  not null default '',
    `status`        varchar(32)  not null default '',
    `text`         text not null,
    `delivers_at`   int(10) unsigned,
    `expires_at`   int(10) unsigned,
    `created_at`   int(10) unsigned not null,
    `updated_at`   int(10) unsigned not null,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
