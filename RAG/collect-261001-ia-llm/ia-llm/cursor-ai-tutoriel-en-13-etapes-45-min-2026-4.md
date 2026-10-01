---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026-4
title: "AGENTS.md"
domain: ia-llm
role: reference
task: reference
actors: ["Apple", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "cloud agent", "copilot", "diffusion", "mcp"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026.md
source_anchor: ""
source_lines: [196, 280]
sha256: 30de47e36a3c88f7472a65ebd36332c4bfac97d801177760f481b0aaefc163f7
---

# AGENTS.md

Les Cloud Agents, anciennement appelés Background Agents, changent d’échelle par rapport au mode Agent local. D’après la documentation officielle, ils **« reposent sur les mêmes principes que l’agent classique, mais s’exécutent dans des machines virtuelles isolées dans le cloud, avec un environnement de développement complet, plutôt que sur la machine locale »**. Cela veut dire un dépôt cloné, les dépendances installées, les secrets nécessaires et un accès réseau, le tout sans mobiliser votre ordinateur.

Un Cloud Agent se déclenche de plusieurs façons : depuis l’application de bureau en sélectionnant “Cloud” dans le menu de l’agent, depuis l’interface web sur cursor.com/agents, depuis l’application iOS, ou directement depuis Slack, GitHub, Bitbucket ou Linear en mentionnant `@cursor`. Sur une pull request GitHub par exemple :

```
@cursor corrige le bug décrit dans ce ticket, ajoute un test de
non-régression, et pousse le correctif sur une nouvelle branche
```
L’agent clone le dépôt, travaille sur une branche séparée, exécute les tests, puis pousse ses modifications sous forme de pull request. Plusieurs agents peuvent tourner en parallèle sans bloquer votre machine, et chacun fournit des artefacts (captures d’écran, vidéos, journaux d’exécution) pour vérifier le travail effectué. Il est même possible de prendre la main à distance sur le bureau de l’agent pour tester une modification avant de la fusionner.

Un usage typique en équipe consiste à laisser tourner un Cloud Agent la nuit sur des tâches de fond : mise à jour de dépendances mineures, correction de tests instables, ou génération d’un rapport de couverture de code. Le développeur retrouve le lendemain matin une pull request déjà ouverte, avec les artefacts d’exécution attachés, plutôt qu’une tâche encore à démarrer.

## Projet complet : construire une API de bout en bout avec l’agent Cursor

Pour relier toutes les étapes précédentes, voici un projet minimal mais complet : une route d’API qui valide une inscription utilisateur, construite avec FastAPI et pilotée par le mode Agent de Cursor AI.

Première étape, créez la règle de projet qui cadre le style attendu :

```
---
description: "Conventions API FastAPI"
globs: app/**/*.py
alwaysApply: true
---
Toutes les routes utilisent des schémas Pydantic pour la validation d'entrée et de sortie.
Les erreurs métier lèvent une HTTPException avec un code et un message explicites.
Chaque nouvelle route s'accompagne d'un test pytest dans le dossier tests/.
```
Ensuite, en mode Agent, le prompt suivant suffit à démarrer le travail :

*“Crée une route POST /users qui valide un email et un mot de passe (12 caractères minimum, au moins un chiffre), retourne une erreur 422 explicite si la validation échoue, et ajoute un test pytest qui couvre le cas valide et deux cas d’échec.”*

L’agent lit d’abord la structure existante du projet, applique la règle définie plus haut, puis génère un résultat proche de celui-ci :

```
from fastapi import APIRouter, HTTPException
from pydantic import BaseModel, EmailStr, field_validator
router = APIRouter()
class UserCreate(BaseModel):
    email: EmailStr
    password: str
    @field_validator("password")
    @classmethod
    def password_strength(cls, value: str) -> str:
        if len(value) < 12 or not any(char.isdigit() for char in value):
            raise ValueError("Le mot de passe doit contenir 12 caractères et un chiffre minimum")
        return value
@router.post("/users", status_code=201)
def create_user(payload: UserCreate):
    if payload.email.endswith("@exemple-interdit.com"):
        raise HTTPException(status_code=422, detail="Domaine email refusé")
    return {"email": payload.email, "status": "créé"}
```
Dernière étape, relisez le diff proposé, lancez les tests en local pour confirmer, puis démarrez le serveur :

```
pytest -v
uvicorn app.main:app --reload
curl -X POST http://localhost:8000/users \
  -H "Content-Type: application/json" \
  -d '{"email":"[email protected]","password":"motdepasse123"}'
```
Ce cycle complet, règle de projet puis prompt agentique puis relecture du diff, reste le même quelle que soit la taille de la tâche. C'est cette boucle qu'il faut automatiser dans vos réflexes pour tirer parti de Cursor AI sans accumuler de dette technique invisible. Pour aller plus loin sur ce même projet, la suite logique consiste à demander à l'agent d'ajouter un endpoint de connexion avec génération de jeton, puis de connecter un serveur MCP vers votre base de données de test pour qu'il vérifie lui-même le schéma des tables existantes avant d'écrire une nouvelle migration.

## Cursor AI vs GitHub Copilot vs Windsurf : comparatif et tarifs 2026

Trois éditeurs IA dominent les discussions des équipes de développement en 2026. Voici comment ils se comparent sur les critères qui comptent le plus au moment de choisir.

| Critère | Cursor AI | GitHub Copilot | Windsurf | 
|---|---|---|---|
| Base technique | Fork complet de VS Code | Extension pour VS Code, JetBrains, Neovim | Fork de VS Code (marque conservée sous Cognition AI) | 
| Entrée de gamme | Gratuit (Hobby) | Gratuit (2 000 complétions/mois, 50 requêtes chat) | Gratuit | 
| Palier individuel courant | 20 $/mois (Pro) | 10 $/mois (Pro) | environ 20 $/mois (Pro, depuis mars 2026) | 
| Palier le plus élevé (individuel) | 200 $/mois (Ultra) | 100 $/mois (Max) | 200 $/mois (Max) | 
| Palier équipe | 40 $/utilisateur/mois | 19 $/utilisateur/mois (Business) | 40 $/utilisateur/mois | 
| Mode Agent natif | Oui | Oui | Oui | 
| Support MCP | Oui, natif | Oui, natif | Partiel selon les sources sectorielles | 
| Société éditrice | Anysphere | Microsoft / GitHub | Cognition AI (racheté fin 2025) | 

GitHub Copilot conserve un avantage de diffusion difficile à battre : selon la page officielle de GitHub Copilot, l'outil revendique "des millions d'utilisateurs individuels et des dizaines de milliers de clients entreprise", ce qui en fait l'outil IA pour développeurs le plus adopté au monde. GitHub cite aussi des gains mesurés en interne : jusqu'à 55 % de productivité en plus sur l'écriture de code sans perte de qualité, et jusqu'à 75 % de satisfaction professionnelle supplémentaire chez les développeurs qui l'utilisent. L'entreprise brésilienne Grupo Boticário est citée en exemple avec une hausse de productivité développeur de 94 % après adoption. Si vous démarrez tout juste avec l'IA en entreprise, notre tutoriel GitHub Copilot détaille sa propre installation pas à pas.

Windsurf occupe une position plus mouvementée : après l'échec d'un rachat par OpenAI à 3 milliards de dollars mi-2025 et un accord de licence de 2,4 milliards de dollars signé par Google DeepMind pour récupérer l'équipe fondatrice, la marque et la base de clients sont finalement passées sous le contrôle de Cognition AI, la société derrière l'agent autonome Devin, fin 2025. Les tarifs ont changé en mars 2026 avec un passage à des quotas d'usage plutôt qu'à des crédits fixes, ce qui rend certaines comparaisons de prix datées rapidement obsolètes. Pour une équipe qui évalue les trois options, mieux vaut vérifier les tarifs affichés le jour même sur chaque site officiel plutôt que de se fier à un comparatif figé, y compris celui-ci.

Le vrai critère de choix reste l'usage. Une équipe qui veut le maximum d'autonomie agentique et une gestion fine des règles de contexte penche vers Cursor AI. Une équipe déjà largement équipée en environnement Microsoft et JetBrains, avec un budget serré, reste souvent sur GitHub Copilot. Et pour les cas où le code ne doit jamais quitter l'infrastructure de l'entreprise, l'exécution locale via Ollama mérite d'être évaluée en complément, même si elle sacrifie une partie des capacités agentiques cloud.

