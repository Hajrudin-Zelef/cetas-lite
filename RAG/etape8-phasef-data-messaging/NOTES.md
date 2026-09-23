# NOTES — corpus `etape8-phasef-data-messaging`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

La source contient 7 titres H1 : un dossier par document H1 (un entête éventuel devient `00-front-matter`).

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 8 Phase F — Data & Messaging: Developer Angle` — 757 lignes, 15 chunks.

- 0. Quick map of this file
- 1. SQL dialects and the modern SQL landscape
- 2. ORMs, query builders, and schema migrations
- 3. SQLite ecosystem: WAL, Litestream, Turso/libSQL
- 4. DuckDB and MotherDuck
- 5. ClickHouse
- 6. Search engines
- 7. Event streaming and messaging
- 8. Event-driven patterns and delivery semantics
- 9. Redis/Valkey developer angle
- 10. MongoDB developer angle
- 11. Time-series databases
- 12. Data pipelines: dbt, Airflow, Dagster, Spark, ETL vs ELT
- 13. Version-release timeline (2025 → 2026-09-22)
- 14. Conflicts, gaps, and unverified claims register
- 15. Glossary and verbatim source index
- 16. SQL dialect comparison matrix (developer-relevant)
- 17. Decision matrices
- 18. Developer code patterns
- 19. CDC tooling and schema registries
- 20. Cross-cutting developer checklist

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
