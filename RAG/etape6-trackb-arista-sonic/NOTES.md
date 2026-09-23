# NOTES — corpus `etape6-trackb-arista-sonic`

## Génération

Mode auto (`RAG/_tools/build_rag.py`) avec `first_is_content: true` (H1 unique = contenu),
cibles **45–90 lignes**. Dossier `00-arista-sonic` (`folder_name`), fichiers par titre H2.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 6 — Track B : Arista + SONiC + Cumulus (fabric datacenter & réseau ouvert) » (2026).

- Arista Networks ;
- SONiC (Software for Open Networking in the Cloud) ;
- NVIDIA Cumulus Linux ;
- paysage du réseau ouvert 2026 ;
- positionnement concurrentiel Arista vs NVIDIA Spectrum ;
- points de vérification ouverts, métadonnées, passe de recherche supplémentaire
  (second passage indépendant, 2026-09-22).

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
