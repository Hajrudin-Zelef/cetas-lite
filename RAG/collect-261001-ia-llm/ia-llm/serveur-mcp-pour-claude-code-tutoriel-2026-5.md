---
id: collect-261001-ia-llm/ia-llm/serveur-mcp-pour-claude-code-tutoriel-2026-5
title: "La sortie doit afficher Python 3.10.x ou une version supérieure"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Hugging Face", "Microsoft", "OpenAI", "Stripe"]
dates: ["2026-07-28"]
keywords: ["agent", "aws", "claude", "copilot", "mcp", "open source"]
source: docs/RAG/collect-261001-ia-llm/serveur-mcp-pour-claude-code-tutoriel-2026.md
source_anchor: ""
source_lines: [284, 343]
sha256: 66f64c2adb250b3f3dfad2cdefc60ddc4c335a9a1cb5b2241507b6d60231299c
---

# La sortie doit afficher Python 3.10.x ou une version supérieure

Une fois le serveur local `git-journal` maîtrisé, plusieurs pistes permettent d’aller plus loin. La première consiste à basculer vers un transport HTTP/SSE plutôt que stdio, ce qui permet d’héberger le serveur sur une machine partagée et de le rendre accessible à toute une équipe plutôt qu’à un seul poste. Ce choix impose en retour une couche d’authentification, généralement fondée sur des jetons, pour éviter qu’un serveur exposé sur le réseau ne devienne un point d’accès non contrôlé aux systèmes qu’il expose.

Le principe de moindre privilège mérite d’être appliqué systématiquement : un outil ne devrait exposer que le strict nécessaire, jamais une commande shell arbitraire. Ajouter une journalisation des appels d’outils, même minimale, facilite également l’audit a posteriori si un comportement inattendu doit être analysé, un réflexe d’autant plus utile à mesure qu’un serveur MCP gagne en périmètre et en nombre d’utilisateurs.

Sur un serveur partagé par plusieurs développeurs, pensez aussi à borner le débit des appels par utilisateur, en particulier si un outil déclenche une requête coûteuse vers une base de données de production ou vers une API tierce facturée à l’appel. Une limite simple, par exemple un nombre maximal de requêtes par minute et par jeton d’authentification, évite qu’une session mal configurée ou qu’une boucle involontaire côté agent ne sature le serveur ou ne génère une facture inattendue. Ce point est souvent négligé lors d’un premier déploiement en HTTP/SSE, alors qu’il se règle en quelques lignes dès la conception du serveur.

### Compatibilité avec Cursor, VS Code et les autres clients MCP

L’un des intérêts pratiques de MCP tient à sa portabilité : le serveur `git-journal` construit dans ce tutoriel fonctionne sans modification avec n’importe quel client compatible, pas seulement Claude Code. Cursor prend en charge le même mécanisme de configuration par serveur, tout comme certaines extensions liées à GitHub Copilot en mode agent. Un seul serveur MCP écrit une fois peut donc être partagé entre plusieurs outils utilisés au sein d’une même équipe, plutôt que de dupliquer la même logique d’intégration pour chaque agent.

Il existe par ailleurs de nombreux serveurs MCP déjà écrits et maintenus, qu’il n’est pas nécessaire de redévelopper soi-même pour des besoins courants. Le mouvement s’est nettement accéléré courant 2026 : le 23 juillet 2026, GitHub a annoncé dans son changelog que son propre serveur MCP prenait déjà en charge la spécification 2026-07-28, avant même le déploiement officiel du 28 juillet, et son serveur officiel github/github-mcp-server est désormais accessible en HTTP distant avec authentification OAuth 2.1 depuis avril 2026. Firecrawl recensait dès le 24 février 2026 huit services majeurs proposant un serveur MCP distant prêt à l’emploi — GitHub, Vercel, Linear, Notion, Supabase, Stripe, Figma et Hugging Face —, une liste à laquelle Snowflake a depuis ajouté son propre serveur MCP géré, entré en avant-première le 2 octobre 2025 pour donner un accès sécurisé aux comptes Snowflake directement depuis une session d’agent. AWS avait ouvert la voie bien plus tôt, le 14 juillet 2025, avec son propre serveur MCP « AWS Price List », qui donne accès en temps réel aux tarifs de centaines de services AWS directement depuis une session d’agent, avant de compléter l’offre dès le 16 juillet 2025 avec AWS Knowledge MCP Server, lancé en préversion publique gratuite mais soumise à des limites de quotas.

