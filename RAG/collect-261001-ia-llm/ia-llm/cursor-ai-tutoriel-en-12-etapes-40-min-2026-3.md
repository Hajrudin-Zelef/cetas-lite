---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-12-etapes-40-min-2026-3
title: "Windows (winget)"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "embeddings", "mcp", "model context protocol"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [154, 262]
sha256: 9f7da31411ce55cad7910d988e5c4f2d6e8f5c526f50c46e2749990b0f7882ba
---

# Windows (winget)

```
# app/models.py
from datetime import datetime
from sqlmodel import SQLModel, Field
class Tache(SQLModel, table=True):
    """Modèle représentant une tâche à accomplir."""
    id: int | None = Field(default=None, primary_key=True)
    titre: str = Field(index=True)
    terminee: bool = Field(default=False)
    date_creation: datetime = Field(default_factory=datetime.utcnow)
class TacheCreation(SQLModel):
    """Schéma d'entrée pour créer une tâche."""
    titre: str
class TacheMaj(SQLModel):
    """Schéma d'entrée pour mettre à jour une tâche."""
    titre: str | None = None
    terminee: bool | None = None
```
**Piège n°2 :** ne validez jamais un gros diff sans le lire. Le mode Agent peut créer ou écraser plusieurs fichiers d’un coup. Parcourez chaque changement, surtout sur un projet existant, et utilisez le contrôle de version Git comme filet de sécurité.

## Étape 6 – Accélérer avec Tab et l’édition en ligne (Cmd+K)

Une fois le squelette en place, le travail au fil de l’eau s’appuie surtout sur **Tab** et **Ctrl/Cmd + K**. Ouvrez `app/main.py` et commencez à écrire une nouvelle route : après quelques caractères, Cursor propose souvent l’implémentation complète, que vous acceptez avec *Tab*. La fonction d’autocomplétion prédit aussi les éditions suivantes (import manquant, retour de type, gestion d’erreur).

Pour une modification ciblée, sélectionnez un bloc, appuyez sur *Ctrl/Cmd + K* et décrivez le changement. Exemples de consignes efficaces :

```
Ajoute une pagination (paramètres limit et offset) à GET /taches.
Transforme cette fonction en async et gère l'erreur 404
si la tâche n'existe pas.
Ajoute une validation : le titre ne doit pas dépasser 200 caractères.
```
L’édition en ligne est idéale pour les retouches chirurgicales, là où l’Agent est plus adapté aux tâches transversales touchant plusieurs fichiers. Alterner intelligemment entre les deux est le vrai secret de la vitesse avec **Cursor AI**. Pour une approche 100 % terminal, notre tutoriel Claude Code montre une philosophie complémentaire.

## Étape 7 – Interroger son code : @Codebase, @Web et @Docs

Dans le Chat (*Ctrl/Cmd + L*), les symboles **@** permettent d’injecter du contexte précis dans vos requêtes. C’est ce qui distingue Cursor d’un chatbot générique : il travaille sur *votre* code, pas dans le vide.

