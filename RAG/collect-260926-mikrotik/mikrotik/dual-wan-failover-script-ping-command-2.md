---
id: collect-260926-mikrotik/mikrotik/dual-wan-failover-script-ping-command-2
title: "software = RouterOS 6.47.10"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/dual-wan-failover-script-ping-command.md
source_anchor: ""
source_lines: [161, 336]
sha256: 20fd325a2932d877e29056591da1bb9e6e0d46d52f8ad8e2a19855d79cd3da5b
---

# software = RouterOS 6.47.10

```
#
# software = RouterOS 6.47.10
# model    = RouterBOARD 750UP r2 (hEX PoE Lite)
###########################################################
# Test router representing two ISP providers.
# ISP1 on ether3 and ISP2 on ether4
# See the two firewall rules that simulate network down.
###########################################################
/system identity set name=ISP_Provider
/interface ethernet
set [ find default-name=ether3 ] comment=ISP1
set [ find default-name=ether4 ] comment=ISP2
###########################################################
# Example environment provided by this Test ISP
###########################################################
/interface bridge
add fast-forward=no name=BR1 protocol-mode=none vlan-filtering=yes
/interface vlan
add interface=BR1 name=VLAN_ISP1 vlan-id=30
add interface=BR1 name=VLAN_ISP2 vlan-id=40
/interface list
add name=WAN
add name=VLAN
add name=BASE
/ip pool
add name=POOL1 ranges=10.3.30.2-10.3.30.254
add name=POOL2 ranges=10.4.40.2-10.4.40.254
/ip dhcp-server
add address-pool=POOL1 disabled=no interface=VLAN_ISP1 name=DHCP1
add address-pool=POOL2 disabled=no interface=VLAN_ISP2 name=DHCP2
/interface bridge port
add bridge=BR1 interface=ether3 pvid=30
add bridge=BR1 interface=ether4 pvid=40
/interface bridge vlan
add bridge=BR1 tagged=BR1 vlan-ids=30
add bridge=BR1 tagged=BR1 vlan-ids=40
/interface list member
add interface=ether1 list=WAN
add interface=VLAN_ISP1 list=VLAN
add interface=VLAN_ISP2 list=VLAN
/ip address
add address=10.3.30.1/24 interface=VLAN_ISP1 network=10.3.30.0
add address=10.4.40.1/24 interface=VLAN_ISP2 network=10.4.40.0
/ip dns set allow-remote-requests=yes
/ip dhcp-client add dhcp-options=hostname disabled=no interface=ether1
/ip dhcp-server network
add address=10.3.30.0/24 dns-server=10.3.30.1 gateway=10.3.30.1
add address=10.4.40.0/24 dns-server=10.4.40.1 gateway=10.4.40.1
/ip firewall filter
add action=accept chain=input connection-state=established,related comment="Allow Estab & Related"
add action=accept chain=input in-interface-list=VLAN comment="Allow VLANs"
add action=drop chain=input comment=Drop
###########################################################
# These two rules test causing failover to occur
###########################################################
add action=drop chain=forward disabled=yes dst-address=8.8.8.8 src-address=10.3.30.0/24 comment="Enable to test Host1 Failure"
add action=drop chain=forward disabled=yes dst-address=9.9.9.9 src-address=10.3.30.0/24 comment="Enable to test Host2 Failure"
add action=accept chain=forward connection-state=established,related comment="Allow Estab & Related"
add action=accept chain=forward connection-state=new in-interface-list=VLAN out-interface-list=WAN comment="Allow VLANs"
add action=drop chain=forward comment=Drop
/ip firewall nat add action=masquerade chain=srcnat comment="Default masquerade" out-interface-list=WAN
```

             
            
           
          
            
            
              
how do I know when to fire a script to do this?


When the failover is active… and when go back online the main line, and the failover is not active

```
:global something
:if ([:len [/ip route find where comment="ISP2" and active=yes]] > 0) do={
    :if ($something != true) do={
        /ip fire conn
        :foreach idc in=[find where timeout>60] do={ remove [find where .id=$idc] }
        :set something true
    }
} else={
    :if ($something != false) do={
        /ip fire conn
        :foreach idc in=[find where timeout>60] do={ remove [find where .id=$idc] }
        :set something false
    }
}
```

             
            
           
          
            
            
              rextended,

Oh my goodness. This is awesome! It works excellent. You should make a separate post about WAN Failover and update the link in your signature to point to that new dedicated topic. It takes a good while to write up topics, so no pressure. Just a grateful user.

Note, I changed the timeout to 0, vs 60. This is because ping sessions timeout at 10s by default. I suppose people could tune the value to their application behavior.

             
            
           
          
            
            
              NO, leave it to 60 seconds, or at least 20 seconds (check gateway 10 ping off, 10 ping on)

I do not write the things without reason.

If just one connection on connection tracking  is already closed for timeout (or other reasons) during the execution of the clean,

the script stop with error because when try the connection is already closed, and do not finish his works.

             
            
           
          
            
            
              
Okay, if that is the case, would it be possible to close connections in a sub function? Perhaps store in memory (whatever MikroTik scripting calls it) a list of all connections, then loop over that list closing them down? If the connection is no longer present, move on to the next item in the list? I need > 9 seconds of closing capability.

             
            
           
          
            
            
              
???

I do not insist further, I have already written you the script that does the right job,

based on the real traffic of equipment in production and not only theoretically simulated.

             
            
           
          
            
            
              
Its okay, I can sort it out. If you have a ping session, not stop (ping 8.8.8.8 -t), when the change over occurs it will hang because it is on a 9s timeout.

             
            
           
          
            
            
              pcunite, I do not understand what you are discussing regarding clearing connections.

Is this something I should be worried about on my setup???

Typically my issue is not failover perse but when the primary come back online,  the router was not able to route the traffic properly if the gateway had changed.

THus my script is to ensure that the new gateway is used for routing.

Where does your connection revelation belong in this??

             
            
           
          
            
            
              
Well, I don’t know. It comes down to how ISP1 fails and the applications you are using. After the failover to the different ISP, if your application times out on its own, you’ll be fine. However, if you have long running applications (VPN or a never ending ping session), those packet flows will not automatically change over if the interface is still up. At least in my testing they did not. So, I just go ahead and clear out any and all connections. This provides the same effect had you disabled the interface.

I think most people must test their failover by unplugging a cable or disabling an interface. In my case, I’m testing hosts that go down, and that is why I want the change over to occur. For all intents, the first ISP still works, in some ways. However, I want to fail over because the second ISP gives me access to those hosts. Changing the route out from under all the other connections leaves them in an unstable state. So, I close them out. Then the applications reconnect on their own.
