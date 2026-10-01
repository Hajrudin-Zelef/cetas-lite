---
id: collect-261001-ia-llm/ia-llm/serveur-mcp-pour-claude-code-tutoriel-2026-1
title: "La sortie doit afficher Python 3.10.x ou une version supérieure"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft"]
dates: ["2025-11-25", "2026-07-28"]
keywords: ["agent", "agentic", "agents", "claude", "copilot", "mai", "mcp", "model context protocol"]
source: docs/RAG/collect-261001-ia-llm/serveur-mcp-pour-claude-code-tutoriel-2026.md
source_anchor: ""
source_lines: [1, 24]
sha256: a4cc8baebd9281774e228abcd34ab366761bbe5cf5ba6d42a19c13ca3c1d055e
---

# La sortie doit afficher Python 3.10.x ou une version supérieure

Le protocole MCP (Model Context Protocol) est devenu, en quelques mois à peine, le langage commun qui permet à Claude Code, Cursor, GitHub Copilot et une bonne partie des outils de codage IA de se connecter aux mêmes sources de données externes : dépôts Git, bases de données, outils de tickets, documentation interne. Lancé par Anthropic comme standard ouvert le 25 novembre 2024, il revendique aujourd’hui plus de 10 000 serveurs publics actifs selon l’annonce d’Anthropic sur sa donation du protocole à l’Agentic AI Foundation — un chiffre déjà dépassé sur le terrain, puisque le registre communautaire PulseMCP en recensait 22 311 en juillet 2026, contre environ 14 000 à peine deux mois plus tôt en mai, selon le suivi publié par ShareuHack — et les kits de développement officiels cumulent désormais plus de 97 millions de téléchargements au total sur l’ensemble des plateformes, toujours d’après ShareuHack, confirmant la tendance déjà repérée début 2026 par digitalapplied.com.

Ce tutoriel détaille la construction d’un serveur MCP fonctionnel, de zéro jusqu’à sa connexion effective à Claude Code : prérequis techniques, installation du kit de développement officiel, écriture des outils et de la ressource exposés par le serveur, tests avec l’Inspecteur MCP, déclaration dans la configuration de l’agent, pièges à éviter, dépannage et astuces avancées pour évoluer vers un serveur distant sécurisé. Treize étapes au total, pour un temps de mise en place d’environ 50 minutes.

Ce guide s’adresse à un développeur qui a déjà pris en main Claude Code, Cursor ou un autre client compatible MCP et qui veut désormais lui connecter ses propres outils, plutôt qu’à quelqu’un qui découvre ces agents pour la première fois. Chaque étape reste valable sur Windows, macOS et Linux, les variantes de commande étant précisées quand elles diffèrent d’un système à l’autre. Les exemples de sortie terminal cités restent illustratifs : le comportement exact peut varier légèrement selon la version du kit de développement installée au moment de la lecture.

## Qu’est-ce que le Model Context Protocol (MCP) ?

MCP est un standard ouvert qui définit comment un modèle de langage, via un client comme Claude Code, peut découvrir et appeler des outils externes, lire des ressources et récupérer des invites prédéfinies, sans que chaque éditeur d’outil ait à développer une intégration sur mesure pour chaque agent IA. Avant MCP, connecter un agent à une base de données ou à un système de tickets supposait souvent un connecteur propriétaire par paire outil/agent. Le protocole standardise cette relation une fois pour toutes : un serveur MCP écrit pour exposer un dépôt Git fonctionne aussi bien avec Claude Code qu’avec un autre client compatible, sans modification.

