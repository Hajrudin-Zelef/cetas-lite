---
id: collect-261001-huawei/huawei/kubernetes-vs-docker-differences-que-tout-developpeur-doit-connaitre-2
title: "Use the official Python base image with version 3.9"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-huawei/kubernetes-vs-docker-differences-que-tout-developpeur-doit-connaitre.md
source_anchor: ""
source_lines: [90, 180]
sha256: b7b44ea619ae566821d6a0aabc2a79e7d7f47ef670ec1de00e844ab381f88a0c
---

# Use the official Python base image with version 3.9

*Évolution des stratégies de déploiement dans le temps. Source de l'image : Kubernetes.io*

### Comment fonctionne Kubernetes

Kubernetes repose sur le concept de clusters, de nœuds et de pods, formant une architecture en couches qui offre flexibilité et évolutivité. Un cluster représente l'ensemble de l'infrastructure, composée de plusieurs nœuds (machines virtuelles ou physiques).

Ces nœuds travaillent ensemble pour héberger et gérer des applications conteneurisées. Les nœuds peuvent être soit des nœuds maîtres, qui contrôlent et gèrent la grappe, soit des nœuds ouvriers, qui exécutent les charges de travail de l'application. Le nœud maître est chargé de gérer l'état de la grappe, de prendre des décisions en matière de programmation et de surveiller son état de santé.

Chaque nœud de travailleur exécute un ou plusieurs pods, qui sont les plus petites unités déployables dans Kubernetes et consistent en un ou plusieurs conteneurs.

Les pods agissent comme des hôtes logiques pour les conteneurs et partagent le même réseau et le même stockage, ce qui facilite la communication entre les conteneurs au sein d'un pod. Les pods sont éphémères par nature, ce qui signifie qu'ils peuvent être créés, détruits ou répliqués dynamiquement en fonction des besoins de l'application.

Vue d'ensemble de l'architecture de Kubernetes. Source de l'image : Kubernetes.io

Kubernetes fait abstraction de la complexité de la gestion de l'infrastructure en fournissant une API puissante et une suite d'outils pour gérer les applications conteneurisées. Il assure le bon fonctionnement des applications en répartissant les charges de travail, en adaptant les ressources en fonction de la demande et en redémarrant les conteneurs en cas de défaillance.

Kubernetes gère également l'état souhaité de vos applications, en veillant à ce que le nombre de pods et leur configuration correspondent toujours à ce que vous spécifiez et que les perturbations soient automatiquement corrigées. Cette automatisation réduit les efforts manuels nécessaires à la gestion de l'infrastructure et améliore la fiabilité et la résilience de vos applications.

### Caractéristiques de Kubernetes

- Mise à l'échelle automatisée: Kubernetes peut faire évoluer les applications automatiquement en fonction des demandes de ressources, en optimisant l'utilisation et en maintenant des performances constantes.
- Équilibrage de la charge: Kubernetes répartit efficacement le trafic réseau entrant entre plusieurs conteneurs, ce qui garantit la disponibilité et la résilience.
- Découverte des services: Kubernetes fournit des services pour découvrir automatiquement les conteneurs, ce qui élimine la nécessité de gérer manuellement les points d'extrémité.
- Mises à jour en continu: Kubernetes permet de mettre à jour les applications avec un temps d'arrêt minimal, ce qui garantit la stabilité et la fiabilité lors des mises à niveau.

## Kubernetes vs Docker : Différences fondamentales

Nous avons maintenant une meilleure compréhension de Docker et de Kubernetes, il est donc temps de mettre en évidence leurs principales différences :

### 1. Objectif et fonction

Docker et Kubernetes résolvent des problèmes différents dans le processus de conteneurisation. Docker est utilisé pour construire, expédier et exécuter des conteneurs. Il permet de créer des environnements isolés pour les applications.

