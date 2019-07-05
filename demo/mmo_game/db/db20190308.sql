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