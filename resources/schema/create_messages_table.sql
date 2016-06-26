CREATE TABLE IF NOT EXISTS messages (
    `id` int(10) unsigned not null auto_increment,
    `payload_json_object` text not null,
    `recipient_json_object` text not null,
    `status` enum('scheduled', 'waiting', 'delivering', 'finished', 'failed') not null default 'waiting',
    `delivers_at`  int(10) unsigned not null,
    `expires_at`   int(10) unsigned,
    `thread_id` int(10) unsigned,
    `created_at`   int(10) unsigned not null,
    `updated_at`   int(10) unsigned not null,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
