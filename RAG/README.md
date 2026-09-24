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
├── ai-industry-kb-2026-wave6/    # corpus n°4 — AI Industry KB Wave 6 delta (136 chunks)
├── frontier-models-2026/         # corpus n°5 — Frontier AI Models, Vague 1 (12 chunks)
├── labs-grok-platforms-2026/     # corpus n°6 — Labs, Grok & platforms, Vague 2 (17 chunks)
├── open-local-models-2026/       # corpus n°7 — Open / Local AI Models (26 chunks)
├── tools-platforms-2026/         # corpus n°8 — AI Tools & Platforms, Step 2 (8 chunks)
├── labs-hyperscalers-2026/       # corpus n°9 — Labs & Hyperscalers (47 chunks)
│
├── etape4-*/                     # corpus n°10-13 — inférence & outillage (4 corpus)
├── etape5-*/                     # corpus n°14-17 — silicium & serveurs (4 corpus)
├── etape6-*/                     # corpus n°18-36 — réseau & sécurité (19 corpus)
├── etape7-*/                     # corpus n°37-45 — OS, containers, stockage, bases (9 corpus)
├── etape8-*/                     # corpus n°46-51 — langages & frameworks (6 corpus)
├── etape9-*/                     # corpus n°52-56 — stockage & mémoire (5 corpus)
├── etape10-*/                    # corpus n°57-63 — actualités 2026 (7 corpus)
├── collect-korben/                # corpus n°64 — Korben.info (18 fiches, mode files)
├── collect-mindstudio/            # corpus n°65 — MindStudio (170 fiches, mode files)
├── collect-huggingface/           # corpus n°66 — Hugging Face (125 fiches, mode files)
├── collect-opencode-docs/         # corpus n°67 — opencode docs (4 fiches, mode files)
├── collect-presse-fr/             # corpus n°68 — Presse FR (8 fiches, mode files)
├── collect-benchmarks/            # corpus n°69 — Benchmarks (5 fiches, mode files)
├── collect-tutoriels/             # corpus n°70 — Tutoriels & reviews (6 fiches, mode files)
├── vague2-briefia/                # corpus n°71 — Briefia (1 fiche, mode files)
├── vague2-datacamp/               # corpus n°72 — DataCamp (68 fiches, mode files)
├── vague2-nerdykings/             # corpus n°73 — NerdyKings (34 fiches, mode files)
├── vague2-vision-ia/              # corpus n°74 — Vision-IA (55 fiches, mode files)
├── collect-240926-mindstudio/     # corpus n°75 — MindStudio (170 fichiers, mode files)
├── collect-240926-huggingface/    # corpus n°76 — Hugging Face (125 fichiers, mode files)
├── collect-240926-storagereview/  # corpus n°77 — StorageReview (74 fichiers, mode files)
├── collect-240926-datacamp/       # corpus n°78 — DataCamp (69 fichiers, mode files)
├── collect-240926-tomshardware/   # corpus n°79 — Tom's Hardware (68 fichiers, mode files)
├── collect-240926-vision-ia/      # corpus n°80 — Vision-IA (55 fichiers, mode files)
├── collect-240926-nerdykings/     # corpus n°81 — NerdyKings (34 fichiers, mode files)
├── collect-240926-korben/         # corpus n°82 — Korben (18 fichiers, mode files)
├── collect-240926-frandroid/      # corpus n°83 — FrAndroid (14 fichiers, mode files)
├── collect-240926-hardwarecooking/ # corpus n°84 — HardwareCooking (12 fichiers, mode files)
└── collect-240926-misc/           # corpus n°85 — longue traîne (29 fichiers, 14 sites)
```

Détail par étape (séries `etape*`) : chaque source de `docs/RAG/` donne **un corpus**
autonome. Liste complète et à jour : `INDEX.md` (racine) et `manifest.json`.
Séries `collect-*` / `vague2-*` / `collect-240926-*` : dossiers de fichiers individuels,
**un fichier = un chunk** (mode `files`, voir plus bas). Total **85 corpus / 2 487 chunks**.

La vague `collect-240926` (668 articles scrapés, `docs/RAG/clean_en/`) est découpée
**par site** — un dossier par site ≥ 10 fichiers, le reste regroupé en `misc` (14 sites).
La facette `domain` porte le vrai site de chaque chunk, y compris ceux de `misc`
(lue dans le marqueur `<!-- source: URL -->` de la source).

### Trois modes de découpe

| mode | source | découpe | cas d'usage |
|---|---|---|---|
| **explicite** | fichier unique | mapping figé (dossier, slug, ligne) | briefings annotés avec ancres |
| **auto** | fichier unique | partition par titres H1→H2→H3, cibles de taille | fiches de recherche longues, sans ancre |
| **files** | **dossier de fiches** | 1 fichier = 1 chunk (verbatim) | collecte organisée en fiches courtes (template) |

Mode **files** : un corpus pointe vers un `source_dir` (dossier de fiches `.md`).
Chaque fiche est copiée verbatim dans un chunk ; `sha256` vérifie chaque copie.
`task` est dérivée du champ `Type` de la fiche (`article`, `model-card`, `benchmark`,
`documentation`, `tutorial`, `review`). Les fichiers `_*.md` (template, index) sont
exclus automatiquement.

Corpus n°1–2 : découpe « explicite » par H3 (listes de mapping figées).
Corpus n°3–63 : découpe « auto » (partition par titres H1→H2→H3, cibles de taille), les
sources ne portant aucune ancre. Cibles 70–170 lignes pour n°3–4 (fichiers de 10 000–
14 000 lignes), 50–110 pour n°5–8 (500–1 100 lignes) et 45–90 pour n°9–63 (48–3 000 lignes).

Options du mode auto : `first_is_content` (le H1 unique est du contenu, pas un entête),
`folder_name` (nom court du dossier pour ce cas), et rattachement automatique des titres H1
« séparateurs » (ex. `# PART 1 — vLLM`) au bloc suivant.

