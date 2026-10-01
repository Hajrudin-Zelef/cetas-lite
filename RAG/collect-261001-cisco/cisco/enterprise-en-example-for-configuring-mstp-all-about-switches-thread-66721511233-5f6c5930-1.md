---
id: collect-261001-cisco/cisco/enterprise-en-example-for-configuring-mstp-all-about-switches-thread-66721511233-5f6c5930-1
title: "enterprise-en-example-for-configuring-mstp-all-about-switches-thread-66721511233-5f6c5930"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost", "ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-example-for-configuring-mstp-all-about-switches-thread-66721511233-5f6c5930.md
source_anchor: ""
source_lines: [1, 50]
sha256: 200f41685aa8938add8e3719fcc668aa6309bccaf9651db61b8ec9479d72fe7e
---

# enterprise-en-example-for-configuring-mstp-all-about-switches-thread-66721511233-5f6c5930

Hello, everybody!
This post contains an example for configuring MSTP, as part of the All About Switches section on our Community. Please see below for more details.
The Multiple Spanning Tree Protocol (MSTP) defined in IEEE 802.1s combines the Spanning Tree Protocol (STP) and Virtual Local Area Network (VLAN). Simply speaking, STP or RSTP is implemented based on ports, whereas MSTP based on instances. MSTP inherits fast transition of RSTP and solves the problem of RSTP in which where all VLANs share one spanning tree.
Compared with STP/RSTP, MSTP uses instances and regions.
? One Multiple Spanning Tree Instance (MSTI) is a collection of multiple VLANs. Multiple VLANs where the same service is transmitted can be mapped to one MSTI. The topology of each MSTI is calculated independent of one another. Traffic can be load-balanced through the mapping between VLANs and MSTIs. By default, all VLANs map to MSTI 0.
? A region consists of the region name, revision level, and mapping between VLANs and MSTIs. All switches in a region must have the same MST region configuration. That is, the switches that have the same region name, revision level, and mapping between VLANs and MSTIs belong to the same region. By default, the region name is the bridge MAC address of a switch, the revision level is 0, and all VLANs are mapped to MSTI 0.
MSTP defines the following spanning trees:
1. The Common Internal Spanning Tree (CIST) is composed of the Common Spanning Tree (CST) and Internal Spanning Tree (IST).
2. The CST connects all the MST regions on a switching network.
3. The IST is located in an MST region.
4. A Single Spanning Tree (SST) is formed when a switching device running STP or RSTP belongs to only one spanning tree or an MST region has only one switching device.
? For all connected switches running STP/RSTP/MSTP, there is only one CIST root, that is, root of MSTI 0. In the preceding figure, the root of A0 is the CIST root.
? Each MSTI in each region has a regional root. That is, the number of regional roots in each region depends on the number of MSTIs.
? Master port: Among all edge ports in an MST region, the master port has the smallest cost to the CIST root. The master port is located on the shortest path connecting MST regions to the CIST root. The master port is the root port in a CIST.
? Regional edge port: A regional edge port is located at the edge of an MST region and connects to another MST region. Multiple regional edge ports may exist in an MST region.
? CIST calculation
The following vectors participate in the CIST calculation to trim a ring network into a tree network:
{root ID, external root path cost, region root ID, internal root path cost, designated switching device ID, designated port ID, receiving port ID}
CIST calculation procedure:
1. Devices on a network send and receive BPDUs. After comparing the configuration BPDUs, the device with the highest priority on the entire network is selected as the CIST root.
2. MSTP calculates an IST for each MST region.
3. MSTP considers each MST region as a single switch and computes a CST to interconnect MST regions.
The CST and ISTs constitute a CIST for the entire network.
? MSTI calculation
The following vectors participate in the MSTI calculation to trim a ring network into a tree network:
{regional root ID, internal root path cost, designated switching device ID, designated port ID, receiving port ID}
In an MST region, MSTP calculates an MSTI for each VLAN based on the mapping between VLANs and MSTIs. Each MSTI is calculated independently. The calculation process is similar to the process for STP/RSTP to calculate a spanning tree.
? The VLANs bound to MSTIs must have been created, and ports have been added to VLANs. The display stp brief command displays only information about all STP-enabled ports in Up state.
? After configuring parameters of an MST region, you must run the active region-configuration command to activate the MST region configuration.
? The default region name is the MAC address of a device. The MAC address of each device is different, so you must manually specify a region to keep region information consistency on switches.
? On fixed switches, ports send STP BPDUs to the CPU after the bpdu enable command is configured globally or on ports. Without this configuration, the ports cannot send BPDUs to the CPU after receiving the BPDUs. STP then cannot converge.
For S2700 switches in V100R005 and earlier versions, check whether the bpdu enable command has been configured globally. For other fixed switches in V100R005 and earlier versions, check whether the bpdu enable command has been configured on physical ports and the Eth-Trunk.
For fixed switches in V100R006 and later versions, the bpdu enable command is configured globally or on ports by default on corresponding models.
? After STP is enabled on an Eth-Trunk, change the cost of the Eth-Trunk to a value smaller than the cost of physical ports so that the Eth-Trunk will be not negotiated as the backup port. The cost of the Eth-Trunk is the cost of a member port divided by the number of member ports. Therefore, the Eth-Trunk cost changes after the status of member ports changes. A smaller cost indicates better link quality.
As shown in the figure, SwitchA, SwitchB, SwitchC, and SwitchD constitute a ring network to back up links and load balance traffic of two departments. MSTP is required to block a port and prune the network into a loop-free tree network.
? Traffic in department 1 and department 2 needs to be load balanced.
? SwitchA and SwitchB function as the root bridge and secondary root bridge of the two departments.
? On SwitchC and SwitchD, GE0/0/3 ports connected to users cannot participate in STP calculation.
The configuration roadmap is as follows:
1. Create VLANs and add ports to VLANs.
2. Configure switches to work in MSTP mode.
3. Configure the region named RG1 and specify the mapping between VLANs and MSTIs.
4. Configure SwitchA and SwitchB as the root bridge and secondary root bridge of the two departments.
5. Configure switches to use the default path cost. Huawei switches use IEEE 802.1T standard to calculate path costs by default. By default, the cost of a GE port is 20000, and the cost of an Ethernet port is 200000. In eNSP, the default STP path cost is 1.
6. Disable STP on GE0/0/3 ports on SwitchC and SwitchD.
7. Enable STP on SwitchA, SwitchB, SwitchC, and SwitchD.
Step 1: Create VLANs and add ports to VLANs.
| <HUAWEI> system-view [HUAWEI] sysname SwitchA [SwitchA] vlan batch 2 to 4094 [SwitchA] interface gigabitethernet0/0/1 [SwitchA-GigabitEthernet0/0/1] port link-type trunk [SwitchA-GigabitEthernet0/0/1] port trunk allow-pass vlan 2 to 4094 [SwitchA-GigabitEthernet0/0/1] quit [SwitchA] interface gigabitethernet0/0/2 [SwitchA-GigabitEthernet0/0/2] port link-type trunk [SwitchA-GigabitEthernet0/0/2] port trunk allow-pass vlan 2 to 4094 [SwitchA-GigabitEthernet0/0/2] quit | 
| <HUAWEI> system-view [HUAWEI] sysname SwitchB [SwitchB] vlan batch 2 to 4094 [SwitchB] interface gigabitethernet0/0/1 [SwitchB-GigabitEthernet0/0/1] port link-type trunk [SwitchB-GigabitEthernet0/0/1] port trunk allow-pass vlan 2 to 4094 [SwitchB-GigabitEthernet0/0/1] quit [SwitchB] interface gigabitethernet0/0/2 [SwitchB-GigabitEthernet0/0/2] port link-type trunk [SwitchB-GigabitEthernet0/0/2] port trunk allow-pass vlan 2 to 4094 [SwitchB-GigabitEthernet0/0/2] quit | 
| <HUAWEI> system-view [HUAWEI] sysname SwitchC [SwitchC] vlan batch 2 to 4094 [SwitchC] interface gigabitethernet0/0/1 [SwitchC-GigabitEthernet0/0/1] port link-type trunk [SwitchC-GigabitEthernet0/0/1] port trunk allow-pass vlan 2 to 4094 [SwitchC-GigabitEthernet0/0/1] quit [SwitchC] interface gigabitethernet0/0/2 [SwitchC-GigabitEthernet0/0/2] port link-type trunk [SwitchC-GigabitEthernet0/0/2] port trunk allow-pass vlan 2 to 4094 [SwitchC-GigabitEthernet0/0/2] quit | 
