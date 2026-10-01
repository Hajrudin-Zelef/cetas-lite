---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-508653-configuring-sd-wan-in-an-ha-5ebe936e-1
title: "document-fortigate-7-4-0-administration-guide-508653-configuring-sd-wan-in-an-ha-5ebe936e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-508653-configuring-sd-wan-in-an-ha-5ebe936e.md
source_anchor: ""
source_lines: [1, 67]
sha256: d6c22406a7f7ad05a530312ffb2c06937cc441931726082cf539d88e00f18fec
---

# document-fortigate-7-4-0-administration-guide-508653-configuring-sd-wan-in-an-ha-5ebe936e

Configuring SD-WAN in an HA cluster using virtual VLAN switch
In this SD-WAN configuration, two FortiGates in an active-passive (A-P) HA pair are used to provide hardware redundancy. Instead of using external switches to provide a mesh network connection to the ISP routers, each FortiGate connects to an ISP and uses a virtual VLAN switch to connect to each other. Virtual VLAN switch mode allows 802.1Q VLANs to be assigned to it, and the configuration of one interface as a trunk port.
|  | Only FortiGate models that support the virtual VLAN switch feature can be used for this solution. See Virtual VLAN switch for details. | 
In this topology:
- 
                                                    An HA cluster of two FortiGate 40Fs is used as an example
- 
                                                    The lan interface is broken up into lan1, lan2 and lan3
- 
                                                    lan1 is connected between FGT_A and FGT_B and designated as a virtual VLAN switch trunk port
- 
                                                    lan2 is used for HA heartbeat, connecting the two FortiGates in HA
- 
                                                    lan3 is connected to the internal network
- 
                                                    wan is connected to ISP1 on FGT_A
- 
                                                    port A is connected to ISP on FGT_B
Two VLANs are created (VLAN 10 and VLAN 20), and assigned to two VLAN switches. Each VLAN switch has the WAN interfaces (wan and port A) as its member.
When FGT_A is the primary device, it reaches ISP1 directly via the local wan interface and reaches ISP2 via the trunk port towards FGT_B. Packets are tagged/untagged as VLAN 10 as it passes through the trunk before egressing on port A on FGT_B.
When FGT_B is the primary device, it reaches ISP2 directly via the local port A interface and reaches ISP1 via the trunk port towards FGT_A. Packets are tagged/untagged as VLAN 20 as it passes through the trunk before egressing on the wan interface on FGT_A.
|  | Using virtual VLAN switches for an SD-WAN HA configuration, as described in this example, requires fewer interfaces than the hardware switch configuration described in Configuring SD-WAN in an HA cluster using internal hardware switches A virtual VLAN switch configuration with 2 WAN connections requires 5 interfaces. With 3 WAN connections, a total of 6 interfaces are required. For the internal hardware switch solution, each WAN connection requires a corresponding hardware switch interface. With 2 WAN connections, a total of 6 interfaces are needed. With 3 WAN connections, a total of 8 interfaces are needed. | 
HA failover
This is not a standard HA configuration with external switches. In the case of a device failure, one of the ISPs will no longer be available because the switch that is connected to it will be down.
For example, if FGT_A loses power, HA failover will occur and FGT_B will become the primary unit. Its connection to VLAN 10 over the trunk port will be down, so it will be unable to connect to ISP 1. Its SD-WAN SLAs will be broken, and traffic will only be routed through ISP 2.
|  | SD-WAN SLA health checks should be used to monitor the health of each ISP. HA link monitoring will not be possible, since at least one of the WAN interfaces will be down on each FortiGate at all times. | 
Configuration
In the configuration example, it is assumed HA A-P mode is already configured and FGT_A and FGT_B are in sync.
To configure the virtual VLAN switch, VLANs and interface settings in the GUI:
- 
                                                    Enable Virtual VLAN switch from the GUI under the System > Settings page. In the View Settings section, enable VLAN switch mode. Click Apply.
- 
                                                    Create a new VLAN switch and assign the wan interface in the GUI. 
  - 
                                                            Go to Network > Interfaces and click Create New > Interface.
  - 
                                                            Enter the name ISP1
  - 
                                                            Set the Type to VLAN Switch.
  - 
                                                            Enter the VLAN ID 10.
  - 
                                                            Click the + and add the Interface Members. Select the interface wan.
  - 
                                                            Configure the Address (192.168.10.99/24) and Administrative Access settings as needed.
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Repeat the same steps for a new VLAN switch and assign the port A interface in the GUI. 
  - 
                                                            Go to Network > Interfaces and click Create New > Interface.
  - 
                                                            Enter the name ISP2
  - 
                                                            Set the Type to VLAN Switch.
  - 
                                                            Enter the VLAN ID 20.
  - 
                                                            Click the + and add the Interface Members. Select the interface 'a'.
  - 
                                                            Configure the Address (192.168.20.99/24) and Administrative Access settings as needed.
  - 
                                                            Click OK
- 
                                                            
