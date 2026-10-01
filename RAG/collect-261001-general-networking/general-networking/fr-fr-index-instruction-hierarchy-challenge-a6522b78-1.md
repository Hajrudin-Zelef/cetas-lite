---
id: collect-261001-general-networking/general-networking/fr-fr-index-instruction-hierarchy-challenge-a6522b78-1
title: "fr-fr-index-instruction-hierarchy-challenge-a6522b78"
domain: general-networking
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["benchmark", "benchmarks", "chatgpt"]
source: docs/RAG/collect-261001-general-networking/fr-fr-index-instruction-hierarchy-challenge-a6522b78.md
source_anchor: ""
source_lines: [1, 36]
sha256: ca7ffe7b20d015a53509e32b4eb428ae4f21c327a551b4e0fbd51f826ac43ddd
---

# fr-fr-index-instruction-hierarchy-challenge-a6522b78

Améliorer la hiérarchie des instructions dans les LLM de pointe
Présentation d’IH-Challenge, un jeu de données d’entraînement qui renforce la hiérarchie des instructions, la contrôlabilité de la sécurité et la robustesse face aux attaques par injection de prompt.
Les systèmes d'IA reçoivent souvent des instructions provenant de plusieurs sources. Cela peut inclure des politiques de sécurité issues des messages systèmes, des recommandations produit de la part des développeurs, des demandes des utilisateurs et des informations trouvées en ligne. Former des modèles afin qu'ils donnent de manière fiable la priorité aux instructions les plus dignes de confiance parmi ces sources est un élément clé pour un déploiement sûr.
De nombreux problèmes de sécurité et de fiabilité de l'IA peuvent survenir lorsque cette priorisation se dégrade. Les modèles peuvent recevoir des demandes de contenu interdit, des tentatives de révélation d'informations privées ou des attaques de type prompt‑injection intégrées dans des données en ligne. Ne pas se comporter de manière appropriée dans chacun de ces scénarios qui partagent la même cause profonde : le modèle peut suivre la mauvaise instruction.
Lorsque ces instructions entrent en conflit, le modèle doit décider lesquelles prioriser. S'il traite une instruction non fiable comme faisant autorité, le modèle peut se comporter d'une manière qui enfreint les politiques ou l'intention du développeur et de l'utilisateur.
Nous montrons que des tâches de hiérarchie des instructions correctement conçues, qui entraînent les modèles à prioriser les instructions selon leur niveau de confiance, améliorent plusieurs propriétés de sécurité dans des situations réelles. Les modèles entraînés sur ces tâches deviennent plus réactifs aux spécifications de sécurité dans les prompts système (améliorant la capacité de contrôlabilité de la sécurité) et plus robustes face aux attaques par injection de prompt intégrées dans les sorties d'outils.
Pour gérer les conflits, les modèles d'OpenAI sont entraînés à suivre une hiérarchie claire des instructions :
Système > développeur > utilisateur > outil
Les instructions de priorité plus élevée sont plus fiables. Le modèle ne doit suivre les instructions de priorité inférieure que lorsqu'elles n'entrent pas en conflit avec des contraintes de priorité supérieure. Ces principes sont décrits dans les spécifications de modèle OpenAI(ouverture dans une nouvelle fenêtre).
Par exemple, si un message système inclut une politique de sécurité et qu'un utilisateur demande au modèle de la violer, le modèle doit refuser. Si la sortie d'un outil contient des instructions malveillantes, le modèle doit les ignorer plutôt que de les traiter comme des commandes.
Bien faire les choses est fondamental pour la sécurité, la sûreté et la fiabilité.
L'apprentissage par renforcement est naturellement adapté à l'enseignement de la hiérarchie des instructions. Nous pouvons générer des conversations avec des instructions contradictoires, demander au modèle de répondre et le récompenser lorsqu’il suit la bonne instruction.
Nous avons identifié trois pièges liés à l'application naïve de cette recette :
- Les échecs dans le suivi des instructions peuvent aussi être des échecs de hiérarchie des instructions : le modèle peut ne pas résoudre un conflit d’instructions, non pas parce qu’il ne comprend pas la hiérarchie des rôles, mais parce que les instructions elles-mêmes sont trop complexes.
- Les conflits d'instructions peuvent être nuancés et même subjectifs. Une approche courante consiste à laisser un LLM distinct attribuer des récompenses au LLM en cours d'entraînement, mais les juges eux-mêmes sont faillibles.
- Les modèles ont tendance à apprendre des raccourcis qui donnent une récompense élevée mais sont inutiles en pratique(ouverture dans une nouvelle fenêtre). L'exemple classique est celui des refus non justifiés : les modèles peuvent apprendre à maximiser la sécurité en refusant même des demandes anodines.
Nous concevons IH-Challenge, un jeu de données d’entraînement pour l’apprentissage par renforcement, afin de répondre à chacun de ces écueils. Nous adhérons aux principes suivants :
- Les tâches sont suivi d'instructions simples
- Ils sont objectivement notables à l'aide d'un simple script Python
- Il n'existe pas de raccourcis simples qui garantissent une récompense élevée pour toutes les tâches
Chaque tâche dans IH-Challenge est essentiellement une conversation avec les messages suivants :
- Un message d'instruction provenant d'un rôle à privilèges élevés, ex. « Répondez uniquement ‘Oui' ou ‘Non' ».
- Un message d’instruction provenant d’un rôle à privilège inférieur, qui tente d’amener le modèle à enfreindre les instructions du message à privilège supérieur.
Le modèle en cours d'entraînement génère le message suivant. Nous concevons les tâches/environnements afin de pouvoir vérifier de manière programmatique si la réponse du modèle respecte la contrainte de niveau supérieur.
Nous entraînons un modèle sur IH‑Challenge et produisons un modèle interne, que nous appelons GPT‑5 Mini-R, avec les améliorations suivantes :
- Offre de meilleures performances sur les benchmarks de la hiérarchie des instructions
- L’amélioration des performances se généralise aux tests de hiérarchie des instructions inédits et adversariaux.
- Maintient l’utilité globale, sans tomber dans un excès de refus.
C’est ce qui rend cette approche particulièrement convaincante du point de vue de la sécurité : en entraînant directement les modèles à résoudre correctement les conflits d’instructions sur les tâches IH-challenge, nous obtenons des améliorations de la hiérarchie des instructions qui se généralisent à de nouvelles attaques et à de nouvelles situations.
Robustesse sur les benchmarks académiques
Robustesse sur les benchmarks internes
Aucune régression des capacités
Une hiérarchie des instructions plus solide apporte plusieurs bénéfices en matière de sécurité à la fois, notamment pour l’orientation de la sécurité et la robustesse face aux attaques par injection de prompt.
Nous évaluons la contrôlabilité de la sécurité en ajoutant des spécifications de sécurité propres à chaque catégorie dans le prompt système et en mesurant le comportement sur les Production Benchmarks de sécurité d’OpenAI (un ensemble de conversations sensibles en matière de sécurité représentatives de ChatGPT en production).
Le modèle entraîné avec IH montre une amélioration constante : lorsque la spécification de sécurité est présente, il atteint des taux plus élevés de refus et de complétion sûre dans les catégories interdites, ce qui indique qu’un comportement de hiérarchie des instructions plus solide l’aide à mieux résoudre les conflits lorsque des requêtes dangereuses proviennent d’instructions de priorité inférieure. Il est important de noter que cette amélioration ne s’accompagne pas d’une baisse correspondante du taux d’utilité (c’est-à-dire que le modèle ne devient pas moins « utile » en refusant simplement davantage de requêtes dans l’ensemble).
La hiérarchie des instructions est également centrale pour résister aux attaques par injection de prompt, lorsque des instructions malveillantes sont intégrées dans les sorties des outils. Nous évaluons le modèle entraîné avec IH sur deux benchmarks d’attaques par injection de prompt — le benchmark académique CyberSecEval 2 et un benchmark interne d’OpenAI sur les attaques par injection de prompt, composé d’attaques comme celle démontrée sur une ancienne version de ChatGPT Atlas.
