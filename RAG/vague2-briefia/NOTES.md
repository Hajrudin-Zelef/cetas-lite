# NOTES — corpus `vague2-briefia`

## Génération

Mode **files** (`RAG/_tools/build_rag.py`) : **une fiche = un chunk**, copié verbatim de
`docs/RAG/Collect RAG Vague 2/01_briefia` (1 fiches). Source sans ancre HTML : navigation via `INDEX.md` et
`manifest.json`. Les fichiers de scaffolding (`_TEMPLATE.md`, `00_INDEX.md`,
`sources*.txt`) sont exclus (hors dossier source / non `.md` / préfixe `_`).

## Contenu

Site **Briefia** — collecte « Vague 2 » (template : Metadata + Full summary + Key points +
Technical data + Why it matters), rédigée le 2026-09-23. Fiches citées :

- OpenAI classe GPT-6 Astra « Critique » en cybersécurité

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (1 fiche = 1 chunk,
sha256 de la fiche == sha256 du chunk, corps verbatim).

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
