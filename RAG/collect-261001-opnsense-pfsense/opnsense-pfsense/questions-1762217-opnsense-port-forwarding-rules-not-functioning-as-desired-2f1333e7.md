---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1762217-opnsense-port-forwarding-rules-not-functioning-as-desired-2f1333e7
title: "OPNsense port forwarding rules not functioning as desired"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1762217-opnsense-port-forwarding-rules-not-functioning-as-desired-2f1333e7.md
source_anchor: ""
source_lines: [1, 34]
sha256: b0bacc2c6f592dfc23b81aaebd7f1380a28672b2f48b79fe760299c663a46d64
---

# OPNsense port forwarding rules not functioning as desired

*Score : 1 | Source : https://superuser.com/questions/1762217/opnsense-port-forwarding-rules-not-functioning-as-desired*

I've got a virtual network setup inside proxmox. I've installed OPNsense as the router to connect traffic from my regular network to the virtual network. The LAN side of OPNsense is on a bridge network in proxmox assigned to the 10.0.42.0/24 network with a statically assigned IP. The WAN side is connected to the virtual bridge bound to the physical interface on the proxmox server. The WAN interface is assigned an IP via DHCP from my internet router.
The virtual network is able to receive IPs from OPNsense with no problem. Those same systems are able to access the internet through OPNsense without problems. I can access the OPNsense web interface from my desktop with no problems.
Inside the virtual network I've got portainer up and running. I'm not able to access the management page for this portainer system from my desktop. I've confirmed that the portainer address was accessible from inside the network.
I've tried adding port forwarding rules on OPNsense. I've used the following pages to try and get a port forwarding rule in place. Most of the pages have instructions very similar to each other
When I attempt to access the portainer page from the desktop, I see the following in the logs of the OPNsense firewall:
Source: 10.0.42.11:9443
Destination: [desktop IP]:51918
Proto: tcp
Label: Default deny / state violation rule
Firewall rule I've tried: Interface: WAN TCP/IP Version: IPv4 Protocol: TCP Source: any Source port range: any Destination: WAN Address Destination port range: any Redirect target IP: portainer IP Redirect target port: 9443 NAT Reflection: tried both enable and disable Filter rule association: Tried both Pass and Add associated filter rule
I've also tried using the reverse of the above with no change to the errors in the firewall logs.
Also tried adding WAN rules. I've used portainer as the source IP/port and WAN Address as the destination with all ports. I've also tried adding in a reverse rule as well.
If I disable the firewall completely, I'm able to access the portainer web interface, but the virtual machines are no longer able to access the internet. This is due to NAT being disabled along with the firewall.
EDIT
The configuration for the networks/devices are as follows:
Modem->(public IP) Netgear router (192.168.x.x)->TP-link Switch->Desktop
Modem->(public IP) Netgear router (192.168.x.x)->TP-link Switch->Proxmox->(192.168.x.x) OPNsense (10.0.42.x)->Portainer
I've created a second bridged interface in proxmox that is not bound to the physical interface. That bridge is the gateway for packets headed to proxmox to be forwarded to OPNsense for accessing the virtual network. Achieved by adding a new ip table and rule on proxmox. I've added the bridge bound to the physical interface as the WAN interface in OPNsense. The unbound bridge is assigned to the LAN interface in OPNsense. On the windows desktop, I've added a route to point 10.0.42.x traffic to the IP of the virtual network bridge on proxmox.
This all works to allow me to access the management webpage for OPNsense, due to it's anti-lockout rule passing all traffic to port 80 on itself. I've tried mimicking that rule to pass all traffic from my desktop, but still getting blocked by the firewall based on log entries.
Ultimate goal is to have a virtual network separate from the LAN that I can play around in as the beginnings of a homelab. Inside the homelab, I'm planning to install truenas, jellyfin, and bind among other tools. Most of these will need to be accessible from the physical LAN at some point.

---

### Reponse — score 0

I've figured out how to solve the problem for portainer. I logged into portainer via the console in proxmox. Then I added a route into the routing table to pass traffic bound for 192.168.0.0/16 via the 10.0.42.x port on proxmox.
ip route add 192.168.0.0/16 via 10.0.42.x
Once that route was in place, access to portainer's web interface was successful from my desktop. This means that for every server inside the virtual LAN I want to access, a route needs to be in place for the "WAN" network.
I've tested adding the same route on a fedora VM inside the virtual network. My desktop is able to ping and connect to that VM now as well. This is without a firewall rule in place. So it looks like adding that route bypasses the firewall completely. In essence, turning the proxmox server into a router as well. I suspect OPNsense needs to remain in place to provide NAT so the VMs can access the internet.
I'm sure there's a way to achieve this same result without needing to add the route on each individual server. That will probably require an enterprise grade router instead of a firewall appliance like OPNsense or my home router. I'll also need to figure out where to add the route command into the startup files so that it'll persist upon reboot.
