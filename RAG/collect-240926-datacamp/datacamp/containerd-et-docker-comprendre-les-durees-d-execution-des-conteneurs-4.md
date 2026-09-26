---
id: collect-240926-datacamp/datacamp/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs-4
title: "containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs"
domain: datacamp
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "training"]
source: docs/RAG/clean_en/datacamp/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs.md
source_anchor: ""
source_lines: [275, 307]
sha256: de7d081e29ea9cacaff78cd196043a0f6d01fed65d1ffb9d0fcd7759c8d4cd95
---

# containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs

To continue your training, we invite you to sign up for our skills course "Containerization and virtualization with Docker and Kubernetes."

## FAQ on Containerd and Docker

### Is it possible to use Docker images with containerd?

**Yes, absolutely. Containerd supports all container images that comply with the OCI standard, including those created with Docker. Since Docker creates OCI-compliant images, they work perfectly with containerd and all other OCI-compatible runtime environments. You can use `docker build` locally and run those images with containerd in production without any compatibility issues.**

### Does Docker use containerd in the background?

**Yes, Docker Engine uses containerd as its primary container runtime. Starting from version 1.11, Docker integrated containerd to manage container lifecycle operations, such as creation, execution, and management. When you run `docker run`, the Docker daemon (`dockerd`) delegates the actual execution of the container to containerd, which then uses runc to interact with the Linux kernel.**

### Why did Kubernetes remove Docker support?

**In 2022, Kubernetes removed dockershim (the Docker compatibility layer) in order to eliminate an unnecessary translation layer. Docker predates the Container Runtime Interface (CRI), so Kubernetes needed dockershim to handle the conversion between its APIs and Docker. By communicating directly with containerd via CRI, Kubernetes achieves better performance, greater stability, and a simpler runtime stack. Docker images continue to work perfectly in Kubernetes.**

### Should I switch from Docker to containerd for local development?

**No, Docker remains the most appropriate choice for local development. Docker provides an integrated toolchain with Docker Compose, the Docker Desktop graphical interface, and extensive ecosystem support that speeds up development workflows. Please use containerd for production Kubernetes clusters, where its low overhead and direct integration with CRI offer clear advantages, but keep Docker on developers' laptops to benefit from its superior user experience.**

### What is nerdctl and do I need it?

**Nerdctl is a Docker-compatible CLI for containerd that offers the same user experience as Docker (supports most common commands and flags) but uses containerd as the runtime. You need it if you want to interact directly with containerd using the usual Docker commands. It is particularly useful for development environments that use containerd or when teams are transitioning from Docker to containerd-based workflows.**

As the founder of Martin Data Solutions and a freelance Data Scientist, ML and AI engineer, I bring a diverse portfolio in regression, classification, NLP, LLM, RAG, neural networks, ensemble methods, and computer vision.

- Successfully developed several end-to-end ML projects, including data cleaning, analysis, modeling, and deployment on AWS and GCP, delivering impactful and scalable solutions.
- Created interactive and scalable web applications using Streamlit and Gradio for various use cases in the industry.
- Teaches and mentors students in data science and analytics, fostering their professional development through personalized learning approaches.
- Designed course content for retrieval-augmented generation (RAG) applications tailored to enterprise requirements.
- Wrote high-impact technical blogs on AI and ML, covering topics such as MLOps, vector databases, and LLMs, with significant engagement.

In every project I take on, I make sure to apply up-to-date software engineering and DevOps practices, such as CI/CD, code linting, formatting, model monitoring, experiment tracking, and robust error handling. I am committed to delivering comprehensive solutions, transforming data insights into practical strategies that help businesses grow and make the most of data science, machine learning, and AI.
