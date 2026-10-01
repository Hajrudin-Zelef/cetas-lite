---
id: collect-261001-rattrapage/rattrapage/comment-apprendre-docker-depuis-zero-guide-pour-les-professionnels-des-donnees-2
title: "Use an official Python runtime as a parent image"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/comment-apprendre-docker-depuis-zero-guide-pour-les-professionnels-des-donnees.md
source_anchor: ""
source_lines: [162, 328]
sha256: 41d36a02d866b1edd384d41baf7b626e11021cdb51b2f59b31e3e238e573060e
---

# Use an official Python runtime as a parent image

### Étape 3 : exécuter votre premier conteneur

Maintenant que Docker est installé, il est temps d’exécuter votre premier conteneur. Nous allons démarrer avec la simple image hello-world, parfaite pour débuter.

1. Ouvrez votre terminal (ou l’invite de commandes sous Windows) et exécutez :

`sudo docker run hello-world`
Cette commande indique à Docker de chercher l’image `hello-world` en local. Si elle est introuvable, Docker la téléchargera depuis Docker Hub puis l’exécutera.

Si Docker est correctement installé, vous verrez le message « Hello from Docker! » accompagné d’explications sur le fonctionnement du processus. Cette sortie confirme que Docker a bien récupéré l’image, créé un nouveau conteneur et exécuté le code à l’intérieur.

### Étape 4 : construire votre première image Docker

Dans cette étape, vous allez créer une image Docker personnalisée pour un mini projet de data science. Elle inclura Python et des bibliothèques courantes comme pandas, NumPy et scikit-learn.

1. Créer un nouveau répertoire :

- 
  - Commencez par créer un répertoire pour votre projet et placez-vous dedans :

```
mkdir my-data-science-app
cd my-data-science-app
```
2. Créer un Dockerfile :

Dans ce répertoire, créez un fichier nommé `Dockerfile` (sans extension) :

```
# Use an official Python runtime as a parent image
FROM python:3.8-slim
 
# Set the working directory in the container
WORKDIR /app
 
# Copy the current directory contents into the container at /app
COPY . /app
 
# Install any needed packages specified in requirements.txt
RUN pip install --no-cache-dir -r requirements.txt
 
# Make port 8888 available to the world outside this container
EXPOSE 8888
 
# Run a Jupyter notebook server when the container launches
CMD ["jupyter", "notebook", "--ip=0.0.0.0", "--port=8888", "--no-browser", "--allow-root"]
```
Ce Dockerfile effectue les actions suivantes :

- `FROM python:3.8-slim` : définit l’image de base sur Python 3.8, variante « slim », idéale pour un environnement data léger.
- `WORKDIR /app` : fixe le répertoire de travail du conteneur à /app. Les commandes suivantes s’y exécuteront.
- `COPY . /app` : copie tous les fichiers du répertoire courant de l’hôte vers`/app` dans le conteneur (scripts, fichiers de configuration, etc.).
- `RUN pip install --no-cache-dir -r requirements.txt` : installe les bibliothèques Python listées dans`requirements.txt` (pandas, NumPy, scikit-learn, ...).
- `EXPOSE 8888` : expose le port 8888 à l’hôte, port par défaut des notebooks Jupyter.
- `CMD ["jupyter", "notebook", "--ip=0.0.0.0", "--port=8888", "--no-browser", "--allow-root"]` : lance un serveur Jupyter Notebook au démarrage du conteneur, accessible depuis votre navigateur.

3. Créer le fichier `requirements.txt` :

Dans le même répertoire, créez un fichier `requirements.txt` listant les paquets Python requis :

```
pandas
numpy
scikit-learn
matplotlib
```
Ce fichier énumère les bibliothèques Python qui seront installées dans votre conteneur Docker.

4. Construire l’image Docker :

- Avec le Dockerfile et `requirements.txt` prêts, construisez l’image avec :

`sudo docker build -t my-data-science-app .`
Cette commande fait :

- `docker build` : lance le processus de construction d’une image.
- `-t my-data-science-app` : étiquette l’image « my-data-science-app ».
- `.` : le point final désigne le contexte de construction (le répertoire courant).

5. Exécuter l’image Docker :

Une fois la construction terminée, exécutez votre nouvelle image en tant que conteneur :

`sudo docker run -p 8888:8888 my-data-science-app`
Détails de la commande :

