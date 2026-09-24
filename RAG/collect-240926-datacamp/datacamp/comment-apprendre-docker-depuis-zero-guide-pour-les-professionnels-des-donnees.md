---
id: collect-240926-datacamp/datacamp/comment-apprendre-docker-depuis-zero-guide-pour-les-professionnels-des-donnees
title: "Use an official Python runtime as a parent image"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Microsoft"]
dates: []
keywords: ["accelerator", "apache", "aws", "distribution", "memory", "open source", "packaging", "training"]
source: docs/RAG/clean_en/datacamp/comment-apprendre-docker-depuis-zero-guide-pour-les-professionnels-des-donnees.md
source_anchor: ""
source_lines: [1, 466]
sha256: 8d1926cd6d0959ec858e084a0e573ba3a297910bc7c84809c0657da8bec7a16d
---

# Use an official Python runtime as a parent image

<!-- source: https://www.datacamp.com/fr/blog/learn-docker -->

Cursus

Containerization has transformed the way engineering teams manage and scale applications, particularly for data management, analytics, and machine learning. By packaging applications into isolated, lightweight environments, containers ensure consistent performance from development to production.

Among the various platforms available, Docker stands out as the most popular solution. Its flexibility and simplicity allow data professionals to create reproducible, scalable, and efficient pipelines, while fostering collaboration.

In this article, we offer a concrete Docker learning plan, with steps to deploy your first simple application. Let's go!

## What is Docker and why learn it?

Docker is an open-source platform that simplifies the deployment, scaling, and management of applications through containerization.

Containers are portable, lightweight environments that bundle everything needed to run an application — code, runtime, libraries, and settings — to ensure consistent performance across different systems. In data projects, Docker is used to build and manage these containers, allowing applications to run reliably on any infrastructure.

Unlike virtual machines (VMs), which require their own operating system and a hypervisor, Docker only virtualizes the application layer. The result: containers that start faster, consume fewer resources, and are simpler to configure.

*Containerized applications versus virtual machines. Source: Docker*

For data professionals, Docker makes it possible to create reproducible environments to run pipelines with the same reliability from development to production. It reduces dependency issues, streamlines workflows, and fosters collaboration through standardized, shareable environments.

Additionally, Docker integrates with popular data tools like Jupyter, TensorFlow, and Apache Hadoop.

Mastering Docker can accelerate your productivity, optimize your workflows, and make your projects easily deployable and scalable!

## Learning Docker from scratch: your first deployment

The best way to learn Docker is to practice. Let us guide you through a first simple deployment. Then, we'll look at learning plans to deepen your knowledge.

### Step 1: understand the key concepts

Before moving to practice, it's important to grasp a few fundamentals. Here are the main Docker concepts:

- Containers: lightweight, isolated units that package an application with all its dependencies, ensuring identical operation across different environments.
- Images: a read-only template used to create containers. It includes everything needed (code, libraries, system tools). Images are generally built from Dockerfiles.
- Dockerfile: a text file containing the instructions for building a Docker image (installing software, copying files, configuring the environment, etc.).
- Docker Hub: Docker Hub is a public registry for storing, sharing, and downloading Docker images. It facilitates the distribution and reuse of preconfigured environments.
- Volumes: a mechanism for persisting data generated and used by Docker containers, storing it outside the container lifecycle to prevent any loss.
- Networks: Docker networks facilitate communication between containers. Each container can be connected to one or more networks to exchange data securely.

Overview of Docker architecture. Source: *Docker*

Understanding these basic notions is essential before deploying applications with Docker. Mastering them will give you a solid foundation and make practice much more effective.

The Introduction to Docker course can significantly help you consolidate your knowledge.

### Step 2: install Docker

To start using Docker, you need to install it on your system. Here are the instructions depending on your platform. For more details, refer to the official Docker documentation via the links provided.

#### 1. Install Docker on Windows

Prerequisites:

- Windows 10 64-bit: Pro, Enterprise, or Education (Build 19041 or higher)
- Windows 11 64-bit: Home, Pro, Enterprise, or Education
- WSL 2 backend

Steps:

1. Enable WSL 2 (Windows Subsystem for Linux):

- Open PowerShell as administrator.
- Run the following commands:

```
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart 
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart wsl --set-default-version 2
```
2. Install Docker Desktop for Windows:

- Download Docker Desktop from the official Docker website.
- Run the installer and follow the instructions.
- Choose the WSL 2 option as the default backend during installation.

3. Start Docker Desktop:

- Launch Docker Desktop from the Start menu.
- Docker should start automatically; if not, start it manually.

4. Verify the installation:

- Open PowerShell or the command prompt.
- Run the following command to verify that Docker is properly installed:

`sudo docker --version`
Official documentation: Docker Desktop for Windows

#### 2. Install Docker on macOS

Prerequisites:

- macOS 10.15 or later

Steps:

1. Download Docker Desktop for macOS:

- Go to the official Docker website and download the Docker Desktop installer.

2. Install Docker Desktop:

- Open the downloaded `.dmg` file.
- Drag the Docker icon into the Applications folder.

3. Start Docker Desktop:

- Open Docker from the Applications folder.
- Follow the wizard to finalize the installation.

4. Verify the installation:

- Open a terminal.
- Run the following command to verify the Docker installation:

`docker --version`
Official documentation: Docker Desktop for Mac

#### 3. Install Docker on Linux

Supported distributions:

- Ubuntu
- Debian
- Fedora
- CentOS
- RHEL

Steps for Ubuntu/Debian:

1. Uninstall old versions:

- Before installing Docker Engine, remove conflicting packages with this command:

`for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do sudo apt-get remove $pkg; done`
2. Set up Docker's apt repository:

- Run the following commands:

```
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo \"$VERSION_CODENAME\") stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
```
3. Install Docker packages:

- Install Docker Engine and the necessary components:

`sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin`
4. Verify the installation:

- Run the following command to verify that Docker is installed:

`sudo docker --version`
Official documentation: Docker Engine on Debian

## Become a data engineer

### Step 3: Run your first container

Now that Docker is installed, it's time to run your first container. We'll start with the simple hello-world image, perfect for getting started.

1. Open your terminal (or the command prompt on Windows) and run:

`sudo docker run hello-world`
This command tells Docker to look for the `hello-world` image locally. If it can't be found, Docker will download it from Docker Hub and then run it.

If Docker is installed correctly, you will see the message "Hello from Docker!" along with explanations about how the process works. This output confirms that Docker successfully retrieved the image, created a new container, and ran the code inside it.

### Step 4: Build your first Docker image

In this step, you will create a custom Docker image for a mini data science project. It will include Python and common libraries such as pandas, NumPy, and scikit-learn.

1. Create a new directory:

- 
  - Start by creating a directory for your project and navigating into it:

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

### Week 5: Deploy and manage containers in production

- Objective: Understand the deployment and operation of containers in production.
- Tasks:
- Apply deployment best practices (Docker Hub or private registry).
- Explore orchestration with Docker Swarm: creating and managing services.
- Understand container scaling and log management.
- Deploy a containerized application to a staging or production environment.
- Resources: Intermediate Docker.

### Week 6: Introduction to Kubernetes for orchestration

- Objective: Become familiar with Kubernetes to scale and manage containerized applications.
- Tasks:
- Install and configure a local Kubernetes environment (Minikube or Kind).
- Learn the basics: pods, services, deployments, namespaces.
- Deploy a simple containerized application on Kubernetes.
- Explore application scaling and management, for example for a tool like Jupyter.
- Resources: Introduction to Kubernetes.

## Tips and methods for learning Docker

To finish, here are a few tips and best practices to accelerate your upskilling in Docker.

### Practice regularly

The key to mastering Docker is regular, hands-on practice. Start with personal projects, such as containerizing a simple website or setting up a local development environment.

As you gain confidence, take on more complex tasks: building a multi-container application or deploying a web server.

