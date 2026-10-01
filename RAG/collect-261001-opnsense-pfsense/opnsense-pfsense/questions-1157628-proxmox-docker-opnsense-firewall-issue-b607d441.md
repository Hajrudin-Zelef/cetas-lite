---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1157628-proxmox-docker-opnsense-firewall-issue-b607d441
title: "Proxmox Docker OPNSense Firewall issue"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1157628-proxmox-docker-opnsense-firewall-issue-b607d441.md
source_anchor: ""
source_lines: [1, 20]
sha256: f1cacffe5e0044ee8ee4898b8742743f1c7c04dde4c2f1bac0de79dfdf6e4fc8
---

# Proxmox Docker OPNSense Firewall issue

*Score : 0 | Source : https://serverfault.com/questions/1157628/proxmox-docker-opnsense-firewall-issue*

I'm currently configuring my root server and I'm hitting a roadblock. My current setup looks like the following:
Now I try to install roundcube on the docker host and have it connect to the mailcow vm. But I only get
errors: <e30816ed> IMAP Error: Login failed for e@mail.com against mail.host.name from 10.0.10.10 (X-Forwarded-For: 1.2.3.4). Could not connect to mail.host.name:143: Connection refused in /var/www/html/program/lib/Roundcube/rcube_imap.php on line 211 (POST /?_task=login&_action=login)
The connection from the mailcow VM to the docker vm via IP works just fine. A netcat connection to the mailcow vm on port 143 is also working but using the mail.host.name results in a connection refused.
How do I configure OPNSense to allow for connections to mail.host.name from within my network or is there an alternative to perform some local dns resolution.
If I change the roundcube config to use the ip, the connection works but gets aborted because the ssl certificate does not match the hostname.
Thans for you help!

---

### Reponse (acceptee) — score 1

What you're looking for is split horizon DNS probably.
The simplest (non-scaling) way to solve this is to add mail.host.name to /etc/hosts, as it by default takes precedence over DNS for name resolution. That doesn't scale very well however.
mail.host.name
/etc/hosts
