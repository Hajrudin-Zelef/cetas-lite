---
id: collect-261001-general-networking/general-networking/send-feedback-35-1
title: "Inter-VDOM routing configuration example: Internet access"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/send-feedback-35.md
source_anchor: ""
source_lines: [1, 133]
sha256: 2a3d8f6cf5f5f4de5b65c996122a57d0e308f58e513158b06594af8cd1876e5d
---

# Inter-VDOM routing configuration example: Internet access

This example shows how to configure a FortiGate unit to use inter-VDOM routing to route outgoing traffic from individual VDOMs to a root VDOM with Internet access. See Inter-VDOM routing for more information.

Two departments of a company, Accounting and Sales, are connected to one FortiGate. The company uses a single ISP to connect to the Internet. This is an example of the Internet access configuration. See Topologies for details.


This example assumes that the interfaces of the FortiGate have already been configured with the IP addresses depicted in the preceding diagram.

## General steps for this example

This example includes the following general steps. We recommend following the steps in the order below:

This example demonstrates how to configure these steps first using the GUI and then, at the end of the section, using the CLI. See Configuration with the CLI for details.

Create the Accounting and Sales VDOMs.

###### To enable VDOMs in the GUI:

1. 
                                                    Go to *System > Settings* .
2. 
                                                    In the *System Operation Settings* section, enable*Virtual Domains* .
3. 
                                                    Click *OK* .

|  | On FortiGate 90 series models and lower, VDOMs can only be enabled using the CLI. | 

###### To create the Sales and Accounting VDOMs in the GUI:

1. 
                                                    In the *Global* VDOM, go to*System > VDOM* .
2. 
                                                    Click *Create New* .
3. 
                                                    In the *Virtual Domain* field, enter*Sales* .
4. 
                                                    If required, set the *NGFW Mode* . If the*NGFW Mode* is*Profile-based* ,*Central SNAT* can be enabled.
5. 
                                                    Click *OK* to create the VDOM.
6. 
                                                    Repeat the above steps for *Accounting* .

This example uses three interfaces on the FortiGate unit: port2 (AccountingLocal), port3 (SalesLocal), and port1 (WAN). Port2 and port3 interfaces each have a department's network connected. Port1 is for all traffic to and from the Internet and uses DHCP to configure its IP address, which is common with many ISPs.

###### To assign interfaces to VDOMs in the GUI:

1. 
                                                    In the *Global* VDOM, go to*Network > Interfaces* .
2. 
                                                    Select *port2* and click*Edit* .
3. 
                                                    From the *Virtual domain* list, select*Accounting* .
4. 
                                                    Click *OK* .
5. 
                                                    Repeat the preceding steps to assign *port3* to the*Sales* VDOM.
6. 
                                                    Repeat the preceding steps to assign *port1* to the*root* VDOM.

To complete the connection between each VDOM and the management VDOM, add the two VDOM links. One pair is the Accounting - management link and the other is the Sales - management link. Each side of these links will be assigned IP addresses since they will be handy in configuring inter-VDOM routing in the next step.

###### To configure the Accounting and management VDOM link in the GUI:

1. 
                                                    In the *Global* VDOM, go to*Network > Interfaces* .
2. 
                                                    Select *Create New > VDOM Link* .
3. 
                                                    Enter the following information: **Name**AccountVlnk **Interface 0****Virtual Domain**Accounting **IP/Netmask**11.11.11.2/255.255.255.252 **Administrative Access**HTTPS, PING, SSH **Comment**Accounting side of the VDOM link **Interface 1****Virtual Domain**root **IP/Netmask**11.11.11.1/255.255.255.252 **Administrative Access**HTTPS, PING, SSH **Comment**Management side of the VDOM link
4. 
                                                    Click *OK* .

###### To configure the Sales and management VDOM link in the GUI:

1. 
                                                    In the *Global* VDOM, go to*Network > Interfaces* .
2. 
                                                    Select *Create New > VDOM link* .
3. 
                                                    Enter the following information: **Name**SalesVlnk **Interface 0****Virtual Domain**Sales **IP/Netmask**12.12.12.2/255.255.255.252 **Administrative Access**HTTPS, PING, SSH **Comment**Accounting side of the VDOM link **Interface 1****Virtual Domain**root **IP/Netmask**12.12.12.1/255.255.255.252 **Administrative Access**HTTPS, PING, SSH **Comment**Management side of the VDOM link
4. 
                                                    Click *OK* .

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
                                                            
- 
                                                    IP address: 0.0.0.0/0.0.0.0 (default)

###### To configure the default static route to the Internet in the Accounting VDOM:

1. 
                                                    In the *Accounting* VDOM, go to*Network > Static Routes* .
2. 
                                                    Click on *Create New* and select the version you need.
3. 
                                                    Enter the following information: **Destination**Subnet **IP address**0.0.0.0/0.0.0.0 **Gateway**11.11.11.1 **Interface**AccountVlink0 **Administrative Distance**10
4. 
                                                    Click *OK* .

###### To configure the default static route to the Internet in the Sales VDOM:

1. 
                                                    In the *Sales* VDOM, go to*Network > Static Routes* .
2. 
                                                    Click on *Create New* and select the version you need.
3. 
                                                    Enter the following information: **Destination**Subnet **IP address**0.0.0.0/0.0.0.0 **Gateway**12.12.12.1 **Interface**SalesVlink0 **Administrative Distance**10
4. 
                                                    Click *OK* .

With the VDOMs, physical interfaces, VDOM links, and static routes configured, the firewall must now be configured to allow the proper traffic. Firewalls are configured per-VDOM, and firewall objects and routes must be created for each VDOM separately.

###### To configure the firewall policies from AccountingLocal to Internet in the GUI:

