---
id: collect-260926-mikrotik/mikrotik/mikrotik-firewall-not-working-as-expected
title: "mikrotik-firewall-not-working-as-expected"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/mikrotik-firewall-not-working-as-expected.md
source_anchor: ""
source_lines: [1, 128]
sha256: a2c6453f33bc700baf5771a8e5f0659c929f52bcf5fc1f085f1db70b35aba7a7
---

# mikrotik-firewall-not-working-as-expected

If firewall rules you posted (and I quoted below) are indeed all of them … then no wonder it doesn’t work for you. I highlited the rule which makes me think you’ve had many more rules on your old device … it refers to kid-control firewall chain … which doesn’t exist in shown config.

If I were you, I’d take this router off-line imediately and rework the config. This time keeping as much default setup as possible and only add what’s necessary.

             
            
           
          
            
            
              @mkx - Thank you very much for your reply. I have no other firewall rules, only the posted ones. The KidControl rule appears dynamically, if i set the rule to “ON” or “Resume”, so if i would block the connection of my childrens device. Otherwise it is empty. But also the previous rules for blocking my IoT devices from connection are not working here  Maybe the order is confusing you, which was done by me as a try to get work the previous ones, but it didnt helped… Normally is the KidControl on first position.

The configuration was done exactly like that. Configured from scratch, only the needed options. The additional part are only those 3rules rules. I dont understand why should i take it offline immediately…

             
            
           
          
            
            
              When you decide to take advice and put a proper firewall on your device, I would be interested, but it seems you are here to prove your 3 rules iare good enough or something chow!

             
            
           
          
            
            
              @anav - Thank you for your reply !

Exactly as you described. Im here ONLY to kindly ask for assistance about 4 mentioned items in my firewall, and nothing else ! But this forum behaves as usually, as its well known  Im not interested with “support” like this anymore

             
            
           
          
            
            
              
So what is the posted config: complete config or “additional part”? If the former, then your router is pretty much wide open to the attacks from internet.

If the later, I’ll paraphrase a sentence, posted recently in another thread: how can you know what to take out if you didn’t know what to put in? In other words: whole setup in ROS is interleaved and if you don’t show us the whole of it, we can’t get the big picture and we can’t help you. In essence you’re wasting our time. Since you’re asking for hell and we’re the kind volunterrs trying to help you, you should at least respect our attempts.

             
            
           
          
            
            
              @mkx - Thank you for your feedback !

The posted config is the full config of my FW rules. KidControl is on last position just because of my attempts to make first three work… I dont have any FW rules, because my router is behind ISP ONT device with Dynamic IP changing couple of times per day, no ports on my router are opened-forwarded, and im charged monthly by ISP for “secure connection (whatever it  means)” besides costs for fibre internet. I am connected like this to the web on two different locations for 7 / 3years without issues. My only connection from outside is via Zerotier (just because SmartHome control). Just to clarify, i dont say that any additional security is not welcome from my side, but i never had configured anything additional. Was not needed, had never troubles, and honestly i cant do that correctly. I was able on my previous router to block mentioned DNS requests and internet connection with same config for couple of devices from my network with those rules (only to avoid connecting IoT devices to cloud services, since im using them locally and they cant “talk” simoutaneusly to two destinations) and KidControl to have a magic weapon against my children  Those are actually my goals only, to get back the “block” functionality.

If im wasting your time, please feel free to scroll further and please ignore my question. Im also pissed off, if somebody-something steals my time, of course. Since my lack of deeper knowledge here is a fact, i wasted lot of time to try searching for solution, documentations, made attemps to achieve my goal. Since i still failed, i made a last try here. The request was aimed for somebody, who can spend bit time and knowledges providing some assistance without feeling of waisting it.

Thank you !

             
            
           
          
            
            
              Nothing like a slovakian lovefest!

Regardless if nothing has happened,  it only takes one incident to ruin ones life,  what MKX is suggesting is common sense and prudent firewall rules, that are easy to implement and that ALSO include your blocking strategy by default.

THis is a stock setup that keeps the required default rules and keep access to the router, solely to the admin and blocks all traffic except that specifically permitted.

```
/ip firewall address-list
add address=admin-IP1  list=Authorized
add address=admin-IP2  list=Authorized
{ etc....... }
/ip firewall filter
{Input Chain}
(default rules to keep)
add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=accept chain=input comment="defconf: accept to local loopback:"  dst-address=127.0.0.1
(user rules)
add action=accept chain=input src-address-list=Admin 
add action=accept chain=input comment="Allow LAN DNS queries-UDP" \ 
dst-port=53 in-interface-list=LAN protocol=udp
add action=accept chain=input comment="Allow LAN DNS queries - TCP" \
dst-port=53 in-interface-list=LAN protocol=tcp
add action=drop chain=input comment="drop all else"
{forward chain}
(default rules to keep)
add action=accept chain=forward comment="defconf: accept in ipsec policy" ipsec-policy=in,ipsec
add action=accept chain=forward comment="defconf: accept out ipsec policy" ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related
add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
(user rules)
add action=accept chain=forward comment="allow internet traffic" in-interface-list=LAN out-interface-list=WAN
add action=accept chain=forward comment="allow port forwarding" connection-nat-state=dstnat  { disable or remove if not req'd }
add action=drop chain=forward comment="drop all else"
```

…

It should be fairly easy to add what you mean by kid control, but you stated it in config speak and not requirement speak.

1. identify user(s)/device(s),  groups of users/devices
2. identify what traffic they should be allowed to execute.
3. the rules above block all other traffic automatically.

For example the above rules

input chain →  allow the admin to config the router

input chain →  allow all LAN users access to router DNS services

forward chain →  allow all LAN users access to internet

forward chain → allow port forwarding, in case you have any port forwarding or dst-nat requirements  ( can be disable or removed if not useful ).

What I dont understand is the purpose of kid control.

Are you attempting to control  a number of IP/mac addresses as to

a.  where they can go in terms of internet?

b.  what time they can execute traffic?

etc…
