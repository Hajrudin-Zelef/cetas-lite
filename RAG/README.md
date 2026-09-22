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
│   ├── build_rag.py       # générateur (déterministe, rejouable)
│   └── verify_rag.py      # vérificateur indépendant
└── briefing-ia-2026/      # corpus n°1
    ├── INDEX.md           # navigation rapide (domaines, tâches, acteurs, dates, ancres)
    ├── manifest.json      # inventaire machine-readable
    ├── NOTES.md           # défauts de source (à ne pas corriger) + régénération
    └── <dossiers>/*.md    # chunks avec en-tête YAML
```

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
sha256: <hex>
```

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
