---
id: collect-261001-huawei/huawei/les-26-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-concernant-d-3
title: "Step 1: Choose a base image"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/les-26-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-concernant-docker-pour.md
source_anchor: ""
source_lines: [205, 305]
sha256: addf5de9a5e9e06acb8877c819f9c30631fd6f15c92c41ff997825bf9a34b08c
---

# Step 1: Choose a base image

Ce contrôle de santé interroge le point de terminaison de santé du conteneur toutes les 30 secondes et marque le conteneur comme non fonctionnel s'il échoue à trois tentatives consécutives. Cette surveillance proactive permet d'identifier et de résoudre rapidement les problèmes.

### 18. Que sont les images suspendues dans Docker et comment peut-on les supprimer ?

Les images suspendues dans Docker sont des couches d'images inutilisées qui ne sont plus associées à aucune balise. Ils s'accumulent souvent lorsque vous créez de nouvelles images avec le même nom et la même balise, laissant les anciens calques sans références. Ces images peuvent occuper un espace disque important, il est donc essentiel de les supprimer. Voici comment procéder :

1. Veuillez exécuter la commande `docker images -f dangling=true` pour identifier les images en suspens.
2. Ensuite, veuillez exécuter la commande `docker image prune -f` pour supprimer toutes les images en une seule fois.
3. Si vous souhaitez supprimer des images manuellement, veuillez utiliser la commande `docker rmi -f $(docker images -f dangling=true -q)` .

Ces étapes contribuent à maintenir votre système propre et à libérer efficacement de l'espace de stockage.

## Questions d'entretien sur Docker et Kubernetes

Docker et Kubernetes sont souvent utilisés conjointement, il n'est donc pas surprenant de rencontrer des questions relatives à Kubernetes lors d'un entretien d'embauche pour un poste chez Docker, en particulier si le poste est orienté DevOps. Voici quelques questions qui pourraient vous être posées :

### 19. Quelle est la principale différence entre Docker et Kubernetes ?

Docker est une plateforme de conteneurisation qui permet de créer, d'expédier et d'exécuter des conteneurs. Il se concentre sur la création et la gestion de conteneurs individuels.Kubernetes, quant à lui, est une plateforme d'orchestration conçue pour gérer plusieurs conteneurs à grande échelle. Il gère le déploiement, la mise à l'échelle, l'équilibrage de charge et l'auto-réparation à travers des clusters de nœuds.

Pour en savoir plus sur les différences entre Kubernetes et Docker, veuillez consulter l'article de blog.

### 20. Veuillez comparer Docker Swarm et Kubernetes.

Kubernetes et Docker Swarm gèrent les conteneurs, mais fonctionnent différemment :

- Kubernetes gère des configurations de conteneurs volumineuses et complexes. Ses fonctionnalités d'auto-réparation et de surveillance intégrée en font une option plus adaptée aux environnements complexes.
- Docker Swarm convient aux configurations plus petites ou moins complexes, car il n'offre pas de fonctionnalités intégrées comme Kubernetes. Nous pouvons facilement l'intégrer à des outils Docker tels que Docker CLI et Docker Compose.

### 21. Comment Kubernetes gère-t-il un grand nombre de conteneurs Docker ?

Bien que Docker soit un excellent outil pour créer et exécuter des conteneurs, la gestion d'un grand nombre d'entre eux nécessite l'utilisation de Kubernetes. Kubernetes coordonne efficacement les conteneurs en :

- Définition des limites de ressources: Il alloue le CPU, la mémoire et d'autres ressources à chaque conteneur afin d'éviter toute surconsommation.
- Planification des conteneurs: Kubernetes détermine où exécuter chaque conteneur, optimisant ainsi l'utilisation des ressources sur l'ensemble des nœuds d'un cluster.
- Mise à l'échelle automatique: En fonction de la charge de travail, il augmente ou diminue le nombre de pods (groupes d'un ou plusieurs conteneurs) afin de maintenir les performances et l'efficacité.

