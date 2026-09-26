---
id: collect-260926-mikrotik/mikrotik/github-oallart-mikrotik-routeros7-vlan-documentation-and-code-to-manage-vlans-in-routeros--1
title: "github-oallart-mikrotik-routeros7-vlan-documentation-and-code-to-manage-vlans-in-routeros-7-via-manu"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "license", "research"]
source: docs/RAG/lot-mikrotik/RouterOS/github-oallart-mikrotik-routeros7-vlan-documentation-and-code-to-manage-vlans-in-routeros-7-via-manu.md
source_anchor: ""
source_lines: [1, 145]
sha256: a2344d320a1e8a85b92e542ee4b36f0d60cc8d5db649f27bf79f78db08b245ca
---

# github-oallart-mikrotik-routeros7-vlan-documentation-and-code-to-manage-vlans-in-routeros-7-via-manu

*This document exists because I never found a reference on the topic in a concise written form, so I took all my research and experience and wrote the missing manual I wish I was given when I started. I hope it helps you on your journey.*

Defining VLANS in RouterOS v7 is notoriously complicated, under-documented, and incredibly easy to do wrong. I have found countless incomplete or wrong examples and tutorials online.

The only proper reference is in the form of a couple youtube videos, especially this one

The configuration of VLANs can be achieved and even automated (see further down) provided a few precautions are taken, and the right order of configuration is respected:

- bridge creation with vlan filtering
- bond (LAG) creation if needed
- untagged ingress configuration
- tagged ingress configuration
- tagged egress configuration

*Note 1:* untagged egress does not need configuration, it's dynamically done by ROS

*Note 2:* the above order also covers hybrid port definition (more on this below)

**It may be useful to think of VLAN configuration on RouterOS 7 as a firewall configuration with filtering defined as egress and ingress rules.**

- trunk: Unlike cisco wants to pretend, a trunk is not a bond (LAG) but a tagged VLAN interface
- access port: an untagged VLAN interface
- hybrid port: an interface both tagged and untagged

It matters *enormously* because if not configured properly, VLAN traffic will be CPU processed, not switch-chip processed, meaning it will be slow and use resources unnecessarily.
Most of the work here ensure we get L2 hardware offload. The first requirement is to have `vlan-filtering=yes` on the bridge.

A RouterOS device very likely has defaults you don't want if you are to use and manage VLANs, especially in an automated way. A new device is also very likely to have outdated software.

Starting clean is the best way in my opinion. There are two main ways to do that in my opinion:

1. reset on boot
2. netinstall (bonus flashing a new ROS version at the same time)

Read below for more on that. I can't guarantee the steps are 100% accurate for your use case.

- Stop the device and power it back on while having the reset button pressed for 12-21 seconds (see manual)
- Connect port 1 of the device directly to your machine
- Use winbox to connect to the device and set your password, skip the license, don't use defaults etc.
- Download the appropriate update for your arch from the mikrotik website and drop it in files via winbox, restart the device
- Dont't forget to run `/system routerboard upgrade`

- Download the appropriate update for your arch from the mikrotik website
- Download Netinstall (CLI Linux) under RouterOS v7, X86
- Poweroff the device and connect it to your machine directly on port 1 or the boot port of your device
- Run `netinstall` by supplying the name of your ethernet device (configure it to be in 192.168.88.0/24) and the update package, IE

```
sudo ./netinstall-cli -e -b -i enp3s0 -a 192.168.88.101 -v routeros-7.19.3-arm64.npk
```
- Power on the device by pressing and holding the reset button for at least 23 seconds to place it in boot mode. `netinstall` should pick it and proceed.
- Dont't forget to run `/system routerboard upgrade`

To enable hardware offloading, the bridge **must** have vlan filtering enabled

```
/interface bridge add name=bridge vlan-filtering=yes
```
If you use link aggregation (LAG), define it now

