---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2-2
title: "how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2.md
source_anchor: ""
source_lines: [18, 68]
sha256: c7392741b8496840f3d4ba8e7a7f94cb5d3b69903479b9a7cc5d6c97c5fd0a87
---

# how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2

By default, OPNsense creates a LAN network so for this example, you will need to create a DMZ and an APP network. You could create the networks as VLANs using the spare OPT interface(s) on your firewall appliance or use the LAN interface as the parent interface (Note: I have read that it may be better to put the VLANs on the OPT interfaces since you have to be careful with the firewall rules in regards to tagged and untagged network traffic on the same interface). In an effort to keep this guide at a reasonable length, I am not going to describe the full details for creating VLANs in OPNsense or on your network switch(es). I have written several guides about VLANs if you need more assistance.
Create Firewall Rule to Allow Access to Nginx Proxy Manager from LAN
Access to the Nginx Proxy Manager needs to be allowed from the LAN (and any other network which needs access to the apps/services). Go to the “Firewall > Rules > [LAN]” page, and click on the “+” button to add a new rule.
In rule below substitute the “LAN” network for the appropriate network which you are using. You will need to decide which IP address you plan to use for the reverse proxy and enter it into in the “Destination” field. It is ok to set up the rules before you set up the reverse proxy. For the “Destination Port”, I am using an alias called NPMPorts which includes ports 80, 81, and 443. Port 81 is used as the management port for the Nginx Proxy Manager.
| Option | Value | 
|---|---|
| Action | Pass | 
| Interface | LAN | 
| TCP/IP Version | IPv4+IPv6 (IPv6 is optional) | 
| Protocol | TCP | 
| Source | LAN net | 
| Source Port | any | 
| Destination | 192.168.2.50 (or use an alias which may include the IPv6 address) | 
| Destination Port | NPMPorts (an alias for port 80, 81, and 443) | 
| Description | Allow access to Nginx Proxy Manager administration | 
If the LAN is used as your management network, you will need to include port 81, but if you are allowing access from other networks such as an IOT network, you should omit port 81 to block access to the management interface. Also, you may omit port 80 if you have everything hosted via HTTPS unless you need to have port 80 redirected to port 443 (for HTTP to HTTPS redirection).
Create Firewall Rule to Allow Access to App Server from Nginx Proxy Manager
The Nginx Proxy Manager needs to have access to the app server(s) in the APP network so a rule needs to be created to allow that access. I am going to assume that every app/service that is placed into the APP network is allowed to be access by the reverse proxy. It is simple to create a rule that uses “APP net” as the destination rather than create an alias with every IP address of every service in the APP network. Using “APP net” alleviates the need to keep an alias up to date every time you spin up a new service. If all services in the APP network needs to be proxied, then it no less secure than specifying each individual IP address. If you wish to nitpick, there is a small chance that you will forget about using “APP net” and deploy a service in the APP network that you do not want proxied. The risk is minimal since no traffic will be redirected to that service unless it is first set up in the proxy manager.
On the “Firewall > Rules > [DMZ]” page and click on the “+” button to add a new rule.
| Option | Value | 
|---|---|
| Action | Pass | 
| Interface | DMZ | 
| TCP/IP Version | IPv4+IPv6 (IPv6 is optional) | 
| Protocol | TCP | 
| Source | 192.168.2.50 (or use an alias which may include the IPv6 address) | 
| Source Port | any | 
| Destination | APP net | 
| Destination Port | WebServerPorts (an alias for port 80 and 443) | 
| Description | Allow access to app servers | 
Create NAT Port Forward Rule to Allow External Network Access
A reverse proxy can be useful even if it is only used internally for all of your self-hosted services, but you will need to allow external WAN access if you wish for your services to be publicly accessible. On the “Firewall > NAT > Port Forward” page, add the following rule.
| Option | Value | 
|---|---|
| Interface | WAN | 
| TCP/IP Version | IPv4+IPv6 (IPv6 is optional) | 
| Protocol | TCP | 
| Source | any | 
| Source Port | any | 
| Destination | WAN address | 
| Destination Port | WebServerPorts (an alias for port 80 and 443) | 
| Redirect target IP | 192.168.2.50 (or use an alias which may include the IPv6 address) | 
| Redirect target port | WebServerPorts (an alias for port 80 and 443) | 
| Description | Allow external access to Nginx Proxy Manager | 
| Filter rule association | Add associated filter rule (or Pass) | 
If you are not going to create Let’s Encrypt certificates using a DNS challenge, you will need to open port 80 so Let’s Encrypt can create the certificate. Otherwise, you only need to open port 443 to the public. In the example above, I am using an alias with both port 80 and 443 since creating certificates via HTTP is the easiest method. You can still host all of your services on port 443/HTTPS and port 80 will only be used for generating and renewing certificates as well as for redirection of HTTP to HTTPS. If your ISP blocks port 80, then you must use a DNS challenge to generate certificates.
Important: If you expose services publicly, security needs to be considered. Some advice you may find is put everything behind a VPN. While that is sound advice, it is not quite as convenient especially if you have others using your service such as family members/relatives/friends (you have to explain to them how to set up a VPN or set it up for them and expect them to actually use it). VPNs can have their own vulnerabilities even though it is often proclaimed as the ultimate security measure. You should also consider using firewall rules and other means to secure your services even if you are using a VPN. Keep in mind that by exposing the reverse proxy rather than the service itself, security is already better than directly exposing the service to the Internet. You could further improve security by using a Cloudflare proxy and only allowing Cloudflare to access your hosted services so external users only see the Cloudflare proxy addresses.
Use Split DNS to Resolve Hostnames to the Reverse Proxy
After configuring your proxy, you will most likely want to use the Unbound DNS override functionality in OPNsense to utilize the split DNS capability so you can redirect clients on your network to the local IP address of the Nginx Proxy Manager rather than your external WAN address. Since I am using a hostname that will be used for external access, the WAN IP address will be used as the IP for that hostname. I found that using the external IP address from within the network does not work properly (perhaps it could be made to work via NAT reflection/outbound NAT rule), but that is ok since a DNS override will resolve the issue.
Because I plan to put my services behind Cloudflare, I would need to use split DNS anyway so that I do not lose access to my local services if my Internet is down. If you set up a Cloudflare proxy (see below), the Cloudflare IPs will be used which means your local traffic will be routed through Cloudflare’s servers and back to your local network.
If you are only using the reverse proxy for local services that you are not exposing to the world and you are using a hostname which is not a real host on your network, you will need to use a DNS override so the hostname can be resolved. You may want to use a hostname that is different than any actual hosts on your network as described in the Determine Hostnames for the Proxy Host and Services section.
