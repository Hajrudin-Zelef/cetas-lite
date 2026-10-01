---
id: collect-261001-general-networking/general-networking/questions-901172-clients-not-authenticating-with-domain-controller-via-wifi-serv-872ae510
title: "questions-901172-clients-not-authenticating-with-domain-controller-via-wifi-serv-872ae510"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-901172-clients-not-authenticating-with-domain-controller-via-wifi-serv-872ae510.md
source_anchor: ""
source_lines: [1, 6]
sha256: 1ffe83067dc93aea0325b64b129d497f6e2c2513d3ca057e39e3361ed5f0345a
---

# questions-901172-clients-not-authenticating-with-domain-controller-via-wifi-serv-872ae510

I run a bit of a complicated setup, let me give you a quick rundown:
Local pfSense Firewall -> IPSec Tunnel to AWS -> Server 2016 DC.
pfSense runs DNS in the local network, forwards domain queries to the DC via Domain Override.
If I am connected via ethernet, everything works fine. Soon as I use WiFi (UniFi), I can't run gpupdate or print. When running gpupdate I get the error that the computer name can not be resolved.
When I reboot the machine, I get a Netlogon ID 5719, the computer can't establish a safe connection with the DC.
Why does this happen, and why does it work via ethernet and not WiFi?
