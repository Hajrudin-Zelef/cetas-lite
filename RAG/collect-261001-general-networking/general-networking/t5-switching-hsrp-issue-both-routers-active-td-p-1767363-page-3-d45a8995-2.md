---
id: collect-261001-general-networking/general-networking/t5-switching-hsrp-issue-both-routers-active-td-p-1767363-page-3-d45a8995-2
title: "t5-switching-hsrp-issue-both-routers-active-td-p-1767363-page-3-d45a8995"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-general-networking/t5-switching-hsrp-issue-both-routers-active-td-p-1767363-page-3-d45a8995.md
source_anchor: ""
source_lines: [73, 177]
sha256: cf1b4e9c1b92e61e8536a04e0dbb747cd60c7fd65eebf815974b868df74b8ae4
---

# t5-switching-hsrp-issue-both-routers-active-td-p-1767363-page-3-d45a8995

			Other Switching
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-01-2012 04:39 AM
Hi Shane
5.1(3)N2(1) is a Nexus 5000-Release, the above Problem occured on a Nexus 7000, where 5.1(3) was another NX-OS-Generation.
AFAIK, 5.1(3)N2(1) still is considered a good Release. The bug which stood behind my issue never was present on this release.
As mentioned in my last post, I resolved my issue not with upgrading, but with disallowing/allowing the one faulty VLAN on the vPC-Link between both Nexus. Please be aware that this would not help if a) you have Dual-Active-HSRP on multiple Vlans or b) if neither of both Nexus receive HSRP-Packets from the other (debug hsrp engine packet hello)
Greetings from Berne, Switzerland
Stefan Mueller
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-02-2012 07:06 AM
Thanks Stefan, I see the error of my ways, this does not relate to 5K as you have pointed out.
My issue turned out to be related to LANBase license not installed on L3 daughter cards out of the box. Once installed, my issues were resolved.
Kind regards.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-15-2013 12:59 AM
Hi,
We encouner similar issue that 2 x N7K w/ SUP1/M1/FAB1 claim itself HSRP active after NXOS upgrade from 6.1.x to 6.2.x. Anyone has clue?
Tks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-10-2013 06:57 PM
Leon....I just had this same issue (active/active HSRP) today on my two 7010's running 6.2.2a code. To resolve, I made sure that my VPC role priority for my vpc domain was such that my 'by design' active HSRP device and the Primary VPC role were in fact one in the same device. I then bounced my peer link port channel.
If you have a case where your hsrp config is set to a high priority on switch 'A' and your vpc role is primary on switch 'B', you may experience issues with the HSRP multicast traffic traversing the peer link due to the loop prevention methodology within vpc. At least that is my hypothesis from reading miscellaneous articles and troubleshooting threads. I did not have time to open a TAC case for a root cause but the above steps resolved my issue on 6.2.x code.
HTH
JW
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-10-2013 09:31 PM
Hi John,
Do you mean that HSRP master must be VPC primary device?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-10-2013 10:12 PM
It appears so. I also employed the peer-gateway component under my vpc domain but I don't think it was the root cause of the active/active HSRP scenario that I saw.
Take a look at this doc and I carved out the relevant quotes below:
"Features That You Must Manually Configure on the Primary and Secondary Devices
You must manually configure the following features to conform to the primary/secondary mapping of each of the vPC peer devices:
•HSRP active—If you want to use Hot Standby Router Protocol (HSRP) and VLAN interfaces on the vPC peer devices, configure the primary vPC peer device with the HSRP active highest priority. Configure the secondary device to be the HSRP standby and ensure that you have VLAN interfaces on each vPC device that are in the same administrative and operational mode. (See the "vPC Peer Links and Routing" section for more information on vPC and HSRP.)
vPC Peer Links and Routing
To simplify initial configuration verification and vPC/HSRP troubleshooting, you can configure the primary vPC peer device with the FHRP active router highest priority.
JW
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-18-2015 02:02 AM
I hit the same on
show version internal build-identifier
Kickstart image file: bootflash:///n7000-s1-kickstart.6.0.2.bin :  S31
System image file: bootflash:///n7000-s1-dk9.6.0.2.bin :  S31
Removing and adding the vlans solved the issue
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2021 05:57 AM
i have/had this same issue, on a Nexus 9k, NXOS: version 7.0(3)I7(2)
no hsrp hellos were being sent to the other host.
no vpc involved here, just a regular port-channel to an identical 9k.
removing the vlan from the port-channel and then re-adding, kicked hsrp into life and the hellos immediately started sending again, thus resolving the hsrp active/active issue.
it was only affecting one specific vlan, others were unaffected.
thanks to Stefan for providing the answer.
this bug still seems to be lurking in NX-OS.
- « Previous
- Next »
