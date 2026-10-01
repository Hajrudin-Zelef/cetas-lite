---
id: collect-261001-general-networking/general-networking/google-antigravity-tutoriel-en-12-etapes-2026-1
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft"]
dates: []
keywords: ["agent", "agents", "arr", "claude", "copilot", "gemini", "mai"]
source: docs/RAG/collect-261001-general-networking/google-antigravity-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 35]
sha256: efbf0fe6cb0a33bf392579e1cab4d19641b5475e43bb86688437b2821dc2fe40
---

# macOS (Homebrew)

Depuis le 18 novembre 2025, **Google Antigravity** a changé la donne pour les développeurs francophones. Lancé le même jour que Gemini 3 en version 1.0 et proposé sans frais (0 $) pendant l’aperçu public – avec, selon Google, un quota Gemini 3.1 Pro qui se rafraîchit toutes les 5 heures –, cet **IDE agentique** ne se contente pas de compléter votre code : il planifie, écrit, exécute, teste et vérifie des tâches entières pendant que vous supervisez. C’est le passage de l’autocomplétion à l’agent autonome, et l’outil a depuis bien évolué : une application de bureau et une CLI Antigravity 2.0 sont arrivées le 19 mai 2026 lors de Google I/O, la version stable est passée à la 2.1.1 le 22 juin 2026 selon Wikipédia, puis Google a enchaîné les versions 2.9.1, 2.10.0 et 2.11.0 en seulement sept jours entre le 20 et le 26 août 2026 – cette dernière ajoutant une UI générative pour du HTML inline. Dans ce tutoriel mis à jour en août 2026, vous allez installer Google Antigravity, le configurer et construire un projet complet – une application web de gestion de tâches – en 12 étapes, en environ 40 minutes.

Ce guide s’adresse aux développeurs de France et d’Europe qui veulent comprendre concrètement comment fonctionne un IDE « agent-first ». Nous couvrons les prérequis avec versions, le choix des modèles (Gemini 3 Pro, Claude Sonnet 4.6, GPT-OSS-120B), l’écriture d’un fichier `AGENTS.md`, la différence entre l’Éditeur et le Manager d’agents, la vérification dans le navigateur, plus une liste de pièges, un dépannage détaillé et une section tarifs. À la fin, vous disposerez d’un projet fonctionnel versionné avec Git.

## Qu’est-ce que Google Antigravity ?

**Google Antigravity** était à l’origine une plateforme de développement agentique construite sur un fork profondément modifié de Visual Studio Code. Concrètement, cela signifiait que vos extensions, thèmes et raccourcis VS Code fonctionnaient, tandis que l’expérience restait réorganisée autour d’agents IA autonomes. Mais depuis l’annonce d’Antigravity 2.0 le 19 mai 2026 lors de Google I/O, Google indique que cette base VS Code a été remplacée par une application de bureau autonome dédiée à l’orchestration des agents, tout en conservant la compatibilité avec vos extensions, thèmes et raccourcis existants. D’après le blog officiel Google pour les développeurs, l’outil reste « conçu pour vous faire opérer à un niveau supérieur, orienté tâche » plutôt que ligne par ligne.

La plateforme repose sur deux surfaces de travail. La première, l’**Éditeur** (Editor View), est une IDE classique dopée à l’IA, avec complétion par tabulation et commandes en ligne pour un flux synchrone – proche de ce que propose Cursor. La seconde, le **Manager** (Manager Surface, parfois appelée « Mission Control »), est une interface dédiée où vous lancez, orchestrez et observez plusieurs agents travaillant en parallèle et de manière asynchrone, sur différents espaces de travail. Depuis le lancement d’Antigravity 2.0 le 19 mai 2026 à l’occasion de Google I/O, ce Manager existe même en application de bureau autonome accompagnée d’une CLI dédiée, une première pour une surface entièrement consacrée à l’orchestration d’agents, selon le blog d’antigravity.google. C’est cette seconde surface qui distingue vraiment Antigravity de ses concurrents.

