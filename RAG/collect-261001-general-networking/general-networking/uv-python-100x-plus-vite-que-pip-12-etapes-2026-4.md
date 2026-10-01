---
id: collect-261001-general-networking/general-networking/uv-python-100x-plus-vite-que-pip-12-etapes-2026-4
title: "macOS et Linux"
domain: general-networking
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["agents", "packaging"]
source: docs/RAG/collect-261001-general-networking/uv-python-100x-plus-vite-que-pip-12-etapes-2026.md
source_anchor: ""
source_lines: [347, 479]
sha256: 5ef56a42e3d0f84c4580771fef38a8bb3e9ae8dab268e714845983c3599684fe
---

# macOS et Linux

```
FROM ghcr.io/astral-sh/uv:python3.13-bookworm-slim
# Optimisations recommandees par Astral
ENV UV_COMPILE_BYTECODE=1
ENV UV_LINK_MODE=copy
WORKDIR /app
# 1) Installer d'abord les dependances (couche mise en cache)
COPY pyproject.toml uv.lock ./
RUN uv sync --frozen --no-dev --no-install-project
# 2) Copier le code puis installer le projet
COPY . .
RUN uv sync --frozen --no-dev
ENTRYPOINT ["uv", "run", "pypeek"]
```
Construisez et lancez l’image :

```
docker build -t pypeek .
docker run --rm pypeek version fastapi
```
Côté intégration continue, l’action officielle `astral-sh/setup-uv` installe uv et active le cache en une ligne. Voici un workflow GitHub Actions qui installe les dépendances, lance Ruff et exécute les tests :

```
name: CI
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Installer uv
        uses: astral-sh/setup-uv@v6
        with:
          enable-cache: true
      - name: Synchroniser (build reproductible)
        run: uv sync --locked
      - name: Linter
        run: uv run ruff check .
      - name: Tests
        run: uv run pytest -q
```
Vérifiez la dernière version majeure de l’action `setup-uv` dans sa documentation avant de l’épingler. Combinée au cache, cette configuration ramène le temps d’installation des dépendances en CI à quelques centaines de millisecondes. Si vous conteneurisez d’autres services, notre tutoriel Docker complète parfaitement cette étape.

## Migrer depuis pip et Poetry vers uv

Vous avez un projet existant ? La transition est indolore. Pour importer un ancien `requirements.txt` dans un projet uv, utilisez :

```
# Importer les dependances d'un requirements.txt
uv add -r requirements.txt
# Ou, en mode pip pur, generer un requirements.txt verrouille
uv pip compile pyproject.toml -o requirements.txt
# Installer depuis un requirements.txt (interface pip)
uv pip install -r requirements.txt
```
Depuis Poetry, la structure `pyproject.toml` est largement compatible : dans la plupart des cas, il suffit de déplacer les dépendances vers la section standard `[project]`, de supprimer les sections spécifiques à Poetry, puis de lancer `uv lock`. Beaucoup d’équipes conservent d’ailleurs Poetry et uv en parallèle pendant la période de transition, sans conflit, car uv respecte les standards de packaging Python (PEP 621). Selon le guide 2026 de Pynions, cette bascule est désormais si aboutie qu’un unique fichier `uv.lock` remplace purement et simplement les anciens flux `requirements.txt` et pip-tools pour la gestion des dépendances d’un projet.

## uv contre Poetry, PDM et Rye : lequel choisir en 2026 ?

uv n’est pas seul sur le créneau des gestionnaires de projets Python. Poetry reste très répandu, PDM a des adeptes, et Rye – dont la maintenance a été confiée à Astral – fait figure de précurseur que uv est venu remplacer. Une comparaison indépendante de Techplained, menée en avril 2026 sur uv 0.5+ face à pip 25+, Poetry 2.x et PDM 2.x, a confirmé des accélérations de **5 à 8 fois** par rapport à Poetry. Le tableau suivant résume les différences structurelles qui font aujourd’hui pencher la balance vers uv.

