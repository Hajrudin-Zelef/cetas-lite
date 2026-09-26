---
id: collect-240926-datacamp/datacamp/les-26-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-concernant-d-2
title: "Step 1: Choose a base image"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["memory"]
source: docs/RAG/clean_en/datacamp/les-26-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-concernant-docker-pour.md
source_anchor: ""
source_lines: [142, 243]
sha256: 1a917690b78dfdf5026204a912bac3b41afa0a242b5bae7397f1353e353e2df8
---

# Step 1: Choose a base image

No, Docker Swarm does not natively support automatic scaling. To achieve autoscaling, it is necessary to integrate monitoring tools and use scripts to manually adjust the number of instances. Here is how to proceed:

- Please install a monitoring tool, such as Prometheus or Grafana, to track resource usage, such as CPU and memory.
- Please define the scaling triggers. For example, we can define that CPU usage above 82% will trigger an increase in capacity.
- Then, please write a script using the `docker service scale` command to adjust the number of replicas. For example, to scale a service to 5 replicas:`docker service scale =5`

By combining monitoring tools, triggers, and scripts, it is possible to implement a form of autoscaling in Docker Swarm, even though this feature is not built in.

### 12. How would you use Docker Compose to scale services?

To scale services using Docker Compose, we can use the ` `--scale` ` flag with the ` `docker-compose up` ` command. This is generally used for stateless services such as web servers. For example, to scale a web service to 3 instances:

`docker-compose up --scale web=3`

It is essential to ensure that the `docker-compose.yml` file correctly defines the services and uses an external load balancer or supports variable-scale instances. Scaling stateful services (for example, databases) requires additional configuration to ensure data consistency.

### 13. Can a container restart by itself? Define its default and permanent policies.

Yes, a container can restart on its own. However, it is necessary to define a restart policy for this purpose.

Docker has different restart policies that determine when and how containers should restart. The default policy is "no," which means that a container will not restart if it stops. With the "always" policy, Docker will automatically restart the container whenever it stops.

We can use this command to apply the "always" policy:

`docker run --restart=always` 

## Advanced Docker Interview Questions

Let's now move on to advanced Docker interview questions.

### 14. Please explain the Docker container lifecycle.

A Docker container follows a lifecycle that defines the states it can be in and how it operates in those states. The stages of a Docker container's lifecycle are as follows:

- Create: In this state, we configure a container from an image using the `docker create` command.
- **Run**: Here, we use the `docker start` command to run the container, which performs tasks until we stop or pause it.
- **Pause**: We use the `docker pause` command to suspend the process. This state preserves memory and disk. If you want to restart the container, please use the `docker unpause` command.
- **Stop**: If the container is inactive, it enters the stop phase, but this can happen for several reasons:
- **Immediate stop**: The `docker kill` command stops the container without cleanup.
- **Process completion**: Once the task is finished, the container stops automatically.
- **Insufficient memory**: The container stops when it uses an excessive amount of memory.
- **Remove**: In the last step, we remove the stopped or created container using the `docker rm` command.

### 15. What is a Docker image repository?

A Docker image repository stores and shares multiple container images with the same name with clients or the community. We can tag them using tags to distinguish their different versions. For example, `app/marketing_campaign:v1` will be the first version of a marketing application, and `app/marketing_campaign:v2` will be the second version.

Docker Hub, the most popular Docker image repository, allows users to host, share, and retrieve container images publicly or privately. Other alternatives include Amazon ECR, Google Artifact Registry, and GitHub Container Registry.

### 16. Please tell me three best practices for ensuring the security of a Docker container.

In order to strengthen container security and minimize common vulnerabilities, I follow the following best practices:

1. **Please select lightweight images**: Please use minimal base images such as Alpine to reduce the attack surface.
2. **Limit system calls**: Since Docker containers can access unnecessary calls, it is recommended to use tools such as Seccomp to limit these calls.
3. **Secure sensitive data**: Please use Docker secrets to manage API keys or passwords. They encrypt secrets and make them accessible only during runtime.

### 17. Why do Docker containers need health checks?

Docker containers rely on health checks to ensure they are working properly. Deploying a container that is running but not processing requests can cause problems for deployment teams. Health checks monitor these issues in real time and inform us immediately.

For example, a health check can be added in a Dockerfile as follows:

`HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://localhost:8080/health || exit 1`

This health check queries the container's health endpoint every 30 seconds and marks the container as non-functional if it fails three consecutive attempts. This proactive monitoring helps identify and resolve issues quickly.

### 18. What are dangling images in Docker and how can they be removed?

Dangling images in Docker are unused image layers that are no longer associated with any tag. They often accumulate when you build new images with the same name and tag, leaving the old layers without references. These images can take up significant disk space, so it is essential to remove them. Here is how to do it:

1. Please run the `docker images -f dangling=true` command to identify dangling images.
2. Then, please run the `docker image prune -f` command to remove all images at once.
3. If you want to remove images manually, please use the `docker rmi -f $(docker images -f dangling=true -q)` command.

These steps help keep your system clean and free up storage space efficiently.

## Docker and Kubernetes Interview Questions

Docker and Kubernetes are often used together, so it is not surprising to encounter Kubernetes-related questions during a job interview for a Docker position, especially if the role is DevOps-oriented. Here are some questions you might be asked:

### 19. What is the main difference between Docker and Kubernetes?

Docker is a containerization platform that allows you to create, ship, and run containers. It focuses on creating and managing individual containers. Kubernetes, on the other hand, is an orchestration platform designed to manage multiple containers at scale. It handles deployment, scaling, load balancing, and self-healing across clusters of nodes.

To learn more about the differences between Kubernetes and Docker, please see the blog article.

### 20. Please compare Docker Swarm and Kubernetes.

Kubernetes and Docker Swarm manage containers, but they work differently:

- Kubernetes manages large, complex container configurations. Its self-healing and built-in monitoring features make it a more suitable option for complex environments.
- Docker Swarm is suitable for smaller or less complex configurations, as it does not offer built-in features like Kubernetes. We can easily integrate it with Docker tools such as Docker CLI and Docker Compose.

### 21. How does Kubernetes handle a large number of Docker containers?

Although Docker is an excellent tool for creating and running containers, managing a large number of them requires the use of Kubernetes. Kubernetes efficiently coordinates containers by:

- Defining resource limits: It allocates CPU, memory, and other resources to each container to prevent overconsumption.
- Scheduling containers: Kubernetes determines where to run each container, optimizing resource usage across all nodes in a cluster.
- Automatic scaling: Depending on the workload, it increases or decreases the number of pods (groups of one or more containers) to maintain performance and efficiency.