Les agents opèrent de bout en bout sur trois terrains : l’éditeur, le terminal et le navigateur. Ils écrivent du code, lancent l’application, testent des composants et itèrent sur l’interface sans intervention humaine constante. Plutôt que de vous noyer sous des journaux bruts, ils produisent des **Artifacts** : listes de tâches, plans d’implémentation, captures d’écran et même enregistrements du navigateur. Vous laissez vos commentaires directement sur ces Artifacts, et l’agent les intègre sans s’arrêter. Une base de connaissances permet en outre aux agents de sauvegarder du contexte et des extraits utiles pour améliorer les tâches suivantes.

Selon la page Wikipédia consacrée à Google Antigravity, la plateforme englobe une IDE, une CLI et un SDK pour construire vos propres agents ; TechCrunch confirme que le lancement d’Antigravity 2.0, le 19 mai 2026, a formalisé cette offre autour de trois briques – l’application de bureau, la CLI et le SDK – toutes dédiées à l’orchestration d’agents. Comme l’analyse le média spécialisé The New Stack, Antigravity a démarré comme un concurrent de Cursor avant de devenir une plateforme de développement à part entière, avec orchestration multi-agents. Elle s’inscrit dans la vague du vibe coding qui a redéfini le développement logiciel en 2026, où l’on décrit une intention et où l’agent produit le code. Antigravity pousse cette logique jusqu’à l’orchestration de plusieurs agents simultanés.

## Google Antigravity face à Cursor, Windsurf et Copilot

Avant de plonger dans l’installation, il faut situer **Google Antigravity** dans l’écosystème. Trois différences majeures ressortent : l’orchestration multi-agents asynchrone via le Manager, la gratuité pendant l’aperçu public, et l’intégration native à l’écosystème Google (AI Studio, Firebase, Android). Là où Cursor et Windsurf misent sur un flux à agent unique enrichi, Antigravity assume une posture d’agent autonome de premier niveau, avec une IDE reléguée au rang d’outil parmi d’autres.

| Critère | Google Antigravity | Cursor | Windsurf | GitHub Copilot | 
|---|---|---|---|---|
| Approche | Agent-first, multi-agents | Éditeur + agent | Éditeur + Cascade | Assistant intégré | 
| Orchestration parallèle | Oui (Manager) | Limitée | Limitée | Non | 
| Modèles principaux | Gemini 3, Claude 4.6, GPT-OSS | Multi-modèles | Multi-modèles | GPT, Claude, Gemini | 
| Contrôle du navigateur | Natif (sous-agent) | Partiel | Partiel | Non | 
| Base VS Code | Fork modifié | Fork modifié | Fork modifié | Extension | 
| Prix de départ | Gratuit (aperçu) | Payant après essai | Gratuit limité | Payant | 

Si vous hésitez encore entre ces outils, deux comparatifs approfondis existent sur le site : Windsurf vs Cursor 2026 et Claude Code vs Cursor 2026. Ils détaillent les scores SWE-bench et les fenêtres de contexte. Antigravity se positionne comme une troisième voie, plus orientée délégation. À noter : la plateforme a subi plusieurs réductions de quotas depuis son lancement, un point sur lequel nous reviendrons dans la section tarifs.

Un dernier repère : contrairement à GitHub Copilot, étroitement lié à la chaîne d’outils Microsoft, Antigravity mise sur l’écosystème Google et sur le modèle Gemini 3 Pro comme moteur par défaut. Le résultat est un outil plus autonome, mais aussi plus gourmand en quotas – d’où l’importance de bien le configurer, ce que nous faisons dès maintenant.

## Prérequis : matériel, système et versions

Dès son lancement public le 18 novembre 2025, Google Antigravity était disponible sur les trois grandes familles de systèmes d’exploitation – Windows, macOS et Linux –, selon le blog officiel Google Antigravity. Vérifiez que votre machine respecte les versions minimales officielles ci-dessous avant de télécharger quoi que ce soit. Pour le projet de démonstration, nous utiliserons également Python et un navigateur Chromium pour le sous-agent navigateur.

