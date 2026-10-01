---
id: collect-261001-fortinet/fortinet/will-coquillette-src-content-blog-google-fiber-with-fortigate-90d-md-at-5adbf607d3d22288a0
title: "will-coquillette-src-content-blog-google-fiber-with-fortigate-90d-md-at-5adbf607d3d22288a0db9df9b9df"
domain: fortinet
role: reference
task: reference
actors: ["Google"]
dates: ["2016-06-13", "2016-07-08"]
keywords: ["latency"]
source: docs/RAG/collect-261001-fortinet/will-coquillette-src-content-blog-google-fiber-with-fortigate-90d-md-at-5adbf607d3d22288a0db9df9b9df.md
source_anchor: ""
source_lines: [1, 51]
sha256: f3669b1dc14184c25ccbc3f8149428b96a9898d196292d42ca3d6a2b3f595979
---

# will-coquillette-src-content-blog-google-fiber-with-fortigate-90d-md-at-5adbf607d3d22288a0db9df9b9df

| title | Google Fiber with Fortigate 90D | 
|---|---|
| description | Tutorial for configuring a Fortigate 90D firewall router in place of a Google Fiber network box for use with the Google Fiber ISP | 
| pubDate | 2016-06-13 | 
| updatedDate | 2016-07-08 | 
| hero | ~/assets/heros/fortigate_googlefiber.jpg | 
| heroAlt | Fortinet logo with Google Fiber rabbit | 
| tags |  | 

| CoS | firewall | firewall policy | fortinet | google fiber | fiber | QoS | vlan | 

*2016/07/08 UPDATE - PLEASE READ: The following configuration will allow your FortiGate to work with Google Fiber. However, after reviewing the hit count for the three policies outlined in this guide, it appears that **my DHCP and IGMP policies are not matching any traffic**. Apparently, absolutely all traffic from all services, including DHCP and IGMP, is being assigned an 802.1p CoS bit of 3.*

*I don't know why DHCP and IGMP traffic is not matching my policies, but I have great download/upload speeds and low latency in games, so I am not concerned. That being said, I am leaving my first two rules enabled so I can continue to monitor. When I perform the unencrypted Gigabit/Fiber speedtest at DSLReports, my bufferbloat exceeds +1400ms and gets an F rating.*

**In this tutorial, I am using the following hardware:**

- Google Fiber fiber jack
- Fortinet FortiGate 90D-POE Firewall

**Before we begin:**

- I am using the Web-based Manager in FortiExplorer version 2.6.1083 while connected via the USB management port
- My internal LAN is already set up to use the default hardware switch utilizing all ten LAN ports and custom DHCP settings for my personal network
- I configured the internalA interface for my FortiAP 321C and disabled internalB, C, and D.
- I disabled wan2
- I have successfully tested the following configuration under firmwares 5.4.0-build1011 and 5.4.1-build1064.

1. Network > Interfaces

1. Policy & Objects > Services
  1. Create New > Service
  2. Assign IGMP as the name, protocol type IP, Protocol Number 2, click OK.
2. Policy & Objects > IPv4 Policy
  1. Create New.
  2. Assign a name (I chose "QoS DHCP"), select your internal subnet(s) for incoming interface, select the VLAN interface you created in last step for outgoing interface, source all, destination all, schedule always, service DHCP, action accept, NAT enabled, do not assign any security policies, click OK.
3. Policy & Objects > IPv4 Policy
  1. Create New.
  2. Assign a name (I chose "QoS IGMP"), select your internal subnet(s) for incoming interface, select the VLAN interface you created in last step for outgoing interface, source all, destination all, schedule always, service IGMP, action accept, NAT enabled, do not assign any security policies, click OK.
4. Policy & Objects > IPv4 Policy
  1. Create New.
  2. Assign a name (I chose "QoS All Others"), select your internal subnet(s) for incoming interface, select the VLAN interface you created in last step for outgoing interface, source all, destination all, schedule always, service ALL, action accept, NAT enabled, select the security policies you wish to use on the traffic between your LAN and the WAN, click OK.
5. Make sure "QoS All Others" comes after the first two policies in the list of policies. I put them in the order of IGMP>DHCP>ALL. I'm not sure if the order of the first two matters, but both of them must be processed before the third policy, as you want IGMP and DHCP services to match in their own respective policies before the third policy, which includes them, has a chance to match them. This is because each policy will have a unique QoS bit.

*Note: The QoS bits come from here. Each policy has a unique number assigned to it, and they may be different from mine. My IGMP policy is 9, DHCP is 8, and All Others is 2.*

1. Login to the CLI of your FortiGate
2. Enter the following commands **using your own policy numbers as determined by the  "show" command after  "config firewall policy":**

`config firewall policy``show``edit 9``set vlan-cos-fwd 6``end``config firewall policy``edit 8``set vlan-cos-fwd 2``end``config firewall policy``edit 2``set vlan-cos-fwd 3``end`
That's it! With these firewall rules in place, I get over 900 Mbps down/up. Please leave a comment if this worked for you!