Regularly experimenting with different use cases — such as creating custom Docker images or configuring a CI/CD pipeline — will strengthen your understanding and anchor the concepts. Also take part in online Docker challenges to practice in a structured and motivating way.

### Make use of online resources

Online resources are valuable for learning Docker. The official documentation is an excellent starting point, reliable and up to date.

Beyond that, there are many comprehensive courses on platforms like DataCamp that will guide you step by step.

If you prefer video format, YouTube tutorials from the official Docker channel and other instructors offer useful and accessible content to consolidate what you have learned.

### Join the Docker community

Getting involved in the Docker community is an excellent accelerator. Participate in forums (Docker Community Forums, Reddit) to exchange ideas, ask your questions, and share your feedback with other learners and professionals.

Attending meetups, webinars, or Docker conferences is also a very good way to stay up to date and expand your network.

### Contribute to open source

Contributing to open source projects that use Docker is an excellent way to gain concrete and relevant experience.

Platforms like GitHub offer opportunities to collaborate, learn from others, and apply your skills in varied contexts. This collaboration exposes you to different approaches and solutions, enriching your understanding of Docker's capabilities.

### Stay up to date

Finally, staying informed about the latest Docker developments is essential to maintaining your level. Follow the official blog and release notes to discover new features. Community discussions on forums and social media will also help you spot trends and best practices to continually refine your skills.

With regular practice, active engagement in the community, and continuous learning, you will build solid foundations in Docker and remain ready to adopt the latest advances in the field.

## Conclusion

Docker has revolutionized application deployment and data management by packaging applications and dependencies into isolated containers. It solves classic problems of environment inconsistency and guarantees consistent, scalable, and efficient workflows for data professionals.

This guide has presented the essential steps to mastering Docker, from key concepts to practical deployments. To go further, explore resources like the Introduction to Docker and Intermediate Docker courses on DataCamp. As you progress, it is equally important to develop your orchestration skills with Kubernetes — discover the Introduction to Kubernetes course to get started.

By practicing regularly, making use of available resources, and staying active in the community, you will unlock the full potential of Docker to manage complex data environments.

## Get certified for the Data Engineer job of your dreams

Our certification programs help you stand out and prove to potential employers that your skills are suited to the job.

## FAQs

### How long does it take to learn Docker?

**The time needed to learn Docker depends on your experience. For beginners, a few days are enough to become familiar with the basics (containers, images, Dockerfiles). To go further — Docker Compose, Kubernetes — count on a few weeks of regular practice.**

### Should I learn Kubernetes after Docker?

**Docker is essential for containerization, while Kubernetes is valuable for managing and orchestrating large-scale deployments across multiple containers. After mastering Docker, learning Kubernetes is highly recommended for complex applications or when you need to automate container scaling and management.**

### What projects should I do to practice Docker?

**To gain hands-on experience with Docker, start with small projects: containerize a personal website or create a local development environment for a web application. Then, build a multi-container application with Docker Compose or deploy a simple CI/CD pipeline. You can also practice by contributing to open source projects that use Docker.**

### Is Docker essential for data professionals?

**Yes, Docker is extremely useful for data professionals. It allows you to create reproducible environments to ensure your pipelines run consistently from development to production. Docker also simplifies dependency management and strengthens collaboration between teams: an essential tool for modern data workflows.**

### What are the best resources for learning Docker?

Among the best resources to get started with Docker:

- The official Docker documentation
- The DataCamp courses Introduction to Docker and Intermediate Docker
- YouTube tutorials from the official Docker channel
- Docker community forums and GitHub repositories to practice and collaborate

AWS Certified Cloud Solutions Architect, DevOps, cloud engineer with a deep understanding of architecture and high availability concepts. I have knowledge in cloud engineering and DevOps and I know how to use open-source resources to run enterprise applications. I build cloud-based applications using AWS, AWS CDK, AWS SAM, CloudFormation, Serverless Framework, Terraform and Django.