```
/interface bonding add name=lag1 slaves=sfp-sfplus1,sfp-sfpplus2 mode=802.3ad
```
The untagged interface definition is done when adding a port to the bridge, specifying the PVID that will internally mark the VLAN traffic. Filtering exludes tagged traffic (unless you need a hybrid port, see further).

```
/interface bridge port add bridge=bridge frame-types=admit-only-untagged-and-priority-tagged interface=ether1 pvid=20
```
As mentioned earlier, **there is no need for configuring untagged egress**.

The tagged interface definition is done similarly to the untagged but with a different filter accepting only tagged frames (again unless you need a hybrid port).

```
/interface bridge port add bridge=bridge frame-types=admit-only-vlan-tagged interface=ether5
```
you can also define your trunk LAG by specifying `interface=lag1` here.

This is probably where most of the confusion takes place. No, you **don't need to define VLANs in `/interface/vlan`.** (although it helps with clarity) but you DO have to be careful when creating the entries so that they are grouped by *VLAN*, and not by *interface*.
The configuration is done in the subset `vlan` of `bridge`

```
/interface bridge vlan add bridge=bridge vlan-ids=20 tagged=ether5,lag1
/interface bridge vlan add bridge=bridge vlan-ids=50 tagged=ether5,lag1
/interface bridge vlan add bridge=bridge vlan-ids=100 tagged=ether5,lag1
/interface bridge vlan add bridge=bridge vlan-ids=201 tagged=lag1
/interface bridge vlan add bridge=bridge vlan-ids=202 tagged=lag1
/interface bridge vlan add bridge=bridge vlan-ids=204 tagged=lag1
```
*Note:* grouping by VLAN will be necessary for automating (see further)

A *hybrid* interface is one that is both tagged and untagged at the same time. Therefore it has the `PVID` set to the *untagged* VLAN and has one or more *tagged* VLANs defined.

The first thing to note about a hybrid port is that the *ingress filter* has to accomodate both types.

```
/interface bridge port add bridge=bridge frame-types=admit-all interface=ether10 pvid=204
```
The *egress* configuration is the same as a normal tagged port

Updating requires the annoying id number which is **dynamically printed and unpredictable**, meaning we need to run a `print` command for every operation

Ingress does not need to be updated unless the port is removed Egress is updated by doign the following steps

1. `/interface/bridge/vlan/export` to show the running config
2. `/interface/bridge/vlan/print` to show the dynamic entries and IDs
3. `/interface/bridge/vlan/set tagged=lag1,ether1,ether2,ether10 numbers=3,4,5`

`numbers` are a single or collection of IDs from the `print` command executed before
`tagged` are a single or collection of interface names associated with a given VLAN as per your requirement (add or delete) from the configuration shown in `export`

**Be especially carful to match the ID line to the VLAN config and interfaces you need**

Much like the trunk, we need a few steps

1. `/interface/bridge/port/export` to show the running config
2. `/interface/bridge/port/print where interface=ether5` to show the dynamic entries and IDs
3. `/interface/bridge/port/set pvid=120 numbers=12`

I highly recomend using a single entry in `numbers` to the ID of the matching entry in `print`. Also double check that it matches your target interface against the running config from `export` when setting the `pvid`.

Some useful commands to check for proper operation of VLANs

*Note:* some commands only show data when the interface is **active**. `export` is generally a safer way to see what is actually configured.

Check for **H** *hardware-offload* as flag on the left in the output of

```
/interface/bridge/port/print
```
**Important:** This table is only populated by *active* interfaces, aka those who have the **R** *Running* in `/interface/print`

```
/interface/bridge/vlan/print
```
The easiest way is to add an IP address on the VLAN interface. This requires seting up a VLAN first

```
/interface vlan add interface=bridge name=test vlan-id=800
/interface bridge port add bridge=bridge frame-types=admit-only-untagged-and-priority-tagged interface=ether8 pvid=800
/ip address add address=192.168.1.25/24 interface=test network=192.168.1.0
```
You may need to add routing information if you are trying to reach from a different subnet.

*Note 1:* At this stage, automation of hybrid ports is still a work in progress.

