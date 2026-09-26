---
id: collect-260926-mikrotik/mikrotik/mikrotik-port-forwarding-routeros-port-forwarding-1
title: "Mikrotik Port Forwarding - RouterOS Port Forwarding"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/mikrotik-port-forwarding-routeros-port-forwarding.md
source_anchor: ""
source_lines: [1, 112]
sha256: 58f95daebb0b43c3cbfd544e989fa77bb20d3ccc9d42c28de4a0b4ac569e4403
---

# Mikrotik Port Forwarding - RouterOS Port Forwarding

MikroTik RouterOS is the stand-alone **Linux operating system** that is used with MikroTik’s networking equipment. However, that is not the only function it can perform. It is more than just an OS for routers. In fact, this software can even be installed on regular PCs to turn them into dedicated routers.

Nevertheless, the blog is a step-by-step tutorial on **how to configure port forwarding in MikroTik**. Before we get into it, however, we will first discuss a little more about the OS itself and give you a general idea of what port forwarding id & Mikrotik port Forwarding, and RouterOS Port Forwarding actually is.

## What is Mikrotik?

MikroTik itself is a Latvian network equipment manufacturing company. They develop and sell wired and wireless network routers, network switches, access points, operating systems, and auxiliary software for their products.

MikroTik’s RouterOS is the operating system that powers its devices and has a very high level of flexibility when it comes to network management. RouterOS can also be installed on a PC turning it into a router with all the necessary features - routing, firewall, bandwidth management, wireless access point, backhaul link, hotspot gateway, VPN server, and more.

The feature that we will delve into today is RouterOS port forwarding. Unlike with other operating systems, the way to set it up might not be very intuitive, but not difficult at its core. Before setting up the configuration, let us first explain what port forwarding is.

## What is Port forwarding?

Port forwarding is the process of intercepting data traffic headed for a computer's IP/port combination and redirecting it to a different IP and/or port. This process can be easily done using a MikroTik router or any system running RouterOS.

Before we get to that, try to imagine the following situation:

You are an IT administrator. You created a large network, and someone wants to connect remotely to your VPS server or dedicated server to work remotely. You can’t share the server IP with that person for security reasons. What should you do? In this situation, you should use port forwarding on the MikroTik router to handle all the requests. Also, You can check out our Mikrotik VPS Server.

## How to Configure Mikrotik Port Forwarding?

In this section, we'll provide detailed step-by-step instructions for configuring MikroTik port forwarding for various purposes. We'll cover three common scenarios: forwarding to an internal web server, FTP server, and SSH server. Each set of instructions will be presented clearly and concisely, catering to readers with varying levels of technical expertise.

**Step 1:** Log in to your own MikroTik server with admin privileges.

**Step 2:** Click on **IP** from the left side panel.

**Step 3:** In the newly opened submenu, click on **Firewall**.

**Step 4:** Head over to the **NAT** tab in the Firewall window.

**Step 5:** Click on the **+** button to create a new rule.

**Note**: In this scenario, assume the router connects to IP (10.10.10.10), and we want to forward all requests from (10.10.10.10:5847) to the (20.20.20.20:4324).

**Step 6:** Click on the **General** tab and select **dstnat** from the **chain** drop-down list.

**Step 7:** In the **Dst. Address** field, type the IP you wish to forward all requests from (i.e., 10.10.10.10 in our case).

**Step 8:** From the **Protocol** list, select the connection protocol, such as TCP.

**Step 9:** In the **Dst. Port** field, type the port you wish to forward requests from (i.e., 5847 in this example).**Step 10:** Now, navigate to the **Action** tab.

**Step 11:** From the **Action** drop-down list, select **dst-nat**.

**Step 12:** In the **To Addresses** field, type the IP to which you wish to forward all requests (i.e., 20.20.20.20 in our case).

**Step 13:** In the **To Ports** field, type the port you want to forward requests to (i.e., 4324 in this example).

**Step 14:** Click on **Apply** and then on **OK** to save and add the new rule.

And that’s it. You have successfully configured your first port forwarding rule on MikroTik. To add new port forwarding rules, simply follow the steps with new ports or IPs.

## Troubleshooting Tips

While configuring port forwarding is a valuable skill, issues may arise during the setup process. Here are some common problems users may encounter and tips on how to troubleshoot them:

### Port Conflicts

Problem: Another device or application is already using the port you intend to forward.

Solution: Check for port conflicts by verifying that the chosen port is not in use elsewhere on your network. Consider changing the forwarded port if necessary.

### Incorrect IP Address

Problem: Forwarding to the wrong internal IP address can lead to connection failures.

Solution: Double-check and ensure that you've entered the correct internal IP address of the target device or server in your port forwarding rule.

### Firewall Blocking

Problem: The target server's firewall may be blocking incoming traffic on the forwarded port.

Solution: Adjust the firewall settings on the target server to allow incoming connections on the forwarded port. Ensure both your router's and server's firewalls are correctly configured.

### Dynamic IP Address

Problem: If your internal device has a dynamic IP address that changes, the port forwarding rule may become invalid.

Solution: Configure your router to assign a static or reserved IP address to the internal device. This ensures that the port forwarding rule remains effective even if the IP address changes.

### ISP Restrictions

Problem: Some Internet Service Providers (ISPs) may impose restrictions on certain ports.

Solution: Contact your ISP to inquire about any port restrictions. Consider using alternative ports for forwarding if necessary.

## Common Use Cases of Port Forwarding

Port forwarding is a versatile networking technique with a wide range of applications. Here, we'll explore some common scenarios where port forwarding becomes a crucial tool:

### Remote Desktop Access

In today's interconnected world, the need for remote access to computers is paramount. Whether you're an IT administrator managing servers or a remote worker accessing your office desktop, port forwarding enables secure access to remote desktops. By forwarding the appropriate port (e.g., TCP port 3389 for RDP) on your MikroTik router to the target computer's IP, you can access it from anywhere, enhancing productivity and flexibility.

### Online Gaming

Online gaming enthusiasts often rely on port forwarding to achieve optimal gaming experiences. Popular multiplayer games, like those on gaming consoles or PC platforms, require specific ports to be open for seamless connectivity. Port forwarding ensures that game traffic flows smoothly between players, reducing lag and enhancing the gaming experience.

Optimize your gameplay with a **Gaming VPS Server**! Deploy yours now and elevate your gaming experience with powerful resources. Explore the potential of Linux commands for ultimate customization and control.

### Web Server Hosting

For businesses and individuals hosting websites or web applications, port forwarding is indispensable. By forwarding HTTP (port 80) or HTTPS (port 443) requests to a web server's internal IP, you can make your website accessible to users on the internet. This is essential for businesses looking to establish an online presence and for developers testing their web applications.

### File Sharing and FTP Servers

Sharing files across the internet or running an FTP server for file transfers often requires port forwarding. Whether you're sharing files with colleagues, clients, or friends, port forwarding allows external users to connect to your FTP server securely. FTP typically uses port 21 for control connections and a range of ports for data transfers, making precise port forwarding crucial.

