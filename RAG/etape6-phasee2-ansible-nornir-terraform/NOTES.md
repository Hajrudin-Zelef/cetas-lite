# NOTES — corpus `etape6-phasee2-ansible-nornir-terraform`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

La source contient 8 titres H1 : un dossier par document H1 (un entête éventuel devient `00-front-matter`).

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries` — 755 lignes, 16 chunks.

- Wave 0 — Scope, method, provenance
- Wave 1 — ansible-core releases & lifecycle (2026)
- Wave 2 — Ansible network collections (2026 versions)
- Wave 3 — AWX, Ansible Automation Platform & Event-Driven Ansible
- Wave 4 — Nornir ecosystem
- Wave 5 — Terraform / OpenTofu for networks
- Wave 6 — Python network libraries
- Wave 7 — Cross-cutting patterns
- Wave 8 — Verification log, gaps & conflicts
- Wave 9 — Ansible networking internals & validated patterns
- Wave 10 — Nornir architecture & usage patterns
- Wave 11 — Terraform/OpenTofu deep dive for networks
- Wave 12 — Python libraries deep dive
- Wave 13 — Operational patterns & CI/CD
- Wave 14 — Selection guide: which tool when
- Wave 15 — Example workflows (placeholders only, no real secrets)
- Wave 16 — Source index (verbatim URLs)
- Wave 17 — Vendor provider & collection catalog details
- Wave 18 — 2026 release timeline (network automation)
- Wave 19 — Glossary & concept map
- Wave 20 — Collection module catalog (notable modules, 2026)
- Wave 21 — Hardening checklist for network automation (2026)
- Wave 22 — Migration playbooks
- Wave 23 — Open items carried forward
- Wave 24 — Head-to-head comparison matrices
- Wave 25 — Key numbers at a glance (2026-09-22)
- Wave 26 — FAQ (field questions answered from research)
- Wave 27 — Adoption signals & community health (2026)
- Wave 28 — Quick-reference commands (2026 tooling)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
