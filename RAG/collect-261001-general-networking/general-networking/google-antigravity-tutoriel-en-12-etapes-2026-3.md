---
id: collect-261001-general-networking/general-networking/google-antigravity-tutoriel-en-12-etapes-2026-3
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: []
keywords: ["agent", "agents", "claude", "gemini", "mai"]
source: docs/RAG/collect-261001-general-networking/google-antigravity-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [110, 228]
sha256: bde506d9ac2aae57ca3e5d31c12aabd2437ac967bb25e5708916c0a08cfed7d4
---

# macOS (Homebrew)

Les modèles Claude peuvent transiter par votre propre compte Anthropic. Dans *Settings > Model providers > Anthropic*, collez votre clé API. Cela déporte la facturation Claude sur votre compte Anthropic et libère du quota Google. C’est utile si vous atteignez régulièrement les limites de l’aperçu gratuit. Définissez ensuite Claude Sonnet 4.6 comme modèle par défaut d’un agent donné pour tester ce routage.

```
# Exemple : définir la clé Anthropic via une variable d'environnement
# (à placer avant de lancer Antigravity depuis un terminal)
export ANTHROPIC_API_KEY="sk-ant-votre-cle-ici"
# Vérifier qu'elle est bien exportée
echo $ANTHROPIC_API_KEY | cut -c1-7
# Sortie attendue : sk-ant-
```
## Étape 6 : cadrer l’agent avec un fichier AGENTS.md

Un agent autonome sans règles produit du code incohérent. La solution est le fichier `AGENTS.md`, une convention désormais partagée par de nombreux outils agentiques : un document Markdown, placé à la racine du projet, qui décrit les conventions, la pile technique, les commandes de test et les interdits. Antigravity lit ce fichier au début de chaque mission et s’y conforme. C’est l’équivalent d’un cahier des charges permanent pour vos agents.

Créez le fichier `AGENTS.md` à la racine du dossier `gestionnaire-taches`. Voici un modèle complet pour notre projet :

```
# AGENTS.md – Gestionnaire de tâches
## Objectif
Construire une application web de gestion de tâches (to-do) avec une API REST
et une interface HTML/JS minimale, sans framework front lourd.
## Pile technique
- Backend : Python 3.11+, FastAPI, base SQLite (fichier taches.db)
- Front : HTML + JavaScript vanilla (fetch), aucun build
- Tests : pytest
## Conventions
- Code et commentaires en français.
- Noms de variables explicites, pas d'abréviations obscures.
- Chaque endpoint documenté avec un docstring.
- Aucune dépendance non listée dans requirements.txt.
## Commandes
- Installer : pip install -r requirements.txt
- Lancer : uvicorn main:app --reload
- Tester : pytest -q
## Interdits
- Ne pas ajouter de base de données externe (Postgres, MySQL).
- Ne pas exposer de secret en clair dans le code.
- Ne pas modifier ce fichier AGENTS.md sans validation.
```
Ce fichier fait deux choses essentielles. D’abord, il verrouille la pile technique : sans lui, un agent pourrait choisir Django, Flask ou Node selon son humeur. Ensuite, il impose des garde-fous (« Interdits ») qui réduisent drastiquement les dérapages. Complétez-le avec la **base de connaissances** intégrée : quand un agent découvre une information utile (une commande, un choix de conception), il peut la sauvegarder pour les tâches suivantes, ce qui renforce la cohérence au fil du projet.

Astuce : gardez `AGENTS.md` court et impératif. Un fichier de 30 lignes bien ciblées est plus efficace qu’un pavé de 300 lignes que l’agent survole. Mettez-le à jour à mesure que le projet grandit, mais restez concis.

## Étape 7 : Éditeur ou Manager, deux façons de travailler

Google Antigravity propose deux modes de travail qu’il faut bien distinguer pour être efficace. Le choix dépend de la nature de la tâche : édition fine et synchrone, ou délégation autonome et asynchrone.

### La vue Éditeur : le flux synchrone

La vue Éditeur ressemble à VS Code : arborescence de fichiers, éditeur central, terminal intégré, et une barre latérale d’agent. Vous y bénéficiez de la complétion par tabulation et de commandes en ligne (sélectionnez du code, décrivez la modification, l’agent l’applique). C’est le mode idéal pour les retouches précises, la lecture de code, le débogage pas à pas – bref, tout ce qui demande votre présence active. Comparé à Cursor, l’expérience est très familière.

