---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-caddy-reverse-proxy-with-lets-encrypt-and-crowdsec-using-opnsense-fb0acb0a-1
title: "caddy.service"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agents"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-caddy-reverse-proxy-with-lets-encrypt-and-crowdsec-using-opnsense--fb0acb0a.md
source_anchor: ""
source_lines: [1, 60]
sha256: d9cfbd03f47d44175a7fa7d67680a8217131d72343195d074729cb10c336c028
---

# caddy.service

Set Up a Caddy Reverse Proxy with Let's Encrypt and CrowdSec Using OPNsense LAPI
Table of Contents
Ever since setting up CrowdSec in OPNsense, I have had the desire to set up a CrowdSec bouncer on my reverse proxy to enable extra protection for apps/services hosted behind the proxy.
Even though I have CrowdSec protecting my entire network from crowd sourced malicious IPs as well as protecting my OPNsense web UI and SSH service, CrowdSec is also able to protect various apps/services running on your network by parsing logs, etc.
I was using Nginx Proxy Manager for my reverse proxy because I enjoyed the simplicity of the web UI. However, there is no official way to add CrowdSec support. You have to use a community fork which has CrowdSec integrated.
On top of that, I seen concerns that NPM may be more vulnerable than simply using Nginx directly as a reverse proxy since the NPM project does not always keep up to date when there are vulnerabilities found in Nginx.
I decided that I would try Caddy as my new reverse proxy due to its simplistic configuration and automatic SSL provisioning. While the default usage of Caddy is pretty simple, the process in this guide will be more involved because it requires manually building Caddy to include additional functionality. Also, I do not plan to use Docker so that adds a few steps (some of which only need to be done once).
In this guide, I will be demonstrating how to set up a Caddy reverse proxy server using Let’s Encrypt with DNS challenges as well as how to set up a CrowdSec multi-server configuration. So there is a lot to cover!
Before Getting Started
To keep this guide on topic, there are a number of assumptions I will make for the example scenario I have described in the introduction:
- 
You have an OPNsense installation with the CrowdSec plugin installed.
- 
SSH is set up on OPNsense or you can log in via the console to run various commands for CrowdSec.
- 
You have an existing service on your network with a static IP address that you can access successfully via HTTP (HTTPS may require some configuration changes to trust self-signed certificates in Caddy). This HTTP service will be placed behind the Caddy reverse proxy.
- 
I will be using a LXC in Proxmox with a static IP address on the DMZ network for the Caddy installation (but you may install Caddy on a VM or a dedicated system), and I will not be using Docker.
- 
I will be using Cloudflare as the domain registrar with an API key generated for “Edit DNS Zone” so certificates can be generated without needing to open up the firewall to allow ACME challenges, which is great when your reverse proxy instance is not directly connected to the Internet (if you are using the reverse proxy for internal network services but still desire valid certificates).
Prepare the OPNsense CrowdSec Configuration
Before setting up the Caddy reverse proxy, some settings for CrowdSec and firewall rules can be configured in OPNsense to prepare for a CrowdSec multi-server environment.
Update the Existing CrowdSec Plugin Configuration
The first thing you can do is change your CrowdSec plugin settings in OPNsense to allow other CrowdSec agents/bouncers to use the LAPI (Local API) on OPNsense. By default, the LAPI on OPNsense only listens on localhost (127.0.0.1).
In this example, I am setting the IP address to be on the LAN interface of 192.168.1.1. You may wish to put it on a different interface.
You will need to start/stop the CrowdSec plugin for changes to take effect.
Create API Key for Caddy CrowdSec Bouncer
To prepare for setting up the CrowdSec bouncer for Caddy in a later step, you will need an API key generated for the bouncer.
Log into your OPNsense system via SSH or the console and issue the following command to create an API key. You may use any name you wish in place of caddyDmz. I used that name since this will be for a Caddy instance on the DMZ network.
sudo cscli bouncers add caddyDmz
You should see the API key in the console output. Copy/paste this key until it is needed later.
API key for 'caddyDmz':
   kLI6ljt3B/zWsBvRxu68vP8rL/cNj3O0hgjJ6B2s7Yk
Please keep this key since you will not be able to retrieve it!
Add the Appropriate Firewall Rule(s)
You will need to add a couple of firewall rules if your OPNsense LAPI, Caddy server, and hosted apps/services are on different networks.
You may skip this entire section if they are all on the same network.
Allow Caddy Server Access to OPNsense LAPI
If Caddy is on another network/VLAN than the OPNsense LAPI which is on the LAN interface, such as the DMZ network for instance, you will need to create a firewall rule on that interface to allow access to 192.168.1.1 for TCP port 8080.
Go to the “Firewall > Rules > [DMZ]” page. In the example firewall rule below, a firewall alias called Caddy should contain the IP address of your Caddy reverse proxy such as 192.168.10.10.
| Action | TCP/IP Version | Protocol | Source | Destination | Dest Port | Description |  | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4 | TCP | Caddy | 192.168.1.1 | 8080 | Allow access to CrowdSec LAPI on OPNsense |  | 
Allow Caddy Server Access to App/Service
Additionally, if your hosted app or service is on a different network/VLAN, you will need a firewall rule to allow access. To potentially decrease the attack surface of hosting apps on a DMZ, you could expose only the reverse proxy directly on the DMZ while the app/service lives on another network protected by the OPNsense firewall. You will only need to open the port(s) required by the service.
For the example scenario, you may use the following firewall rule (while still on the “Firewall > Rules > [DMZ]” page) to allow access to the Homepage dashboard from the Caddy reverse proxy. Caddy is a firewall alias for the IP address of the Caddy reverse proxy such as 192.168.10.10 while HomepageServer is an alias for an IP address such as 192.168.20.10.
| Action | TCP/IP Version | Protocol | Source | Destination | Dest Port | Description |  | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4 | TCP | Caddy | HomepageServer | 3000 | Allow access to Homepage dashboard server |  | 
Allow Devices to Access the Caddy Server
Any devices which exist on other networks which need access to your apps/services behind the proxy will need to reach the Caddy reverse proxy via HTTPS. The rule below will allow devices in the USER network to access the apps/services to reach Caddy so you can access your apps/services.
| Action | TCP/IP Version | Protocol | Source | Destination | Dest Port | Description |  | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4 | TCP | USER net | Caddy | 443 (HTTPS) | Allow access to Caddy reverse proxy |  | 
If you have multiple networks which need access, you could create a firewall group and include multiple interfaces or you could make use of floating rules.
Unbound DNS Overrides
When using a reverse proxy, you will need to add Unbound DNS overrides. Unbound DNS overrides allow you to point all of your local hostnames to the IP address of the Caddy server so that you can access multiple apps/services behind the reverse proxy.
Normally you would want hostnames to uniquely refer to clients on your network but in the case of a reverse proxy, you will want all of the hostnames of the frontend of your apps/services to be using the reverse proxy IP address. The proxy will handle forwarding the requests to the proper backend server based on the hostnames.
Create an Unbound DNS override entry for the Caddy server hostname such as caddy-server for the appropriate domain name such as homenetworkguy.com.
For the aliases you may add all of the hostnames you will use for the servers hosted behind the reverse proxy. You should use a different hostname than the actual backend app/server.