| Critère | uv | Poetry | PDM | pip + venv | 
|---|---|---|---|---|
| Vitesse d’installation | Très élevée (Rust) | Moyenne | Moyenne | Faible | 
| Versions de Python intégrées | Oui | Non | Non | Non | 
| Lockfile universel | Oui | Oui (poetry.lock) | Oui | Non | 
| Remplace pipx / pyenv / twine | Oui | Non | Partiel | Non | 
| Standard PEP 621 | Oui | Oui (récent) | Oui | – | 
| Langage d’implémentation | Rust | Python | Python | Python | 

La différence décisive tient en deux points. D’abord, uv est le seul à intégrer nativement la gestion des versions de Python : nul besoin d’installer pyenv en parallèle. Ensuite, sa vitesse – héritée de Rust et de son résolveur – le place dans une catégorie à part pour les gros projets et la CI. Poetry demeure un excellent choix si votre équipe l’a déjà standardisé et que la vitesse n’est pas critique, mais pour un nouveau projet en 2026, **uv Python** est le point de départ recommandé. Sa reprise par OpenAI ne fait que renforcer la solidité de ce pari à long terme.

## Exemple concret : gérer un service FastAPI avec uv

Notre CLI illustre bien le cycle de vie d’un paquet, mais uv brille tout autant pour les applications web. Voici comment démarrer un microservice FastAPI en quelques secondes, dépendances verrouillées comprises :

```
uv init --package web-api
cd web-api
uv add fastapi "uvicorn[standard]"
```
Dans `src/web_api/__init__.py`, un point d’entrée minimal :

```
from fastapi import FastAPI
app = FastAPI(title="web-api")
@app.get("/sante")
def sante() -> dict[str, str]:
    return {"statut": "ok"}
```
Lancez le serveur de développement, toujours via `uv run`, qui garantit le bon environnement :

```
uv run uvicorn web_api:app --reload
# Uvicorn running on http://127.0.0.1:8000
```
La force de cette approche est la reproductibilité : le `uv.lock` généré capture non seulement FastAPI et Uvicorn, mais aussi toute leur chaîne de dépendances transitives, avec leurs empreintes cryptographiques. Vos collègues, vos serveurs et votre pipeline CI obtiendront tous exactement les mêmes versions. Le même mécanisme s’applique aux projets d’intelligence artificielle : un service qui appelle un modèle local via Ollama ou orchestre des agents avec LangGraph se gère de la même manière, avec les mêmes garanties.

## Structure finale du projet et bonnes pratiques d’équipe

Au terme des douze étapes, votre projet `pypeek` présente une structure propre, standard et reproductible :

```
pypeek/
├── .github/workflows/ci.yml
├── .gitignore
├── .python-version
├── Dockerfile
├── README.md
├── pyproject.toml
├── uv.lock
├── src/
│   └── pypeek/
│       └── __init__.py
└── tests/
    └── test_cli.py
```
Un point souvent négligé : le fichier `.gitignore`. Il doit exclure l’environnement virtuel et le cache, mais surtout **pas** le lockfile ni le fichier de version Python :

```
# .gitignore
.venv/
__pycache__/
dist/
*.egg-info/
# NE PAS ignorer : uv.lock et .python-version (a versionner !)
```
Pour une équipe, quelques règles simples garantissent une expérience fluide. Versionnez toujours `uv.lock` et `.python-version` : ils font partie du contrat du projet. En CI, utilisez systématiquement `uv sync --locked` afin d’échouer explicitement si le lock n’est pas à jour, plutôt que d’installer silencieusement des versions divergentes. Traitez les modifications de `uv.lock` comme du code : elles doivent apparaître dans les revues de pull request, car un changement de dépendance transitive peut avoir des conséquences. Enfin, encouragez chacun à utiliser `uv run` pour toute commande liée au projet, ce qui élimine les différences d’environnement entre les postes. Avec ces habitudes, la promesse de reproductibilité de uv devient une réalité opérationnelle plutôt qu’un simple argument marketing.

## Aide-mémoire des commandes uv essentielles

Gardez ce tableau sous la main : il couvre 90 % de votre usage quotidien de uv Python.