## Garanties

- **Aucune perte** : les plages de lignes de tous les chunks couvrent `[1, n]` sans trou
  ni chevauchement ; la concaténation des corps est identique à la source.
  Mode `files` : chaque fiche du dossier a exactement un chunk, sha256 correspondant.
- **Aucune réécriture** : le texte source est copié verbatim, titres d'origine conservés.
  Le seul ajout est un en-tête YAML délimité (`--- … ---`) et un H1 de titre.
- **Traçabilité** : chaque chunk porte `source`, `source_lines`, `source_anchor` et
  `sha256` (du corps source).

## Périmètre des sources

Toutes les sources de corpus vivent sous **`docs/RAG/`**. Les documents de l'application
— **`README.md` (racine) et `docs/index.md`** — ne sont **jamais** des sources de corpus :
ils sont privés et ne doivent pas être utilisés pour répondre. Ne pas les déposer dans
`$CETAS_LITE_HOME/rag/` non plus.

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

## Intégration runtime (cetas)

cetas charge ce corpus dans un **index local en mémoire** (`internal/rag`) et l'utilise
pour répondre avant d'aller sur le web. Aucun service externe, aucun réseau.

- **Emplacement** : `$CETAS_LITE_HOME/rag/` (défaut) ou `CETAS_LITE_RAG_DIR`.
  Le dossier est créé au démarrage ; **vide ⇒ RAG inactif, coût nul**.
- **Format riche** : y déposer les dossiers de corpus générés (`<corpus>/manifest.json`
  + chunks), exactement la sortie de `build_rag.py`.
- **Format brut** : un dossier sans `manifest.json` est lu au plus simple — **un fichier
  `.md`/`.txt` = un chunk** (titre = premier `#`, sinon nom du fichier). Pratique pour des
  notes perso : déposer les fichiers, rien d'autre à faire.
- **Au runtime** : les extraits pertinents sont injectés automatiquement dans le tour ;
  si la base couvre la requête, la pré-recherche web est **sautée** (économie de coût et
  de latence). Outils agent `rag_search` / `rag_read` pour creuser.
- **Sûreté** : index immuable, chargement en tâche de fond, recherche bornée à 120 ms,
  fail-open (une erreur RAG ne bloque jamais la réponse).

Vérifier au démarrage : le log affiche `rag: N chunks, M corpus, K termes, X Mo de texte en …`
ou `rag: aucun corpus … (inactif)`.

## Ressources (mesuré)

L'index est en mémoire et **proportionnel au texte des corpus** ; il ne dépend pas du
nombre de requêtes (index immuable, aucune écriture, aucun cache qui grandit).

| mesure | 21 corpus | 63 corpus | 70 corpus | 74 corpus | **85 corpus** |
|---|---|---|---|---|---|
| texte des corpus | 6,4 Mo | 10,4 Mo | 12,1 Mo | 13,0 Mo | **21,6 Mo** |
| chunks / termes | 707 / 25 789 | 1 325 / 41 707 | 1 661 / 43 577 | 1 819 / 45 319 | **2 487 / 56 375** |
| **mémoire heap** | 32 Mo | 71 Mo (6,8×) | 69 Mo (5,7×) | 84 Mo (6,5×) | **140 Mo** (6,5×) |
| construction (arrière-plan) | 0,3 s | 0,7 s | 0,7 s | 0,75 s | **1,2 s** |
| recherche (BM25, plafond 120 ms) | 207 µs | 302 µs | 296 µs | 428 µs | **571 µs** |

Ordres de grandeur : ~1 Mo de corpus ⇒ ~7 Mo de heap. Le service complet (RAG inclus)
est passé à **~140 Mo de heap** (RSS mesurée à suivre au démarrage). Sur un VPS sans swap,
prévoir la marge : un corpus de 100 Mo de texte demanderait ~0,7 Go de heap. Au-delà,
réduire les cibles de taille (`max_lines`) ou ne charger que les corpus utiles.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py     # (re)génère les chunks + INDEX + manifest
python3 RAG/_tools/verify_rag.py    # vérification indépendante (doit finir « ALL CHECKS PASSED »)
```

Le générateur est **idempotent** : relancer écrase les fichiers générés mais préserve
`README.md` et `NOTES.md`. Ajouter un corpus = ajouter sa source dans `docs/RAG/` puis
une entrée dans la table `CORPORA` de `build_rag.py`.

## Ajouter un corpus

1. Placer la source dans `docs/RAG/` :
   - un **fichier** `.md` (intact) ;
   - ou un **dossier de fiches** `docs/RAG/<mon-dossier>/` (template de fiche courant).
2. Ajouter une entrée dans `CORPORA` (`build_rag.py`) :
   - source **fichier unique, sans ancre** → `"mode": "auto"` + `max_lines`/`min_lines`
     (et `first_is_content: True` si le fichier n'a qu'un seul H1 qui est du contenu) ;
   - source **fichier unique, avec ancres** → table de chunks explicite
     `(dossier, slug, ligne, titre)` ;
   - source **dossier de fiches** → `"mode": "files"` + `source_dir` + `folder`
     (une fiche = un chunk ; `task` déduite du champ `Type`).
3. Lancer `build_rag.py` puis `verify_rag.py`.
