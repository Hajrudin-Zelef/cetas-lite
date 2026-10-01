---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026-6
title: "AGENTS.md"
domain: ia-llm
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "agents", "copilot"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026.md
source_anchor: ""
source_lines: [339, 367]
sha256: b66177f6500c5e7dc4953bfae5febd47a909a5fdf998137f13d22e19aad4635d
---

# AGENTS.md

AGENTS.md est un fichier unique et simple, sans métadonnées, idéal pour démarrer. Les fichiers .mdc dans .cursor/rules permettent un ciblage précis par type de fichier grâce aux globs, et une activation conditionnelle via alwaysApply. Les deux mécanismes peuvent cohabiter dans un même projet.

**Cursor AI peut-il fonctionner hors ligne ?**

Non. L'éditeur lui-même s'ouvre sans connexion, mais toutes les fonctions IA (Tab, chat, Agent) nécessitent une connexion Internet pour interroger les modèles hébergés dans le cloud.

**Le code envoyé à Cursor AI sert-il à entraîner les modèles ?**

Cela dépend de vos réglages de confidentialité et de la formule choisie. Les formules Teams et Enterprise proposent un mode confidentialité renforcé et une rétention de données réduite. Vérifiez les paramètres de rétention avant de connecter un dépôt contenant des données sensibles ou soumises au RGPD.

**Peut-on utiliser sa propre clé API avec Cursor AI ?**

Oui, l'option est disponible dans Settings > Models. Elle permet de faire transiter les requêtes par votre propre contrat avec le fournisseur du modèle plutôt que de consommer les crédits inclus dans votre formule Cursor.

**Cursor AI remplace-t-il un développeur ?**

Non. L'outil accélère l'écriture, le refactoring et les tâches répétitives, mais les choix d'architecture, la relecture de sécurité et la validation métier restent une responsabilité humaine. Les pièges listés plus haut dans cet article montrent justement ce qui arrive quand cette relecture disparaît.

**Comment migrer de GitHub Copilot vers Cursor AI ?**

Installez Cursor AI, importez vos paramètres VS Code au premier lancement (étape 2 de ce tutoriel), puis reconstituez progressivement vos habitudes de prompt sous forme de règles .mdc ou AGENTS.md. La bascule technique prend quelques minutes, l'adaptation des réflexes d'équipe prend généralement une à deux semaines.

**Combien coûte Cursor AI pour une équipe de dix développeurs ?**

Sur la formule Teams à 40 $ par utilisateur et par mois, une équipe de dix développeurs paie 400 $ mensuels hors taxes, facturation centralisée incluse. Ce montant n'inclut pas les dépassements éventuels de crédits ni la revue de code Bugbot facturée à l'usage.

### Couverture connexe

Pour suivre l'actualité complète des outils de développement, retrouvez tous nos tutoriels dans la rubrique Logiciels.
