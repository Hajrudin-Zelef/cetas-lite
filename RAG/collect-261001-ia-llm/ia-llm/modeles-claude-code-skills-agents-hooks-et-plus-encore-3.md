---
id: collect-261001-ia-llm/ia-llm/modeles-claude-code-skills-agents-hooks-et-plus-encore-3
title: "Database Migration Skill"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "diffusion", "mcp"]
source: docs/RAG/collect-261001-ia-llm/modeles-claude-code-skills-agents-hooks-et-plus-encore.md
source_anchor: ""
source_lines: [278, 377]
sha256: 9511a3c7bd4417cd59013c7fc7d232d74cb285033fa6b8c38fc68cdcdcac702e
---

# Database Migration Skill

Pour en créer un de zéro, consultez le guide pas-à-pas des plugins Claude Code de DataCamp.

## Quel type de modèle choisir ?

On le comprend : tous ces types peuvent prêter à confusion. Ils modifient tous le comportement de Claude. La différence principale tient à leur mode de déclenchement et au degré de contrôle qu’ils apportent.

Pour clarifier, voici une comparaison entre eux :

| **Type de modèle** | **Déclenché par** | **Idéal pour** | **Peu adapté à** | **Cas d’usage** | 
| Skill | Claude automatiquement ou utilisateur via `/skill-name` | Workflows multi-étapes répétables | Tâches ponctuelles | Appliquer automatiquement une checklist de migration quand Claude modifie des fichiers de schéma | 
| Agent | Demande utilisateur ou délégation par Claude | Expertise métier et isolement des permissions | Sessions généralistes | Un auditeur sécurité qui peut lire les fichiers mais ne doit pas les éditer | 
| Commande | Commande slash utilisateur | Actions à la demande et points de contrôle | Garde-fous automatiques | /generate-tests quand vous êtes prêt à tester | 
| Hook | Événement du cycle de vie de Claude | Garde-fous et contrôles qualité automatisés | Tâches nécessitant un raisonnement interactif | Formater les fichiers après chaque édition | 
| MCP | Appel d’outil par Claude | Accès à des systèmes externes | Workflows simples limités au local | Interroger PostgreSQL ou créer un ticket GitHub | 
| Plugin | Installation ou activation | Diffusion équipe et workflows groupés | Ajustements locaux à but unique | Un bundle frontend avec agents, skills et hooks | 

Une règle simple pour décider :

- Si vous souhaitez **imposer automatiquement une règle** chaque fois que Claude touche au code, utilisez un**hook** .
- Si vous voulez que Claude **adopte une expertise métier poussée** pour une tâche précise, utilisez un**agent** .
- Si vous voulez **encoder un workflow** que Claude doit répéter fidèlement, utilisez un**skill** .
- Si vous voulez **déclencher vous-même une action** au bon moment, utilisez une**commande** ou un skill de type commande.
- Si Claude a besoin de **services externes ou de données en direct** , utilisez**MCP** .
- Si vous souhaitez installer ou partager une **configuration de workflow complète** , utilisez un**plugin** .

En pratique, ces types se combinent souvent. Un plugin sécurité peut regrouper un agent `security-auditor`, un skill `audit-findings`, une commande de contrôle des dépendances et un hook pré-commit. L’agent définit le rôle, le skill définit la structure du rapport, la commande offre un point de contrôle explicite, et le hook applique le garde-fou.

Pour les workflows où Claude doit suivre un plan formel avant mise en œuvre, le spéc-driven development est souvent plus adapté que des prompts au fil de l’eau.

## Où trouver des modèles Claude Code ?

Trois sources pratiques : Anthropic, les collections communautaires et vos propres créations.

Commencez par les ressources officielles d’Anthropic et la documentation. La documentation Claude Code d’Anthropic couvre les skills, sous-agents, hooks, MCP et plugins ; c’est l’endroit idéal pour vérifier les formats de fichiers et comportements à jour avant toute mise en production.

Ensuite, utilisez les collections communautaires. Le hub le plus visible est aitmpl.com, qui se présente comme un catalogue de configurations prêtes à l’emploi pour les projets Claude Code. Sa navigation propose actuellement Skills, Agents, Commands, Settings, Hooks, MCPs et Plugins.

