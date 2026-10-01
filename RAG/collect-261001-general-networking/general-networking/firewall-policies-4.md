---
id: collect-261001-general-networking/general-networking/firewall-policies-4
title: "firewall-policies"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/firewall-policies.md
source_anchor: ""
source_lines: [343, 365]
sha256: c52fce9b2410c8111d374aba44d2d98d8709bdd7f1f3543299f6859d2dcab74f
---

# firewall-policies

If your FortiGate is operating in NAT mode, rather than enabling source NAT in individual NGFW policies you go to **Policy & Objects > Central SNAT** and add source NAT policies that apply to all matching traffic. In many cases you may only need one SNAT policy for each interface pair. For example, if you allow users on the internal network (connected to port1) to browse the Internet (connected to port2) you can add a port1 to port2 Central SNAT policy similar to the following:

### Application control in NGFW policy mode

You configure **Application Control** simply by adding individual applications to security policies. You can set the action to accept or deny to allow or block the applications.

Policy modes

### Web filtering in NGFW mode

You configure **Web Filter** by adding URL categories to security policies. You can set the action to accept or deny to allow or block the applications.


### Other NGFW policy mode options

You can also combine both application control and web filtering in the same NGFW policy mode policy. Also if the policy accepts applications or URL categories you can also apply Antivirus, DNS Filtering, and IPS profiles in NGFW mode policies as well a logging and policy learning mode.

TJ
Hi,

In security profiles > custom signatures, I only see options to create a new IPS signature or a new application signature. I am using NGFW policy-based mode.

I no longer see a way to ready the applications or categories, which existed before changing to policy-based mode.
