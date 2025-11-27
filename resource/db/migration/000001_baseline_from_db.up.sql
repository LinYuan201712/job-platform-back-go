-- =================================================================
-- 000001_baseline_from_db.up.sql
-- =================================================================

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- -----------------------------------------------------------------
-- 1. 基础字典表模块
-- -----------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `t_industries` (
                                              `id` INT NOT NULL AUTO_INCREMENT COMMENT '行业唯一ID',
                                              `name` VARCHAR(100) NOT NULL COMMENT '行业名称',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_industries_name` (`name`)
    ) COMMENT='行业领域字典表';

INSERT INTO `t_industries` (`id`, `name`) VALUES
                                              (1, '互联网/电子商务'), (2, '计算机软件'), (3, 'IT服务/系统集成'), (4, '通信/电信运营'),
                                              (5, '电子/半导体/集成电路'), (6, '硬件/智能设备'), (7, '金融(银行/证券/保险)'), (8, '房地产/建筑/物业'),
                                              (9, '制造业/工业自动化'), (10, '广告/传媒/公关'), (11, '教育/培训/科研'), (12, '医疗/制药/生物工程'),
                                              (13, '专业服务(咨询/法律/会计)'), (14, '快速消费品(FMCG)'), (15, '零售/批发'), (16, '交通/物流/仓储'),
                                              (17, '能源/化工/环保'), (18, '汽车及零部件'), (19, '服务业(生活/娱乐/餐饮)'), (20, '政府/非营利组织'),
                                              (99, '其他');

CREATE TABLE IF NOT EXISTS `t_company_natures` (
                                                   `id` INT NOT NULL AUTO_INCREMENT COMMENT '企业性质唯一ID',
                                                   `name` VARCHAR(100) NOT NULL COMMENT '企业性质名称',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_company_natures_name` (`name`)
    ) COMMENT='企业性质字典表';

INSERT INTO `t_company_natures` (`id`, `name`) VALUES
                                                   (1, '国有企业'), (2, '集体企业'), (3, '私营企业'), (4, '联营企业'), (5, '股份制企业'),
                                                   (6, '上市公司'), (7, '中外合资企业'), (8, '外商独资企业'), (9, '港澳台商投资企业'), (10, '事业单位'),
                                                   (11, '政府机关'), (12, '非营利组织'), (13, '个人独资企业'), (14, '合伙企业'), (99, '其他');

CREATE TABLE IF NOT EXISTS `t_company_scales` (
                                                  `id` INT NOT NULL AUTO_INCREMENT COMMENT '公司规模唯一ID',
                                                  `scale` VARCHAR(50) NOT NULL COMMENT '规模范围',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_company_scales_scale` (`scale`)
    ) COMMENT='公司规模字典表';

INSERT INTO `t_company_scales` (`id`, `scale`) VALUES
                                                   (1, '0-20人'), (2, '20-99人'), (3, '100-499人'), (4, '500-999人'),
                                                   (5, '1000-4999人'), (6, '5000-9999人'), (7, '10000人以上');

CREATE TABLE IF NOT EXISTS `t_job_categories` (
                                                  `id` INT NOT NULL AUTO_INCREMENT COMMENT '类别唯一ID',
                                                  `name` VARCHAR(50) NOT NULL COMMENT '类别名称 (如: 研发, 产品)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_job_categories_name` (`name`)
    ) COMMENT='岗位职能类别字典表';

INSERT INTO `t_job_categories` (`id`, `name`) VALUES (1, '产品'), (2, '测试'), (3, '开发'), (4, '算法');

