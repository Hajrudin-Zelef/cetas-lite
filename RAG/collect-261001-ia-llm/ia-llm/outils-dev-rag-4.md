---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-4
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei", "OpenAI"]
dates: []
keywords: ["agent", "embedding", "embeddings"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [568, 830]
sha256: 5388d419da6221c3b0af9d8829ad75cf9ee082e81267423bb02502539f45ff8b
---

# Outils dev + ingénierie RAG (chunk & corpus)

> Alternative « gros corpus » : **DVC** (Data Version Control, stocke les
> pointeurs dans Git et les données sur S3/disque/NAS). À évaluer si ton
> corpus dépasse quelques Go — voir §98.

## 17. Secrets GitHub : les clés API ne sont JAMAIS en dur

```bash
# Côté code : lis depuis l'environnement, TOUJOURS
# ❌ INTERDIT :  client = OpenAI(api_key="sk-...")
# ✅ OBLIGATOIRE :
import os
from openai import OpenAI
client = OpenAI(api_key=os.environ["OPENAI_API_KEY"])  # lève une erreur si absent
```

Enregistrer un secret (repo → **Settings → Secrets and variables →
Actions → New repository secret**), ou en CLI :

```bash
gh secret set OPENAI_API_KEY        # demande la valeur interactivement
gh secret set SUPABASE_URL --body "https://xyz.supabase.co"
gh secret list                      # liste les NOMS (jamais les valeurs)
```

Usage dans un workflow (§20) :

```yaml
env:
  OPENAI_API_KEY: ${{ secrets.OPENAI_API_KEY }}
```

**Garde-fous :**
- `.env` est dans `.gitignore` (§15). Un `git log -S "sk-"` permet de
  vérifier qu'aucune clé n'a fuité dans l'historique.
- Si une clé fuite (commit poussé) : **révoque-la immédiatement** côté
  fournisseur, ne te contente pas de la retirer du code (l'historique
  la contient toujours).
- En local, charge `.env` via `python-dotenv` (`load_dotenv()`) — jamais
  de `export` avec la clé dans ton `.bashrc` versionné.

## 18. Issues : suivre tes collectes comme des tickets

Crée un template `.github/ISSUE_TEMPLATE/collecte.md` :

```markdown
---
name: "Suivi de collecte"
about: "Suivre une passe de collecte du pipeline RAG"
title: "[collecte] "
labels: ["collecte"]
---

## Source
<!-- ex. Huawei EDOC, lot 3 -->

## Fichiers
- [ ] liste d'URL : `listes/huawei_lot3.txt` (N URL)
- [ ] passe 1 : `collect_failed.py`
- [ ] passe 2 : `collect_pw.py`
- [ ] passe 3 : `collect_still.py`
- [ ] passe 4 : `collect_final.py`
- [ ] passe 5 : cloud (Leo)

## Résultat
- Récupérées : x/N
- `failed_final.txt` : y URL restantes
- Notes :
```

```bash
# Créer / lister / fermer depuis le terminal
gh issue create --title "[collecte] Huawei EDOC lot 3" --body-file /tmp/issue.md
gh issue list --label collecte
gh issue close 12 --comment "452/452 récupérées, corpus mergé"
```

**Pourquoi c'est utile :** dans 6 mois, tu sauras exactement quelle passe a
produit quel fichier, et pourquoi il reste 14 URL en `failed_dead.txt`.

## 19. README : le modèle qui rend ton repo utilisable

```markdown
# rag-perso

Pipeline RAG personnel : collecte → nettoyage → chunking → embeddings
(`text-embedding-3-small`) → pgvector → interrogation.

## Quickstart

\`\`\`bash
python -m venv .venv && source .venv/bin/activate
pip install -r requirements.txt
cp .env.example .env   # renseigne OPENAI_API_KEY, SUPABASE_URL, SUPABASE_KEY
python scripts/chunk_md.py --input corpus/ --output build/chunks.jsonl
python scripts/embed_corpus.py --input build/chunks.jsonl
python scripts/query_rag.py --q "configurer iStack sur S310"
\`\`\`

## Architecture

collecte (scripts/collect_*.py) → nettoyage → chunk (scripts/chunk_*.py)
→ dedup (scripts/dedup.py) → embed → pgvector (sql/) → query

## État des collectes

| Lot | URL | Récupérées | Reste |
|-----|-----|------------|-------|
| Huawei EDOC | 452 | 452 | 0 |

## Conventions

- Commits : `type : objet (périmètre)` (voir docs/pipeline.md)
- Corpus brut : jamais dans git (voir .gitignore) ; manifestes dans corpus/
- Secrets : variables d'environnement uniquement
```

Un bon README contient : **quoi**, **comment lancer en 3 commandes**,
**où sont les données**, **conventions**. C'est aussi ce que lit un agent
IA en premier quand il découvre ton repo.

## 20. GitHub Actions : CI lint + tests pour tes scripts Python

**Principe (vérifié sept 2026) :** un workflow = fichier YAML dans
`.github/workflows/`. Déclencheur (`on:`) → jobs (machines `ubuntu-latest`)
→ steps (actions versionnées + commandes). Les versions ci-dessous
(`actions/checkout@v4`, `actions/setup-python@v5`) sont les plus documentées
en 2026 ; une source 2026 mentionne des majeures `@v6` — **à vérifier**
au moment où tu crées le fichier, et épingle toujours une version.

`.github/workflows/ci.yml` — **complet et fonctionnel** :

```yaml
name: CI — lint & tests

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
  workflow_dispatch:        # lancement manuel depuis l'onglet Actions

# Évite les runs qui se marchent dessus : un nouveau push annule l'ancien
concurrency:
  group: ci-${{ github.ref }}
  cancel-in-progress: true

jobs:
  lint:
    name: Lint (ruff)
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Python
        uses: actions/setup-python@v5
        with:
          python-version: "3.12"
          cache: pip                  # cache pip → builds plus rapides

      - name: Install lint deps
        run: |
          python -m pip install --upgrade pip
          pip install ruff

      - name: Ruff check
        run: ruff check scripts/ tests/

      - name: Ruff format (vérif seule)
        run: ruff format --check scripts/ tests/

  test:
    name: Tests (pytest)
    runs-on: ubuntu-latest
    needs: lint                        # ne tourne que si le lint passe
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Python
        uses: actions/setup-python@v5
        with:
          python-version: "3.12"
          cache: pip

      - name: Install deps
        run: |
          python -m pip install --upgrade pip
          pip install -r requirements.txt
          pip install pytest

      - name: Run pytest
        run: pytest -q --tb=short
        env:
          # Les tests ne doivent JAMAIS appeler la vraie API OpenAI :
          # le code lit OPENAI_API_KEY mais les tests mockent le client.
          OPENAI_API_KEY: dummy-pour-tests

  # Variante moderne avec uv (à vérifier : astral-sh/setup-uv@v5, vu en 2026)
  # test-uv:
  #   runs-on: ubuntu-latest
  #   steps:
  #     - uses: actions/checkout@v4
  #     - uses: actions/setup-python@v5
  #       with: { python-version: "3.12" }
  #     - uses: astral-sh/setup-uv@v5
  #     - run: uv sync
  #     - run: uv run pytest -q
```

`pyproject.toml` minimal (config ruff) :

```toml
[project]
name = "rag-perso"
version = "0.1.0"
requires-python = ">=3.12"
dependencies = [
  "openai>=1.0",
  "psycopg[binary]>=3.1",
  "tiktoken>=0.7",
  "beautifulsoup4>=4.12",
  "python-dotenv>=1.0",
]

[tool.ruff]
line-length = 100
target-version = "py312"

[tool.ruff.lint]
select = ["E", "F", "I", "W"]
ignore = ["E501"]   # la longueur est gérée par le formateur
```

Et un **vrai** test que la CI exécutera (`tests/test_chunk_md.py`) :

```python
from scripts.chunk_md import chunk_markdown_sections

def test_chunk_md_ne_coupe_pas_les_sections():
    md = "# Titre\n\n## A\n\nTexte A.\n\n## B\n\nTexte B."
    chunks = chunk_markdown_sections(md, source="test.md")
    assert len(chunks) == 2
    assert chunks[0]["metadata"]["section"] == "A"
    assert chunks[1]["metadata"]["section"] == "B"
```

> **Testabilité :** pour que ce test passe en CI, `chunk_markdown_sections`
> doit être **importable sans effet de bord** : pas d'appel réseau ni de
> lecture de clé API au niveau du module. Mets le code CLI sous
> `if __name__ == "__main__":`.

## 21. Branch protection : interdire le push direct sur main

Repo → **Settings → Branches → Add branch protection rule** :
- Branch name pattern : `main`
- ☑ *Require a pull request before merging*
- ☑ *Require status checks to pass* → sélectionne `Lint (ruff)` et
  `Tests (pytest)` (noms des jobs du §20)
- ☑ *Do not allow bypassing the above settings*

Workflow quotidien :

