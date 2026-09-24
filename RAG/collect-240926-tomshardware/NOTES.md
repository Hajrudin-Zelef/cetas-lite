# NOTES — corpus `collect-240926-tomshardware`

## Génération

Mode **files** (`RAG/_tools/build_rag.py`) : **une source = un chunk**, copie verbatim.

- source : `docs/RAG/clean_en/tomshardware` (68 fichiers)
- dossier interne : `tomshardware` · domaine : tomshardware (100 %)
- total : ~121,403 mots

## Contenu

Tom's Hardware — CPU, GPU, PC, hardware grand public et serveurs.

## Titres

57 fichier(s) sur 68 n'ont pas de titre `# H1` dans la source : le titre
YAML retombe sur le nom de fichier (lisible). **Aucun réécriture** du texte source.

## Provenance

`source:` dans l'en-tête YAML = chemin dans `docs/RAG/` (espace perso, gitignoré).
Cette ligne est une **référence morte** à l'exécution : le runtime ne la lit jamais
(strip du front-matter au chargement), seule `verify_rag.py` la relit à la génération.

## Non audité

Aucun défaut de source recensé ; 10 fichiers restés en FR sont laissés verbatim
(décision utilisateur).

## Vérification

`verify_rag.py` : `ALL CHECKS PASSED` — correspondance 1:1 fichier ↔ chunk,
sha256, corps verbatim.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