La commande d’installation interactive actuelle est :

`npx claude-code-templates@latest`
La documentation du projet présente aussi un alias plus court :

`npx cct@latest`
Pour des composants spécifiques, le README GitHub en ligne propose des commandes d’installation comme :

```
npx claude-code-templates@latest --agent development-tools/code-reviewer --yes
npx claude-code-templates@latest --command performance/optimize-bundle --yes
npx claude-code-templates@latest --hook git/pre-commit-validation --yes
npx claude-code-templates@latest --mcp database/postgresql-integration --yes
```
Vous pouvez aussi installer un stack complet en lot via plusieurs options dans une seule commande.

En évaluant des modèles communautaires, vérifiez quelques signaux de qualité :

- 
La `description` est-elle assez précise pour que Claude déclenche correctement le skill ou l’agent ?
- 
Les `allowed-tools` sont-elles bien circonscrites, ou le modèle demande-t-il inutilement de larges permissions d’écriture et de bash ?
- 
Le dépôt est-il entretenu récemment ?
- 
Le modèle explique-t-il ce qu’il modifie ?
- 
Inclut-il des hooks ou des serveurs MCP qui exécutent du code que vous n’avez pas audité ?

Enfin, rédigez les vôtres. C’est souvent la meilleure option pour des workflows très liés à votre stack. Un modèle communautaire offre une bonne base, mais il ne connaît pas votre politique de migration interne, vos conventions de nommage, votre modèle de données ni votre tolérance au risque de déploiement.

## Pour conclure

Les modèles Claude Code permettent de passer d’un assistant à la session à un environnement de développement persistant.

Les six catégories présentées s’empilent en couches : les skills encodent les workflows, les agents définissent les rôles, les commandes créent des actions explicites, les hooks imposent des garde-fous, MCP connecte des systèmes externes et les plugins emballent le tout pour la réutilisation.

Le meilleur point de départ n’est pas un énorme bundle de plugins. Commencez par un skill pour votre workflow le plus répétitif. Dès que vous identifiez les frictions restantes dans le comportement par défaut de Claude, ajoutez un agent pour une revue spécialisée, un hook pour l’application automatique, ou un serveur MCP pour un accès en direct aux systèmes.

Pour aller plus loin avec Claude Code, découvrez nos cours Claude Code 101 et Claude Code in Action.

## FAQ sur les modèles Claude Code

### Les modèles Claude Code sont-ils identiques à CLAUDE.md ?

**Non. CLAUDE.md sert surtout à des consignes générales au niveau du projet : stack technique, conventions de code, structure du projet et commandes préférées. Les modèles Claude Code sont plus modulaires. Ils emballent des workflows, rôles, commandes, hooks ou intégrations que Claude peut utiliser au besoin.**

### Dois-je utiliser un skill ou un agent ?

**Utilisez un skill lorsque vous voulez que Claude suive un processus répétable, par exemple générer des tests, écrire des changelogs ou revoir des migrations. Utilisez un agent lorsque vous voulez que Claude adopte un rôle spécifique, par exemple auditeur sécurité, relecteur de documentation ou architecte frontend. Dans de nombreux workflows, vous pouvez employer les deux de concert.**

### Les modèles Claude Code sont-ils propres au projet ou globaux ?

**Les deux sont possibles, selon l’emplacement de stockage. Les modèles spécifiques à un projet se trouvent généralement dans le répertoire `.claude/` du projet. Des modèles globaux sont utiles si vous souhaitez le même comportement sur plusieurs projets.**

### Les modèles Claude Code communautaires sont-ils sûrs à installer ?

**Pas automatiquement. Les modèles communautaires peuvent être très utiles, mais ils peuvent inclure des permissions d’outils, des commandes shell, des hooks ou des configurations MCP ayant un impact sur votre environnement local.**

### Quel est le meilleur type de modèle pour démarrer ?

**Commencez par un skill. C’est généralement la façon la plus simple de transformer des consignes répétées en workflows réutilisables, sans complexifier votre setup. Une fois un premier skill utile en place, vous pourrez ajouter des agents, des hooks, MCP et des plugins.**

