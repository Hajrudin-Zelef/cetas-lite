---
id: collect-260926-mikrotik/mikrotik/request-for-review-mikrotik-firewall-filter-rules
title: "request-for-review-mikrotik-firewall-filter-rules"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/request-for-review-mikrotik-firewall-filter-rules.md
source_anchor: ""
source_lines: [1, 96]
sha256: 4810eec76d3197b385625d9ac09f484cce52828604e6daf29d77bb887abbf9be
---

# request-for-review-mikrotik-firewall-filter-rules

I am seeking feedback and suggestions to improve my current MikroTik firewall filter rules setup. Below are the rules I have configured on my router according AI suggestion:

```
/ip/firewall/filter/print                                                                     
Flags: X - disabled, I - invalid; D - dynamic 
 0  D ;;; special dummy rule to show fasttrack counters
      chain=forward action=passthrough 
 1    ;;; defconf: accept established, related, untracked
      chain=input action=accept connection-state=established,related,untracked 
 2    ;;; defconf: drop invalid
      chain=input action=drop connection-state=invalid 
 3    ;;; defconf: accept ICMP
      chain=input action=accept protocol=icmp 
 4    ;;; defconf: accept to local loopback (for CAPsMAN)
      chain=input action=accept dst-address=127.0.0.1 
 5    ;;; defconf: drop all not coming from LAN
      chain=input action=drop in-interface-list=!LAN 
 6    ;;; defconf: accept in IPsec policy
      chain=forward action=accept ipsec-policy=in,ipsec 
 7    ;;; defconf: accept out IPsec policy
      chain=forward action=accept ipsec-policy=out,ipsec 
 8    ;;; defconf: fasttrack
      chain=forward action=fasttrack-connection hw-offload=yes connection-state=established,related 
 9    ;;; defconf: accept established, related, untracked
      chain=forward action=accept connection-state=established,related,untracked 
10    ;;; defconf: drop invalid
      chain=forward action=drop connection-state=invalid 
11    ;;; defconf: drop all from WAN not DSTNATed
      chain=forward action=drop connection-state=new connection-nat-state=!dstnat in-interface-list=WAN 
12    ;;; accept IGMP traffic for multicast
      chain=input action=accept protocol=igmp 
13    ;;; accept IGMP traffic for multicast forwarding
      chain=forward action=accept protocol=igmp 
14    ;;; accept UDP traffic for multicast
      chain=forward action=accept protocol=udp dst-address=224.0.0.0/4 
15    ;;; accept UDP traffic for mDNS (Multicast DNS)
      chain=input action=accept protocol=udp dst-port=5353 
16    ;;; accept UDP traffic for mDNS (Multicast DNS) forwarding
      chain=forward action=accept protocol=udp dst-port=5353
```




- My home network has a Google Chromecast, which needs to communicate smoothly within the LAN.
- I aim to balance security and functionality, ensuring essential services work without exposing the network to unnecessary risks.

I appreciate your time and expertise in reviewing my setup. Any advice or recommendations to improve these firewall rules would be greatly appreciated.

             
            
           
          
            
            
              Thank you to everyone who has viewed my initial post. I haven’t received any responses yet, so I wanted to add some clarifying questions to hopefully garner some feedback.

Specifically, I am interested in the following:

1. 
**Multicast Rules Positioning:** The rules for handling multicast traffic (rules 12, 13, 14, 15, and 16) are located at the end of my firewall filter list. Will these rules be processed correctly in this position, or should they be moved to ensure they function as intended?
2. 
**Security Concerns:** Could these multicast rules potentially pose a security threat to my network? If so, what modifications would you recommend to mitigate any risks?
3. 
**Relevance for Chromecast:** How relevant are these multicast rules for ensuring that my Google Chromecast functions properly on my network? Are there any additional or alternative rules that would better support Chromecast while maintaining network security?

 
            
           
          
            
            
              
I got a good chuckle off this. Thank you.

To OP; my immediate question would be as to why you've attempted opened port 5353 on the router?

```
15    ;;; accept UDP traffic for mDNS (Multicast DNS)
      chain=input action=accept protocol=udp dst-port=5353 
      
16    ;;; accept UDP traffic for mDNS (Multicast DNS) forwarding
      chain=forward action=accept protocol=udp dst-port=5353
```

The *input* chain is for requests **to** the router. *forward* is **through** the router. Thankfully, rule 15 won't expose your router to the internet since rule 5 does drop requests from WAN that aren't dstnatted, but something as simple as reordering your firewall rules could cause issues.

Likewise, rule 11 should prevent rule 16 from being actioned for the WAN.

Generally, I'd say your rules aren't specific enough. You need to restrict rules based on interfaces to be a bit more secure and so that you're not opening your network for attacks when reordering rules.

For now I recommend disabling your 5 rules (12 - 16) and reading through the following guide to understand the firewall setup.

There is also this site that contains Iptables documentation. RouterOS uses Iptables behind the scenes for the firewall. This contains more in-depth info for things like chains, dstnat, srcnat etc.

https://www.frozentux.net/iptables-tutorial/chunkyhtml/c962.html
