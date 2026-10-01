---
id: collect-261001-huawei/huawei/les-26-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-concernant-d-1
title: "Step 1: Choose a base image"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/les-26-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-concernant-docker-pour.md
source_anchor: ""
source_lines: [1, 121]
sha256: bf929196ce6583ec4acf62e745134b6de0e893006c2339639a95f18c9d83c817
---

# Step 1: Choose a base image

Cursus

Docker est devenu l'outil de conteneurisation le plus populaire dans le développement logiciel moderne, en particulier dans les workflows DevOps et CI/CD. Il simplifie le déploiement et la gestion des applications grâce à des conteneurs, ce qui permet de fournir des logiciels rapidement et de manière cohérente.

Sa scalabilité et sa flexibilité font de Docker un outil indispensable pour les postes liés aux données, tels que l'ingénierie des données, le MLOps et même la science des données. C'est pourquoi j'ai compilé les questions fréquemment posées lors des entretiens d'embauche liés à Docker, couvrant les concepts fondamentaux et les scénarios concrets.

## Devenez ingénieur en données

## Qu'est-ce que Docker ?

Docker est une plateforme de conteneurs que les développeurs utilisent pour regrouper des applications avec toutes leurs dépendances afin qu'elles puissent fonctionner correctement dans différents environnements.

Bien que les conteneurs partagent le même noyau de système d'exploitation, chacun fonctionne dans son propre environnement isolé. Cette configuration minimise les problèmes de compatibilité, réduit les retards et améliore la communication entre les équipes de développement, de test et d'exploitation.

*Logo Docker. Source de l'image*

En 2023, Docker était leader sur le marché de la conteneurisation avec plus de 32 % de parts de marché. Cela souligne son importance dans le développement logiciel moderne. C'est pourquoi vous pouvez vous attendre à ce que les recruteurs évaluent votre expertise en matière de Docker lors d'entretiens d'embauche liés aux données.

## Questions d'entretien de base sur Docker

Tout d'abord, familiarisez-vous avec certains concepts fondamentaux de Docker. Ces questions fondamentales vous aideront à approfondir votre compréhension et à vous préparer pour la phase initiale de l'entretien.

### 1. Qu'est-ce qu'une image Docker ?

Une image Docker est comparable à un plan qui permet de créer des conteneurs. Il contient tous les éléments nécessaires à un développeur pour exécuter une application, notamment :

- Code
- Bibliothèques
- Paramètres

Lorsque vous utilisez une image Docker, Docker la transforme en conteneur, qui est un environnement entièrement isolé. C'est là que l'application fonctionne de manière autonome.

### 2. Qu'est-ce qu'un hôte Docker ?

Un hôte Docker est le système sur lequel nous installons Docker. Il sert d'environnement principal pour l'exécution et la gestion des conteneurs Docker. Nous pouvons configurer un hôte Docker sur un appareil local ou dans un environnement virtuel ou cloud.

### 3. En quoi un client Docker diffère-t-il d'un démon Docker ? Pourriez-vous nous donner un exemple ?

Le client Docker et le démon Docker fonctionnent en parallèle, mais ont des rôles distincts. Le client Docker est l'outil qui envoie les commandes et le démon Docker est le moteur qui exécute ces commandes.

Par exemple, si nous saisissons la commande `docker run` pour démarrer un conteneur, le client prendra la requête et l'enverra au démon Docker. Le démon Docker se chargera ensuite du travail réel en démarrant le conteneur. 

### 4. Pourriez-vous expliquer ce qu'est le réseau Docker et quelles commandes permettent de créer un pont et un réseau superposé ?

Le réseau Docker permet aux conteneurs de se connecter et de communiquer avec d'autres conteneurs et hôtes. La commande ` `docker network create` ` nous permet de configurer des réseaux définis par l'utilisateur.  

