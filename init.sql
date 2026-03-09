use gin_gateway;
-- user表数据
insert into `users` (`id`, `name`, `nick_name`, `email`,`password`, `avatar`, `mobile`, `employ_date`,`resign_date` ,`department_id`,`position_id`, `description`,
                     `created_at`, `updated_at`, `is_deleted`)
values ('1', 'Sutter Wu', 'sutter', 'sutter.wu@itforce-tech.com', '$2a$04$toEgtXNyEIVIcsBD8RJkEONlfBU37l65pTwjcLsfuOgNvZp4tbd4C','https://youke1.picui.cn/s1/2025/10/25/68fc6b8106bed.jpeg', '18862857104', '2024-08-08', '2024-12-10',1,
        1,'啦啦啦', '2024-08-08', '2024-08-08', 0);
insert into `users` (`id`, `name`, `nick_name`, `email`,`password`, `avatar`, `mobile`, `employ_date`,`resign_date`, `department_id`,`position_id`,
                        `description`,`created_at`, `updated_at`, `is_deleted`)
values ('2', 'sutter2', 'sutter2', 'sutter2.wu@itforce-tech.com', '$2a$04$toEgtXNyEIVIcsBD8RJkEONlfBU37l65pTwjcLsfuOgNvZp4tbd4C','https://youke1.picui.cn/s1/2025/10/25/68fc6b8106bed.jpeg', '18862857104', '2024-08-08', '2024-08-10',1,
        1, '666','2024-08-08', '2024-08-08', 0);

-- 角色表
insert into `roles` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (1, '超级管理员', 'administrator','超级管理员', '2024-08-08 17:23:00','2024-08-08 08:00:23',  0);
insert into `roles` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (2, '系统管理员', 'System Administrator', '系统管理员','2024-08-08 17:23:00','2024-08-08 08:00:23',  0);

-- 用户角色表
insert into `user_role` (`user_id`, `role_id`) values ( 1, 1);
insert into `user_role` (`user_id`, `role_id`) values ( 1, 2);

