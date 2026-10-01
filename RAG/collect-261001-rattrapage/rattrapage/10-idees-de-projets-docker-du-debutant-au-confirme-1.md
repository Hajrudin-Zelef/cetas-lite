---
id: collect-261001-rattrapage/rattrapage/10-idees-de-projets-docker-du-debutant-au-confirme-1
title: "Stage 1: Build"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/10-idees-de-projets-docker-du-debutant-au-confirme.md
source_anchor: ""
source_lines: [1, 192]
sha256: e3a9b85a4d33fbe317f2233c5e36773e7959cbce7fb5d83ffaa3d35aaa68ec42
---

# Stage 1: Build

Cours

L'expérience pratique est essentielle pour maîtriser Docker. Docker est un outil essentiel au développement de logiciels modernes et à la science des données, qui vous permet de créer, de déployer et de gérer des applications dans des conteneurs.

Dans cet article, je vous propose des exemples de projets Docker de niveau débutant, intermédiaire et avancé, axés sur les constructions en plusieurs étapes, l'optimisation des images Docker et l'application de Docker dans le domaine de la science des données. Ces projets sont conçus pour approfondir votre compréhension de Docker et améliorer vos compétences pratiques.

## Démarrer avec les projets Docker

Avant de vous lancer dans les projets, assurez-vous que Docker est installé sur votre machine. En fonction de votre système d'exploitation (Windows, macOS, Linux), vous pouvez télécharger Docker depuis le site officiel de Docker.

Vous devrez également avoir des connaissances de base dans les domaines suivants