- **@Codebase** – Cursor recherche dans l’intégralité du projet indexé pour répondre (« Où est gérée la connexion à la base ? »).
- **@Files** /**@Folders** – cibler un fichier ou un dossier précis comme contexte.
- **@Web** – Cursor effectue une recherche web pour intégrer des informations à jour.
- **@Docs** – pointer vers une documentation officielle (FastAPI, SQLModel, etc.) indexée par Cursor.
- **@Git** – analyser des commits ou des différences pour comprendre l’historique.

```
@Codebase Explique comment la session SQLite est ouverte et fermée.
@Docs FastAPI Montre-moi comment ajouter une gestion CORS.
@Web Quelle est la dernière version stable de SQLModel et ses changements ?
```
Cette capacité à interroger le code repose sur l’**indexation** de votre dépôt. À l’ouverture d’un projet, Cursor construit un index sémantique (embeddings) en arrière-plan. Sur un très gros dépôt, patientez quelques minutes le temps que l’indexation se termine, sinon les réponses @Codebase seront partielles.

## Étape 8 – Laisser l’Agent exécuter le terminal et corriger les erreurs

C’est ici que **Cursor** dépasse le simple assistant. En mode Agent, l’IA peut proposer et exécuter des commandes dans le terminal intégré, lire la sortie, détecter les erreurs et se corriger toute seule. Demandez-lui d’installer les dépendances et de lancer le serveur :

```
Crée un environnement virtuel .venv, installe les dépendances
de requirements.txt, puis démarre le serveur avec uvicorn
et vérifie que /docs répond.
```
L’agent exécute alors les commandes (après votre autorisation) et affiche la sortie. En cas d’échec – par exemple un module manquant – il lit le message d’erreur et propose la correction. Exemple de sortie type dans le terminal :

```
$ python -m venv .venv && source .venv/bin/activate
$ pip install -r requirements.txt
Successfully installed fastapi-0.115.x sqlmodel-0.0.x uvicorn-0.3x.x
$ uvicorn app.main:app --reload
INFO:     Uvicorn running on http://127.0.0.1:8000 (Press CTRL+C to quit)
INFO:     Application startup complete.
# Documentation interactive disponible sur http://127.0.0.1:8000/docs
```
**Piège n°3 :** activez le mode d’exécution automatique (« auto-run ») en connaissance de cause. Il est très pratique, mais il autorise l’agent à lancer des commandes sans confirmation. Sur une machine de production ou un dépôt sensible, gardez la validation manuelle activée et surveillez chaque commande, en particulier les suppressions de fichiers.

## Étape 9 – Générer et lancer les tests avec pytest

Un projet sérieux se teste. Demandez à l’Agent de créer une suite `pytest` couvrant les principaux endpoints. Grâce à la règle de projet définie à l’étape 4, Cursor sait déjà qu’il doit générer des tests pour chaque route.

```
# tests/test_taches.py
from fastapi.testclient import TestClient
from app.main import app
client = TestClient(app)
def test_creation_tache():
    """Vérifie la création d'une tâche."""
    reponse = client.post("/taches", json={"titre": "Écrire l'article"})
    assert reponse.status_code == 201
    donnees = reponse.json()
    assert donnees["titre"] == "Écrire l'article"
    assert donnees["terminee"] is False
def test_liste_taches():
    """Vérifie que la liste des tâches est renvoyée."""
    reponse = client.get("/taches")
    assert reponse.status_code == 200
    assert isinstance(reponse.json(), list)
```
Lancez les tests avec `pytest -q`. Si un test échoue, ne corrigez pas à la main : collez la sortie dans le Chat ou demandez à l’Agent « les tests échouent, corrige jusqu’à ce que tout passe ». Cursor lit la trace, identifie la cause (souvent un code de statut mal configuré) et itère. C’est la boucle « écrire → tester → corriger » automatisée, l’un des grands atouts du mode agentique.

## Étape 10 – Réviser le code avec Bugbot et les agents en arrière-plan

Au-delà de l’écriture, Cursor aide à **relire** le code. **Bugbot** est la fonction de revue automatisée d’Anysphere : branchée sur vos *pull requests*, elle repère les bugs potentiels, les problèmes de sécurité et les régressions, puis propose des correctifs. C’est un garde-fou précieux, notamment quand une partie du code a été générée par IA.

Les **agents en arrière-plan** (et les agents cloud) constituent l’autre grande nouveauté de l’ère Cursor 2.0 : vous lancez une tâche longue (refactorisation, migration, écriture de tests sur tout un module) et l’agent travaille pendant que vous continuez sur autre chose. L’interface multi-agents permet de suivre plusieurs de ces chantiers en parallèle et de comparer les résultats avant d’accepter.

Pour notre projet, une consigne typique en arrière-plan serait : « Ajoute la journalisation (logging) structurée sur tous les endpoints et écris les tests correspondants. » Pendant l’exécution, vous restez libre de coder ailleurs. À la fin, vous relisez le diff proposé, exactement comme une revue de code humaine.

## Étape 11 – Étendre Cursor avec MCP (Model Context Protocol)

Le **Model Context Protocol (MCP)** est un standard ouvert qui permet de connecter Cursor à des outils et sources de données externes : bases de données, gestionnaires de tickets, documentation interne, services cloud. Concrètement, un serveur MCP expose des « outils » que l’agent peut appeler, comme il appellerait le terminal.

