---
id: collect-261001-general-networking/general-networking/uv-python-100x-plus-vite-que-pip-12-etapes-2026-3
title: "macOS et Linux"
domain: general-networking
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/uv-python-100x-plus-vite-que-pip-12-etapes-2026.md
source_anchor: ""
source_lines: [204, 346]
sha256: eda51e2b733691d2d7cd94511b93b17dd45bfd91b96f7ef11a2f1e1975663b26
---

# macOS et Linux

```
import httpx
import typer
app = typer.Typer(help="Affiche la derniere version d'un paquet PyPI.")
@app.command()
def version(paquet: str) -> None:
    """Recupere la derniere version publiee d'un paquet."""
    url = f"https://pypi.org/pypi/{paquet}/json"
    reponse = httpx.get(url, timeout=10.0)
    if reponse.status_code == 404:
        typer.secho(f"Paquet introuvable : {paquet}", fg=typer.colors.RED)
        raise typer.Exit(code=1)
    reponse.raise_for_status()
    info = reponse.json()["info"]
    typer.secho(f"{info['name']} {info['version']}", fg=typer.colors.GREEN)
    if info.get("summary"):
        typer.echo(info["summary"])
def main() -> None:
    app()
if __name__ == "__main__":
    main()
```
Exécutez l’application avec `uv run`. Cette commande synchronise l’environnement si nécessaire, puis lance votre code dans le bon interpréteur – sans que vous ayez à activer quoi que ce soit :

`uv run pypeek version requests`
Sortie attendue :

```
requests 2.32.3
Python HTTP for Humans.
```
Toujours utiliser `uv run` plutôt que `python` directement : c’est la garantie que le script s’exécute avec les bonnes dépendances et la bonne version de Python, celle épinglée à l’étape 2.

## Étape 7 – Le lockfile universel et les builds reproductibles

Le fichier `uv.lock`, généré et géré exclusivement par uv, est un **lockfile universel** au format TOML lisible par un humain : il capture les versions exactes et les empreintes (hashes) de chaque dépendance, pour toutes les plateformes à la fois. C’est ce qui rend un build reproductible d’une machine à l’autre. Depuis novembre 2025, la documentation d’Astral détaille en outre l’export de ce lockfile vers trois formats – `requirements.txt`, `pylock.toml` et une nomenclature CycloneDX – dont un SBOM CycloneDX 1.5 généré via `uv export --format cyclonedx1.5`, un atout pour la conformité logicielle en entreprise. Vous pouvez régénérer le lockfile explicitement :

```
# (Re)generer le fichier de verrouillage
uv lock
# Synchroniser l'environnement a partir du lock, sans le modifier
uv sync --frozen
# Verifier que le lock est a jour (echoue sinon) – ideal en CI
uv sync --locked
```
Règle d’or : **versionnez `uv.lock` dans Git, mais jamais le dossier `.venv`**. Ajoutez ce dernier à votre `.gitignore`. Le lockfile est la source de vérité ; l’environnement virtuel n’est qu’un artefact reconstructible en une seconde.

### Pourquoi le lockfile universel change la donne

Contrairement à un `requirements.txt` figé pour une seule plateforme, `uv.lock` encode les résolutions conditionnelles (par système d’exploitation, architecture ou version de Python) grâce à ce qu’Astral appelle la **résolution universelle**, un mécanisme que la documentation de résolution d’Astral, mise à jour en septembre 2026, continue de présenter comme le socle du lockfile. Un même fichier fonctionne donc à l’identique sur le portable macOS d’un développeur, sur un runner Linux en CI et dans une image Docker. En cas de conflit Git sur `uv.lock`, ne le résolvez jamais à la main : relancez simplement `uv lock`.

## Étape 8 – Exécuter des outils sans les installer (uvx)

uv remplace pipx pour exécuter des outils Python en ligne de commande. `uvx` (alias de `uv tool run`) lance un outil dans un environnement éphémère, sans polluer votre projet. Par exemple, pour vérifier du code avec Ruff sans même l’ajouter au projet :