En revanche, Kubernetes se concentre sur l'orchestration des conteneurs, ce qui signifie qu'il aide à gérer, à mettre à l'échelle et à assurer le bon fonctionnement de grandes collections de conteneurs.

### 2. Gestion des conteneurs

Docker gère des conteneurs individuels, tandis que Kubernetes gère plusieurs conteneurs à travers des clusters.

Docker offre des capacités d'orchestration de base grâce à Docker Compose et Docker Swarm, mais Kubernetes fait passer l'orchestration au niveau supérieur, en gérant des scénarios complexes impliquant des milliers de conteneurs.

### 3. Orchestration d'applications

En ce qui concerne l'orchestration avancée, Kubernetes offre des fonctionnalités telles que l'autorégénération, l'équilibrage de charge, les déploiements automatisés et la mise à l'échelle.

Docker Swarm est le propre outil d'orchestration de Docker, mais Kubernetes s'est imposé comme la solution privilégiée pour l'orchestration d'environnements complexes à grande échelle en raison de ses capacités avancées et de la prise en charge plus large de l'écosystème.

Docker vs Kubernetes. Source de l'image : Alex Xu / ByteByteGo

## Cas d'utilisation de Docker

En gardant à l'esprit les informations précédentes sur Docker, voici quelques-uns des cas d'utilisation les plus courants :

### 1. Développement et essais locaux

Docker est un outil précieux pour le développement local. Les développeurs peuvent créer des environnements conteneurisés qui imitent les paramètres de production, garantissant ainsi un comportement cohérent tout au long du cycle de développement du logiciel.

### 2. Applications légères

Docker est un excellent choix pour les cas d'utilisation plus simples qui ne nécessitent pas d'orchestration. Sa simplicité s'illustre dans des scénarios tels que l'exécution d'applications à petite échelle ou le déploiement de services autonomes.

### 3. Pipelines CI/CD

Docker est largement utilisé dans les pipelines d'intégration et de déploiement continus (CI/CD). Il garantit que chaque étape - de la création du code aux tests - est réalisée dans un environnement cohérent et reproductible, ce qui permet de réduire les surprises lors de la production.

## Cas d'utilisation de Kubernetes

Kubernetes est le plus souvent utilisé dans les scénarios suivants :

### 1. Gérer des applications conteneurisées à grande échelle

Kubernetes brille dans les environnements à grande échelle. Il peut gérer des milliers de conteneurs répartis sur plusieurs nœuds dans un cluster distribué. Des organisations comme Spotify et Airbnb utilisent Kubernetes pour assurer le bon fonctionnement de leurs applications complexes basées sur des microservices.

### 2. Mise à l'échelle et résilience automatisées

Kubernetes fait évoluer automatiquement les conteneurs en fonction des exigences du système, répondant ainsi de manière dynamique aux demandes fluctuantes. En outre, Kubernetes dispose de mécanismes d'autoréparation intégrés - redémarrage des conteneurs défaillants et remplacement des nœuds non réactifs pour maintenir le temps de fonctionnement des applications.

### 3. Architecture microservices

Kubernetes est idéal pour gérer les microservices dans les environnements de production. Sa capacité à gérer de nombreux services et leurs dépendances tout en facilitant la communication entre eux en fait un outil idéal pour les applications complexes et distribuées.

## Kubernetes et Docker peuvent-ils fonctionner ensemble ?

À présent, il est facile de voir que Docker et Kubernetes sont destinés à travailler ensemble.

Kubernetes utilise des runtimes de conteneurs pour exécuter des conteneurs individuels, et Docker est traditionnellement l'un de ces runtimes de conteneurs. Bien que Kubernetes et Docker aient des rôles distincts, ils fonctionnent très bien ensemble ! Docker construit et exécute des conteneurs, tandis que Kubernetes orchestre ces conteneurs à travers des clusters.

### Docker Swarm vs Kubernetes

Docker Swarm est l'outil d'orchestration natif de Docker, adapté aux environnements plus simples et moins exigeants.

