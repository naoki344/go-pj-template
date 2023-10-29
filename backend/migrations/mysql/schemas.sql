CREATE TABLE `notes`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Noteの識別子',
    `title`    varchar(20) NOT NULL COMMENT 'タイトル',
    `content` VARCHAR(80) NOT NULL COMMENT 'コンテンツ',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ノート';
