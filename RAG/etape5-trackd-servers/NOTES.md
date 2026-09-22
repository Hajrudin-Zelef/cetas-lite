# NOTES — corpus `etape5-trackd-servers`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**. La source contient
5 titres H1 : un entête de document (`00-front-matter`) puis quatre rapports. Un dossier
par rapport.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 5 — Track D : serveurs IA, marché et réseau datacenter » (2026).

- Wave 1 : Supermicro + Dell ;
- Wave 2 : HPE + Giga Computing (Gigabyte) + ASUS + MSI + xFusion ;
- réseau datacenter pour l'IA ;
- chiffres du marché des serveurs IA (brouillon de recherche).

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
