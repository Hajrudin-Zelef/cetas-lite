---
id: collect-261001-general-networking/general-networking/2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac-3
title: "2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac.md
source_anchor: ""
source_lines: [390, 463]
sha256: 38d59db7ac04ceab997241de7a65c66d8b692364f438a9bdf5bb57e88f126512
---

# 2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac

To complete the setup, configure devices on VLAN_100 and VLAN_200 with default gateways. The default gateway for VLAN_100 is the FortiGate VLAN_100 subin- terface. The default gateway for VLAN_200 is the FortiGate VLAN_200 subinterface.


**T****es****t the configuration**

Use diagnostic commands, such as tracert, to test traffic routed through the FortiGate unit and the Cisco switch.


**T****es****t****i****n****g traffic from VLAN_100 to VLAN_200**

In this example, a route is traced between the two internal networks. The route target is a host on VLAN_200. Access a command prompt on a Windows computer on the VLAN_100 network, and enter the following command:


C:\>tracert 10.1.2.2

Tracing route to 10.1.2.2 over a maximum of 30 hops:

1 <10 ms <10 ms <10 ms 10.1.1.1

2 <10 ms <10 ms <10 ms 10.1.2.2

Trace complete.


**T****es****t****i****n****g traffic from VLAN_200 to the external network**

In this example, a route is traced from an internal network to the external network. The route target is the external network interface of the FortiGate-800 unit.

From VLAN_200, access a command prompt and enter this command:

C:\>tracert 172.16.21.2

Tracing route to 172.16.21.2 over a maximum of 30 hops:

1 <10 ms <10 ms <10 ms 10.1.2.1

2 <10 ms <10 ms <10 ms 172.16.21.2

Trace complete.

ja
Sir, i have a Question below:

1. do we need to add any routing in the Fortiget for each vlan? if yes, what will be the routing.

2. how to allow Internet in one PC only in vlan 200 (IP 10.1.2.10)

MikePost author
If you are creating the VLANs via the FortiGate then they are connected routes.

KC
Hello Mike,

I currently have a Fortinet 80C that is configured with 192.168.25.1 on port 2 (LAN) with no VDOM configured. I now need to make this a VLAN and add two other VLAN’s going to a layer 2 switch (Router on a stick). I created two of the VLAN’s but I’m unable to change 192.168.25.1 to a VLAN on the Fortinet. Can you tell me how to accomplish this. Thank you in advance for any assistance you can offer.

MikePost author
You are going to have to remove the IP from the physical interface (which uses only native tagging) and manually recreate using the create interface process. Schedule a maintenance window.

– Leave all of your existing policies alone (hopefully your address objects and VIPs etc aren’t tied to the interface specifically)

– Set physical interface the 192.168.25.1 address is currently assigned to so that it shows 0.0.0.0/0.0.0.0

– Create a new interface like you did the others, set it to VLAN xyz or whatever you like, name it accordingly, tag it accordingly, and place that IP address there

– Switch the source and destination interfaces of your previous policies accordingly so that they all transfer over 🙂

artur
hello ,

i dont fully understand why in your example the internal intrafce has an ip ? if you configuring the internal interface as a trunk port so its need to be with no address on it like 0.0.0.0/0 , and then create vlans assign the valns to the internal ports and give to the vlans you created the IP addresses.

please correct me ig i worng ?

thank you
