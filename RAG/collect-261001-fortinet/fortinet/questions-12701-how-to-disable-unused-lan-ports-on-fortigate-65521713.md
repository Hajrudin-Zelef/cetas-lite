---
id: collect-261001-fortinet/fortinet/questions-12701-how-to-disable-unused-lan-ports-on-fortigate-65521713
title: "questions-12701-how-to-disable-unused-lan-ports-on-fortigate-65521713"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-12701-how-to-disable-unused-lan-ports-on-fortigate-65521713.md
source_anchor: ""
source_lines: [1, 46]
sha256: 4cc63b9a4e21acae25239afa6857b3fae0f579dcec9c7b58ac3dfa2482e89875
---

# questions-12701-how-to-disable-unused-lan-ports-on-fortigate-65521713

I know I can set the administrative status of WAN/DMZ ports to down to disable them, but I can only see how to set the entire internal interface up/down. Is there a way to set individual LAN ports down so someone cannot plug into them and access the network?
- 
        What type of fortinet device is it?David– David2014-10-27 22:18:58 +00:00Commented Oct 27, 2014 at 22:18
- 
        It is a Fortigate-60D.Paul Kearney - pk– Paul Kearney - pk2014-10-28 00:01:48 +00:00Commented Oct 28, 2014 at 0:01
3 Answers 3
It's look like that theinternal-switch-mode is set as switch¹ (by default). That means that all port on the internal interface are configured as they are only one:
  Switch mode combines FortiGate unit interfaces into one switch with one address. Interface mode gives each internal interface its own address.²
so, as I understand, if in system global configuration you set: internal-switch-mode interface, you shall configure each port independently, so you will able to reconfigure port 1 and 2 then disable the other as @David say.
NB Before switching modes, all configuration settings for the interfaces affected by the switch must be set to defaults.²
Ref:
This can be accomplished through the GUI by
- navigating to System -> Network -> interface
- select the desired interface
- select edit for that interface
- change "Administrative Access" to Down.
This can also be accomplished through the CLI with these commands:
config system interface
edit (desired interface)
set status down
(this works for devices running fortios)
- 
        That lets me turn down the entire internal interface (all 8 internal ports on the Fortigate-60D), but not individual ports on the internal interface.Paul Kearney - pk– Paul Kearney - pk2014-10-28 00:03:02 +00:00Commented Oct 28, 2014 at 0:03
- 
            
            
- 
        Furthering on my comment, I have a Fortigate device at a client site. I have two devices connected to internal ports 1 and 2. I want to disable ports 3-7 so that they are unusable.Paul Kearney - pk– Paul Kearney - pk2014-10-28 00:05:55 +00:00Commented Oct 28, 2014 at 0:05
- 
        See @gio900's answer for how to take the device off internal switch modeDavid– David2014-10-28 16:00:29 +00:00Commented Oct 28, 2014 at 16:00
You can not do this by GUI, it can be done only by CLI:
config system virtual-switch
edit internal
config port
  edit internal3
    set status down
  next
  edit internal4
     set status down
  next
end
end
end
Of course internal-switch-mode must be set to interface first.
Hope this helps.
Wlodek W.
