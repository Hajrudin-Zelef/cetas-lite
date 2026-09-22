# NOTES — corpus `ai-industry-kb-2026-wave6`

## Génération

Produit en **mode auto** (`RAG/_tools/build_rag.py`, clé `mode: auto`) : partition par
titres **H1 → H2 → H3**, cibles 70–170 lignes par chunk, un dossier par section H1
(`00-front-matter`, `01-compute-and-capital-deals`, …, `26-annex-c-keyword-list`), un
fichier par bloc.

La source ne contient **aucune ancre** HTML : pas de table ancre→fichier dans `INDEX.md`.

## Nature du document

**Delta** par rapport au corpus `ai-industry-kb-2026` : ce volume ne contient que les faits
nouveaux ou corrigés (consolidation de 5 fichiers de vague 6, vérifiés le 2026-09-22). Il
conserve les renvois « covered in main KB §X » et les **labels de provenance**
(`[VENDOR]`, `[SECONDARY]`, `[COMMUNITY]`, `[UNVERIFIED]`, `[DIRECTIONAL]`,
`[CONTRADICTED]`).

Relation encodée dans les métadonnées : `relationship: delta` + `delta_of:
ai-industry-kb-2026` au niveau corpus, et `delta_of: ai-industry-kb-2026` dans l'en-tête de
**chaque** chunk. Vérifié : **aucun** chunk n'est un doublon texte (sha256) d'un chunk de la
base — les deux corpus se recouvrent thématiquement, pas textuellement.

Les sections suivent le même template que le KB principal (`Summary`, `Key dated facts`,
`Figures and metrics`, `Main actors`, `Timeline and context`, `Implications`, `Sources and
URLs`), avec des blocs « expansion » (`new verified …`). Les annexes B (index alphabétique)
et C (liste de mots-clés) sont des listes monolithiques découpées sur lignes vides.

## Détail technique

La source n'a **pas de saut de ligne final** : le fichier compte 14333 retours ligne mais
14334 lignes logiques (dernière ligne non terminée). La partition et le `sha256` portent sur
le contenu exact (`14334` lignes dans le manifest).

## Non audité

Pas d'audit sémantique section par section (contrairement aux corpus `briefing-*`) ; la
fidélité est garantie par le vérificateur (concaténation == source).

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
