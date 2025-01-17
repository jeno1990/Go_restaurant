CREATE TABLE IF NOT EXISTS `companies` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
`description` varchar(255) NOT NULL,
`contact_email` varchar(255) NOT NULL,
`contact_phone` varchar(20) NOT NULL
);