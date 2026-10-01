---
id: collect-261001-ia-llm/ia-llm/deepseek-harness-explique-cordis-et-architecture-par-plugins-3
title: "deepseek-harness-explique-cordis-et-architecture-par-plugins"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "OpenAI"]
dates: []
keywords: ["deepseek", "agent", "agents", "benchmark", "claude", "mcp", "tool calling"]
source: docs/RAG/collect-261001-ia-llm/deepseek-harness-explique-cordis-et-architecture-par-plugins.md
source_anchor: ""
source_lines: [189, 253]
sha256: fd2efd1ab6aadb6b1e83f8a582c01d7ac41c509381750ccdd37f9b0a213d87c2
---

# deepseek-harness-explique-cordis-et-architecture-par-plugins

Cela peut être superflu pour de simples appels modèle ou pour des équipes qui veulent un agent de code prêt à l’emploi sans toucher à ses entrailles. Remplacer davantage de pièces ne vaut l’effort que si ce contrôle résout un vrai problème.

## Limites de DeepSeek Harness : statut de préversion et risques de sécurité

Toute l’architecture ci-dessus importe peu sans une vision claire de ses limites actuelles.

### C’est toujours une préversion développeur

Le dépôt DeepSeek indique clairement que des changements cassants surviendront. C’est déjà arrivé : le renommage de Code en PTC s’est accompagné de changements d’API de session et du retrait d’une option SQLite facultative. Figez vos versions. Sauter cette étape en espérant une stabilité de configuration n’est pas une stratégie.

### Plus de contrôle, plus de complexité

Rendre davantage de couches remplaçables implique davantage à apprendre : dépendances de plugins, paramètres, différences entre fournisseurs, compatibilités de versions. C’est le compromis habituel entre confort et contrôle.

### DeepSeek Harness est-il local ?

DeepSeek Harness stocke par défaut le contenu des sessions, les traces d’outils et les paramètres en local, conformément à sa déclaration de traitement des données. Vous pouvez désactiver ses rapports anonymes sur les paramètres et les listes de projets.

Mais un fournisseur de modèle externe, un outil web, un serveur MCP ou un plugin peuvent toujours envoyer des données hors de votre machine selon leur propre politique. « Local-first » ne couvre pas tous les services que vous connectez.

### Exécuter des agents comporte des risques de sécurité

Un runtime capable d’éditer des fichiers, d’exécuter des commandes et de charger des plugins tiers peut causer de vrais dégâts. L’avis de sécurité de DeepSeek précise qu’aucun audit n’a été mené. Bacs à sable, approbations et contrôles d’autorisations réduisent le risque sans garantir l’isolation.

Exécuter le logiciel sur votre propre machine n’élimine pas ce risque. Utilisez des permissions limitées et un environnement jetable pour les tâches non fiables, et soyez vigilant avec les contenus susceptibles de contenir des instructions cachées.

## Pourquoi le comportement d’un agent dépend de plus que du modèle

Le comportement d’un agent dépend du runtime autant que du modèle. On revient à « Agent = Modèle + Harness », et cette séparation vaut pour les agents LLM au-delà de DeepSeek.

Ce qu’un modèle peut produire dépend de ses poids. Ce que fait un agent dépend aussi du contexte transmis au modèle, des actions autorisées et du degré de contrainte de l’exécution. Rien de cela ne réside dans les poids.

DeepSeek Harness met en lumière cette couche environnante en la découpant en composants nommés et remplaçables. Le mode Minimal illustre pourquoi cela dépasse DeepSeek : un score de benchmark reflète en partie le harness utilisé pour le test, pas seulement le modèle. Le harness ne rend pas un modèle plus « intelligent ». Il change le cadre dans lequel il opère.

## Conclusion

La phrase d’ouverture est à retenir : le modèle raisonne, mais le runtime décide à quoi ce raisonnement peut accéder et ce qu’il peut faire. DeepSeek Harness rend ce runtime modifiable, de l’adaptateur de modèle et des outils jusqu’au magasin de sessions et à la boucle d’agent.

Ce contrôle a un coût. Remplacer davantage de briques signifie assumer davantage la configuration, les évolutions de versions et les frontières de sécurité. Une préversion développeur avec accès shell n’est pas un outil à « installer et oublier ».

Mon avis est simple : utilisez DeepSeek Harness quand le runtime fait partie du travail. Si vous avez seulement besoin de modifications dans un dépôt, un agent de code prêt à l’emploi vous en demandera moins.

Notre tutoriel DeepSeek Harness couvre la configuration. Le guide sur les alternatives à Claude Code compare d’autres agents de code, tandis que Introduction to AI Agents revient sur les bases supposées connues ici.

## FAQ sur DeepSeek Harness

### DeepSeek Harness est-il la même chose qu’un modèle DeepSeek ?

**Non, le modèle et le runtime sont distincts. Harness n’inclut pas les poids d’un modèle et n’exécute pas l’inférence lui-même ; il envoie des requêtes à DeepSeek, Anthropic, OpenAI ou à un modèle local.**

### L’utilisation de DeepSeek Harness est-elle gratuite ?

**Le logiciel lui-même est gratuit et sous licence MIT. Ce qui ne l’est pas, c’est le fournisseur de modèle auquel vous vous connectez, puisque l’inférence est facturée séparément par l’opérateur du modèle, ainsi que tout coût d’infrastructure lié aux bacs à sable ou services externes que vous ajoutez par-dessus.**

### Que signifie réellement l’acronyme PTC ?

**Les notes de version de DeepSeek utilisent « mode PTC » sans en donner une forme développée officielle, même si le comportement correspond à « programmatic tool calling ». J’en ferais une définition de travail, pas un sigle confirmé, tant que DeepSeek n’en précise pas un explicitement.**

### Puis-je confier à DeepSeek Harness un dépôt qui compte pour moi ?

**Certaines limites demeurent. Pour un dépôt important, travaillez sur une copie ou une branche séparée, gardez les identifiants de production hors de l’environnement et examinez chaque plugin avant de le charger.**

### « Tout est un plugin » signifie-t-il que je peux en faire n’importe quel type d’agent ?

**Pas sans un vrai travail d’ingénierie. Remplacer l’adaptateur de modèle ou la boucle d’agent nécessite un plugin qui respecte le contrat de service adéquat. Le système de plugins vous donne accès à plus de pièces ; il ne fait pas disparaître le travail.**
