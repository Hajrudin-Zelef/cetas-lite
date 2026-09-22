# NOTES — corpus `ai-industry-kb-2026`

## Génération

Ce corpus a été produit en **mode auto** (`RAG/_tools/build_rag.py`, clé `mode: auto`) :

- partition par titres **H1 → H2 → H3**, sans jamais couper au milieu d'un paragraphe ;
- cibles de taille : 70–170 lignes par chunk (fusion des petits blocs, découpe des gros
  aux H3, et à défaut sur lignes vides pour les listes) ;
- un dossier par section H1 (`00-front-matter`, `01-…`, …, `23-annex-a-…`) ; un fichier par
  bloc (nommé d'après son premier titre, `overview` pour l'entête de section).

La source ne contient **aucune ancre** HTML (`<a id=…>`) : l'index n'a donc pas de table
ancre→fichier. La navigation repose sur `INDEX.md` (domaines, tâches, acteurs, dates) et
`manifest.json`.

## Contenu

- Base de connaissances consolidant 6 vagues de corpus. Chaque section suit un template
  (`Summary`, `Key dated facts`, `Figures and metrics`, `Main actors`, `Timeline and
  context`, `Implications`, `Sources and URLs`), parfois dupliqué en blocs « (continued) ».
- Les **labels de provenance** (`[VENDOR]`, `[PRIMARY]`, `[VERIFIED]`, `[SECONDARY]`,
  `[COMMUNITY]`, `[UNVERIFIED]`, `[DIRECTIONAL]`, `[EDITORIAL]`…) sont dans le texte et
  conservés verbatim.
- Les annexes B (index global) et C (liste de mots-clés) sont des listes monolithiques sans
  sous-titres : elles sont découpées sur lignes vides (`overview.md`, `part-2.md`, …).

## Non audité

Contrairement aux corpus `briefing-*`, le contenu de ce fichier n'a **pas** été audité
section par section. Aucun défaut de source n'est donc recensé ici ; la fidélité est
garantie par le vérificateur (concaténation == source), pas par une relecture sémantique.

## Redondance

Le corpus est redondant par construction (consolidation multi-vagues). Rien n'est supprimé.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