CREATE TABLE IF NOT EXISTS `tag_categories` (
                                                `id` INT NOT NULL AUTO_INCREMENT COMMENT '分类唯一ID',
                                                `code` VARCHAR(50) NOT NULL COMMENT '分类短码',
    `name` VARCHAR(100) NOT NULL COMMENT '分类名称',
    `description` VARCHAR(255) NULL COMMENT '分类描述',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_tag_categories_code` (`code`)
    ) COMMENT='标签分类字典表';

INSERT INTO `tag_categories` (`id`, `code`, `name`, `description`) VALUES
                                                                       (1, 'lang', '编程语言', '例如 C++, Java, Python 等'),
                                                                       (2, 'job_function', '职能', '例如 测试工程师, 项目管理 等'),
                                                                       (3, 'framework', '技术框架/库', '例如 React, SpringBoot, TensorFlow 等'),
                                                                       (4, 'tool', '工具/平台', '例如 Git, Docker, K8s 等'),
                                                                       (5, 'soft_skill', '软技能', '例如 沟通能力, 团队协作 等'),
                                                                       (99, 'other', '未分类', '用于归档所有未分类的标签');

CREATE TABLE IF NOT EXISTS `application_statuses` (
                                                      `id` INT NOT NULL AUTO_INCREMENT COMMENT '状态唯一ID',
                                                      `code` INT NOT NULL COMMENT '状态代码',
                                                      `name` VARCHAR(50) NOT NULL COMMENT '状态名称',
    `detail` TEXT NOT NULL COMMENT '状态详情描述',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_application_statuses_code` (`code`)
    ) COMMENT='投递状态字典表';

INSERT INTO `application_statuses` (`code`, `name`, `detail`) VALUES
                                                                  (10, '已投递', '您的简历已成功投递至企业，请耐心等待企业审核，后续通知将通过平台或预留联系方式发送。'),
                                                                  (20, '候选人', '企业已将您加入候选人名单。若后续岗位匹配，将会通过平台或联系方式继续与您沟通。'),
                                                                  (30, '面试邀请', '您的简历已通过初筛，请留意平台及预留的联系方式，企业将向您发送具体的面试安排和通知。'),
                                                                  (40, '通过', '恭喜您已通过本次招聘流程！企业将通过平台或您的联系方式与您进一步沟通入职相关事宜。'),
                                                                  (50, '拒绝', '很遗憾，您的简历未能通过本次筛选。建议您继续关注其他岗位或优化简历后再次申请，祝您求职顺利。');

CREATE TABLE IF NOT EXISTS `t_provinces` (
                                             `id` INT NOT NULL AUTO_INCREMENT COMMENT '省份ID',
                                             `code` VARCHAR(20) NULL COMMENT '行政区划代码',
    `name` VARCHAR(50) NOT NULL COMMENT '省份名称',
    PRIMARY KEY (`id`)
    ) COMMENT='省份字典表';

INSERT INTO `t_provinces` (`id`, `code`, `name`) VALUES
                                                     (1, '110000', '北京市'), (2, '120000', '天津市'), (3, '130000', '河北省'), (4, '140000', '山西省'),
                                                     (5, '150000', '内蒙古自治区'), (6, '210000', '辽宁省'), (7, '220000', '吉林省'), (8, '230000', '黑龙江省'),
                                                     (9, '310000', '上海市'), (10, '320000', '江苏省'), (11, '330000', '浙江省'), (12, '340000', '安徽省'),
                                                     (13, '350000', '福建省'), (14, '360000', '江西省'), (15, '370000', '山东省'), (16, '410000', '河南省'),
                                                     (17, '420000', '湖北省'), (18, '430000', '湖南省'), (19, '440000', '广东省'), (20, '450000', '广西壮族自治区'),
                                                     (21, '460000', '海南省'), (22, '500000', '重庆市'), (23, '510000', '四川省'), (24, '520000', '贵州省'),
                                                     (25, '530000', '云南省'), (26, '540000', '西藏自治区'), (27, '610000', '陕西省'), (28, '620000', '甘肃省'),
                                                     (29, '630000', '青海省'), (30, '640000', '宁夏回族自治区'), (31, '650000', '新疆维吾尔自治区'), (32, '710000', '台湾省'),
                                                     (33, '810000', '香港特别行政区'), (34, '820000', '澳门特别行政区');

CREATE TABLE IF NOT EXISTS `t_cities` (
                                          `id` INT NOT NULL AUTO_INCREMENT COMMENT '城市ID',
                                          `province_id` INT NOT NULL COMMENT '所属省份ID',
                                          `code` VARCHAR(20) NULL COMMENT '行政区划代码',
    `name` VARCHAR(50) NOT NULL COMMENT '城市名称',
    PRIMARY KEY (`id`),
    KEY `fk_cities_province` (`province_id`),
    CONSTRAINT `fk_cities_province` FOREIGN KEY (`province_id`) REFERENCES `t_provinces` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
    ) COMMENT='城市字典表';

