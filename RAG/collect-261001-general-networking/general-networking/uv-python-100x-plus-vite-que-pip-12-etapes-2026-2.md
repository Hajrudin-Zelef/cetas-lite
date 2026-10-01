---
id: collect-261001-general-networking/general-networking/uv-python-100x-plus-vite-que-pip-12-etapes-2026-2
title: "macOS et Linux"
domain: general-networking
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/uv-python-100x-plus-vite-que-pip-12-etapes-2026.md
source_anchor: ""
source_lines: [55, 203]
sha256: 337751f42da30aac1e9a7ca2f37833f7d06c6e499516f34457892765933451ff
---

# macOS et Linux

Ce tutoriel a été rédigé et testé avec la dernière version stable de uv, **0.12.7**, publiée par Astral le 27 août 2026 (comme le relève PythonDataBench). Voici ce dont vous avez besoin avant de commencer. Bonne nouvelle : la liste est courte, car uv n’exige quasiment rien en amont. Le rythme de publication d’Astral reste soutenu – entre la 0.12.2 du 5 août 2026, recensée par Modern Python Weekly, et la 0.12.6 du 25 août 2026, ce ne sont pas moins de trois versions qui sont sorties en l’espace de trois semaines – ce qui garantit des correctifs et améliorations quasi hebdomadaires.

| Élément | Version / exigence | Remarque | 
|---|---|---|
| uv | 0.11.26 (dernière stable) | Binaire Rust autonome | 
| Système d’exploitation | macOS, Linux ou Windows 10+ | Multiplateforme | 
| Python | Aucun préinstallé requis | uv le télécharge à la demande | 
| Terminal | bash, zsh ou PowerShell | Toutes les commandes sont en CLI | 
| Docker (facultatif) | 24+ | Uniquement pour l’étape 12 | 
| Compte PyPI / TestPyPI (facultatif) | – | Uniquement pour la publication | 
| Espace disque | Quelques centaines de Mo | Pour le cache global de uv | 

Si vous suivez d’autres tutoriels de développement sur Tech Insider, sachez que uv s’intègre parfaitement à des piles comme FastAPI ou aux projets d’IA locale de type Ollama. Nous y reviendrons.

## Étape 1 – Installer uv (macOS, Linux, Windows)

La méthode recommandée est l’installateur autonome, qui télécharge le binaire et l’ajoute au PATH – une approche documentée dès novembre 2025 par le tutoriel du CNRS SARI, qui présente uv comme le gestionnaire de projets Python moderne d’Astral et détaille cet installateur basé sur curl. Astral continue d’affiner ces scripts : la version 0.11.28, publiée le 7 juillet 2026, a notamment republié des scripts d’installation précompilés le même mois. Sur macOS et Linux, ouvrez un terminal et exécutez :

```
# macOS et Linux
curl -LsSf https://astral.sh/uv/install.sh | sh
```
Sous Windows, utilisez PowerShell :

```
# Windows (PowerShell)
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```
Vous préférez passer par un gestionnaire de paquets existant ? uv est aussi disponible via pip, pipx et Homebrew. Les installateurs shell et PowerShell officiels, eux, existent depuis longtemps – ils ont été formalisés dès la version 0.9.27 du 26 janvier 2026 – et la 0.11.8, publiée le 27 avril 2026, a ajouté de nouveaux binaires d’installation pour davantage de plateformes :

```
# Alternatives
pip install uv        # via pip
pipx install uv       # via pipx
brew install uv       # via Homebrew (macOS/Linux)
```
Vérifiez ensuite que tout fonctionne et mettez uv à jour si besoin :

```
uv --version
# uv 0.11.26
# Mettre a jour uv (si installe via l'installateur autonome)
uv self update
```
Si la commande `uv` reste introuvable après l’installation, redémarrez votre terminal ou rechargez votre shell (`source ~/.bashrc` ou `source ~/.zshrc`). L’installateur ajoute uv à `~/.local/bin`, qui doit figurer dans votre PATH.

## Étape 2 – Gérer les versions de Python avec uv

C’est ici que uv Python remplace pyenv. uv sait télécharger, installer et épingler n’importe quelle version de l’interpréteur, sans compilation ni configuration système. Listez d’abord les versions disponibles et installées :

