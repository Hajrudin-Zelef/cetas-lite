---
id: vague2-datacamp/datacamp/docker-commands
title: "Les 18 principales commandes Docker pour construire, exécuter et gérer des conteneurs"
domain: datacamp
role: reference
task: tutorial
actors: []
dates: ["2026-09-23"]
keywords: []
source: docs/RAG/Collect RAG Vague 2/02_datacamp/docker-commands.md
source_anchor: ""
source_lines: [1, 63]
sha256: 2d39b4bde4ec59c9b828dd0257d3bf28eb21cca509146d153b1de55d767125cb
---

# Les 18 principales commandes Docker pour construire, exécuter et gérer des conteneurs

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/docker-commands
- **Site** : DataCamp
- **Type** : Tutorial / Guide
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp guide by Laiba Siddiqui presents 18 essential Docker commands for building, running, and managing containers, organized into categories: basics, images, containers, networking, volumes, and Compose. It begins with an overview of Docker as a platform for developing, shipping, and running applications in lightweight, isolated containers that bundle dependencies, and notes that every Docker command starts with `docker`.

The basic commands section covers `docker --version` and `docker version` (client and server info), `docker info` (full configuration overview), `docker pull` (downloading images from Docker Hub, defaulting to `:latest`), `docker run` (creating and starting containers, with flags like `--name`, `-w/--workdir`, `--pid`, `--cidfile`), and `docker stop`/`docker start`.

The image commands cover `docker build` (constructing images from a Dockerfile, e.g. a Node.js example, tagged with `-t`), `docker images` (listing top-level images with aliases), and `docker rmi` (removing images, with `-f` to force).

Container management commands include `docker exec` (running commands inside a running container, useful for debugging), `docker logs` (viewing container logs with `--details`, `--until`, `-f`), `docker rm` (removing containers), and `docker restart` (with `-s/--signal` and `-t/--timeout`).

Networking commands cover `docker network ls` and `docker network create` (bridge for single-host, overlay for multi-host Swarm). Volume commands cover `docker volume ls` and `docker volume create` (with a busybox mount example). Compose commands cover `docker-compose up` (builds/recreates/starts/attaches, with `--detach`) and `docker-compose down` (stops and removes containers, networks, volumes).

The guide ends with best practices (use volumes for persistent data, automate with Docker Compose) and an extensive FAQ covering common command questions.

## Key points

- 18 essential Docker commands grouped by images, containers, networking, volumes, and Compose.
- `docker run` creates and starts a new container; `docker start` restarts an existing stopped container.
- `docker build -t <tag> .` builds an image from a Dockerfile in the current directory.
- `docker exec` runs commands inside a running container without restarting it (PID 1 must be running).
- `docker network create` supports `bridge` (single host) and `overlay` (multi-host Swarm) drivers.
- Volumes persist data beyond container lifecycle and are more portable/secure than bind mounts.
- `docker compose up --detach` runs multi-container services in the background; `docker compose down` tears them down.
- `docker ps -a` lists all containers including stopped ones; `docker image prune` cleans dangling images.

## Technical data / figures

| Command | Purpose |
|---------|---------|
| `docker --version` / `docker version` | CLI version / detailed client+server info |
| `docker info` | Full Docker configuration overview |
| `docker pull <image>` | Download image (default `:latest`) |
| `docker run <image>` | Create and start a container |
| `docker start` / `docker stop` | Start/stop containers |
| `docker build -t <tag> .` | Build image from Dockerfile |
| `docker images` (`docker image ls`) | List top-level images |
| `docker rmi <image>` | Remove image (`-f` to force) |
| `docker exec <container> <cmd>` | Run command inside running container |
| `docker logs <container>` | View container logs |
| `docker rm <container>` | Remove container |
| `docker restart <container>` | Stop then restart (`-s`, `-t`) |
| `docker network ls` | List networks |
| `docker network create -d bridge <name>` | Create bridge network |
| `docker network create --scope=swarm --attachable -d overlay <name>` | Create overlay network |
| `docker volume ls` / `docker volume create <name>` | List/create volumes |
| `docker compose up --detach` / `docker compose down` | Start/stop multi-container app |
| `docker ps -a` | List all containers |
| `docker image prune` | Remove dangling images |

## Why this source matters for the RAG

This command reference provides a compact, categorized catalog of Docker CLI syntax and options with concrete examples, making it excellent for retrieval-based answers on "how do I do X in Docker." Its coverage of images, containers, networking, volumes, and Compose directly supports technical Q&A and troubleshooting use cases.
