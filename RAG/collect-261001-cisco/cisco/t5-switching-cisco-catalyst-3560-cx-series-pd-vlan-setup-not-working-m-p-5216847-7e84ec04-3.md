---
id: collect-261001-cisco/cisco/t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04-3
title: "t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04.md
source_anchor: ""
source_lines: [316, 382]
sha256: 0f2eb631e749534c9d15ad164ea26d94f074bb184d18590ac9a604131d8dc76a
---

# t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04

I might say the problem does not seems to be on the switch any more. If you have interface up, vlans UP and ip routing. there is no reason to not ping.
"The hosts I am testing them with have some bridges setup for QEMU and maybe expecting tags - so if the switch configuration looks good, I can switch start looking for issues on the connected hosts. so far, all I have done on the hosts are these commands
sudo ip addr add 192.168.10.2/255.255.255.0 dev enp0s25 (a static ip in the 10 subnet)
sudo ip route add 0.0.0.0 via 192.168.10.1 --> which is Vlan 10's IP on the switch"
As you are running linux, I think it is a good idea to check firewall on the host.
Eventually, test with windows os to make sure. The switch seems to be fine now to me.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 09:30 AM - edited 10-29-2024 10:55 AM
Hi,
Can you ping from the switch to any host in any of the VLAN's? Please provide confirmation. From switch side, provide following outputs: "show vlan brief", "show vlan id 10", "show vlan id 20", "show interfaces trunk", "show ip interface brief", "show spanning-tree vlan 10", "show spanning-tree vlan 20", "show ip route", "show ip cef", "show ip arp", "show mac address-table", "show version".
Best,
Cristian.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 12:29 PM - edited 10-30-2024 05:41 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-30-2024 06:06 AM
Hi,
It all looks good, meaning you have ports Gi0/1, Gi0/2 in VLAN10 and Gi0/3, Gi0/4 in VLAN 20 (except that on STP output for VLAN 20, port Gi0/4 does not show up; did you not paste complete output or have disabled STP on that port via BPDFilter?); assuming hosts connected to these ports have IP addresses from the correct subnet, you should be able to have connectivity between PC's and default gateway which is the switch, while based on routing also between hosts in different VLAN's assuming you'v set the correct gateway IP address (the switch) on the hosts. Not sure what OS are the hosts running, however, try to ping the switch from hosts (the other way around, from switch to hosts may not work as maybe hosts have firewall turned on which filters ICMP packets). Based on the ARP and MAC table, you should at leat be able to ping from host 192.168.10.2 to switch which is 192.168.10.1.
ARP entries on the switch will show up only if there is IP communication between switch and host (if you ping the switch from all hosts, the switch should have all ARP entries); MAC entries will show up on switch only if the hosts sends any kind of traffic for switch to learn the MAC address.
Best,
Cristian.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-30-2024 07:34 AM
@Cristian Matei Thank you so much, I believe I am one step closer to the solution at this point. my hosts are linux machines, and had someone make some changes to remove bridges running on both the hosts today, so now
the host connected to VLAN 10 is able to successfully ping the gateway IP - that was not working so far
the host connected to VLAN 20 is NOT able to successfully ping the gateway IP "Destination Host Unreachable" which means there is some issue on the switch. Here are the ONLY differences I saw today, with your list of commands (ran all of them again this morning)
192.168.10.0/24 attached Vlan10
192.168.10.0/32 receive Vlan10
192.168.10.1/32 receive Vlan10
192.168.10.2/32 attached Vlan10
192.168.10.255/32 receive Vlan10
192.168.20.0/24 attached Vlan20
192.168.20.0/32 receive Vlan20
192.168.20.1/32 receive Vlan20
192.168.20.2/32 attached Vlan20 <----- NEW LINE TODAY, AFTER CHANGES TO HOST CONNECTED ON VLAN 20
192.168.20.255/32 receive Vlan20
Protocol Address Age(min) Hardware Addr Type Interface
Internet 192.168.10.1 - 308b.b2e2.c3c1 ARPA Vlan10
Internet 192.168.10.2 44 20c5.eb9a.299c ARPA Vlan10
Internet 192.168.20.1 - 308b.b2e2.c3c2 ARPA Vlan20
Internet 192.168.20.2 47 20c5.eb9a.298d ARPA Vlan20 <----- NEW LINE TODAY, AFTER CHANGES TO HOST CONNECTED ON VLAN 20
Can you please explain "(except that on STP output for VLAN 20, port Gi0/4 does not show up; did you not paste complete output or have disabled STP on that port via BPDFilter?);" I did not have any hosts connected on Gi0/4 - can you explain this in a bit more detail ? One last step before I can 100% declare the Switch is good, mark yours as the right answer (helping with all the commands to make sure I can verify if the config was good). Or any other suggestions on why I would get a "Destination Host Unreachable" for a host on Vlan 20 and not on Vlan 10.
Once again, thanks much !
