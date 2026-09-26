---
id: collect-260926-mikrotik/mikrotik/questions-1019156-ansible-error-with-auto-provision-mikrotik-router-75802b02
title: "questions-1019156-ansible-error-with-auto-provision-mikrotik-router-75802b02"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1019156-ansible-error-with-auto-provision-mikrotik-router-75802b02.md
source_anchor: ""
source_lines: [1, 5]
sha256: 235a79ceb938bc12f02dcb78532b2825af65b5402fb78737a9e938eddf19630a
---

# questions-1019156-ansible-error-with-auto-provision-mikrotik-router-75802b02

I have a problem that I cannot understand how to solve ...
From my ansible central machine I have written a routine for Mikrotik, to execute a series of configuration commands and thus auto provision my RBs, but it returns the following error, and I just don't know how to pass the OS detection configuration to me asks, do I have to pass it to it as a variable in the Ansible configuration file? About Ansible's own script? I just couldn't find much info about it.
Error :
fatal: [10.0.21.10]: FAILED! => {"msg": "Unable to automatically determine host network os. Please manually configure ansible_network_os value for this host"}
If someone has already worked on Ansible + Mikrotik automation I would really appreciate the clarification!
