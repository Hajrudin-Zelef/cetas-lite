---
id: collect-240926-datacamp/datacamp/10-idees-de-projets-docker-du-debutant-au-confirme-1
title: "Stage 1: Build"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean_en/datacamp/10-idees-de-projets-docker-du-debutant-au-confirme.md
source_anchor: ""
source_lines: [1, 225]
sha256: 31e335e966a1f4a3574417491773d2519422a4a68c9c02f3c4fcb868a59c640a
---

# Stage 1: Build

<!-- source: https://www.datacamp.com/fr/blog/docker-projects -->

Course

Practical experience is essential for mastering Docker. Docker is an essential tool for modern software development and data science, allowing you to build, deploy, and manage applications in containers.

In this article, I offer you examples of Docker projects at beginner, intermediate, and advanced levels, focused on multi-stage builds, optimizing Docker images, and applying Docker in the field of data science. These projects are designed to deepen your understanding of Docker and improve your practical skills.

## Getting Started with Docker Projects

Before diving into the projects, make sure Docker is installed on your machine. Depending on your operating system (Windows, macOS, Linux), you can download Docker from the official Docker website.

You will also need basic knowledge in the following areas

- Dockerfiles (to define what is inside your containers)
- Docker Compose (for multi-container applications)
- Basic CLI commands such as `docker build`, `docker run`, `docker-compose up`, etc.

If you need to refresh your knowledge of the above concepts, check out the Introduction to Docker or Containerization and Virtualization Concepts courses.

Let's get started!

## Docker Projects for Beginners

When you're new to Docker, it's important to choose projects that match your skill level while pushing you to learn new concepts. Here are some project ideas to help you get started:

### Project 1: Setting Up a Simple Web Server

In this project, you will create a Docker container that runs a basic web server using Nginx. Nginx is one of the most popular open-source web servers for reverse proxying, load balancing, etc. By the end of this project, you will have learned how to create and run containers with Docker and expose ports so that the application is accessible from your local machine.

Difficulty level: Beginner

Technologies used: Docker, Nginx

#### Step-by-step instructions

- Install Docker: Make sure Docker is installed on your system.
- Create the project directory: Create a new folder and an `index.html` file inside it that will be served by Nginx.
- Write the Docker file: A Dockerfile is a script that defines the container's environment. It tells Docker which base image to use, which files to include, and which ports to expose:

```
FROM nginx:alpine
COPY ./index.html /usr/share/nginx/html
EXPOSE 80
```
- Build the Docker image: Navigate to your project folder and build the image using:

`docker build -t my-nginx-app .`
- Run the container: Start the container and map port 80 of the container to port 8080 of your machine:

`docker run -d -p 8080:80 my-nginx-app`
- Access the web server: Open your browser and navigate to http://localhost:8080 to see the page you created.

### Project 2: Dockerizing a Python Script

This project involves containerizing a simple Python script that processes data from a CSV file using the pandas library. The goal is to learn how to manage dependencies and run Python scripts inside Docker containers, making the script portable and executable in any environment.

Difficulty level: Beginner

Technologies used: Docker, Python, pandas

#### Step-by-step instructions

- Write the Python script: Create a script named `process_data.py` that reads and processes a CSV file. Here is an example script:

```
import pandas as pd
df = pd.read_csv('data.csv')
print(df.describe())
```
- Create a `requirements.txt` file: This file lists the Python libraries that the script needs. In this case, we only need `pandas`:

`pandas`
- Write the Docker file: This file defines the Python script's environment:

```
FROM python:3.9-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
CMD ["python", "process_data.py"]
```
- Build the Docker image:

`docker build -t python-script .`
- Run the container:

`docker run -v $(pwd)/data:/app/data python-script`
### Project 3: Building a Simple Multi-Container Application

This project will help you become familiar with Docker Compose by building a multi-container application. You will create a simple web application using Flask as the interface and MySQL as the database. Docker Compose allows you to manage multiple containers that work together.

Difficulty level: Beginner

Technologies used: Docker, Docker Compose, Flask, MySQL

#### Step-by-step instructions

- Write the Flask application: Create a simple Flask application that connects to a MySQL database and displays a message. Here is an example:

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
- Create the `docker-compose.yml` file: Docker Compose defines and runs multi-container Docker applications. In this file, you will define the Flask application and the MySQL database services:

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
- Write the Docker file for Flask: This will create the Docker image for the Flask application:

```
FROM python:3.9-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
CMD ["python", "app.py"]
```
- Build and run the containers: Use Docker Compose to bring up the entire application:

`docker-compose up --build`
- **Access the Flask application:** Go to http://localhost:5000 in your browser.

## Become a data engineer

## Intermediate-level Docker projects

The following projects are intended for those who have a solid understanding of the basics of Docker. They will introduce more complex concepts, such as multi-stage builds and optimization techniques.

### Project 4: Multi-stage build of a Node.js application

Multi-stage builds make it possible to reduce the size of Docker images by separating the build and runtime environments. In this project, you will containerize a Node.js application using multi-stage builds.

Difficulty level: Intermediate

Technologies used: Docker, Node.js, Nginx

#### Step-by-step instructions

- Create a simple Node.js application: Write a basic Node.js server that returns a simple message. Here is an example:

```
const express = require('express');
const app = express();
 
app.get('/', (req, res) => res.send('Hello from Node.js'));
 
app.listen(3000, () => console.log('Server running on port 3000'));
```
- Write the Docker file with a multi-stage build: The first stage consists of building the application and the second of running it with a lighter base image.

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
- Build the image:

`docker build -t node-multi-stage .`
- Run the container:

`docker run -p 3000:3000 node-multi-stage`
### Project 5: Dockerize a machine learning model with TensorFlow

This project will involve containerizing a machine learning model using TensorFlow. The goal is to create a portable environment in which you can run TensorFlow models on different systems without worrying about the underlying configuration.