- Dockerfiles (pour définir ce qui se trouve à l'intérieur de vos conteneurs)
- Docker Compose (pour les applications multi-conteneurs)
- Commandes CLI de base telles que `docker build` ,`docker run` ,`docker-compose up` , etc.

Si vous avez besoin de rafraîchir vos connaissances sur les concepts ci-dessus, consultez les cours Introduction à Docker ou Concepts de conteneurisation et de virtualisation.

Commençons !

## Projets Docker pour les débutants

Lorsque vous débutez avec Docker, il est important de choisir des projets qui correspondent à votre niveau de compétence tout en vous incitant à apprendre de nouveaux concepts. Voici quelques idées de projets pour vous aider à démarrer :

### Projet 1 : Mise en place d'un serveur web simple

Dans ce projet, vous allez créer un conteneur Docker qui exécute un serveur web de base à l'aide de Nginx. Nginx est l'un des serveurs web open-source les plus populaires pour le reverse proxying, l'équilibrage de charge, etc. À la fin de ce projet, vous aurez appris à créer et à exécuter des conteneurs avec Docker et à exposer des ports pour que l'application soit accessible depuis votre machine locale.

Niveau de difficulté : Débutant

Technologies utilisées : Docker, Nginx

#### Instructions étape par étape

- Installez Docker : Assurez-vous que Docker est installé sur votre système.
- Créez le répertoire du projet : Créez un nouveau dossier et un fichier `index.html` à l'intérieur qui sera servi par Nginx.
- Écrivez le fichier Docker : Un Dockerfile est un script qui définit l'environnement du conteneur. Il indique à Docker l'image de base à utiliser, les fichiers à inclure et les ports à exposer :

```
FROM nginx:alpine
COPY ./index.html /usr/share/nginx/html
EXPOSE 80
```
- Construisez l'image Docker : Naviguez dans le dossier de votre projet et construisez l'image à l'aide de :

`docker build -t my-nginx-app .`
- Exécutez le conteneur : Démarrez le conteneur et mappez le port 80 du conteneur au port 8080 de votre machine :

`docker run -d -p 8080:80 my-nginx-app`
- Accédez au serveur web : Ouvrez votre navigateur et naviguez vers http://localhost:8080 pour voir la page que vous avez créée.

### Projet 2 : Dockeriser un script Python

Ce projet consiste à conteneuriser un simple script Python qui traite les données d'un fichier CSV à l'aide de la bibliothèque pandas. L'objectif est d'apprendre à gérer les dépendances et à exécuter des scripts Python à l'intérieur de conteneurs Docker, ce qui rend le script portable et exécutable dans n'importe quel environnement.

Niveau de difficulté : Débutant

Technologies utilisées : Docker, Python, pandas

#### Instructions étape par étape

- Écrivez le script Python : Créez un script nommé `process_data.py` qui lit et traite un fichier CSV. Voici un exemple de script :

```
import pandas as pd
df = pd.read_csv('data.csv')
print(df.describe())
```
- Créez un fichier `requirements.txt` : Ce fichier répertorie les bibliothèques Python dont le script a besoin. Dans ce cas, nous n'avons besoin que de`pandas` :

`pandas`
- Écrivez le fichier Docker : Ce fichier définit l'environnement du script Python :

```
FROM python:3.9-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
CMD ["python", "process_data.py"]
```
- Construisez l'image Docker :

`docker build -t python-script .`
- Exécutez le conteneur :

`docker run -v $(pwd)/data:/app/data python-script`
### Projet 3 : Construire une application multi-conteneurs simple

Ce projet vous aidera à vous familiariser avec Docker Compose en construisant une application multi-conteneurs. Vous créerez une application web simple en utilisant Flask comme interface et MySQL comme base de données. Docker Compose vous permet de gérer plusieurs conteneurs qui fonctionnent ensemble.

Niveau de difficulté : Débutant

Technologies utilisées : Docker, Docker Compose, Flask, MySQL

#### Instructions étape par étape

- Écrivez l'application Flask : Créez une application Flask simple qui se connecte à une base de données MySQL et affiche un message. En voici un exemple :

```
from flask import Flask
import mysql.connector
 
app = Flask(__name__)
 
def get_db_connection():
 	connection = mysql.connector.connect(
	 host="db",
	 user="root",
	 password="example",
	 database="test_db"
 	)
 	return connection
 
@app.route('/')
def hello_world():
 	connection = get_db_connection()
 	cursor = connection.cursor()
 	cursor.execute("SELECT 'Hello, Docker!'")
 	result = cursor.fetchone()
 	connection.close()
 	return str(result[0])
 
if __name__ == "__main__":
 	app.run(host='0.0.0.0')
```
- Créez le fichier `docker-compose.yml` : Docker Compose définit et exécute des applications Docker multi-conteneurs. Dans ce fichier, vous définirez l'application Flask et les services de base de données MySQL :

```
version: '3'
services:
  db:
    image: mysql:5.7
    environment:
      MYSQL_ROOT_PASSWORD: example
      MYSQL_DATABASE: test_db
    ports:
      - "3306:3306"
    volumes:
      - db_data:/var/lib/mysql
  web:
    build: .
    ports:
      - "5000:5000"
    depends_on:
      - db
    environment:
      FLASK_ENV: development
    volumes:
      - .:/app
volumes:
  db_data:
```
- Écrivez le fichier Docker pour Flask : Cela créera l'image Docker pour l'application Flask :

```
FROM python:3.9-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
CMD ["python", "app.py"]
```
- Construisez et exécutez les conteneurs : Utilisez Docker Compose pour afficher l'ensemble de l'application :

`docker-compose up --build`
- **Accédez à l'application Flask :** Allez sur http://localhost:5000 dans votre navigateur.

## Devenez ingénieur en données

## Projets Docker de niveau intermédiaire

Les projets suivants sont destinés à ceux qui ont une solide compréhension des bases de Docker. Ils introduiront des concepts plus complexes, tels que les constructions en plusieurs étapes et les techniques d'optimisation.

### Projet 4 : Construction en plusieurs étapes d'une application Node.js

Les constructions en plusieurs étapes permettent de réduire la taille des images Docker en séparant les environnements de construction et d'exécution. Dans ce projet, vous allez conteneuriser une application Node.js en utilisant des constructions en plusieurs étapes.

Niveau de difficulté : Intermédiaire

Technologies utilisées : Docker, Node.js, Nginx

#### Instructions étape par étape

- Créez une application Node.js simple : Ecrivez un serveur Node.js basique qui renvoie un message simple. En voici un exemple :

```
const express = require('express');
const app = express();
 
app.get('/', (req, res) => res.send('Hello from Node.js'));
 