INSERT INTO `t_cities` (`province_id`, `code`, `name`) VALUES
                                                           (1, '110100', '北京市'), (2, '120100', '天津市'), (9, '310100', '上海市'), (19, '440100', '广州市'), (19, '440300', '深圳市');

-- -----------------------------------------------------------------
-- 2. 用户与权限模块
-- -----------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `users` (
                                       `id` INT NOT NULL AUTO_INCREMENT COMMENT '用户唯一ID',
                                       `email` VARCHAR(255) NOT NULL COMMENT '登录邮箱',
    `password_hash` VARCHAR(255) NOT NULL COMMENT '加密后的密码',
    `role` INT NOT NULL COMMENT '用户角色 1=student, 2=hr',
    `status` INT NOT NULL DEFAULT '0' COMMENT '账户状态 0=pending, 1=active, 2=disabled',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `last_login_at` DATETIME NULL COMMENT '最近登录时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`),
    KEY `idx_users_email` (`email`)
    ) COMMENT='用户主表';

CREATE TABLE IF NOT EXISTS `students` (
                                          `user_id` INT NOT NULL COMMENT '关联 users.id',
                                          `student_id` VARCHAR(50) NULL COMMENT '学号',
    `avatar_url` VARCHAR(1024) NULL COMMENT '头像',
    `full_name` VARCHAR(100) NULL COMMENT '姓名',
    `phone_number` VARCHAR(20) NULL COMMENT '手机号',
    `gender` INT NULL COMMENT '性别 0=男, 1=女',
    `date_of_birth` DATE NULL COMMENT '出生日期',
    `job_seeking_status` INT NULL COMMENT '求职状态',
    `expected_position` VARCHAR(100) NULL COMMENT '期望岗位',
    `expected_min_salary` INT NULL COMMENT '期望最少薪资(k)',
    `expected_max_salary` INT NULL COMMENT '期望最多薪资(k)',
    `skills_summary` TEXT NULL COMMENT '技能掌握',
    `current_template_id` BIGINT NULL COMMENT '当前选用的模板ID',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `student_id` (`student_id`),
    KEY `idx_students_student_id` (`student_id`),
    CONSTRAINT `students_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
    ) COMMENT='学生信息表';

-- -----------------------------------------------------------------
-- 3. 企业与招聘模块
-- -----------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `companies` (
                                           `company_id` INT NOT NULL AUTO_INCREMENT COMMENT '公司唯一ID',
                                           `user_id` INT NOT NULL COMMENT '关联 users.id',
                                           `company_name` VARCHAR(255) NULL COMMENT '公司名称',
    `description` TEXT NULL COMMENT '公司介绍',
    `logo_url` VARCHAR(255) NULL COMMENT '公司Logo地址',
    `industry_id` INT NULL COMMENT '行业领域ID',
    `nature_id` INT NULL COMMENT '企业性质ID',
    `company_scale_id` INT NULL COMMENT '公司规模ID',
    `company_address` VARCHAR(255) NULL COMMENT '公司地址',
    `contact_person_name` VARCHAR(100) NULL COMMENT '联系人姓名',
    `contact_person_phone` VARCHAR(50) NULL COMMENT '联系人电话',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`company_id`),
    UNIQUE KEY `user_id` (`user_id`),
    KEY `idx_companies_company_name` (`company_name`),
    CONSTRAINT `companies_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_companies_industry` FOREIGN KEY (`industry_id`) REFERENCES `t_industries` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT `fk_companies_nature` FOREIGN KEY (`nature_id`) REFERENCES `t_company_natures` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT `fk_companies_scale` FOREIGN KEY (`company_scale_id`) REFERENCES `t_company_scales` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
    ) COMMENT='存储企业的基本信息';

CREATE TABLE IF NOT EXISTS `company_links` (
                                               `id` INT NOT NULL AUTO_INCREMENT,
                                               `company_id` INT NOT NULL,
                                               `link_name` VARCHAR(100) NOT NULL,
    `link_url` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `company_id` (`company_id`),
    CONSTRAINT `company_links_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`company_id`) ON DELETE CASCADE
    ) COMMENT='企业链接表';

