# NOTES — corpus `etape6-phased2-evpn-vxlan`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-evpn-vxlan` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 6 — Phase D wave 2: EVPN-VXLAN overlay` — 879 lignes, 16 chunks.

- Wave 1 — VXLAN encapsulation fundamentals
- Wave 2 — EVPN control plane
- Wave 3 — EVPN multihoming
- Wave 4 — DCI: EVPN multi-site, OTV comparison
- Wave 5 — Vendor implementations and interop
- Wave 6 — Troubleshooting and operations
- Coverage audit and open items

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
