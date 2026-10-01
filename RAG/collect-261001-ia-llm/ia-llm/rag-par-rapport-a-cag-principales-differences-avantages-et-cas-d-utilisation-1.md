---
id: collect-261001-ia-llm/ia-llm/rag-par-rapport-a-cag-principales-differences-avantages-et-cas-d-utilisation-1
title: "rag-par-rapport-a-cag-principales-differences-avantages-et-cas-d-utilisation"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-ia-llm/rag-par-rapport-a-cag-principales-differences-avantages-et-cas-d-utilisation.md
source_anchor: ""
source_lines: [1, 70]
sha256: 6cc8983ac0504a401ca677099ac0de93bef790ffe9edc4d931df3bc2637b85cc
---

# rag-par-rapport-a-cag-principales-differences-avantages-et-cas-d-utilisation

Cours

Alors que l'intelligence artificielle continue d'évoluer, l'un des principaux défis consiste à déterminer comment intégrer efficacement les connaissances dans les grands modèles linguistiques (LLM), compte tenu de leurs connaissances limitées. Afin de surmonter ces contraintes, les chercheurs et les praticiens ont exploré différentes approches en matière d'intégration des connaissances.

Deux des approches les plus importantes à l'heure actuelle sont la génération augmentée par la récupération (RAG) et la génération augmentée par le cache (CAG). J', j'ai travaillé avec les deux approches, et bien qu'elles soient souvent présentées comme concurrentes, j'ai constaté qu'elles s'apparentent davantage à des outils différents pour des tâches différentes, parfois même plus efficaces lorsqu'elles sont utilisées conjointement.

Dans cet article, je vais vous présenter une comparaison entre RAG et CAG, en explorant la signification de chaque concept, leur fonctionnement et leur utilisation optimale dans des applications concrètes. À la fin, vous comprendrez en quoi ces approches diffèrent, où elles se recoupent et comment choisir entre elles, voire les combiner, lors de la conception de systèmes d'IA.

Si vous souhaitez aller au-delà des concepts et commencer à développer vous-même ces systèmes, je vous recommande de suivre notre cours pratique intitulé « Retrieval Augmented Generation (RAG) avec LangChain ».

## Qu'est-ce que la génération augmentée par la récupération (RAG) ?

La génération augmentée par récupération est une technique d'te qui permet aux modèles d'IA d'aller au-delà de leurs données d'entraînement fixes et d'intégrer dynamiquement des informations externes. Au lieu de se fier uniquement à ce qui a été codé dans le modèle pendant l'entraînement,

RAG relie le modèle à des bases de données externes et à des mécanismes de recherche, lui permettant ainsi de récupérer des documents ou des connaissances pertinents au moment d'une requête.

Cette idée a gagné en popularité lorsque les organisations ont réalisé que les données de formation statiques devenaient rapidement obsolètes. J'ai observé comment les informations évoluent quotidiennement dans de nombreux secteurs, et un modèle dépourvu de couche de récupération externe ne peut pas suivre le rythme.

RAG a été développé pour combler cette lacune et intégrer des connaissances nouvelles, spécifiques à un domaine ou dynamiques directement dans le processus de génération.

### Comment fonctionne le RAG ?

Le flux de travail RAG commence par une requête utilisateur. La requête est d'abord codée en une représentation vectorielle, qui est ensuite utilisée pour effectuer une recherche dans une base de données vectorielle (système de recherche) contenant des documents, des enregistrements ou d'autres sources de connaissances. Cette étape de récupération garantit que le modèle identifie les informations externes les plus pertinentes avant de poursuivre.

À ce stade, il est essentiel de mettre en place des stratégies de segmentation efficaces : les documents sont divisés en unités de sens plus petites, généralement comprises entre 100 et 1 000 tokens, afin que le système de recherche puisse faire apparaître le contexte le plus pertinent sans surcharger le modèle de génération.

