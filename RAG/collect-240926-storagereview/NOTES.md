# NOTES — corpus `collect-240926-storagereview`

## Génération

Mode **files** (`RAG/_tools/build_rag.py`) : **une source = un chunk**, copie verbatim.

- source : `docs/RAG/clean_en/storagereview` (74 fichiers)
- dossier interne : `storagereview` · domaine : storagereview (100 %)
- total : ~196,866 mots

## Contenu

StorageReview — SSD, stockage entreprise, serveurs et infra IA.

## Titres

71 fichier(s) sur 74 n'ont pas de titre `# H1` dans la source : le titre
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
