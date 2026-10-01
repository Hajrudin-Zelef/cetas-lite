---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/pando85-homelab-blob-head-docs-deployment-unifi-md-67571ebd
title: "pando85-homelab-blob-head-docs-deployment-unifi-md-67571ebd"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/pando85-homelab-blob-head-docs-deployment-unifi-md-67571ebd.md
source_anchor: ""
source_lines: [1, 18]
sha256: 33ec6353ac4dd1f5b13def8eaadbaa014f4c2fd3b634297668d9333334460b6a
---

# pando85-homelab-blob-head-docs-deployment-unifi-md-67571ebd

- 
ServicesDNS->ResolverGeneral->Settings->Host Overrides :- host: unifi-controller domain: grigri ip_address: 192.168.193.2 description: Deployed on k8s.grigri in his own load balancer additional_names_for_this_host: - host: unifi domain: grigri description: Used for automatically adopt devices
» dog unifi-controller.grigri
A unifi-controller.grigri. 1h00m00s   192.168.193.2
» dog unifi
A unifi.grigri. 59m50s   192.168.193.2
- Factory reset. One of this:
  - 10 seconds button
  - ssh and run syswrapper.sh restore-default
- Connect to the DMZ network.
- Click adopt from unifi web interface
Note: For wireless mesh use same channel for both APs (disable Channel Optimization). If still doesn't work go to legacy UI and set up same Channel and width in Settings->RADIOS
Add SSH key to System->Network Device SSH Authentication->SSH Keys or use the configured password (iot/unifi-device-authentication).
ssh {{ hostname }}
From this doc, using SSH to adopt devices from the DMZ network.
After SSH into the device:
set-inform http://unifi-controller.grigri:8080/inform
And then you will see the device now show up for adoption.
