---
id: collect-261001-huawei/huawei/kubernetes-vs-docker-differences-que-tout-developpeur-doit-connaitre-1
title: "Use the official Python base image with version 3.9"
domain: huawei
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/kubernetes-vs-docker-differences-que-tout-developpeur-doit-connaitre.md
source_anchor: ""
source_lines: [1, 89]
sha256: d8eda8795e1683c87385026ffd3be9f7d4e54ca1e1c16473563c94747432ec64
---

# Use the official Python base image with version 3.9

Cours

La conteneurisation est devenue la solution idéale pour créer, déployer et mettre à l'échelle des applications modernes de manière efficace. Les deux grands noms de cet espace sont Kubernetes et Docker, et bien qu'ils soient souvent mentionnés ensemble, ils ont en fait des objectifs différents. Ces deux éléments sont essentiels, mais ils ont des fonctions distinctes.

Dans ce guide, je vais vous aider à comprendre ce qui différencie Kubernetes et Docker, leurs fonctionnalités, et quand utiliser chacun d'entre eux.

## Qu'est-ce que la conteneurisation ?

Avant d'aborder Docker et Kubernetes,commençons par comprendre de quoi il s'agit, à savoir la conteneurisation.

La conteneurisation est une forme légère de virtualisation qui regroupe une application et ses dépendances dans une unité unique appelée conteneur.

Contrairement aux machines virtuelles traditionnelles, les conteneurs partagent le système d'exploitation du système hôte mais maintiennent l'isolation entre les applications. Ils sont donc plus efficaces, plus légers et plus rapides à démarrer !

La conteneurisation aide les développeurs à créer des environnements cohérents, portables et faciles à gérer, quel que soit l'endroit où ils s'exécutent, que ce soit sur l'ordinateur portable d'un développeur, dans un centre de données ou dans le cloud.

### Conteneurisation ou virtualisation

Il est utile de comparer la conteneurisation à la virtualisation traditionnelle pour mieux la comprendre. Les machines virtuelles (VM) virtualisent des systèmes matériels entiers, ce qui signifie que chaque VM comprend un système d'exploitation complet ainsi que les binaires et les bibliothèques nécessaires. Cette approche offre l'isolation mais s'accompagne d'une surcharge de ressources importante - chaque VM nécessite son propre système d'exploitation, ce qui la rend gourmande en ressources et plus lente à démarrer.

Les conteneurs, en revanche, partagent le noyau du système d'exploitation hôte, ce qui les rend beaucoup plus légers et rapides à démarrer. Au lieu de virtualiser le matériel, les conteneurs virtualisent le système d'exploitation. Cela permet aux conteneurs d'exécuter des processus isolés sans avoir à gérer un système d'exploitation complet pour chaque instance, ce qui se traduit par une meilleure utilisation des ressources et une plus grande efficacité.

Alors que les machines virtuelles sont idéales pour une isolation totale et l'exécution de plusieurs systèmes d'exploitation différents sur le même matériel, les conteneurs sont mieux adaptés à un déploiement d'applications efficace, évolutif et cohérent.

Machines virtuelles vs. Conteneurs. Source de l'image : contentstack.io

Si vous souhaitez en savoir plus sur l'essentiel des VM, des conteneurs, de Docker et de Kubernetes, consultez le cours gratuit sur les concepts de conteneurisation et de virtualisation sur DataCamp.

Entrons maintenant dans les détails de Docker et Kubernetes !

## Devenez ingénieur en données

## Qu'est-ce que Docker ?

Docker est une plateforme open-source qui offre un moyen léger et portable de créer, déployer et gérer des conteneurs. Contrairement aux machines virtuelles traditionnelles, les conteneurs Docker regroupent tout, y compris le code de l'application, le moteur d'exécution, les outils système et les bibliothèques, ce qui permet aux applications de fonctionner de manière cohérente dans différents environnements.

### Comment fonctionne Docker

Docker fonctionne en créant des conteneurs qui, comme nous l'avons vu précédemment, sont des paquets légers qui encapsulent tous les composants nécessaires à l'exécution d'une application.

Les conteneurs sont construits à partir d'images Docker, qui agissent comme un plan définissant ce qui se trouve à l'intérieur de chaque conteneur. Une image Docker peut inclure un système d'exploitation, des binaires d'application et des fichiers de configuration, ce qui facilite la réplication des environnements.

Une fois l'image créée, les développeurs peuvent utiliser Docker pour exécuter des conteneurs basés sur cette image. L'un des principaux atouts de Docker est sa simplicité et sa cohérence : quel que soit l'endroit où un conteneur est exécuté (sur la machine locale d'un développeur, dans un centre de données sur site ou dans le cloud), le comportement reste le même.

Vue d'ensemble de l'architecture Docker. Source de l'image : Documentation Docker

L'exemple suivant donne un aperçu de la manière dont les images Docker sont mises en œuvre. Jetez un coup d'œil au fichier Docker ci-dessous :

```
# Use the official Python base image with version 3.9
FROM python:3.9
# Set the working directory within the container
WORKDIR /app
# Copy the requirements file to the container
COPY requirements.txt .
# Install the dependencies
RUN pip install -r requirements.txt
# Copy the application code to the container
COPY . .
# Set the command to run the application
CMD ["python", "app.py"]
```
Un `Dockerfile` est un script qui contient une série d'instructions permettant à Docker de construire une image, qui peut ensuite être utilisée pour créer un conteneur.

Après avoir créé un fichier Docker dans votre projet, l'étape suivante consiste à construire l'image Docker. Cette opération s'effectue à l'aide de la commande `docker build`, qui lit les instructions contenues dans le document `Dockerfile` pour assembler l'image. 

Par exemple, l'exécution de `docker build -t my-app .` dans le terminal indique à Docker de construire une image avec l'étiquette `my-app` à partir du répertoire actuel (indiqué par `.`).

Au cours du processus de construction, Docker exécute chaque étape du fichier Docker, comme l'extraction de l'image de base, l'installation des dépendances et la copie du code de l'application dans l'image. Une fois l'image construite, elle sert de modèle qui peut être réutilisé pour créer plusieurs conteneurs.

Une fois l'image construite avec succès, vous pouvez créer et exécuter des conteneurs à partir de celle-ci à l'aide de la commande `docker run`. Par exemple, `docker run my-app` démarre un nouveau conteneur basé sur l'image `my-app`, lançant ainsi votre application dans l'environnement isolé fourni par Docker.

Si vous souhaitez en savoir plus sur les commandes Docker courantes et les meilleures pratiques du secteur, consultez le blog Docker for Data Science : An Introduction.

### Caractéristiques de Docker

- Portabilité: Les conteneurs Docker peuvent fonctionner de manière cohérente sur différents systèmes, offrant ainsi une expérience transparente dans les environnements de développement, de test et de production.
- Facilité d'utilisation: L'interface en ligne de commande de Docker et son ensemble complet d'outils le rendent accessible aux développeurs, même à ceux qui ne connaissent pas encore la conteneurisation.
- Léger: Les conteneurs Docker partagent le même noyau de système d'exploitation, ce qui réduit la surcharge de ressources par rapport aux machines virtuelles complètes.
- Temps de démarrage rapide: Les conteneurs Docker peuvent être démarrés en quelques secondes, ce qui les rend très efficaces pour les applications qui nécessitent un démarrage et un démontage rapides.

Consultez l'antisèche Docker de DataCamp, qui donne un aperçu de toutes les commandes Docker disponibles.

## Qu'est-ce que Kubernetes ?

Kubernetes est une puissante plateforme d'orchestration de conteneurs open-source conçue pour gérer des applications conteneurisées sur des clusters de machines.

Initialement développé par Google, Kubernetes, communément appelé K8s, gère le déploiement, la mise à l'échelle et les opérations des conteneurs d'applications, ce qui en fait un outil essentiel pour la gestion des conteneurs à l'échelle.

