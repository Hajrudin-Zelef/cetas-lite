---
id: collect-261001-fortinet/fortinet/blog-how-to-configuration-pppoe-dhcp-nat-lan-ip-internet-on-fortigate-firewall-4484924c
title: "blog-how-to-configuration-pppoe-dhcp-nat-lan-ip-internet-on-fortigate-firewall-4484924c"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/blog-how-to-configuration-pppoe-dhcp-nat-lan-ip-internet-on-fortigate-firewall-4484924c.md
source_anchor: ""
source_lines: [1, 51]
sha256: 6eb583ba21e69eb01f3503d33866aaa47653945efe79cf30a513cb0f4ec9dc51
---

# blog-how-to-configuration-pppoe-dhcp-nat-lan-ip-internet-on-fortigate-firewall-4484924c

If you have a new FortiGate Firewall, You should read this article. How to Configuration PPPoE DHCP NAT LAN IP Internet On Fortigate Firewall.

We practice with a small example like above. You said a static IP address for your PC as above. The article guides pppoe dialing configuration for WAN ports on Fortigate devices.

Connect the FortiGate's interface port 2 to your ISP-supplied equipment and Connect a PC to the FortiGate using an internal port (port1).

### Enable PPPoE


Configure PPPoE dial-up connection on 1 interface with CLI. If this step is not configured, there are only 2 modes on the interface: Static and DHCP.

### PPPoE Client configuration


Configure PPPoE dialing using the web interface. Go to Network -> Select Interface -> Select the interface you want as an WAN port to dial the PPPoE -> Click Edit.

- Role: Choose " **WAN"**
- Addressing mode: Choose " **PPPoE** "
- Enter **username** and**password** provided by your carrier

Click "**OK**" to save you settings.

### Configure the internal interface

Select the interface you want as an LAN port.

In this case, DHCP is enabled. Click "OK" button to save your settings.

### Configure default route static routes

- Set the Interface to WAN
- Set the gateway address to Dynamic

### Firewall Policy & NAT


Creating a policy to allow traffic from the internal Network to the internet.

Go to Policy and Objects Firewall Policy and create a New policy.

- Set the incoming interface to the Lan interface in the outgoing interface to the internet facing interface.
- Set Source, Destination Address, Schedule, and Services is **all** .
- Make sure the Action is set to **ACCEPT** .
- Turn on **NAT** and make sure Use**Outgoing Interface Address** is selected

You can now browse the Internet using any computer that connects to the FortiGate's internal interface. I hope this article helped you connect and configure a new FortiGate so that computers can connect to the internet.


Facebook: https://www.facebook.com/routerbest

Twitter: https://twitter.com/routerbestcom