### Le Manager : l’orchestration asynchrone

Le Manager (ou « Mission Control ») est le cœur de la promesse d’Antigravity. Vous y lancez une **mission** : une description de haut niveau que l’agent transforme en plan, puis exécute pendant que vous faites autre chose. Vous pouvez démarrer plusieurs missions en parallèle – par exemple un agent qui écrit le backend et un autre qui prépare les tests – chacune sur son propre espace de travail ; dès le lancement d’Antigravity 2.0 le 19 mai 2026, Google a présenté sa plateforme comme un environnement agent-first unifié capable de faire tourner plusieurs agents IA en parallèle – jusqu’à cinq de front, propulsés par Gemini 3.5 Flash, selon Emergent. Le Manager affiche l’avancement, les Artifacts produits et les points où votre validation est requise.

Pour ce tutoriel, nous utiliserons le Manager pour la génération initiale du projet (une tâche large et autonome), puis la vue Éditeur pour les corrections fines. C’est la combinaison la plus productive : déléguer le gros œuvre, garder la main sur les finitions.

## Étape 8 : lancer votre première mission (le projet complet)

C’est le moment clé. Ouvrez le Manager, cliquez sur « New mission » et rédigez une consigne claire. La qualité du prompt détermine la qualité du résultat : soyez précis sur le quoi, laissez l’agent décider du comment (le fichier `AGENTS.md` cadre déjà les choix techniques). Voici la mission que nous utilisons :

```
Mission : Créer une application de gestion de tâches complète.
Besoins fonctionnels :
1. API REST FastAPI avec ces endpoints :
   - GET /taches : liste toutes les tâches
   - POST /taches : crée une tâche { "titre": str, "faite": bool }
   - PUT /taches/{id} : met à jour une tâche
   - DELETE /taches/{id} : supprime une tâche
2. Stockage SQLite dans taches.db, créé automatiquement au démarrage.
3. Une page index.html servie à la racine, avec un champ d'ajout,
   la liste des tâches, une case à cocher "faite" et un bouton supprimer.
4. Un fichier requirements.txt et un fichier de tests pytest couvrant
   la création et la suppression d'une tâche.
Respecte AGENTS.md. Propose d'abord un plan, puis implémente.
```
Sélectionnez Gemini 3 Pro pour cette mission (planification + génération), puis lancez. L’agent commence par produire un **plan d’implémentation** sous forme d’Artifact : arborescence des fichiers, ordre des tâches, choix de conception. Validez le plan (ou corrigez-le) avant l’exécution. L’agent écrit ensuite les fichiers un à un. Voici un extrait représentatif du backend généré :

```
# main.py – extrait généré par l'agent
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import sqlite3
app = FastAPI(title="Gestionnaire de taches")
def get_db():
    conn = sqlite3.connect("taches.db")
    conn.row_factory = sqlite3.Row
    return conn
def init_db():
    with get_db() as db:
        db.execute(
            "CREATE TABLE IF NOT EXISTS taches ("
            "id INTEGER PRIMARY KEY AUTOINCREMENT, "
            "titre TEXT NOT NULL, faite INTEGER DEFAULT 0)"
        )
class Tache(BaseModel):
    titre: str
    faite: bool = False
@app.on_event("startup")
def demarrage():
    init_db()
@app.get("/taches")
def lister_taches():
    """Retourne la liste complète des tâches."""
    with get_db() as db:
        lignes = db.execute("SELECT * FROM taches").fetchall()
        return [dict(l) for l in lignes]
@app.post("/taches")
def creer_tache(tache: Tache):
    """Crée une nouvelle tâche et renvoie son identifiant."""
    with get_db() as db:
        cur = db.execute(
            "INSERT INTO taches (titre, faite) VALUES (?, ?)",
            (tache.titre, int(tache.faite)),
        )
        return {"id": cur.lastrowid, "titre": tache.titre, "faite": tache.faite}
```
Remarquez que le code respecte le `AGENTS.md` : français, SQLite, docstrings, aucune dépendance surprise. L’agent poursuit avec les endpoints PUT et DELETE, la page `index.html`, le `requirements.txt` et les tests. Pendant qu’il travaille, il produit des Artifacts que nous allons maintenant lire.

