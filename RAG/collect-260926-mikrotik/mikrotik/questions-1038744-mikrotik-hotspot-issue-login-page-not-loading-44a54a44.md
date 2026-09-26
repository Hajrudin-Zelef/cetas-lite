---
id: collect-260926-mikrotik/mikrotik/questions-1038744-mikrotik-hotspot-issue-login-page-not-loading-44a54a44
title: "questions-1038744-mikrotik-hotspot-issue-login-page-not-loading-44a54a44"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-1038744-mikrotik-hotspot-issue-login-page-not-loading-44a54a44.md
source_anchor: ""
source_lines: [1, 7]
sha256: a86b2da83c44b09c4aa7673b616904fa9f06c3ecf74602d2d1f8bdaa64f79ee9
---

# questions-1038744-mikrotik-hotspot-issue-login-page-not-loading-44a54a44

I've configured a hotspot on a Mikrotik CCR1036-12G (OS 6.47.4), but after the device connects to the wifi and asks to authenticate in the network, the login page remains loading for a while and returns an error page with net:ERR_CONNECTION_TIMED_OUT. The hotspot worked perfectly already, but suddently stopped.
1 Answer 1
Solved! The problem wasn't on Mikrotik configuration, that was OK. The problem was on the wifi configuration, in Unifi solution:
i) In the network configuration, the used network was marked as "Guest". I changed to "corporate".
ii) In the wireless configuration, there was enabled the "Gest Policy" option. I disabled it.
Both options (or maybe just one - the guest policy, I think) was resulting in the described problem because I didn't configure the guest portal in the Unifi solution. My fault... :´/
Well, we learn from our mistakes too...
