use polyswap;
ALTER TABLE `token_basic` ADD COLUMN `total_amount` varchar(64);
ALTER TABLE `token_basic` ADD COLUMN `total_count` bigint(20);
ALTER TABLE `token_basic` ADD COLUMN `stats_update_time` bigint(20);
ALTER TABLE `token_basic` ADD COLUMN `social_twitter` varchar(256);
ALTER TABLE `token_basic` ADD COLUMN `social_telegram` varchar(256);
ALTER TABLE `token_basic` ADD COLUMN `social_website` varchar(256);
ALTER TABLE `token_basic` ADD COLUMN `social_other` varchar(256);
