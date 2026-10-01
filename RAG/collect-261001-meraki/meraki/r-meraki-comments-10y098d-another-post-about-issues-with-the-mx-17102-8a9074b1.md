---
id: collect-261001-meraki/meraki/r-meraki-comments-10y098d-another-post-about-issues-with-the-mx-17102-8a9074b1
title: "r-meraki-comments-10y098d-another-post-about-issues-with-the-mx-17102-8a9074b1"
domain: meraki
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["incident", "license", "throughput"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-10y098d-another-post-about-issues-with-the-mx-17102-8a9074b1.md
source_anchor: ""
source_lines: [1, 34]
sha256: df7e04a3674966ebd57955f4578fd6a1d63f0098ea7863a2bd25a542fea7aebd
---

# r-meraki-comments-10y098d-another-post-about-issues-with-the-mx-17102-8a9074b1

Another post about issues with the MX 17.10.2 firmware... All AutoVPN peers dropped VPN connectivity 
        
        
        
    
    
    On Monday, all was well. Life was good. We had 60 remote peers (MX64s, MX67s & MX84s) all working fine, all AutoVPN'd to our Head Office (2x MX250 in a HA pair).
Around came 2PM & our support team became inundated with calls. Each call was logged & during the process, each call was escalated to my team via Slack. Within less than 5 minutes, we had around 10 calls from different sites stating they couldn't access internal network resources.
All sites were showing green on Meraki, getting a good throughput - however the VPN status page showed all of the remote peers with a red dot, suggesting disconnection.
These remote sites were able to contact the internet and internet based resources (Google/BBC/Office 365 etc) however anything hosted internally was unable to be contacted.
We disconnected the site-to-site VPN connection and re-enabled it, hoping it would re-establish the AutoVPN tunnels to no avail. We rebooted the MX250 and that brought everything back online. We contacted Meraki support to make them aware however as the router had been rebooted, certain logs and files weren't able to be assessed and viewed.
At 4:30PM, the same issue occurred. Meraki support were immediately called, who weren't able to see anything. The support engineer insinuated the MX could have been faulty, suggesting that we fail over onto the secondary MX in our HA setup - and this was related to a hardware fault. I asked him to summarise his suspicions in an email/comment, which never happened.
Nonetheless, we failed over onto the secondary MX and all seems to be well. This further suggested to me that it was a hardware fault on the primary MX250.
I've been doing a lot of to-ing and fro-ing regarding this issue with Meraki support. I yesterday was advised that in actual fact, they believe the issue is related to the firmware version. We upgraded two weeks ago to 17.10.2 and they've advised us to downgrade from MX 17.10.2 back to MX 16.16.9. I called again to schedule the downgrade in & the engineer further advised that since the general public release of MX 17, they've had lots of calls about lots of different issues, this being one of them.
Saying that, I find it odd that since failing over to the secondary MX250 in the HA pair, the issue has since disappeared and hasn't reoccurred - and this secondary MX250 is also on 17.10.2.
Surely if they've had lots of calls about issues with MX 17.10.2, they'd have either placed the release on hold or at the very least, put a disclaimer on the changelog stating that there is a chance this can cause issues?
Section des commentaires
I have a single data point to share about 17.10.2
A customer with a single MX105 had the 16 to 17 update scheduled overnight a couple weeks ago. The MX was not passing passing traffic in the morning and showed offline in the dashboard. A power cycle restored it and it seemed like maybe just a one time glitch. I opened a case with Meraki and they told me even though the dashboard showed 17 the MX was really running 16 and we should try the update again.
I was game, but of course when I tried I could not upgrade because the dashboard showed I was already on 17. So, I called Meraki back and they scheduled on the back end a downgrade to 16. This ran, but from this point on the system never came back online after an upgrade, downgrade or even soft restart from the dashboard. Any of those events would result in the system not passing traffic (local status via maintenance port access fine though) and require power cycle.
I requested RMA and so far the new MX is fine. It updated itself and rebooted and I confirmed soft reboot though dashboard worked as expected.
I suspect, but did not wish to spend my time troubleshooting for Cisco that something corrupted in firmware during the first update attempt to cause this. Possibly a factory reset of the original MX may have helped, but the end customer was not willing nor was I to go through that.
As opposed to your description, this customer's configuration was VERY simple. Single MX, Enterprise license, one big flat LAN, single WAN uplink. Single LAN link to switches. No traffic shaping or firewall rules, a single site to site VPN tunnel is about the only thing not default on the box.
All that to say, I agree there may be a serious issue with 17.10.2 or at least the process of upgrade to it.
This scares me. Genuinely scares me. This will cause a huge outage if this is the case...
No doubt. I really hope it was an isolated incident, but I am starting to wonder.
I had this happen on previous versions as well. The company I had worked for had about 40 sites all were HA pairs. Somewhere about 5-6 of them had the issue. Two I couldn't bring back to life needed a RMA. Since then I don't update any MX right away. Start at the small less important sites and slowly work your way to the hubs and distros. Trust no meraki firmware.
Do the MX250's have a public IP on them, or do they sit behind something doing NAT?
They share the same uplink & present with public IP's.
So they are sitting behind something doing NAT?
Man, that sucks... For what it's worth, I have been running 17.10.2 for over two months with similar hardware models (MX64, MX67, MX68, and two MX250s at HQ). We haven't seen any issues like this.
Same
I proactively cancelled our main hubs from upgrading to 17.10.2 next week. Staying on 16.16.9 for the foreseeable future
Just had to downgrade our mx250 down to 16.16.9 due to clientvpn issues. Straight up would not see connected clients or apply policy but would log in the event log. Meraki support advised a "wide range of unexpected issues" with 17.10.2 so far.
