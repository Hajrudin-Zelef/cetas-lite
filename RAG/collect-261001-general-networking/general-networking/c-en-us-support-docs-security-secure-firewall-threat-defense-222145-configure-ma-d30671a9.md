---
id: collect-261001-general-networking/general-networking/c-en-us-support-docs-security-secure-firewall-threat-defense-222145-configure-ma-d30671a9
title: "c-en-us-support-docs-security-secure-firewall-threat-defense-222145-configure-ma-d30671a9"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/c-en-us-support-docs-security-secure-firewall-threat-defense-222145-configure-ma-d30671a9.md
source_anchor: ""
source_lines: [1, 90]
sha256: 57616f6e4c0c4a66b6c86a3d61d8636a4dc92cbcbabd16858cf87f5e95e38a8e
---

# c-en-us-support-docs-security-secure-firewall-threat-defense-222145-configure-ma-d30671a9

Introduction
This document describes the process for modifying the Manager Access on the Firepower Threat Defense (FTD) from a Management to a Data interface.
Prerequisites
Requirements
Cisco recommends that you have knowledge of these topics:
- Firepower Threat Defense (FTD)
- Firepower Management Center (FMC)
Components Used
- Firepower Management Center Virtual 7.4.1
- Firepower Threat Defense Virtual 7.2.5
The information in this document was created from the devices in a specific lab environment. All of the devices used in this document started with a cleared (default) configuration. If your network is live, ensure that you understand the potential impact of any command.
Background Information
Each device includes a single dedicated Management interface for communicating with the FMC. You can optionally configure the device to use a data interface for management instead of the dedicated Management interface, The FMC access on a data interface is useful if you want to manage the Firepower Threat Defense remotely from the outside interface, or you do not have a separate management network. This change has to be performed on the Firepower Management Center for FTD managed by FMC.
The FMC access from a data interface has a few limitations:
- You can only enable manager access on one physical, data interface. You cannot use a subinterface or EtherChannel. 
- Routed firewall mode only, using a routed interface.
- PPPoE is not supported. If your ISP requires PPPoE, you have to put a router with PPPoE support between the Firepower Threat Defense and the WAN modem.
- You cannot use separate management and event only interfaces.
Configure
Proceed with Interface Migration
Note: It is strongly recommended to have the latest backup of both FTD and FMC before proceeding with any changes.
 
 
- 
Navigate to Devices > Device Management page, click Edit for the device you are making changes.
2. Go to the  section, and click the link for Manager Access Interface.
The Manager Access Interface field displays the existing Management interface. Click the link to select the new interface type, which is the Data Interface option in the Manage device by drop down list and click Save.
3. You must now proceed to Enable management access on a data interface, navigate to Devices > Device Management > Interfaces > Edit Physical Interface > Manager Access.
Note: (Optional) If you use a secondary interface for redundancy, enable management access on the interface used for redundancy purpose.
(Optional) If you use DHCP for the interface, enable the web type DDNS method on the Devices > Device Management > DHCP > DDNS dialog.
(Optional) Configure DNS in a Platform Settings policy, and apply it to this device at Devices > Platform Settings > DNS.
 
 4. Make sure the threat defense can route to the management center through the data interface; add a static route if necessary on .
- Click IPv4or IPv6 depending on the type of static route that you are adding.
- Choose the Interface to which this static route applies.
- In the Available Network list, choose the destination network.
- In the Gateway or IPv6 Gateway field, enter or choose the gateway router which is the next hop for this route.
  (Optional) To monitor route availability, enter or choose the name of an Service Level Agreement (SLA) Monitor object that defines the monitoring policy, in the Route Tracking field.
5. Deploy configuration changes. The configuration changes are now deployed over the current Management interface.
6. At the FTD CLI, set the Management interface to use a static IP address and the gateway to be data-interfaces.
- configure network {ipv4 | ipv6} manual ip_address netmask data-interfaces
Note: Although you do not plan to use the Management interface, you must set a static IP address. For example, a private address so that you can set the gateway to data-interfaces. This management is used to forward the management traffic to data interface using tap_nlp interface.
 
 7. Disable the Management in the Management Center. Click Edit and update the Remote Host Address IP addressand (Optional)Secondary Address for the threat defense in the  section, and enable the connection.
Enable SSH on Platform Settings
Enable SSH for the data interface in Platform Settings policy, and apply it to this device at 
- The hosts or networks you are allowing to make SSH connections.
- Add the zones that contain the interfaces to allow SSH connections. For interfaces not in a zone, you can type the interface name into the field Selected Zones/Interfaces list and click Add.
- Click OK. Deploy the changes.
Note: SSH is not enabled by default on the data interfaces, so if you want to manage the threat defense using SSH, you need to explicitly allow it.
 
 Verify
Ensure that the management connection is established over the Data interface.
Verify from FMC Graphical User Interface (GUI)
In the management center, check the management connection status on the  page.
Verify from FTD Command Line Interface (CLI)
At thethreat defenseCLI, enter thesftunnel-status-briefcommand to view the management connection status.
The status shows a successful connection for a data interface, showing the internal tap_nlp interface.
Troubleshoot
In the management center, check the management connection status on the  page.
At thethreat defenseCLI, enter thesftunnel-status-briefcommand to view the management connection status. You can also usesftunnel-statusto view more complete information.
Management Connection Status
Working Scenario
Non-Working Scenario
Validate the Network Information
At thethreat defenseCLI, view the Management and manager access data interface network settings:
> show network
Validate the Manager State
At the threat defense CLI, check that the management center registration was completed.
  > show managers
Note: This command does not show the current status of the management connection.
 
 Validate Network Connectivity
Ping the Management Center
At thethreat defenseCLI, use the command to ping themanagement centerfrom the data interfaces:
  >  ping fmc_ip
At thethreat defenseCLI, use the command to ping themanagement centerfrom the Management interface, which routes over the backplane to the data interfaces:
  >  ping system fmc_ip
Check Interface Status, Statistics, and Packet Count
At thethreat defenseCLI, see information about the internal backplane interface, nlp_int_tap:
  >  show interface detail
Validate Route on FTD to Reach FMC
At thethreat defenseCLI, check that the default route (S*) was added and that internal NAT rules exist for the Management interface (nlp_int_tap).
  > show route
  >  show nat
Check Sftunnel and Connection Statistics
  >  show running-config sftunnel
Warning: Throughout the process of changing manager access, refrain from deleting the manager on the FTD or unregistering/force deleting the FTD from FMC.
 
 Related Information
