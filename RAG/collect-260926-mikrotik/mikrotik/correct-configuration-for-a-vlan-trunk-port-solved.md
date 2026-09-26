---
id: collect-260926-mikrotik/mikrotik/correct-configuration-for-a-vlan-trunk-port-solved
title: "make ports slaves. The default configuration probably already has this."
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/correct-configuration-for-a-vlan-trunk-port-solved.md
source_anchor: ""
source_lines: [1, 62]
sha256: b7ed35edc92045cd97275c1fecf1ca68017e1085de00ede1668d0f7ee8241156
---

# make ports slaves. The default configuration probably already has this.

Thanks. I have seen the wiki, but the examples there are incomplete and don’t work without additional configurations (e.g. IPs or trunk bridge, if you have more than one switch chip). This is really annoying for beginners.

However, I recreated the configuration on the “main router” and now it works. It’s ridiculous how straight forward the final configuration is 

If somebody ever stumbles over this thread and is looking for a vlan configuration which is not based on “bridges, bridges, bridges” and actually uses the switch chip, here it is.

This is example is for an RB2011U, which has two switch chips (and therefore two master ports). It’s based on the “home AP” default config.

```
# make ports slaves. The default configuration probably already has this.
/interface ethernet
set [ find default-name=ether1 ] name=eth01-outside
set [ find default-name=ether2 ] name=eth02-master
set [ find default-name=ether3 ] master-port=eth02-master name=eth03
set [ find default-name=ether4 ] master-port=eth02-master name=eth04
set [ find default-name=ether5 ] master-port=eth02-master name=eth05
set [ find default-name=ether6 ] name=eth06-master
set [ find default-name=ether7 ] master-port=eth06-master name=eth07
set [ find default-name=ether8 ] master-port=eth06-master name=eth08
set [ find default-name=ether9 ] master-port=eth06-master name=eth09
set [ find default-name=ether10 ] master-port=eth06-master name=eth10
# trunk bridge, probably already exists
/interface bridge add comment=defconf name=bridge
# the master ports are probably already member of that bridge
/interface bridge port add bridge=bridge comment=defconf interface=eth02-master
/interface bridge port add bridge=bridge interface=eth06-master
# now, let's create the vlans on the bridge
/interface vlan add interface=bridge name=vlan-guest-200 vlan-id=200
/interface vlan add interface=bridge name=vlan-int-100 vlan-id=100
# vlan port config on the switch
/interface ethernet switch vlan add ports=eth02-master,eth03,eth04,eth05,switch1-cpu switch=switch1 vlan-id=100 independent-learning=no
/interface ethernet switch vlan add ports=eth02-master,switch1-cpu switch=switch1 vlan-id=200 independent-learning=no
# if you have a device (like the RB2011U) with 2 switch chips, you need an additional switch config for the second switch:
/interface ethernet switch vlan add ports=eth06-master,eth07,eth08,eth09,switch2-cpu switch=switch2 vlan-id=100
/interface ethernet switch vlan add ports=eth06-master,eth07,eth10,switch2-cpu switch=switch2 vlan-id=200
# define the adresses of your vlans
/ip address add address=10.10.100.1/24 interface=bridge-vlan-int-100 network=10.10.100.0
/ip address add address=10.10.200.1/24 interface=bridge-vlan-guest-200 network=10.10.200.0
# you might want to create a dhcp on each vlan, I'll skip that here.
```

For trunk ports, we do

```
/interface ethernet switch port set 2 default-vlan-id=1 vlan-header=add-if-missing vlan-mode=secure
```

Now for every access port:

```
/interface ethernet switch port set 3 default-vlan-id=100 vlan-header=always-strip vlan-mode=secure
```

If there are wifi AP stations on a vlan, we can define a bridge for said vlan and just add the wifi to it:

```
# create a bridge for each vlan you need
/interface bridge add name=bridge-vlan-guest-200
# add the vlan and wifi to the bridge
/interface bridge port add bridge=bridge-vlan-guest-200 interface=vlan-guest-200
/interface bridge port add bridge=bridge-vlan-guest-200 interface=wifi-guest-200
```
