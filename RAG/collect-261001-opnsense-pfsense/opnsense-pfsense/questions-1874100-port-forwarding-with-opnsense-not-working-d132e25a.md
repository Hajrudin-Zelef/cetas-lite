---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1874100-port-forwarding-with-opnsense-not-working-d132e25a
title: "Port Forwarding with OpnSense not working"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1874100-port-forwarding-with-opnsense-not-working-d132e25a.md
source_anchor: ""
source_lines: [1, 20]
sha256: 71edb96883aef47d7c5a489351f7618a57c3892f11cefe8c2371f621662976ff
---

# Port Forwarding with OpnSense not working

*Score : 0 | Source : https://superuser.com/questions/1874100/port-forwarding-with-opnsense-not-working*

I’m trying to port forward 25292/tcp to 192.168.1.111/32. Both source and destination ports are identical. The device is an Unraid 7 server running Docker. The target container is in bridge mode with port 25292/tcp allocated.
In OPNsense, I have created a NAT rule along with the corresponding firewall rule. However, when I try to access http://MY-WAN-IP:25292 from another network using curl, I get the following error:
user@device ~ % curl http://MY-WAN-IP:25292 --verbose   
*   Trying 0.0.0.0:25292...
* connect to 0.0.0.0 port 25292 from 192.168.2.49 port 57851 failed: Network is unreachable
* Failed to connect to 0.0.0.0 port 25292 after 4018 ms: Couldn't connect to server
* Closing connection
curl: (7) Failed to connect to 0.0.0.0 port 25292 after 4018 ms: Couldn't connect to server
Edit: I'm not behind a second NAT (DS-Lite etc).
Screenshots:

---

### Reponse (acceptee) — score -1

I don’t know what caused the problem, but it no longer exists. I didn’t change anything – it just started working again.
