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
  "globalData" text
);
CREATE INDEX index_dynamic_activity_id ON dynamic_activity(id);

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

DROP TABLE IF EXISTS "whitelist_table";
CREATE TABLE "whitelist_table" (
  "ip" text NOT NULL UNIQUE PRIMARY KEY
);
CREATE INDEX index_whitelist_table_ip ON whitelist_table(ip);