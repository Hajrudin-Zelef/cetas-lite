# INDEX — Step 8 Phase F — Data & Messaging: Developer Angle

Corpus `etape8-phasef-data-messaging` · **17 fichiers** · 757 lignes source · ~8431 mots · partition exacte de `docs/RAG/etape8_phaseF_data_messaging.md`.

## Mode d'emploi

1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur `domain`, `task`, `actors`, `dates` ou `keywords`.
2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité thématique auto-suffisante avec un en-tête YAML.
3. Pour un événement répété dans plusieurs sections, préférer le fichier marqué `canonical_for` (voir la table Événements canoniques).

## Domaines (dossiers → fichiers)

### `00-front-matter/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Step 8 Phase F — Data & Messaging: Developer Angle](00-front-matter/overview.md) | 1–55 | reference | reference |
| 02 | [2. ORMs, query builders, and schema migrations](00-front-matter/2-orms-query-builders-and-schema-migrations.md) | 56–100 | reference | reference |
| 03 | [3. SQLite ecosystem: WAL, Litestream, Turso/libSQL](00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md) | 101–149 | reference | reference |
| 04 | [5. ClickHouse](00-front-matter/5-clickhouse.md) | 150–212 | reference | reference |
| 05 | [7.5 Apache Pulsar](00-front-matter/7-5-apache-pulsar.md) | 213–232 | reference | reference |
| 06 | [8. Event-driven patterns and delivery semantics](00-front-matter/8-event-driven-patterns-and-delivery-semantics.md) | 233–290 | reference | reference |
| 07 | [10. MongoDB developer angle](00-front-matter/10-mongodb-developer-angle.md) | 291–353 | reference | reference |
| 08 | [12.4 Apache Spark 4.x](00-front-matter/12-4-apache-spark-4-x.md) | 354–373 | reference | reference |
| 09 | [13. Version-release timeline (2025 → 2026-09-22)](00-front-matter/13-version-release-timeline-2025-2026-09-22.md) | 374–418 | reference | reference |
| 10 | [15. Glossary and verbatim source index](00-front-matter/15-glossary-and-verbatim-source-index.md) | 419–488 | reference | reference |
| 11 | [16. SQL dialect comparison matrix (developer-relevant)](00-front-matter/16-sql-dialect-comparison-matrix-developer-relevant.md) | 489–591 | reference | reference |

### `01-producer-idempotence-durability-baseline/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [producer: idempotence + durability baseline](01-producer-idempotence-durability-baseline/overview.md) | 592–596 | deep-dive | reference |

### `02-consumer-control-your-own-offsets-for-effectively-once/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [consumer: control your own offsets for effectively-once](02-consumer-control-your-own-offsets-for-effectively-once/overview.md) | 597–609 | deep-dive | reference |

### `03-crashed-consumer-recovery/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [crashed consumer recovery:](03-crashed-consumer-recovery/overview.md) | 610–631 | deep-dive | reference |

### `04-models-marts-schema-yml/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [models/marts/schema.yml](04-models-marts-schema-yml/overview.md) | 632–668 | deep-dive | reference |

### `05-review-alembic-versions-rev-add-is-admin-to-users-py-autogen/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [REVIEW alembic/versions/<rev>_add_is_admin_to_users.py -- autogenerate is a draft](05-review-alembic-versions-rev-add-is-admin-to-users-py-autogen/overview.md) | 669–682 | deep-dive | reference |

### `06-spark-4-default-ansi-mode-on-divide-by-zero-and-bad-casts-ra/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Spark 4 default: ANSI mode on -> divide-by-zero and bad casts raise](06-spark-4-default-ansi-mode-on-divide-by-zero-and-bad-casts-ra/overview.md) | 683–757 | deep-dive | reference |

## Par tâche

