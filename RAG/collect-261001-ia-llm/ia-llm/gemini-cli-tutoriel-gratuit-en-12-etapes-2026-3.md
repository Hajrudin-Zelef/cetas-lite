---
id: collect-261001-ia-llm/ia-llm/gemini-cli-tutoriel-gratuit-en-12-etapes-2026-3
title: "Vérifier les versions actuelles"
domain: ia-llm
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agent", "agents", "gemini", "mcp", "memory", "model context protocol"]
source: docs/RAG/collect-261001-ia-llm/gemini-cli-tutoriel-gratuit-en-12-etapes-2026.md
source_anchor: ""
source_lines: [158, 268]
sha256: d01e5b4364a10e2d5d14bf443ec988ef20503010c4e05251c22dfd827bc366b9
---

# Vérifier les versions actuelles

Le fichier `GEMINI.md` est la clé d’un agent réellement utile. Il s’agit d’un fichier de **contexte persistant** – des instructions que Gemini CLI charge automatiquement à chaque session pour comprendre votre projet : conventions de code, architecture, commandes de build, style attendu. C’est l’équivalent d’un onboarding permanent pour votre assistant.

La façon la plus rapide d’en générer un est la commande `/init` : l’agent analyse le dossier courant et produit un `GEMINI.md` sur mesure. Vous pouvez ensuite l’affiner à la main. Le chargement est **hiérarchique** : Gemini CLI fusionne un fichier global (`~/.gemini/GEMINI.md`), un fichier à la racine du projet, et d’éventuels fichiers dans les sous-dossiers, du plus général au plus spécifique.

```
# Générer automatiquement un fichier de contexte
> /init
# Exemple de GEMINI.md rédigé à la main
# ---------------------------------------------
# Projet : API de facturation (Node.js / Express 5)
#
# ## Conventions
# - TypeScript strict, pas de "any".
# - Tests avec Vitest ; couvrir chaque route.
# - Messages de commit en français, format Conventional Commits.
#
# ## Commandes
# - Installer : npm ci
# - Développer : npm run dev
# - Tester : npm test
# - Lint : npm run lint
#
# ## Architecture
# - src/routes  : points d'entrée HTTP
# - src/services: logique métier
# - src/db      : accès Postgres via le client pg
#
# ## À NE PAS FAIRE
# - Ne jamais logger de données personnelles (RGPD).
# - Ne pas modifier les migrations déjà appliquées.
```
Gérez ces fichiers avec la commande `/memory` : `/memory show` affiche le contexte chargé, `/memory refresh` le recharge après modification. Un bon `GEMINI.md` transforme radicalement la qualité des réponses : l’agent respecte vos conventions, utilise vos commandes et évite les pièges spécifiques à votre code. C’est l’investissement de cinq minutes le plus rentable de tout ce tutoriel.

## Étape 7 – Maîtriser les commandes slash essentielles

Les commandes slash contrôlent la session sans interrompre votre flux de travail. Il en existe une quarantaine, mais une douzaine couvre 95 % des besoins quotidiens. Mémorisez-les : elles vous feront gagner un temps considérable, notamment pour gérer le contexte (et donc votre consommation de tokens) sur les longues sessions.

| Commande | Effet | 
|---|---|
| `/help` | Affiche l’aide et la liste des commandes disponibles | 
| `/tools` | Liste les outils intégrés (fichiers, shell, recherche, web) | 
| `/memory` | Gère le contexte issu des fichiers GEMINI.md (show / refresh / add) | 
| `/init` | Analyse le dossier et génère un fichier de contexte adapté | 
| `/model` | Affiche et change le modèle Gemini utilisé | 
| `/mcp` | Gère les serveurs MCP (liste, authentification, rechargement) | 
| `/compress` | Remplace le contexte par un résumé pour économiser des tokens | 
| `/restore` | Restaure les fichiers à leur état avant une action de l’agent | 
| `/rewind` | Revient en arrière dans l’historique de la conversation | 
| `/stats` | Affiche les statistiques de session, de modèle et d’outils | 
| `/chat` | Sauvegarde, liste et reprend des points de contrôle de conversation | 
| `/clear` | Efface l’écran et l’historique visible de la session | 
| `/theme` | Change le thème visuel de l’interface | 
| `/quit` | Quitte Gemini CLI | 

