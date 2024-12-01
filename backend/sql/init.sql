-- ----------------------------
-- Table structure for "permission"
-- ----------------------------
DROP TABLE IF EXISTS "permission";
CREATE TABLE "permission" (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  code VARCHAR(50) NOT NULL UNIQUE,
  type VARCHAR(255) NOT NULL,
  parentId INTEGER,
  path VARCHAR(255),
  redirect VARCHAR(255),
  icon VARCHAR(255),
  component VARCHAR(255),
  layout VARCHAR(255),
  keepAlive SMALLINT,
  method VARCHAR(255),
  description VARCHAR(255),
  "show" SMALLINT NOT NULL DEFAULT 1, 
  enable SMALLINT NOT NULL DEFAULT 1,
  "order" INTEGER 
);

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO "permission" (id, name, code, type, parentId, path, redirect, icon, component, layout, keepAlive, method, description, "show", enable, "order")
VALUES
(1, '资源管理', 'Resource_Mgt', 'MENU', 2, '/pms/resource', NULL, 'i-fe:list', '/src/views/pms/resource/index.vue', NULL, NULL, NULL, NULL, 1, 1, 1),
(2, '系统管理', 'SysMgt', 'MENU', NULL, NULL, NULL, 'i-fe:grid', NULL, NULL, NULL, NULL, NULL, 1, 1, 2),
(3, '角色管理', 'RoleMgt', 'MENU', 2, '/pms/role', NULL, 'i-fe:user-check', '/src/views/pms/role/index.vue', NULL, NULL, NULL, NULL, 1, 1, 2),
(4, '用户管理', 'UserMgt', 'MENU', 2, '/pms/user', NULL, 'i-fe:user', '/src/views/pms/user/index.vue', NULL, 1, NULL, NULL, 1, 1, 3),
(5, '分配用户', 'RoleUser', 'MENU', 3, '/pms/role/user/:roleId', NULL, 'i-fe:user-plus', '/src/views/pms/role/role-user.vue', 'full', NULL, NULL, NULL, 0, 1, 1),
(6, '业务示例', 'Demo', 'MENU', NULL, NULL, NULL, 'i-fe:grid', NULL, NULL, NULL, NULL, NULL, 1, 1, 1),
(7, '图片上传', 'ImgUpload', 'MENU', 6, '/demo/upload', NULL, 'i-fe:image', '/src/views/demo/upload/index.vue', '', 1, NULL, NULL, 1, 1, 2),
(8, '个人资料', 'UserProfile', 'MENU', NULL, '/profile', NULL, 'i-fe:user', '/src/views/profile/index.vue', NULL, NULL, NULL, NULL, 0, 1, 99),
(9, '基础功能', 'Base', 'MENU', NULL, '', NULL, 'i-fe:grid', NULL, '', NULL, NULL, NULL, 1, 1, 0),
(10, '基础组件', 'BaseComponents', 'MENU', 9, '/base/components', NULL, 'i-me:awesome', '/src/views/base/index.vue', NULL, NULL, NULL, NULL, 1, 1, 1),
(11, 'Unocss', 'Unocss', 'MENU', 9, '/base/unocss', NULL, 'i-me:awesome', '/src/views/base/unocss.vue', NULL, NULL, NULL, NULL, 1, 1, 2),
(12, 'KeepAlive', 'KeepAlive', 'MENU', 9, '/base/keep-alive', NULL, 'i-me:awesome', '/src/views/base/keep-alive.vue', NULL, 1, NULL, NULL, 1, 1, 3),
(13, '创建新用户', 'AddUser', 'BUTTON', 4, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 1, 1, 1),
(14, '图标 Icon', 'Icon', 'MENU', 9, '/base/icon', NULL, 'i-fe:feather', '/src/views/base/unocss-icon.vue', '', NULL, NULL, NULL, 1, 1, 0),
(15, 'MeModal', 'TestModal', 'MENU', 9, '/testModal', NULL, 'i-me:dialog', '/src/views/base/test-modal.vue', NULL, NULL, NULL, NULL, 1, 1, 5);

-- ----------------------------
-- Table structure for profile
-- ----------------------------
DROP TABLE IF EXISTS "profile";
CREATE TABLE "profile" (
  id SERIAL PRIMARY KEY,
  gender INTEGER,
  avatar VARCHAR(255) NOT NULL DEFAULT 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80',
  address VARCHAR(255),
  email VARCHAR(255),
  userId INTEGER NOT NULL UNIQUE,
  nickName VARCHAR(10)
);

-- ----------------------------
-- Records of profile
-- ----------------------------
INSERT INTO "profile" VALUES (1, NULL, 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80', NULL, NULL, 1, 'Admin');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS "role";
CREATE TABLE "role" (
  id SERIAL PRIMARY KEY,
  code VARCHAR(50) NOT NULL,
  name VARCHAR(50) NOT NULL,
  enable SMALLINT NOT NULL DEFAULT 1
);

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO "role" (id, code, name, enable) VALUES (1, 'SUPER_ADMIN', '超级管理员', 1);
INSERT INTO "role" (id, code, name, enable) VALUES (2, 'ROLE_QA', '质检员', 1);

-- ----------------------------
-- Table structure for role_permissions
-- ----------------------------
DROP TABLE IF EXISTS "role_permissions";
CREATE TABLE "role_permissions" (
  roleId INTEGER NOT NULL,
  permissionId INTEGER NOT NULL,
  PRIMARY KEY (roleId, permissionId)
);

-- ----------------------------
-- Records of role_permissions
-- ----------------------------
INSERT INTO "role_permissions" VALUES (2, 1);
INSERT INTO "role_permissions" VALUES (2, 2);
INSERT INTO "role_permissions" VALUES (2, 3);
INSERT INTO "role_permissions" VALUES (2, 4);
INSERT INTO "role_permissions" VALUES (2, 5);
INSERT INTO "role_permissions" VALUES (2, 9);
INSERT INTO "role_permissions" VALUES (2, 10);
INSERT INTO "role_permissions" VALUES (2, 11);
INSERT INTO "role_permissions" VALUES (2, 12);
INSERT INTO "role_permissions" VALUES (2, 14);
INSERT INTO "role_permissions" VALUES (2, 15);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR(50) NOT NULL UNIQUE,
  "password" VARCHAR(255) NOT NULL,
  "enable" SMALLINT NOT NULL DEFAULT 1,
  "createTime" TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  "updateTime" TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)
);
-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO "user" VALUES (1, 'admin', '$2a$10$FsAafxTTVVGXfIkJqvaiV.1vPfq4V9HW298McPldJgO829PR52a56', 1, '2023-11-18 16:18:59.150632', '2023-11-18 16:18:59.150632');

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS "user_roles";
CREATE TABLE "user_roles" (
  "userId" INT NOT NULL,
  "roleId" INT NOT NULL,
  PRIMARY KEY ("userId", "roleId")
);

-- ----------------------------
-- Records of user_roles
-- ----------------------------
INSERT INTO "user_roles" VALUES (1, 1);
INSERT INTO "user_roles" VALUES (1, 2);
