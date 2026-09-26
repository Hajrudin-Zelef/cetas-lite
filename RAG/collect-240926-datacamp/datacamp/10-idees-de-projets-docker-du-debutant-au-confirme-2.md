---
id: collect-240926-datacamp/datacamp/10-idees-de-projets-docker-du-debutant-au-confirme-2
title: "Stage 1: Build"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "inference"]
source: docs/RAG/clean_en/datacamp/10-idees-de-projets-docker-du-debutant-au-confirme.md
source_anchor: ""
source_lines: [226, 470]
sha256: 1beb95b3ca8188d4666a079bee25d4b04fd3dbdb1f3c9cd23444faf66ede0a2b
---

# Stage 1: Build

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

