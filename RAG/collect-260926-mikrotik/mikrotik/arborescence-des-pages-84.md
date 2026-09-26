---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-84
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-84.md
source_anchor: ""
source_lines: [1, 101]
sha256: 9f17ce7ccd5316938095dfe99574f2ec4a2b73cc0a6f64b2c35908be2f6ec4ee
---

# Introduction

The introduction of the container feature into the RouterOS made it possible to run all kinds of servers for all sorts of tasks inside the router. This is especially relevant for people, who want to reduce the number of devices in their network. Instead of running a server on a separate device/machine, why not run it inside the router?

Radius is short for Remote Authentication Dial-In User Service. RouterOS has a RADIUS client feature supported that can authenticate for HotSpot, PPP, PPPoE, PPTP, L2TP, and ISDN connections. Basically, this feature allows you to connect RouterOS to a Radius Server, and then, utilize the user database from the server for client authentication.

In our example, we will showcase **freeradius/freeradius-server** image installation.

# Summary

**Sub-menu:** `/container`

**note**: **container** package is required.

Make sure to study our container guide before proceeding with the configuration. Make sure to check the disclaimer and requirements sections to understand all the risks and necessary steps you might be required to do.

At the time, when the guide was published, the image was available for linux/**amd64** OS/architecture **only** (usable by CHR and x86 devices). For arm64 devices you will need to make your own container from FreeRADIUS source. For both arm64 and arm32 devices, you can also use **freeradius/freeradius-server-dev** image **at your own risk** (as it is "experimental/development" version of the image).

To help you set up a CHR in a Virtual Box, please check our youtube tutorial, or Make your own x86 router.

This guide demonstrates a basic example! The tests were performed in a local environment! This guide is meant for basic RADIUS "testing" purposes! Not all "freeradius" features were tested!

# Configuration

## Container mode

Enable container mode:

You will need to confirm the device-mode with a cold reboot if using the container on X86.

## Networking

Add veth interface for the container:

Create a bridge for the container, assign an IP network to it, and add veth to the bridge:

Setup NAT for outgoing traffic if required:

## Getting image

To simplify the configuration, we will get the image from an external library but you can also import it via the .tar file.

Make sure that you have "Registry URL" set accordingly, limit RAM usage (if necessary), and set up a directory for the image:

Pull image with the help of the command:

where `cmd="-X"` enables debug logging (per the "freeradius" documentation).

After running the command, RouterOS should start "extracting" the package. Check "File System" for newly created folders and monitor container status with the command `/container/print`.

## Starting the container

After you make sure that the container has been added and the status changed to `status=stopped` after using `/container/print` → you can initiate it:

## Altering the server's configuration files

To access the server's configuration files (**clients.conf** and **authorize**), we will need to use SFTP (file transfer over SSH) protocol, so make sure that SSH service is enabled.

Open your command terminal ("CMD", as Administrator, for Windows users, or "Linux Shell or Command Terminal" for Linux users) and navigate it to the directory where you want to download the configuration files. For example, to "radius" folder on your "Desktop":

Initiate SFTP to the device's IP address:

Go to the server's configuration file folder (use `dir` or `ls` command to see the content of the folder you are in and `cd` command to go to the folder of our choice).

The first file, "clients.conf" allows you to define RADIUS clients. Per the "freeradius" documentation, it should be under the "/etc/freeradius" directory...so, navigate there and use `get` command to download it:

Open "**clients.conf**" via your preferred text editor (notepad or any other). You can study the file to see all the options that you have (additionally, check freeradius.org). This example shows a basic setup, so, we will just overwrite the whole file with the lines shown below:

where we indicate, that our radius client can connect using any possible IP address (**ipaddr=0.0.0.0/0** ensures that, but you also can change it to the actual ip address/mask of your radius client if you require to do so) and that our secret is "client_password" (you can change it to any other secret).

Save the file/overwrite it.

The second file, "authorize" allows you to set up users. Per the "freeradius" documentation, it should be under "/etc/freeradius/mods-config/files". Go there and `get` the file:

Open "**authorize**" via your preferred text editor (notepad or any other). This example shows a basic setup, so, we will just uncomment (remove the "#" symbol from) the line shown below (leave the rest of the configuration/lines as they are):

which creates a username "bob" and sets the password to "hello" (you can change the username and password).

Save the file/overwrite it.

Upload both files back/overwrite the default files with the help of the `put` command:

Restart the container:

Make sure to wait for the container to stop (`status=stopped` should be shown after using `/container/print` command) before initiating it again.

# Result verification

In RouterOS, add a new RADIUS client configuration:

,where the `address` is the IP address of the veth3 interface, `secret` is the secret that we configured in the **clients.conf** file and `service` is the allowed service that you wish to use.

Allow "login" with RADIUS users via the command:

We have allowed the "login" service for the RADIUS and we can test it using ssh/winbox/webfig connection. For SSH test, issue the command (where you need to indicate the device's management IP and input bob's password "hello" after):

You should be able to verify, that the terminal user changed from "admin@MikroTik" to "bob@MikroTik":

If you issue the command `/user/active/print`:

you will be able to verify, that a new user "bob" is "active" and has a flag "R" assigned, which indicates it is a RADIUS user.