-- 菜单表
insert into `menus` (`id`, `name`, `code`,`path`, `parent_menu_id`, `menu_status`, `icon_name`,`description`,`created_at`, `updated_at`, `is_deleted`) values (1, '个人中心', 'personal', '/personal',0,1,'HomeOutlined','个人中心','2024-08-08 17:23:00','2024-08-08 08:00:23',  0);
insert into `menus` (`id`, `name`, `code`,`path`, `parent_menu_id`, `menu_status`, `icon_name`,`description`,`created_at`, `updated_at`, `is_deleted`) values (2, '系统看板', 'dashboard', '/dashboard',0,1,'AreaChartOutlined','系统看板','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `menus` (`id`, `name`, `code`,`path`, `parent_menu_id`, `menu_status`, `icon_name`,`description`,`created_at`, `updated_at`, `is_deleted`) values (3, '系统设置', 'systems', '/systems',0,1,'SettingOutlined','系统设置','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `menus` (`id`, `name`, `code`,`path`, `parent_menu_id`, `menu_status`, `icon_name`,`description`,`created_at`, `updated_at`, `is_deleted`) values (4, '用户管理', 'users', '/systems/users',3,1,'SettingOutlined','用户管理','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `menus` (`id`, `name`, `code`,`path`, `parent_menu_id`, `menu_status`, `icon_name`,`description`,`created_at`, `updated_at`, `is_deleted`) values (5, '角色管理', 'roles', '/systems/roles',3,1,'TeamOutlined','角色管理','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `menus` (`id`, `name`, `code`,`path`, `parent_menu_id`, `menu_status`, `icon_name`,`description`,`created_at`, `updated_at`, `is_deleted`) values (6, '菜单管理', 'menus', '/systems/menus',3,1,'MenuOutlined','菜单管理','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `menus` (`id`, `name`, `code`,`path`, `parent_menu_id`, `menu_status`, `icon_name`,`description`,`created_at`, `updated_at`, `is_deleted`) values (7, '按钮管理', 'buttons', '/systems/buttons',3,1,'MenuUnfoldOutlined','按钮管理','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);

-- 角色菜单表
-- insert into `role_menu` (`id`, `role_id`, `menu_id`,`created_at`, `updated_at`, `is_deleted`) values (1, 1, 1, '2024-08-08 17:23:00','2024-08-08 08:00:23',  0);
-- insert into `role_menu` (`id`, `role_id`, `menu_id`,`created_at`, `updated_at`, `is_deleted`) values (1, 1, 2, '2024-08-08 17:23:00','2024-08-08 08:00:23',  0);

insert into `role_menu` (`role_id`, `menu_id`) values ( 1, 1);
insert into `role_menu` (`role_id`, `menu_id`) values ( 1, 2);
insert into `role_menu` (`role_id`, `menu_id`) values ( 1, 3);
insert into `role_menu` (`role_id`, `menu_id`) values ( 1, 4);
insert into `role_menu` (`role_id`, `menu_id`) values ( 1, 5);
insert into `role_menu` (`role_id`, `menu_id`) values ( 1, 6);
insert into `role_menu` (`role_id`, `menu_id`) values ( 1, 7);














-- 部门表
insert into `departments` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('1', '人事行政部', '2024-08-08', '2024-09-10', 0);
insert into `departments` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('2', '财务部', '2024-08-08', '2024-09-10', 0);
insert into `departments` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('3', '研发部', '2024-08-08', '2024-09-10', 0);
insert into `departments` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('4', '信息技术部', '2024-08-08', '2024-09-10', 0);

-- 岗位表
insert into `positions` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('1', '人事', '2024-08-08', '2024-09-10', 0);
insert into `positions` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('2', '开发', '2024-08-08', '2024-09-10', 0);
insert into `positions` (`id`, `name`, `created_at`, `updated_at`, `is_deleted`) values ('3', '行政', '2024-08-08', '2024-09-10', 0);

-- 食谱表
insert into `cook_books` (`id`, `name`, `price`, `image`,`description`,`created_at`, `updated_at`, `is_deleted`) values ('1', '糯米糕',6.00, 'https://youke1.picui.cn/s1/2025/10/28/68ffa0e8826a2.png', '糯米糕是中国传统糕类小吃，以糯米（或糯米粉）为核心原料，辅以面粉、豆沙、红枣、红糖、坚果（如核桃、花生）等配料制成，属于蒸制或油炸类甜点。部分地区也会加入野草汁（如马田糯米糕）、牛奶（如奶香红枣糯米糕）等提升风味。', '2024-08-08', '2024-09-10', 0);
insert into `cook_books` (`id`, `name`, `price`, `image`,`description`,`created_at`, `updated_at`, `is_deleted`) values ('2', '水煮虾滑',10.00, 'https://youke1.picui.cn/s1/2025/10/28/68ffa1729bc3e.png', '水煮虾滑是一道菜品，主料是虾，配料是油菜、鸡蛋白等，调料为淀粉、色拉油、食盐等，通过加热水煮的做法而成。', '2024-08-08', '2024-09-10', 0);
insert into `cook_books` (`id`, `name`, `price`, `image`,`description`,`created_at`, `updated_at`, `is_deleted`) values ('3', '鸡蛋羹',3.00, 'https://youke1.picui.cn/s1/2025/10/28/68ffa1d45ac1f.png', '鸡蛋是人类最好的营养来源之一，鸡蛋中含有大量的维生素和矿物质及有高生物价值的蛋白质。对人而言，鸡蛋的蛋白质品质最佳，仅次于母乳。一个鸡蛋所含的热量，相当于半个苹果或半杯牛奶的热量，但是它还拥有8%的磷、4%的锌、4%的铁、12.6%的蛋白质、6%的维生素D、3%的维生素E、6%的维生素A、2%的维生素B、5%的维生素B2.4%的维生素B6。这些营养都是人体必不可少的，它们起着极其重要的作用，如修复人体组织、形成新的组织、消耗能量和参与复杂的新陈代谢过程等。', '2024-08-08', '2024-09-10', 0);
insert into `cook_books` (`id`, `name`, `price`, `image`,`description`,`created_at`, `updated_at`, `is_deleted`) values ('4', '烤鸡腿',8.00, 'https://youke1.picui.cn/s1/2025/10/28/68ffa22740090.png', '鸡腿是鸡从脚到腿的部位，及腿根一带的肉，其肉质颇坚硬，连皮吃的话脂肪含量较多。鸡腿的肉质细嫩，滋味鲜美，味道较淡，可使用于各种料理中。一般人群均可食用，老人、病人、体弱者、贫血患者更宜食用。鸡腿肉高蛋白低脂肪，鸡腿皮的脂肪含量较高，因此不可把带皮的鸡腿视为低热量的食品。同时鸡腿含有丰富的维生素B12、维生素B6、维生素A、维生素D和维生素K等，也是磷、铁、铜和锌的良好来源。', '2024-08-08', '2024-09-10', 0);

-- 车牌表
insert into `plate_numbers` (`id`, `name`, `user_id`, `created_at`, `updated_at`, `is_deleted`) values (1, '粤E9UH30', 1, '2024-08-08', '2024-09-10', 0);
insert into `plate_numbers` (`id`, `name`, `user_id`, `created_at`, `updated_at`, `is_deleted`) values (2, '粤E9UH31', 2, '2024-08-08', '2024-09-10', 0);

-- 停车记录表
insert into `park_records` (`id`, `plate_number_id`, `in_time`,`out_time`,`created_at`, `updated_at`, `is_deleted`) values (1, 1, '2024-08-08 08:00:23', '2024-08-08 17:00:00','2024-08-08 08:00:23', '2024-08-08 08:00:23', 0);
insert into `park_records` (`id`, `plate_number_id`, `in_time`,`out_time`,`created_at`, `updated_at`, `is_deleted`) values (2, 2, '2024-08-08 08:23:23', '2024-08-08 17:23:00','2024-08-08 08:00:23', '2024-08-08 08:00:23', 0);



