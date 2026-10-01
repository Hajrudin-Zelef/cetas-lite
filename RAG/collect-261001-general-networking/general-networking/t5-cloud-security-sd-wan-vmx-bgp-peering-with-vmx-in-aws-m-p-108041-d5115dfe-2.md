---
id: collect-261001-general-networking/general-networking/t5-cloud-security-sd-wan-vmx-bgp-peering-with-vmx-in-aws-m-p-108041-d5115dfe-2
title: "t5-cloud-security-sd-wan-vmx-bgp-peering-with-vmx-in-aws-m-p-108041-d5115dfe"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Lambda"]
dates: []
keywords: ["aws", "cost"]
source: docs/RAG/collect-261001-general-networking/t5-cloud-security-sd-wan-vmx-bgp-peering-with-vmx-in-aws-m-p-108041-d5115dfe.md
source_anchor: ""
source_lines: [22, 161]
sha256: 632b2c400e1e383d9845d4730a0d51e535e90fc238344cecdb9842d59cc61d32
---

# t5-cloud-security-sd-wan-vmx-bgp-peering-with-vmx-in-aws-m-p-108041-d5115dfe

			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-16-2021 08:44 AM
After weighing options like spinning up a csr1000v and other things, we decided to implement an ASAv firewall in AWS. This could be done with PaloAlto and other devices. This gets us around the need to do BGP peering as we have done in other traditional data centers. This should satisfy the routing issue by having a single default route from the vMX to the firewall and the firewall with a default route to the Internet and routes to our internal AWS resources. We also have an ASAv that is used for RA-VPN in AWS so putting both the vMX and the RA-VPN ASAv into a DMZ made a bit more sense for us. This provides an additional layer of security outside of what we may have configured with the AWS Security Groups.
Below is a basic diagram for anyone else looking for a deployment solution for vMX in AWS.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 07:36 AM
Meraki Support hasn't been much help either. All they have done is directed me to the BGP configuration guide.
https://documentation.meraki.com/MX/Networks_and_Routing/BGP
The AWS deployment guide also does not contain information about routing other than adding the AutoVPN subnets to the VPC route table. That appears to be more of a static routing configuration.
The issue with that is we are in a transition between physical data centers and cloud. I need the subnets that are used by our office or Z3s to be advertised from the MX or vMX that they are connected to. Adding a blanket static route to a VPC would seem to break dynamic routing.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 07:44 AM
Hi @PatrickBixler22547 , are you able to reach out to your Meraki AM to engage with an SE for this requirement?
https://www.linkedin.com/in/darrenoconnor
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 07:55 AM
Thanks for the reply UCcert. I just sent am email to my account manager requesting the assistance of an SE.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 08:00 AM
Good luck with the project.
https://www.linkedin.com/in/darrenoconnor
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 12:45 PM
BGP is not supported on the VMX (at least, none of the VMXs I have access to show the BGP menu options).
Additionally, even if the VMX can do BGP, you can't do BGP to AWS inside of a VPC. You have to do something more complicated like run a GRE tunnel over IPSec to an AWS VPN gateway, and then run BGP over that.
In the past when needing to do HA with VMXs in AWS with static routing I have used a Lambda script to detect failure and swap our the routing. It's complicated. I created some instructions on how to do this.
https://www.ifm.net.nz/cookbooks/meraki-ha-vmx-amazon-aws.html
However, since I did my last one, there is now an AWS Gateway Loadbalancer service, and this is probably going to be a better way of doing it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 01:00 PM
Hi Philip,
I was hoping you would show up here.
This is not a saved configuration, but when I set this vMX to Hub, I have the option to enable and configure BGP. I have not saved this to know if it will take. I will be doing that in a change window tonight.
My plan was to keep this as simple as possible. I read your article about HA and the use of Lambda script. We are opting to do this as if they are 2 different hubs. They are both in the same AWS VPC, but in 2 different AZs. The idea is to configure the MX and Z3 to have both hubs listed like a primary and standby hub. We do that today with MX100s in 2 different data centers.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 01:02 PM
Even if you can enable BGP on the VMX (which looks like you can) there is no way you can BGP with AWS inside of the VPC. Amazon AWS doesn't offer this as an option.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 01:04 PM
One option I have often thought about but never done was to write a script to grab the routeing table from each VMX, and then add/remove/update static routes in the VPC. Then just run the script every minute as a Lambda script.
It would be like dynamic routing then.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 01:07 PM
Another idea we have been discussing is spinning up a Cisco CSR1000v inside the same VPC and BGP peering off of that.
I was looking for more of an AWS native solution instead of adding virtual routers into the mix. I saw that AWS has BGP for Direct Connects, but haven't found any other document that states where else in AWS BGP is supported.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 01:13 PM
AWS supports BGP for DirectConnect and over VPN tunnels. Nowhere else.
You can spin up a CSR1000v. You'll need a VPN licence for it. You would need to BGP peer with it (from the VMX), and then have it build a VPN to the VPC, and then run BGP over the VPN to the VPC.
Typically, people deploy a "transit vpc" when doing this. The VMXs and CSR1000Vs would go into this. Your current VPC would become a spoke of the transit VPC.
https://docs.aws.amazon.com/solutions/latest/cisco-based-transit-vpc/architecture.html
I would guess that the cost of running the CSR1000Vs in Amazon plus their purchase cost for three months would cover the cost of writing a Lambda script to dynamically update the routing tables, and be tremendously less complex ...
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 01:17 PM
Scratch the script idea. I just checked the API and you can only retrieve static routes, not the current dynamic routing table.
So you would have to do something much more complicated like run Zebra on Linux (Zebra is a very popular BGP routing engine), have the VMX peer with it, and then run a script on that box to dynamically update the AWS route tables from there.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-16-2021 08:44 AM
After weighing options like spinning up a csr1000v and other things, we decided to implement an ASAv firewall in AWS. This could be done with PaloAlto and other devices. This gets us around the need to do BGP peering as we have done in other traditional data centers. This should satisfy the routing issue by having a single default route from the vMX to the firewall and the firewall with a default route to the Internet and routes to our internal AWS resources. We also have an ASAv that is used for RA-VPN in AWS so putting both the vMX and the RA-VPN ASAv into a DMZ made a bit more sense for us. This provides an additional layer of security outside of what we may have configured with the AWS Security Groups.
Below is a basic diagram for anyone else looking for a deployment solution for vMX in AWS.
