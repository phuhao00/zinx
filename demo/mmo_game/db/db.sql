PRAGMA main.page_size = 65536;
pragma main.default_cache_size=10000;

pragma main.locking_mode=exclusive;
pragma main.synchronous=normal;
pragma main.journal_mode=wal;
pragma main.temp_store = memory;

-- Table: user_table
DROP TABLE IF EXISTS `user_table`;
CREATE TABLE user_table ( 
    userid   INT       NOT NULL
                            UNIQUE
                            PRIMARY KEY,
	account  VARCHAR(64) NOT NULL, 						
    username NVARCHAR( 20 )  NOT NULL 
							 UNIQUE,
    nick     NVARCHAR( 20 )	 NOT NULL 
							 UNIQUE,
	channel   VARCHAR(20),
	platform  VARCHAR(20),
	login  INT,
	lv		 INT,
	viplv	 INT,
	power    INT DEFAULT 0,
	mountid  INT DEFAULT 0,
	mountlv  INT DEFAULT 0,
	mountexp  INT DEFAULT 0,
	petnum   INT DEFAULT 0,
	dressnum  INT DEFAULT 0,
	onlinetime INT,
	offlinetime INT,
	basedata TEXT,
	battledata TEXT,
	moduledata TEXT,
	deviceid TEXT
);
CREATE INDEX index_userid ON user_table(userid);
CREATE INDEX index_username ON user_table(username);
CREATE INDEX index_account ON user_table(account);
CREATE INDEX index_power ON user_table(power);
CREATE INDEX index_mountlv ON user_table(mountlv);
CREATE INDEX index_mountexp ON user_table(mountexp);
CREATE INDEX index_petnum ON user_table(petnum);
CREATE INDEX index_dressnum ON user_table(dressnum);
CREATE INDEX index_lv ON user_table(lv);
CREATE INDEX index_mountlv_mountexp_userid ON user_table(mountlv, mountexp, userid);
CREATE INDEX index_power_userid ON user_table(power, userid);
CREATE INDEX index_dressnum_userid ON user_table(dressnum, userid);
CREATE INDEX index_petnum_userid ON user_table(petnum, userid);


DROP TABLE IF EXISTS `global_table`;
CREATE TABLE global_table ( 
  server_version   INT    NOT NULL,
	name_init		INT,
	name_data Text,
	open_server_time  INT,
	family_battle Text
);
INSERT INTO global_table(server_version,name_init,name_data,open_server_time,family_battle) VALUES(2,0, "",0,"");
DROP TABLE IF EXISTS `mail_template_table`;
CREATE TABLE mail_template_table ( 
    id   INT   NOT NULL UNIQUE  PRIMARY KEY,
	title	NVARCHAR( 20 ),
	content Text,
	attaches Text,
	starttime INT,
	endtime INT
);
CREATE INDEX index_mail_template_id ON mail_template_table(id);


DROP TABLE IF EXISTS `arena_globaldata_table`;
CREATE TABLE arena_globaldata_table ( 
	id 	 INT  NOT NULL UNIQUE  PRIMARY KEY,
	globalData Text
);
CREATE INDEX index_arena_globaldata_id ON arena_globaldata_table(id);

DROP TABLE IF EXISTS `worldboss_globaldata_table`;
CREATE TABLE worldboss_globaldata_table ( 
	id 	 INT  NOT NULL UNIQUE  PRIMARY KEY,
	globalData Text
);
CREATE INDEX index_worldboss_globaldata_id ON worldboss_globaldata_table(id);

DROP TABLE IF EXISTS `roadhall_globaldata_table`;
CREATE TABLE roadhall_globaldata_table ( 
	id 	 INT  NOT NULL UNIQUE  PRIMARY KEY,
	globalData Text
);
CREATE INDEX index_roadhall_globaldata_id ON roadhall_globaldata_table(id);

DROP TABLE IF EXISTS `exchange_saleitem_table`;
CREATE TABLE "exchange_saleitem_table" (
  "id" Text NOT NULL UNIQUE PRIMARY KEY,
  "mainType" integer,
  "subType" integer,
  "itemProtoId" integer,
  "sellerId" integer,
  "curCount" integer,
  "maxCount" integer,
  "unitPrice" integer,
  "saleType" integer,
  "endTime" integer,
  "bookUserIds" text,
  "itemData" text
);
CREATE INDEX index_exchange_saleitem_id ON exchange_saleitem_table(id);
CREATE INDEX index_exchange_saleitem_sellerId ON exchange_saleitem_table(sellerId);
CREATE INDEX index_exchange_saleitem_itemProtoId ON exchange_saleitem_table(itemProtoId);

DROP TABLE IF EXISTS "family_table";
CREATE TABLE "family_table" (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "level" integer,
  "memberCount" integer,
  "icon" integer,
  "name" text,
  "fightValue" integer,
  "detail" text,
  "members" text,
  "status" integer,
  "familyBoss" text,
  "familyHelp" text
);

DROP TABLE IF EXISTS recharge_table;
create table recharge_table
(
  id   Text NOT NULL UNIQUE PRIMARY KEY,
  userId int,
  time  INT
);
create index recharge_table_id_uindex on recharge_table (id);


DROP TABLE IF EXISTS `camp_table`;
CREATE TABLE "camp_table" (
  "id" integer NOT NULL UNIQUE PRIMARY KEY,
  "globalData" text
);
CREATE INDEX index_camp_table_id ON camp_table(id);

DROP TABLE IF EXISTS `camp_member`;
CREATE TABLE "camp_member" (
  "id" integer NOT NULL UNIQUE PRIMARY KEY,
  "campId" integer,
  "data" text
);
CREATE INDEX index_camp_member_id ON camp_member(id);
CREATE INDEX index_camp_member_campId ON camp_member(campId);

DROP TABLE IF EXISTS "whitelist_table";
CREATE TABLE "whitelist_table" (
  "ip" text NOT NULL UNIQUE PRIMARY KEY
);
CREATE INDEX index_whitelist_table_ip ON whitelist_table(ip);

DROP TABLE IF EXISTS `dynamic_activity`;
CREATE TABLE "dynamic_activity" (
  "id" integer NOT NULL UNIQUE PRIMARY KEY,
  "type" integer,
  "name" text,
  "desc" text,
  "showTime" integer,
  "startTime" integer,
  "endTime" integer,
  "removeTime" integer,
  "extraData" text,
  "globalData" text,
  "clientCfgData" text,
  "mgrBgId" integer
);
CREATE INDEX index_dynamic_activity_id ON dynamic_activity(id);
CREATE INDEX index_dynamic_activity_mgrBgId ON dynamic_activity(mgrBgId);

DROP TABLE IF EXISTS `dynamic_activity_userdata`;
CREATE TABLE "dynamic_activity_userdata" (
  "id" Text NOT NULL UNIQUE PRIMARY KEY,
  "activityid" integer,
  "type" integer,
  "userid" integer,
  "updateTime" integer,
  "data" text
);
CREATE INDEX index_dynamic_activity_userdata_id ON dynamic_activity_userdata(id);
CREATE INDEX index_dynamic_activity_userdata_activityid ON dynamic_activity_userdata(activityid);
CREATE INDEX index_dynamic_activity_userdata_userid ON dynamic_activity_userdata(userid);