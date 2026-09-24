# NOTES — corpus `collect-240926-misc`

## Génération

Mode **files** (`RAG/_tools/build_rag.py`) : **une source = un chunk**, copie verbatim.

- source : `docs/RAG/clean_en/misc` (29 fichiers)
- dossier interne : `misc` · domaine : unsloth (6), blogdumoderateur (4), opencode (4), orcarouter (3), artificialanalysis (2), benchlm (2), numerama (1), lesnumeriques (1), koul (1), poyo (1), commandcode (1), briefia (1), amd (1), simonwillison (1)
- total : ~39,923 mots

## Contenu

longue traîne — 14 sites isolés : unsloth, opencode, blogdumoderateur, orcarouter, benchlm, artificialanalysis, amd, poyo, koul, numerama, simonwillison, briefia, lesnumeriques, commandcode.

## Titres

20 fichier(s) sur 29 n'ont pas de titre `# H1` dans la source : le titre
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
