---
id: collect-261001-meraki/meraki/r-meraki-comments-1av4jf9-file-transfer-corruption-over-autovpn-d517dbb3
title: "r-meraki-comments-1av4jf9-file-transfer-corruption-over-autovpn-d517dbb3"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["data centre"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-1av4jf9-file-transfer-corruption-over-autovpn-d517dbb3.md
source_anchor: ""
source_lines: [1, 54]
sha256: 0b547075778968a25bc971227476e4ae4d04d6e9d029969d12ab93954b7767de
---

# r-meraki-comments-1av4jf9-file-transfer-corruption-over-autovpn-d517dbb3

File transfer corruption over AutoVPN 
        
    Hi all,
Sorry for the long rambling post, it's been a long issue!!
We've recently been migrating from a MPLS config with a single internet breakout to a Meraki SDWan config with DIA at each site. In the old MPLS setup, everything worked great, so we assumed that SDWan would be a decent upgrade, as it would give us direct internet breakouts at each site for all those pesky teams calls and office 355 things we all seem to need now.
Under advise from our reseller, we've setup a hub + spoke approach, with 5 of our larger sites as Hubs, and 30 odd smaller sites as spokes
This has been working perfectly at all but 2 of the hub sites, which do a socket based FTP file transfer of encrypted database zip files between 2 servers from site "N" to site "O"
The majority of the files received are crc checked and unpacked correctly, however a small percentage are failing crc checks when they reach site "O". This causes the database to fall out of sync, and our reports about site N stop working at site O
Site O is the largest of the sites, and has a 200Mbps fibre DIA line, and a MX105 - this was chosen as the site to connect any client VPN users to, so we needed the larger MX to cope with a peak of 200 VPN users along with the LAN devices.
All of the other sites have MX67 devices, the remainder of the hub sites have 100Mbps fibre DIA (including site N), and the remainder have a mix of 100Mbps DIA, FttP, ADSL and EFM lines dependant on what they have available and price etc.
We have ubiquiti switches and WiFi at all sites, with a self hosted controller VM at site O to manage these devices.
We have purchased the new sdwan solution (internet connections at each site, Meraki MX at each site, full installation and support everywhere etc)as package deal from our existing MPLS provider - one of the major UK based telecoms providers, let's call them Provider for this discussion.
For migration purposes, the MPLS has needed to stay connected alongside the new Meraki MX boxes, so on each MX, a spare Lan port has been configured as VLAN 100 and connected to the MPLS routers, and appropriate firewall rules added to ensure traffic is routed over MPLS to sites that haven't migrated to SDWan yet, or for home VPN users on the old client etc.
Since the implementation, the issues started with the file transfers our database software uses. For reference, the database software is HQBird, and their support team have been very helpful trying to determine the issues, resulting in a diagnosis of an issue with the new internet setup.
Through troubleshooting steps with the implementation team at our provider and with Meraki tech support (Inc the escalation team) we seem to have gotten nowhere.
Things we have tried:
- 
      adding traffic shaping rules to limit heavy internet users, such as our off-site backup software veeam for taking up too much, no effect
- 
      changing back to MPLS only at site O - this fixed the issue
- 
      doubling the SDWan dia line speed to 200Mbps at site O, no effect
- 
      hardware swap out of MX 105 at site O - Meraki support say they saw tiny 100% CPU spikes at times, although web log shows around 20%, no effect
- 
      turn snort engine to passive mode at site O, no effect
- 
      upgrade to 18.2, whitelist source server from snort engine via web gui, no effect
- 
      remove traffic shaping rules again
The provider have been very good, and upgraded the internet at site O for free as part of the troubleshooting steps, and have organised a MX250 as a trial to see if that would help, but ideas seem to be running dry.
This has all been ongoing since the beginning of January, so we understandably want a resolution soon, hence turning to all of you lovely people who may have seen similar issues before.
Thanks in advance!!
Section des commentaires
What link speeds are using, AUTO?
What about MTU settings?
Our servers on both sites are MS Hyper-V VMs running Server 2019 data centre. We use jumbo frames (MTU 9000) on all of our servers devices to take advantage of the 10GbE Lan / switches we have at each site
Meraki have ruled out MTU as being an issue on the wan links, our ISP uses a MTU of 1500 on the internet connections
Although they can see packets fragmenting and being reassembled over the autoVPN tunnel
For reference, the files that are being transferred / corrupted range from 2KB to 4096KB in size
Then what about link speeds you mentioned some 100mbps links, is the other side of that connection hardcoded to match or set to Auto?
This sounds like a reassembly issue with autoVPN. What is your packet loss? Do you see retransmission of fragments due to PL?
I’ve seen issues when small fragments are lost leading to overlapping retransmission. Overlapping fragments require a decision on which payload to use and there are 5 different permutations of reassembly ( see Paxson Shankar fragmentation ) and can lead to corruption if reassembled by something other than the destination
Check your mtu settings. We had a similar issue on Meraki for one of our customers and lowering the mtu to conform to isp settings worked for us. I would start with lowering mtu significantly to 1100 then tweaking upwards from there.
I can't see where to change MTU settings on the MX? We use jumbo frames within the LAN at each site, and don't want to lower the MTU of LAN traffic, so any MTU change needs to be for WAN side only if we do make changes
You can't, have to ask support to do it. And yeah, I've had exactly the same issue multiple times and it's been MTU related. In my case it was always some sort of PPP connection in the mix, which needed an MTU of 1492 instead of 1500.
I believe this is the tool I used to track it down real quick:
https://www.elifulkerson.com/projects/mturoute.php
Smells like mtu problem
CRC checks or CRC errors on a network port?
The database software runs crc /md5 checks on the file that is received before unpacking it,
Are you using content filtering, AMP, or any security features on the MX, what do the event logs say for the problem device?
Sound like mss related issue... To big to fit in mtu
Try lower mss