CREATE TABLE IF NOT EXISTS `jobs` (
                                      `id` INT NOT NULL AUTO_INCREMENT COMMENT '岗位唯一ID',
                                      `company_id` INT NOT NULL COMMENT '关联 companies',
                                      `posted_by_user_id` INT NOT NULL COMMENT '关联 users',
                                      `title` VARCHAR(255) NULL COMMENT '岗位名称',
    `description` TEXT NULL COMMENT '岗位描述',
    `tech_requirements` TEXT NULL COMMENT '岗位要求',
    `min_salary` INT NULL COMMENT '最少薪资(k)',
    `max_salary` INT NULL COMMENT '最多薪资(k)',
    `province_id` INT NULL COMMENT '省份ID',
    `city_id` INT NULL COMMENT '城市ID',
    `address_detail` VARCHAR(255) NULL COMMENT '详细地址',
    `work_nature` INT NULL COMMENT '岗位性质 1=实习, 2=全职',
    `deadline` DATE NULL COMMENT '截止日期',
    `status` INT NULL DEFAULT '10' COMMENT '状态 10=pending, 20=approved...',
    `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `type` INT NULL COMMENT '岗位类别ID',
    `department` VARCHAR(100) NULL COMMENT '所属部门',
    `headcount` INT NULL COMMENT '招聘人数',
    `view_count` INT NOT NULL DEFAULT '0' COMMENT '浏览次数',
    `required_degree` INT NULL COMMENT '学历要求',
    `required_start_date` DATE NULL COMMENT '要求入职时间',
    `bonus_points` TEXT NULL COMMENT '加分项',
    PRIMARY KEY (`id`),
    KEY `idx_jobs_company_id` (`company_id`),
    KEY `idx_jobs_status` (`status`),
    KEY `idx_jobs_work_nature` (`work_nature`),
    KEY `idx_jobs_type` (`type`),
    KEY `idx_jobs_view_count` (`view_count`),
    KEY `posted_by_user_id` (`posted_by_user_id`),
    KEY `fk_jobs_province` (`province_id`),
    KEY `fk_jobs_city` (`city_id`),
    CONSTRAINT `jobs_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`company_id`) ON DELETE CASCADE,
    CONSTRAINT `jobs_ibfk_2` FOREIGN KEY (`posted_by_user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION,
    CONSTRAINT `fk_jobs_category` FOREIGN KEY (`type`) REFERENCES `t_job_categories` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT `fk_jobs_province` FOREIGN KEY (`province_id`) REFERENCES `t_provinces` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT `fk_jobs_city` FOREIGN KEY (`city_id`) REFERENCES `t_cities` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
    ) COMMENT='岗位信息表';

CREATE TABLE IF NOT EXISTS `job_views` (
                                           `id` BIGINT NOT NULL AUTO_INCREMENT,
                                           `job_id` INT NOT NULL,
                                           `viewer_user_id` INT NULL,
                                           `viewed_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                           `client_ip` VARCHAR(64) NULL,
    `user_agent` VARCHAR(255) NULL,
    PRIMARY KEY (`id`),
    KEY `idx_job_views_job` (`job_id`),
    KEY `idx_job_views_viewer` (`viewer_user_id`),
    KEY `idx_job_views_viewed_at` (`viewed_at`),
    CONSTRAINT `fk_job_views_job` FOREIGN KEY (`job_id`) REFERENCES `jobs` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_job_views_viewer` FOREIGN KEY (`viewer_user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
    ) COMMENT='岗位浏览记录表';

