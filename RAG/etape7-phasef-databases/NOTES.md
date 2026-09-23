# NOTES — corpus `etape7-phasef-databases`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-databases` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 7 — Phase F: Databases & Cache (Operations Angle)` — 797 lignes, 14 chunks.

- 1. PostgreSQL — releases, lifecycle, 2026 status
- 2. PostgreSQL — high availability, replication, failover
- 3. PostgreSQL — connection pooling, extensions, tuning
- 4. PostgreSQL — backup, WAL, point-in-time recovery
- 5. MySQL — releases, lifecycle, 2026 status
- 6. MySQL/MariaDB — high availability and replication
- 7. Supabase — 2026 status (ops angle)
- 8. Redis / Valkey and the in-memory ecosystem
- 9. Managed databases — pricing and positioning (2026 signals)
- 10. Monitoring, observability and slow-query work
- 11. Security hardening checklist (ops)
- 12. Decision guides (ops)
- 13. Conflicts, gaps and unverified claims
- 14. Glossary (ops terms)
- 15. Source index (verbatim URLs, access 2026-09-22)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
