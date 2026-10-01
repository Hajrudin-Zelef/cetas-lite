---
id: collect-261001-cisco/cisco/r-cisco-comments-12gem8-vlan-troubleshooting-5b707af0
title: "r-cisco-comments-12gem8-vlan-troubleshooting-5b707af0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "pruning", "voice"]
source: docs/RAG/collect-261001-cisco/r-cisco-comments-12gem8-vlan-troubleshooting-5b707af0.md
source_anchor: ""
source_lines: [1, 26]
sha256: a79811c307bdb8a7e843de80e956e38d012b5bb1c3e3792e113788e9607a37fc
---

# r-cisco-comments-12gem8-vlan-troubleshooting-5b707af0

VLAN troubleshooting 
        
    Hello all. I'm pretty new in the Cisco world as our CCNA quit earlier this year and I've had to take up all the networking duties now. I've created a set of new VLANs on our 4510r (10,11,12 - 172.16.3.1/24, 172.16.8.1/21, and 172.16.6.1/24 respectively) but am having some odd occurrences. Whenever I change my port on the switch to access one of the new VLANs (switchport access vlan XX) it seems to kill one of our existing VLANs (172.16.0.1/24) that we use for our phone systems from accessing our DHCP server in another VLAN (128.239.0.1/16). As soon as I move my port out of that VLAN the phones can register IPs again. If you want me to post some configs I'd be happy to, as this is really confusing me and I can't seem to make any progress getting this working correctly. Thanks for any help anyone could provide.
Section des commentaires
Are you using data and phone off of the same switch port? If so you will probably need to configure your switchports as trunks in order to allow use of more than one vlan.
this! if your using cisco phones put the commands
int gig 0/1 switchport mode dynamic switchport access vlan xx switchport voice vlan yy
if you are not using cisco phones but phones that tag the vlan themselves do
int gig 0/1 switchport mode trunk switchport trunk encapsulation dot1q switchport trunk native vlan xx
where xx is your DATA vlan -native means treat any untagged packets as being on the xx vlan
Thank you both! This was exactly what the problem was. It sucks being completely over my head sometimes in the networking field, but this really is the best way to learn how everything works. Thanks again for all your help.
I believe we are using the voice vlan command on all the PoE ports on the switch (although I'm not certain how to confirm this in bulk). We are using helper addresses to point to the DHCP servers. I'll post a follow up with the vlan and port configs shortly
are you using the switchport voice vlan command for the phones? is the 4510 doing DHCP or are you using a helper command? examples of your vlan config and port settings would be nice..
VLAN I just configured
Vlan11 is up, line protocol is up Internet address is 172.16.8.1/21 Broadcast address is 255.255.255.255 Address determined by setup command MTU is 1500 bytes Helper addresses are 128.239.0.203 128.239.0.25 128.239.0.26 Directed broadcast forwarding is disabled Outgoing access list is not set Inbound access list is not set Proxy ARP is enabled Local Proxy ARP is disabled Security level is default Split horizon is enabled ICMP redirects are always sent ICMP unreachables are always sent ICMP mask replies are never sent IP fast switching is enabled IP Flow switching is disabled IP CEF switching is enabled IP CEF switching turbo vector IP multicast fast switching is enabled IP multicast distributed fast switching is disabled IP route-cache flags are Fast, CEF Router Discovery is disabled IP output packet accounting is disabled IP access violation accounting is disabled TCP/IP header compression is disabled RTP/IP header compression is disabled Probe proxy name replies are disabled Policy routing is disabled Network address translation is disabled WCCP Redirect outbound is disabled WCCP Redirect inbound is disabled WCCP Redirect exclude is disabled BGP Policy Mapping is disabled
VLAN for the phones
Vlan4 is up, line protocol is up Internet address is 172.16.0.1/24 Broadcast address is 255.255.255.255 Address determined by non-volatile memory MTU is 1500 bytes Helper addresses are 128.239.0.203 128.239.0.25 128.239.0.22 Directed broadcast forwarding is disabled Outgoing access list is not set Inbound access list is not set Proxy ARP is enabled Local Proxy ARP is disabled Security level is default Split horizon is enabled ICMP redirects are always sent ICMP unreachables are always sent ICMP mask replies are never sent IP fast switching is enabled IP Flow switching is disabled IP CEF switching is enabled IP CEF switching turbo vector IP multicast fast switching is enabled IP multicast distributed fast switching is disabled IP route-cache flags are Fast, CEF Router Discovery is disabled IP output packet accounting is disabled IP access violation accounting is disabled TCP/IP header compression is disabled RTP/IP header compression is disabled Probe proxy name replies are disabled Policy routing is disabled Network address translation is disabled WCCP Redirect outbound is disabled WCCP Redirect inbound is disabled WCCP Redirect exclude is disabled BGP Policy Mapping is disabled
And the configuration for my port in the switch
Voice VLAN: none Administrative private-vlan host-association: none Administrative private-vlan mapping: none Administrative private-vlan trunk native VLAN: none Administrative private-vlan trunk Native VLAN tagging: enabled Administrative private-vlan trunk encapsulation: dot1q Administrative private-vlan trunk normal VLANs: none Administrative private-vlan trunk private VLANs: none Operational private-vlan: none Trunking VLANs Enabled: ALL Pruning VLANs Enabled: 2-1001 Capture Mode Disabled Capture VLANs Allowed: ALL
Unknown unicast blocked: disabled Unknown multicast blocked: disabled Appliance trust: none
can you show the "show run" information of the vlans and port setting for one working and non-working port? we dont need the whole config just the parts about the interfaces, dhcp, and vlan info..
Can we see the following command:
sh run int gig 9/2
This will show us the switchport config for that interface, I have a sneaking suspicion you don't have a voice vlan configured and that's why stuff is breaking. This was my first clue :)
Voice VLAN: none
I did change the voice vlan to 4 for this port and it brought all the phones down in our little area. I just realized that port is plugged into a little Dell Powerconnect which may be part of my problem.
