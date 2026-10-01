---
id: collect-261001-ia-llm/ia-llm/serveur-mcp-pour-claude-code-tutoriel-2026-2
title: "La sortie doit afficher Python 3.10.x ou une version supérieure"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "agents", "chatgpt", "claude", "copilot", "distribution", "gemini", "mai", "mcp"]
source: docs/RAG/collect-261001-ia-llm/serveur-mcp-pour-claude-code-tutoriel-2026.md
source_anchor: ""
source_lines: [25, 72]
sha256: 95a824ef179f015b8a8589b24072337b0d11e7f74646e1f30baa3fbe58a82615
---

# La sortie doit afficher Python 3.10.x ou une version supérieure

| Date | Étape de l’adoption | 
|---|---|
| 25 novembre 2024 | Anthropic publie MCP comme standard ouvert | 
| Mars 2025 | OpenAI adopte MCP dans ChatGPT Desktop, son kit Agents et l’API Responses | 
| Avril 2025 | Google DeepMind confirme la prise en charge de MCP pour ses modèles Gemini | 
| Mai 2025 | Microsoft annonce un support préliminaire à Build 2025, dans Windows 11, Copilot Studio et Azure OpenAI Service | 
| Septembre à novembre 2025 | Le registre de serveurs MCP progresse de 407 %, pour approcher 2 000 entrées | 
| Décembre 2025 | Anthropic transfère la gouvernance de MCP à l’Agentic AI Foundation, hébergée par la Linux Foundation | 
| Début 2026 | 28 % des entreprises du Fortune 500 ont déployé des serveurs MCP en production | 
| Mi-2026 | Plus de 10 000 serveurs MCP publics actifs recensés par Anthropic ; environ 97 millions de téléchargements mensuels cumulés pour les kits Python et TypeScript | 

Ce calendrier explique pourquoi la question ne se limite plus à l’écosystème Anthropic. OpenAI a rallié MCP dès mars 2025 selon l’article de TechCrunch consacré à cette annonce, un choix notable puisqu’il s’agit d’un protocole conçu par un concurrent direct. Nous avions d’ailleurs détaillé le fonctionnement de Codex CLI, l’agent de codage en ligne de commande d’OpenAI, qui bénéficie lui aussi de cette compatibilité croissante. Du côté des chiffres d’usage en entreprise, le rapport State-of-AI 2026 évoque une adoption de MCP dans 74 % des systèmes d’IA en production, avec une réduction du temps d’intégration de nouveaux outils de 73 % et des coûts de maintenance annuels réduits de 62 % par rapport à des connecteurs propriétaires développés en interne, selon les données reprises par l’analyse de WorkOS sur l’état de MCP en 2026.

Ces gains se retrouvent aussi côté déploiement d’agents. Début 2026, 80 % des entreprises du Fortune 500 font tourner des agents IA actifs en production, et 78 % des équipes IA d’entreprise déclarent au moins un agent adossé à MCP parmi ces déploiements. Le nombre exact de serveurs publics varie toujours sensiblement selon la source consultée, mais toutes convergent vers un ordre de grandeur nettement supérieur à celui de 2025 : Anthropic évoquait plus de 10 000 serveurs dans son annonce officielle, un suivi indépendant publié sur dev.to situait la fourchette entre 9 400 et plus de 17 000 serveurs selon les registres pris en compte pour environ 300 applications clientes recensées, le registre communautaire Glama en dénombrait 22 775 dès mai 2026, et PulseMCP en recensait 22 311 en juillet 2026, contre environ 14 000 seulement deux mois plus tôt, d’après ShareuHack. Cet écart tient surtout à la méthode de comptage (registre central unique ou somme de plusieurs registres communautaires), mais l’ordre de grandeur reste le même : le nombre de serveurs MCP publics a été multiplié par plusieurs dizaines depuis le lancement du protocole fin 2024. Cette dispersion des chiffres devrait toutefois s’atténuer : depuis le 9 décembre 2025, la gouvernance de MCP est officiellement passée sous l’Agentic AI Foundation, hébergée par la Linux Foundation, selon Praxena, et le registre officiel doit basculer avant avril 2026 vers une adresse unique, registry.modelcontextprotocol.io, coportée par Anthropic, GitHub, Microsoft et PulseMCP — de quoi, à terme, remplacer la mosaïque actuelle de registres communautaires comme Glama, ShareuHack ou le MCP Server Hub, qui recensait encore des services isolés comme Amap Maps au 30 décembre 2025.

