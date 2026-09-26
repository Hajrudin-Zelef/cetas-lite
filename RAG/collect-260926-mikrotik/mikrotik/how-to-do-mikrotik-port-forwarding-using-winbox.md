---
id: collect-260926-mikrotik/mikrotik/how-to-do-mikrotik-port-forwarding-using-winbox
title: "how-to-do-mikrotik-port-forwarding-using-winbox"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/tools/how-to-do-mikrotik-port-forwarding-using-winbox.md
source_anchor: ""
source_lines: [1, 37]
sha256: a46580cc70ca37f6a02dee887086fddca2593c100107424eb532dc04ca24b896
---

# how-to-do-mikrotik-port-forwarding-using-winbox

1

Ans

3 years ago

MikroTik Port Forwarding is a procedure of capturing the movement of data ruled for a processor’s port grouping and forwarding the data to a diverse IP and port. To do MikroTik Port Forwarding using Winbox you need a public IP from your ISP, a network set up behind a router and a firewall and a proxy server.

Now login to the Winbox and choose Firewall by clicking on IP. Then click on NAT then tick on +symbol, by clicking on the general tab pick ‘dstnat’ from chain drop-down list. Target address and choose DST-nat.

Finally click on apply and hence your MikroTik Port Forwarding using Winbox is completed.

It is the procedure of capturing the data movement ruled for a processer's port grouping and forwarding it to a diverse IP and port. This can be completed by means of a MikroTik-router. This complete artifact is a guide on how to do the port-forwarding on a MikroTik router. What you want to set up the Port-Forwarding

**Prior to getting to that, do imagining a state as follows:**

Visualize an IT manager. You have formed a big net, and someone needs to distantly link to your VPS server and work tenuously. You cannot share the IP address with him for the safety drives. What would you prepare? In this state, you should use up the port-forwarding on the MikroTik Router to deal with all requests.

**Step 1: Login to the Winbox** & Click on IP, then **choose** Firewall.

**Step 2:** Click on **NAT** and then tick on the **+ symbol**.

**Step 3:** Click on the General tab. Pick ‘dstnat’ from the chain drop-down list. In the “**Dst.-Address**” field, write this IP (10.10.10.10). From the Protocol list, **pick** the protocol link like (TCP, xdp, ddp). In the “Dst.-Port” field, enter 5847.

**Step 4:** Target Address: Write your communal IP as given by your ISP.

**Step 5:** Next, go to Action.

**Step 6:** Action: **choose ‘DST-na**t’

**Step 7:** Click **Appy** & **OK.**

If you have set-up a device, that is, a DVR, or might be a Biometric device at your workplace or cafeteria and would like to access it outwardly, you will first need a communal inert Internet Protocol (IP) on your internet linking, and then, port forwarding guidelines have to be set-up on the router, so that the traffic, directed to that still IP over those specific ports, is furthered to the Biometric device itself.

MikroTik is a Latvian web apparatus industrial company. MikroTik has a remarkably high level of flexibility when it comes to net supervision with its own classy router. The Router OS can be connected on a PC and it will change it into a router with all the essential facets - direction-finding, firewall, bandwidth control, wireless entrée point, hotspot entryway, VPN server, and more.

Router Operating System has all the needed facets of an internet service provider. MikroTik is able to manage everything related to networking and in this post, we will concentrate on how to effectively apply the Port forwarding on MikroTik by using Router OS.
