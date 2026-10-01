---
id: collect-261001-rattrapage/rattrapage/10-idees-de-projets-docker-du-debutant-au-confirme-2
title: "Stage 1: Build"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-rattrapage/10-idees-de-projets-docker-du-debutant-au-confirme.md
source_anchor: ""
source_lines: [193, 424]
sha256: 01fd37eef7d677df124e775a6d9b6b9cf0d2b526b043b8341df0e7dfe1e99ecc
---

# Stage 1: Build

app.listen(3000, () => console.log('Server running on port 3000'));
```
- Écrire le fichier Docker avec une construction en plusieurs étapes : La première étape consiste à créer l'application et la seconde à l'exécuter avec une image de base plus légère.

```
# Stage 1: Build
FROM node:14 as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
# Add the following line if there's a build step for the app
# RUN npm run build
# Stage 2: Run
FROM node:14-slim
WORKDIR /app
COPY --from=build-stage /app .
EXPOSE 3000
ENV NODE_ENV=production
CMD ["node", "server.js"]
```
- Construisez l'image :

`docker build -t node-multi-stage .`
- Exécutez le conteneur :

`docker run -p 3000:3000 node-multi-stage`
### Projet 5 : Dockeriser un modèle d'apprentissage automatique avec TensorFlow

Ce projet consistera à conteneuriser un modèle d'apprentissage automatique à l'aide de TensorFlow. L'objectif est de créer un environnement portable dans lequel vous pouvez exécuter des modèles TensorFlow sur différents systèmes sans vous soucier de la configuration sous-jacente.

Niveau de difficulté : Intermédiaire

Technologies utilisées : Docker, TensorFlow, Python

#### Instructions étape par étape

- Installez TensorFlow dans un script Python : Créez un script Python `model.py` qui charge et exécute un modèle TensorFlow pré-entraîné :

```
import tensorflow as tf
model = tf.keras.applications.MobileNetV2(weights='imagenet')
print("Model loaded successfully")
```
- Écrivez le fichier Docker : Définissez l'environnement de TensorFlow dans Docker :

```
FROM tensorflow/tensorflow:latest
WORKDIR /app
COPY . .
CMD ["python", "model.py"]
```
- Construisez l'image :

`docker build -t tensorflow-model .`
- Exécutez le conteneur :

`docker run tensorflow-model`
### Projet 6 : Créer un environnement de science des données avec Jupyter et Docker

Ce projet se concentre sur la création d'un environnement de science des données reproductible en utilisant Docker et les carnets Jupyter. L'environnement comprendra des bibliothèques Python populaires telles que pandas, NumPy et scikit-learn.

Niveau de difficulté : Intermédiaire

Technologies utilisées : Docker, Jupyter, Python, scikit-learn.

#### Instructions étape par étape

- Créez le fichier `docker-compose.yml` : Définissez le service Jupyter Notebook et les bibliothèques nécessaires. En voici un exemple :

```
version: '3'
services:
  jupyter:
    	image: jupyter/scipy-notebook
    	ports:
    	- "8888:8888"
    	volumes:
    	- ./notebooks:/home/joelwembo/work
