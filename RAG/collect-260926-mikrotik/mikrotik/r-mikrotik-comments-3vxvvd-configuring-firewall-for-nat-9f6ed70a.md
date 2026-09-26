---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-3vxvvd-configuring-firewall-for-nat-9f6ed70a
title: "r-mikrotik-comments-3vxvvd-configuring-firewall-for-nat-9f6ed70a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/r-mikrotik-comments-3vxvvd-configuring-firewall-for-nat-9f6ed70a.md
source_anchor: ""
source_lines: [1, 102]
sha256: 6007225578aaba315ca8f69c018570c1745b8f375bf2f53e378f7907997ff0aa
---

# r-mikrotik-comments-3vxvvd-configuring-firewall-for-nat-9f6ed70a

Configuring firewall for NAT? 
        
    Hi all,
I've got a question regarding my firewall configuration and NAT, with a bonus interface question.
Here are my firewall filters:
[admin@MikroTik] /ip> /ip firewall filter print 
Flags: X - disabled, I - invalid, D - dynamic 
 0  D ;;; special dummy rule to show fasttrack counters
      chain=forward 
 1    ;;; invalid
      chain=input action=drop connection-state=invalid log=no 
      log-prefix="" 
 2    ;;; internal network
      chain=input action=accept src-address=192.168.88.0/24 
      in-interface=!ether1-gateway log=no log-prefix="" 
 3    ;;; established and related
      chain=input action=accept 
      connection-state=established,related log=no log-prefix="" 
 4    ;;; pings
      chain=input action=accept protocol=icmp log=no 
      log-prefix="" 
 5    ;;; Drop everything else
      chain=input action=drop log=no log-prefix="" 
 6    ;;; invalid
      chain=forward action=drop connection-state=invalid log=no 
      log-prefix="" 
 7    chain=forward action=fasttrack-connection 
      connection-state=established,related log=no log-prefix="" 
 8    ;;; Established and related
      chain=forward action=accept 
      connection-state=established,related log=no log-prefix="" 
 9    ;;; internal network
      chain=forward action=accept src-address=192.168.88.0/24 
      in-interface=!ether1-gateway log=no log-prefix="" 
10    ;;; Drop Bogon
      chain=forward action=drop src-address-list=Bogon 
      in-interface=ether1-gateway log=yes 
      log-prefix="Bogon Forward Drop"     Here is my NAT table:
[admin@MikroTik] /ip> /ip firewall nat print 
Flags: X - disabled, I - invalid, D - dynamic 
 0    ;;; default configuration
      chain=srcnat action=masquerade to-addresses=0.0.0.0 
      out-interface=ether1-gateway log=no log-prefix="" 
 1    chain=dstnat action=dst-nat to-addresses=192.168.88.246 
      to-ports=16881 protocol=tcp dst-address-type=local 
      dst-port=16881 log=no log-prefix="" 
 2    chain=dstnat action=dst-nat to-addresses=192.168.88.246 
      to-ports=5000-5006 protocol=tcp dst-port=5000-5006 log=no 
      log-prefix="" 
 3  D chain=dstnat action=dst-nat to-addresses=192.168.88.226 
      to-ports=63760 protocol=udp dst-address=70.171.209.165 
      dst-port=63760 log=no log-prefix="" 
 4  D chain=dstnat action=dst-nat to-addresses=192.168.88.251 
      to-ports=22818 protocol=udp dst-address=70.171.209.165 
      dst-port=22818 log=no log-prefix="" 
 5  D chain=dstnat action=dst-nat to-addresses=192.168.88.251 
      to-ports=22818 protocol=tcp dst-address=70.171.209.165 
      dst-port=22818 log=no log-prefix="" 
 6  D chain=dstnat action=dst-nat to-addresses=192.168.88.226 
      to-ports=52514 protocol=udp dst-address=70.171.209.165 
      dst-port=52514 log=no log-prefix="" My question is: I've read various basic firewall tutorials and it's my understanding that one has to add a final firewall filter that drops all chain=forward traffic if it doesn't match any of the other chain=forward rules. However, in this case, if I add such a rule, it seems to block any incoming NAT'd traffic. I think I need to add a rule to allow incoming NAT traffic but I'm not sure if that's true or how to do that?
