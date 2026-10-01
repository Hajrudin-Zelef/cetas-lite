---
id: collect-261001-general-networking/general-networking/can-diagnose-debug-application-be-filtered-community
title: "can-diagnose-debug-application-be-filtered-community"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/can-diagnose-debug-application-be-filtered-community.md
source_anchor: ""
source_lines: [1, 4]
sha256: bf29285835fc685b3351efd03f063f749837d687f65a8c77125fa6d21b331503
---

# can-diagnose-debug-application-be-filtered-community

I am trying to debug some ssl-vpn connection stuff. If I run "diag debug application sslvpn -1" it generates a lot of debug lines. Downloading the output and filtering through it to find what I need is not fun.


Is there a way to filter this by the source IP of the remote VPN client? Or by some sort of VPN session ID? Or something so that I can focus on troubleshooting a single user without having to wade through all the other connection data?