Les algorithmes de recherche, souvent basés sur la recherche approximative du plus proche voisin, garantissent que les informations pertinentes sont récupérées rapidement, même à partir de bases de connaissances à grande échelle.

Une fois les documents pertinents récupérés, ils sont transmis à l'étape de génération, où le modèle linguistique intègre ces informations dans sa réponse. Ce processus permet au système de fournir des réponses non seulement plus cohérentes, mais également fondées sur des connaissances externes actualisées.

*Flux de travail RAG*

Les sources de connaissances externes peuvent inclure des bases de données propriétaires, des articles scientifiques, des archives juridiques ou même des API en temps réel. Le moteur de recherche constitue le pont qui permet au modèle linguistique de combiner sa capacité générative avec des données factuelles. Considérez cela comme si vous fournissiez à votre IA une carte de bibliothèque plutôt que d'espérer qu'elle mémorise chaque ouvrage.

Cette structure de base peut être affinée en appliquant des techniques RAG avancées ou en utilisant Corrective RAG (CRAG), uneversion améliorée de RAG optimisée pour la précision.

Maintenant que nous avons examiné la structure et le fonctionnement du RAG, il est plus facile d'évaluer ce qui rend cette méthode particulièrement efficace dans des scénarios réels.

### Points forts de RAG

Ce que j'apprécie le plus chez RAG, c'est sa capacité à gérer le changement. Votre service juridique met à jour une politique à 15 h ? Votre système RAG en est informé à 15 h 01, sans qu'aucune formation supplémentaire ne soit nécessaire.

D'après mon expérience, RAG se distingue dans trois domaines principaux :

- 
**Mises à jour en temps réel :** La couche de récupération se connecte à des connaissances externes, fournissant des réponses basées sur les données les plus récentes. Cela rend le RAG particulièrement précieux dans des domaines en constante évolution tels que la médecine, la finance ou la technologie.
- 
**Réduction des hallucinations:** Les modèles de langage à grande échelle (LLM) génèrent souvent des textes qui semblent plausibles, mais qui sont factuellement incorrects. En fondant ses réponses sur des documents récupérés, le RAG garantit que les résultats sont ancrés dans la réalité, ce qui renforce leur fiabilité.
- 
**Intégration flexible des données :** Les connaissances externes peuvent provenir de multiples sources, telles que des bases de données structurées, des API semi-structurées ou des référentiels de texte non structurés. Les organisations peuvent adapter les pipelines de récupération à leurs besoins spécifiques.

Bien que le RAG offre des avantages convaincants, il est tout aussi important de comprendre les défis et les contraintes liés à cette approche.

### Limites du RAG

C'est là que cela devient délicat, et c'est ce dont j'informe mes clients dès le départ. RAG implique de véritables compromis :

- 
**Complexité du système :** Il est nécessaire d'orchestrer le système de récupération, la base de données vectorielle et le modèle de génération. Cette complexité engendre des points de défaillance supplémentaires et augmente les frais généraux liés à la maintenance.
- 
**Problèmes de latence :** Le processus de récupération ajoute une charge informatique supplémentaire à chaque requête. La recherche dans de vastes bases de connaissances et la récupération de documents pertinents prennent du temps, ce qui peut nuire à l'expérience utilisateur dans les applications en temps réel.
- 
**Dépendance de la qualité de la récupération :** La qualité de vos réponses dépend de l'efficacité de votre mécanisme de récupération. Une récupération inadéquate implique que des informations contextuelles non pertinentes sont transmises au modèle linguistique, ce qui peut nuire à la qualité de la réponse.

Cependant, il existe plusieurs techniques clés pour améliorer les performances du RAG et traiter efficacement ces problèmes.

Après avoir examiné l'approche axée sur la récupération, je vais maintenant aborder la génération augmentée par cache, qui suit une voie très différente pour améliorer les performances du modèle.

## Qu'est-ce que la génération augmentée par cache (CAG) ?

