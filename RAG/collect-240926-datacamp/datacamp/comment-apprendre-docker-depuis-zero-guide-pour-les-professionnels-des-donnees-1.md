---
id: collect-240926-datacamp/datacamp/comment-apprendre-docker-depuis-zero-guide-pour-les-professionnels-des-donnees-1
title: "Use an official Python runtime as a parent image"
domain: datacamp
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["apache", "distribution", "packaging"]
source: docs/RAG/clean_en/datacamp/comment-apprendre-docker-depuis-zero-guide-pour-les-professionnels-des-donnees.md
source_anchor: ""
source_lines: [1, 183]
sha256: e5910c9534990baddd54fad5ab2e0e6e4bcc8d7308ac67f4cd2775e59d71fb53
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

