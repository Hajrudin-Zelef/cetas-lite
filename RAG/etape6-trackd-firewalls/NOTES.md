# NOTES — corpus `etape6-trackd-firewalls`

## Génération

Mode auto (`RAG/_tools/build_rag.py`) avec `first_is_content: true` (H1 unique = contenu),
cibles **45–90 lignes**. Dossier `00-firewalls` (`folder_name`), fichiers par titre H2.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 6 — Track D : FortiGate + pfSense + OPNsense (pare-feu & sécurité réseau) »
(février → 22 septembre 2026).

- Fortinet / FortiGate ;
- pfSense (Netgate), OPNsense (Deciso) ;
- marché pare-feu / sécurité réseau 2026 ;
- CVE et avis notables affectant ces éditeurs en 2026 ;
- zero-trust / SASE 2026 (bref) ;
- points de vérification ouverts, métadonnées de collecte ;
- passe de recherche supplémentaire (second agent, 2026-09-22).

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
