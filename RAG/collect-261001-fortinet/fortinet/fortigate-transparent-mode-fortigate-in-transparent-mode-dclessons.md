---
id: collect-261001-fortinet/fortinet/fortigate-transparent-mode-fortigate-in-transparent-mode-dclessons
title: "fortigate-transparent-mode-fortigate-in-transparent-mode-dclessons"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-transparent-mode-fortigate-in-transparent-mode-dclessons.md
source_anchor: ""
source_lines: [1, 24]
sha256: 2a7289872e30d984c2b28dd042b0a276db64ef6fe563d0a17598fd870439999b
---

# fortigate-transparent-mode-fortigate-in-transparent-mode-dclessons

## LAB Configuration Fortigate Transparent Mode

In the **fortigate transparent mode** all interface of the Fortigate are on same network and appliance does not do routing or NAT, It just act as L2 Firewall. The Fortigate unit acts as bridge between different network segments.

**Task :**

- Set the IP address of the Fortinet in your management LAN 10.10.11.0/24
- Create a policy to allow from Internal Interface to External Interface with Source All and Destination All
- Power off and Power on the Device.

#### **Solution**

**Step :1** Configure Management IP address :

Go to the Dashboard and enter the following command into the CLI console widget, substituting your own IP addresses where necessary:

You can now access the FortiGate using the new Management IP address (in the example, https:// 10.10.11.30).

Go to the Dashboard. The System Information widget shows the Operation Mode is Transparent.


## LEAVE A COMMENT

Please login here to comment.
