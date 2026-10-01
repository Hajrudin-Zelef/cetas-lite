---
id: collect-261001-cisco/cisco/configuring-bgp-neighbors-on-cisco-routers-flackbox-2
title: "configuring-bgp-neighbors-on-cisco-routers-flackbox"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/configuring-bgp-neighbors-on-cisco-routers-flackbox.md
source_anchor: ""
source_lines: [148, 256]
sha256: ad5bc9a3f74d7b9b99cba04dcb3415fae8352897d87f20299c64060f5712b38c
---

# configuring-bgp-neighbors-on-cisco-routers-flackbox

**R2(config)#router bgp 65002**

**R2(config-router)#neighbor 172.16.0.1 remote-as 65002**


#### The BGP Update-Source Command


Using loopback addresses for our neighbor statements leads to a common gotcha with BGP. BGP has a security mechanism where it will only peer with another router if it has a matching neighbor statement for that peer. The source address in the packets received from the neighbor must match the exact IP address in the neighbor statement on this router.


When a router sends packets from itself, it uses the IP address of the exit interface as the source address by default.


This is not just for BGP. This is for all traffic. If you send a ping from the command line on the router, the source address will be the address of the interface that the packet goes out of. This is going to cause BGP peering to fail between loopback addresses if we don't do something about it.


Let’s see why…


**R1(config)#router bgp 65002**

**R1(config-router)#neighbor 172.16.0.2 remote-as 65002**


**R2(config)#router bgp 65002**

**R2(config-router)#neighbor 172.16.0.1 remote-as 65002**


On R1, we point it at the neighbor of 172.16.0.2. On R2, we've got a neighbor statement for 172.16.0.1.


When R1 sends the BGP packet out to R2, it's load balanced and is either going to go along the top path or the bottom path. If it goes along the top path, it goes out interface FastEthernet 1/0 with IP address 10.0.0.1. That packet will be sent with a source address of 10.0.0.1.


When it reaches R2, R2 sees that another router is trying to form a BGP relationship with it. R2 checks to see if it has a matching BGP neighbor statement. R2 has a neighbor statement for 172.16.0.1, but it does **not** have a neighbor statement for 10.0.0.1 so the BGP peering will be rejected.


R2 is not smart enough to realise it's the same router, just using a different IP address. The IP address has to match **exactly**.


The ‘show ip bgp summary’ command will show whether the neighbor BGP peering succeeded or not. Here we can see there's been no messages sent or received and it's never come up.



What we have to do is tell R1 and R2, "When you send that BGP traffic to each other, use your loopback address rather than the address that's on the exit interface."


We do this by adding an additional command an additional command to the BGP configuration to specify the source IP address that we want to use. This overrides the default of using the IP address of the exit interface.


**R1(config)#router bgp 65002**

**R1(config-router)#neighbor 172.16.0.2 remote-as 65002**

**R1(config-router)#neighbor 172.16.0.2 update-source loopback 0**


**R2(config)#router bgp 65002**

**R2(config-router)#neighbor 172.16.0.1 remote-as 65002**

**R2(config-router)#neighbor 172.16.0.1 update-source loopback 0**


Now when R1 sends BGP traffic to R2 it uses a source address of 172.16.0.1, which matches the BGP neighbor statement on R2. And when R2 sends BGP traffic to R1 it uses a source address of 172.16.0.2, which matches the BGP neighbor statement on R1. Everything ties up on both sides so BGP peering will come up successfully.



We can see that there are being messages sent and received, and it's been up for nearly a minute. So that is all good now.

#### BGP Neighbor Verification


Notice the verification is a little bit different than with our IGPs. After we’ve configured OSPF, normally the first thing we'll do is a ‘show IP OSPF neighbor’ and for EIGRP we'll do a ‘show IP EIGRP neighbor’. That will show us a short summary of the neighbors and whether they're up or not.


But in BGP, we don't use ‘show IP BGP neighbor’ to get a summary, we use ‘show IP BGP summary’.

‘show IP BGP neighbors’ is a valid command as well, but it doesn't just give a short summary. It gives really long, verbose output.



You can use the command if you want to get more detailed information, but if you just want the summary information then ‘show IP BGP summary’ is best.


Note in the output of the ‘show IP BGP neighbors’ command above that the state of the neighbor 172.16.0.1 is ‘Established’. That is good and means the neighbor relationship is up. If you see BGP state ‘Active’, it sounds good but it's not. ‘Active’ means the router is actively trying to establish peering - and is failing. So ‘Active’ is bad, ‘Established’ is good.

#### eBGP Neighbor Configuration


Okay, so we've got our iBGP neighbors configured. The next thing to do is configure our eBGP neighbors.



We use the same configuration again. The only difference is the ‘remote-as’ on the neighbor statements is different than our own AS of 65002. This tells the router we’re configuring eBGP neighbors and to apply the eBGP rules.



Once that’s done we have all the BGP neighbors configured and working. There’s not much point in doing this on its own because they’re not sharing any routing information with each other yet. We’ll cover that in the next post in this series.

#### Additional Resources

Part 1: Why We Need BGP

Part 2: BGP Routing and Path Selection for Service Providers

Basic BGP Configuration from Cisco Press
