---
id: etape7-phasef-databases/00-databases/10-monitoring-observability-and-slow-query-work
title: "10. Monitoring, observability and slow-query work"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["latency", "memory"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [579, 625]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: aae9fc23a86fa2cb180293ce032f0be3e5ba17a14d891edf7d7cbe8aa528c779
---

# 10. Monitoring, observability and slow-query work

## 10. Monitoring, observability and slow-query work

### 10.1 PostgreSQL

- **pg_stat_statements**: tracks execution statistics per normalized query
  (calls, total/mean time, rows, shared blocks). Standard production setup:
  `shared_preload_libraries='pg_stat_statements'`,
  `pg_stat_statements.track='all'`; query the view ordered by
  `total_exec_time` to find top offenders [secondary].
- **auto_explain**: logs execution plans of slow queries automatically
  (`auto_explain.log_min_duration`) — pairs with pg_stat_statements for
  "which query" + "why slow" [secondary].
- **Prometheus exporters**: `prometheuscommunity/postgres-exporter` (SQL-
  driven metrics), `prometheuscommunity/pgbouncer-exporter`; typical stack
  adds Prometheus + Grafana with per-database dashboards; production Docker
  Compose patterns wire exporter → Prometheus → Grafana with TLS in 2026
  guides [secondary].
- **RDS Performance Insights**: long-term retention billed at **$7.50/
  vCPU/month** [secondary].
- Key metrics: connections vs `max_connections`, cache hit ratio
  (`blks_hit/(blks_hit+blks_read)`), bloat (dead tuples), replication lag
  (`pg_stat_replication`), checkpoint frequency, temp files, lock waits
  [secondary].

### 10.2 MySQL/MariaDB

- **Performance Schema** + `sys` schema: `events_statements_summary_by_digest`
  is the MySQL answer to pg_stat_statements [secondary].
- **Slow query log** with `long_query_time`, `log_queries_not_using_indexes`;
  `pt-query-digest` (Percona Toolkit) for aggregation [secondary].
- Exporters: `prometheus/mysqld_exporter`; Percona Monitoring and Management
  (**PMM**) bundles exporters + Grafana dashboards + Query Analytics
  [vendor-reported].
- Key metrics: InnoDB buffer pool hit rate, redo log pressure, replica lag
  (`Seconds_Behind_Master` / GTID gaps), thread pool saturation, table locks
  (MyISAM/Aria legacy) [secondary].

### 10.3 Redis/Valkey

- `INFO` sections (stats, replication, persistence), `SLOWLOG`,
  `LATENCY HISTORY`; exporters: `oliver006/redis_exporter` [secondary].
- Key metrics: hit rate, evicted/expired keys, connected clients, blocked
  clients, replication offset lag, RDB/AOF fork latency, memory
  fragmentation ratio [secondary].

---

