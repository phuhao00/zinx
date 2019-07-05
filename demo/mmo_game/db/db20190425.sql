CREATE INDEX index_lv ON user_table(lv);
CREATE INDEX index_mountlv_mountexp_userid ON user_table(mountlv, mountexp, userid);
CREATE INDEX index_power_userid ON user_table(power, userid);
CREATE INDEX index_dressnum_userid ON user_table(dressnum, userid);
CREATE INDEX index_petnum_userid ON user_table(petnum, userid);