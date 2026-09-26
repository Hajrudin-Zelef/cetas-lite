---
id: collect-260926-mikrotik/mikrotik/site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it-3
title: "site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it.md
source_anchor: ""
source_lines: [375, 446]
sha256: 022b40d27f94a8d60b6f835e61291a60a20eca8ae6b09d2bc70f64241cc2c5f3
---

# site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it

```
/interface wireguard add listen-port=13231 mtu=1420 name=wireguard1
/ip address add address=10.255.255.2/30 interface=wireguard1 network=10.255.255.0
/ip route add dst-address=192.168.100.0/24 gateway=wireguard1
/interface wireguard peers add allowed-address=192.168.100.0/24 endpoint-address=blablabla.sn.mynetname.net endpoint-port=42346 interface=wireguard1 public-key="BLABLABLA="
```

Now it’s working as intended. I have removed the SSTP configuration and will keep Wireguard in its place!

Thanks everyone, and sorry for the inconveniences!

             
            
           
          
            
            
              Very common mistake is to mix up the keys, its confusing the first time and then becomes clearer.

Note when using a third party provider often they will give you the private key to use in your router vice having the router generate a private key like we do normally

WHY →  because the private key generates  a constant known public key, so the ISP provider an give you the private key to use which means they already know the public key to setup on their side and you dont have to send them anything.

             
            
           
          
            
            
              Mmm, that’s interesting. I’m wondering if Mikrotik could implement a simple “wizard-style” process like they do for DHCP servers?

Maybe you could run the “Wizard”, it’ll ask you a few questions, it’ll save the configuration on your current Mikrotik, and then give you a Summary of the commands that you would need to copy & paste into the other Mikrotik?

That would make it much easier for beginners, and would eliminate human errors like mine. Also, if the Wizard could cover the two most common scenarios: site-to-site and roadwarriors, it would help a LOT of people!

             
            
           
          
            
            
              Hi Filament,

i am planning to do such a config for my configs and would like to post complete setups here for both router A and router B. In my case, i will have something like 9 VLANs and will need to setup several Wireguard connections, one for each VLAN.

The thing is i would like to use your basic setup and try to setup wireguard interfaces 1-9 for each VLAN.

Could you pls post it here or send it to me?

tx

korg

             
            
           
          
            
            
              Start your own thread, your scenario bares little to no resemblance to the original threads situation.

State the traffic flow requirements and the design will fall out naturally, for example, there is probably no need to have a different wireguard per vlan approach.

Post a a diagram of your intentions and also post both configs to see what you have done so far.

             
            
           
          
            
            
              Really,  you think keeping track of extra keys is FUN?    You need a vacation LOL.
