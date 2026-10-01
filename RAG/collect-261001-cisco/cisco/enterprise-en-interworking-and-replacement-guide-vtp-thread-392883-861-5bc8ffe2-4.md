---
id: collect-261001-cisco/cisco/enterprise-en-interworking-and-replacement-guide-vtp-thread-392883-861-5bc8ffe2-4
title: "Run the show running-config command to check the interface configuration."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-interworking-and-replacement-guide-vtp-thread-392883-861-5bc8ffe2.md
source_anchor: ""
source_lines: [249, 327]
sha256: a7517b77dbb966747e6dddb5fe2dddfe797537f260bec2c55a8abf35bd1b34db
---

# Run the show running-config command to check the interface configuration.

l This example applies to Huawei S series switches of V200R005 and later versions.
In Figure 1-8, a Huawei S series switch and a Cisco switch are deployed on a network. To reduce the configuration and maintenance workload, the Huawei S series switch uses VCMP and the Cisco switch uses VTP to synchronize VLAN information to other switches. The Cisco switch and user hosts connected to the Huawei S series switch need to communicate in VLAN 10.
Figure 1-8 Hybrid networking of the C-H-H-C model
3. Configure VCMP and LNP on the Huawei S series switch.
Step 3 Configure the VCMP server on the Huawei S series switch.
# Configure Layer 2 transparent transmission on the VCMP server.
<HUAWEI> system-view 
<HUAWEI> sysname Server 
[Server] l2protocol-tunnel vtp group-mac 0100-5e00-0011   
[Server] interface GigabitEthernet1/0/1 
[Server-GigabitEthernet1/0/1] l2protocol-tunnel vtp vlan 1   
[Server-GigabitEthernet1/0/1] quit 
[Server] interface GigabitEthernet1/0/2 
[Server-GigabitEthernet1/0/2] l2protocol-tunnel vtp vlan 1   
[Server-GigabitEthernet1/0/2] quit 
# Configure VCMP on the VCMP server.
[Server] vcmp domain huawei   
[Server] vcmp role server   
[Server] vcmp authentication sha2-256 password huawei   
[Server] vlan 10   
[Server-vlan10] quit
# Add interfaces on the VCMP server to the VLAN.
[Server] interface GigabitEthernet1/0/1 
[Server-GigabitEthernet1/0/48] port link-type trunk   
[Server-GigabitEthernet1/0/48] port trunk allow-pass vlan 2 to 4094   
[Server-GigabitEthernet1/0/48] quit 
[Server] interface GigabitEthernet1/0/2 
[Server-GigabitEthernet1/0/2] port link-type trunk   
[Server-GigabitEthernet1/0/2] port trunk allow-pass vlan 2 to 4094   
[Server-GigabitEthernet1/0/2] quit 
[Server] interface GigabitEthernet1/0/3 
[Server-GigabitEthernet1/0/3] port link-type negotiation-desirable   
[Server-GigabitEthernet1/0/3] port default vlan 10   
[Server-GigabitEthernet1/0/3] quit 
Step 4 Configure VCMP client 1 on the Huawei S series switch.
# Configure Layer 2 transparent transmission on VCMP client 1.
<HUAWEI> system-view 
<HUAWEI> sysname Client1 
[Client1] l2protocol-tunnel vtp group-mac 0100-5e00-0011   
[Client1] interface GigabitEthernet1/0/48 
[Client1-GigabitEthernet1/0/48] l2protocol-tunnel vtp vlan 1   
[Client1-GigabitEthernet1/0/48] quit 
[Client1] interface GigabitEthernet1/0/46 
[Client1-GigabitEthernet1/0/46] l2protocol-tunnel vtp vlan 1   
[Client1-GigabitEthernet1/0/46] quit 
# Configure VCMP on VCMP client 1.
[Client1] vcmp domain huawei   
[Client1] vcmp role client   
[Client1] vcmp authentication sha2-256 password huawei   
# Add interfaces on VCMP client 1 to the VLAN.
[Client1] interface GigabitEthernet1/0/48 
[Client1-GigabitEthernet1/0/48] port link-type trunk   
[Client1-GigabitEthernet1/0/48] port trunk allow-pass vlan 2 to 4094   
[Client1-GigabitEthernet1/0/48] quit 
[Client1] interface GigabitEthernet1/0/46 
[Client1-GigabitEthernet1/0/46] port link-type trunk   
[Client1-GigabitEthernet1/0/46] port trunk allow-pass vlan 2 to 4094   
[Client1-GigabitEthernet1/0/46] quit 
[Client1] interface GigabitEthernet1/0/1 
[Client1-GigabitEthernet1/0/1] port link-type access   
[Client1-GigabitEthernet1/0/1] port default vlan 10   
[Client1-GigabitEthernet1/0/1] quit 
Step 5 Configure VCMP client 2 on the Huawei S series switch.
# Configure VCMP on VCMP client 2.
<HUAWEI> system-view 
<HUAWEI> sysname Client2 
[Client2] vcmp domain huawei   
[Client2] vcmp role client    
[Client2] vcmp authentication sha2-256 password huawei   
# Add interfaces on VCMP client 2 to the VLAN.
[Client2] interface GigabitEthernet1/0/1 
[Client2-GigabitEthernet1/0/1] port link-type access   
[Client2-GigabitEthernet1/0/1] port default vlan 10   
[Client2-GigabitEthernet1/0/1] quit 
[Client2] interface GigabitEthernet1/0/2 
[Client2-GigabitEthernet1/0/2] port default vlan 10   
[Client2-GigabitEthernet1/0/2] quit 
Step 6 Verify the configuration.
l Run the display vcmp status command to check the VCMP configuration on the Huawei S series switch.
