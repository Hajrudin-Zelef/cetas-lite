---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-2-2
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages.md
source_anchor: ""
source_lines: [96, 213]
sha256: 798596dd655cb6c7cfc349d18f0efc5376c158d4fa6f216dae44ef4c4b153d37
---

# Summary

1. Enable Container mode and follow the instructions the command gives you (read more about Device-mode). You will need to confirm the device-mode with a press of the reset button, or a cold reboot (if using Containers on x86):
2. Create a new veth interface and assign an IP address in a range that is unique in your network:The following configuration is equivalent to "bridge" networking mode in other Container engines such as Docker. It is possible to create a "host" equivalent configuration as well.
3. Create a new bridge that is going to be used for your Containers and assign the same IP address that was used for the veth interface's gateway:
4. Add the veth interface to your newly created bridge:
5. Create a NAT for outgoing traffic:
6. Create environment variables for the Container:
7. Create mounted volumes for the Container:`src=` points to RouterOS location (could also be`src=disk1/etc_pihole` if, for example, You decide to put configuration files on external USB media),`dst=` points to defined location (consult containers manual/wiki/github for information on where to point). If`src` directory does not exist on first time use then it will be populated with whatever container have in`dst` location.It is highly recommended to place any Container volume on an attached disk to your RouterOS device. Avoid placing Container volumes on the built-in storage.
8. Configure to use a specific Container repository, for example, to use Docker.io:
9. Add a Containter:If You wish to see container output in `/log print` , then add`logging=yes` when creating a Container, root-dir should point to an external drive. It's not recommended to use internal storage for Containers.Adding a Containter will start downloading or extracting it, the Container itself will not be started after it has been added, you need to start it manually for the first time after it has been downloaded/extracted.
10. Check the status of your Container and wait until downloading/extracting has been finished and the `status=stopped` :
11. Start the Containter:
12. Create a port forwarding for your Container:
13. You should be able to access the Pi-hole web panel by navigating to `http://192.168.88.1/admin/` in your web browser.
14. To start using Pi-hole on your devices, change their DNS configuration to use `192.168.88.1` as your DNS server.

## Adding a Container image

There are multiple ways you can get a Container image running on your RouterOS device. Check the examples below.

### Option A: Get an image from an external library

Set registry-url (for downloading containers from Docker registry) and set extract directory (tmpdir) to attached USB media:

pull image:

The image will be automatically pulled and extracted to root-dir, status can be checked by using

### Option B: Import image from PC

Your can use your PC running either Docker or Podman to download your required container image and save it to an archive. We recommend using Podman since it is easier to build and download containers for specific architectures using Podman.

1. Download your required image based on the architecture of your RouterOS device:
2. Save the container image to an archive:
3. Upload the archive to your RouterOS device, for example:You can also use Winbox to upload files!
4. Create a Container on your RouterOS device using the uploaded container image archive file:

### Option C: Build an image on PC

You can build your own Containers and use them on your RouterOS device. While you can build Containers using Docker, we recommend using Podman since it is easier to build Containers for a specific architecture using Podman.

1. Get source files for your required Container image, for example by using git:
2. Build the Container image by specifying the Dockerfile or Containerfile and the target archiceture:
3. Save the container image to an archive:
4. Upload the archive to your RouterOS device, for example:You can also use Winbox to upload files!
5. Create a Container on your RouterOS device using the uploaded container image archive file:

#### Alternative: Using Docker to build Container images

To use Dockerfile and make your own docker package - docker needs to be installed as well as buildx or other builder toolkit.

Easiest way is to download and install Docker Engine:

https://docs.docker.com/engine/install/

After install check if extra architectures are available:

should return:

If not - install extra architectures:

pull or create your project with Dockerfile included and build, extract image (adjust --platform if needed):

Upload *pihole.tar* to Your RouterOS device.

Images and objects on the Linux system can be `pruned` 

Create a container from the tar image

## Networking examples

### Bridge with NAT

In this networking setup, all Containers use the same veth interface and communicate with each other without any Firewall restrictions, but you need to forward ports in order to allow access to a Container's port.

For example, a database Container needs to communicate with a web application Container, the web application needs the port `80` to be exposed to the world, but the database Container does not need any ports to be exposed to the world.

- The network configuration:
- The database Container configuration:
- The webapp Container configuration:

In this example, the `pgadmin` port 80 is accessible to everyone, but the `postgres` port 5432 is not accessible to everyone, it can only be accessed through either `pgadmin` as `127.0.0.1` or through the RouterOS device running the Containter as `172.17.0.2` .

### Isolated Containers

In this networking setup, you have multiple Containers and you want to make sure that some of them can communicate without Firewall restrictions, but some need to be isolated from other Containers. For example, you might want to create two database Containers and isolate them.

- The network configuration:
- The first and second database Container configuration:
- The first and second webapp Container configuration:

In this example, `pgadmin1` is able to reach `postgres1` , but is not able to reach `postgres2` . Similarly `pgadmin2` is able to reach `postgres2` , but is not able to reach `postgres1` . 

### Container in Layer2 network

In this networking setup, your Container is directly attached to a Layer2 network with other physical network devices. This networking setup is equivalent to "host" networking mode on other Container engines such as Docker.

In this networking setup, all the ports on your Container are exposed. This is considered insecure, but does slightly improve the Container's networking performance.

- The networking configuration:
- In case your RouterOS device has services running on the same port, you need to disable them:
- The webapp configuration:

In this example, `pgadmin` Container does not need port forwarding, but all other ports that the Container is using are now accessible to others on the same Layer2 network. This type of setup should only be used when your application requires that the Container has an IP address in the same Layer2 network such as application that use broadcast traffic for service discovery (in most cases such requirements can still be bypassed by using NAT).

## IPv4 and IPv6 for Container

In this networking setup your Container will be able to communicate over IPv4 and IPv6. The solution is based on Bridge with NAT networking setup.

- The network configuration:
- The webapp Container configuration:

# Tips and tricks

- Containers use up a lot of disk space, USB/SATA, NVMe attached media is highly recommended. For devices with USB ports - USB to SATA adapters can be used with 2.5" drives - for extra storage and faster file operations.
- RAM usage can be limited by using:

this will soft limit RAM usage - if a RAM usage goes over the high boundary, the processes of the cgroup are throttled and put under heavy reclaim pressure.

