if exist Rose.db (
	del Rose.db
)
if exist Rose.db-shm (
	del Rose.db-shm
)
if exist Rose.db-wal (
	del Rose.db-wal
)
sqlite3 Rose.db ".read db.sql"
pause