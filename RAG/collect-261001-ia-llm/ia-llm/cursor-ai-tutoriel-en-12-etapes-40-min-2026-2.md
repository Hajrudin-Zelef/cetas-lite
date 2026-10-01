---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-12-etapes-40-min-2026-2
title: "Windows (winget)"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Mistral", "Moonshot", "OpenAI", "Z.ai", "xAI"]
dates: []
keywords: ["agent", "chatgpt", "claude", "gemini", "glm", "grok", "grok 4", "kimi", "mistral", "open-weight", "opus 4", "sonnet 5"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [41, 153]
sha256: 910fdf5f70f9b42a67b3939944dc3d9fba26e0e3aaeccc59f7c3f7331c70b58b
---

# Windows (winget)

À l’issue des 12 étapes, vous disposerez d’une API exposant les opérations CRUD (créer, lire, mettre à jour, supprimer) sur des tâches, documentée automatiquement via OpenAPI, couverte par des tests `pytest`, et prête à être versionnée. L’arborescence cible ressemblera à ceci :

```
gestion-taches/
├── .cursor/
│   └── rules/
│       └── regles-projet.mdc     # règles de projet pour l'agent
├── .cursorignore                 # fichiers exclus de l'indexation
├── app/
│   ├── __init__.py
│   ├── main.py                   # point d'entrée FastAPI
│   ├── models.py                 # modèles Pydantic + SQLModel
│   └── database.py               # connexion SQLite
├── tests/
│   └── test_taches.py            # tests pytest
├── requirements.txt
└── README.md
```
## Étape 1 – Télécharger et installer Cursor

Rendez-vous sur le site officiel cursor.com et cliquez sur *Download*. Le site détecte votre système et propose l’installeur adapté. Vous pouvez aussi passer par un gestionnaire de paquets, plus pratique pour scripter l’installation ou la maintenir à jour.

```
# Windows (winget)
winget install Anysphere.Cursor
# macOS (Homebrew)
brew install --cask cursor
# Linux : télécharger l'AppImage depuis cursor.com, puis
chmod +x Cursor-*.AppImage
./Cursor-*.AppImage
```
Sous Windows, l’installeur classique `.exe` reste le plus simple : double-cliquez et suivez l’assistant. Sur macOS, glissez l’application dans le dossier *Applications*. Sous Linux, l’AppImage est la voie officielle, mais des paquets `.deb` et `.rpm` existent également. Au premier lancement, Cursor vérifie les mises à jour : l’éditeur suit un rythme de versions très soutenu (quasi hebdomadaire) depuis la série 2.x, aussi acceptez systématiquement les mises à jour proposées pour bénéficier des derniers modèles.

### Piège n°1 : ne pas confondre Cursor et l’extension VS Code

Cursor est un logiciel **autonome**, pas une extension à installer dans VS Code. Vous n’avez pas besoin d’avoir VS Code au préalable. Beaucoup de débutants cherchent « Cursor » dans la place de marché des extensions VS Code et ne trouvent rien : c’est normal, il faut télécharger l’application depuis cursor.com.

## Étape 2 – Importer VS Code, se connecter et choisir son modèle

Au premier démarrage, Cursor propose d’**importer votre configuration VS Code** : extensions, thèmes, raccourcis clavier et paramètres. Acceptez si vous venez de VS Code, l’expérience sera immédiatement familière. Connectez-vous ensuite avec GitHub, Google ou une adresse e-mail. La connexion est nécessaire car les modèles s’exécutent dans le cloud.

Ouvrez ensuite le sélecteur de modèle (en haut du panneau de discussion). En 2026, **Cursor** donne accès à un catalogue impressionnant de grands modèles de langage, que vous pouvez changer à la volée selon la tâche. Les principaux disponibles à la mi-2026 sont résumés ci-dessous.

| Fournisseur | Modèles disponibles dans Cursor (2026) | Usage type | 
|---|---|---|
| Anysphere (maison) | Composer 2, Composer 2.5 | Vitesse maximale, tâches agentiques | 
| Anthropic | Claude Opus 4.8, Claude Sonnet 5, Claude 4.7 Opus | Raisonnement, refactorisations complexes | 
| OpenAI | GPT-5.5, GPT-5.x Codex | Génération de code polyvalente | 
|  | Gemini 3.1 Pro, Gemini 3.5 Flash | Grands contextes, rapidité | 
| xAI | Grok 4.3 | Alternative généraliste | 
| Autres | GLM 5.2 (Z.ai), Kimi K2.5 (Moonshot) | Options open-weight / coût réduit | 

Notre conseil pratique : utilisez **Composer 2** (ou le mode « Auto ») pour les tâches courantes où la réactivité prime, et basculez sur **Claude Opus 4.8** ou **Claude Sonnet 5** pour les refactorisations délicates ou le débogage difficile. Pour approfondir les différences entre ces grands modèles, consultez notre comparatif Claude vs ChatGPT vs Gemini vs Mistral.

## Étape 3 – Maîtriser l’interface : Tab, Cmd+K, Chat et Agent

Quatre points d’entrée résument l’expérience **Cursor**. Les mémoriser dès le départ change radicalement votre productivité. Sur macOS, la touche *Cmd* remplace *Ctrl*.

- **Tab** – l’autocomplétion prédictive. Cursor ne complète pas seulement la ligne en cours, il prédit votre prochaine modification (y compris à plusieurs lignes) et vous laisse l’accepter d’une pression sur*Tab* .
- **Ctrl/Cmd + K** – l’édition en ligne. Sélectionnez du code, appuyez sur K, décrivez la modification en langage naturel : Cursor réécrit la sélection sur place.
- **Ctrl/Cmd + L** – le Chat. Un panneau latéral pour poser des questions sur votre code, coller des erreurs ou demander des explications.
- **Ctrl/Cmd + I** – l’Agent (Composer). Le mode le plus puissant : l’IA planifie, crée et modifie plusieurs fichiers, exécute des commandes et itère jusqu’au résultat.

Prenez cinq minutes pour ouvrir un fichier de test et essayer chacun de ces raccourcis. La bascule entre le mode *Ask* (Cursor répond sans rien modifier) et le mode *Agent* (Cursor agit sur le projet) se fait directement dans le panneau : c’est un réflexe essentiel pour garder la maîtrise de ce que fait l’outil.

## Étape 4 – Définir les règles du projet avec .cursor/rules

C’est l’étape que les débutants sautent, et c’est une erreur. Les **règles de projet** (Project Rules) sont des instructions persistantes qui orientent chaque réponse de l’IA dans votre dépôt. Elles évitent que l’agent parte dans tous les sens et réduisent drastiquement le nombre d’allers-retours.

L’ancienne approche reposait sur un fichier unique `.cursorrules` à la racine. Depuis 2025, la méthode recommandée utilise des fichiers `.mdc` (Markdown enrichi) placés dans `.cursor/rules/`, ce qui permet d’organiser des règles par type de fichier ou par contexte. Créez le fichier suivant :

```
# .cursor/rules/regles-projet.mdc
---
description: Règles pour l'API de gestion de tâches
alwaysApply: true
---
- Langage : Python 3.12, framework FastAPI.
- Utiliser SQLModel pour les modèles et la persistance SQLite.
- Typage strict : annoter tous les paramètres et retours de fonction.
- Écrire les commentaires et docstrings EN FRANÇAIS.
- Chaque endpoint doit renvoyer des codes HTTP explicites.
- Générer des tests pytest pour toute nouvelle route.
- Ne jamais coder en dur de secrets : utiliser des variables d'environnement.
```
Ajoutez également un fichier `.cursorignore` pour empêcher l’indexation de fichiers volumineux ou sensibles (l’équivalent d’un `.gitignore` pour l’IA). Cela accélère l’indexation et protège vos secrets.

```
# .cursorignore
.venv/
__pycache__/
*.db
.env
node_modules/
dist/
```
## Étape 5 – Générer le squelette du projet avec l’Agent

Place à la magie. Ouvrez l’Agent (*Ctrl/Cmd + I*), vérifiez que vous êtes bien en mode **Agent** et non *Ask*, puis saisissez une consigne précise. La qualité du résultat dépend directement de la clarté de la demande : décrivez le quoi, le comment et les contraintes.

```
Crée la structure d'une API REST de gestion de tâches avec FastAPI.
Contraintes :
- Utilise SQLModel + SQLite (fichier taches.db).
- Modèle Tache : id, titre (str), terminee (bool, défaut False),
  date_creation (datetime auto).
- Endpoints CRUD complets : POST /taches, GET /taches,
  GET /taches/{id}, PATCH /taches/{id}, DELETE /taches/{id}.
- Crée app/main.py, app/models.py, app/database.py, requirements.txt.
- Respecte les règles du projet (.cursor/rules).
```
Cursor affiche un **plan**, puis propose les fichiers à créer sous forme de diff. Vous validez ou refusez chaque modification. En quelques secondes, l’agent génère une base cohérente. Voici, par exemple, le type de code produit pour `app/models.py` :

