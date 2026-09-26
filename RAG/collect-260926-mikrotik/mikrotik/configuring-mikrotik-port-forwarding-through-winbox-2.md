---
id: collect-260926-mikrotik/mikrotik/configuring-mikrotik-port-forwarding-through-winbox-2
title: "Configuring MikroTik port forwarding through Winbox"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/tools/configuring-mikrotik-port-forwarding-through-winbox.md
source_anchor: ""
source_lines: [86, 150]
sha256: 7ee4b019f932c5d1462e2b0bb7ab867abb477d440d418cc26c5640e60559f6bd
---

# Configuring MikroTik port forwarding through Winbox

**Step 6**: In the input field Dst. Address, Enter **MikroTik WAN IP**(120.50.–.198).

**Step 7**: To send the protocol packets, you need to continue configuration by selecting the Protocol option and selecting TCP from Protocol drop-down list.

**Step 8**: In the Dst Port field, you must enter the port to which they will be connected from the public network. You can specify a port according to the server type and forward it, but FTP Server usually works on TCP port 21, so put the number 21 in this field.

**Step 9**: Select the Action tab and set it to the dst-nat option.

**Step 10**: In the **To Addresses** field, enter the IP of the **FTP server** inside the network (193.168.20.20.20).

**Step 11**: Fill the **To port** field with the port of the internal server to which we want to transfer requests and traffic. (You can enter the number **21** in this field)

**Note**: The input and output ports can be the same or different.

**Step 12**: To confirm the new rules, select apply and then OK.

The process of configuring port forwarding to your internal network’s FTP server has been completed. Now you can connect to the internal network FTP server from the Internet by entering ftp://mikrotik-wan-ip (ftp:// 120.50.–.198) in FTP clients or any web browser.

**Note**: Don’t forget to allow FTP service or TCP port 21 in your FTP server firewall to connect to the FTP server from the public network.

### Configuring Mikrotik port Forwarding to Internal SSH server through Winbox

This is how to configure Mikrotik port forwarding to connect to the SSH server in the internal network through the public network. To do this, follow the steps below:

**Step 1**: Log in to the Mikrotik server through Winbox with administrative privileges

**Step 2**: By referring to the IP option on the left side of the panel, select Firewall from the IP menu.

**Step 3**: Click the NAT tab in the Firewall window.

**Step 4**: Open the New NAT Rule window by clicking the PLUS sign (+).

**Step 5**: By selecting the General tab, select the dstnat option in the Chain drop-down menu.

**Step 6**: Enter the MikroTik WAN IP address in our network diagram (120.50.–.198) in the Dst address section.

**Step 7**: By selecting the protocol option to send packets from the Protocol list, click on the TCP option.

**Step 8**: Fill the **Dst Port** input field with the port from which you intend to transfer requests. Usually, **TCP port 22** is chosen to run the SSH server, so enter the number 22 in the Dst Port field.

**Step 9**: Now go to the Action tab and select the dst-nat option from the Action list.

**Step 10**: In the **To Addresses** field, enter the **SSH Server IP** (193.168.20.20.21).

**Step 11**: Complete the **To Ports** field with port **22**.

**Step 12**: Finally, press the Apply and Ok buttons to confirm.

Finally, you were able to configure Mikrotik port Forwarding to the Internal SSH server successfully, and if you followed the steps correctly, you can now connect to your internal SSH server through any SSH client (Putty or SSH Secure Shell Client) from the Internet.

**Note**: For communicating with the SSH server from the public network, it is necessary to allow the SSH service or TCP port 22 in the firewall of the SSH server.

There is no difference in how they work, Port Forwarding is actually the same process as Destination NAT in a Mikrotik router.

Configuring Mikrotik Port forwarding through Winbox for the web server, FTP server, and SSH server has the same basic steps. Their difference is in entering the IP address for the To Addresses and Dst Address sections and their ports. You can use the ports you specified or the commonly used ports for web servers, usually TCP80 port, for FTP server, TCP21 port, and for SSH Server, TCP22 port.

We hope that by reading this article, you will be more familiar with Mikrotik’s port forwarding function and you will be able to configure Port forwarding in Mikrotik easily. If you need guidance in this field, ask us your question in the comments section so that we can guide you well and solve your problem.

Thank you for staying with us until the end of the article.

Hi. Thanks for your efforts. I think it would be better to provide images within the article for each step if possible.

How can I do nat configuration Mikrotik?

1. Login to MikroTik RouterOS 2. Configure WAN (Public) and LAN (Private) Interfaces 3. Set Up WAN IP Address 4. Add NAT Rule for Source NAT (Masquerade) 5. Add NAT Rule for Destination NAT (Port Forwarding, if needed) 6. Verify and Save Configurations by Double-check your NAT rules using /ip firewall nat print in the CLI 7. Reboot Router
