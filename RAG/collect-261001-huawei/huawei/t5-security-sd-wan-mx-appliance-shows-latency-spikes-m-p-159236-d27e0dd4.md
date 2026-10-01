---
id: collect-261001-huawei/huawei/t5-security-sd-wan-mx-appliance-shows-latency-spikes-m-p-159236-d27e0dd4
title: "t5-security-sd-wan-mx-appliance-shows-latency-spikes-m-p-159236-d27e0dd4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "license", "throughput"]
source: docs/RAG/collect-261001-huawei/t5-security-sd-wan-mx-appliance-shows-latency-spikes-m-p-159236-d27e0dd4.md
source_anchor: ""
source_lines: [1, 171]
sha256: 993d7b0f2c8ca8c706f5f4bb9893de0df3eed591178d3cf6f4e7095bf94e8843
---

# t5-security-sd-wan-mx-appliance-shows-latency-spikes-m-p-159236-d27e0dd4

MX appliance shows latency spikes
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2022 10:03 AM
Hello,
Have two ISP connections to an MX250 as WAN 1 and WAN 2 in a load balanced mode. Time to time latency spikes are noticed on both the ISPs at the same time. Checked with ISPs and both report stable, low latency internet connection. So, why would such latency spikes show up intermittently. Could a sudden burst of high uplink traffic lead to such latency spikes ? Uplink is mostly around 1.5 Gbps on a 4 Gbps internet connection.
Thanks,
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2022 10:37 AM
Are you running autovpn or just internet traffic.
The uplink page latency and loss are based on icmp packets. So most destinations drop icmp traffic or queue/buffer this traffic in case its busy.
You could add another destination to compare if you see the same loss/latency at the same time.
In case you run aotuvpn then use this statistics https://documentation.meraki.com/MX/Monitoring_and_Reporting/SD-WAN_Monitoring
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-16-2022 12:26 PM
So, on Uplink Latency page of MX 250 if destination is a google DNS server i.e 8.8.8.8 and we see latency spike on both carriers, is it due ICMP packets unable to hit DNS server?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2022 12:14 PM
Active-Active Auto VPN is On and load balancing enabled. VPN traffic is for control of some devices but we see hits on all our video streams traffic both downstream and upstream via internet.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-10-2022 10:56 PM
@TheAlchemist1 wrote:
Active-Active Auto VPN is On and load balancing enabled. VPN traffic is for control of some devices but we see hits on all our video streams traffic both downstream and upstream via internet. camzap
It's not always web browsing traffic, but visitors streaming media to the site are mostly UDP-based and with each delay spike of around five ms or more, a hit is achieved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2022 12:54 PM
Is this just web browsing traffic that is affected, and are you using content filtering? If so, if may be affected by your choice of "URL category list size". "Full List" defers the user's web browsing session till the MX completes looking up the content category.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2022 04:07 PM
@Philip D'Ath it is not web browsing traffic but media streaming traffic based of UDP and with every latency spike of even around 5 ms or more, a hit is taken.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2022 01:30 PM
@TheAlchemist1 the MX250 has the following recommended capacity limits:
VPN 1Gb/s
Advanced features 2Gb/s
Enterprise (basic) features 4Gb/s
If you are running 1.5Gb/s of traffic including VPN and advanced features then the box might simply be overloaded.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2022 01:50 PM
@CMR We are not using Advanced features so have 4 Gb/s of throughput available for Uplink traffic. Also, isn't the VPN traffic also added up in the Uplink traffic ? I thought all traffic (including VPN) shows up as uplink traffic which for is is around 1.5 Gb/s.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2022 02:21 PM
@TheAlchemist1 The VPN traffic uses the most resources, so an MX250 is designed for up to 1Gb/s of it.
As an approximate example, if you had 750Mb/s of VPN traffic you could have an additional 1Gb/s of basic traffic. If you had 900Mb/s of VPN traffic, you could have an additional 400Mb/s of basic traffic.
Im guessing you have the Enterprise license (like we do), as you mention you aren't using advanced features.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2022 03:05 PM
@CMR yes, we have enterprise license with 4 Gb/s throughput, so ok at that end. Also, most of our VPN traffic has been less than 10 mbps. So, for some reason, latency spikes are hard to understand.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-09-2022 06:38 AM
What firmware are you currently running? Since upgrading to 16.X, we've seen weekly spikes of latency (1500-3000ms) on our MX250 stack. Working with support we also saw CPU spiking in the same way. A reboot or swapping primary/secondary was almost always required to level out the MX
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-09-2022 01:47 PM
@finzwake1 we are running on version 16 as well. But, have not done a reboot in a while. CPU utilization has been been below 50% even during high upstream traffic.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-11-2022 11:12 AM
I'm seeing this on multiple networks, each showing both ISPs with similar spikes.
I've narrowed this down to the IDS feature:
Setting "Intrusion detection and prevention" to disabled returns latencies to normal values.
I've opened a support case on two different networks, but the support response did not really assist.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-11-2022 11:57 AM
@ewcs our setup does not have Advanced Security License so cannot see this Threat Protection>Intrusion detection and prevention but still get these latency spikes time to time on both ISPs at the same time. Definitely a Meraki issue rather than ISPs.