-- -----------------------------------------------------------------
-- 4. 简历与投递模块
-- -----------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `resume` (
                                        `id` BIGINT NOT NULL AUTO_INCREMENT,
                                        `student_user_id` INT NOT NULL,
                                        `file_name` VARCHAR(255) NULL,
    `file_url` VARCHAR(1024) NOT NULL,
    `file_size` BIGINT NULL,
    `usage_type` VARCHAR(50) NULL,
    `uploaded_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_resume_student_user_id` (`student_user_id`),
    KEY `idx_resume_uploaded_at` (`uploaded_at` DESC),
    CONSTRAINT `resume_ibfk_1` FOREIGN KEY (`student_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
    ) COMMENT='pdf简历表';

CREATE TABLE IF NOT EXISTS `applications` (
                                              `id` INT NOT NULL AUTO_INCREMENT,
                                              `job_id` INT NOT NULL,
                                              `student_user_id` INT NOT NULL,
                                              `resume_id` BIGINT NOT NULL,
                                              `status` INT NOT NULL,
                                              `submitted_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                              `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                              PRIMARY KEY (`id`),
    KEY `idx_applications_job_id` (`job_id`),
    KEY `idx_applications_student_user_id` (`student_user_id`),
    KEY `idx_applications_status` (`status`),
    KEY `idx_applications_job_student` (`job_id`,`student_user_id`),
    KEY `resume_id` (`resume_id`),
    CONSTRAINT `applications_ibfk_1` FOREIGN KEY (`job_id`) REFERENCES `jobs` (`id`) ON DELETE CASCADE,
    CONSTRAINT `applications_ibfk_2` FOREIGN KEY (`student_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `applications_ibfk_3` FOREIGN KEY (`resume_id`) REFERENCES `resume` (`id`) ON DELETE NO ACTION
    ) COMMENT='投递记录表';

CREATE TABLE IF NOT EXISTS `job_favorites` (
                                               `id` BIGINT NOT NULL AUTO_INCREMENT,
                                               `student_user_id` INT NOT NULL,
                                               `job_id` INT NOT NULL,
                                               `saved_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                               PRIMARY KEY (`id`),
    UNIQUE KEY `uk_student_job` (`student_user_id`,`job_id`),
    KEY `idx_job_favorites_student_user_id` (`student_user_id`),
    KEY `idx_job_favorites_saved_at` (`saved_at`),
    KEY `fk_jobfav_job` (`job_id`),
    CONSTRAINT `fk_jobfav_job` FOREIGN KEY (`job_id`) REFERENCES `jobs` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_jobfav_student` FOREIGN KEY (`student_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
    ) COMMENT='岗位收藏表';

-- -----------------------------------------------------------------
-- 5. 学生经历模块
-- -----------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `education_experiences` (
                                                       `id` BIGINT NOT NULL AUTO_INCREMENT,
                                                       `student_user_id` INT NOT NULL,
                                                       `school_name` VARCHAR(255) NULL,
    `degree_level` INT NULL,
    `major` VARCHAR(100) NULL,
    `start_date` DATE NULL,
    `end_date` DATE NULL,
    `major_rank` VARCHAR(50) NULL,
    PRIMARY KEY (`id`),
    KEY `idx_education_experiences_student_user_id` (`student_user_id`),
    KEY `idx_education_experiences_student_start_date` (`student_user_id`,`start_date` DESC),
    CONSTRAINT `education_experiences_ibfk_1` FOREIGN KEY (`student_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
    ) COMMENT='教育经历表';

CREATE TABLE IF NOT EXISTS `work_experiences` (
                                                  `id` BIGINT NOT NULL AUTO_INCREMENT,
                                                  `student_user_id` INT NOT NULL,
                                                  `company_name` VARCHAR(255) NULL,
    `position_title` VARCHAR(100) NULL,
    `start_date` DATE NULL,
    `end_date` DATE NULL,
    `description` TEXT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_work_experiences_student_user_id` (`student_user_id`),
    KEY `idx_work_experiences_student_start_date` (`student_user_id`,`start_date` DESC),
    CONSTRAINT `work_experiences_ibfk_1` FOREIGN KEY (`student_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
    ) COMMENT='实习/工作经历表';

CREATE TABLE IF NOT EXISTS `project_experiences` (
                                                     `id` BIGINT NOT NULL AUTO_INCREMENT,
                                                     `student_user_id` INT NOT NULL,
                                                     `project_name` VARCHAR(255) NULL,
    `role` VARCHAR(100) NULL,
    `project_link` VARCHAR(1024) NULL,
    `start_date` DATE NULL,
    `end_date` DATE NULL,
    `description` TEXT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_project_experiences_student_user_id` (`student_user_id`),
    KEY `idx_project_experiences_student_start_date` (`student_user_id`,`start_date` DESC),
    CONSTRAINT `project_experiences_ibfk_1` FOREIGN KEY (`student_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
    ) COMMENT='项目经历表';

CREATE TABLE IF NOT EXISTS `organization_experiences` (
                                                          `id` BIGINT NOT NULL AUTO_INCREMENT,
                                                          `student_user_id` INT NOT NULL,
                                                          `organization_name` VARCHAR(255) NULL,
    `role` VARCHAR(100) NULL,
    `start_date` DATE NULL,
    `end_date` DATE NULL,
    `description` TEXT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_organization_experiences_student_user_id` (`student_user_id`),
    KEY `idx_organization_experiences_student_start_date` (`student_user_id`,`start_date` DESC),
    CONSTRAINT `organization_experiences_ibfk_1` FOREIGN KEY (`student_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
    ) COMMENT='社团组织经历表';

CREATE TABLE IF NOT EXISTS `competition_experiences` (
                                                         `id` BIGINT NOT NULL AUTO_INCREMENT,
                                                         `student_user_id` INT NOT NULL,
                                                         `competition_name` VARCHAR(255) NULL,
    `role` VARCHAR(100) NULL,
    `award` VARCHAR(100) NULL,
    `date` VARCHAR(20) NULL,
    PRIMARY KEY (`id`),
    KEY `idx_competition_experiences_student_user_id` (`student_user_id`),
    KEY `idx_competition_experiences_student_date` (`student_user_id`,`date` DESC),
    CONSTRAINT `competition_experiences_ibfk_1` FOREIGN KEY (`student_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
    ) COMMENT='竞赛经历表';

-- -----------------------------------------------------------------
-- 6. 平台标签模块
-- -----------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tags` (
                                      `id` INT NOT NULL AUTO_INCREMENT,
                                      `name` VARCHAR(100) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `category_id` INT NOT NULL,
    `created_by` INT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`),
    KEY `fk_tags_category` (`category_id`),
    KEY `fk_tags_created_by` (`created_by`),
    CONSTRAINT `fk_tags_category` FOREIGN KEY (`category_id`) REFERENCES `tag_categories` (`id`) ON UPDATE CASCADE,
    CONSTRAINT `fk_tags_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
    ) COMMENT='技术与能力标签表';

CREATE TABLE IF NOT EXISTS `job_tags` (
                                          `job_id` INT NOT NULL,
                                          `tag_id` INT NOT NULL,
                                          PRIMARY KEY (`job_id`,`tag_id`),
    KEY `idx_job_tags_tag_id` (`tag_id`),
    CONSTRAINT `job_tags_ibfk_1` FOREIGN KEY (`job_id`) REFERENCES `jobs` (`id`) ON DELETE CASCADE,
    CONSTRAINT `job_tags_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE
    ) COMMENT='岗位-标签关联表';

CREATE TABLE IF NOT EXISTS `students_tags` (
                                               `user_id` INT NOT NULL,
                                               `tag_id` INT NOT NULL,
                                               PRIMARY KEY (`user_id`,`tag_id`),
    KEY `tag_id` (`tag_id`),
    CONSTRAINT `students_tags_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `students` (`user_id`) ON DELETE CASCADE,
    CONSTRAINT `students_tags_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE
    ) COMMENT='个人-标签关联表';

-- -----------------------------------------------------------------
-- 7. 活动模块
-- -----------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `events` (
                                        `id` BIGINT NOT NULL AUTO_INCREMENT,
                                        `admin_user_id` INT NOT NULL,
                                        `event_title` VARCHAR(255) NOT NULL,
    `event_summary` VARCHAR(500) NULL,
    `event_start_time` DATETIME NOT NULL,
    `event_end_time` DATETIME NULL,
    `event_location` VARCHAR(255) NULL,
    `event_type` VARCHAR(50) NULL,
    `target_audience` VARCHAR(255) NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `admin_user_id` (`admin_user_id`),
    CONSTRAINT `events_ibfk_1` FOREIGN KEY (`admin_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
    ) COMMENT='求职活动表';

CREATE TABLE IF NOT EXISTS `event_related_jobs` (
                                                    `event_id` BIGINT NOT NULL,
                                                    `job_id` INT NOT NULL,
                                                    PRIMARY KEY (`event_id`,`job_id`),
    KEY `job_id` (`job_id`),
    CONSTRAINT `event_related_jobs_ibfk_1` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE,
    CONSTRAINT `event_related_jobs_ibfk_2` FOREIGN KEY (`job_id`) REFERENCES `jobs` (`id`) ON DELETE CASCADE
    ) COMMENT='活动-岗位关联表';

SET FOREIGN_KEY_CHECKS = 1;