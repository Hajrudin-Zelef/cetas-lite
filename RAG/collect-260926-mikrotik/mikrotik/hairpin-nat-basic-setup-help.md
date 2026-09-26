---
id: collect-260926-mikrotik/mikrotik/hairpin-nat-basic-setup-help
title: "hairpin-nat-basic-setup-help"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/hairpin-nat-basic-setup-help.md
source_anchor: ""
source_lines: [1, 65]
sha256: f388ebd4306c34da94423581672836f0c911caa004709f786c3a46c945d04eea
---

# hairpin-nat-basic-setup-help

Hi guys,

this configuration doesn’t seem to work for Hairpin NAT. Client is .227 and Server .248




```
/ip firewall nat print
Flags: X - disabled, I - invalid; D - dynamic 
 0    ;;; Hairpin NAT
      chain=srcnat action=masquerade protocol=udp src-address=192.168.88.0/24 
      dst-address=192.168.88.0/24 log=no log-prefix="" 
 1    ;;; defconf: masquerade
      chain=srcnat action=masquerade out-interface-list=WAN 
      ipsec-policy=out,none 
 2    ;;; p1
      chain=dstnat action=dst-nat to-addresses=192.168.88.248 to-ports=7777 
      protocol=udp in-interface-list=WAN dst-port=7777 log=no log-prefix="" 
 3    ;;; p2
      chain=dstnat action=dst-nat to-addresses=192.168.88.248 to-ports=8177 
      protocol=udp in-interface-list=WAN dst-port=8177 log=no log-prefix="" 
 4    ;;; p3
      chain=dstnat action=dst-nat to-addresses=192.168.88.248 to-ports=25565 
      protocol=tcp in-interface-list=WAN dst-port=25565 log=no log-prefix=""
```

Any idea what could be wrong, I also tried another rule like described in the official video tutorial, which also didn’t work.

https://www.youtube.com/watch?v=1I5FywY6opQ




```
;;; YT rule
      chain=srcnat action=masquerade protocol=udp src-address=192.168.88.0/24 
      dst-address=192.168.88.248 out-interface=bridge log=no log-prefix=""
```

I am also uncertain where the rule has to be placed, cannot find info about that in the documentation.

             
            
           
          
            
            
              If anyone is interested, the issue was using in-interface-list=WAN for the dstnat. This way it was not done after the hairpin, because the request came from inside and not from outside WAN network.

             
            
           
          
            
            
              https://forum.mikrotik.com/viewtopic.php?t=179343



Yes,  the dst-nat rules require something other than in-interface-lis=WAN ( we are talking dynamic WANIPs here) and as well one needs to ensure the firewall rule is agnostic.

Best is

*add chain=forward action=accept connection-nat-state=dstnat*
