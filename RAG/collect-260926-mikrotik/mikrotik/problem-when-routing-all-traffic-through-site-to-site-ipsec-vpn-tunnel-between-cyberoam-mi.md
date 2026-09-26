---
id: collect-260926-mikrotik/mikrotik/problem-when-routing-all-traffic-through-site-to-site-ipsec-vpn-tunnel-between-cyberoam-mi
title: "problem-when-routing-all-traffic-through-site-to-site-ipsec-vpn-tunnel-between-cyberoam-mikrotik"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/ipsec/problem-when-routing-all-traffic-through-site-to-site-ipsec-vpn-tunnel-between-cyberoam-mikrotik.md
source_anchor: ""
source_lines: [1, 127]
sha256: bd913e2e4b76433c9e37da3ff4d4eb9bed05aee4e032072480de592fa2f159a4
---

# problem-when-routing-all-traffic-through-site-to-site-ipsec-vpn-tunnel-between-cyberoam-mikrotik

Hello everyone,

I’m trying to create site-to-site IPSEC VPN tunnel between two sites, and pass LAN and Internet traffic from Site2 to Site1. The Configuration is working; the tunnel is being created, and all traffic passed, but several issues are occurring with the Mikrotik when the IPSEC policy has the dst-adderss=0.0.0.0/0 .

-Configuration:

- Site1: Cyberoam v10.6 with LAN 192.168.0.0/24
 
 
- Site2: Mikrotik v6.41.4 with LAN 192.168.1.0/24
 
 
- NAT is not being used, because there is route between the two sites
 
 
- IPSEC Policy on Mikrotik: src-address=192.168.1.0/24 dst-address=0.0.0.0/0
 
 
- IPSEC Policy on Cyberoam: src-address=0.0.0.0/0 dst-address=192.168.1.0/24

-Issues:

- Some TCP traffic is getting “lost”, and others getting TCP RST flag after some time (weird part is TCP handshake is always completing without any issues)
 
 
- Can’t access the WebFig from the Mikrotik LAN side
 
 
- Can’t access web page located in the Mikrotik site (Site2) from Site1

Note: The tunnel is working well without any issues when the IPSEC Policy in Mikrotik is following: src-address=192.168.1.0/24 dst-address=192.168.0.0/24. Since we want to route all traffic, this configuration does not help us in any way.

Any kind of help,insight, alternative configuration for site-to-site VPN with passing internet traffic will be appreciated.

             
            
           
          
            
            
              
The point is you actually **do not** want to route all traffic to the tunnel - namely, you do not want the wide IPsec policy (****

```
0.0.0.0/0 -> 0.0.0.0/0
```

) to “steal” your local traffic, which it currently does. IPsec policies match on packets which have already been routed and are just about to be sent out via a physical interface, and this is true for any outgoing traffic, including the one between your Mikrotik and the devices on its LAN.

So you have to add another policy as an exception from the wide one:

```
/ip ipsec policy add action=none disabled=no dst-address=your.lan.subnet dst-port=any protocol=all src-address=your.lan.subnet src-port=any
```

and place it before the wide one.

The IPsec policies work similar to firewall rules in terms that they are matched top to bottom. Unlike in case of routes, it is impossible to sort them automatically as there is more than one prefix size to compare. So e.g. if one policy would have ****

```
src-address=x.x.x.x/24
```

and

```
dst-address=y.y.y.y/16
```

, and another policy would have

```
src-address=x.x.x.x/16
```

and

```
dst-address=y.y.y.y/24
```

, no “proper” order could be found automatically.

             
            
           
          
            
            
              
The point is you actually do not want to route all traffic to the tunnel - namely, you do not want the wide IPsec policy (0.0.0.0/0 → 0.0.0.0/0) to “steal” your local traffic, which it currently does.


Are you sure? Only MT is working curiously with this setting (ipsec all traffic). From the routing point view: local networks are the best metric, so this traffic should has predomination than another. After routing traffic should be encrypted. Packets from local network to local gateway should not never routed outside of local interface.

             
            
           
          
            
            
              The point is that in RouterOS, IPsec policies do their job after all routing and even firewalling has been done, just before the packet is about to be sent out via the physical interface. See this diagram and especially the Routing Diagram details (points IJKL) below.

So even though the “normal” routing says “ok, this is a packet for a host on connected subnet, let’s send it out the corresponding Ethernet interface rather than via some gateway”, the IPsec policy still grabs the packet if its traffic selector matches the packet fields, and lets it be encrypted and sent out the associated SA.

The easiest way to check this is to give it a try. As soon as you activate a policy whose *src-address* matches the LAN address of the Mikrotik and whose *dst-address* matches the same LAN subnet, hosts in that subnet stop being reachable for the Mikrotik. If you deactivate that policy, they become reachable again. Beware - if you access the Mikrotik from such device, by activating the policy you lose access. So activate the IPsec peer enabling that policy in safe mode in that case.

             
                
            
           
          
            
            
              
Beware - if you access the Mikrotik from such device, by activating the policy you lose access. So activate the IPsec peer enabling that policy in safe mode in that case.


I know it, but still not understand this “mis-conception” from reasons:

1. packet into locally network should not leave this network; moreover in the same network communication relies on the L2, not L3
2. access to management of router is not forward, only input traffic
3. I tested it on the R11e - it can not manage of MKT even if SIM card will be removed (lte intereface down), but does IPSEC still “working”?

I think that is should be fixed. Its not normal behaviour.

There is not problem with traffic from local-connected hosts to gateway, “reply-traffic” from gateway is directed to the ipsec, not to the local connected network.
