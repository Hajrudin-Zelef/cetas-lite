---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-1
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit", "jailbreak", "memory", "research"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages.md
source_anchor: ""
source_lines: [1, 95]
sha256: 2483745187670793e68d4af3f4ed3dc7a16166dc104b59ffaa8882bbbdd0c4ce
---

# Summary

**Sub-menu:** `/container`

Packages requied: `container`

A container is MikroTik's implementation of Linux containers, allowing users to run containerized environments within RouterOS. The container feature works in the latest MikroTik RouterOS v7.x version. Containers are compatible with images from Docker Hub, GCR, Quay, or other providers, as well as those built on other devices, using the same formats supported by these providers. While RouterOS uses different syntax compared to Docker, it still achieves similar functionality.

# Disclaimer

you need physical access to your RouterOS device to enable support for the container feature, it is disabled by default;

- once the container feature is enabled, containers can be added/configured/started/stopped/removed remotely!
- if your RouterOS device is compromised, containers can be used to easily install malicious software in your RouterOS device and over network;
- your RouterOS device is as secure as anything you run in container;
- if you run container, there is no security guarantee of any kind;
- running a 3rd party container image on your RouterOS device could open a security hole/attack vector/attack surface;
- an expert with knowledge how to build exploits will be able to jailbreak/elevate to root;

## **Security risks:**

- When a security expert publishes his exploit research - anyone can apply such an exploit;
- Someone can build a container image that can use the exploit AND provide a Linux root shell;
- By using a root shell someone may leave a permanent backdoor/vulnerability in your RouterOS system even after the container image is removed and the container feature disabled;
- If a vulnerability is injected into the primary or secondary RouterBOOT (or vendor pre-loader), then even Netinstall may not be able to fix it;

# Requirements

Container package is compatible with **arm, arm64** and **x86** architectures. Using of remote-image (similar to docker pull) functionality requires a lot of free space in main memory, 16MB SPI flash boards may use pre-build images on USB or other disk media.

- An external disk supporting at least 100MB/s sequential read/write speed and 10K random iops recommended. When using slower disks, container extraction times may become longer.
- Container package needs to be installed
- For devices with EN7562CT CPU like the hEX Refresh, only arm32v5 container images are supported, meaning a limited number of containers can be run.

# Properties

| Property | Description | 
|---|---|
| **auto-restart-interval**  (string; Default: ) | Specify an interval at which Container will be restarted on Container failure. Example: 10s | 
| **cmd** (string; Default: ) | The main purpose of a CMD is to provide defaults for an executing container. These defaults can include an executable, or they can omit the executable, in which case you must specify an ENTRYPOINT instruction as well. | 
| **comment** (*string* ; Default: ) | Short description | 
| **dns** (string; Default: ) | If container needs different DNS, it can be configured here | 
| **domain-name** (string; Default: ) |  | 
| **entrypoint (** string; Default:**)** | An ENTRYPOINT allows to specify executable to run when starting container. Example: /bin/sh | 
| **envlist** (*string* ; Default: ) | list of environmental variables (configured under */container envs* ) to be used with container | 
| **file** (*string; Default:* ) | container *tar.gz tarball if the container is imported from a file | 
| **hostname** (*string; Default:* ) | Assigning a hostname to a container helps in identifying and managing the container more easily | 
| **interface** (*string; Default:* ) | veth interface to be used with the container | 
| **logging** (*string; Default:* ) | if set to yes, all container-generated output will be shown in the RouterOS log | 
| **start-on-boot** (*string; Default:* ) | if set to yes, the container will be started automatically on device start-up. | 
| **mountlists** (*string; Default:* ) | mounts from /container/mounts/ sub-menu to be used with this container | 
| **mount** (*string; Default:* ) | specify directory to be used as a mount | 
| **remote-image** (*string; Default:* ) | the container image name to be installed if an external registry is used (configured under /container/config set registry-url=...) | 
| **root-dir** (*string; Default:* ) | used to save container store outside main memory | 
| **stop-signal** (*string; Default: 15* ) | Type of Linux signal to send when container was not stopped after 10 seconds | 
| **workdir** (*string* ; Default: ) | the working directory for cmd entrypoint | 
| **devices**  (*string* ; Default: ) | passes through physical device to the container | 
| **cpu-list**  (*string* ; Default: ) | specifies which CPU cores the container is allowed to run on | 
| **user**  (*string* ; Default: ) | sets the user and group the container process runs as before execution. | 
| **memory-high**  (*int* ; Default: ) | RAM usage limit in bytes for a specific container | 
| **memory-max** (*int* ; Default: ) | max RAM usage limit in bytes per container (The container process will be terminated if the memory-max value is smaller than the container memory-current.) | 

**Menu specific commands**

| Property | Description | 
|---|---|
| **update** | Updates the container image. Automatically pulls from container repository and extracts it, replacing the original image. | 
| **kill** | Kills the specified running container. | 
| **repull** | Re-pulls/extracts the container image. | 
| **shell**  | Enters the container shell of a running container. | 
| **run** | Starts the container and enters it's shell. Useful if container shuts down after running. | 

# Container configuration

| Property | Description | 
|---|---|
| **registry-url** | external registry url from where the container will be downloaded (default : https://lscr.io/ ) | 
| **tmpdir** | container extraction directory | 
| **memory-high** | RAM usage limit in bytes 1 - unlimited | 
| **username** | Specifies the username for authentication ( starting from ROS 7.8) | 
| **password** *sensitive* | Specifies the password for authentication ( starting from ROS 7.8) | 

# Examples

## Running Pi-hole

### Prerequisites

  1. RouterOS device with RouterOS v7.4beta or later and **installed Container package -** How to install packages
  2. Physical access to a device to enable container mode - will be explained down bellow
  3. Attached HDD, SSD or USB drive for storage - formatted with a filesystem supported by RouterOS - How to format/manage disks
2. RouterOS device with RouterOS v7.4beta or later and 

### Steps to run Pi-hole

