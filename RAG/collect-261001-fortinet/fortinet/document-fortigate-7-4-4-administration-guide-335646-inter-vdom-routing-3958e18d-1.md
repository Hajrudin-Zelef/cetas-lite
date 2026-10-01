---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d-1
title: "document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d.md
source_anchor: ""
source_lines: [1, 84]
sha256: 1630426db856b72f44abcd1711a81be065b1784d117c2f9074bb62103e33f40a
---

# document-fortigate-7-4-4-administration-guide-335646-inter-vdom-routing-3958e18d

Inter-VDOM routing configuration example: Internet access
This example shows how to configure a FortiGate unit to use inter-VDOM routing to route outgoing traffic from individual VDOMs to a root VDOM with Internet access. See Inter-VDOM routing for more information.
Two departments of a company, Accounting and Sales, are connected to one FortiGate. The company uses a single ISP to connect to the Internet. This is an example of the Internet access configuration. See Topologies for details.
This example assumes that the interfaces of the FortiGate have already been configured with the IP addresses depicted in the preceding diagram.
General steps for this example
This example includes the following general steps. We recommend following the steps in the order below:
This example demonstrates how to configure these steps first using the GUI and then, at the end of the section, using the CLI. See Configuration with the CLI for details.
Enable multi-VDOM mode and create the VDOMs
Create the Accounting and Sales VDOMs.
To enable VDOMs in the GUI:
- 
                                                    Go to System > Settings.
- 
                                                    In the System Operation Settings section, enable Virtual Domains.
- 
                                                    Click OK.
|  | On FortiGate 90 series models and lower, VDOMs can only be enabled using the CLI. | 
To create the Sales and Accounting VDOMs in the GUI:
- 
                                                    In the Global VDOM, go to System > VDOM.
- 
                                                    Click Create New.
- 
                                                    In the Virtual Domain field, enter Sales.
- 
                                                    If required, set the NGFW Mode. If the NGFW Mode is Profile-based, Central SNAT can be enabled.
- 
                                                    Click OK to create the VDOM.
- 
                                                    Repeat the above steps for Accounting.
Assign interfaces to VDOMs
This example uses three interfaces on the FortiGate unit: port2 (AccountingLocal), port3 (SalesLocal), and port1 (WAN). Port2 and port3 interfaces each have a department's network connected. Port1 is for all traffic to and from the Internet and uses DHCP to configure its IP address, which is common with many ISPs.
To assign interfaces to VDOMs in the GUI:
- 
                                                    In the Global VDOM, go to Network > Interfaces.
- 
                                                    Select port2 and click Edit.
- 
                                                    From the Virtual domain list, select Accounting.
- 
                                                    Click OK.
- 
                                                    Repeat the preceding steps to assign port3 to the Sales VDOM.
- 
                                                    Repeat the preceding steps to assign port1 to the root VDOM.
Configure the VDOM links
To complete the connection between each VDOM and the management VDOM, add the two VDOM links. One pair is the Accounting - management link and the other is the Sales - management link. Each side of these links will be assigned IP addresses since they will be handy in configuring inter-VDOM routing in the next step.
To configure the Accounting and management VDOM link in the GUI:
- 
                                                    In the Global VDOM, go to Network > Interfaces.
- 
                                                    Select Create New > VDOM Link.
- 
                                                    Enter the following information: Name AccountVlnk Interface 0 Virtual Domain Accounting IP/Netmask 11.11.11.2/255.255.255.252 Administrative Access HTTPS, PING, SSH Comment Accounting side of the VDOM link Interface 1 Virtual Domain root IP/Netmask 11.11.11.1/255.255.255.252 Administrative Access HTTPS, PING, SSH Comment Management side of the VDOM link
- 
                                                    Click OK.
To configure the Sales and management VDOM link in the GUI:
- 
                                                    In the Global VDOM, go to Network > Interfaces.
- 
                                                    Select Create New > VDOM link.
- 
                                                    Enter the following information: Name SalesVlnk Interface 0 Virtual Domain Sales IP/Netmask 12.12.12.2/255.255.255.252 Administrative Access HTTPS, PING, SSH Comment Accounting side of the VDOM link Interface 1 Virtual Domain root IP/Netmask 12.12.12.1/255.255.255.252 Administrative Access HTTPS, PING, SSH Comment Management side of the VDOM link
- 
                                                    Click OK.
Configure inter-VDOM routing
A default static route can be configured on each VDOM to provide Internet access. In other words, this static route would provide inter-VDOM routing between each department VDOM and the root VDOM.
For this static route, these settings are used:
- 
                                                    Default Gateway: IP address of the management side of the VDOM link 
  - 
                                                            Accounting VDOM: 11.11.11.1
  - 
                                                            Sales VDOM: 12.12.12.1
- 
                                                            
- 
                                                    Interface: Interface on the department VDOM side of the VDOM link 
  - 
                                                            Accounting VDOM: AccountVlnk0
  - 
                                                            Sales VDOM: SalesVlnk0
- 
                                                            
