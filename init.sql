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
insert into `menus` (`id`, `name`, `code`,`path`, `parent_menu_id`, `menu_status`, `icon_name`,`description`,`created_at`, `updated_at`, `is_deleted`) values (8, '字典管理', 'mappings', '/systems/mappings',3,1,'SafetyOutlined','字典管理','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `menus` (`id`, `name`, `code`,`path`, `parent_menu_id`, `menu_status`, `icon_name`,`description`,`created_at`, `updated_at`, `is_deleted`) values (9, '接口管理', 'connectors', '/systems/connectors',3,1,'SafetyOutlined','接口管理','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);

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
insert into `role_menu` (`role_id`, `menu_id`) values ( 1, 8);
insert into `role_menu` (`role_id`, `menu_id`) values ( 1, 9);
-- 按钮表
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (1, '添加用户', 'system:user:create', '添加用户','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (2, '添加角色', 'system:role:create', '添加角色','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (3, '添加菜单', 'system:menu:create', '添加菜单','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (4, '添加按钮', 'system:button:create', '添加按钮','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (5, '添加字典', 'system:mapping:create', '添加字典','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (6, '添加接口', 'system:connector:create', '添加接口','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (7, '编辑用户', 'system:user:edit', '编辑用户','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (8, '编辑角色', 'system:role:edit', '编辑角色','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (9, '编辑菜单', 'system:menu:edit', '编辑菜单','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (10, '编辑按钮', 'system:button:edit', '编辑按钮','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (11, '编辑字典', 'system:mapping:edit', '编辑字典','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (12, '编辑接口', 'system:connector:edit', '编辑接口','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (13, '删除用户', 'system:user:delete', '删除用户','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (14, '删除角色', 'system:role:delete', '删除角色','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (15, '删除菜单', 'system:menu:delete', '删除菜单','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (16, '删除按钮', 'system:button:delete', '删除按钮','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (17, '删除字典', 'system:mapping:delete', '删除字典','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);
insert into `buttons` (`id`, `name`, `code`,`description`,`created_at`, `updated_at`, `is_deleted`) values (18, '删除接口', 'system:connector:delete', '删除接口','2024-08-08 17:24:00','2024-08-08 08:55:23',  0);

-- 角色按钮表
insert into `role_button` (`role_id`, `button_id`) values ( 1, 1);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 2);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 3);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 4);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 5);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 6);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 7);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 8);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 9);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 10);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 11);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 12);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 13);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 14);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 15);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 16);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 17);
insert into `role_button` (`role_id`, `button_id`) values ( 1, 18);

-- 接口表
insert into `connectors` (`id`, `name`, `method`, `path`,`description`, `created_at`, `updated_at`,  `is_deleted`)
values (1, '查询用户列表', 1, '/user/page','查询用户列表','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (2, '添加用户', 2,'/user/create', '添加用户','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (3, '删除用户', 3, '/user/delete','删除用户','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (4, '编辑用户', 2,'/user/update', '编辑用户', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (5, '查询全量用户', 1,'/user/list', '查询全量用户','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (6, '查询用户详情',1,'/user/detail', '查询用户详情','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (7, '查询角色列表',1,'/role/page', '查询角色列表','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (8, '添加角色',2,'/role/create', '添加角色','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (9, '删除角色',3,'/role/delete', '删除角色','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (10, '编辑角色',2,'/role/update', '编辑角色','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (11, '查询全量角色', 1,'/role/list', '查询全量角色','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (12, '查询角色详情',1,'/role/detail', '查询角色详情','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (13, '查询菜单列表',1,'/menu/page', '查询菜单列表','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (14, '添加菜单',2,'/menu/create', '添加菜单','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (15, '删除菜单',3,'/menu/delete', '删除菜单','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (16, '编辑菜单',2,'/menu/update', '编辑菜单','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (17, '查询全量菜单', 1,'/menu/list', '查询全量菜单','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (18, '查询菜单详情',1,'/menu/detail', '查询菜单详情','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (19, '查询菜单树形列表',1,'/menu/treeLists', '查询菜单树形列表','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (20, '查询接口列表',1,'/connector/page', '查询接口列表','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (21, '添加接口',2,'/connector/create', '添加接口', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (22, '删除接口',3,'/connector/delete', '删除接口', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (23, '编辑接口',2,'/connector/update', '编辑接口','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (24, '查询全量接口', 1,'/connector/list', '查询全量接口','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (25, '查询接口详情',1,'/connector/detail', '查询菜单详情','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (26, '查询字典列表',1,'/mapping/page', '查询字典列表','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (27, '添加字典',2,'/mapping/create', '添加字典','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (28, '删除字典',3,'/mapping/delete', '删除字典', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (29, '编辑字典',2,'/mapping/update', '编辑字典', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (30, '查询全量字典类型', 1,'/mapping/types', '查询全量字典类型','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (31, '查询字典详情',1,'/mapping/detail', '查询字典详情', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (32, '查询按钮列表',1,'/button/page', '查询按钮列表', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (33, '添加按钮',2,'/button/create', '添加按钮','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (34, '删除按钮',3,'/button/delete', '删除按钮', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (35, '编辑按钮',2,'/button/update', '编辑按钮','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (36, '编辑按钮',2,'/button/update', '编辑按钮','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (37, '查询全量按钮', 1,'/button/list', '查询全量按钮', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (38, '查询按钮详情',1,'/button/detail', '查询按钮详情','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0);


-- 角色接口表
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 1);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 2);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 3);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 4);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 5);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 6);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 7);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 8);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 9);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 10);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 11);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 12);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 13);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 14);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 15);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 16);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 17);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 18);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 19);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 20);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 21);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 22);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 23);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 24);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 25);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 26);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 27);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 28);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 29);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 30);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 31);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 32);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 33);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 34);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 35);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 36);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 37);
insert into `role_connector` (`role_id`, `connector_id`) values ( 1, 38);
-- 字典表
insert into `mappings` (`id`, `code`,`value`,  `name`,   `color`, `background_color`,`description`, `created_at`, `updated_at`,  `is_deleted`)
values (1, 'status', 0,'禁用', '','','禁用','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (2, 'status', 1, '启用','','', '启用','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (3, 'menuType',1, '目录','','', '目录', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (4, 'menuType',2, '菜单','','', '菜单','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (5, 'connectorType',1,'GET','','', 'GET','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (6, 'connectorType',2,'POST','','', 'POST', '2024-08-08 17:24:00', '2024-08-08 17:24:00', 0),
       (7, 'connectorType',3,'DELETE','','', 'DELETE','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0),
       (8, 'connectorType',4,'PUT','','', 'PUT','2024-08-08 17:24:00', '2024-08-08 17:24:00',  0);











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



