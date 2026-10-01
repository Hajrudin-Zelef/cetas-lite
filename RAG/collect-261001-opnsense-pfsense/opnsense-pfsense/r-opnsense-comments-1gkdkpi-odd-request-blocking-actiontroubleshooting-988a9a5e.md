---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/r-opnsense-comments-1gkdkpi-odd-request-blocking-actiontroubleshooting-988a9a5e
title: "r-opnsense-comments-1gkdkpi-odd-request-blocking-actiontroubleshooting-988a9a5e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/r-opnsense-comments-1gkdkpi-odd-request-blocking-actiontroubleshooting-988a9a5e.md
source_anchor: ""
source_lines: [1, 31]
sha256: c4e89099de5f574aae8e34dea8147012e73fff78062c5a61bd9d2d169c7b6388
---

# 
       Odd Request Blocking Action/Troubleshooting 

    
    This one has me almost beating my head into the wall, and I'm hoping someone here as a better idea.

When playing spotify on my phone and using the DJ function the app locks up. I can play music fine but when the DJ starts a new section manual intervention is required.

I have crowdsec installed, have a list of blocked malicious IPs, and was using DNS blocking at point point.

Crowdsec and malicious IPs are on floating rules on all of my VLANS. So as I understand if it was being blocked on those it would apply to everyting on any of my VLANS.

My phone operates correctly on celluler so clearly its a local network issue. Spotify on another vlan works without issue.

This would lead me it's an issue with DNS on my lan (VLAN 1). (It's always DNS...)

Here is my confusion. In an attempt to trouble shoot the issue I disabled all the Unbound DSN blocklists and disabled it and my issue remains. I have also diabled my malicious IPs floating rules and restarted the firewall but the issue persists.

Does anyone have any insight into what is happening or have any ideas? This is very annoying but fortunately does not majorly break anything.

Edit: As a side note if I change the DNS on the playback device to PiHole (setup just for troubleshooting this issue) it work without problem. (Again points to DNS) But I don't know whats holding me out with unbound on opnsense.

Is the DNS server on your firewall and PiHole using the same external DNS?

Do you have any custom DNS block/allow rules on your main lan could be causing an issue?

opnsense is using unbound. Pihole was set to opendns.

I don’t have any DNS firewall rules for lan. Its only default allow all.

You could try using an external DNS like NextDns and enable logging so you could see what name requests are being made. With a little extra work you can get it set so you can see the requests by device. The extra logging could be helpful with your troubleshooting.
