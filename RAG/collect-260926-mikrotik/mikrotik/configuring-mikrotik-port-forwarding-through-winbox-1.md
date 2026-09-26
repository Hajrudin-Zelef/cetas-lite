---
id: collect-260926-mikrotik/mikrotik/configuring-mikrotik-port-forwarding-through-winbox-1
title: "Configuring MikroTik port forwarding through Winbox"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["training"]
source: docs/RAG/lot-mikrotik/tools/configuring-mikrotik-port-forwarding-through-winbox.md
source_anchor: ""
source_lines: [1, 85]
sha256: 5074b0850a2c91caaeb2263285957551de9680fd7078366310611e29a9a93268
---

# Configuring MikroTik port forwarding through Winbox

MikroTik RouterOS is a powerful Linux-independent operating system designed to interact with MikroTik network equipment. MikroTik RouterOS is more than a simple router operating system and has many functions and applications. This software is capable of turning it into dedicated routers by running on a normal and personal computer.


Port Forwarding or Port Mapping is a NAT program used in MikroTik to redirect traffic from an IP address and port number on a remote network to an IP address and port number on a local network. In other words, Port Forwarding is the most secure way to connect to your servers from outside your local network without risking network security.

Using MikroTik port forwarding, you can easily connect from outside your private/local zone (from the internet/public) to a server in your private/local zone, such as web server, game server, FTP server, and NVR and DVR.

If you have servers and applications on a private network that you need to access over the Internet and public networks, the best way to log into MikroTik RouterOS is through Winbox. Configuring Mikrotik port forwarding is very simple. In this article, we will introduce you to Port Forwarding in Mikrotik and teach you how to configure port forwarding in the Mikrotik router step.

Before starting the port forwarding configuration process in Mikrotik, we will briefly introduce the Mikrotik operating system and a definition of Port forwarding. Stay with us until the end of the article.

MikroTik is a hardware and network equipment manufacturer in Latvia that is the main provider of Internet access infrastructure (hardware and software) in most countries. They design and market hardware and software for computer networks, including routers, switches, access points, utility software, and operating systems.

RouterOS is the operating system used by MikroTik. MikroTik provides a lot of freedom to manage networks with its advanced router. Installing RouterOS on a PC turns the device into a fully functional router with features such as routing, firewall, bandwidth control, wireless access point, backhaul link, hotspot gateway, VPN server, and more. In order to provide automatic operation, the boot time can be reduced by using Mikrotik, a very effective router.

Port forwarding refers to the technique of routing data traffic on a network that is usually directed to a specific IP address and port on a particular computer and then to a new destination. A MikroTik router or any other device running RouterOS makes this process easier. Therefore, assigning a specific port to a specific service in the private network makes it possible to use that specific service for another user by entering that port in their browser or software.

Mikrotik is placed as a router between the internal path of your private network and the public network (Internet). For example, suppose you have a site inside the organization with a specific port. In that case, if you are thinking of outsourcing your IP address or domain with a port to the software inside the organization, you should use port forwarding.

Imagine that you are the IT manager of an extensive network looking for a safe, low-risk method for network security to allow someone to remotely access your VPS on your network while you don’t want to share the server’s IP with that remote person. In this situation, port forwarding in the Mikrotik router is the safest solution to solve your problem, which provides the possibility of connecting to the VPS in the local network for people outside your private network through the port. For this, you need to buy the desired Mikrotik VPS first so that you can benefit from Mikrotik Port forwarding.

The ether1 interface of the MikroTik router in this network is connected to a wide area network (WAN) with IP address120.50.–.198, while the ether2 interface is connected to a LAN switch with IP address193.168.20. Only those on the local area network (LAN) can access the three servers (web server, FTP server, and SSH server) located on the internal network. By using MikroTik Port Forwarding, you can make these servers available to users outside your local network; we will teach how to configure MikroTik Port Forwarding to connect to servers inside the local network through the Internet for users outside the local network.

Mikrotik Port forwarding is a widely used method to respond to various purposes, which are considered three common purposes in our training:

- Port Forwarding to Internal Web Server
- Port Forwarding to Internal FTP Server
- Port Forwarding to Internal SSH Server

### Configuring MikroTik port forwarding to the internal Web server through Winbox

To go through the port forwarding steps, note that you have installed the latest version of MikroTik RouterOS.

To provide access to a web server inside the internal network from outside the internal network, we provided the following steps to configure MikroTik Port Forwarding to connect from the public network to a specific web server with Ip Address (193.168.20.10) according to the network diagram as an example.

**Step 1**: Log in to the Mikrotik server using **Winbox** as a server administrator with the required permissions

**Step 2**: Click on **IP** on the left side of the panel

**Step 3**: In the opened menu, select **Firewall**.

**Step 4**: In the Firewall window, click on the **NAT** tab.

**Step 5**: The **NAT Rule window** will appear by clicking on the **PLUS (+)** sign.

**Step 6**: By clicking on the **General** tab, the Chain drop-down menu will open; select the **dstnat** option from the list you see.

**Step 7**: Type the **MikroTik WAN IP Address**(120.50.–.198) in the **Dst. Address** input field. (Enter the IP Address from which you intend to forward all requests.)

**Step 8**: Click on the **TCP** connection protocol from the **Protocol** drop-down menu.

**Step 9**: In the **Dst Port** field, enter the port from which you decide to forward requests. Usually, the number **80** is entered in the Dst Port field because web servers run on TCP port 80.


**Step 10**: Now go to the **Action** tab.

**Step 11**: Select the **dst-nat** option from the Action drop-down list.

**Step 12**: In the **To Addresses** input field, enter the desired **web server** **IP** to which you want requests to be forwarded.

**Step 13**: Enter the port to which all requests are forwarded in the To **Ports field**. (You can type the number **80** in the To Ports field.)


**Step 14**: To confirm and save the information we entered, click **Apply** and then **Ok**.

Finally, the port forwarding configuration to the internal web server has been completed successfully. Now you can access the web server by sharing the MikroTik WAN IP in web browsers outside your local network. If you want to add new rules to port forwarding, you can add new ports or IPs by going through the steps we explained.

### Configuring Mikrotik port Forwarding to the Internal FTP server through Winbox

FTP(file transfer protocol) server is used to share files through client and server programs. You can also use port forwarding to create Nat rules allowing Internet users to access your FTP server. This configuration method is for the situation the FTP server is set in our local network, and we intend to provide access to the FTP server from the public network through the Mikrotik port forwarding configuration.

For this purpose, follow the steps we will explain to reach your goal.

**Step 1**: Log in to the Mikrotik router via Winbox as a server administrator.

**Step 2**: In the left part of the panel, by clicking on IP, select Firewall in the IP menu list.

**Step 3**: In the Firewall window, open the NAT tab.

**Step 4**: Create a New NAT Rule by clicking the PLUS (+) sign.

**Step 5**: Click on the General option, and from the chain drop-down menu, click on dstnat.