- **reference** — [Step 8 Phase F — Data & Messaging: Developer Angle](00-front-matter/overview.md), [2. ORMs, query builders, and schema migrations](00-front-matter/2-orms-query-builders-and-schema-migrations.md), [3. SQLite ecosystem: WAL, Litestream, Turso/libSQL](00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md), [5. ClickHouse](00-front-matter/5-clickhouse.md), [7.5 Apache Pulsar](00-front-matter/7-5-apache-pulsar.md), [8. Event-driven patterns and delivery semantics](00-front-matter/8-event-driven-patterns-and-delivery-semantics.md), [10. MongoDB developer angle](00-front-matter/10-mongodb-developer-angle.md), [12.4 Apache Spark 4.x](00-front-matter/12-4-apache-spark-4-x.md), [13. Version-release timeline (2025 → 2026-09-22)](00-front-matter/13-version-release-timeline-2025-2026-09-22.md), [15. Glossary and verbatim source index](00-front-matter/15-glossary-and-verbatim-source-index.md), [16. SQL dialect comparison matrix (developer-relevant)](00-front-matter/16-sql-dialect-comparison-matrix-developer-relevant.md), [producer: idempotence + durability baseline](01-producer-idempotence-durability-baseline/overview.md), [consumer: control your own offsets for effectively-once](02-consumer-control-your-own-offsets-for-effectively-once/overview.md), [crashed consumer recovery:](03-crashed-consumer-recovery/overview.md), [models/marts/schema.yml](04-models-marts-schema-yml/overview.md), [REVIEW alembic/versions/<rev>_add_is_admin_to_users.py -- autogenerate is a draft](05-review-alembic-versions-rev-add-is-admin-to-users-py-autogen/overview.md), [Spark 4 default: ANSI mode on -> divide-by-zero and bad casts raise](06-spark-4-default-ansi-mode-on-divide-by-zero-and-bad-casts-ra/overview.md)

## Par acteur

- **AWS** (2) — [00-front-matter/10-mongodb-developer-angle.md](00-front-matter/10-mongodb-developer-angle.md), [06-spark-4-default-ansi-mode-on-divide-by-zero-and-bad-casts-ra/overview.md](06-spark-4-default-ansi-mode-on-divide-by-zero-and-bad-casts-ra/overview.md)
- **Oracle** (1) — [00-front-matter/overview.md](00-front-matter/overview.md)

## Par date

