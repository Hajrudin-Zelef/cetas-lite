---
id: collect-261001-general-networking/general-networking/manual-how-tos-ipsec-swanctl-rw-ikev2-eap-mschapv2-html-81aa2fcc-4
title: "Add IPv4 route"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-ipsec-swanctl-rw-ikev2-eap-mschapv2-html-81aa2fcc.md
source_anchor: ""
source_lines: [613, 616]
sha256: dc5bd70c69183317fb522e219490eeab2138a5ed4330b9c1a287f565aafda0db
---

# Add IPv4 route

- Use tcpdump on the OPNsense to look for incoming packets on port 500 and port 4500 when you connect your VPN client. If you cannot see any, your firewall blocks them, or the remote client cannot send them due to a remote firewall. There could also be a wrong IP Address the packets are sent to.
- If there are packets received, but no packets sent, look into the VPN log files.
- Check /var/logs/ipsec/latest.log or for the connection being processed. Most of the time you can see errors in there you can search on the internet.
- The easiest tool to troubleshoot the connection is the Android StrongSwan Client or the Windows NCP Secure Entry Client. They have powerful inbuild logging so you can check both sides of the connection. In IPsec, you need the log of the server and the client to find the true cause of a connection error.
