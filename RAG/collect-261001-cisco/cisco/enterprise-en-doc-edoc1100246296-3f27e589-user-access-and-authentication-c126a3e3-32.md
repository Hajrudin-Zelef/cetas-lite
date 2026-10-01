---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-32
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [5660, 5811]
sha256: aa563af19be764174684c174de93d4af07602314faa67e5a7391c1089cceafaa
---

# Configure DeviceA to generate a local key pair.

In addition, to meet high security requirements, 802.1X authentication is used for access control of user PCs through the RADIUS server, and the authentication point is deployed on M-LAG member devices (DeviceA and DeviceB).
In this example, interfaces 1, 2, 3, 4, 5, 6, and 7 on DeviceA represent 10GE 1/0/1, 10GE 1/0/2, 10GE 1/0/3, 10GE 1/0/4, 10GE 1/0/5, 10GE 1/0/6, and MEth 0/0/0, respectively.
Interfaces 1, 2, 3, 4, 5, 6, and 7 on DeviceB represent 10GE 1/0/1, 10GE 1/0/2, 10GE 1/0/3, 10GE 1/0/4, 10GE 1/0/5, 10GE 1/0/6, and MEth 0/0/0, respectively.
Interface 1 and interface 2 on DeviceC represent 10GE 1/0/1 and 10GE 1/0/2, respectively.
Interface 1 and interface 2 on DeviceD represent 10GE 1/0/1 and 10GE 1/0/2, respectively.
On DeviceA and DeviceB, Eth-Trunk 0 serves as a peer-link interface, M-LAG interface Eth-Trunk 1 connects to downlink terminals, and Eth-Trunk 2 connects to the uplink server.
Configure DeviceA and DeviceB as root bridges with the same bridge MAC address so that both devices function as root bridges on the Layer 2 network.
Configure an IP address for the management interface on both DeviceA and DeviceB to ensure Layer 3 connectivity for heartbeat packet exchange between them.
Configure M-LAG on DeviceA and DeviceB so that terminals can be dual-homed to DeviceA and DeviceB.
Configure RADIUS authentication and 802.1X authentication on DeviceA and DeviceB, enable NAC for M-LAG globally, and enable NAC authentication on M-LAG interfaces.
When configuring 802.1X authentication in an M-LAG scenario:
If the downstream device dual-homed to the M-LAG member devices is a switching device, root protection must be configured.
# Configure DeviceA.
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] stp root primary
[DeviceA] stp bridge-address 00e0-fc12-3458  //Configure the bridge MAC address of the root bridge (MAC address of the M-LAG master device).
[DeviceA] interface eth-trunk 1
[DeviceA-Eth-Trunk1] trunkport 10ge 1/0/2
[DeviceA-Eth-Trunk1] trunkport 10ge 1/0/5
[DeviceA-Eth-Trunk1] stp edged-port enable
[DeviceA-Eth-Trunk1] quit
# Configure DeviceB.
<HUAWEI> system-view
[HUAWEI] sysname DeviceB
[DeviceB] stp root primary
[DeviceB] stp bridge-address 00e0-fc12-3458   //Configure the bridge MAC address of the root bridge.
[DeviceB] interface eth-trunk 1
[DeviceB-Eth-Trunk1] trunkport 10ge 1/0/2
[DeviceB-Eth-Trunk1] trunkport 10ge 1/0/5
[DeviceB-Eth-Trunk1] stp edged-port enable
[DeviceB-Eth-Trunk1] quit
Ensure that DeviceA and DeviceB can communicate at Layer 3 through their management interfaces.
[DeviceA] interface meth 0/0/0
[DeviceA-MEth0/0/0] ip address 10.1.1.1 24
[DeviceA-MEth0/0/0] quit
[DeviceB] interface meth 0/0/0
[DeviceB-MEth0/0/0] ip address 10.1.1.2 24
[DeviceB-MEth0/0/0] quit
[DeviceA] dfs-group 1
[DeviceA-dfs-group-1] dual-active detection source ip 10.1.1.1 peer 10.1.1.2
[DeviceA-dfs-group-1] priority 150
[DeviceA-dfs-group-1] authentication-mode hmac-sha256 password YsHsjx_202206
[DeviceA-dfs-group-1] dfs-group state switchover disable   //Disable the proactive master/backup status switchback function of a DFS group. This step is required when both M-LAG and NAC are configured.
[DeviceA-dfs-group-1] quit
[DeviceB] dfs-group 1
[DeviceB-dfs-group-1] dual-active detection source ip 10.1.1.2 peer 10.1.1.1
[DeviceB-dfs-group-1] priority 120
[DeviceB-dfs-group-1] authentication-mode hmac-sha256 password YsHsjx_202206
[DeviceB-dfs-group-1] dfs-group state switchover disable   //Disable the proactive master/backup status switchback function of a DFS group. This step is required when both M-LAG and NAC are configured.
[DeviceB-dfs-group-1] quit
[DeviceA] interface eth-trunk 0
[DeviceA-Eth-Trunk0] mode lacp-static
[DeviceA-Eth-Trunk0] trunkport 10ge 1/0/3
[DeviceA-Eth-Trunk0] trunkport 10ge 1/0/4
[DeviceA-Eth-Trunk0] undo stp enable
[DeviceA-Eth-Trunk0] peer-link 1
[DeviceA-Eth-Trunk0] quit
[DeviceB] interface eth-trunk 0
[DeviceB-Eth-Trunk0] mode lacp-static
[DeviceB-Eth-Trunk0] trunkport 10ge 1/0/3
[DeviceB-Eth-Trunk0] trunkport 10ge 1/0/4
[DeviceB-Eth-Trunk0] undo stp enable
[DeviceB-Eth-Trunk0] peer-link 1
[DeviceB-Eth-Trunk0] quit
The uplink interfaces that connect the terminals to DeviceA and DeviceB must be added to an Eth-Trunk interface, and the working mode of the Eth-Trunk interface must be the same as that of the Eth-Trunk interfaces on both devices. In this example, the Eth-Trunk interfaces on both devices are configured to work in static LACP mode.
[DeviceA] vlan batch 10 11
[DeviceA] interface eth-trunk 1
[DeviceA-Eth-Trunk1] mode lacp-static
[DeviceA-Eth-Trunk1] port link-type access
[DeviceA-Eth-Trunk1] port default vlan 11
[DeviceA-Eth-Trunk1] dfs-group 1 m-lag 1
[DeviceA-Eth-Trunk1] quit
[DeviceB] vlan batch 10 11
[DeviceB] interface eth-trunk 1
[DeviceB-Eth-Trunk1] mode lacp-static
[DeviceB-Eth-Trunk1] port link-type access
[DeviceB-Eth-Trunk1] port default vlan 11
[DeviceB-Eth-Trunk1] dfs-group 1 m-lag 1
[DeviceB-Eth-Trunk1] quit
[DeviceA] interface eth-trunk 2
[DeviceA-Eth-Trunk2] mode lacp-static
[DeviceA-Eth-Trunk2] port link-type trunk
[DeviceA-Eth-Trunk2] port trunk allow-pass vlan 10 11
[DeviceA-Eth-Trunk2] trunkport 10ge 1/0/1
[DeviceA-Eth-Trunk2] trunkport 10ge 1/0/6
[DeviceA-Eth-Trunk2] quit
[DeviceB] interface eth-trunk 2
[DeviceB-Eth-Trunk2] mode lacp-static
[DeviceB-Eth-Trunk2] port link-type trunk
[DeviceB-Eth-Trunk2] port trunk allow-pass vlan 10 11
[DeviceB-Eth-Trunk2] trunkport 10ge 1/0/1
[DeviceB-Eth-Trunk2] trunkport 10ge 1/0/6
[DeviceB-Eth-Trunk2] quit
# Configure DeviceC.
<HUAWEI> system-view
[HUAWEI] sysname DeviceC
[DeviceC] vlan batch 10 11
[DeviceC] interface eth-trunk 2
[DeviceC-Eth-Trunk2] mode lacp-static
[DeviceC-Eth-Trunk2] port link-type trunk
[DeviceC-Eth-Trunk2] port trunk allow-pass vlan 10 11
[DeviceC-Eth-Trunk2] trunkport 10ge 1/0/1
[DeviceC-Eth-Trunk2] trunkport 10ge 1/0/2
[DeviceC-Eth-Trunk2] quit
# Configure DeviceD.
<HUAWEI> system-view
[HUAWEI] sysname DeviceD
[DeviceD] vlan batch 10 11
[DeviceD] interface eth-trunk 2
[DeviceD-Eth-Trunk2] mode lacp-static
[DeviceD-Eth-Trunk2] port link-type trunk
[DeviceD-Eth-Trunk2] port trunk allow-pass vlan 10 11
[DeviceD-Eth-Trunk2] trunkport 10ge 1/0/1
[DeviceD-Eth-Trunk2] trunkport 10ge 1/0/2
[DeviceD-Eth-Trunk2] quit
<HUAWEI> system-view
[HUAWEI] sysname DeviceE
[DeviceE] l2protocol-tunnel user-defined-protocol 802.1X protocol-mac 0180-c200-0003 group-mac 0100-0000-0002
[DeviceE] interface eth-trunk 1
[DeviceE-Eth-Trunk1] l2protocol-tunnel user-defined-protocol 802.1X enable
[DeviceE-Eth-Trunk1] port link-type access
[DeviceE-Eth-Trunk1] port default vlan 11
[DeviceE-Eth-Trunk1] quit
[DeviceE] interface 10ge 1/0/2
[DeviceE-10GE1/0/2] l2protocol-tunnel user-defined-protocol 802.1X enable
[DeviceE-10GE1/0/2] port link-type access
[DeviceE-10GE1/0/2] port default vlan 11
[DeviceE-10GE1/0/2] quit
# Configure the source address for communicating with the RADIUS server. When both M-LAG and NAC are configured, DeviceA and DeviceB must have the same source address configured.
[DeviceA] interface vlanif 10
[DeviceA-Vlanif10] ip address 192.168.1.1 24
[DeviceA-Vlanif10] mac-address 0000-0000-0011
[DeviceB] interface vlanif 10
[DeviceB-Vlanif10]ip address 192.168.1.1 24
[DeviceB-Vlanif10] mac-address 0000-0000-0011
Create and configure the RADIUS server template rd1.
[DeviceA] radius-server template rd1
[DeviceA-radius-rd1] radius-server authentication 192.168.10.1 1812 source ip-address 192.168.1.1
[DeviceA-radius-rd1] radius-server accounting 192.168.10.1 1813 source ip-address 192.168.1.1
[DeviceA-radius-rd1] radius-server shared-key cipher YsHsjx_202206mc@1
[DeviceA-radius-rd1] quit
[DeviceA] access-user m-lag enable    //Enable NAC in an M-LAG scenario.
[DeviceA] interface eth-trunk 1
[DeviceA-Eth-Trunk1] authentication-profile p1
[DeviceA-Eth-Trunk1] quit
#
sysname DeviceA
#
dfs-group 1
 priority 150
 dual-active detection source ip 10.1.1.1 peer 10.1.1.2
