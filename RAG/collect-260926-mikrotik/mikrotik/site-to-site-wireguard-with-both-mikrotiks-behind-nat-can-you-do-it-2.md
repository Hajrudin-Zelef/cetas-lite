---
id: collect-260926-mikrotik/mikrotik/site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it-2
title: "site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/wireguard/site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it.md
source_anchor: ""
source_lines: [204, 374]
sha256: f7776f58c8794e027ead7569ebc2826c4e53fabcc66d25e490325017bec13c69
---

# site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it

add action=accept chain=forward comment=“defconf: accept established,related, untracked” connection-state=established,related,untracked

add action=drop chain=forward comment=“defconf: drop invalid” connection-state=invalid

add action=accept chain=forward comment=“allow internet traffic” in-interface-list=LAN out-interface-list=WAN

add acction=accept chain=forward comment=“allow remote wg inbound”  in-interface=wireguardNAME dst-address=192.168.100.0/24

add action=accept chain-forward comment=“allow local wg outboud”  out-interface=wireguardNAME src-address=192.168.100.0/24

add action=accept chain=forward comment=“allow port forwarding” connection-nat-state=dstnat

add action=drop chain=forward comment=“drop all else”

/ip firewall nat

add action=masquerade chain=srcnat comment=“defconf: masquerade” out-interface-list=WAN

             
            
           
          
            
            
              
The config I’ve posted is full and complete, these routers are currently working perfectly (using SSTP), but I would like to change to Wireguard. That’s why I’m asking for help. The examples in Mikrotik’s website do not work for me. I tried to use them, but I reverted back to my normal SSTP configuration, because I couldn’t get it to work.

I don’t use the firewall in any of the Mikrotik (firewall is already in place by the ISP router). The firewall is completely blank on both Mikrotiks. It’s not needed in my situation.

I also don’t use any NAT or sourcenat rule. Again, the NAT is handled by the ISP routers. Mikrotiks don’t need to do NAT at all. The config I posted works perfectly with SSTP!!

ether1 is a member of the bridge, like all the other ethernet ports. So, you’re right, I can delete the WAN group in the interface list. It’s useless because there’s no interface in the WAN group, it’s empty. I forgot to delete it, sorry.

So, again, my question is: can Wireguard do what SSTP is doing for me right now? Is Wireguard equally capable in my situation?

             
            
           
          
            
            
              You seem to know everything about how to setup the MT router. Clearly you dont need help.

Also Holvoe provided a link with information, just read that, you should be up and running in no  time.

As I stated I would use the input chain rule to ensure the wireguard handshake is successful.  Why you dont want confirming information is illogical.

No, the router needs to do Sourcnat.

Follow the logic tree,  the subnet on the mikrotik does not exist on the ISP router, it knows nothing about the subnet.

So when subnetA goes out to the internet by way of the WANIP of the mikortik it leaves as its subnet address and heads out the ISP router to the net.

When the return traffic arrives at the ISP router with dst-address of the subnet on the Mikrotik, the ISP router will drop it as it doesnt reccognize this as a valid local address.

The purpose of sourcenat is to change the srce address to the WANIp of the MT which is also its LANIP on the ISP router.  Thus return traffic is recognized and then sent to the MT router by the ISP router.  The MT router un sourcenats back the traffic to the originator.

             
            
           
          
            
            
              I wish I could knew everything about Mikrotik! I wouldn’t be asking for help in that case!

The ISP router in site A has a LAN interface, and a WAN interface.

The LAN interface of the ISP router has IP address 192.168.100.1 (it serves as the gateway for Mikrotik in site A). The WAN inteface of the ISP router is connected to a fiber ONT box, and receives a public dynamic address using PPPoE.

This configuration works perfectly well with SSTP. No firewall rules, no NAT rules, no sourcenat. Everything works perfectly.

My question is simple: Why is Wireguard unable to work in this situation? Or isn’t it? I’m not familiar with the Wireguard protocol and I’m struggling to understand what is needed to make it work.

It puzzles me that SSTP in my situation works so easily and effortlessly, why does Wireguard need so much complexity???

             
            
           
          
            
            
              Well using SSTP probably means your LAN users never go out the ISP WAN, but instead go out the SSTP tunnel?

OTherwise you wouldnt have any working internet.

I am trying to ensure your router is setup properly in general and also to ensure wireguard success.

You can drive your car to the story everyday but never get into the accident and ask WHY SHOULD I wear my seatbelt, I havent needed it thus far>…

Or perhaps you want to drive to a different state, and it has snow, so maybe one should consider better tires…

             
            
           
          
            
            
              All the computers in the LAN are running Windows 10. They have DHCP client active, and they receive a dynamic IP address from the ISP router. For example, in site A the computers receive private addresses in the range from 192.168.100.33 to 192.168.100.254.

The DHCP server in the ISP router in site A sets the gateway as 192.168.100.1

That’s why I’m forced to type this command in all the computers in site A:

```
route -P ADD 192.168.200.0 MASK 255.255.255.0 192.168.100.2
```

That command creates a static route inside Windows (the -P parameter makes it permanent, so it survives a reboot), and tells the computer to route the packets to site B using Mikrotik.

Basically, the Mikrotiks are only dealing with the SSTP tunnel and nothing more. The rest of the surfing of the internet is done using the ISP router.

             
            
           
          
            
            
              hello filament,

The config I’ve posted is full and complete, these routers are currently working perfectly (using SSTP), but I would like to change to Wireguard. That’s why I’m asking for help. The examples in Mikrotik’s website do not work for me. I tried to use them, but I reverted back to my normal SSTP configuration, because I couldn’t get it to work.


did you try to re-use your isp sstp port dnat for the wg port?

router a is the server.

router b is the client. cgnat-ed

so what is the error log says?

             
            
           
          
            
            
              In my view, Wireguard is conceptually not more complex then SSTP.

The only thing where you need to be careful is the correct entry of public keys (easy to goof up) and the definition of allowed address on the peer side.

Other then that, it’s very simple (a LOT easier then e.g. IPSEC/IKE2).

What’s so difficult about MT’s own Wireguard instructions ?

They are really dead simple.

The point anav wants to make (and me too) is spoon-feeding is not going to help you.

You will not learn from it.

So try those instructions again, post the config you have made and indicate what doesn’t work at that point.

             
            
           
          
            
            
              Arghhg, jesuschrist! I’m now seeing the mistake. I had created several peers in Mikrotik A during the testing and I was copying and pasting another peer’s public key into Mikrotik A (instead of the public key from the Wireguard interface of Mikrotik B). F*ck my life… 

The example from Mikrotik’s official web DOES in fact work. Here are the lines needed in Mikrotik site A:

```
/interface wireguard add listen-port=42346 mtu=1420 name=wireguard1
/ip address add address=10.255.255.1/30 interface=wireguard1 network=10.255.255.0
/ip route add dst-address=192.168.200.0/24 gateway=wireguard1
/interface wireguard peers add allowed-address=192.168.200.0/24 interface=wireguard1 public-key="BLABLABLAMAKESUREYOUCOPYTHERIGHTKEYDAMNIT="
```

And these are the lines needed in Mikrotik site B:

