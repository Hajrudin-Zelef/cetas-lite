# NOTES — corpus `vague2-vision-ia`

## Génération

Mode **files** (`RAG/_tools/build_rag.py`) : **une fiche = un chunk**, copié verbatim de
`docs/RAG/Collect RAG Vague 2/04_vision_ia` (55 fiches). Source sans ancre HTML : navigation via `INDEX.md` et
`manifest.json`. Les fichiers de scaffolding (`_TEMPLATE.md`, `00_INDEX.md`,
`sources*.txt`) sont exclus (hors dossier source / non `.md` / préfixe `_`).

## Contenu

Site **Vision-IA (beehiiv)** — collecte « Vague 2 » (template : Metadata + Full summary + Key points +
Technical data + Why it matters), rédigée le 2026-09-23. Fiches citées :

- Alibaba met un générateur d'images de 7 milliards de paramètres sur vo
- Anthropic accuse Alibaba du plus grand pillage jamais mené contre Clau
- Anthropic blacklistée par Trump ? la justice annule la décision
- Anthropic lance Fable 5.1 et réduit jusqu'à 45 % le coût de ses agents
- Anthropic prêchait la prudence, son filtre anti-armes biologiques est 
- Anthropic vs Trump
- 49 autres fiches (voir `INDEX.md`).
## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (1 fiche = 1 chunk,
sha256 de la fiche == sha256 du chunk, corps verbatim).

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
