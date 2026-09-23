# NOTES — corpus `etape6-phasee4-validation-cicd`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

La source contient 4 titres H1 : un dossier par document H1 (un entête éventuel devient `00-front-matter`).

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Phase E4 — Network Validation, Observability & CI/CD` — 782 lignes, 14 chunks.

- E4.1 — Batfish: open-source network configuration analysis
- E4.2 — SuzieQ: multi-vendor network observability
- E4.3 — IP Fabric: automated network assurance platform
- E4.4 — Forward (formerly Forward Networks): network digital twin
- E4.5 — Cisco Crosswork and other commercial assurance
- E4.6 — Lab and emulation environments
- E4.7 — Testing frameworks
- E4.8 — CI/CD and GitOps for networks
- E4.9 — Event-driven automation
- E4.10 — Selection guidance (synthesis)
- E4.11 — Open items and gaps log
- E4.12 — Source index (verbatim URLs)
- E4.13 — Deep-dive expansions
- E4.14 — Verification
- E4.15 — Reference tables (detail)
- E4.16 — Final additions

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
