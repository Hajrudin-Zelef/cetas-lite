---
id: collect-261001-general-networking/general-networking/questions-596-how-to-find-the-port-a-device-is-connected-to-based-on-ip-on-a-cis-3a572f19
title: "questions-596-how-to-find-the-port-a-device-is-connected-to-based-on-ip-on-a-cis-3a572f19"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-596-how-to-find-the-port-a-device-is-connected-to-based-on-ip-on-a-cis-3a572f19.md
source_anchor: ""
source_lines: [1, 48]
sha256: b7746f8fc24d052cad6bfe51089c895483cc10c98f15202ed1733c70fa905eda
---

# questions-596-how-to-find-the-port-a-device-is-connected-to-based-on-ip-on-a-cis-3a572f19

On a Cisco Catalyst switch, I need to find what port an end device is connected to.
I have the IP address/host name.
How do I quickly find the port?
The answer depends on whether the switch is a Layer 2 or a Layer 3 switch. That is to say, is the switch only switching and relaying traffic on to a different device for routing, or, is it doing the routing decisions itself via SVIs (switched virtual interfaces).
On a layer 3 switch, the port can be found by using a few simple commands on the device. However on a layer 2 switch, you have to log into both the switch and whatever device is doing the routing to locate the port.
In either case, the commands are the same, just run on two different boxes for the layer 2 switch.
Log into the switch and issue the following command (where ipaddress is the ip address of the host you are trying to locate:
show ip arp *ipaddress*
The output should look similar to below, and give you the mac-address of the device (I've highlighted the mac-address below in bold).
LYKINS-1861#show ip arp 172.20.1.100
Protocol  Address          Age (min)  Hardware Addr   Type   Interface
Internet  172.20.1.100            0   **28cf.da1d.1b05**  ARPA   Vlan10
Now issue one of the following commands (where mac-address is the hardware address from previous step). Depending on the Cisco platform, sometimes the command is listed in either form.
show mac address-table address *mac-address*
show mac-address-table address *mac-address*
The output should look similar to below, the port you are looking for should be listed as the Destination Port:
LYKINS-1861#show mac-address-table address 28cf.da1d.1b05
Destination Address  Address Type  VLAN  Destination Port
-------------------  ------------  ----  --------------------
28cf.da1d.1b05          Dynamic      10     **FastEthernet0/1/1**
Find out what device is doing the routing for this switch (you may have to look at the network documentation). Sometimes it is a "Router on a Stick", where the Layer 2 VLANs are being trunked up to the router for the Layer 3 decisions.
Log into the routing device, and issue the following command (where ipaddress is the ip address of the host you are trying to locate:
The output should look similar to below, and give you the mac-address of the device (listed below in bold).
Now log into the Layer 2 switch
On that device issue one of the following commands (where mac-address is the hardware address from previous step). Depending on the Cisco platform, sometimes the command is listed in either form.
Firstly, you need to get the MAC address, so get into a machine on the same VLAN and look at its neighbour table - Windows is netsh int ipv4 show neigh, Linux: ip nei Cisco: show ip arp x.x.x.x. Once you have that...
If this is a discovery job on a layer 2 switch, do show mac address-table | i 0011.2233.4455 - replacing the mac address bytes as appropriate.
If on the other hand it's a router, use show ip arp | i 0011.2233.4455 - again replacing the MAC as appropriate. Obviously when you were getting the MAC, if it turned out to be directly connected to that router, you're already done.
Long-term however, I heartily recommend that you setup LLDP (failing that, CDP) to your hosts so that you can identify them from either side. lldpd is an absolutely excellent LLDP daemon for Linux that also supports CDP, EDP, SONMP and FDP. If you're currently able to reach the host and it does happen to run linux/BSD, I'd recommend skipping the above and just turn on LLDP.
Regardless of whether your switch is performing frame forwarding (layer 2) or packet routing (layer 3), the following should work if the switch has a management IP address in the same subnet as the host you want to find:
show ip arp) and its corresponding interface.show mac-address table address [the address]
Quickly?
terminal monitor; press enter.
00:00:47: %LINK-3-UPDOWN: Interface GigabitEthernet2/0/1, changed state to down
Note, if you are connecting to the switch from the endpoint you are unplugging you will not get the syslog message because you will have cut connectivity.
Just telnet yourself into the switch, then unplug the device, then replug it. The switch will write a message which port was down and then up.
I hope this helps.
Enable ip device tracking on the switch. USe the following commands
ip device tracking ip device tracking probe delay 30
The second commands prevents the duplicate ip error from popping up on pc's. There is a Cisco bug for it but can't remember the id.
Will want older 12.4 code or 15 code to enable the above.
The CAM table on the switch matches the MAC address to the device port. Use 'show mac address-table' on the switch. You can also configure port descriptions to help find devices later too.
you want to find: 1.At the switch, ping the IP address you want to find. If the address is on the same subnet as the switch's management address, an ARP request will be sent looking for the MAC address of the host. 2.Now you can look at the ARP cache and find the MAC address of the host you are looking for (show ip arp) and its corresponding interface. 3.You can also look for its MAC address table entry by issuing: show mac-address table address [the address]
I have created a few python scripts to help with this and to document the devices connected to the edge switches.
The specific repositories are: pingSVI - takes the output of "sh run | i ^interface|^_ip address" parses the subnets and pings all hosts. Populates switch's arp table for devices that have timed out.
ARP-Sort - creates a json database of IP-Mac addresses from the core switch
MAC2Manuf - takes the output of "show mac add int" and uses the json database to create a list of IP-MAC-Port-Manufacture off the edge switch.
If you issue the command show mac address-table address x, you could also see this MAC being learned on a trunk port. If that is the case (and assuming you're connected to another Cisco device), use "show cdp nei" and log into that device and perform the same steps as above. Keep doing this until you see the MAC address being learned on a non-trunk port.
