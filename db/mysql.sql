/*
警告: 字段名可能非法 - status
*/
create table  ti_menu
(
       menu_id           INTEGER(10) not null 	/*编号 功能菜单编号*/,
       menu_name         VARCHAR(30) 	/*菜单功能名 功能菜单名称*/,
       route             VARCHAR(64) 	/*路由*/,
       color             VARCHAR(32) 	/*菜单颜色 菜单颜色*/,
       createdate        DATETIME 	/*创建日期*/,
       newdate           DATETIME 	/*更新日期*/,
       user_id           VARCHAR(10) 	/*用户ID 具有权限的用户，逗号分割*/,
       status            INTEGER(32) 	/*按钮状态 按钮状态，0启用，1失效*/
);
alter  table ti_menu
       add constraint PK_ti_menu_menu_id primary key (menu_id);
insert into `ti_menu` values(0, "IT服务-测试", "/repair/list", "#FA8C72","2022-05-23 11:23:30","2022-05-23 11:23:30",0,0);
insert into `ti_menu` values(1, "超快反馈工具", "/reply", "#FCC662","2022-05-23 11:23:30","2022-05-23 11:23:30",0,0);
insert into `ti_menu` values(2, "物品借用", "/borrow", "#C1B9AE","2022-05-23 11:23:30","2022-05-23 11:23:30",0,0);
insert into `ti_menu` values(3, "资产、耗材领用登记", "/check_in", "#6C5B7B","2022-05-23 11:23:30","2022-05-23 11:23:30",0,0);
insert into `ti_menu` values(4, "工作交接助手", "/hand_over", "#0E9283","2022-05-23 11:23:30","2022-05-23 11:23:30",0,0);