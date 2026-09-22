# RAG — corpus de référence Cetas

Corpus documentaires découpés pour la récupération (RAG) : chaque corpus est une
**partition exacte** de sa source Markdown, découpée sur les frontières de titres,
avec un index et un manifeste.

## Organisation

```
RAG/
├── README.md              # ce fichier (conventions)
├── INDEX.md               # index global (tous les corpus)
├── manifest.json          # inventaire global
├── _tools/
│   ├── build_rag.py       # générateur (déterministe, rejouable, multi-corpus)
│   └── verify_rag.py      # vérificateur indépendant (tous les corpus)
├── briefing-ia-2026/      # corpus n°1 — AI News 2026 (126 chunks)
│   ├── INDEX.md           # navigation rapide (domaines, tâches, acteurs, dates, ancres)
│   ├── manifest.json      # inventaire machine-readable
│   ├── NOTES.md           # défauts de source (à ne pas corriger) + régénération
│   └── <dossiers>/*.md    # chunks avec en-tête YAML
├── briefing-general-tech-2026/   # corpus n°2 — General Tech News 2026 (112 chunks)
├── ai-industry-kb-2026/          # corpus n°3 — AI Industry KB 2026 (123 chunks)
└── ai-industry-kb-2026-wave6/    # corpus n°4 — AI Industry KB Wave 6 delta (136 chunks)
```

Corpus n°1–2 : découpe « explicite » par H3 (listes de mapping figées).
Corpus n°3–4 : découpe « auto » (partition par titres H1→H2→H3, cibles de taille), les
sources ne portant aucune ancre.

## Garanties

- **Aucune perte** : les plages de lignes de tous les chunks couvrent `[1, n]` sans trou
  ni chevauchement ; la concaténation des corps est identique à la source.
- **Aucune réécriture** : le texte source est copié verbatim, titres d'origine conservés.
  Le seul ajout est un en-tête YAML délimité (`--- … ---`) et un H1 de titre.
- **Traçabilité** : chaque chunk porte `source`, `source_lines`, `source_anchor` et
  `sha256` (du corps source).

## Conventions d'en-tête

```yaml
id: <corpus>/<dossier>/<slug>
title: "…"
domain: <dossier>
role: timeline | deep-dive | reference | appendix
task: <intention>
actors: [ … ]
dates: [ YYYY-MM-DD, … ]
keywords: [ … ]
source: docs/RAG/<fichier>
source_anchor: "#sNN-M"
source_lines: [start, end]
canonical_for: [ … ]      # optionnel : version de fond d'un événement répété
delta_of: <corpus>        # optionnel : ce chunk est un delta d'un corpus de base
sha256: <hex>
```

## Relations entre corpus

Deux corpus peuvent se recouvrir **thématiquement** sans se recouvrir **textuellement**
(cas `ai-industry-kb-2026` → `ai-industry-kb-2026-wave6`). Pour que la recherche
arbitre, la relation est explicite :

- `relationship: base` (manifest de corpus) — corpus de fond, source canonique.
- `relationship: delta` + `delta_of: <base>` (manifest de corpus) — n'apporte que les
  faits nouveaux/corrigés ; ne remplace pas la base.
- `delta_of: <base>` (en-tête de **chaque** chunk du corpus delta) — permet de filtrer
  ou de pondérer au niveau chunk.

Le vérificateur contrôle que tout corpus `delta` référence une base existante, que tous
ses chunks portent `delta_of`, et qu'aucun de ses chunks n'est un doublon **texte**
(sha256) d'un chunk de la base.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py     # (re)génère les chunks + INDEX + manifest
python3 RAG/_tools/verify_rag.py    # vérification indépendante (doit finir « ALL CHECKS PASSED »)
```

Le générateur est **idempotent** : relancer écrase les fichiers générés mais préserve
`README.md` et `NOTES.md`. Ajouter un corpus = ajouter sa source dans `docs/RAG/` puis
étendre la table `CHUNKS` de `build_rag.py`.

## Ajouter un corpus

1. Placer la source dans `docs/RAG/<fichier>.md` (intacte).
2. Définir dans `build_rag.py` la table `(dossier, slug, ligne_du_titre, titre)`.
3. Lancer `build_rag.py` puis `verify_rag.py`.