Deux réflexes à adopter. Sur une session longue qui commence à ralentir, tapez `/compress` : l’agent remplace tout l’historique par un résumé compact, ce qui libère de la fenêtre de contexte et réduit la consommation. Et pour repartir sur un sujet totalement différent, `/clear` remet la session à zéro. Ces deux commandes évitent l’écueil du contexte saturé, principale cause de réponses lentes ou hors-sujet.

## Étape 8 – Changer de modèle : Gemini 3 Pro contre Flash

En 2026, Gemini CLI donne accès à deux familles de modèles : **Gemini 3 Pro**, taillé pour le raisonnement complexe et le code difficile, et **Gemini 3 Flash**, plus rapide et plus économe, idéal pour les tâches répétitives. Par défaut, l’agent utilise un mode « Auto » qui route intelligemment vos requêtes vers le modèle adapté. Vous pouvez toutefois forcer un choix.

```
# Choisir le modèle de façon interactive
> /model
# Forcer Gemini 3 Pro au lancement (drapeau -m / --model)
gemini -m gemini-3.1-pro-preview
# Lancer directement une requête sur un modèle donné
gemini -m gemini-3.1-pro-preview -p "Analyse la complexité de src/parser.ts"
```
Quand privilégier chaque modèle ? Réservez **Pro** aux refactorisations d’envergure, au débogage d’algorithmes ardus, à la conception d’architecture et à l’analyse de code volumineux où la profondeur de raisonnement fait la différence. Basculez sur **Flash** pour la génération de tests unitaires, la rédaction de documentation, les commits, les petites corrections et tout travail par lots où la vitesse et le coût priment. Le tableau ci-dessous résume l’arbitrage.

| Aspect | Gemini 3 Pro | Gemini 3 Flash | 
|---|---|---|
| Point fort | Raisonnement profond, code complexe | Vitesse, coût réduit | 
| Latence | Plus élevée | Faible | 
| Accès gratuit | Selon quota et routage | Inclus dans tous les paliers | 
| Cas idéal | Refactor, architecture, débogage | Tests, docs, tâches par lots | 
| Contexte | 1 M tokens | 1 M tokens | 

Notez que la disponibilité exacte des modèles Pro sur le palier gratuit dépend du routage et des quotas en vigueur, susceptibles d’évoluer. Pour un accès garanti et prioritaire aux modèles Pro, activez la facturation sur votre clé API ou passez par Vertex AI. Pour une comparaison des modèles côté API, consultez la documentation officielle des modèles Gemini.

## Étape 9 – Étendre Gemini CLI avec les serveurs MCP

Le **Model Context Protocol (MCP)** est un standard ouvert qui connecte les agents IA à des outils et sources de données externes : bases de données, API, systèmes de fichiers, dépôts Git, services tiers. Gemini CLI supporte MCP nativement. En déclarant un serveur MCP, vous donnez à l’agent de nouvelles capacités sans modifier son code. C’est le mécanisme d’extension le plus puissant de l’outil.

La configuration se fait dans un fichier `settings.json`. Trois emplacements existent, du plus général au plus spécifique : `~/.gemini/settings.json` (utilisateur), `.gemini/settings.json` (projet) et une configuration système. Ajoutez un bloc `mcpServers` décrivant chaque serveur par sa commande de lancement, ses arguments et ses variables d’environnement.

```
// .gemini/settings.json – déclaration de serveurs MCP
{
  "theme": "Default Dark",
  "mcpServers": {
    "filesystem": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-filesystem", "./data"]
    },
    "github": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-github"],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "$GITHUB_TOKEN"
      }
    }
  }
}
```
Après avoir enregistré le fichier, relancez Gemini CLI ou tapez `/mcp` pour vérifier que les serveurs sont bien chargés et voir les outils qu’ils exposent. L’agent peut désormais, par exemple, lire des fichiers d’un dossier autorisé via le serveur *filesystem*, ou ouvrir une pull request via le serveur *github*. Pour découvrir l’écosystème complet de connecteurs, référez-vous à la spécification officielle du Model Context Protocol.

**Bonne pratique :** n’ajoutez que des serveurs MCP de confiance et accordez-leur le périmètre minimal (un dossier précis plutôt que tout le disque). Un serveur MCP mal configuré peut donner à l’agent un accès trop large. Le principe du moindre privilège s’applique ici comme partout ailleurs en sécurité.

## Étape 10 – Automatiser : mode non-interactif et scripts

