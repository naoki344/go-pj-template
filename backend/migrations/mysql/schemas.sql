CREATE TABLE `customers`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Customerの識別子',
    `name`    VARCHAR(255) NOT NULL COMMENT '顧客名',
    `name_kana` VARCHAR(255) COMMENT '顧客名（カナ）',
    `telephone` VARCHAR(255) NOT NULL COMMENT '電話番号',
    `email` VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    `person_in_charge_name` VARCHAR(255) NOT NULL COMMENT '担当者名',
    `person_in_charge_name_kana` VARCHAR(255) COMMENT '担当者名（カナ）',
    `postal_code` VARCHAR(255) NOT NULL COMMENT '郵便番号',
    `pref_id` INT UNSIGNED NOT NULL COMMENT '都道府県ID',
    `address1` VARCHAR(255) NOT NULL COMMENT '市区町村',
    `address2` VARCHAR(255) NOT NULL COMMENT '番地・建物名・部屋',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='顧客';
