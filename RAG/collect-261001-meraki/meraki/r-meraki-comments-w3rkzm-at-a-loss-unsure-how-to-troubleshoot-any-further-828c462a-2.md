---
id: collect-261001-meraki/meraki/r-meraki-comments-w3rkzm-at-a-loss-unsure-how-to-troubleshoot-any-further-828c462a-2
title: "r-meraki-comments-w3rkzm-at-a-loss-unsure-how-to-troubleshoot-any-further-828c462a"
domain: meraki
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agents", "ethernet", "license", "voice"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-w3rkzm-at-a-loss-unsure-how-to-troubleshoot-any-further-828c462a.md
source_anchor: ""
source_lines: [7, 56]
sha256: 3a69eba7c9374d3aed8a3decf5d6ad5e12b662467de76259c6168e1a285028dc
---

# r-meraki-comments-w3rkzm-at-a-loss-unsure-how-to-troubleshoot-any-further-828c462a

    Good afternoon,
I come to you r/Meraki in search of help. I (and multiple Meraki engineers) are at a complete loss, and I feel we're grasping at straws as to how proceed.
To keep things as simple as possible, there are (2) Meraki devices. Head Office MX84 (10.0.0.0/16) and Satellite office (10.20.0.0/16). Head Office has a 100M/100M fiber connection, Satellite Office has a 120M/20M cable connection. Both have static IP's -- direct connection to the internet (cable modem is bridged). Both offices maintain a hub/spoke configuration.
VLAN configuration is the exact same both all locations:
- 
      VLAN 10 = x.x.10.x/24 = LAN
- 
      VLAN 20 = x.x.20.x/24 = Voice
VLAN 10 has "VPN mode" enabled; internet traffic is routed the local MX.
VLAN 20 has "VPN mode" disabled. Phones are 100% cloud based and do not need to interact with anything on either side of the tunnel. DHCP is handled by the Meraki, pushing Google DNS, and has a Group Policy applied on the VLAN blocking any RFC1918 traffic.
No inter-VLAN rules, and very basic configuration on both sides. Very simple stuff. Maybe 2-3 users maximum at the Satellite Office.
I can sit at the Head Office, at Home, or at another customers site and successfully ping the Satellite Office modem and Meraki MX67W WAN interface without issue. It does not miss a beat. However, whenever the Meraki decides to execute the chaos monkey protocol -- all heck breaks loose. If I'm sitting on VLAN10 in the Head Office and ping VLAN10 in the Satellite Office -- I will randomly just start dropping packets. When this happens, it's almost as if the internal software bridge just dies for [random interval] until a service restarts, as when I drop packets on the other side of the tunnel, the users on VLAN20 will have their phones disconnect (which should not be possible).
I've tried everything, new cables, new power adapters, RMA'd MX's, factory resets, creating a new site, upgrading to various firmware builds (currently on 16.16.3 as the 17.8 beta randomly broke everything), moving to different registration servers and nothing helps.
Whilst I was waiting for my RMA I had to provide a workaround to keep the packets flowing, so I dug out some old Cisco 881's and set them up at both locations without a single issue for over a month straight. They both just ran without any drops, disconnects, or issues of any kind so I have an incredibly difficult time pointing the finger at the ISP -- however, I honestly just don't understand what to do next short of just calling it and leaving the 881's in place -- but it seems silly when I've paid for an MX67W (with advanced security license) and have to treat it as a doorstop.
Any help or next steps would be appreciated. I'm truly at a loss -- it just doesn't make sense.
Cheers
Edit: Throwing some additional random bits of info here
- 
      I have another office (in a different city) with the same hardware as this satellite office -- no issues
- 
      There is a APC UPS connected correctly
- 
      The room is air conditioned -- hardware (should) not be overheating
Section des commentaires
Were there any issues with the setup on MX versions prior to 16?
I had a few customers who had to revert due to application weirdness once the MX 16.x firmwares were released.
Yes. 15.44 (and below) were absolutely the most stable -- but even then drops would still happen. Part of the Meraki troubleshooting was to upgrade everything to the latest build, multiple times -- hence why I'm almost completely on 17.8.
Does the packet loss happen when you start the ping from the satellite office to the head office?
Is the VPN set as full mesh or spoke to hub?
Im probably sure youve tried this but just want to verify. Disable all advanced security settings on both MX units and test from both directions to see if it drops again.
Support has the ability to take Pcaps on the VPN interface so I would have them take that capture while pinging across the tunnel from both a client device behind each MX and also from the MX units themselves to help eliminate some variables.
I vaguely remember something similar I worked on back in 2018/2019 with MX84s doing weird things with the VPN interface due to how that unit was made. Basically they threw a switch interface in the MX unit that had a ton of software to hardware issues like the VPN interface dropping when you make any DHCP changes.
My best guess its something with the MX84 itself
Yes. It is seemingly random. I can go a day, week, month without issue. However when chaos monkey is let loose, watch out.
Hub and spoke.
Correct. Same thing.
Unfortunately when chaos monkey is released and/or I start losing any type of connectivity the "internal" side of the MX just dies. No VPN traffic, nothing. The WAN is the only thing that responds oddly enough.
MX84 has a hardware issue that will always exist -- depending on the change you make, you can drop a couple packets on any/all interfaces (it's super fun to work around)
I considered that, but I do have another Satellite Office (not outlined in this scenario) which has been rock solid (up until 17.8 -- which required me to do a factory reset and has been fine since). It uses the 10.10.0.0/16 range and actually uses the same ISP/service (Rogers Cable - 120/20 w/ 1 static IP). Just different city.
Ty for the response! I'm glad it wasn't something stupid right off the hop.
Just to confirm no other DNS servers other then Google’s? What’s your lease time and when do your start dropping packets? Do you drop packets when you ping every device or is it only a single device? Are you specifying packet size in your ping commands?
Nope. I originally had phones proxy to upstream DNS (Meraki is configured to use Google DNS) so figured it was prettier. Part of my troubleshooting process, I removed that from the equation and just pass "Google DNS" directly -- on the phones, I can see 8.8.8.8/8.8.4.4 being used.
DHCP lease time is 1 week across the board. (Guest WiFi is 1 hour, but that's out of scope).
No -- default ping sizes were used; but that's not really the issue; rather just one of the tools I would use to initially detect the outage. I would see the agents drop offline in the call center and have all (3) pings going -- modem external IP, WAN external IP or Meraki, and Meraki VLAN IP -- and without fail, when the agents would disappear so would my pings to the Meraki VLAN IP (Data VLAN -- VLAN 10).
And I’m guessing the issue occurs one a week or so? Do yourself a favour and go to that site and double check all of your phones only have a single Ethernet connection. I’ve run into this issue with Meraki so many times. If you have a switch that has spanning tree enabled and isn’t configured for port isolation and one of your phones is double connected your Meraki will flip out and seemingly shutdown at lease renewal time.
Hi
Sorry youre going through so much grief.
I had not yet read the whole thread, so if its solved, please disregard.
BUT, regarding the use of Google as your DNS, and sending ping (IGMP) traffic. We have done extensive work with our enterprise ISP and they informed us that google is constantly hammered by ping traffic. Therefore they tend to have high loss rates. Im not saying this is the cause of your issue, but you might consider using a different DNS and see if it helps, such as your ISPs, or Cloudflare (1.1.1.1, or 4.2.2.2 ) These have a much friendlier public DNS. Google is great, but not if your sending IGMP traffic.
Check the monthly or weekly resolution graph in the SD-WAN, Appliance Status, Uplink tab.