Bonus question: I have an RB2011. Here is my interface list. I read in an earlier post that only master ports should be connected to the bridge, with all other ports slave to a master. This is the default interface for the RB2011. It shows ether2-5 to be connected to the bridge with no master port. Is this correct, or should I reconfigure ether2-5 to be slave to ether1 master, with ether1 connected to bridge?
[admin@MikroTik] > /interface ethernet print 
Flags: X - disabled, R - running, S - slave 
 #    NAME        MTU MAC-ADDRESS       ARP        MASTER-PORT      SWITCH     
 0 R  ether1...  1500 4C:5E:0C:44:97:F4 enabled    none             switch1    
 1 RS ether2     1500 4C:5E:0C:44:97:F5 enabled    none             switch1    
 2 RS ether3     1500 4C:5E:0C:44:97:F6 enabled    none             switch1    
 3 RS ether4     1500 4C:5E:0C:44:97:F7 enabled    none             switch1    
 4 RS ether5     1500 4C:5E:0C:44:97:F8 enabled    none             switch1    
 5  S ether6...  1500 4C:5E:0C:44:97:F9 enabled    none             switch2    
 6  S ether7...  1500 4C:5E:0C:44:97:FA enabled    ether6-master... switch2    
 7  S ether8...  1500 4C:5E:0C:44:97:FB enabled    ether6-master... switch2    
 8  S ether9...  1500 4C:5E:0C:44:97:FC enabled    ether6-master... switch2    
 9  S ether1...  1500 4C:5E:0C:44:97:FD enabled    ether6-master... switch2    
10 XS sfp1       1500 4C:5E:0C:44:97:F3 enabled    none             switch1
[admin@MikroTik] > /interface print 
Flags: D - dynamic, X - disabled, R - running, S - slave 
 #     NAME                                TYPE       ACTUAL-MTU L2MTU
 0  R  ether1-gateway                      ether            1500  1598
 1  RS ether2                              ether            1500  1598
 2  RS ether3                              ether            1500  1598
 3  RS ether4                              ether            1500  1598
 4  RS ether5                              ether            1500  1598
 5   S ether6-master-local                 ether            1500  1598
 6   S ether7-slave-local                  ether            1500  1598
 7   S ether8-slave-local                  ether            1500  1598
 8   S ether9-slave-local                  ether            1500  1598
 9   S ether10-slave-local                 ether            1500  1598
10  XS sfp1                                ether            1500  1598
11  R  bridge-local                        bridge           1500  1598
Section des commentaires
Disable all your rules. Re-enable them one at a time until your traffic stops.
My apologies if I wasn't being clear. My traffic, with the rules as posted, gets through just fine.
However, it's my understanding that the default behavior is that packets are allowed through the firewall unless specifically denied. Without a final rule in the forward chain of "chain=forward action=drop", then having a forward chain firewall is essentially pointless, because the router will allow all packets (except for invalid and bogon as I've specified).
However, if I add such a rule to my firewall, it seems to have the undesirable effect of also preventing any forward packets from hitting the NAT table. I want packets to hit the NAT table and be forwarded to the appropriate IP address if the port matches. What is the rule I need to add to allow such traffic before the final rule "chain=forward action=drop"?
add chain=forward action=accept comment="allow port forwards from dstnat" protocol=tcp dst-port="16881,63760,22818,52514"
add chain=forward action=accept comment="allow port forwards from dstnat" protocol=tcp dst-port="5000-5006"
I'm not sure that 2nd rule will work. You may have to do "5000,5001,5002..." etc.
Rules 4 & 5 in your Nat list appear to be duplicates
Ether1 is is for your isp router in the default rb2011 config. You could set ether2 as a master and slave 3,4,5 to it.
Having ports slaved saves CPU workload by having the switch chips do a lot of the work. This only matters if your router is going to busy.
