---
id: collect-261001-general-networking/general-networking/questions-1121945-i-am-unable-to-access-services-on-the-wan-ip-from-within-the-n-36b278f4
title: "I am unable to access services on the WAN IP from within the network"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1121945-i-am-unable-to-access-services-on-the-wan-ip-from-within-the-n-36b278f4.md
source_anchor: ""
source_lines: [1, 11]
sha256: a4fe3936d7e024eb1ed8e31830a2942949d6718ef4b8060596953cf57610ee81
---

# I am unable to access services on the WAN IP from within the network

*Score : 2 | Source : https://serverfault.com/questions/1121945/i-am-unable-to-access-services-on-the-wan-ip-from-within-the-network*

Normally, this would not be a desired configuration, but I am setting up a NextCloud server, and to validate the domain, it requires that it be able to access it through the public IP address. No matter what I do, I cannot get this to work. It specifically needs port 443, but I cannot reach port 80, 8080, nor 443 from inside the firewall (OPNSense), when using the FQDN. DNS queries are resolving properly with the WAN IP, I have opened the ports outgoing in order to let the server bypass the transparent proxy, and have even port forwarded port 443 outgoing for the server IP to push it past the proxy, but nothing works. If I try to access these ports from outside the firewall (from my cell phone), I have no trouble at all. I know this is unusual, but is there any way to make this work? Someone has to have been able to get NextCloud working at some point, right?

---

### Reponse — score 4

In order to access other internal LAN resources within your network using your external IP address through OPNSense, you need to enable the NAT reflection feature. It will rewrite such requests so that they use the internal IP in order to avoid taking a detour and applying rules meant for actual outside traffic. More information on NAT reflection can be found here.