- `docker run` : lance un conteneur à partir de l’image spécifiée.
- `-p 8888:8888` : mappe le port 8888 de votre machine locale vers le port 8888 du conteneur pour accéder au serveur Jupyter via`localhost:8888` .
- `my-data-science-app` : image à exécuter.

### Étape 5 : utiliser Docker Compose

Docker Compose est indispensable pour gérer des applications multi-conteneurs. Dans un projet de data science, vous pouvez avoir des conteneurs distincts pour un notebook Jupyter, une base de données et un outil de visualisation.

Avec Docker Compose, vous définissez ces services dans un unique fichier YAML, ce qui permet de lancer l’ensemble de votre environnement en une seule commande.

1. Créer un fichier `docker-compose.yml` :

- Dans votre répertoire de projet, créez un fichier `docker-compose.yml` et ajoutez les lignes suivantes :

```
version: '3.8'
services:
  jupyter:
        image: jupyter/scipy-notebook:latest
        volumes:
        - ./notebooks:/home/joelwembo/work
        ports:
        - "8888:8888"
        environment:
        - JUPYTER_ENABLE_LAB=yes
 
  postgres:
        image: postgres:13-alpine
        environment:
        POSTGRES_USER: myuser
        POSTGRES_PASSWORD: mypassword
        POSTGRES_DB: mydatabase
        volumes:
        - postgres_data:/var/lib/postgresql/data
        ports:
        - "5432:5432"
 
  redis:
        image: redis:alpine
        ports:
        - "6379:6379"
 
volumes:
  postgres_data:
 
```
Voici ce que fait ce fichier Docker Compose :

- `version: '3.8'` : version du format de fichier Docker Compose, largement utilisée.
- `services:` : section qui définit les différents conteneurs (services) de votre application.
- `jupyter:` : service pour l’environnement Jupyter, essentiel à l’analyse interactive.
- `image: jupyter/scipy-notebook:latest` : image Docker pour Jupyter incluant NumPy, pandas, matplotlib, etc.
- `volumes: - ./notebooks:/home/joelwembo/work` : monte le dossier`notebooks` de l’hôte dans`/home/joelwembo/work` du conteneur pour persister et partager vos notebooks. Remarque : il s’agit de mon chemin personnel, le vôtre sera différent.
- `ports: - "8888:8888"` : mappe le port 8888 de l’hôte sur celui du conteneur (port Jupyter par défaut).
- `environment: - JUPYTER_ENABLE_LAB=yes` : active JupyterLab, interface plus puissante pour travailler avec Jupyter.
- `postgres:` : service PostgreSQL, souvent utilisé pour stocker et gérer des jeux de données volumineux.
- `image: postgres:13-alpine` : image PostgreSQL 13 basée sur Alpine Linux, légère.
- `environment:` : variables d’environnement pour configurer l’utilisateur, le mot de passe et le nom de la base.
- `volumes: - postgres_data:/var/lib/postgresql/data` : volume Docker pour persister les données de la base au-delà du cycle de vie du conteneur.
- `ports: - "5432:5432"` : expose le port PostgreSQL par défaut (5432).
- `redis:` : service Redis, utilisé comme magasin de données en mémoire ou cache rapide.
- `image: redis:alpine` : image Redis basée sur Alpine, légère.
- `ports: - "6379:6379"` : expose le port Redis par défaut (6379).
- `volumes:` : section qui définit les volumes nommés partagés entre services ou persistants entre redémarrages.
- `postgres_data:` : crée un volume nommé pour les données PostgreSQL.

Pour lancer votre application multi-conteneurs, exécutez simplement :

`sudo docker compose up`
Cette commande construit les images (si nécessaire) et démarre les conteneurs définis dans votre `docker-compose.yml`. Elle met en route l’ensemble des services et leur permet de communiquer comme prévu.

Vous venez de déployer votre première application Docker ! Plutôt motivant, n’est-ce pas ? Passons maintenant à un plan pour votre formation continue.

## Exemple de plan d’apprentissage Docker

Voici un plan semaine par semaine que vous pouvez suivre et adapter à vos besoins.

Si vous préférez une voie plus balisée, le parcours Containerization and Virtualization with Docker and Kubernetes est fait pour vous. Il contient quatre cours essentiels.

### Semaine 1 : se familiariser avec les bases de Docker

