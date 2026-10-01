---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad-1
title: "docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "cyber"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad.md
source_anchor: ""
source_lines: [1, 63]
sha256: bb5ae7739d4bf62810073eef714532bc4eb2d6fe311bfd65ed3216b01c4271b0
---

# docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad

Network address translation is the process of mapping one [Internet Protocol (IP) address] to another by modifying the header of IP packets while they are in transit across a router. As part of this technique, NAT settings can expose only one IP address for an entire network to the outside world, effectively masking the entire internal network and increasing security. Network address translation is widely used in remote-access scenarios because it conserves addresses while also increasing security. This improves security while also reducing the number of IP addresses required by a business.

Network Address Translation (NAT) is a method of separating external and internal networks (WANs and LANs) and sharing an external IP address among clients on the internal network. NAT can be used on both IPv4 and IPv6 networks. Network Prefix Translation is also available for IPv6.

In addition the its NAT features, OPNsense also provides next-generation firewall capabilities such as web control and application control. This is provided by an external tool called Zenarmor.

Zenarmor NGFW Plug-in for OPNsense is one of the most popular OPNsense plug-ins and allows you to easily upgrade your firewall to a Next Generation Firewall in seconds. NG Firewalls empower you to combat modern-day cyber attacks that are becoming more sophisticated every day.

Some of the capabilities are layer-7 application/user aware blocking, granular filtering policies, commercial-grade web filtering utilizing cloud-delivered AI-based Threat Intelligence, parental controls, and the industry's best network analytics and reporting.

Zenarmor Free Edition is available at no cost for all OPNsense users.

The majority of the options below make use of three distinct addresses: the source, destination, and redirect address. These addresses will be used for the following purposes:

| Address | Description | 
|---|---|
| Source | From where the traffic is coming. This is frequently left on "any". | 
| Destination | Where the traffic is going. This is typically your external IP address for incoming traffic from the outside world. | 
| Redirect | Where traffic should be rerouted | 

Disabling `pf` disables `NAT` on OPNsense.

- 
**BINAT** : NAT typically operates in only one direction. But, if your networks are of equal size, you can also use bidirectional BINAT. This can help to simplify your setup. You can only use regular NAT if your networks are not of equal size.
- 
**NAT reflection** : When a user on the internal network attempts to connect to a local server by using the external IP address rather than the internal one, NAT reflection can rewrite the request to use the internal IP address, avoiding a detour and applying rules designed for actual outside traffic.
- 
**Pool Options** : When there are multiple IPs to choose from, this option allows you to control which IP is used. The default, Round Robin, simply sends packets to one server after another. This option has no effect if you only have one external IP address.

OPNsense firewall provides the following types of NAT configurations:

1. 
Port Forwarding NAT (DNAT)
2. 
One-to-One NAT (1:1 NAT)
3. 
Outbound NAT (SNAT)

In this article, we will cover all these NAT configurations on OPNsense shortly and give the following real-world examples.

- 
Port forwarding configuration in OPNsense for a web server accessible from the Internet.
- 
Port forwarding configuration in OPNsense for ssh and RDP servers accessible by a specific IP
- 
Outbound NAT configuration in OPNsense for allowing specific local servers to access a remote service.

## Configure Port Forwarding (DNAT)

OPNsense Port Forwarding is a utility that facilitates the routing of incoming internet traffic from external sources to particular devices on your local network. It enables you to provide hosting services for websites or games, allowing them to be accessed outside. Any connections to the internal network from the Internet are blocked on the OPNsense firewall by default. When an internal system behind a firewall needs to be configured for remote access, port forwarding NAT should be configured. You may use the OPNsense port forwarding feature to allow certain services(ports) from the external network.

Port forwarding is a method that allows for the configuration of certain destination ports to always be directed to specific nodes. Port forwarding approach enables full IP masquerading while maintaining the ability for services to respond to incoming traffic.

Port forwarding is also known as "Destination NAT" or "DNAT." When multiple servers in a LAN share the same external IP address, any connection that is not initiated by one of the servers will fail because the firewall will not know where to send the traffic. This can be remedied by establishing port forwarding rules. For example, to make your organization's web server behind the firewall accessible from the internet, you must redirect HTTP(s) ports (80/443) to the server.

To configure the port forwarding in OPNsense you may navigate to **Firewall** → **NAT** → **Port Forward**. An overview of port forwarding rules can be found here.

**Figure 1.** *Port forwarding configuration in OPNsense*

To add new port forwarding rules, you may click the `+` button in the upper right corner.

The following fields are available when adding a port forwarding rule on OPNsense:

