use gin_gateway;
-- user表数据
insert into `users` (`id`, `name`, `nick_name`, `email`,`password`, `avatar`, `mobile`, `employ_date`,`resign_date` ,`department_id`,`position_id`,
                     `created_at`, `updated_at`, `is_deleted`)
values ('1', 'sutter', 'sutter', 'sutter.wu@itforce-tech.com', '$2a$04$toEgtXNyEIVIcsBD8RJkEONlfBU37l65pTwjcLsfuOgNvZp4tbd4C','', '18862857104', '2024-08-08', '2024-12-10',1,
        1, '2024-08-08', '2024-08-08', 0);
insert into `users` (`id`, `name`, `nick_name`, `email`,`password`, `avatar`, `mobile`, `employ_date`,`resign_date`, `department_id`,`position_id`,
                     `created_at`, `updated_at`, `is_deleted`)
values ('2', 'sutter2', 'sutter2', 'sutter2.wu@itforce-tech.com', '$2a$04$toEgtXNyEIVIcsBD8RJkEONlfBU37l65pTwjcLsfuOgNvZp4tbd4C','', '18862857104', '2024-08-08', '2024-08-10',1,
        1, '2024-08-08', '2024-08-08', 0);