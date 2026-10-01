---
id: collect-261001-ia-llm/ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026-6
title: "claude-skills-tutoriel-en-13-etapes-70-min-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["claude", "agents", "arr", "distribution", "mcp", "sonnet 5"]
source: docs/RAG/collect-261001-ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026.md
source_anchor: ""
source_lines: [331, 369]
sha256: eabc3b1a38ee50485a490cb149850217b449acb2f064ebd8a74f9c760aed2c2a
---

# claude-skills-tutoriel-en-13-etapes-70-min-2026

Enfin, pensez à la commande `/run-skill-generator`, décrite dans la documentation de Claude Code sur les Skills, une Skill embarquée dans Claude Code qui observe comment lancer votre application depuis un environnement propre puis enregistre la recette (commandes d’installation, variables d’environnement, script de lancement) comme Skill de projet réutilisable dans `.claude/skills/run-<nom>/`. Toute la logique de découverte manuelle que vous referiez sinon à chaque nouvelle session se transforme ainsi en Skill persistante, partagée avec le reste de l’équipe dès qu’elle est committée.

## Foire aux questions

**Les Claude Skills sont-elles gratuites ?**

Créer et utiliser une Skill personnalisée dans Claude Code ne coûte rien au-delà de votre abonnement Claude existant. Sur claude.ai, le téléversement de Skills personnalisées nécessite un plan Pro, Max, Team ou Enterprise avec l’exécution de code activée. Les Skills pré-construites (pptx, xlsx, docx, pdf) sont incluses sur claude.ai et facturées à l’usage normal des tokens sur l’API, où le tarif standard de Claude Sonnet 5 est fixé depuis le 10 août 2026 à 2 $ par million de tokens en entrée et 10 $ par million de tokens en sortie, d’après les notes de version d’Anthropic.

**Claude Skills fonctionne-t-il avec d’autres assistants IA que Claude ?**

Le format a été publié comme standard ouvert le 18 décembre 2025 sur agentskills.io, précisément pour permettre cette portabilité. Plusieurs rapports publiés fin 2025 et courant 2026 indiquent qu’OpenAI a annoncé une prise en charge du format dans Codex. Le niveau de compatibilité réel dépend cependant de l’implémentation de chaque outil tiers.

**Quelle est la différence entre une Skill et un plugin Claude Code ?**

Une Skill est une unité de savoir procédural : un dossier avec un SKILL.md. Un plugin est un paquet de distribution qui peut regrouper plusieurs Skills avec des agents, des hooks et des serveurs MCP. Une Skill peut exister seule ; un plugin est une façon d’en empaqueter et d’en partager plusieurs ensemble.

**Où stocker mes Skills pour qu’elles soient partagées avec toute l’équipe ?**

Committez le dossier dans `.claude/skills/` à la racine du dépôt Git. Chaque membre de l’équipe qui clone le projet et lance Claude Code y a accès automatiquement, sans étape d’installation manuelle.

**Les Skills sont-elles sûres à utiliser en production ?**

Une Skill que vous avez écrite vous-même, testée et versionnée dans Git ne pose pas plus de risque qu’un script interne classique. Le risque vient des Skills récupérées de sources tierces non vérifiées : Anthropic recommande explicitement de les traiter comme n’importe quel logiciel externe, avec audit du code avant toute utilisation en contexte sensible.

**Combien de Skills puis-je installer sans ralentir Claude ?**

Grâce à la divulgation progressive, chaque Skill installée ne coûte qu’environ 100 tokens au repos (son nom et sa description). Le nombre de Skills qui reste confortable dépend surtout de la lisibilité de vos descriptions : au-delà de plusieurs dizaines, il devient plus difficile d’écrire des descriptions suffisamment distinctes pour que Claude choisisse la bonne à chaque fois.

**Puis-je utiliser une Skill écrite pour Claude Code dans l’API ?**

Pas directement si elle utilise des champs propres à Claude Code comme `context`, `argument-hint` ou l’injection dynamique `` !`commande` ``. Pour la rendre compatible avec l’API ou claude.ai, limitez le frontmatter aux six champs portables du standard et retirez toute syntaxe spécifique à Claude Code du corps du fichier.

**Comment mettre à jour une Skill déjà installée ?**

Dans Claude Code, modifiez directement le fichier SKILL.md : la détection à chaud applique le changement dans la session en cours, sans redémarrage. Sur claude.ai ou via l’API, il faut retéléverser une nouvelle archive pour remplacer la version existante.

## Ce qu’il faut retenir

Claude Skills résout un problème précis : arrêter de répéter les mêmes instructions à chaque conversation, sans pour autant les charger en permanence comme le ferait un CLAUDE.md surdimensionné. La divulgation progressive à trois niveaux, le frontmatter YAML normalisé et le statut de standard ouvert publié en décembre 2025 en font une brique durable de l’écosystème Claude plutôt qu’une fonctionnalité isolée. Commencez petit — une seule Skill personnelle, testée localement — avant de la faire évoluer vers un partage d’équipe via Git ou un plugin.
