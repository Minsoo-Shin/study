
# Lock 조회
```sql
SELECT t.relname, l.locktype, page, virtualtransaction, pid, mode, granted FROM pg_locks l, pg_stat_all_tables t WHERE l.relation = t.relid ORDER BY relation ASC;
```

# Lock 해제
```sql
SELECT pg_cancel_backend(PID);
```

# 실행중인 query 상태 조회
```sql
select * from pg_stat_activity;
```