## Prérequis techniques : versions, comptes et outils nécessaires

Avant de lancer la moindre commande, vérifiez les éléments suivants. Un environnement mal préparé reste la cause la plus fréquente des blocages recensés dans la section dépannage plus bas.

| Élément | Exigence minimale | Recommandation | 
|---|---|---|
| Python | Version 3.10 | Version 3.12 ou plus récente | 
| Node.js (pour l’Inspecteur MCP) | Version 18 | Version 20 ou supérieure | 
| Claude Code | Installé et authentifié | Dernière version disponible | 
| Gestionnaire de paquets Python | pip | uv, plus rapide pour isoler l’environnement | 
| Système d’exploitation | Windows 10, macOS 12 ou distribution Linux récente | Dernière version stable de l’OS | 
| Git | Recommandé (utilisé par le projet exemple) | Dernière version stable | 
| Éditeur de code | Facultatif | Utile pour relire le code du serveur | 
| Espace disque | 200 Mo disponibles | 1 Go si plusieurs serveurs testés en parallèle | 

Deux précisions méritent d’être posées avant de continuer. D’abord, Python et Node.js jouent chacun un rôle distinct ici : le serveur lui-même est écrit en Python dans ce tutoriel, tandis que Node.js n’intervient que pour exécuter l’Inspecteur MCP, un outil de test en ligne de commande distribué comme paquet npm. Vous pouvez très bien écrire un serveur MCP entièrement en TypeScript si vous préférez rester dans un seul écosystème, le kit de développement officiel existant dans les deux langages. Ensuite, si Claude Code n’est pas encore installé sur votre machine, notre tutoriel dédié à l’installation de Claude Code en 13 étapes couvre ce prérequis avant de revenir ici.

## Vue d’ensemble du projet : le serveur MCP que nous allons construire

Plutôt qu’un exemple jouet du type “bonjour le monde”, ce tutoriel construit un serveur réellement utile au quotidien : un serveur MCP baptisé `git-journal`, qui donne à Claude Code une vision structurée de l’activité récente d’un dépôt Git. Concrètement, il expose deux outils et une ressource.

Le premier outil, `derniers_commits`, retourne l’historique récent d’un dépôt sous une forme lisible par l’agent (hash court, auteur, date, message). Le second, `chercher_todo`, parcourt le code source à la recherche de commentaires `TODO` et `FIXME` laissés par l’équipe. La ressource, exposée sous l’identifiant `git://status`, donne l’état courant du dépôt, équivalent à un `git status --short`. Une fois ce serveur connecté, une simple demande comme “résume les cinq derniers commits et signale les TODO non traités” devient une requête que Claude Code peut satisfaire seul, sans que vous ayez à copier-coller la sortie de plusieurs commandes Git dans le terminal au préalable.

Ce choix d’exemple n’est pas anodin. La recherche menée en amont de ce tutoriel confirme que le raisonnement multi-fichiers assisté par une intégration Git profonde figure parmi les usages qui distinguent le plus nettement Claude Code d’un simple auto-complétion : plus l’agent dispose d’un contexte fiable sur l’historique et l’état du dépôt, moins il a besoin de deviner ou de redemander une information déjà présente ailleurs dans le projet. Un serveur MCP dédié au Git constitue ainsi un point d’entrée pédagogique idéal avant de s’attaquer à des intégrations plus ambitieuses, vers un outil de tickets ou une base de données interne par exemple.

## Étapes 1 à 3 : préparer l’environnement Python et l’espace de travail

Les trois premières étapes se jouent avant même d’installer le kit de développement MCP. Elles évitent la majorité des erreurs recensées plus loin dans cet article.

- **Étape 1** : vérifiez la version de Python installée sur votre machine.
- **Étape 2** : créez un dossier de projet dédié et un environnement virtuel Python isolé.
- **Étape 3** : vérifiez Node.js et la présence de Claude Code, tous deux nécessaires plus loin dans ce tutoriel.

