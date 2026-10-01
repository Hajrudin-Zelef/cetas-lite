---
id: collect-261001-ia-llm/ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026-6
title: ".github/copilot-instructions.md"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "Mistral", "OpenAI", "xAI"]
dates: []
keywords: ["copilot", "agent", "arr", "gemini", "grok", "grok 4", "mcp", "mistral", "model context protocol", "open source", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026.md
source_anchor: ""
source_lines: [372, 422]
sha256: 47636436f8ed07057b65a839c5b722a1eac2573bbf2ed860066d38bb95226ca5
---

# .github/copilot-instructions.md

## Sécurité, RGPD et résidence des données en Europe

Pour une équipe qui code en France ou ailleurs en Europe, la question de la conformité se pose dès qu’un agent IA touche à du code source sensible ou à des données de production. Les paliers Business et Enterprise de GitHub Copilot ajoutent des couches de contrôle absentes des offres individuelles : authentification unique (SSO), journaux d’audit détaillés sur chaque action de l’agent, et des politiques de gestion centralisées au niveau de l’organisation. GitHub a par ailleurs annoncé en juin 2026 que les clients Enterprise pourraient activer un Autonomous Agent Mode à partir de juillet 2026, une capacité supplémentaire à évaluer avec la même rigueur de conformité avant tout déploiement sur du code sensible.

Sur le plan du RGPD, le point de vigilance principal reste le même qu’avec n’importe quel outil cloud : vérifiez précisément ce que votre contrat Enterprise garantit en matière de traitement et de conservation des données envoyées au modèle, avant de connecter l’agent à un dépôt contenant des données personnelles ou des secrets de production. Un fichier d’instructions comme celui présenté à l’étape 7 peut d’ailleurs servir à interdire explicitement à l’agent de manipuler certains dossiers sensibles, une pratique simple qui réduit une bonne partie du risque.

Si votre organisation gère déjà l’authentification unique pour d’autres outils internes, notre tutoriel Keycloak détaille la mise en place d’un SSO open source, une brique souvent utilisée en complément des contrôles d’accès proposés par les offres Copilot Business et Enterprise. Pour la partie scan de vulnérabilités du code généré, notre tutoriel Trivy montre comment vérifier automatiquement les dépendances et les images ajoutées à un projet, un réflexe utile quand une bonne partie du code provient désormais d’un agent.

Autre réglage à vérifier avant de généraliser l’outil à toute une équipe : les paramètres de rétention et d’utilisation des données pour l’entraînement des modèles, disponibles dans les réglages d’organisation sur GitHub. Sur les paliers Business et Enterprise, ces réglages permettent d’exclure le code de votre organisation de tout usage au-delà de la génération de réponses, un point que les équipes juridiques et sécurité demandent presque systématiquement avant validation.

## Foire aux questions

### Quelle est la différence entre l’Agent Mode et le mode Edit classique ?

Le mode Edit modifie un ou plusieurs fichiers selon vos instructions, puis s’arrête. L’Agent Mode va plus loin : il planifie la tâche, modifie les fichiers nécessaires, exécute des commandes pour vérifier le résultat, et corrige automatiquement ce qui ne fonctionne pas, sans attendre une nouvelle instruction à chaque étape.

### GitHub Copilot Agent Mode fonctionne-t-il en français ?

Oui. L’interface et les modèles sous-jacents comprennent le français aussi bien que l’anglais pour la formulation des prompts. Le code généré, lui, reste écrit dans la langue de programmation demandée, avec des commentaires que vous pouvez explicitement demander en français dans votre fichier d’instructions.

### Quel est le prix minimum pour utiliser l’Agent Mode ?

Le palier Pro à 10 $/mois est le premier à débloquer l’Agent Mode sans restriction sur les complétions de code. Le palier gratuit donne un accès limité, insuffisant pour un usage agentique régulier.

### L’agent peut-il supprimer ou casser mon code ?

Techniquement oui, s’il propose une commande destructrice que vous approuvez sans la lire. C’est pour cette raison que chaque modification de fichier passe par une validation Keep/Undo, et que chaque commande terminal attend votre confirmation par défaut. Le risque vient presque toujours d’une validation trop rapide, pas de l’agent lui-même.

### L’Agent Mode fonctionne-t-il avec JetBrains ou seulement avec VS Code ?

GitHub Copilot propose des extensions pour VS Code, les IDE JetBrains et Neovim. Les fonctionnalités les plus récentes, dont certaines options de l’Agent Mode, apparaissent généralement d’abord sur VS Code avant d’être portées sur les autres environnements. La documentation officielle de GitHub Copilot tient à jour la liste des fonctionnalités disponibles par IDE.

### Qu’est-ce qu’un serveur MCP, et dois-je en installer un ?

Un serveur MCP (Model Context Protocol) donne à l’agent accès à des outils externes, comme une base de données ou une API interne. Ce n’est pas obligatoire pour débuter : l’Agent Mode fonctionne très bien sans, en s’appuyant sur le système de fichiers local et le terminal. Un serveur MCP devient utile quand vous voulez connecter l’agent à des systèmes en dehors du projet ouvert dans VS Code.

### Comment savoir combien de crédits il me reste ?

Le tableau de bord de facturation de votre compte GitHub affiche la consommation de crédits en cours, détaillée par modèle. Il est accessible depuis les paramètres de facturation, en dehors de VS Code.

### L’Agent Mode remplace-t-il un développeur junior ?

Non, il change la nature du travail plutôt que de le remplacer. L’agent exécute vite des tâches bien définies, mais reste dépendant de la qualité du prompt et d’une relecture humaine à chaque étape. La compétence qui gagne en valeur n’est plus de taper du code ligne par ligne, mais de savoir décomposer un problème, écrire un prompt précis et juger si le résultat produit est correct.

### Contenus similaires

- Cursor AI : Tutoriel en 13 Étapes, 45 Min [2026] : l’installation et la prise en main de l’éditeur concurrent basé sur un fork de VS Code.
- Keycloak : SSO Open Source en 14 Étapes, 60 Min [2026] : pour sécuriser l’accès à vos outils internes en complément des offres Copilot Business et Enterprise.
- Trivy : Scanner de Failles en 13 Étapes, 45 Min [2026] : analyser automatiquement les dépendances et images d’un projet, y compris le code généré par un agent.
- Opus 4.8 vs GPT-5.5 vs Mistral Large 3 : prix x10 [2026] : comparatif des modèles disponibles dans le sélecteur de modèle de Copilot Chat.
- Grok 4.5 vs Opus 4.8 vs Gemini 3.1 Pro : 0 Accès UE [2026] : un autre comparatif de modèles utile pour choisir le bon moteur selon la tâche.
