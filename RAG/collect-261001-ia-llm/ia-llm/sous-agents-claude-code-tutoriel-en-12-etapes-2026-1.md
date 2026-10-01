---
id: collect-261001-ia-llm/ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026-1
title: "./scripts/validate-readonly-query.sh"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "mcp"]
source: docs/RAG/collect-261001-ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 50]
sha256: 5008fa30a9a8943467cb90d56a150022e634cf9deb056055a15710750fc3049b
---

# ./scripts/validate-readonly-query.sh

Claude Code sait déjà lire un dépôt entier, corriger un bug ou lancer une suite de tests. Mais sur un projet volumineux, chaque recherche de code, chaque journal de build et chaque page de documentation consultée s’accumule dans la même fenêtre de contexte que votre conversation principale. Au bout de quelques échanges, elle déborde de détails que vous ne relirez jamais. Les sous-agents corrigent ce défaut d’architecture : avec la version 2.1.271 de Claude Code, sortie le 14 septembre 2026 selon AI-TLDR, chaque sous-agent tourne toujours dans sa propre fenêtre de contexte, avec son propre prompt système et ses propres outils. Ce tutoriel montre, en 12 étapes, comment créer, configurer et déployer des sous-agents Claude Code personnalisés, avec la syntaxe exacte, les pièges à éviter et un projet complet à copier tel quel.

## Sous-agents Claude Code : ce qui change dans votre façon de coder

Un sous-agent est un assistant IA spécialisé qui traite un type de tâche précis dans sa propre fenêtre de contexte, défini comme un fichier Markdown avec un en-tête YAML qui fixe sa configuration et ses permissions. Quand Claude reconnaît une tâche qui correspond à la description d’un sous-agent, il lui délègue le travail. Le sous-agent l’exécute de façon indépendante, avec son propre prompt système, ses propres outils et des permissions indépendantes qui isolent son accès aux outils de celui de toute autre instance de sous-agent, puis renvoie uniquement un résumé à la conversation principale. C’est la définition exacte que donne la documentation officielle d’Anthropic sur les sous-agents personnalisés, mise à jour en août 2026.

Anthropic met en avant cinq bénéfices concrets pour justifier l’usage des sous-agents Claude Code : préserver le contexte principal en gardant l’exploration hors de la conversation, imposer des restrictions d’outils strictes, réutiliser une configuration d’un projet à l’autre grâce aux sous-agents de niveau utilisateur, spécialiser le comportement avec un prompt système ciblé, et maîtriser les coûts en renvoyant certaines tâches vers des modèles plus rapides comme Haiku. Ce dernier point compte double pour une équipe qui facture ses heures de développement assisté par IA au coût réel de l’API.

Claude Code embarque déjà trois sous-agents génériques prêts à l’emploi : Explore pour la recherche en lecture seule, Plan pour la préparation en mode plan, et general-purpose pour les tâches complexes qui combinent recherche et modification. Ce tutoriel va plus loin. Il montre comment écrire des sous-agents adaptés à votre pile technique et à vos conventions d’équipe, avec des exemples réels tirés de la documentation Anthropic : un relecteur de code, un débogueur, un analyste de données et un agent de requêtes en lecture seule protégé par un hook.

## Prérequis : les versions et outils nécessaires

Avant de commencer, la documentation officielle est claire sur un seul point de départ obligatoire : avoir Claude Code installé et authentifié, dans un terminal ouvert sur un répertoire de projet (un dépôt Git n’est pas obligatoire, un dossier vide fonctionne aussi). Si Claude Code n’est pas encore installé sur votre machine, notre tutoriel d’installation de Claude Code couvre la procédure complète en 13 étapes. Le tableau ci-dessous résume ce qu’il vous faut avant d’attaquer l’étape 1.

