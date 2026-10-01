---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/stoicturk182-blog-blob-head-content-posts-self-20hosted-20unifi-20controller-20s-8d088045-2
title: "stoicturk182-blog-blob-head-content-posts-self-20hosted-20unifi-20controller-20s-8d088045"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/stoicturk182-blog-blob-head-content-posts-self-20hosted-20unifi-20controller-20s-8d088045.md
source_anchor: ""
source_lines: [114, 133]
sha256: 5288972d66543473ccebc08926dc9ce77a7af61a61a440ff603d27853177b81c
---

# stoicturk182-blog-blob-head-content-posts-self-20hosted-20unifi-20controller-20s-8d088045

Superuser privileges are required to read both.
| Symptom | Resolution | 
|---|---|
| Service starts then immediately stops | Check server.log for a MongoDB version rejection or a Java version mismatch. Confirm the installed MongoDB against the version table above. | 
| MongoDB fails to start with no clear error | Confirm AVX support on the CPU if running MongoDB 5.0 or later. | 
| Port 8080 reported as in use | Another service holds the port. Identify it with sudo ss -tlnp \| grep 8080 , or change the port in/usr/lib/unifi/data/system.properties and restart. | 
| Web interface unreachable | Confirm the service is running, then confirm 8443 is permitted by the host firewall and any upstream firewall. | 
| Devices not appearing for adoption | Layer 3 separation between device and controller. Use set-inform , DHCP option 43, or aunifi DNS record. | 
| Slow start or intermittent service failure on a VM | Entropy starvation on headless virtual machines. Installing haveged is the documented workaround. | 
| Application will not bind to a chosen port | Ports below 1024 are unavailable, as the service does not run as root. | 
Ubiquiti's newer self-hosting path is UniFi OS Server, which runs the UniFi applications in containers using Podman via an official installation script, rather than the APT package and MongoDB pairing described here. It targets Debian 12 and later and Ubuntu 22.04 and later on x64 hardware, and provides the multi-application experience of a Cloud Gateway on self-managed hardware.
The APT method above remains valid for Network-only deployments and is the lighter option where the host is already running other services. For new builds intended to grow beyond Network alone, UniFi OS Server is the more likely long-term direction, though at the time of writing the APT package continues to be published and documented by Ubiquiti.
- Updating and Installing Self-Hosted UniFi Network Servers (Linux): https://help.ui.com/hc/en-us/articles/220066768-Updating-and-Installing-Self-Hosted-UniFi-Network-Servers-Linux
- Self-Hosting a UniFi Network Server: https://help.ui.com/hc/en-us/articles/360012282453-Self-Hosting-a-UniFi-Network-Server
- Required Ports Reference: https://help.ui.com/hc/en-us/articles/218506997-Required-Ports-Reference
- UniFi Network Application release notes: https://community.ui.com/releases
- UniFi downloads: https://ui.com/download
- MongoDB installation on Ubuntu: https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-ubuntu/
- MongoDB production notes, CPU requirements: https://www.mongodb.com/docs/manual/administration/production-notes/
- Debian sudo documentation: https://wiki.debian.org/sudo