En automatisant ces processus, Kubernetes garantit un fonctionnement fluide, même lors de la gestion de milliers de conteneurs. Bien que des erreurs occasionnelles puissent survenir, ses capacités d'auto-réparation, telles que le redémarrage des conteneurs défaillants, minimisent les perturbations.

### 22. Qu'est-ce qu'un pod dans Kubernetes et en quoi diffère-t-il d'un conteneur ?

Un pod est la plus petite unité déployable dans Kubernetes et représente un groupe d'un ou plusieurs conteneurs qui partagent le même espace de noms réseau, le même stockage et la même configuration.Contrairement aux conteneurs individuels, les pods permettent à plusieurs conteneurs étroitement liés de fonctionner ensemble comme une seule unité (par exemple, un serveur web et un conteneur de journalisation sidecar).

*Présentation d'un nœud Kubernetes, mettant en évidence les pods et les conteneurs. Source de l'image : Kubernetes.*

### 23. Comment pouvez-vous gérer les données sensibles telles que les mots de passe dans Docker et Kubernetes ?

- Dans Docker : Nous pouvons utiliser les secrets Docker, qui chiffrent les données sensibles et les rendent accessibles uniquement aux conteneurs autorisés lors de l'exécution.
- Dans Kubernetes : Nous utilisons des objets Secrets, qui stockent des données sensibles telles que des mots de passe, des jetons et des clés API. Les secrets peuvent être montés en tant que volumes ou exposés en tant que variables d'environnement aux pods de manière sécurisée.

Exemple dans Kubernetes :

```
apiVersion: v1
kind: Secret
metadata:
  name: my-secret
type: Opaque
data:
  password: cGFzc3dvcmQ=  # Base64-encoded "password"
```
## Questions d'entretien sur Docker basées sur des scénarios

L'intervieweur pose des questions basées sur des scénarios et axées sur la résolution de problèmes afin d'évaluer votre approche face à des situations réelles. Veuillez examiner quelques questions pour vous donner une idée :

### 24. Veuillez imaginer que vous créez une image d'une API basée sur Maven. Vous avez déjà configuré le fichier Dockerfile avec les paramètres de base. Vous remarquerez que la taille de l'image est importante. Comment pourriez-vous le réduire ?

Exemple de réponse :

Pour réduire la taille d'une image Docker pour une API basée sur Maven, je suivrais les étapes suivantes :

Veuillez créer un fichier `.dockerignore` dans le répertoire du projet afin de spécifier les fichiers et dossiers qui ne doivent pas être inclus dans le contexte de construction Docker. Cela empêche l'ajout de fichiers inutiles à l'image, réduisant ainsi sa taille. Par exemple, j'ajouterais ce qui suit à `.dockerignore`:

```
.git        # Version control files
target      # Compiled code and build artifacts
.idea       # IDE configuration files
```
Optimisez le fichier Dockerfile à l'aide de builds multi-étapes. Je construirais le projet Maven en une seule étape et ne copierais que les artefacts nécessaires (par exemple, les fichiers JAR compilés) dans l'étape finale afin de conserver une image de petite taille. Exemple de fichier Dockerfile avec compilation en plusieurs étapes :

```
# Stage 1: Build the application
FROM maven:3.8.5-openjdk-11 AS build
WORKDIR /app
COPY pom.xml .
COPY src ./src
RUN mvn clean package
# Stage 2: Create a lightweight runtime image
FROM openjdk:11-jre-slim
WORKDIR /app
COPY --from=build /app/target/my-api.jar .
CMD ["java", "-jar", "my-api.jar"]
```
En ignorant les fichiers inutiles et en utilisant des compilations en plusieurs étapes, il est possible de réduire considérablement la taille de l'image tout en conservant son efficacité.

### 25. Veuillez imaginer que vous devez transférer une image de conteneur Docker vers Docker Hub à l'aide de Jenkins. Comment procéderiez-vous ?

Exemple de réponse :

Voici comment je procéderais pour transférer une image de conteneur Docker vers Docker Hub avec Jenkins :