Concrètement, la différence se voit surtout au quotidien. Avant MCP, un développeur qui voulait que son agent IA tienne compte d’un ticket de support devait ouvrir l’outil de tickets dans un onglet séparé, copier le texte pertinent, puis le coller dans la conversation avec l’agent, en répétant l’opération à chaque nouvelle information nécessaire. Avec un serveur MCP connecté à cet outil de tickets, la même information devient accessible directement depuis la session de l’agent, sur simple demande en langage naturel, sans jamais quitter le terminal ou l’éditeur. Le principe rappelle celui d’un port universel pour les échanges de données entre un modèle et le monde extérieur. La page Wikipédia consacrée au Model Context Protocol retrace la genèse du projet chez Anthropic fin 2024 et son adoption rapide par le reste de l’industrie l’année suivante. Un serveur MCP peut exposer trois types d’éléments : des outils (des fonctions que l’agent peut appeler, comme lister des commits Git), des ressources (des données consultables, comme le contenu d’un fichier ou l’état d’un dépôt) et des invites préconçues (des modèles de requête réutilisables). Le tutoriel qui suit construit un serveur combinant les deux premiers.

### JSON-RPC, stdio et HTTP/SSE : l’architecture en bref

Techniquement, MCP repose sur JSON-RPC 2.0, un format d’échange de messages compact et bien établi, qui sert de base à la communication entre le client (l’agent) et le serveur. Trois transports coexistent désormais, stdio et HTTP restant le duo dominant selon la cartographie de l’écosystème publiée le 16 mai 2026 par Konishi — le registre officiel MCP en dénombrait pour sa part 9 400 en juillet 2026 selon Praxena, un chiffre qui illustre à quel point les méthodes de comptage restent hétérogènes d’une source à l’autre. Le transport stdio fait tourner le serveur comme un simple processus local, l’agent lui parlant via l’entrée et la sortie standard : c’est le mode le plus simple, celui que ce tutoriel utilise de bout en bout. Le transport HTTP, historiquement couplé à des évènements envoyés par le serveur (SSE), permet en revanche d’exposer un serveur MCP à distance, accessible par plusieurs postes ou plusieurs membres d’une équipe à la fois, au prix d’une configuration réseau et d’une authentification plus poussées, abordées plus loin dans la section consacrée aux astuces avancées. La spécification 2026-07-28, publiée fin juillet 2026 par l’Agentic AI Foundation en remplacement de la version 2025-11-25 selon Ottho, impose désormais le transport Streamable HTTP et déprécie officiellement l’ancien couple HTTP+SSE ; ce nouveau transport HTTP devient entièrement sans état (stateless), avec des capacités négociées requête par requête plutôt qu’à l’ouverture d’une session persistante, un changement d’architecture détaillé début août 2026 dans la newsletter de Victor Dibia et déjà commenté par Cloudflare.

Le choix du langage de développement reste ouvert. Les kits officiels couvrent désormais Python et TypeScript en premier lieu, mais aussi C#, publié par Microsoft en partenariat avec Anthropic — un engagement que Microsoft prolonge côté plateforme avec son propre Management MCP Server, passé en aperçu public dès mars 2026 avant d’atteindre la disponibilité générale en avril 2026 dans le cadre de son plan Wave 1 —, et Go, distribué par Block au sein de son agent Goose — les quatre kits reconnus comme SDK « Tier 1 » par le projet. Le 28 juillet 2026, le blog officiel de la spécification MCP a fait basculer ces quatre kits vers la version 2026-07-28 en même temps, gardant ainsi les implémentations Python, TypeScript, C# et Go strictement synchronisées à chaque évolution du protocole. Cette diversité explique en partie pourquoi le protocole s’est répandu aussi vite hors de l’écosystème Anthropic : une équipe déjà organisée autour d’un langage donné n’a pas besoin d’en apprendre un nouveau pour exposer ses propres outils à un agent IA.

## Pourquoi connecter Claude Code à des serveurs MCP en 2026 ?

Un agent de codage comme Claude Code raisonne déjà bien sur le code source qu’il a sous les yeux. Sa limite tient plutôt à tout ce qui vit en dehors du dépôt : un ticket dans un outil de suivi, une conversation dans un canal d’équipe, une table dans une base de données de production. Sans MCP, faire remonter ce contexte suppose des copier-coller manuels entre plusieurs fenêtres. Avec un serveur MCP correctement déclaré, l’agent peut interroger ces systèmes directement pendant une session, ce qui change concrètement la nature des tâches qu’on peut lui confier.

L’adoption du protocole depuis son lancement donne une idée de la vitesse à laquelle ce changement s’est installé dans l’industrie.