- **Réseau de ponts** : Crée un réseau local pour la communication entre les conteneurs sur le même hôte Docker.
- Commande : `docker network create -d bridge my-bridge-network`
- Cela configure un réseau pont appelé `my-bridge-network` pour les conteneurs sur le même hôte.
- **Réseau superposé** : Permet la communication entre les conteneurs sur plusieurs hôtes Docker, souvent utilisée dans une configuration Swarm.
- Commande : `docker network create --scope=swarm --attachable -d overlay my-multihost-network`
- Cela crée un réseau superposé connectable appelé « `my-multihost-network` » pour les conteneurs fonctionnant sur différents hôtes dans un Docker Swarm.

### 5. Veuillez expliquer le fonctionnement du réseau pont Docker.

Le pont réseau est la configuration par défaut utilisée par Docker pour connecter les conteneurs. Si nous ne spécifions pas de réseau, Docker le relie au réseau bridge. Ce pont relie tous les conteneurs sur le même hôte Docker. Chaque conteneur dispose d'une adresse IP unique, ce qui permet aux conteneurs de communiquer directement entre eux.

## Questions d'entretien intermédiaire sur Docker

Ces questions sont posées afin d'évaluer votre connaissance des concepts Docker de niveau intermédiaire.

### 6. Qu'est-ce qu'un fichier Dockerfile ? Veuillez expliquer comment vous l'écririez.

Un fichier Dockerfile est un script qui définit les instructions pour créer une image Docker. Chaque commande du fichier Dockerfile configure une partie spécifique de l'environnement. Lorsque nous exécutons ces commandes, Docker construit une image couche par couche. Voici comment nous pouvons l'écrire :

1. Tout d'abord, veuillez sélectionner une image de base. Il contient les outils essentiels pour l'application.
2. Ensuite, veuillez définir un répertoire de travail à l'intérieur du conteneur. C'est là que les fichiers de l'application seront stockés et exécutés.
3. Dans la troisième étape, veuillez utiliser la commande `COPY . .` pour copier tous les fichiers du projet dans le répertoire de travail du conteneur.
4. Veuillez utiliser la commande « `RUN` » pour installer les dépendances.
5. Veuillez utiliser la commande ` `EXPOSE` ` pour indiquer le port sur lequel votre application est exécutée.
6. Veuillez maintenant définir la commande que Docker doit exécuter lorsqu'il démarre le conteneur.

Voici un exemple simple de fichier Dockerfile pour une application web Python :

```
# Step 1: Choose a base image
FROM python:3.9-slim
# Step 2: Specify the working directory
WORKDIR /app
# Step 3: Copy project files into the container
COPY . .
# Step 4: Install dependencies
RUN pip install -r requirements.txt
# Step 5: Expose the port the app runs on
EXPOSE 5000
# Step 6: Define the default command
CMD ["python", "app.py"]
```
À l'aide du fichier Dockerfile ci-dessus, vous pouvez créer une image avec la commande ` `docker build -t my-python-app .` ` et exécuter un conteneur avec la commande ` `docker run -p 5000:5000 my-python-app``.

### 7. Qu'est-ce que Docker Compose et en quoi diffère-t-il de Dockerfile ?

Docker Compose est un outil permettant de définir et de gérer des applications Docker multi-conteneurs à l'aide d'un fichier YAML (`docker-compose.yml`). Il nous permet de configurer les services, les réseaux et les volumes dans un seul fichier, ce qui facilite la gestion des applications complexes.

Différences par rapport au fichier Dockerfile :

- Un fichier Dockerfile est utilisé pour créer une image Docker unique en définissant ses couches et ses dépendances.
- Docker Compose est utilisé pour exécuter et orchestrer plusieurs conteneurs qui peuvent dépendre les uns des autres (par exemple, un conteneur d'application web et un conteneur de base de données).

Par exemple, un fichier `docker-compose.yml` pourrait se présenter comme suit :

```
version: '3.9'
services:
  web:
    build: .
    ports:
      - "5000:5000"
    depends_on:
      - db
  db:
    image: postgres
    volumes:
      - db-data:/var/lib/postgresql/data
volumes:
  db-data:
```
Ce fichier définit deux services, `web` et `db`, avec des configurations réseau et de volume.

### 8. Pourquoi utilisons-nous des volumes dans Docker ?

