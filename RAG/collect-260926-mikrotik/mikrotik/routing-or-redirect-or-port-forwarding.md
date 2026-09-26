---
id: collect-260926-mikrotik/mikrotik/routing-or-redirect-or-port-forwarding
title: "routing-or-redirect-or-port-forwarding"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/routing-or-redirect-or-port-forwarding.md
source_anchor: ""
source_lines: [1, 162]
sha256: ca4da78d7b488b8b53415517a8a6d2d75e9ae1e0c8d5c185468aee23cadd1db1
---

# routing-or-redirect-or-port-forwarding

Hi,

I need to redirect a incoming traffic through X port to other ip address in the same lan segment and same port…but maintaining the client ip… It’s possible this and how ?

Best regards

             
            
           
          
            
            
              I’m not sure I understand your questing completely, you may want to describe the situation more thoroughly (which port clients connect to, which port router connects to, how are those ports configured in router, how exactly clients try to connect to server, etc.).

Anyway, if I understand your question, then It is possible - just perform DST NAT like this:

```
/ip firewall nat
add chain=dstnat action=dst-nat dst-port=<port number> to-addresses=<some LAN IP address>
```

However, this has a few problems:

- LAN clients will have to make connection to one of router’s IP addresses (not a huge problem)
- probably the deal breaker with LAN-to-LAN connections: server will see incoming IP address from same IP subnet and will respond to client directly. Which actually causes 2 problems:

1. clients will initiate connections towards router’s IP address, but will receive response from server. Which clients will not recognize as valid connection setup and will discard response packets.
2. router’s firewall will not see response packets, hence connection tracking state will not be updated appropriately. Even if previous point would not prevent the connection from being established, firewall in router would break it at this point.

If there would be no NAT in play (clients would try to connect to server’s IP address directly), then the only problem is that traffic should not be routed between clients and server if they are all members of same subnet. Which means that ports, connecting to clients and connecting to server, should be bridged/switched, not routed.

             
            
           
          
            
            
              Yes. You use an *action=dst-nat* rule in *chain=dstnat* of */ip firewall nat*, matching on the destination port and some of (*in-interface*, *dst-address(-list)*, *src-address(-list)*), and setting the new address as *to-addresses*. The routing will use the new destination address. If the default setup of */ip firewall filter* is used, you don’t need to add any filter rules to permit the traffic. The source address remains unchanged unless you use some *action=src-nat* rules.

             
            
           
          
            
            
              Hi,

Thanks for the quickly answer… So far i have the same like this and works… but in the final server (teamspeak) instead of the client ip i have the MK ip… that’s why i want to know if is this possible preserve the client ip address…

The thing is that i need to force to all my clients to connect to my teamspeak lan server, and if they disconnect block the internet access… i have this code but i can’t figure how can force them..

My mikrotik works like GW..

             
            
           
          
            
            
              Maybe if i change the teamspeak server ip address out of my lan segment to force the MK route !

             
            
           
          
            
            
              Do I get you right that the clients are in the same LAN subnet like the Teamspeak server?

             
            
           
          
            
            
              Hi,

ISP global network 10.96.0.0/12 - 10.83.0.0/16

My localnetwork is 10.89.100.X/24

My pc Ip 10.89.100.15

My TS server usually is 10.89.100.55…

ISP have a captive portal in 192.168.100.1:3333

What i need is force my users to logging in the local TS before or after to login or reach the captive portal…

Pd: After login in the captive portal i can connect to other TS 10.96.55.3 and there remains my IP address 10.89.100.15…

That’s what i need, i want simulate my entire network using gns3 and mikrotik !!

Help ?

Best regards

             
            
           
          
            
            
              To move the teamspeak server to its own subnet (also in the real deployment, not just in GNS3) is the only way how to make the TS server see the clients’ actual IP addresses although they will be connecting to the public IP of the Mikrotik which will port-forward these requests to the actual address of the TS server.

The thing is that if the clients and the TS server are in different subnets, all traffic beween them passes through the router “naturally”. If they are in the same subnet, you need to force the router into the response (server->client path) so that it could “un-dst-nat” the responses, and the only way to ensure that is to src-nat clients’ requests on their way to the server. So you can either src-nat the requests to router’s own IP in that subnet, which means that the server won’t see the actual addresses of the clients, or you can src-netmap them with a prefix of another subnet (@Sob’s idea), which means that the server will see a distinctive individual address for each client but it won’t be the real one, it will just inherit last N bits from the real address and the prefix will be different.

             
            
           
          
            
            
              Can you give me an example about the src-netmsp option, pls…

Thanks for the entire explanation… I really appreciate..

Best regards

             
            
           
          
            
            
              */ip firewall nat
add chain=srcnat out-interface=bridge src-address=192.168.192.0/18 action=netmap to-addresses=10.123.64.0/18*

Here, *to-addresses* is not a range or a list but a prefix. The 18 most significant bits of the original source address will be substituted, the 14 least significant ones will be copied to the new source address.

Other than that, the rule behaves exactly the same way like an ordinary *action=src-nat* one.

             
            
           
          
            
            
              
Ok, question…

src-address=10.89.100.0/18        Right ¿?

to-address=¿?    Should i leave that you use for example…

Do i move the TS to another subnet then ? Like 172.16.1.X ¿? and then use this rule ¿?

Best regards

             
            
           
          
            
            
              
The mask size (prefix length) should be the same like in the LAN subnet you want to translate. 10.89.100.0/18 is a nonsense, the minimum mask length for 10.89.100.0 is 22. And normally, the mask length should be the same in both the *src-address* and *to-addresses*.




If you can move the TS to another subnet, you don’t need this rule at all, and the TS will see the actual addresses of the clients rather than translated ones.
