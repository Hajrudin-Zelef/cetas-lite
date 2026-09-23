# NOTES — corpus `vague2-datacamp`

## Génération

Mode **files** (`RAG/_tools/build_rag.py`) : **une fiche = un chunk**, copié verbatim de
`docs/RAG/Collect RAG Vague 2/02_datacamp` (68 fiches). Source sans ancre HTML : navigation via `INDEX.md` et
`manifest.json`. Les fichiers de scaffolding (`_TEMPLATE.md`, `00_INDEX.md`,
`sources*.txt`) sont exclus (hors dossier source / non `.md` / préfixe `_`).

## Contenu

Site **DataCamp** — collecte « Vague 2 » (template : Metadata + Full summary + Key points +
Technical data + Why it matters), rédigée le 2026-09-23. Fiches citées :

- Qu'est-ce qu'un agent harness ? Comment les agents d'IA obtiennent des
- Les 30 questions et réponses les plus fréquentes lors d'entretiens d'e
- Frameworks d’agents IA
- Qu'est-ce que Git ? - le guide du débutant sur le contrôle de version 
- Les meilleurs agents IA en 2026
- Les meilleures alternatives à Cursor en 2026, classées par flux de tra
- 62 autres fiches (voir `INDEX.md`).
## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (1 fiche = 1 chunk,
sha256 de la fiche == sha256 du chunk, corps verbatim).

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