```
- Démarrez le Jupyter Notebook : Utilisez Docker Compose pour démarrer le Jupyter Notebook.

`docker-compose up`
- Accédez au carnet de notes Jupyter : Ouvrez votre navigateur et allez sur http://localhost:8888.

## Projets Docker de niveau avancé

Ces projets de niveau avancé se concentreront sur des applications du monde réel et des concepts Docker avancés, tels que les pipelines d'apprentissage profond et les pipelines de données automatisés.

### Projet 7 : Réduire la taille d'une image Docker pour une application Python.

Dans ce projet, vous optimiserez une image Docker pour une application Python en utilisant des images de base minimales comme Alpine Linux et en mettant en œuvre des constructions en plusieurs étapes pour que la taille de l'image soit la plus petite possible.

Niveau de difficulté : Avancé

Technologies utilisées : Docker, Python, Alpine Linux

#### Instructions étape par étape

- Écrivez le script Python : Créez un script qui analyse les données à l'aide de pandas. Voici un exemple de script :

```
import pandas as pd
df = pd.read_csv('data.csv')
print(df.head())
```
- Optimisez le fichier Docker : Utilisez les constructions en plusieurs étapes et Alpine Linux pour créer une image légère.

```
# Stage 1: Build stage
FROM python:3.9-alpine as build-stage
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY script.py .
# Stage 2: Run stage
FROM python:3.9-alpine
WORKDIR /app
COPY --from=build-stage /app/script.py .
CMD ["python", "script.py"]
```
- Construisez l'image :

`docker build -t optimized-python-app .`
### Projet 8 : Dockeriser un pipeline d'apprentissage profond avec PyTorch

Ce projet consiste à conteneuriser un pipeline d'apprentissage profond en utilisant PyTorch. L'accent est mis sur l'optimisation du fichier Docker en termes de performance et de taille, ce qui facilite l'exécution de modèles d'apprentissage profond dans différents environnements.

Niveau de difficulté : Avancé

Technologies utilisées : Docker, PyTorch, Python

#### Instructions étape par étape

- Installez PyTorch dans un script Python : Créez un script qui charge un modèle PyTorch pré-entraîné et effectue l'inférence. En voici un exemple :

```
import torch
model = torch.hub.load('pytorch/vision', 'resnet18', pretrained=True)
print("Model loaded successfully")
```
- Écrivez le fichier Docker : Définissez l'environnement de PyTorch :

```
FROM pytorch/pytorch:1.9.0-cuda11.1-cudnn8-runtime
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY model.py .
CMD ["python", "model.py"]
```
- Construisez l'image :

`docker build -t pytorch-model .`
- Exécutez le conteneur :

`docker run pytorch-model` ### Projet 9 : Automatisation des pipelines de données avec Apache Airflow et Docker.

Dans ce projet, vous allez configurer et conteneuriser un environnement Apache Airflow pour automatiser les pipelines de données. Apache Airflow est un outil populaire d'orchestration de flux de travail complexes largement utilisé dans l'ingénierie des données.

Niveau de difficulté : Avancé

Technologies utilisées : Docker, Apache Airflow, Python, PostgreSQL

#### Instructions étape par étape

- Créez le fichier `docker-compose.yml` : Définissez les services Airflow et la base de données PostgreSQL :

```
version: '3'
services:
  postgres:
    image: postgres:latest
    environment:
      POSTGRES_USER: airflow
      POSTGRES_PASSWORD: airflow
      POSTGRES_DB: airflow
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
  webserver:
    image: apache/airflow:latest
    environment:
      AIRFLOW__CORE__SQL_ALCHEMY_CONN: postgresql+psycopg2://airflow:airflow@postgres/airflow
      AIRFLOW__CORE__EXECUTOR: LocalExecutor
    depends_on:
      - postgres
    ports:
      - "8080:8080"
    volumes:
      - ./dags:/opt/airflow/dags
    command: ["webserver"]
  scheduler:
    image: apache/airflow:latest
    environment:
      AIRFLOW__CORE__SQL_ALCHEMY_CONN: postgresql+psycopg2://airflow:airflow@postgres/airflow
      AIRFLOW__CORE__EXECUTOR: LocalExecutor
    depends_on:
      - postgres
      - webserver
    volumes:
      - ./dags:/opt/airflow/dags
    command: ["scheduler"]
volumes:
  postgres_data:
```
- Démarrez l'environnement Airflow : Utilisez Docker Compose pour faire apparaître l'environnement Airflow :

`docker-compose up`
- Accédez à l'interface utilisateur Airflow : Ouvrez votre navigateur et allez sur http://localhost:8080.

### Projet 10 : Déployer une API de science des données avec FastAPI et Docker

Créez et déployez une API de science des données à l'aide de FastAPI. Vous conteneuriserez l'API à l'aide de Docker et vous vous attacherez à l'optimiser pour les environnements de production.

Niveau de difficulté : Avancé

Technologies utilisées : Docker, FastAPI, Python, scikit-learn

#### Instructions étape par étape

- Écrivez l'application FastAPI : Créez une API simple qui utilise un modèle d'apprentissage automatique pour les prédictions. En voici un exemple :

```
from fastapi import FastAPI
import pickle
 
app = FastAPI()
with open("model.pkl", "rb") as f:
   model = pickle.load(f)
 
