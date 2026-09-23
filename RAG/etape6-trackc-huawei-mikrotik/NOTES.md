# NOTES — corpus `etape6-trackc-huawei-mikrotik`

## Génération

Mode auto (`RAG/_tools/build_rag.py`) avec `first_is_content: true` (H1 unique = contenu),
cibles **45–90 lignes**. Dossier `00-huawei-mikrotik` (`folder_name`), fichiers par titre H2.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 6 — Track C : Huawei + MikroTik (matériel réseau) » (1er janvier → 22 septembre 2026).

- Huawei : portefeuille réseau 2026 ;
- MikroTik : nouveautés produits 2026 ;
- tendances de prix des switches 25G / 100G / 400G ;
- OS réseau open source (bref) ;
- journal de vérification ouverte, métadonnées de collecte.

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source).

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
