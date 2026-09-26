---
id: collect-260926-mikrotik/mikrotik/a-simple-wan-lan-dmz-vlan-config-to-start-off-2
title: "a-simple-wan-lan-dmz-vlan-config-to-start-off"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/a-simple-wan-lan-dmz-vlan-config-to-start-off.md
source_anchor: ""
source_lines: [159, 257]
sha256: 672365d5db4193cdec785a8babcd3208df31d70fc9e4d7cbb35255691d15eda4
---

# a-simple-wan-lan-dmz-vlan-config-to-start-off

Oh, and I forgot, the NoOp VLAN interface gives you a traffic monitor that only includes LAN traffic, while the bridge shows all bridge traffic:


             
            
           
          
            
            
              
… if that traffic hits CPU-facing bridge port (either due to being CPU communicating with devices on same VLAN or if it’s broadcast traffic). Most of traffic between devices, connected to bridge, does not.

Another thing: if device supports bridge L2 offload to switch chip (seems your drvice doesn’t), most of happening will be hidden from bridge (the switch-like entity), including port and VLAN stats.

             
            
           
          
            
            
              
True, and in my setup, there is plenty of traffic that will never even see that particular router, but it is what it is. I still have a counter of what part of the total traffic of the bridge is non-DMZ/WAN-related. Traffic on a different leg of the network that just gets L2 switched is only ever recorded directly on the interface.

Again, the LAN VLAN interface is a NoOp, but there is also no reason against it. On the other side, you can just do:

```
/ip dhcp-server
add add-arp=yes address-pool=client-pool interface=lan lease-script=update-dhcp-dns lease-time=52w1d name=dhcp-lan-server
```

And that remains the same, no matter if there is a bridge called “lan” or a VLAN interface called “lan”. And without any context, it’s still clear what it is supposed to do. That’s also the basic idea behind putting interfaces into an interface-list and then setting tagged/untagged based on the list. It is far less convoluted than having to fight the brain teaser that is manually setting up tagged/untagged on

```
/interface bridge vlan
```

.

             
            
           
          
            
            
              Your setup probably works fine for you and I’m glad for it. But the problem is when it gets published as a general template for newcomers to grab and blindly apply. Because generally it has a few problems and those will bite a few of those users. And that’s the reason for it getting quite some negative feedback. Don’t take it personally, it’s just that everybody has different experiences and some will spot weaknesses and problems in the code where others won’t. So if you want to make the template as universal and safe as possible, listen to what people say and consider their suggestions to improve your code. As it is now it hardly surpasses some ignorant youtube guides (e.g. on how to “secure” MT router). And this is my subjective opinion, backed by a pinch of experience and knowledge on the matter.

             
            
           
          
            
            
              I’m not taking this negatively or even personally. I want to fix the real issues, but besides the ingress-filtering I haven’t seen any. I’ll be changing that issue accordingly. The rest remains “I wouldn’t do that” without any explanation. The NoOp-VLAN-interface has no issues - it is a matter of choice. Using VLAN ID 1 has no implications, at least none that anyone could point out, and it is again a convenience thing.

             
            
           
          
            
            
              As I already explained, VLAN ID 1 is used in implicit configuration which makes it non-obvious and even non-transparent. And that makes it insecure.

Having the NoOp VLAN interface again makes things a bit muddy, users who don’t understand how bridge and L2HW offload works might jump into wrong conclusions.

Having *ingress-filtering* disabled (default setting, you did it explicitly on a particular bridge port for some random reason you didn’t explain), together with default *frame-types=admit-all* setting, makes setup insecure, I explained it in one of my previous posts.

I’ll say it one more time before I leave this thread: if you consider other people’s objections as irrelevant based on the fact that you didn’t experience anything similar yet, then you’re making your code irrelevant (in best case) or you’re leading less knowkedgeable users into problems by providing substandard solution. And I’m not trying to make you change the code you’re running on your own device, I’m pretty sure you’ll be able to get yourself out of any trouble yourself. I’m trying to persuade you to make published code as bulletproof as possible (because in networking it’s not about ease of configuration, it’s about security).

             
            
           
          
            
            
              
Again no explanation besides, “I feel like that is insecure”?

It feels “muddy”? Talk about objective security analysis.

I assume export of the configuration and switching some stuff around made me eventually show the configuration that way. As you pointed out, the default value would be disabled, so wouldn’t even appear in an export. As I wrote, I fixed that in the template, especially since that is how I run it as well.

You never mentioned

```
frame-types=admit-all
```

, but if you’d care to elaborate?

I consider it irrelevant because you and others fail to point out any concrete consequence, not because of a lack of experience. Your arguments are “non-transparent” and “muddy”. This setup is already a great improvement over the VLAN tutorial articles posted by Mikrotik, since they all manually manage each interface’s untagged/tagged configuration, which is honestly a disgrace seeing ROS’ capabilities in consolidating configurations.

             
            
           
          
            
            
              
In practical testing, I found out that the “NoOp” VLAN interface is necessary for WiFi CAP to operate correctly.
