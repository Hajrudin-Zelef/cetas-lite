---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-928947-how-to-set-up-pfsense-opnsense-on-virtual-host-f5b2c7a2
title: "questions-928947-how-to-set-up-pfsense-opnsense-on-virtual-host-f5b2c7a2"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-928947-how-to-set-up-pfsense-opnsense-on-virtual-host-f5b2c7a2.md
source_anchor: ""
source_lines: [1, 131]
sha256: a8b535b6c975a6576e35490c60bff3e51fffcfbdd9360e8c4a994aaf94577434
---

# questions-928947-how-to-set-up-pfsense-opnsense-on-virtual-host-f5b2c7a2

Unfortunately, the question is a little too imprecise and i don't have the reputation to ask further questions.
A few days ago I set up OPNSense in a virtual environment, so hopefully I can leave a few hints here.
I don't claim that this is a good solution but it works for me. 
In this scenario i have a dedicated server and a /29 public IP subnet.
All traffic from other VM's is routed through the OPNSense-VM. 
The traffic of the host machine cannot be sent through the virtual machine.

For administration of the networks and virtual machines i use WebvirtCloud
but it's also possible to do everything manually. 

1. Create virtual network interfaces for each of your public IPs. ```
/etc/network/interfaces
auto eth0
iface eth0 inet static
        address 103.x.x.104      #Dedicated server IP address
        netmask 255.255.255.255
        gateway 103.x.x.65
        pointopoint 103.x.x.65   #IP of the switch in the data center
auto eth0:0
iface eth0:0 inet static
address 103.xx.77.136             #First IP from the public subnet
netmask 255.255.255.255
auto eth0:1
iface eth0:1 inet static
address 103.xx.77.137
netmask 255.255.255.255
```
 [...]
2. Use libvirt to create the networks as described here: 
(or use WebvirtCloud) 
You need atleast two bridges, one for LAN the other one for WAN.

At this point you should have one network called LAN (Device: virbr1)(Network: 192.168.100.0/24)
and another one called WAN (Device: virbr0) (Network: 192.168.77.0/24)

1. Create your OPNSense Machine configuration
This is the important part:

```
<interface type='network'>
      <mac address='00:52:66:d7:7e:65'/>
      <source network='WAN' bridge='virbr0'/>
      <target dev='vnet0'/>
      <model type='virtio'/>
      <alias name='net0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </interface>
    <interface type='network'>
      <mac address='00:c0:41:50:f9:0b'/>
      <source network='lan' bridge='virbr1'/>
      <target dev='vnet1'/>
      <model type='virtio'/>
      <alias name='net1'/>
      <address type='pci' domain='0x0000' bus='0x02' slot='0x05' function='0x0'/>
    </interface>
    <interface type='network'>
      <mac address='00:98:c3:b1:b6:b8'/>
      <source network='lan' bridge='virbr1'/>
      <target dev='vnet2'/>
      <model type='virtio'/>
      <alias name='net2'/>
      <address type='pci' domain='0x0000' bus='0x02' slot='0x06' function='0x0'/>
    </interface>
```


1. Start OPNSense, assign interfaces according to your machine configuration and set interface IP addresses via the terminal.
I have selected 192.168.77.2 for my OPNSense WAN IP address.
The WAN upstream gateway is set to 192.168.77.1.

I also created seperate LAN's for each of my public IP's in OPNSense.

```
LAN136 IP-Address: 192.168.100.136/24 
LAN137 IP-Address: 192.168.100.137/32
No upstream Gateways.
```

[...]

1. NAT POST- and PREROUTING rules on hostmachine

```
iptables -t nat -A POSTROUTING -s 192.168.77.2 -j SNAT --to-source 103.x.x.104
iptables -t nat -A POSTROUTING -s 192.168.77.136 -j SNAT --to-source 103.xx.77.136
iptables -t nat -A POSTROUTING -s 192.168.77.137 -j SNAT --to-source 103.xx.77.137
iptables -t nat -A PREROUTING -p tcp --dport 10:65530 -d 103.xx.77.136 -j DNAT --to 192.168.77.136
iptables -t nat -A PREROUTING -p udp --dport 10:65530 -d 103.xx.77.136 -j DNAT --to 192.168.77.136
iptables -t nat -A PREROUTING -p tcp --dport 10:65530 -d 103.xx.77.137 -j DNAT --to 192.168.77.137
iptables -t nat -A PREROUTING -p udp --dport 10:65530 -d 103.xx.77.137 -j DNAT --to 192.168.77.137
```


After this step you should be able to open OPNSense from your web browser.
In order to do that you need a virtual machine that is already part of the LAN. 
I recommend using a live CD like grml or Ubuntu.
In this case, OPNSense can be reached via http://192.168.100.136

1. OPNSense (web browser) - Create virtual IPs on WAN interface

```
Virtual IP-address    Interface           TYPE
192.168.77.136/32         WAN          IP Alias
192.168.77.137/32         WAN          IP Alias
```


[...]

1. NAT Rules in OPNSense (web browser)
Create your port forwarding and outgoing rules

example port forwarding:

```
Interface Proto   S-address S-port    D-address       D-port       NAT-Ip         Nat-Port
WAN           TCP     *           *       192.168.77.137  80 (HTTP)   192.168.100.101 80 (HTTP)
```


example outgoing rules:

mode must be set to manual

```
Interface     Source                S-port    Dest.    D-port       NAT-IP        Port   static?
WAN               192.168.100.100/32      *        *        *         192.168.77.136  *       no
```


Make sure all your public IP's have outgoing rules and a unique NAT IP.
