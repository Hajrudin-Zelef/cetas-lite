---
id: collect-261001-ia-llm/ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026-5
title: "./scripts/validate-readonly-query.sh"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "JFrog"]
dates: []
keywords: ["agent", "agents", "claude", "mcp", "memory", "model context protocol"]
source: docs/RAG/collect-261001-ia-llm/sous-agents-claude-code-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [380, 426]
sha256: 0ad7200ca53d1bbb82567ec9a3d2a4d2ea49a4d626494161e013c517cfd8d624
---

# ./scripts/validate-readonly-query.sh

## Les pièges les plus courants (et comment les éviter)

**Oublier qu’un nouveau dossier agents/ demande un redémarrage.** Claude Code surveille les modifications de fichiers dans un dossier `agents` déjà existant, mais ne détecte pas la création du dossier lui-même pendant qu’une session tourne. Si votre sous-agent flambant neuf reste invisible, redémarrez avant de chercher plus loin.

**Donner bypassPermissions à un sous-agent par confort.** Ce mode saute les invites de permission, y compris pour les écritures dans `.git`, `.claude`, `.vscode` ou `.husky`. Une définition de sous-agent partagée dans un dépôt public avec ce réglage devient un vecteur d’exécution silencieuse pour quiconque clone le projet.

**Dupliquer un nom de sous-agent sans le savoir.** Deux fichiers portant le même `name` dans la même arborescence `.claude/agents/` ne provoquent pas d’erreur visible : Claude Code en charge un seul, selon l’ordre de lecture du système de fichiers. Le résultat est un comportement qui semble aléatoire d’une session à l’autre.

**Confondre la restriction tools et le contrôle réellement nécessaire.** Retirer Bash de la liste `tools` empêche toute commande shell, y compris les requêtes en lecture seule légitimes. Quand seule une sous-catégorie de commandes doit être bloquée, comme les écritures SQL, un hook `PreToolUse` fait le travail que `tools` ne peut pas faire seul.

**Ignorer que les sous-agents de plugin n’acceptent pas hooks, mcpServers ni permissionMode.** Ces trois champs sont silencieusement ignorés pour un sous-agent distribué par un plugin, pour des raisons de sécurité. Si vous en avez besoin, copiez le fichier dans `.claude/agents/` ou `~/.claude/agents/`.

**Sous-estimer la consommation de contexte des sous-agents parallèles.** Faire tourner cinq sous-agents en parallèle qui renvoient chacun un résultat détaillé peut remplir la conversation principale presque aussi vite qu’un seul long échange sans délégation. Pour un travail qui a vraiment besoin de parallélisme soutenu, les équipes d’agents donnent à chaque travailleur son propre contexte indépendant : lancées en aperçu de recherche en février 2026, après la disponibilité générale des sous-agents en 2025, selon HatchWorks, elles ciblent justement ce cas où plusieurs sous-agents parallèles saturent la conversation principale.

**Répéter dans le prompt des règles déjà couvertes par CLAUDE.md.** Un sous-agent standard charge déjà toute la hiérarchie de vos fichiers CLAUDE.md au démarrage, donc dupliquer ces règles dans sa description ou son prompt système n’ajoute rien, sauf pour Explore et Plan, qui sautent volontairement CLAUDE.md pour rester rapides. Si une consigne doit absolument atteindre l’un de ces deux agents, elle doit être répétée dans le prompt de délégation que vous donnez à Claude, pas dans le fichier du sous-agent.

## Dépannage : les problèmes fréquents et leurs solutions

Voici les huit situations les plus fréquemment rencontrées en configurant des sous-agents Claude Code, avec la cause exacte et la correction documentée par Anthropic.

