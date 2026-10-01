---
id: collect-261001-fortinet/fortinet/fortinet-document-library
title: "Transparent mode"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortinet-document-library.md
source_anchor: ""
source_lines: [1, 15]
sha256: 60902acd9dbf087dc38ee0c72ac3788558f86515693f4437efbd711306abf228
---

# Transparent mode

# Transparent mode

To configure the FortiGate-VM to operate in transparent mode, you must configure the VMware ESXi server's virtual switches to operate in promiscuous mode to allow traffic that is not addressed to the FortiGate-VM to pass through it.

###### To configure virtual switches to support FortiGate-VM transparent mode:

1. In the vSphere client, select your VMware server, then select the *Configuration* tab.
2. In *Hardware* , select*Networking* .
3. Select *Properties* of vSwitch0.
4. In the *Properties* window, select*vSwitch* , then select*Edit* .
5. Select the *Security* tab, set*Promiscuous Mode* to*Accept* , then select*OK* .
6. Select *Close* .
7. Repeat steps 3 to 6 for other virtual switches that the FortiGate-VM uses.
