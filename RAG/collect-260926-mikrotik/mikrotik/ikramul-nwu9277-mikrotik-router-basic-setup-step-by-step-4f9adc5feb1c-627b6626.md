---
id: collect-260926-mikrotik/mikrotik/ikramul-nwu9277-mikrotik-router-basic-setup-step-by-step-4f9adc5feb1c-627b6626
title: "ikramul-nwu9277-mikrotik-router-basic-setup-step-by-step-4f9adc5feb1c-627b6626"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/ikramul-nwu9277-mikrotik-router-basic-setup-step-by-step-4f9adc5feb1c-627b6626.md
source_anchor: ""
source_lines: [1, 25]
sha256: 22e2b1933955f235dcd118ec12bc9461680d7207ed5f8902bd7b53eacb229d14
---

# ikramul-nwu9277-mikrotik-router-basic-setup-step-by-step-4f9adc5feb1c-627b6626

MikroTik Router Basic Setup — Step by Step
1
Connect the Router to Your PC
Connect a LAN cable from any Ethernet port on the MikroTik router (usually ether2) to your PC. Plug the WAN/Internet cable into the ether1 port.
2
Download and Open Winbox
Download Winbox from MikroTik’s official website (no installation needed, runs directly). Open Winbox and go to the ‘Neighbors’ tab — your router will show up there along with its MAC address.
3
Log In to the Router
Click on the MAC address, then type ‘admin’ as the username, leave the password blank (for a new/reset router), and click Connect. On first login, an option to remove the default RouterOS configuration may appear — select ‘Remove Configuration’ for a clean setup.
4
Configure the WAN (Internet) Interface
Go to IP → DHCP Client, click the ‘+’ icon, select ether1 (WAN port) as the Interface, and click Apply/OK. This will let the router automatically get an internet IP from your ISP. If your ISP provides a static IP instead, you’ll need to manually set the IP/gateway under IP → Addresses.
5
Set the IP Address on the LAN Interface
Go to IP → Addresses and click the ‘+’ icon. Enter a local IP in the Address field, such as 192.168.88.1/24, select ether2 (LAN port) as the Interface, and click Apply. This IP will become the gateway for your local network.
6
Enable the DHCP Server
Go to IP → DHCP Server and click the ‘DHCP Setup’ button. Select ether2 in the wizard and keep clicking Next — it will automatically set the network range, gateway, and DNS. This means any device connected to the LAN will automatically receive an IP address.
7
Configure NAT (Internet Sharing)
Go to Firewall → NAT tab and click the ‘+’ icon. In the General tab, set Chain: srcnat, Out. Interface: ether1 (WAN). In the Action tab, select Action: masquerade and click Apply/OK. Without this step, devices on the LAN won’t get internet access — this is the most common mistake.
8
Test and Confirm
Now connect your PC or any device to the ether2 port. The device should automatically get an IP (in the 192.168.88.x range) and be able to browse the internet. If there’s an issue, check IP → Routes to confirm the default route (0.0.0.0/0) is set correctly.
