---
id: collect-261001-cisco/cisco/questions-686-asa-5550-reboot-worthwhile-5be65210
title: "questions-686-asa-5550-reboot-worthwhile-5be65210"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "research"]
source: docs/RAG/collect-261001-cisco/questions-686-asa-5550-reboot-worthwhile-5be65210.md
source_anchor: ""
source_lines: [1, 25]
sha256: 6fa3faaa393522b99cb9a0073c0806a353ecb0def9ab6899a3782f8e34e2ceda
---

# questions-686-asa-5550-reboot-worthwhile-5be65210

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
13
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I've got an ASA 5550 that is performing loads and loads of operations (AnyConnect, NAT, ACL, RADIUS, etc, etc). It isn't particularly overloaded in terms of CPU & Memory, but it has an uptime of over 3.5 years.
Lately I have been attempting to deploy another IPSEC tunnel (via cryptomap) along with a NAT Exempt rule, but the ASA is exhibiting very strange behaviour. Sometimes when I add ACE's a mass of text pops up from nowhere in the description field. No matter what I do, my tests with the on-box PacketTracer tool do not yield the results I expect (for example - I see the packet hitting the Any/Any rule at the bottom of the ACL, even though there is a specifically configured ACE at the top of said ACL).
Anyway, the question is this: Has anyone ever actually solved anything by rebooting an ASA? It isn't my favourite option, but with the very strange behaviours I am seeing troubleshooting is becoming fruitless.
Longer answer: :-) There are bugs in every piece of software. The longer it runs, the more likely one is going to setup shop in your network. But more to the point, the longer it's gone without a reboot, the more little bits of "old" configuration and/or status will be left lingering. In IOS, no interface foo will emit a warning that it's not completely destroyed and configuration elements may reappear if you recreate the interface -- shouldn't happen in an ASA but in rare cases, it does. I've also seen phantom NAT entries after deleting them from the config. (that one actually is a bug)
When dealing with IPSec/crypto, I've found a whole lot of crazy can be cleared up by a reload. In one case (pix 6.3.5) it wouldn't re-establish a VPN tunnel until I did.
[edit] A word on reboots in general: I tend to reboot things just to make sure they will. All too often I've had various systems (routers, firewalls, servers) running for extended periods -- constantly being modified, and when something ends up restarting them (usually a power outage, but "oops, wrong machine" happens too) they rarely come back up exactly as they were before... someone forgot to make X start at boot, or some odd interaction of parts makes something not startup as expected. I admit, it's less of a concern for more static parts of one's infrastructure.
Generally, I don't recommend a reboot as a resolution to a problem unless you know you are dealing with a bug that introduces something like a memory leak or a cache overflow condition.
With an ASA running an image at least 3.5 years old, have you checked the Cisco bug toolkit? Odds are that any bugs in the platform will be documented and you can see if any look to apply.
I would also recommend opening a TAC case if you have support.
Reboots in my mind gloss over other problems and can make it very difficult (if not impossible) to find root cause. Ultimately without understanding the root cause, you don't know that you fixed anything and I find that very dangerous, especially on a "security" platform.
For instance, maybe you have a security vulnerability in the code that is being exploited by an outside source. While the reboot may cut off their connection and alleviate the symptoms, it does nothing to address the problem.
As mentioned, risk management and vulnerability management should be your concerns. I'd say there are at least 10-20 known vulnerabilities for your ASA software version, assuming you had the latest firmware installed at the time represented by uptime.
Cisco IOS Software Checker. I don't know if there's something similar for the ASA, but perhaps someone could chime in?
Router Configuration Auditing: RedSeal may include version checks (it's been several years since I've used worked with it), as well as plenty of other security tools for networks
Vulnerability Management: Nessus has community and commercial versions, and there is plenty of other software like this out there
I have recently encountered similar problems from an ASA running 8.2(2)16 with ~2.5 years uptime, whereby object-groups specified in crypto map ACLs were not being matched. Adding an ACL statement that the object-group already encompassed caused interesting traffic to be matched. Very frustrating.
A colleague advised they had seen this behaviour previously and that a reload resolved it in that instance.
When you say a load of 'random' text is appearing when adding ACE's, are you manually typing these ACE's in or are you pasting them from some other source (like notepad).
I have seen issues before where if you are pasting a lot of lines into a device it can get overloaded and some corruption occurs, pasting less lines usually fixes it or using a function on your terminal program to 'paste slow' to allow for a small time gap between each line.
