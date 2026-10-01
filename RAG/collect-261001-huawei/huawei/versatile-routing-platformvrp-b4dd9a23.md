---
id: collect-261001-huawei/huawei/versatile-routing-platformvrp-b4dd9a23
title: "versatile-routing-platformvrp-b4dd9a23"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2016-04-11"]
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/versatile-routing-platformvrp-b4dd9a23.md
source_anchor: ""
source_lines: [1, 96]
sha256: 618d19c3645004bec84aa62535df777bf6fa0ce9c083c93f224de125622b82b5
---

# versatile-routing-platformvrp-b4dd9a23

Versatile Routing Platform(VRP)
It is a Command Line Interface which is used for configuration, management and monitoring of devices. It is based on a standardized and hierarchical command line system. Huawei devices use eNSP software for this purpose.
The command line interface
The command line interface has four command views: They are:
User View: The initial command view of VRP is the User View, which operates as an observation command view for observing parameter statuses and general statistical information.
System View: This view is used for configuring the huawei products. For application of changes to system parameters, users must enter the System View
Interface View:
Protocol View:
lThe example demonstrates a selection of common system defined shortcut keys that are widely used to simplify the navigation process within the VRP command line interface. Additional commands are as follows:
- CTRL+B moves the cursor back one character.
- CTRL+D deletes the character where the cursor is located.
- CTRL+E moves the cursor to the end of the current line.
- CTRL+F moves the cursor forward one character.
- CTRL+H deletes the character on the left side of the cursor.
- CTRL+N displays the next command in the historical command buffer.
- CTRL+P displays the previous command in the historical command buffer.
- CTRL+W deletes the word on the left side of the cursor.
- CTRL+X deletes all the characters on the left side of the cursor.
- CTRL+Y deletes all the characters on the right side of the cursor.
- ESC+B moves the cursor one word back.
- ESC+D deletes the word on the right side of the cursor.
- ESC+F moves the cursor forward one word
The system administrator sets user access levels that grant specific users access to specific command levels. The command level of a user is a value ranging from 0 to 3, whilst the user access level is a value ranging from 0 to 15. Level 0 defines a visit level for which access to commands that run network diagnostic tools, (such as ping and traceroute), as well as commands such as telnet client connections, and select display commands.
When you start the huawei router or any product on VRP. The following prompt appear:
<Huawei>
- To get help: ? and press enter to view all commands
<huawei> ?
<huawei> dis? to get help for display
2. to change time zone
<Huawei>clock timezone BJ add 08:00:00
3. to change date and time
<Huawei>clock datetime 10:20:29 2016-04-11
4. <Huawei>display clock
5. To change to system view
<huawei>system-view press enter. prompt changes to:
[Huawei]
6. to change component / system name
[huawei] sysname R1
[R1]
7. to configure interface
[Huawei]interface GigabitEthernet 0/0/0
8. to configure IP
[Huawei-GigabitEthernet0/0/0]ip address 10.0.12.1 255.255.255.0
9. to configure loopback
[Huawei-GigabitEthernet0/0/0]interface loopback 0
[Huawei-LoopBack0]ip address 1.1.1.1 32
10. The user level can be any value in the range of 0 – 15, where values represent a visit level (0), monitoring level (1), configuration level (2), and management level (3)
<Huawei>system-view
[Huawei]user-interface vty 0
[Huawei-ui-vty0]user privilege level 2
Summary:
- Which of the following operating system is used within Huawei product to configure, operate such managed devices:
  - VRP
  - Windows
  - Linux
  - MAC
- Ethernet network is a contentious network meaning that :
  - Host must compete for media access before sending the data
  - Hosts can send the data at any time
- In the case of the switched media such as 100BaseT:
  - Data transmission and reception becomes isolated within channels enabling potential for collision to occur eliminated
  - CSMA /CD is used to avoid collision
- The switch was evolved which is capable of:
  - Breaking down the shared collision domain into multiple collision domain removing threat of collision
  - Increasing collision domain and hence increasing chances of collision
- A broadcast domain is capable of being comprised of a single, or multiple collision domains, and any broadcast transmission is contained within the boundary of a broadcast domain.
- End of the broadcast boundary is connected to:
  - Gateway
  - Switch
  - Hub
- A single IP network can generally be understood to make up a :
Fill-up: broadcast domain, which refers to the scope of a link-layer segment.
- Routers are generally responsible for routing Internet datagrams (IP packets) to a given destination based on the knowledge of a forwarding address for the destination network, found within an internally managed forwarding table. TRUE /FALSE
- Full form of VRP is ……………….Versatile Routing Platform
- VRP uses a command line interface for managing and configuration of Huawei devices
- VRP is a kind of network operating system
- AR series enterprise routers (AR) include the AR150, AR200, AR1200, AR2200, and AR3200. Are all next generation of Huawei products.
- The AR series are positioned between the enterprise network and a public network, functioning as an ingress and egress gateway for data transmitted between the two networks
- The Sx7 Series Ethernet Switch provides data transport functionality
- Management of the ARG3 series routers and Sx7 series of switch can be achieved through establishing a connection to the console interface, and in the case of the AR2200, a connection is also possible to be established via a Mini USB interface
- Different views defined in the hierarchical command structure of VRP are:
  - User View
  - System View
  - Interface View
  - Protocol View
  - All of the above
- Command line views can be determined based on parentheses and information within the parentheses: (Read All of these followings)
  - User View :- View the running status and statistics of the device. <Huawei> user view
  - Systems View:- Set the system parameters of the device. E.g <Huawei>system-view [Huawei]
  - Interface View: Configure interface parameters. e.g. [Huawei]interface GigabitEthernet 0/0/0 [Huawei-GigabitEthernet0/0/0]
  - Protocol View:- Configure most routing protocol parameters.
- Whereas a switch separate the collision domain, the router is used to separate the broadcast domain [TRUE/FALSE]
- Switch is located in one broadcast domain thus broadcasting to all interfaces. Router does not forward the broadcast. Router has one interface and one broadcast domain
- Both the router and the switch are running the OS which is called VRP. VRP is a software running in hardware and providing functions such as routing and switching, network management, security, WLAN and UTM( universal s/w management)
- Version 5 running VRP version 5, and the NE router and the data center CE router which are modular router are running VRP-8. VRP-8 also provide virtualization function
- When configuring router and switch first time, we need to connect physically to a device usually a console
