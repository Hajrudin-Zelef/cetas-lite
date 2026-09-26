---
id: collect-240926-datacamp/datacamp/comment-apprendre-docker-depuis-zero-guide-pour-les-professionnels-des-donnees-2
title: "Use an official Python runtime as a parent image"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "training"]
source: docs/RAG/clean_en/datacamp/comment-apprendre-docker-depuis-zero-guide-pour-les-professionnels-des-donnees.md
source_anchor: ""
source_lines: [184, 368]
sha256: dcde2a1ed5a62b81145136bf2a768ccdb9a466c7d4855f7abab1bc3066624e49
---

# Use an official Python runtime as a parent image

```
mkdir my-data-science-app
cd my-data-science-app
```
2. Create a Dockerfile:

In this directory, create a file named `Dockerfile` (without an extension):

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
This Dockerfile performs the following actions:

- `FROM python:3.8-slim`: sets the base image to Python 3.8, "slim" variant, ideal for a lightweight data environment.
- `WORKDIR /app`: sets the container's working directory to /app. The following commands will run there.
- `COPY . /app`: copies all files from the host's current directory to `/app` in the container (scripts, configuration files, etc.).
- `RUN pip install --no-cache-dir -r requirements.txt`: installs the Python libraries listed in `requirements.txt` (pandas, NumPy, scikit-learn, ...).
- `EXPOSE 8888`: exposes port 8888 to the host, the default port for Jupyter notebooks.
- `CMD ["jupyter", "notebook", "--ip=0.0.0.0", "--port=8888", "--no-browser", "--allow-root"]`: starts a Jupyter Notebook server when the container launches, accessible from your browser.

3. Create the `requirements.txt` file:

In the same directory, create a `requirements.txt` file listing the required Python packages:

```
pandas
numpy
scikit-learn
matplotlib
```
This file lists the Python libraries that will be installed in your Docker container.

4. Build the Docker image:

- With the Dockerfile and `requirements.txt` ready, build the image with:

`sudo docker build -t my-data-science-app .`
This command does the following:

- `docker build`: starts the image build process.
- `-t my-data-science-app`: tags the image "my-data-science-app".
- `.`: the final dot designates the build context (the current directory).

5. Run the Docker image:

Once the build is complete, run your new image as a container:

`sudo docker run -p 8888:8888 my-data-science-app`
Command details:

- `docker run`: starts a container from the specified image.
- `-p 8888:8888`: maps port 8888 of your local machine to port 8888 of the container to access the Jupyter server via `localhost:8888`.
- `my-data-science-app`: the image to run.

### Step 5: Use Docker Compose

Docker Compose is essential for managing multi-container applications. In a data science project, you might have separate containers for a Jupyter notebook, a database, and a visualization tool.

With Docker Compose, you define these services in a single YAML file, which allows you to launch your entire environment with a single command.

1. Create a `docker-compose.yml` file:

- In your project directory, create a `docker-compose.yml` file and add the following lines:

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
Here is what this Docker Compose file does:

- `version: '3.8'`: version of the Docker Compose file format, widely used.
- `services:`: section that defines the different containers (services) of your application.
- `jupyter:`: service for the Jupyter environment, essential for interactive analysis.
- `image: jupyter/scipy-notebook:latest`: Docker image for Jupyter including NumPy, pandas, matplotlib, etc.
- `volumes: - ./notebooks:/home/joelwembo/work`: mounts the `notebooks` folder from the host into `/home/joelwembo/work` of the container to persist and share your notebooks. Note: this is my personal path, yours will be different.
- `ports: - "8888:8888"`: maps port 8888 of the host to that of the container (default Jupyter port).
- `environment: - JUPYTER_ENABLE_LAB=yes`: enables JupyterLab, a more powerful interface for working with Jupyter.
- `postgres:`: PostgreSQL service, often used to store and manage large datasets.
- `image: postgres:13-alpine`: PostgreSQL 13 image based on Alpine Linux, lightweight.
- `environment:`: environment variables to configure the user, password, and database name.
- `volumes: - postgres_data:/var/lib/postgresql/data`: Docker volume to persist the database data beyond the container lifecycle.
- `ports: - "5432:5432"`: exposes the default PostgreSQL port (5432).
- `redis:`: Redis service, used as an in-memory data store or fast cache.
- `image: redis:alpine`: Redis image based on Alpine, lightweight.
- `ports: - "6379:6379"`: exposes the default Redis port (6379).
- `volumes:`: section that defines named volumes shared between services or persistent across restarts.
- `postgres_data:`: creates a named volume for PostgreSQL data.

To launch your multi-container application, simply run:

`sudo docker compose up`
This command builds the images (if necessary) and starts the containers defined in your `docker-compose.yml`. It starts all the services and allows them to communicate as expected.

You have just deployed your first Docker application! Pretty motivating, isn't it? Let's now move on to a plan for your continued training.

## Example Docker learning plan

Here is a week-by-week plan that you can follow and adapt to your needs.

If you prefer a more structured path, the Containerization and Virtualization with Docker and Kubernetes track is for you. It contains four essential courses.

### Week 1: Getting familiar with Docker basics

- Objective: understand the key concepts of Docker and the basic commands.
- Tasks:
- Install Docker and configure Docker Desktop.
- Learn the commands: `docker run`, `docker ps`, `docker stop`, `docker rm`, etc.
- Launch simple containers for exploration (Ubuntu, Alpine) and test basic commands.
- Explore official Docker images and run containers of simple visualization tools (e.g., matplotlib).
- Resources: Containerization and Virtualization Concepts, Introduction to Docker.

### Week 2: Create and run Docker images for data tools

- Objective: learn to create and manage custom Docker images.
- Tasks:
- Write a Dockerfile for an analysis environment (Python + pandas, NumPy).
- Build the image with `docker build`.
- Run containers from your image and test tools like Jupyter Notebook.
- Create Dockerfiles for TensorFlow and PostgreSQL, with correct configuration and chaining.
- Resources: Introduction to Docker.

### Week 3: Discover Docker Compose

- Objective: understand and use Docker Compose for multi-container environments.
- Tasks:
- Install Docker Compose and learn the basics.
- Write `docker-compose.yml` files to define multi-container applications.
- Practice, for example, with a web application and a database (e.g., Flask + PostgreSQL).
- Test different configurations and service dependencies.
- Resources: Intermediate Docker.

### Week 4: Go deeper into Docker networking and volumes

- Objective: master Docker's networking and storage capabilities.
- Tasks:
- Discover Docker network types (bridge, host, overlay) and how to create them.
- Create custom networks and connect containers to them.
- Explore volumes for persistent storage and practice managing them.
- Set up a configuration where containers communicate via a custom network and use volumes for persistence.
- Resources: Intermediate Docker.