```
# Executer un outil de maniere ephemere
uvx ruff check .
# Executer une version precise
uvx [email protected] --version
# Installer un outil de maniere permanente (globale)
uv tool install ruff
```
La distinction est importante : `uvx` convient aux usages ponctuels (une vérification rapide, un générateur de projet), tandis que `uv tool install` place l’exécutable dans votre PATH pour un usage quotidien, à la manière de pipx. Les deux profitent du cache global de uv, ce qui rend les lancements suivants quasi instantanés.

## Étape 9 – Scripts autonomes avec dépendances inline (PEP 723)

uv implémente le standard PEP 723, qui permet d’embarquer les dépendances directement dans l’en-tête d’un script Python isolé. Plus besoin de projet ni de `requirements.txt` pour un utilitaire d’une page. Créez un fichier `rapide.py` :

```
# /// script
# requires-python = ">=3.13"
# dependencies = ["httpx"]
# ///
import httpx
reponse = httpx.get("https://pypi.org/pypi/uv/json", timeout=10.0)
print("Derniere version de uv :", reponse.json()["info"]["version"])
```
Exécutez-le directement : uv lit l’en-tête, crée un environnement temporaire avec httpx, et lance le script – le tout automatiquement :

```
uv run rapide.py
# Derniere version de uv : 0.11.26
```
Vous pouvez aussi ajouter une dépendance à un script via `uv add --script rapide.py rich`, ou fournir des dépendances ad hoc sans modifier le fichier avec `uv run --with rich rapide.py`. Cette fonctionnalité est idéale pour partager des scripts d’automatisation qui « fonctionnent partout » sans instructions d’installation.

## Étape 10 – Tester, linter et formater

Un projet complet a besoin de tests. Créez un dossier `tests/` avec un fichier `test_cli.py` qui vérifie le comportement de notre CLI grâce au `CliRunner` de Typer :

```
from typer.testing import CliRunner
from pypeek import app
runner = CliRunner()
def test_version_paquet_connu():
    resultat = runner.invoke(app, ["version", "pip"])
    assert resultat.exit_code == 0
    assert "pip" in resultat.stdout
def test_paquet_inexistant():
    resultat = runner.invoke(app, ["version", "paquet-qui-nexiste-pas-xyz"])
    assert resultat.exit_code == 1
```
Lancez la suite de tests avec `uv run`, qui utilise le pytest du groupe de développement :

```
uv run pytest -q
# ..                                    [100%]
# 2 passed in 0.42s
```
Enfin, contrôlez la qualité et le formatage avec Ruff – l’autre outil d’Astral, écrit lui aussi en Rust :

```
# Detecter les problemes de style et de logique
uv run ruff check .
# Formater automatiquement le code
uv run ruff format .
```
Puisque uv et Ruff partagent le même éditeur, leur intégration est sans friction. Cette pile Astral (uv + Ruff) est aujourd’hui l’un des socles les plus rapides pour l’outillage Python moderne. Consultez la documentation de Ruff pour affiner vos règles de style.

## Étape 11 – Construire et publier le paquet

uv remplace `build` et `twine`. Pour générer les artefacts de distribution (une archive source « sdist » et un « wheel »), une seule commande suffit :

```
uv build
# Successfully built dist/pypeek-0.1.0.tar.gz
# Successfully built dist/pypeek-0.1.0-py3-none-any.whl
```
Pour la publication, il est fortement recommandé de tester d’abord sur TestPyPI avant de viser le vrai PyPI. uv lit le jeton d’authentification depuis la variable d’environnement `UV_PUBLISH_TOKEN` :

```
# Publier sur TestPyPI pour valider
export UV_PUBLISH_TOKEN="pypi-VOTRE_JETON"
uv publish --publish-url https://test.pypi.org/legacy/
# Puis, une fois valide, sur le vrai PyPI
uv publish
```
En cas d’erreur 403 lors de la publication, vérifiez la validité de votre jeton et que le nom du paquet n’est pas déjà pris sur l’index visé. Le flux `uv build` puis `uv publish` remplace intégralement l’ancien duo `python -m build` + `twine upload`.

## Étape 12 – Conteneuriser avec Docker et automatiser en CI/CD

Pour déployer, rien de tel qu’une image Docker minimale. Astral publie des images officielles contenant le binaire uv. Voici un `Dockerfile` optimisé, qui met en cache les dépendances séparément du code source pour accélérer les reconstructions :