- **2024-03** — [00-front-matter/8-event-driven-patterns-and-delivery-semantics.md](00-front-matter/8-event-driven-patterns-and-delivery-semantics.md)
- **2024-10-21** — [00-front-matter/7-5-apache-pulsar.md](00-front-matter/7-5-apache-pulsar.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2025-04** — [00-front-matter/10-mongodb-developer-angle.md](00-front-matter/10-mongodb-developer-angle.md)
- **2025-05** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/8-event-driven-patterns-and-delivery-semantics.md](00-front-matter/8-event-driven-patterns-and-delivery-semantics.md), [00-front-matter/12-4-apache-spark-4-x.md](00-front-matter/12-4-apache-spark-4-x.md)
- **2025-05-21** — [00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md](00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md), [00-front-matter/15-glossary-and-verbatim-source-index.md](00-front-matter/15-glossary-and-verbatim-source-index.md)
- **2025-05-29** — [00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md](00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2025-09** — [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2025-09-11** — [00-front-matter/15-glossary-and-verbatim-source-index.md](00-front-matter/15-glossary-and-verbatim-source-index.md)
- **2025-09-16** — [00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md](00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md), [00-front-matter/15-glossary-and-verbatim-source-index.md](00-front-matter/15-glossary-and-verbatim-source-index.md)
- **2026-02** — [00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md](00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md), [00-front-matter/8-event-driven-patterns-and-delivery-semantics.md](00-front-matter/8-event-driven-patterns-and-delivery-semantics.md)
- **2026-03** — [00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md](00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md), [00-front-matter/10-mongodb-developer-angle.md](00-front-matter/10-mongodb-developer-angle.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-03-02** — [00-front-matter/15-glossary-and-verbatim-source-index.md](00-front-matter/15-glossary-and-verbatim-source-index.md)
- **2026-03-24** — [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-04** — [00-front-matter/7-5-apache-pulsar.md](00-front-matter/7-5-apache-pulsar.md), [00-front-matter/8-event-driven-patterns-and-delivery-semantics.md](00-front-matter/8-event-driven-patterns-and-delivery-semantics.md)
- **2026-04-14** — [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-04-27** — [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-04-30** — [00-front-matter/5-clickhouse.md](00-front-matter/5-clickhouse.md), [00-front-matter/7-5-apache-pulsar.md](00-front-matter/7-5-apache-pulsar.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-05-13** — [00-front-matter/10-mongodb-developer-angle.md](00-front-matter/10-mongodb-developer-angle.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-05-14** — [00-front-matter/10-mongodb-developer-angle.md](00-front-matter/10-mongodb-developer-angle.md)
- **2026-06** — [00-front-matter/10-mongodb-developer-angle.md](00-front-matter/10-mongodb-developer-angle.md)
- **2026-06-23** — [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-07-06** — [00-front-matter/7-5-apache-pulsar.md](00-front-matter/7-5-apache-pulsar.md)
- **2026-07-14** — [00-front-matter/12-4-apache-spark-4-x.md](00-front-matter/12-4-apache-spark-4-x.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-08-20** — [00-front-matter/10-mongodb-developer-angle.md](00-front-matter/10-mongodb-developer-angle.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-08-27** — [00-front-matter/15-glossary-and-verbatim-source-index.md](00-front-matter/15-glossary-and-verbatim-source-index.md)
- **2026-09** — [00-front-matter/8-event-driven-patterns-and-delivery-semantics.md](00-front-matter/8-event-driven-patterns-and-delivery-semantics.md)
- **2026-09-22** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md), [06-spark-4-default-ansi-mode-on-divide-by-zero-and-bad-casts-ra/overview.md](06-spark-4-default-ansi-mode-on-divide-by-zero-and-bad-casts-ra/overview.md)
- **2026-09-24** — [00-front-matter/7-5-apache-pulsar.md](00-front-matter/7-5-apache-pulsar.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-10-21** — [00-front-matter/7-5-apache-pulsar.md](00-front-matter/7-5-apache-pulsar.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-12-19** — [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2026-12-31** — [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)
- **2027-10-21** — [00-front-matter/7-5-apache-pulsar.md](00-front-matter/7-5-apache-pulsar.md), [00-front-matter/13-version-release-timeline-2025-2026-09-22.md](00-front-matter/13-version-release-timeline-2025-2026-09-22.md)

## Carte de couverture (lignes source)

| plage | fichier |
|---|---|
| 1–55 | etape8-phasef-data-messaging/00-front-matter/overview.md |
| 56–100 | etape8-phasef-data-messaging/00-front-matter/2-orms-query-builders-and-schema-migrations.md |
| 101–149 | etape8-phasef-data-messaging/00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql.md |
| 150–212 | etape8-phasef-data-messaging/00-front-matter/5-clickhouse.md |
| 213–232 | etape8-phasef-data-messaging/00-front-matter/7-5-apache-pulsar.md |
| 233–290 | etape8-phasef-data-messaging/00-front-matter/8-event-driven-patterns-and-delivery-semantics.md |
| 291–353 | etape8-phasef-data-messaging/00-front-matter/10-mongodb-developer-angle.md |
| 354–373 | etape8-phasef-data-messaging/00-front-matter/12-4-apache-spark-4-x.md |
| 374–418 | etape8-phasef-data-messaging/00-front-matter/13-version-release-timeline-2025-2026-09-22.md |
| 419–488 | etape8-phasef-data-messaging/00-front-matter/15-glossary-and-verbatim-source-index.md |
| 489–591 | etape8-phasef-data-messaging/00-front-matter/16-sql-dialect-comparison-matrix-developer-relevant.md |
| 592–596 | etape8-phasef-data-messaging/01-producer-idempotence-durability-baseline/overview.md |
| 597–609 | etape8-phasef-data-messaging/02-consumer-control-your-own-offsets-for-effectively-once/overview.md |
| 610–631 | etape8-phasef-data-messaging/03-crashed-consumer-recovery/overview.md |
| 632–668 | etape8-phasef-data-messaging/04-models-marts-schema-yml/overview.md |
| 669–682 | etape8-phasef-data-messaging/05-review-alembic-versions-rev-add-is-admin-to-users-py-autogen/overview.md |
| 683–757 | etape8-phasef-data-messaging/06-spark-4-default-ansi-mode-on-divide-by-zero-and-bad-casts-ra/overview.md |