| Symptôme | Cause probable | Solution | 
|---|---|---|
| Le sous-agent n’apparaît pas après sa création | Le dossier agents/ vient d’être créé pendant que la session tournait | Redémarrez Claude Code pour que le nouveau dossier soit détecté | 
| « Agent would be spawned with zero tools » | Aucune entrée de la liste tools ne correspond à un outil réel | Vérifiez l’orthographe exacte de chaque nom d’outil dans le frontmatter | 
| « Subagent spawn limit reached » | Plus de 200 sous-agents ont été générés dans la session (limite par défaut) | Lancez /clear pour réinitialiser le compteur, ou augmentez CLAUDE_CODE_MAX_SUBAGENTS_PER_SESSION | 
| « Concurrent subagent limit reached » | 20 sous-agents tournent déjà en même temps (limite par défaut) | Attendez qu’un sous-agent se termine, ou augmentez CLAUDE_CODE_MAX_CONCURRENT_SUBAGENTS | 
| Deux sous-agents semblent entrer en conflit | Le même name est déclaré deux fois dans la même arborescence agents/ | Lancez /doctor pour repérer les doublons et renommer l’un des deux fichiers | 
| Le sous-agent tourne sur le mauvais modèle | CLAUDE_CODE_SUBAGENT_MODEL est définie et prend le dessus sur le frontmatter | Vérifiez la variable d’environnement en premier, elle a la priorité la plus haute | 
| La mémoire du sous-agent ne persiste pas | La mémoire automatique est désactivée globalement (autoMemoryEnabled ou la variable dédiée) | Réactivez la mémoire automatique, le champ memory dépend entièrement d’elle | 
| Un sous-agent repris perd son modèle personnalisé | Comportement des versions antérieures à 2.1.211 lors d’une reprise | Mettez à jour vers 2.1.211 ou plus récent, la valeur par invocation est alors conservée à la reprise | 

Pour retrouver l’identifiant exact d’un sous-agent et consulter son historique complet, les transcriptions se trouvent dans `~/.claude/projects/{project}/{sessionId}/subagents/`, chacune stockée sous la forme `agent-{agentId}.jsonl`. Ces fichiers persistent indépendamment de la conversation principale et sont nettoyés automatiquement après 30 jours par défaut.

## Astuces avancées : sécurité, coûts et bonnes pratiques

Anthropic recommande quatre principes pour concevoir des sous-agents qui restent utiles sur la durée : des sous-agents focalisés sur une seule tâche précise, des descriptions détaillées puisque Claude s’appuie dessus pour décider quand déléguer, un accès aux outils limité au strict nécessaire pour la sécurité et la concentration, et des sous-agents de projet versionnés pour que toute l’équipe les améliore collectivement.

### La sécurité des serveurs MCP à l’intérieur d’un sous-agent

Un sous-agent qui embarque un serveur MCP via `mcpServers` hérite du même risque que n’importe quelle connexion MCP. En juillet 2025, une vulnérabilité critique référencée CVE-2025-6514 a touché le paquet npm `mcp-remote`, utilisé pour relier des clients MCP locaux à des serveurs distants. Avec un score CVSS de 9,6, la faille permettait une injection de commandes système à distance quand le client se connectait à un serveur malveillant qui renvoyait une URL `authorization_endpoint` piégée, comme l’a détaillé The Hacker News en citant l’équipe de recherche JFrog à l’origine de la découverte. Le paquet, qui comptait plus de 437 000 téléchargements, a été corrigé en version 0.1.16. La leçon pour un sous-agent : ne connectez un serveur MCP distant que si sa provenance est fiable, et gardez vos dépendances à jour.

Plus largement, une analyse de sécurité 2025 relayée par Wikipedia sur le Model Context Protocol pointe des risques structurels dans l’écosystème MCP, dont l’injection de prompt et les « outils empoisonnés » capables d’exfiltrer des données via un outil connecté. Un sous-agent en lecture seule, sans Write ni Edit, réduit mécaniquement la surface d’attaque même si l’un de ses outils MCP se révèle compromis.

### Isolation Git pour les sous-agents qui modifient du code

Par défaut, un sous-agent démarre dans le répertoire de travail de la conversation principale, et une commande `cd` ne persiste pas d’un appel d’outil à l’autre. Pour donner à un sous-agent une copie isolée du dépôt plutôt qu’un accès direct à votre copie de travail, ajoutez `isolation: worktree` dans son frontmatter. Le worktree est nettoyé automatiquement si le sous-agent ne modifie rien, ce qui le rend pratique pour des sous-agents qui testent des correctifs en parallèle sans risquer d’écraser votre travail en cours.