| Serveur MCP | Éditeur / mainteneur | Usage principal | 
|---|---|---|
| Filesystem | Anthropic (référence officielle) | Lire et écrire des fichiers locaux de façon contrôlée | 
| GitHub | Communauté open source | Lire les tickets, pull requests et fichiers d’un dépôt distant | 
| Slack | Communauté open source | Lire et publier des messages depuis une session agent | 
| Postgres | Communauté open source | Interroger un schéma de base de données en lecture | 
| Stripe | Stripe | Consulter des données de paiement et de facturation | 
| Notion | Notion | Lire et modifier des pages et bases de données Notion | 

La liste officielle des serveurs MCP de référence, maintenue sur GitHub, reste le point de départ le plus fiable avant d’envisager d’écrire un serveur sur mesure : mieux vaut vérifier qu’un serveur existant ne couvre pas déjà le besoin avant de partir de zéro.

## Foire aux questions sur les serveurs MCP

**Qu’est-ce qu’un serveur MCP, en une phrase ?**

C’est un programme qui expose des outils, des ressources ou des invites à un agent IA compatible, comme Claude Code, selon un format d’échange standardisé plutôt qu’une intégration propriétaire.

**MCP est-il réservé à Claude Code ?**

Non. Bien qu’Anthropic ait créé le protocole, il a été adopté par OpenAI, Google DeepMind et Microsoft, et fonctionne aussi avec Cursor, VS Code et d’autres clients compatibles, comme détaillé dans la section sur la compatibilité plus haut.

**Faut-il savoir coder en Python pour créer un serveur MCP ?**

Le kit de développement officiel existe aussi en TypeScript. Ce tutoriel utilise Python car son kit `FastMCP` réduit le code nécessaire au minimum, mais le choix du langage reste libre.

**Un serveur MCP personnalisé est-il sûr à connecter à mon code source ?**

Oui, à condition de respecter le principe de moindre privilège détaillé plus haut : un outil ne devrait exposer que des actions précises et contrôlées, jamais une exécution de commande arbitraire non filtrée. Sur un dépôt sensible, testez d’abord le serveur sur un projet non critique, exactement comme on le ferait avant d’accorder à Claude Code lui-même des permissions étendues sur un dépôt de production.

**Peut-on utiliser des serveurs MCP tout faits sans en écrire un soi-même ?**

Absolument, et c’est même recommandé pour des besoins courants comme la lecture de fichiers, l’accès à GitHub ou à une base Postgres : la table de serveurs prêts à l’emploi plus haut couvre les cas les plus fréquents.

**Quelle différence entre un serveur MCP local (stdio) et un serveur distant (HTTP) ?**

Un serveur en stdio tourne comme un processus local propre à un poste, simple à mettre en place mais limité à un seul utilisateur. Un serveur en HTTP/SSE tourne sur une machine accessible par le réseau et peut être partagé par une équipe entière, moyennant une authentification adaptée.

**MCP va-t-il remplacer les extensions d’éditeur classiques comme GitHub Copilot ?**

Non, les deux répondent à des besoins différents. Copilot reste centré sur l’auto-complétion dans l’éditeur, tandis que MCP standardise la façon dont un agent accède à des données et outils externes, quel que soit l’agent utilisé par ailleurs. Les deux approches se combinent d’ailleurs très bien : rien n’empêche une équipe d’utiliser Copilot au quotidien pour l’auto-complétion tout en connectant Claude Code à des serveurs MCP pour les tâches qui demandent un contexte plus large que le seul fichier ouvert.

**Où trouver d’autres serveurs MCP prêts à l’emploi ?**

Le dépôt GitHub officiel `modelcontextprotocol/servers` référence les serveurs de démonstration maintenus par Anthropic et par la communauté, un bon point de départ avant d’en écrire un sur mesure.

### Pour aller plus loin

D’autres tutoriels publiés sur tech-insider.org complètent ce guide sur les outils de codage IA :
