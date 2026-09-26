---
id: collect-260926-mikrotik/mikrotik/questions-1156640-how-to-wipe-and-completely-reinstall-all-firmware-on-mikrotik-cbad9897
title: "questions-1156640-how-to-wipe-and-completely-reinstall-all-firmware-on-mikrotik--cbad9897"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/lot-mikrotik/troubleshooting/questions-1156640-how-to-wipe-and-completely-reinstall-all-firmware-on-mikrotik--cbad9897.md
source_anchor: ""
source_lines: [1, 17]
sha256: c43908ba245b621982e42b397630d8ca21f9e42080f2e9f6e221fa9cff3f9eb9
---

# questions-1156640-how-to-wipe-and-completely-reinstall-all-firmware-on-mikrotik--cbad9897

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a brand new MikroTik router that behaves strangely out of the box (came with RouterOS version 7.11, upgraded to 7.14.1, default admin account disabled) - even when its configuration is completely reset, it sends out SYN packets to random IP addresses on the Internet (yes, the packets are going through the output chain, not input). Either I don't understand something (which is possible) or the router has some malware on it. So at this point I have two questions:
Given a prior history of multiple massive infections of MikroTik routers worldwide in 2018 and 2021, were these infections limited to router configurations (which can be easily inspected in WinBox and erased via configuration reset) or were they implemented as a binary code that was injected via a vulnerability and secretly executing in parallel with RouterOS and persisted after configuration reset?
How can I perform a complete wipe of everything on the router including all installed firmware so that I can then reinstall an image of RouterOS I can trust via NetInstall? I have read manufacturer's documentation on router reset (https://wiki.mikrotik.com/wiki/Manual:Reset) and found it incomplete and confusing, as it does not clearly explain what will be reset when using one or another reset option.
"even when its configuration is completely reset, it sends out SYN packets to random IP addresses on the Internet" does not automatically indicate compromise, if a) those are SYN/ACK reply packets and you overlooked the ACK part – you did not specify what the source and destination port numbers are, so it's not obvious from the post – or if b) those IP addresses belong to Mikrotik's "cloud" service which RouterOS uses to update its time and determine its external IP address.
I faintly remember that there have been both kinds, at different times.
You can already do that with Netinstall without performing a complete wipe. The Netinstall docs explicitly say that it "re-formats the device's disk" during the process.
I bought more than 100 mikrotik devices and see this problem just 2 times (and I have no reason for it)!
if you have reset your router (as you said), Connect directly via LAN and to the first port.
open winbox and use your device information that labeled under it or the box and write first ethernet mac address manually as address.
for username and password, if nothing printed on your device label, just use "admin" as username and leave the password blank.
Now you can connect to your router.
