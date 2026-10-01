---
id: collect-261001-meraki/meraki/questions-81045-cisco-meraki-versus-cloud-based-wlc-deployment-69cf8d5c
title: "Cisco Meraki versus Cloud-Based WLC Deployment"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-81045-cisco-meraki-versus-cloud-based-wlc-deployment-69cf8d5c.md
source_anchor: ""
source_lines: [1, 16]
sha256: 71f3509752d05495a2fad3e4bf129e24073924fe889713eb3bbf767bff9aa7a5
---

# Cisco Meraki versus Cloud-Based WLC Deployment

*Score : 0 | Source : https://networkengineering.stackexchange.com/questions/81045/cisco-meraki-versus-cloud-based-wlc-deployment*

A networking course that I follow mentions that we can centrally manage our APs in the cloud and that Cisco Meraki is a popular solution for that.
But then again, we can also have a Cloud-Based WLC deployment where we can, again, centrally manage our APs in the cloud.
Is there any major difference between these 2 services? When would I want to pick one over another?

---

### Reponse (acceptee) — score 3

You can make a similar comparison of any cloud managed service vs. maintaining your own server.
With the managed service (Meraki) you connect your APs, configure your policies, and youre done.
With your own WLC, you have to maintain the server, configure backups, install the WLC, manage upgrades, patches, etc. You also have to manage security for your WLC.
Usually, it comes down to price vs convenience. Every network is different, and everyone has different budgets.
