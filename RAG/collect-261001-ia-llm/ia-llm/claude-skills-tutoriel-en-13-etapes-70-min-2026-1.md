---
id: collect-261001-ia-llm/ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026-1
title: "claude-skills-tutoriel-en-13-etapes-70-min-2026"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Microsoft", "OpenAI"]
dates: []
keywords: ["claude", "agent", "agents", "aws", "disclosure", "foundry", "mcp", "model context protocol"]
source: docs/RAG/collect-261001-ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026.md
source_anchor: ""
source_lines: [1, 42]
sha256: 6a18346f82d5efba7e198c1abf3ca4a429a1ab45b92d2e558e4472ba6297513a
---

# claude-skills-tutoriel-en-13-etapes-70-min-2026

Un développeur qui colle la même checklist de déploiement dans Claude pour la dixième fois de la semaine a un problème d’organisation, pas un problème d’IA. Anthropic a résolu ce problème avec **Claude Skills** (ou Agent Skills), un système de dossiers que Claude charge automatiquement quand il en a besoin, au lieu de tout garder en mémoire à chaque conversation. Introduites le 16 octobre 2025 et publiées comme standard ouvert le 18 décembre 2025 sur agentskills.io, les Skills fonctionnent désormais sur claude.ai, Claude Code, l’API et le Claude Agent SDK. Ce tutoriel montre comment en créer une de zéro, la tester, la sécuriser et la déployer en 13 étapes, avec un projet complet à la fin que vous pourrez copier tel quel.

## Qu’est-ce que Claude Skills ? Définition, lancement et standard ouvert

Une Claude Skill est un dossier contenant un fichier `SKILL.md`, accompagné en option de scripts et de ressources annexes. Dans son billet d’ingénierie de lancement, Anthropic les décrit comme des “dossiers organisés d’instructions, de scripts et de ressources que les agents peuvent découvrir et charger dynamiquement” pour mieux exécuter une tâche précise. L’analogie que l’entreprise utilise dans son billet d’ingénierie est parlante : une Skill ressemble au guide d’intégration qu’on écrirait pour un nouveau collègue, avec la procédure, les pièges connus et les bonnes pratiques du métier.

La mécanique qui rend tout ça utile s’appelle la divulgation progressive (*progressive disclosure*). Claude ne charge pas le contenu complet de chaque Skill installée à chaque tour de conversation, ce qui saturerait rapidement la fenêtre de contexte. Il charge trois niveaux de contenu à des moments différents : le nom et la description de la Skill sont toujours présents dans le prompt système (environ 100 tokens par Skill), le corps du `SKILL.md` ne se charge que lorsque la Skill se déclenche (moins de 5 000 tokens recommandés), et les fichiers annexes ou scripts ne se chargent que si Claude les référence explicitement. Ce découpage permet d’installer des dizaines de Skills sans pénaliser le contexte disponible pour le reste de la conversation.

Anthropic fournit quatre Skills prêtes à l’emploi pour la bureautique : PowerPoint (`pptx`), Excel (`xlsx`), Word (`docx`) et PDF (`pdf`), disponibles sur claude.ai, l’API, Claude Platform sur AWS et Microsoft Foundry. L’écosystème s’est étoffé rapidement depuis le lancement : Anthropic recensait déjà 17 Skills jugées prêtes pour la production en juin 2026 selon The Data Scientist, tandis que Collective Brain cataloguait de son côté 56 Skills validées réparties en 7 catégories à la même période. Le 18 décembre 2025, Anthropic a publié la spécification Agent Skills comme standard ouvert et portable entre outils IA. Plusieurs rapports parus fin 2025 et début 2026 indiquent qu’OpenAI a annoncé la prise en charge du format dans Codex peu après cette publication, ce qui en ferait l’un des rares formats d’extension IA à circuler entre éditeurs concurrents plutôt que de rester propriétaire.

## Claude Skills vs MCP vs Projects Claude : les différences clés

