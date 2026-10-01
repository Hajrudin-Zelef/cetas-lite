---
id: collect-261001-huawei/huawei/r-networking-comments-199m1xv-any-ideas-to-troubleshoot-100g-link-between-bfe4e840
title: "r-networking-comments-199m1xv-any-ideas-to-troubleshoot-100g-link-between-bfe4e840"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-01-18"]
keywords: ["license"]
source: docs/RAG/collect-261001-huawei/r-networking-comments-199m1xv-any-ideas-to-troubleshoot-100g-link-between-bfe4e840.md
source_anchor: ""
source_lines: [1, 40]
sha256: d2db41e9a68dfa51cb69529eb5fcf354694ea01f627f603d6b912aabfcebb749
---

# r-networking-comments-199m1xv-any-ideas-to-troubleshoot-100g-link-between-bfe4e840

Any ideas to troubleshoot 100G link between Huawei switches? 
        
        
        
    
    
    SOLVED: S6730 switches can use only 1m QSPF28 DAC, that was the reason
I am trying to configure 100G link between two Huawei s6730 switches. Both devices have 100G interfaces and apropriate licence to use them. Switches are connected with 100G QSPF28 DAC.
However, this is interface information displayed by "display interface" command:
100GE1/0/2 current state : UP
Line protocol current state : UP
Description:
Switch Port, Link-type : trunk(configured),
PVID : 1, TPID : 8100(Hex), The Maximum Frame Length is 9216
Last physical up time : 2024-01-18 08:41:12
Last physical down time : 2024-01-18 08:40:35
Current system time: 2024-01-18 08:45:53
Port Mode: COMMON FIBER, Transceiver: 40GBASE_CR4_QSFP28
Speed : 40000, Loopback: NONE
Duplex: FULL, Negotiation: DISABLE
Mdi : -, Flow-control: DISABLE
FEC : NONE
Command output is the same for both switches. As you can see, interfaces work in 100G mode, but displayed speed is 40G. Also, transceiver type 40GBASE_CR4_QSFP28 looks strange to me, because my transceivers are 100G according to label.
There is no speed auto negotiation for 100G interfaces, and i can't configure their speed manually.
The question is how to get 100G speed on interfaces instead of 40G speed.
I would appreciate any ideas and advices
Thanks
Section des commentaires
Sounds a lot like this:
https://forum.huawei.com/enterprise/en/how-to-use-40ge-in-s6730-with-100ge-license-activated/thread/667231577618399232-667213852955258880
This is what I started with. I applied the license and used "assign port-type 100ge" command, than rebooted switches. It changed interfaces names from 40ge x/x/x to 100ge x/x/x, so I guess it is not a cause of my issue
Have you tried non dac? We had so many issues with DACs, we did a refresh and completely changed to dedicated transceivers and fiber.
I've not tried non DAC yet, because I dont't have 100G capable fiber transceivers right now. I will try to borrow a pair of transceivers tomorrow and test them
By the way, can you tell about your issues with DACs? We didn't use them a lot, and didn't have 100Gb links at all
They would randomly drop backbone connections. We were able to find some syslog fingerprints within the Cisco devices that pointed to the dacs just not transmitting the data like it should have been.
This was in a service provider environment so maybe the combination of different vendors and very high speeds made them less stable, but we have not had a problem after migrating away from them.
Which software is running on the switch? VRP5 or YunShan? Some Huawei switches only support DACs for stacking but not for „classic“ traffic
There is VRP5 software on my switches.
Could you please tell where to find more information about this limitation? I didn't see it in the device specification and my distributor didn't mention it.
I will check if I will find it again. But I know for S5732v2 with YunShan its the case, that DAC is only supported for stacking. As I wrote, I think the limitation does not affect all switches
