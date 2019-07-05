alter table `dynamic_activity` add column mgrBgId integer default 0;
CREATE INDEX index_dynamic_activity_mgrBgId ON dynamic_activity(mgrBgId);