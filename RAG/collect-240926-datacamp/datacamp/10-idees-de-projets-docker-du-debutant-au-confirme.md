---
id: collect-240926-datacamp/datacamp/10-idees-de-projets-docker-du-debutant-au-confirme
title: "Stage 1: Build"
domain: datacamp
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["apache", "aws", "inference"]
source: docs/RAG/clean_en/datacamp/10-idees-de-projets-docker-du-debutant-au-confirme.md
source_anchor: ""
source_lines: [1, 501]
sha256: 7f0ec6276612291ff2aebc00f31f6444e4d5a7aba4ef77827a8211774b349b75
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

Difficulty level: Intermediate

Technologies used: Docker, TensorFlow, Python

#### Step-by-step instructions

- Install TensorFlow in a Python script: Create a Python script `model.py` that loads and runs a pre-trained TensorFlow model:

```
import tensorflow as tf
model = tf.keras.applications.MobileNetV2(weights='imagenet')
print("Model loaded successfully")
```
- Write the Docker file: Define the TensorFlow environment in Docker:

```
FROM tensorflow/tensorflow:latest
WORKDIR /app
COPY . .
CMD ["python", "model.py"]
```
- Build the image:

`docker build -t tensorflow-model .`
- Run the container:

`docker run tensorflow-model`
### Project 6: Create a data science environment with Jupyter and Docker

This project focuses on creating a reproducible data science environment using Docker and Jupyter notebooks. The environment will include popular Python libraries such as pandas, NumPy, and scikit-learn.

Difficulty level: Intermediate

Technologies used: Docker, Jupyter, Python, scikit-learn.

#### Step-by-step instructions

- Create the `docker-compose.yml` file: Define the Jupyter Notebook service and the necessary libraries. Here is an example:

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
- Start the Jupyter Notebook: Use Docker Compose to start the Jupyter Notebook.

`docker-compose up`
- Access the Jupyter notebook: Open your browser and go to http://localhost:8888.

## Advanced-level Docker projects

These advanced-level projects will focus on real-world applications and advanced Docker concepts, such as deep learning pipelines and automated data pipelines.

### Project 7: Reduce the size of a Docker image for a Python application.

In this project, you will optimize a Docker image for a Python application by using minimal base images such as Alpine Linux and implementing multi-stage builds so that the image size is as small as possible.

Difficulty level: Advanced

Technologies used: Docker, Python, Alpine Linux

#### Step-by-step instructions

- Write the Python script: Create a script that analyzes data using pandas. Here is an example script:

```
import pandas as pd
df = pd.read_csv('data.csv')
print(df.head())
```
- Optimize the Docker file: Use multi-stage builds and Alpine Linux to create a lightweight image.

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
- Build the image:

`docker build -t optimized-python-app .`
### Project 8: Dockerizing a Deep Learning Pipeline with PyTorch

This project involves containerizing a deep learning pipeline using PyTorch. The focus is on optimizing the Docker file in terms of performance and size, which makes it easier to run deep learning models in different environments.

Difficulty level: Advanced

Technologies used: Docker, PyTorch, Python

#### Step-by-step instructions

- Install PyTorch in a Python script: Create a script that loads a pre-trained PyTorch model and performs inference. Here is an example:

```
import torch
model = torch.hub.load('pytorch/vision', 'resnet18', pretrained=True)
print("Model loaded successfully")
```
- Write the Docker file: Define the PyTorch environment:

```
FROM pytorch/pytorch:1.9.0-cuda11.1-cudnn8-runtime
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY model.py .
CMD ["python", "model.py"]
```
- Build the image:

`docker build -t pytorch-model .`
- Run the container:

`docker run pytorch-model` ### Project 9: Automating Data Pipelines with Apache Airflow and Docker.

In this project, you will set up and containerize an Apache Airflow environment to automate data pipelines. Apache Airflow is a popular orchestration tool for complex workflows widely used in data engineering.

Difficulty level: Advanced

Technologies used: Docker, Apache Airflow, Python, PostgreSQL

#### Step-by-step instructions

- Create the `docker-compose.yml` file: Define the Airflow services and the PostgreSQL database:

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
- Start the Airflow environment: Use Docker Compose to bring up the Airflow environment:

`docker-compose up`
- Access the Airflow user interface: Open your browser and go to http://localhost:8080.

### Project 10: Deploy a Data Science API with FastAPI and Docker

Create and deploy a data science API using FastAPI. You will containerize the API using Docker and focus on optimizing it for production environments.

Difficulty level: Advanced

Technologies used: Docker, FastAPI, Python, scikit-learn

#### Step-by-step instructions

- Write the FastAPI application: Create a simple API that uses a machine learning model for predictions. Here is an example:

```
from fastapi import FastAPI
import pickle
 
app = FastAPI()
with open("model.pkl", "rb") as f:
   model = pickle.load(f)
 
@app.post("/predict/")
def predict(data: list):
   return {"prediction": model.predict(data)}
```
- Write the Docker file: Create a Docker file that defines the environment for FastAPI:

```
FROM python:3.9-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
```
- Build the image:

`docker build -t fastapi-app .`
- Run the container:

`docker run -p 8000:8000 fastapi-app`
## Tips for working on Docker projects

While you work on these projects, keep the following tips in mind:

- Start modestly: Start with slightly difficult projects, then move on to more complex tasks. It is essential to gain confidence with the simplest tasks.
- Log your progress: Keep a detailed record of your projects so that you can track your learning and use it as a reference for your future projects.
- Join Docker communities: Participate in online forums and local meetups to share your experiences, ask questions, and learn from others.
- Experiment and customize: Don't be afraid to modify the projects, try different approaches, and explore new Docker features.
- Keep learning: Continue to build your knowledge of Docker by exploring advanced topics and tools such as Kubernetes, Docker Swarm, or microservices architecture.

## Conclusion

Mastering Docker is not limited to learning commands and configurations. It is about understanding how Docker fits into modern application development, data science workflows, and infrastructure management.

The projects presented in this guide give you some ideas for acquiring the basic skills and hands-on experience needed to excel in real-world scenarios.

At this point, I suggest you consolidate your knowledge by taking these courses:

## Become a data engineer

## FAQ

### What are the best practices for writing efficient Dockerfiles?

**Best practices for writing efficient Dockerfiles include minimizing the number of layers by combining commands, using multi-stage builds to reduce image size, selecting lightweight base images, caching dependencies, and avoiding including unnecessary files in the final image.**

### What is a multi-stage build in Docker?

**A multi-stage build is a method for optimizing Docker images by separating the build and runtime environments. This results in smaller and more secure images.**

### How do you reduce the size of a Docker image?

**Use minimal base images, manage dependencies efficiently, and use multi-stage builds to reduce image size and improve performance.**

### How do you troubleshoot common errors when building Docker images?

**The most common errors when building Docker images are permission issues, incorrect Dockerfile syntax, and failure to install dependencies. To troubleshoot, check the Docker build logs, make sure you are using the correct base image, and confirm that file paths or permissions are set correctly. Tools such as `docker build --no-cache` can help you identify caching issues.**

### Can I use Docker with Kubernetes for these projects?

**Yes, once you are comfortable with Docker, Kubernetes can be the next step. Kubernetes allows you to manage containerized applications at scale. You can deploy your Docker projects on a Kubernetes cluster to manage multiple instances, handle scaling, and automate deployments.**

### What are the best practices for managing Docker volumes and persistent data?

**When working with Docker volumes, it is important to use named volumes to ensure data persistence across container restarts. Back up your volumes regularly and monitor for disk I/O bottlenecks. Avoid storing sensitive data directly in containers; instead, use secure storage solutions or external databases.**

### What is the purpose of the ENTRYPOINT directive in a Dockerfile?

**The `ENTRYPOINT` directive in a Dockerfile specifies the command that will always be executed when the container starts. It allows the container to be treated as an executable, where arguments can be passed at runtime, which improves flexibility.**

### What is the difference between CMD and ENTRYPOINT in a Dockerfile?

**The `CMD` and `ENTRYPOINT` directives specify the commands to run when a container starts. However, `CMD` provides default arguments that can be overridden, while `ENTRYPOINT` defines the command that always runs. `ENTRYPOINT` is useful for creating containers that act as executables, while `CMD` is more flexible for specifying default commands.**

AWS Certified Cloud Solutions Architect, DevOps, cloud engineer with a deep understanding of architecture and high-availability concepts. I have knowledge in cloud engineering and DevOps and I know how to use open-source resources to run enterprise applications. I build cloud-based applications using AWS, AWS CDK, AWS SAM, CloudFormation, Serverless Framework, Terraform, and Django.
