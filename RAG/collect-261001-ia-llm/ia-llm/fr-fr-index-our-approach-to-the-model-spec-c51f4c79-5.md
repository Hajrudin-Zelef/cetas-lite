---
id: collect-261001-ia-llm/ia-llm/fr-fr-index-our-approach-to-the-model-spec-c51f4c79-5
title: "fr-fr-index-our-approach-to-the-model-spec-c51f4c79"
domain: ia-llm
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["agents", "agi", "merger"]
source: docs/RAG/collect-261001-ia-llm/fr-fr-index-our-approach-to-the-model-spec-c51f4c79.md
source_anchor: ""
source_lines: [61, 73]
sha256: 5b81472d20cbd448aa79883efc55f3c3ae0b25d4d1cf31043e6a2f2cbfb3d7d2
---

# fr-fr-index-our-approach-to-the-model-spec-c51f4c79

- Nouvelles fonctionnalités et nouveaux produits. Alors que les modèles acquièrent de nouvelles capacités et que nous lançons de nouveaux produits, nous souhaitons que la spécification du modèle évolue en contenu comme en couverture — par exemple en ajoutant des règles pour les interactions multimodales(ouverture dans une nouvelle fenêtre), les agents autonomes(ouverture dans une nouvelle fenêtre) et les utilisateurs de moins de 18 ans(ouverture dans une nouvelle fenêtre).
Quelques principes de conception guident la manière dont nous rédigeons et révisons la spécification du modèle.
- Clarté et précision. « Être honnête » est une valeur positive, mais pas une procédure de décision complète. La spécification du modèle devrait clarifier les désaccords, et non les dissimuler derrière un langage consensuel. Lorsque c’est possible, nous devrions signaler explicitement les conflits potentiels entre les règles et fournir des indications ou des exemples pour les résoudre. Par exemple, Ne pas mentir(ouverture dans une nouvelle fenêtre) met en évidence un conflit potentiel avec Être chaleureux(ouverture dans une nouvelle fenêtre), en expliquant que l’assistant doit respecter les normes de politesse, sans aller jusqu’à des mensonges par souci de complaire qui pourraient relever de la complaisance(ouverture dans une nouvelle fenêtre) et aller à l’encontre de l’intérêt de l’utilisateur.
- Règles de fond. Un lecteur doit pouvoir prendre un prompt réaliste et produire une réponse qu’un autre lecteur reconnaîtra clairement comme étant dans les limites ou hors des limites (même s’il subsiste des zones de jugement en marge).
- Exemples qui maximisent le rapport signal/bruit. De bons exemples sont souvent essentiels à l’élaboration d’une mise à jour de spécifications de haute qualité. Les exemples doivent aider à mettre en lumière les principales difficultés liées à la définition du comportement des modèles, en faisant émerger les conflits complexes et en adoptant une position claire sur la manière de les résoudre. De manière secondaire, ils doivent aussi servir d’exemples du ton et du style souhaités, ce qui peut être difficile à transmettre uniquement par écrit.
- Robustesse. Nous essayons d’éviter les exemples comportant une ambiguïté ou une complexité superflue, afin que le conflit principal et la résolution attendue soient clairs.
- Cohérence et organisation claire. Nous nous efforçons de garantir que les règles de spécification du modèle soient parfaitement cohérentes entre elles et avec le comportement du modèle que nous visons, tout en rendant l’organisation générale du document claire et accessible.
La spécification du modèle ne prétend pas que nous pouvons formaliser tout ce qui compte, ni que les modèles atteindront toujours l’objectif. C’est l’affirmation que le comportement visé est suffisamment important pour être clair, exploitable et révisable.
Trois critères de réussite guident la manière dont nous le faisons évoluer.
- Lisibilité. Les personnes, qu’elles soient chez OpenAI ou en dehors, peuvent se forger des attentes claires quant au comportement et s’appuyer sur le texte lorsque celui-ci les surprend.
- Capacité d’action. La spécification du modèle peut être utilisée pour concevoir des évaluations, diagnostiquer des incidents et prendre des décisions produit cohérentes—et pas seulement pour exprimer des valeurs.
- Capacité de révision. La spécification du modèle peut évoluer alors que nous apprenons, sans devenir une cible instable en constante évolution.
Alors que les modèles et les produits évoluent, nous nous attendons à ce que la spécification du modèle s’enrichisse et se précise alors que de nouvelles capacités et de nouveaux contextes de déploiement apparaissent. L’objectif est de maintenir une spécification comportementale cohérente, testable et alignée avec notre mission de faire en sorte que l’AGI bénéficie à l’ensemble de l’humanité.