```
# Voir les versions disponibles et installees
uv python list
# Installer une version precise
uv python install 3.13
# Installer plusieurs versions d'un coup
uv python install 3.11 3.12 3.13
```
uv gère les téléchargements de manière transparente : si un projet réclame une version absente, il l’installe automatiquement au moment de l’exécution. Vous pouvez épingler une version pour un projet donné, ce qui crée un fichier `.python-version` lu par tous les collègues :

```
# Epingler la version pour le repertoire courant
uv python pin 3.13
# Cree un fichier .python-version contenant "3.13"
```
Cette approche élimine la classe entière de bugs « ça marche sur ma machine » liés à des interpréteurs divergents. Le fichier `.python-version` se versionne dans Git, garantissant que toute l’équipe – et la CI – utilise exactement le même Python.

## Étape 3 – Initialiser un projet uv Python

Nous allons créer un paquet distribuable, avec un point d’entrée en ligne de commande. L’option `--package` génère une structure « src layout » complète, prête pour la publication :

```
uv init --package pypeek
cd pypeek
```
uv crée l’arborescence suivante :

```
pypeek/
├── .python-version
├── README.md
├── pyproject.toml
└── src/
    └── pypeek/
        └── __init__.py
```
Le fichier `pyproject.toml` est le cœur du projet. Il contient les métadonnées, les dépendances et la déclaration du point d’entrée CLI. Voici à quoi il ressemble après l’initialisation :

```
[project]
name = "pypeek"
version = "0.1.0"
description = "Affiche la derniere version d'un paquet PyPI"
readme = "README.md"
requires-python = ">=3.13"
dependencies = []
[project.scripts]
pypeek = "pypeek:main"
[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"
```
Notez la section `[project.scripts]` : elle déclare que la commande `pypeek` appellera la fonction `main()` du module. Nous écrirons cette fonction à l’étape 6.

## Étape 4 – Ajouter des dépendances avec uv add

Notre CLI a besoin de deux paquets : `httpx` pour les requêtes HTTP et `typer` pour construire l’interface en ligne de commande. La commande `uv add` les ajoute au `pyproject.toml`, met à jour le fichier de verrouillage `uv.lock` et synchronise l’environnement virtuel – le tout en une fraction de seconde :

`uv add httpx typer`
Sortie typique :

```
Resolved 12 packages in 84ms
Installed 12 packages in 21ms
 + httpx==0.28.1
 + typer==0.16.0
 + click==8.1.8
 + ...
```
La section `dependencies` de votre `pyproject.toml` reflète désormais ces ajouts, avec des contraintes de version minimales :

```
dependencies = [
    "httpx>=0.28.1",
    "typer>=0.16.0",
]
```
Point crucial à retenir : `uv add` n’est pas la même chose que `uv pip install`. Le premier modifie `pyproject.toml` et `uv.lock` – c’est le mode « projet », reproductible. Le second se comporte comme un pip classique et n’écrit rien dans les métadonnées. Pour un projet géré, utilisez toujours `uv add`. Vous n’avez d’ailleurs jamais besoin d’activer manuellement l’environnement : uv crée et gère le dossier `.venv` pour vous.

## Étape 5 – Dépendances de développement et groupes

Les outils de test et de qualité de code ne doivent pas être livrés à vos utilisateurs. uv les range dans un groupe de développement séparé grâce à l’option `--dev` :

`uv add --dev pytest ruff`
Ces dépendances atterrissent dans une section dédiée du `pyproject.toml`, distincte des dépendances de production :

```
[dependency-groups]
dev = [
    "pytest>=8.3.0",
    "ruff>=0.14.0",
]
```
Lors d’un déploiement en production, vous pourrez exclure ce groupe avec `uv sync --no-dev`, ce qui allège considérablement l’image finale. À l’inverse, en local et en CI, `uv sync` installe tout par défaut. Cette séparation nette entre dépendances applicatives et outils de développement est l’une des bonnes pratiques que uv rend triviales, là où un simple `requirements.txt` mélangeait souvent les deux.

## Étape 6 – Écrire l’application et l’exécuter avec uv run

Place au code. Ouvrez `src/pypeek/__init__.py` et remplacez son contenu par notre CLI, qui interroge l’API JSON publique de PyPI :

