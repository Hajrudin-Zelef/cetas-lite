---
id: collect-261001-rattrapage/rattrapage/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs-1
title: "containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["arr", "aws"]
source: docs/RAG/collect-261001-rattrapage/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs.md
source_anchor: ""
source_lines: [1, 67]
sha256: 425874685d64f9463bed599643c84ef385f803401386add46ec38838772778a1
---

# containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs

Cours

Dans l'écosystème des conteneurs, une source fréquente de confusion survient lorsque les développeurs comparent Docker (une plateforme complète) à containerd, qui est en réalité un composant d'exécution spécialisé. Cette comparaison revient à comparer une automobile complète à son moteur : les deux sont essentiels, mais ils remplissent des fonctions différentes à des niveaux d'abstraction différents.

Dans ce guide, je vais clarifier la relation entre ces deux technologies, en explorant leurs architectures, l'intégration de Kubernetes et les scénarios spécifiques dans lesquels chaque outil excelle. Que vous construisiez des conteneurs localement ou que vous gériez des clusters Kubernetes de production, comprendre quand utiliser Docker plutôt que containerd peut avoir un impact significatif sur vos décisions en matière d'infrastructure.

À la fin de cet article, vous aurez une compréhension claire de la manière dont Docker et containerd se complètent, et vous saurez comment choisir l'outil adapté à vos besoins spécifiques, qu'il s'agisse de développement local rapide ou de déploiements de production à grande échelle.

Si vous débutez dans le domaine de la conteneurisation, je vous recommande vivement de suivre notre cours sur les concepts de conteneurisation et de virtualisation.

## Qu'est-ce que Docker ?

Docker a révolutionné le paysage du développement logiciel en rendant la conteneurisation accessible et pratique pour les développeurs au quotidien. En tant que plateforme complète, Docker fournit tous les éléments nécessaires pour créer, livrer et exécuter des applications conteneurisées.

Pour les débutants, ce guide pratique sur les conteneursconstitue une excellente introduction à Docker et aux conteneurs.

### Une plateforme de conteneurs complète

Docker fonctionne comme une plateforme de conteneurs tout-en-un , offrant une pile d'outils intégrée qui couvre la majeure partie du cycle de vie des conteneurs. À la base, Docker regroupe les applications et toutes leurs dépendances dans des images portables et autonomes qui peuvent fonctionner de manière cohérente dans n'importe quel environnement, de l'ordinateur portable d'un développeur aux serveurs de production.

Cette approche globale a fait de Docker la norme industrielle en matière de développement local et d'expérience développeur. Parmi les principaux avantages, on peut citer :

- **Écosystème étendu :** Docker Hub héberge des millions d'images pré-construites, allant des bases de données aux serveurs web, ce qui élimine les procédures d'installation complexes.
- **Portabilité universelle :** Un conteneur créé sur macOS fonctionne de manière identique sur Linux ou Windows, à condition que Docker soit installé.
- **Priorité aux développeurs :** Résume les différences d'infrastructure afin que les développeurs puissent se concentrer sur leurs applications plutôt que sur les complexités du déploiement.
- **Intégration rapide :** Les nouveaux membres de l'équipe peuvent mettre en place des environnements de développement en quelques minutes plutôt qu'en plusieurs heures.

### Éléments clés et processus de travail

L'architecture de Docker se compose de plusieurs composants interconnectés qui fonctionnent ensemble de manière transparente.

L'interface CLI Docker fournit l'interface utilisateur à partir de laquelle les développeurs exécutent des commandes. Lorsque vous exécutez une commande telle que ` `docker run``, l'interface CLI communique avec le démon Docker (``dockerd``), qui sert de service central pour gérer les conteneurs, les images, les réseaux et les volumes.

Pour les utilisateurs Windows et macOS, Docker Desktop offre une couche supplémentaire de commodité grâce à une interface graphique, rendant la gestion des conteneurs accessible même à ceux qui sont moins à l'aise avec les outils en ligne de commande. Cet environnement intégré comprend tout ce qui est nécessaire au développement local, du support Kubernetes à la gestion des volumes.

Le flux de travail type suit un schéma clair :

1. Les développeurs créent des images à l'aide de fichiers Dockerfiles.
2. Ils les expédient vers des registres tels que Docker Hub.
3. Enfin, les images sont exécutées en tant que conteneurs sur n'importe quel hôte compatible avec Docker.

Ce que de nombreux développeurs ignorent, c'est que Docker n'exécute pas directement les conteneurs. `dockerd` e la délégation de l'exécution effective du conteneur à containerd, qui s'exécute en arrière-plan en tant que processus distinct.

Cette conception modulaire, dans laquelle Docker utilise containerd pour les opérations de bas niveau, permet à chaque composant de se concentrer sur ce qu'il fait le mieux : Docker offre une interface et un écosystème conviviaux pour les développeurs, tandis que containerd gère les détails techniques de l'exécution des conteneurs.

Que vous soyez novice dans l'utilisation de Docker ou que vous souhaitiez passer à l'étape suivante, nous vous invitons à consulter nos 10 idées de projets Docker adaptés à tous les niveaux.

## Qu'est-ce que Containerd ?

Après avoir examiné l'approche globale de la plateforme Docker, nous allons maintenant nous intéresser à containerd, le runtime spécialisé qui permet l'exécution des conteneurs à la fois dans Docker et dans l'écosystème plus large.

### Un environnement d'exécution de conteneurs conforme aux normes de l'industrie

Containerd est un moteur d'exécution de conteneurs léger, conforme aux normes de l'industrie, qui a été certifié par la Cloud Native Computing Foundation (CNCF), ce qui témoigne de sa maturité, de sa fiabilité et de son adoption généralisée. Initialement intégré à Docker, containerd a été extrait en 2017 et transféré à la CNCF afin de permettre une adoption plus large au sein de l'écosystème.

Contrairement à l'ensemble complet de fonctionnalités de Docker, la portée de containerd est délibérément limitée : il gère les opérations essentielles du cycle de vie des conteneurs telles que le démarrage, l'arrêt, la mise en pause et la suppression, et prend en charge le transfert et le stockage des images. Cette approche ciblée rend containerd exceptionnellement stable et efficace, des qualités essentielles pour les environnements de production.

Considérez containerd comme une « plomberie » : une infrastructure conçue pour être intégrée dans des systèmes plus vastes plutôt que d'être utilisée directement par les utilisateurs. Les principales plateformes telles que Kubernetes, AWS Fargate et Google Kubernetes Engine s'appuient toutes sur containerd pour exécuter des conteneurs, même si les utilisateurs interagissent avec ces plateformes via leurs propres interfaces.

### Architecture et design

Pour appréhender pourquoi containerd est si largement adopté dans les systèmes de production, il est nécessaire d'examiner ses principes architecturaux.

L'architecture de Containerd illustre parfaitement les principes de conception modulaire. Conçu selon les normes de l'Open Container Initiative (OCI), containerd garantit la compatibilité dans l'ensemble de l'écosystème des conteneurs. Cela signifie que toute image conforme à OCI fonctionnera avec containerd, quel que soit l'outil utilisé pour la créer.

L'architecture se compose de plusieurs couches principales :