La confusion la plus fréquente concerne le rôle respectif des Skills et du Model Context Protocol (MCP), que nous avons déjà couvert en détail sur Tech Insider. Anthropic tranche clairement la question dans sa documentation : MCP fournit la connectivité sécurisée vers des outils et des données externes (une base Postgres, un dépôt GitHub, un espace Slack), tandis qu’une Skill fournit le savoir procédural pour utiliser correctement ces outils une fois connectés. Les deux systèmes sont complémentaires, pas concurrents : une Skill peut très bien contenir des instructions expliquant comment enchaîner plusieurs appels à un serveur MCP.

La différence avec les instructions personnalisées façon Projects ou un fichier CLAUDE.md est tout aussi nette. Un CLAUDE.md se charge intégralement au démarrage de chaque session, même si son contenu ne sert jamais dans la conversation en cours. Une Skill, elle, reste invisible (hormis son nom et sa description) jusqu’à ce que la tâche en cours corresponde réellement à son usage. C’est la raison pour laquelle la documentation de Claude Code recommande de transformer en Skill toute section d’un CLAUDE.md qui a fini par ressembler à une procédure de plusieurs étapes plutôt qu’à un simple fait sur le projet.

| Critère | Claude Skills | Serveur MCP | CLAUDE.md / Projects | 
|---|---|---|---|
| Rôle principal | Savoir procédural réutilisable | Connectivité sécurisée à des outils/données externes | Contexte permanent du projet | 
| Chargement | À la demande (divulgation progressive) | Outils déclarés au démarrage, exécutés à la demande | Intégral à chaque session | 
| Coût en contexte au repos | ~100 tokens (nom + description) | Dépend du nombre d’outils exposés | Taille totale du fichier | 
| Format | Dossier + SKILL.md | Serveur (process local ou distant) | Fichier Markdown unique | 
| Portabilité | Standard ouvert (agentskills.io) | Standard ouvert (Anthropic, 2024) | Propre à chaque outil | 
| Cas d’usage typique | Checklist de revue, format de rapport, convention d’équipe | Lire une base de données, interroger une API tierce | Stack technique, conventions de nommage | 

## Prérequis : versions, comptes et outils nécessaires

Avant de commencer, vérifiez que vous disposez de ce qui suit. La plupart des fonctionnalités avancées décrites plus loin (sous-agents, injection de contexte dynamique, champs `effort` et `background`) nécessitent une version relativement récente de Claude Code, alors autant partir sur une base à jour.

- **Claude Code** installé et fonctionnel (dernière version disponible via`npm install -g @anthropic-ai/claude-code` ou l’installeur natif) — vérifiez votre version avec`claude --version` ou la commande`/status` une fois dans une session.
- **Un compte Claude** sur un plan Pro, Max, Team ou Enterprise si vous comptez aussi utiliser des Skills personnalisées sur claude.ai (l’exécution de code doit être activée dans les paramètres).
- **Un accès à l’API Claude** (clé API valide) si vous voulez tester les Skills pré-construites (`pptx` ,`xlsx` ,`docx` ,`pdf` ) par ce canal.
- **Git** installé, pour l’exemple de Skill que nous allons construire (elle s’appuie sur`git diff` ).
- **Un terminal** et un éditeur de texte pour écrire du Markdown et du YAML.
- **Python 3** (optionnel) si vous voulez tester le script`package_skill.py` du dépôt officiel pour empaqueter une Skill destinée à claude.ai ou à l’API.

Aucune de ces briques n’est exotique : si vous avez déjà suivi notre tutoriel d’installation de Claude Code, vous avez déjà tout ce qu’il faut pour la partie locale de ce guide.

## Étape 1 : choisir l’emplacement de votre Skill

Première décision, et elle conditionne tout le reste : où stocker le dossier de la Skill. Claude Code reconnaît quatre emplacements, avec un ordre de priorité strict quand deux Skills portent le même nom : le niveau entreprise prime sur le personnel, qui prime lui-même sur le niveau projet. Le déploiement de Skills au niveau entreprise, qui permet à une organisation entière d’imposer ou de partager des Skills via ses paramètres managés, est disponible depuis le 18 décembre 2025, date à laquelle Anthropic l’a formalisé dans son guide officiel des Skills. Une Skill à n’importe lequel de ces niveaux peut aussi remplacer une Skill “bundled” (embarquée nativement dans Claude Code) qui porterait le même nom.

