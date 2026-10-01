---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-3
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei", "Microsoft"]
dates: []
keywords: ["copilot", "embeddings", "mcp"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [378, 567]
sha256: 8638015a9d3861382337ae8e9b5f99e5fe3927f27f9187f756d4bea5d560fab4
---

# Outils dev + ingénierie RAG (chunk & corpus)

```text
rag-perso/
├── README.md                    # vision, quickstart, état (voir §19)
├── pyproject.toml               # dépendances + config ruff (voir §20)
├── requirements.txt             # ou requirements-dev.txt pour la CI
├── .gitignore                   # voir §15
├── .gitattributes               # règles Git LFS, voir §16
├── .editorconfig                # fin de lignes, indentation
├── .vscode/
│   ├── settings.json            # settings partagés d'équipe (sans secrets)
│   └── mcp.json                 # serveurs MCP (sans secrets, voir §7)
├── .devcontainer/
│   └── devcontainer.json        # pour Codespaces, voir §22
├── .github/
│   ├── workflows/
│   │   └── ci.yml               # lint + tests, voir §20
│   ├── ISSUE_TEMPLATE/
│   │   └── collecte.md          # template suivi de collecte, voir §18
│   └── instructions/
│       └── python.instructions.md  # consignes pour Copilot (§2)
├── scripts/
│   ├── collect_failed.py        # tes scripts de collecte existants
│   ├── collect_pw.py
│   ├── collect_still.py
│   ├── collect_final.py
│   ├── chunk_md.py              # chunker markdown (§85)
│   ├── chunk_fixed.py           # chunker fixe (§86)
│   ├── dedup.py                 # dédup exacte + MinHash (§95-97)
│   ├── embed_corpus.py          # embeddings batch (bloc final, §131+)
│   ├── ingest_pgvector.py       # ingestion (bloc final, §131+)
│   └── query_rag.py             # interrogation (§123)
├── sql/
│   ├── 001_schema.sql           # tables + extension vector (§64)
│   └── 002_indexes.sql          # HNSW / IVFFlat (§65-66)
├── corpus/
│   ├── README.md                # d'où viennent les données, licence
│   ├── manifest.json            # versioning du corpus (§98)
│   └── echantillon/             # petit échantillon versionné (tests)
├── tests/
│   ├── test_chunk_md.py
│   ├── test_dedup.py
│   └── data/                    # fixtures minuscules
├── docs/
│   ├── pipeline.md              # schéma de ton pipeline
│   └── decisions.md             # ADR légers : pourquoi tel chunk size…
└── .code-profile                # profil VS Code exporté (§4)
```

**Principes :**
- `scripts/` = que du code. `corpus/` = que des données (+ manifeste).
- `sql/` versionné = ton schéma est reproductible (`psql -f`).
- `tests/` dès le début, même 3 tests : c'est ce que la CI exécutera.

## 14. Git de base pour ton usage quotidien

```bash
# ── Premier commit ──────────────────────────────────────────
git init -b main
git add README.md .gitignore
git commit -m "Initial commit : structure projet RAG"

# ── Cycle quotidien ─────────────────────────────────────────
git status -sb                    # vue courte : branche + fichiers modifiés
git diff                          # relire AVANT de committer
git add scripts/chunk_md.py
git commit -m "chunk_md : chunker markdown par sections + métadonnées"

# ── Branches : une branche par chantier ─────────────────────
git switch -c feat/chunk-semantique   # crée + bascule
# ... travail ...
git switch main
git merge feat/chunk-semantique       # fusionne
git branch -d feat/chunk-semantique   # nettoie

# ── Historique utile ────────────────────────────────────────
git log --oneline --graph -15
git show <sha> --stat
git restore <fichier>             # annule modif non committée (Git ≥ 2.23)

# ── Synchroniser ────────────────────────────────────────────
git pull --rebase                 # rejoue tes commits sur le distant
git push
```

**Convention de messages** (lisible dans `git log --oneline`) :

```text
<type> : <quoi> (<où>)

feat    : nouvelle fonctionnalité      ex. "feat : chunker sémantique (scripts/)"
fix     : correction de bug            ex. "fix : overlap off-by-one dans chunk_fixed"
docs    : documentation                ex. "docs : README quickstart Codespaces"
chore   : maintenance                  ex. "chore : bump ruff 0.x"
data    : changement de corpus/manifeste (pas de gros binaires !)
```

## 15. .gitignore : ne jamais committer les gros fichiers de corpus

```gitignore
# ── Python ─────────────────────────────────────────────────
__pycache__/
*.py[cod]
.venv/
.venv*/
*.egg-info/
.pytest_cache/
.mypy_cache/
.ruff_cache/

# ── Secrets : JAMAIS dans git ──────────────────────────────
.env
.env.*
*.pem
*.key
auth.json
**/*token*.json

# ── Corpus : données volumineuses hors git ─────────────────
corpus/brut/
corpus/embeddings/
*.zip
*.tar
*.tar.gz
data/
downloads/
*.pdf
!docs/**/*.pdf

# ── OS / éditeurs ──────────────────────────────────────────
.DS_Store
Thumbs.db
.idea/
*.swp
*~

# ── Playwright / navigateurs ───────────────────────────────
# (les binaires vont dans ~/.cache/ms-playwright, pas dans le repo)
```

> **Réflexe :** `git status` avant chaque commit. Si tu vois un `.zip`
> de 400 Mo, c'est que ton `.gitignore` a un trou. Pour les gros fichiers
> que tu veux **vraiment** versionner : Git LFS (§16).

## 16. Git LFS : versionner les gros fichiers proprement

Git casse au-delà de ~100 Mo par fichier (push refusé sur GitHub).
**Git LFS** stocke un pointeur dans Git, le contenu sur un stockage séparé.
Vérifié sept 2026 (commandes stables depuis des années).

```bash
# 1. Installer (une fois par machine)
sudo apt install git-lfs     # Debian/Ubuntu
git lfs install              # initialise les hooks (une fois par utilisateur)

# 2. Déclarer les motifs à tracker — CRÉE .gitattributes
git lfs track "*.zip"
git lfs track "*.tar"
git lfs track "corpus/exports/**"
git lfs track "*.onnx" "*.bin"   # modèles locaux éventuels

# 3. COMMITTER .gitattributes (obligatoire, sinon ça ne marche pas en équipe)
git add .gitattributes
git commit -m "chore : Git LFS pour zip/tar/exports corpus"

# 4. Usage ensuite = git normal
git add corpus/exports/huawei_cli_ref_clean.tar
git commit -m "data : export CLI Huawei nettoyé"
git push   # le pointeur part sur GitHub, le contenu sur le stockage LFS
```

Commandes utiles :

```bash
git lfs ls-files            # fichiers trackés LFS
git lfs ls-files --size     # + tailles (surveille ton quota)
git lfs pull                # récupère les contenus LFS
git lfs fetch               # récupère sans checkout
git lfs prune               # nettoie le cache local
# Cloner SANS télécharger les gros fichiers :
GIT_LFS_SKIP_SMUDGE=1 git clone <url>
```

**Quotas GitHub (à vérifier — non confirmés par mes recherches) :**
chaque compte aurait ~1 Go de stockage LFS + 1 Go de bande passante/mois
en gratuit. **Stratégie pour toi :**
- ✅ LFS : exports de référence stables (ex. un `.tar` de CLI nettoyés que
  tu veux figer comme « v1 »).
- ❌ Pas LFS : corpus brut qui change chaque semaine, embeddings
  régénérables (ça se recalcule, ça ne se versionne pas), logs.