| Élément | Exigence | Pourquoi c’est nécessaire | 
|---|---|---|
| Claude Code CLI | Installé et authentifié | Les sous-agents sont une fonctionnalité native du CLI, pas un module séparé | 
| Version du CLI | 2.1.198 ou plus récente recommandée | Plusieurs comportements cités dans ce guide (modèle hérité par Explore, exécution en arrière-plan par défaut) datent de cette version | 
| Node.js | 18 ou plus récent | Nécessaire si vos sous-agents lancent des serveurs MCP via npx, comme l’exemple Playwright plus bas | 
| Terminal | macOS, Linux, WSL ou Windows PowerShell | La syntaxe `claude mcp` et `claude –agent` fonctionne à l’identique sur les quatre | 
| Éditeur de texte | N’importe lequel | Les sous-agents sont de simples fichiers Markdown avec en-tête YAML | 
| Dépôt Git (optionnel) | Recommandé pour les sous-agents de projet | Permet de versionner `.claude/agents/` et de le partager avec l’équipe | 

Un point de vocabulaire à clarifier avant de continuer : jusqu’à la version 2.1.63 de Claude Code, l’outil qui exécute les sous-agents s’appelait Task. Il a été renommé Agent depuis, mais les anciennes références `Task(…)` dans les paramètres et les définitions de sous-agents continuent de fonctionner comme alias.

## Étape 1 : repérer les sous-agents déjà intégrés à Claude Code

Avant d’écrire le moindre fichier, il vaut mieux savoir ce que Claude Code propose déjà. Trois sous-agents génériques sont enregistrés par défaut dans toute session interactive. Explore est un agent rapide, en lecture seule, optimisé pour la recherche et l’analyse de code : les outils Write et Edit lui sont refusés. Depuis la version 2.1.198, Explore hérite du modèle de la conversation principale au lieu de tourner systématiquement sur Haiku, avec un plafond à Opus sur l’API Claude. Plan sert de sous-agent de recherche pendant le mode plan, également en lecture seule. General-purpose, enfin, reste invocable à tout moment via l’outil Agent pour les tâches complexes qui demandent à la fois de l’exploration et des modifications, avec accès à tous les outils disponibles pour un sous-agent. Ces trois sous-agents intégrés, Explore, Plan et general-purpose, sont livrés en standard avec Claude Code pour couvrir chacun un type de tâche ciblé, comme le confirme la documentation à jour d’août 2026.

Deux agents auxiliaires s’y ajoutent en coulisse : statusline-setup (modèle Sonnet), utilisé quand vous lancez la commande `/statusline`, et claude-code-guide (modèle Haiku), sollicité pour répondre à vos questions sur les fonctionnalités de Claude Code. Vous n’avez normalement pas à les invoquer directement.

Un détail utile pour le dépannage plus loin dans ce guide : Explore et Plan sont les deux seuls sous-agents à ignorer vos fichiers CLAUDE.md et l’état Git de la session parente, afin de garder la recherche rapide et économique. Tous les autres sous-agents, intégrés ou personnalisés, chargent ces deux éléments au démarrage. Si vous voulez bloquer un sous-agent intégré précis, ajoutez-le au tableau `deny` de vos permissions avec la syntaxe `Agent(Explore)`. Pour empêcher Claude de déléguer à n’importe quel sous-agent, refusez directement l’outil Agent.

## Étape 2 : créer votre premier sous-agent personnalisé

Un sous-agent Claude Code est un fichier Markdown avec un en-tête YAML. Deux méthodes permettent de le créer : demander à Claude de l’écrire, ou rédiger le fichier vous-même.

### Méthode A : demander à Claude d’écrire le fichier

Dans une session Claude Code, décrivez le sous-agent voulu et son emplacement :

```
Crée un sous-agent personnel code-improver dans ~/.claude/agents/ qui
scanne les fichiers et suggère des améliorations de lisibilité, de
performance et de bonnes pratiques. Il doit expliquer chaque problème,
montrer le code actuel et fournir une version améliorée. Rends-le en
lecture seule et fais-le tourner sur Sonnet.
```
Claude écrit alors le fichier avec un `name`, une `description`, une liste `tools`, un `model` et un prompt système. Le résultat ressemble à ceci :

