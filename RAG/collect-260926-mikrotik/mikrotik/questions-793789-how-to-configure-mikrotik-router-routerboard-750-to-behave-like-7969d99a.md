---
id: collect-260926-mikrotik/mikrotik/questions-793789-how-to-configure-mikrotik-router-routerboard-750-to-behave-like-7969d99a
title: "questions-793789-how-to-configure-mikrotik-router-routerboard-750-to-behave-like-7969d99a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-793789-how-to-configure-mikrotik-router-routerboard-750-to-behave-like-7969d99a.md
source_anchor: ""
source_lines: [1, 10]
sha256: 2959928d499264bb3ce0c3cb3cc6b2bbbac0311145b46436e591aef2d8405eef
---

# questions-793789-how-to-configure-mikrotik-router-routerboard-750-to-behave-like-7969d99a

In a mikrotik, it makes the distinction between a bridge and a switch, even though functionally they do the same thing.
A switch is a hardware accelerated function that uses an onboard switching chip. It will give you better performance than a bridge, because the CPU for the most part is not involved in processing the packets.
A bridge is a Linux Kernel software bridge. The processor is involved in each packet that goes through. This means it can do filtering, packet capture of traffic between ports, etc. But, you will have lower performance because it is not hardware accelerated.
So, unless you have filtering/packet capture requirements, you want to use the switching function instead of bridge. This will get you performance closer to full wire speed.
So, how does Mikrotik handle the two options in software? Well, as you've seen, with a software bridge, you added each port to the bridge, and then the CPU emulates a switch, creating a new virtual interface with the name of the bridge. This virtual interface is as if it is connected to the emulated switch, and is where you assign layer3 stuff (IP addresses, IP filter rules, etc).
For the hardware switch chip(s), ports are grouped and connected to a switch chip. Usually 4 or 5 ports go to one switch chip, and usually it is the ports physically close to each other. The way the software handles this is that you go to the interfaces you want as the switch, and pick one port as a master port, and set all the others to use that master port. Then, you configure any layer 3 config on the master port and leave the slave ports alone. The switching hardware switches between the ports, and the CPU only deals with packets to/from the address on the master port, offloading everything else to the hardware switch chip.
So, with the current config, traffic from ether1 to any of the other ports on the switch may be slower than traffic between ether 2-5. 
Ideally, you want to have ether1 have ether2 as a master port, and have any management addresses for the mikrotik on ether2 (the master port). 
So, I would start by going into safe mode. Then add another management IP address on the current master port (ether2 in your case). Then remove ether1 and then ether2 from the bridge. At this point, you will be probably disconnected, because you just removed the old management IP. Log back in on the new management IP. Delete the old bridge. Add the old management IP to ether2. Now you should be able to login on the original management IP again. Go to interfaces and change ether1's master port to ether2. Bring up a console, and put in press Ctrl - x  to try to re-enter safe mode. When it says Hijacking Safe Mode from someone - unroll/release/don't take it [u/r/d]: tell it 'r' to re-take over the old safe mode session (so it doesn't undo your changes), and then do Ctrl-X to go out of safe mode and make the changes permanent.
(Safe mode means that if you are disconnected, any changes since you entered safemode will be undone, if session ends or is broken and you wait a specified timeout (usually 9 minutes) without re-picking up the session.)
