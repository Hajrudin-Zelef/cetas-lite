---
id: collect-261001-general-networking/general-networking/dagster-vs-airflow-comparaison-des-meilleurs-outils-dorchestration-de-donnees-pour-des-dat-1
title: "dagster-vs-airflow-comparaison-des-meilleurs-outils-dorchestration-de-donnees-pour-des-data-stacks-m"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "open source"]
source: docs/RAG/collect-261001-general-networking/dagster-vs-airflow-comparaison-des-meilleurs-outils-dorchestration-de-donnees-pour-des-data-stacks-m.md
source_anchor: ""
source_lines: [1, 100]
sha256: b6ec4834cf7e3b2e1264914f28e10d5f5d24b2dca68dbb53940072f53e418a1f
---

# dagster-vs-airflow-comparaison-des-meilleurs-outils-dorchestration-de-donnees-pour-des-data-stacks-m

Cours

À chaque scroll sur LinkedIn, on dirait qu’un nouvel outil apparaît dans un paysage de la donnée déjà bien encombré. À mesure que les équipes data empilent les briques de leur « data stack », elles ont besoin d’un moyen de piloter et de relier ces outils disparates.

L’orchestration de données désigne le fait de concevoir, exécuter et superviser des processus qui combinent, transforment et organisent les données à l’échelle de tout l’écosystème. Pour cela, les data engineers s’appuient sur des outils d’orchestration des données. Ici, nous passons en revue deux des solutions les plus populaires : Apache Airflow et Dagster. C’est parti !

## Qu’est-ce qu’Airflow ?

Airflow est la référence du marché pour construire, exécuter et superviser des pipelines de données en tant que code. Initialement développé en 2014 par l’équipe Data Engineering d’Airbnb, le projet a depuis été adopté par la fondation Apache Software Foundation et est devenu l’offre la plus populaire sous leur licence.

Chaque mois, Airflow est téléchargé plus de trente millions de fois et bénéficie d’une communauté d’utilisateurs et de contributeurs très active. Dans Airflow, les pipelines de données sont appelés DAGs, pour directed acyclic graphs (graphes acycliques dirigés).

Pour en savoir plus, consultez notre tutoriel bien démarrer avec Apache Airflow.

## Qu’est-ce que Dagster ?

Comme Airflow, Dagster est un outil d’orchestration qui permet de définir des pipelines de données en tant que code. Open source, il a été lancé en 2019. Dagster utilise Python pour définir des « assets », les briques de base d’un pipeline de données.

Dagster s’est imposé comme un concurrent d’Airflow dans un espace relativement peu encombré. Le projet est maintenu par sa communauté open source et bénéficie d’un soutien commercial de Dagster Labs.

## Relier une data stack moderne

Alors, où Airflow et Dagster s’inscrivent-ils dans une data stack moderne ? Partons du schéma d’architecture ci-dessous.

Schéma d’architecture présentant un échantillon d’outils courants dans une data stack moderne.

Ce schéma illustre ce que beaucoup considèrent comme une data stack standard. Elle se compose de systèmes sources, d’outils d’entrepôt et de transformation de données, ainsi que de « destinations » en aval comme Tableau et Looker.

Ici, Airflow et Dagster sont les flèches qui relient cet écosystème de données. Sans outil d’orchestration, il est difficile de faire circuler la donnée d’un outil à l’autre.

Des outils comme Airflow et Dagster ajoutent aussi une couche d’observabilité, pour comprendre facilement où, comment et quand les données sont acheminées de la source à la destination, et comment elles sont transformées au passage.

## Dagster vs Airflow : fonctionnalités clés

Vous l’aurez deviné, Airflow et Dagster partagent de nombreux points communs dans leur usage et leur rôle dans une data stack moderne. Parmi lesquels :

- Possibilité de définir des pipelines en tant que code
- Intégrations natives avec la data stack moderne
- Expérience de développement local

Chacun dispose toutefois d’atouts et de fonctionnalités qui lui sont propres.

### Airflow

Commençons par les principales fonctionnalités d’Airflow :

#### Comprendre Airflow et les DAGs

