use gin_gateway;
-- user表数据
insert into `users` (`id`, `name`, `nick_name`, `email`,`password`, `avatar`, `mobile`, `employ_date`,`resign_date` ,`department_id`,`position_id`,
                     `created_at`, `updated_at`, `is_deleted`)
values ('1', 'sutter', 'sutter', 'sutter.wu@itforce-tech.com', '$2a$04$toEgtXNyEIVIcsBD8RJkEONlfBU37l65pTwjcLsfuOgNvZp4tbd4C','https://youke1.picui.cn/s1/2025/10/25/68fc6b8106bed.jpeg', '18862857104', '2024-08-08', '2024-12-10',1,
        1, '2024-08-08', '2024-08-08', 0);
insert into `users` (`id`, `name`, `nick_name`, `email`,`password`, `avatar`, `mobile`, `employ_date`,`resign_date`, `department_id`,`position_id`,
                     `created_at`, `updated_at`, `is_deleted`)
values ('2', 'sutter2', 'sutter2', 'sutter2.wu@itforce-tech.com', '$2a$04$toEgtXNyEIVIcsBD8RJkEONlfBU37l65pTwjcLsfuOgNvZp4tbd4C','https://youke1.picui.cn/s1/2025/10/25/68fc6b8106bed.jpeg', '18862857104', '2024-08-08', '2024-08-10',1,
        1, '2024-08-08', '2024-08-08', 0);

-- 部门表
insert into `departments` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('1', '人事行政部', '2024-08-08', '2024-09-10', 0);
insert into `departments` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('2', '财务部', '2024-08-08', '2024-09-10', 0);
insert into `departments` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('3', '研发部', '2024-08-08', '2024-09-10', 0);
insert into `departments` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('4', '信息技术部', '2024-08-08', '2024-09-10', 0);

-- 岗位表
insert into `positions` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('1', '人事', '2024-08-08', '2024-09-10', 0);
insert into `positions` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('2', '开发', '2024-08-08', '2024-09-10', 0);
insert into `positions` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('3', '行政', '2024-08-08', '2024-09-10', 0);





