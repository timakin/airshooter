CREATE TABLE IF NOT EXISTS message_recipients (
    `id`              int(10) unsigned not null auto_increment,
    `message_id`      int(10) unsigned not null,
    `recipient_id`    int(10) unsigned not null,
    `recipient_type`  varchar(32) not null default "",
    `created_at`      int(10) unsigned not null,
    `updated_at`      int(10) unsigned not null,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
