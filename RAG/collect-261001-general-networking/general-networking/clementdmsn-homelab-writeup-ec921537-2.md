---
id: collect-261001-general-networking/general-networking/clementdmsn-homelab-writeup-ec921537-2
title: "clementdmsn-homelab-writeup-ec921537"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/clementdmsn-homelab-writeup-ec921537.md
source_anchor: ""
source_lines: [188, 303]
sha256: ca2205a3e9014543555e6db5920d159e2c9da6474e459abcf3c31e8678b99122
---

# clementdmsn-homelab-writeup-ec921537

| Network | Action | Rule | Source | Destination | Protocol | Ports | 
|---|---|---|---|---|---|---|
| `SERVERS` | BLOCK | `BLOCK_MGMT_FROM_SERVERS` | `SERVERS net` | `This Firewall` | TCP | `MGMT_PORTS` | 
| `SERVERS` | PASS | `ALLOW_SERVERS_WEB_OUT` | `SERVERS net` | `*` | TCP | `WEB_PORTS` | 
| `SERVERS` | PASS | `ALLOW_SERVERS_DNS_TO_FORWARD` | `SERVERS net` | `This Firewall` | UDP | `DNS_PORT` | 
| `SERVERS` | PASS | `ALLOW_NTP_TO_FIREWALL` | `SERVERS net` | `This Firewall` | UDP | `NTP_PORT` | 

| Network | Action | Rule | Source | Destination | Protocol | Ports | 
|---|---|---|---|---|---|---|
| `DMZ` | BLOCK | `BLOCK_MGMT_FROM_DMZ` | `DMZ net` | `This Firewall` | TCP | `MGMT_PORTS` | 
| `DMZ` | BLOCK | `BLOCK_ACCESS_TO_ALL NETWORKS` | `DMZ net` | `ADMIN net` ,`CLIENTS net` ,`SERVERS net` | `*` | `*` | 
| `DMZ` | PASS | `ALLOW_DMZ_DNS_TO_FW` | `DMZ net` | `This Firewall` | UDP | `DNS_PORT` | 
| `DMZ` | PASS | `ALLOW_NTP_TO_FIREWALL` | `DMZ net` | `This Firewall` | UDP | `NTP_PORT` | 

Several infrastructure services are deployed inside the SERVERS network.

The domain controller provides:

- authentication
- DNS services
- DHCP services
- time services for domain-joined systems

Server configuration:

- Hostname: **SRV-AD**
- IP address: **192.168.20.10**
- Domain name: home.arpa

Windows client machines located in the CLIENTS network are joined to this domain.

DNS and DHCP are managed by the domain controller.

Because DHCP requests do not cross network boundaries by default, the firewall is configured as a **DHCP relay agent**.

DHCP requests originating from the CLIENTS network are forwarded to the domain controller located in the SERVERS network.

Time synchronization was split according to system role. Domain-joined Windows clients obtain time from the Active Directory hierarchy through the domain controller, while non-domain or infrastructure systems use the firewall as an NTP source.

The DMZ network is used to host services intended to be reachable from outside the internal infrastructure.

These services are isolated from internal networks through firewall policies.

Because these services use private internal addresses, port forwarding is required on OPNsense to make selected DMZ hosts reachable from the internet. Inbound WAN traffic is translated and forwarded only to explicitly published services, such as HTTP and HTTPS, while all other unsolicited access remains blocked by default.

A dedicated Linux system is deployed in the SERVERS network to host backup services.

This server is intended to store backups of critical infrastructure machines.

Networking inside the virtualization platform is implemented using a VLAN-aware bridge.

This bridge carries VLAN-tagged traffic between virtual machines and the firewall.

Each virtual machine is connected to the appropriate VLAN corresponding to its network zone.

The firewall has one interface per VLAN and performs routing between network segments.

The first version of the lab used separate virtual bridges to isolate each network. While this provided basic separation, it did not model segmentation over a shared 802.1Q trunk. The design was therefore migrated to VLANs.

A VLAN-aware bridge was introduced in Proxmox, allowing OPNsense to route between tagged networks over a single trunk while preserving isolation at Layer 2 through VLAN tagging.

This change aligned the lab more closely with how segmented networks are commonly deployed on managed switching infrastructure.

A later evolution of the lab introduced a WireGuard VPN service hosted on OPNsense. This addition provided a controlled remote administration path without exposing management interfaces directly to less trusted networks.

VPN clients are restricted to the ADMIN network, which preserves the separation between administrative access and the rest of the infrastructure while more closely reflecting real-world remote management practices.

Several virtual machines were deployed to validate the network configuration.

Test machines were placed in each subnet to verify:

- IP address assignment
- DNS resolution
- connectivity between network zones
- domain join functionality

A temporary ICMP firewall rule was enabled during testing to validate connectivity and isolate routing versus policy issues. It was removed after verification so that the final ruleset remained consistent with the intended security model.

A Windows client machine was successfully joined to the Active Directory domain and received network configuration through DHCP.

*DNS resolution from client :*


*Client network identity :*


Several technical issues were encountered during the deployment process.

During Windows installation, the virtual disk was not detected.

The VirtIO storage drivers had to be manually loaded in the installer.

The Windows installer initially failed to detect the network interface.

Changing the adapter model to **e1000e** resolved the issue.

Accessing the firewall web interface over HTTPS initially failed in Microsoft Edge.

Switching to Firefox resolved the problem.

After migrating the ADMIN network to VLAN segmentation, DHCP requests stopped reaching the domain controller.

Restarting the DHCP relay service restored normal operation.

Multicast discovery traffic such as SSDP and mDNS was explicitly blocked at the firewall boundary. These protocols are useful on flat local networks for device discovery, but they were not required in this lab and would undermine segmentation by allowing unnecessary cross-zone service discovery.

Several improvements are planned for future iterations of the lab.

Possible extensions include:

- centralized logging infrastructure
- monitoring and alerting systems
- additional services hosted in the DMZ
- automated backup workflows

These additions will allow the lab to more closely resemble a complete enterprise infrastructure.
