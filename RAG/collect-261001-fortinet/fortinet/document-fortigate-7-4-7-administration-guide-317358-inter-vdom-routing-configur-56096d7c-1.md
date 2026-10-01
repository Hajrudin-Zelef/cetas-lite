---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c-1
title: "document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c.md
source_anchor: ""
source_lines: [1, 122]
sha256: 52502c4951a7870fb2d32a8d86999782b3c22cb04bcb5e84c1da4e0e6da964d3
---

# document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c

Inter-VDOM routing configuration example: Partial-mesh VDOMs
This example shows how to configure a FortiGate unit to use inter-VDOM routing to route traffic between an internal network and FTP server that are each behind separate VDOMs. See Inter-VDOM routing for more information.
The following example shows how to configure per-VDOM settings, such as operation mode, routing, and firewall policies, in a network that includes the following VDOMs:
- 
                                                    VDOM-A: allows the internal network to access the Internet.
- 
                                                    VDOM-B: allows external connections to an FTP server.
- 
                                                    root: the management VDOM.
You can use VDOMs in either NAT or transparent mode on the same FortiGate. By default, VDOMs operate in NAT mode. In this example, both VDOM-A and VDOM-B use NAT mode. An inter-VDOM link is created and inter-VDOM routes configured to allow users on the internal network to access the FTP server.
This is an example of the partial-mesh VDOMs configuration since only VDOM-A is connected to VDOM-B but neither of those VDOMs are connected to the root VDOM. See Topologies for details.
This example assumes that the interfaces of the FortiGate have already been configured with the IP addresses depicted in the preceding diagram.
General steps for this example
This configuration requires the following general steps:
This example demonstrates how to configure these steps first using the GUI and then, at the end of the section, using the CLI. See Configuration with the CLI for details.
Enable Multi-VDOM mode and create the VDOMs
Multi-VDOM mode can be enabled in the GUI or CLI. Enabling it does not require a reboot, but does log you out of the device. The current configuration is assigned to the root VDOM.
|  | On FortiGate 90 series models and lower, VDOMs can only be enabled using the CLI. | 
To enable multi-VDOM mode in the GUI:
- 
                                                    On the FortiGate, go to System > Settings.
- 
                                                    In the System Operation Settings section, enable Virtual Domains.
- 
                                                    Click OK.
To create the VDOMs in the GUI:
- 
                                                    In the Global VDOM, go to System > VDOM. Click Create New.
- 
                                                    In the Virtual Domain field, enter VDOM-A.
- 
                                                    If required, set the NGFW Mode. If the NGFW Mode is Profile-based, Central SNAT can be enabled.
- 
                                                    Click OK to create the VDOM.
- 
                                                    Repeat the above steps for VDOM-B.
Assign interfaces to VDOMs
This example uses three interfaces on the FortiGate unit: port1 (internal network), port2 (FTP server), wan1 (WAN link for VDOM-A), and wan2 (WAN link for VDOM-B). The port1 and port2 interfaces are connected to the internal network and FTP server, respectively. The wan1 and wan2 interfaces are static assigned with IP addresses and default gateways provided by the ISPs for those WAN links.
To assign interfaces to VDOMs in the GUI:
- 
                                                    In the Global VDOM, go to Network > Interfaces.
- 
                                                    Select port1 and click Edit.
- 
                                                    From the Virtual domain list, select VDOM-A.
- 
                                                    Click OK.
- 
                                                    Repeat the preceding steps to assign port2 to VDOM-B.
- 
                                                    Repeat the preceding steps to assign wan1 to VDOM-A.
- 
                                                    Repeat the preceding steps to assign wan2 to VDOM-B.
Configure VDOM-A
VDOM-A allows connections from devices on the internal network to the Internet. WAN1 and port1 are assigned to this VDOM.
The per-VDOM configuration for VDOM-A includes the following:
- 
                                                    A firewall address for the internal network
- 
                                                    A static route to the ISP gateway
- 
                                                    A firewall policy allowing the internal network to access the Internet
All procedures in this section require you to connect to VDOM-A, either using a global or per-VDOM administrator account.
To add the firewall addresses in the GUI:
- 
                                                    Go to Policy & Objects > Addresses and select Address.
- 
                                                    Click Create new.
- 
                                                    Enter the following information: Name internal-network Type Subnet IP/Netmask 192.168.10.0/255.255.255.0 Interface port1
- 
                                                    Click OK.
To add a default route in the GUI:
- 
                                                    Go to Network > Static Routes and create a new route.
- 
                                                    Enter the following information: Destination Subnet IP address 0.0.0.0/0.0.0.0 Gateway 172.20.201.254 Interface wan1 Administrative Distance 10
- 
                                                    Click OK.
To add the firewall policy in the GUI:
- 
                                                    Go to Policy & Objects > Firewall Policy.
- 
                                                    Click Create New.
- 
                                                    Enter the following information: Name VDOM-A-Internet Incoming Interface port1 Outgoing Interface wan1 Source internal-network Destination all Schedule always Service ALL Action ACCEPT NAT enabled
- 
                                                    Click OK.
Configure VDOM-B
VDOM-B allows external connections to reach an internal FTP server. WAN2 and port2 are assigned to this VDOM.
The per-VDOM configuration for VDOM-B includes the following:
- 
                                                    A firewall address for the FTP server
- 
                                                    A virtual IP address for the FTP server
- 
                                                    A static route to the ISP gateway
- 
                                                    A firewall policy allowing external traffic to reach the FTP server
The procedures described above require you to connect to VDOM-B, either using a global or per-VDOM administrator account.
To add the firewall addresses in the GUI:
- 
                                                    Go to Policy & Objects > Addresses and select Address.
- 
                                                    Click Create new.
- 
                                                    Enter the following information: Name FTP-server Type Subnet IP/Netmask 192.168.20.10/255.255.255.255 Interface port2
- 
                                                    Click OK.
To add the virtual IP address in the GUI:
- 
                                                    Go to Policy & Objects > Virtual IPs and navigate to the Virtual IP tab.
- 
                                                    Click Create new.
- 
                                                    Enter the following information: Name FTP-server-VIP Interface wan2 External IP address/range 172.20.10.2 Map To 192.168.20.10
- 
                                                    Click OK.
To add a default route in the GUI:
- 
                                                    Go to Network > Static Routes and create a new route.
- 
