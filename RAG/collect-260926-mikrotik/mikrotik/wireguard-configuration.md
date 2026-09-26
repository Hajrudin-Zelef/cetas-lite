---
id: collect-260926-mikrotik/mikrotik/wireguard-configuration
title: "wireguard-configuration"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-configuration.md
source_anchor: ""
source_lines: [1, 142]
sha256: b9d0fde37de40c3d5efe2318e78803be6e376d2dfe1158debabcc7523eab43e4
---

# wireguard-configuration

(1) The first thing I would do is get rid of this rule.

Its an advanced rule that is rarely used…  The standard IP firewall rules suffice for 98% of configs.

/interface bridge settings

set use-ip-firewall=yes

(2) The second thing I would get rid of is IP filter strict… its use is definitely not compatible with dual wans and again, rarely used.

/ip settings

set rp-filter=strict

If you are concerned about optimizing the default rule set especially this rule…

*add action=drop chain=input comment=“Drop invalid input packets” 
connection-state=invalid*

Then ensure that you go to IP Firewall menu selection, choose the Connections Tab, and find the **Tracking** sub tab/button (second line).

Then ensure connection tracking is set to auto and the checkbox for LOOSE TRACKING is **NOT** SELECTED, which means it will be set for strict tracking.

(3) Your firewall rules are out of control IMHO.   Also disorganized as you mix forward chain and input chains together.  My recommendation is to simplify.

ex. look at last two rules you have… a symptom of a shit show…

add action=drop chain=input comment=“defconf: drop all from WAN” 

in-interface=ether1-WAN

add action=drop chain=forward comment=“defconf: drop all from WAN” 

in-interface=ether1-WAN



90% of your rules I would throw in the toilet,  not even sure what L2TP brute force protection is about?  If you are using an insecure VPN such as PPP, dont use that type of VPN.  I would focus on

a.  actual required rules first in an organized fashion.

b.  add rules for better security as required

—>  what is SSH used for and by whom?

—>  are you running any servers that require extra security?

—> what issues have you had in the past that drove you to this extreme?

(4) Now lets focus on wireguard 

a.  why are you using port 443 for wireguard ??  443 is a port used for protocol HTTPS traffic?

Suggest you change it.

b.  I dont understand your masquerade rules…

**Lets take a look at the first two:**

*/ip firewall nat
add action=masquerade chain=srcnat comment=“defconf: masquerade” 
out-interface=ether1-WAN
add action=masquerade chain=srcnat comment=“LAN > WG” out-interface=WG_to_VPS*

When viewed in conjunction with an interface list member you have.

add interface=WG_to_VPS list=WAN

This means that any traffic heading out the tunnel will be given the IP address of the interface  10.8.0.2

So what is the point of the second rule???

Finally,  do you really need to masquerade your traffic heading out of the tunnel??

You control the VPS correct?  If so, then you dont have to masquerade your LAN traffic from the MT to the VPS?

All you need to do is ensure the subnets are identified as allowed IPs on the VPS for the MT client peer settings on the VPS

addresses=10.8.0.2/32,  subnetA, subnetB   etc…

I mean its personal choice, as long as you understand what you are doing to the traffic, either way works but at least get rid of the second rule.

++++++++++++++++++++

I have no clue as to what is intended by the next two rules…

Can you please explain what is being accomplished here in some detail??

_add action=netmap chain=dstnat comment=“WG > PC1 RDP” dst-port=40001 

in-interface=WG_to_VPS protocol=tcp to-addresses=192.168.33.2 to-ports=

3389

add action=netmap chain=dstnat comment=“WG > PC2 RDP” dst-port=40002 

in-interface=WG_to_VPS protocol=tcp to-addresses=192.168.33.3 to-ports=_

3389



c.  Since you dont have drop all rules that I can see and the firewall is too messy I will assume there is nothing blocking LAN to WG interface traffic.

d.  Whats left is routes.  All I see is disabled routes… ???

Thus what you should have currently in your IP Tables, which is not shown in the config…

  dst-address=0.0.0.0/0  gateway=ISPgateway_IP  table=main

  dst-address=192.168.33.0/24  gateway=bridge1  table=main

You have currently created the required components for WG.

/routing rule

add action=lookup src-address=192.168.33.0/24 table=use-WG-table

/routing table

add fib name=use-WG-table

Now you have to add the IP route ( i see its there but disabled, so just enable it! )

/ip route

add dst-address=0.0.0.0/0  gateway=WG_to_VPS   table=use-WG-table

Now lets look at the logic,

All traffic coming from the LAN is captured and sent out WG, so to ensure we dont inadvertently do that for LAN to LAN traffic, add a another route rule ORDER IS KEY SUCH THAT..

/routing rule

**add action=lookup-only-in-table dst-address=192.168.33.0/24 table=main
add action=lookup src-address=192.168.33.0/24 table=use-WG-table**

Another option is that we could also simply add another route manually similar to the DAC route and create the same route on the other table (VS_to_WG)

/ip route

add dst-address=192.68.33.0/24  gateway=bridge1 table=WG_to VS
