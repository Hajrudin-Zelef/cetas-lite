---
id: collect-261001-meraki/meraki/r-meraki-comments-hhk461-issues-with-sitetosite-vpn-cannot-reach-clients-dd992307
title: "r-meraki-comments-hhk461-issues-with-sitetosite-vpn-cannot-reach-clients-dd992307"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-hhk461-issues-with-sitetosite-vpn-cannot-reach-clients-dd992307.md
source_anchor: ""
source_lines: [1, 25]
sha256: 361e868baec3472799797cae006aef9577b8a607e358b1bc97a86354bc45c607
---

# r-meraki-comments-hhk461-issues-with-sitetosite-vpn-cannot-reach-clients-dd992307

Issues with Site-to-Site VPN. Cannot reach clients on other network. 
        
        
        
    
    
    Hi!
I'm a sysadmin new to Meraki and I am having trouble with my newly setup site-to-site VPN. The VPN is configured in Hub (Mesh) Between 2 MX64 appliances. From the VPN status page everything shown there indicates that the connection is up and running, I can even, using the ping tool on both appliances, ping devices on both networks across the VPN. However when I try to communicate with device, such as the server over the VPN connection I get nothing. I keep on getting General Failure errors back on my pings from the workstations as well.
I'm sure this is due to a couple of simple settings I am missing because of my inexperience with Meraki. Any help on this would be greatly appreciated. Thanks!
Section des commentaires
Figured it out! Well almost, plugged in my laptop and it worked! Looks like the firewall on the workstation I was using was blocking traffic. Can't believe I didn't think to try that before. Still having issues resolving DNS over VPN, though I suspect it has more to do with my DNS server then the VPN connection.
Another suggestion too.. Since you have enterprise or advanced licensing on the devices you can call into Meraki support to get a TAC engineer to assist you. Thought I'd mention that since you are new to Cisco Meraki.
Check "Security & SDWAN > Site to Site VPN" and make sure VPN participation is ON.
This
It gets me every time I set up a new MX, I forget about this.
That was the resource I was going to exploit if everything else failed. Never had to call in before but I've heard a lot of good things about Meraki support.
Support is hit or miss. All depends if you get an L1 or L2 resource.
Public IP on wan interface?
Yes the VPN is connected to the WAN interface on both devices.
Just wanted to make sure it's not a lte connection without a public IP.
Check your firmware, update it first before trying to troubleshoot. We had issues with auto vpn this weekend that firmware upgrade to latest beta fixed.
What's your current firmware version?
I seem to recall having to put VLAN forwarders in each VLAN at both sites.... but it’s been a while.
Got no VLANs so far on either network so doesn't really apply to me case right now. But I will keep this in mind for when I do. Thanks!
Another thing I run into sometimes even when the VPN shows up, but you cannot ping anything from either side, I turn the VPN off completely at one site and turn it back on. Something happens on Meraki devices sometimes that are strange. Probably not a big deal for most but when you manage 80 MXs you see things every now and then.
