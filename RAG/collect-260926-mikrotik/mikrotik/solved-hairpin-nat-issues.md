---
id: collect-260926-mikrotik/mikrotik/solved-hairpin-nat-issues
title: "solved-hairpin-nat-issues"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/solved-hairpin-nat-issues.md
source_anchor: ""
source_lines: [1, 133]
sha256: bf6a4c1949909ed566b6a4622d5b7a12a1a375cc9059ca1ed8bc86b0ad4e1194
---

# solved-hairpin-nat-issues

Hi,

I followed the guide on https://wiki.mikrotik.com/wiki/Hairpin_NAT, specifically:

```
/ip firewall nat
add chain=srcnat src-address=192.168.1.0/24 \
  dst-address=192.168.1.2 protocol=tcp dst-port=80 \
  out-interface=LAN action=masquerade
```

Except my network is at 192.168.0.0/24 and the server that I wish to actually connect to is at 192.168.0.10, and the port is 999, so:

```
/ip firewall nat
add chain=srcnat src-address=192.168.0.0/24 \
  dst-address=192.168.0.10 protocol=tcp dst-port=999 \
  action=masquerade
```

But it doesn’t work - the port simply appears as closed.

I also had to remove the `out-interface=LAN` part, which could be the issue - if I specified my LAN interface, I just got this error:

```
      ;;; in/out-interface matcher not possible when interface (lan) is slave - use master instead (bridge)
```

But it didn’t work either, when I put `out-interface=bridge`.

What to do now?

             
            
           
          
            
            
              Hello, try this :

/ip firewall nat

add chain=srcnat src-address=192.168.0.0/24 

dst-address=192.168.0.10 protocol=tcp dst-port=999 

out-interface=***bridge*** action=masquerade

Out interface name should be your bridge name.

Regards,

             
            
           
          
            
            
              Also make sure that you have correct dstnat rule. Common mistake is to have it with in-interface=WAN, which can’t work from LAN.

             
            
           
          
            
            
              LAN and WAN are “interface list” in the default configurations.

Bridge and up-link ethernet interface are then member of the LAN and WAN list.

Many default rules, and settings use these WAN and LAN “interface list” names (access and forward rules, NAT, MAC server, discover , … )

Interfaces as ports to the bridge are “slave” interfaces, settings must be on the bridge then, not on the interface.

But … LAN and WAN are just names, they have no other reserved meaning. (They could be assigned to any named object … interface, bridge, switch, address pool, etc ,etc)

             
            
           
          
            
            
              
This was actually the issue, oops! I suppose there’s no harm in just removing the condition?

But thank you for the other responses, I didn’t know about the interface list names.

             
            
           
          
            
            
              ```
/ip firewall nat
add chain=srcnat src-address=192.168.0.0/24 \
  dst-address=192.168.0.10 protocol=tcp dst-port=999 \
  action=masquerade
```

You are mixing stuff up…

Assuming you have a fixed wanip static

Rule1 default or normal source nat type rule

add chain=srcnat action=src-nat out-interface=wan  to address=whatever your fixed WANIP is…

Rule2  new sourcenat rule to cover off Hairpin

add chain=srcnat action=masquerade comment=“HairpinNAT” src-address=192.168.0.0/24  dst-address=192.168.0/24

Rule 3 - Destination NAT Rule

add chain=dstnat action=dst-nat  dst-address=FIXEDWANIP  dst-port=999 protocol=tcp 

to-addresses=192.168.0.10

If your WAN is dynamic, the Destination rule is a bit more complex…

add chain=srcnat action=masquerade out-interface=WAN

add chain=srcnat action=masquerade comment=“HairpinNAT” src-address=192.168.0.0/24 dst-address=192.168.0.0/24  (as above)

Destination Rule

add chain=dstnat action=dst-nat  dst-port=999 protocol=tcp dst-address=**!**192.168.0.1 

**dst-address-type=local** to-addresses=192.168.0.10

Since we don’t know how to identify directly the incoming WANIP,  we get around that with the dst-address-type local which says the destination address is on the router.

We also state that the destination address is not the subnet  (which leaves the router interface and thus the wanip).
