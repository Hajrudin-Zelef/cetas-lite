# NOTES — corpus `briefing-general-tech-2026`

## Défauts de la source (conservés verbatim — NE PAS corriger ici)

La découpe est strictement fidèle : les défauts ci-dessous existent dans la source
`docs/RAG/briefing-general-tech-2026-en.md` et sont **recopiés tels quels**. Toute
correction doit être faite à la source puis le corpus régénéré.

1. **Lignes 7245–7247** — phrase tronquée : « And 201 grams puts the Fold 8 Samsung has
   effectively erased much of the weight penalty… » (mots manquants entre « Fold 8 » et
   « Samsung »).
2. **Ligne 7729** — clause dupliquée : « …as of September 2026, no vendor has volunteered
   that number as of September 2026. »
3. **Table des matières locale du chapitre 7** — entrées avec point final (`7.1.`, `7.2.`,
   …, `7.13.`) alors que les titres réels et la TOC globale n'en ont pas (`7.1`).
4. **Double « Detailed table of contents »** — une pour le dossier (ligne 402) et une pour
   le chapitre 7 (ligne 6649) ; les deux sont conservées dans leurs chunks respectifs.
5. **Titres sans ancre** (frontières non ancrées) : H1 384, H2 385, H2 402, H1 6642,
   H2 6649, H3 « Chapter synthesis » 6487, H3 « Chapter contents » 7920, et les trois H3
   d'ouverture de chapitre 3 (1345, 1349, 1358). Le générateur démarre alors le chunk à la
   ligne du titre.
6. **Ancre inline** — `g06` est collée en fin de ligne 5587 (dernier paragraphe de §5.15),
   et non sur sa propre ligne. Elle reste donc dans le fichier
   `04-training-inference-quantization/15-what-quantization-changes.md` ; la table
   « Ancres source → fichier » de `INDEX.md` fait foi.
7. **Ancre détachée** — `g09-2` (ligne 9241) est séparée de son titre `### 9.2` (ligne 9243)
   par une ligne vide ; le chunk commence à l'ancre (9241).

## Contenu conservé même s'il est redondant

Le dossier est **volontairement redondant** : chaque mois de la chronologie (§2) se termine
par un tableau qui reformule son propre récit ; la section 3.19 (« Competitive landscape »)
récapitule toutes les sections 3.1–3.18 ; les annexes (glossaire, index, notes
méthodologiques) recyclent les définitions et faits du corps. **Rien n'a été supprimé.**
La table « Événements canoniques » (`INDEX.md`) indique la version de fond (`canonical_for`).

## Métadonnées

`actors`, `dates` et `keywords` sont **auto-dérivés** par `RAG/_tools/build_rag.py`
(listes contrôlées + extraction regex), à titre d'aide au filtrage — non exhaustifs.
`task` et `canonical_for` proviennent de tables explicites du générateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