Dans Airflow, les pipelines de données sont appelés DAGs — pour graphes acycliques dirigés. Imaginez un DAG comme un ensemble de tâches reliées dans un ordre précis. Le composant le plus élémentaire d’un DAG est la tâche. Par exemple, dans un pipeline d’extraction, transformation, chargement (ETL), l’étape « transform » est une tâche à part entière.

Ces tâches sont généralement définies via des operators Airflow. Mais, comme vous allez le voir, il existe une autre façon de définir des tâches grâce à la TaskFlow API.

#### Programmer des DAGs dans Airflow

La planification d’un pipeline dans Airflow est d’une flexibilité redoutable. Besoin de l’exécuter chaque jour ? Aucun problème. Le premier vendredi de chaque mois ? Facile aussi. Et si votre manager vous demande de déclencher un DAG à la mise à jour d’un jeu de données ? Airflow sait le faire également.

Airflow offre de nombreuses options de personnalisation pour la planification des DAGs. Avec CRON, les Timetables et le data-aware scheduling, vous pouvez exécuter vos DAGs quand vous en avez besoin.

#### TaskFlow API

Pour rendre l’écriture de DAGs plus accessible aux professionnels de la donnée, la communauté Airflow a introduit la TaskFlow API. Plutôt que d’utiliser uniquement des operators, la TaskFlow API permet de définir des tâches en décorant simplement des fonctions. Cela facilite grandement le partage de données entre tâches et la définition des dépendances.

De plus, des outils comme l’Astro SDK se sont bâtis au-dessus de la TaskFlow API, élargissant encore ses possibilités.

### Dagster

Voyons maintenant comment Dagster se positionne :

#### L’approche centrée sur les assets de Dagster

Dagster adopte une approche fondée sur les assets pour construire des pipelines. Dans Dagster, tout objet de données stocké de manière persistante — fichier, table, etc. — est un asset.

Ces assets sont définis en code via des fonctions Python. Lors de leur exécution, Dagster crée automatiquement les dépendances et matérialise l’asset. Cette approche centrée asset rend le suivi de la production et de la consommation des données au sein d’un pipeline particulièrement simple.

#### Comprendre les ops dans Dagster

Autre concept clé dans Dagster : les « ops », assez proches des tâches dans Airflow. Les ops représentent des étapes individuelles d’un pipeline, avec des entrées et sorties qui peuvent être des assets. Mieux encore, ces entrées et sorties peuvent être typées pour préciser clairement les données traitées.

Quelle relation entre ops et assets ? Les ops correspondent aux étapes (extract, transform, load) qui opèrent sur les assets, c’est-à-dire les objets de données eux-mêmes.

#### Le système de typage de Dagster

Pour clarifier les flux de données d’un pipeline, Dagster s’appuie sur un système de typage robuste afin de valider les entrées et sorties de chaque op. Airflow prend en charge les types Python, mais Dagster en a fait un pilier de l’écriture des pipelines.

Vous pouvez typer non seulement les ops, mais aussi les définitions d’assets, afin de garantir que les données produites par un asset sont correctes et validées.

## Airflow vs Dagster : expérience développeur

Pour choisir un outil d’orchestration, on évalue souvent la richesse fonctionnelle, le coût total de possession et la scalabilité. Cela dit, il est crucial de considérer l’expérience développeur. Une expérience intuitive, efficace et robuste favorise l’itération rapide et anticipe des enjeux comme les tests unitaires et la gestion des dépendances.

### Airflow

#### La communauté Airflow

Une question ou un bug Airflow que vous n’arrivez pas à résoudre ? Il y a de fortes chances que la réponse existe déjà. Airflow possède l’une des plus grandes communautés Slack parmi les projets Apache et compte plus de trois mille contributeurs. Des milliers de blogs, tutoriels et cours ont été consacrés à Airflow. En plus d’aider les utilisateurs, la communauté construit et maintient le projet lui-même.

Astronomer, l’éditeur commercial derrière Apache Airflow, a hissé l’expérience développeur à un autre niveau. La CLI astro ajoute une couche d’abstraction et de fonctionnalités aux utilitaires Airflow. Astronomer propose aussi sa registry et le chatbot « Ask Astro » pour répondre à toutes les questions Airflow qu’un développeur pourrait avoir.

#### Développer et tester en local

