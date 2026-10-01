---
id: collect-261001-general-networking/general-networking/questions-1651515-can-i-block-network-traffic-for-a-specified-ip-range-on-a-ubiq-9177e0f9
title: "questions-1651515-can-i-block-network-traffic-for-a-specified-ip-range-on-a-ubiq-9177e0f9"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1651515-can-i-block-network-traffic-for-a-specified-ip-range-on-a-ubiq-9177e0f9.md
source_anchor: ""
source_lines: [1, 20]
sha256: 5c088a7cc84f82cc6a62485e6cdccf7220961160e8a531efed9fcd5d2e7d2080
---

# questions-1651515-can-i-block-network-traffic-for-a-specified-ip-range-on-a-ubiq-9177e0f9

Kind of new to the whole networking situation, but I'm working on a project that has my home network stretching out over my property to two different locations. So I have my home, a Barn, and a 2nd smaller inlaw house. I can't run wires to each of these buildings, So I have three nanobeams configured in PTMP configuration to connect my home to the two smaller buildings. The main network that connects to the outside world is configured as 10.8.0.*. So this network will be used for the primary devices in all three locations so that it has access to the network gateway.
However, I have a project that is set up in the barn and the inlaw house, that needs to stay local to those buildings. The project is configured as a 192.168.1.* network. The project is a duplicate of each other in both locations, so the project in the barn is using the same IP addresses as what's set up in the inlaw house. I need to prevent this 192.168.1.* traffic from leaking out of the individual locations and into the other.
I thought that when I configured the Nanobeams and set them up with 10.8.0.* addresses that that would have prevented any 192.168.1.* traffic from getting out. However, it appears that I'm having several IP Address conflicts. Is there a simple way, via the NanoBeams only, to prevent 192.168.1.* traffic from getting onto the NanoBeam WAN?
I don't want to use Unifi Software to create any rules. I want the rules only to be configured via the NanoBeams themselves. There are times that I will want to shut down access to the main house from the main house, but still want traffic from the barn to the inlaw house.
As I'm going to get asked, it's a brewing setup. I'm testing temperature, humidity control, and a slew of other sensors that work in a cluster, on the 192.168.1.* network at the individual location only. However, the cluster does share information with the other cluster, but via the 10.8.0.* network only. I need to be able to share the data via the 10.8.0*, but prevent any traffic getting out over the Nanbeam network that's related to the 192.168.1.*.
What's happening is when I have it up and running, the software takes a reading from some of the remote sensors on the network, but it causes the sensors in the other building to stop posting. There are zero issues when the NanoBeam is not connected, but once I start up the PTMP network sharing information, there is IP cross talk on the 192.168.1.*.
What's also weird is that from the main house, I can not ping the 192.168.1.* equipment in either the barn or the inlaw house. However, when if I remote access the Primary SBC at either location, I can ping my desktop in my house. So it seems traffic coming out of the barn and inlaw house allows 192.168.1.* out, but I can't access from the house to the other locations on the 192.168.1.*.
Any suggestions??????? Please????
For those not familiar with the Ubiquiti NanoBeam, let me know and I can provide screen shots of the settings available? However, I'm assuming that there is something under the Network Tab that would allow me to block the traffic. The following are the options available to me.
- IP Aliases
- VLAN Network
- Bridge Network
- Static Routes
- Firewall
- Traffic Shaping
I'm going to start searching for information related to these, but I'm a little lost in the weeds as to what I'm searching for.
UPDATE
So I got some advice and some more digging into the settings of AirOS8. I ended up adding a Firewall rule. [![Firewall rule 1][1]][1] [1]: https://i.sstatic.net/jkeP7.png
This seemed to work in that it no longer allowed me to ping my home computer from either the barn or inlaw space. However, it still is causing issues with the one particular sensor. Its not interfering with any other sensors.
I think I've come to the conclusion, after doing some reading that this particular sensor is doing something called multicasting. I'm going to setup Wireshark and see what information I can find.
