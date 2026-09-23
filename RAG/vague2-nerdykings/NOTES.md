# NOTES — corpus `vague2-nerdykings`

## Génération

Mode **files** (`RAG/_tools/build_rag.py`) : **une fiche = un chunk**, copié verbatim de
`docs/RAG/Collect RAG Vague 2/03_nerdykings` (34 fiches). Source sans ancre HTML : navigation via `INDEX.md` et
`manifest.json`. Les fichiers de scaffolding (`_TEMPLATE.md`, `00_INDEX.md`,
`sources*.txt`) sont exclus (hors dossier source / non `.md` / préfixe `_`).

## Contenu

Site **NerdyKings** — collecte « Vague 2 » (template : Metadata + Full summary + Key points +
Technical data + Why it matters), rédigée le 2026-09-23. Fiches citées :

- ARC-AGI-3
- Bloome
- L'inexplicable découverte cachée dans Claude
- Claude Mythos
- 650 Échecs Plus Tard, Claude Fait Une Découverte Sur Riemann
- Code World Model
- 28 autres fiches (voir `INDEX.md`).
## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (1 fiche = 1 chunk,
sha256 de la fiche == sha256 du chunk, corps verbatim).

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
