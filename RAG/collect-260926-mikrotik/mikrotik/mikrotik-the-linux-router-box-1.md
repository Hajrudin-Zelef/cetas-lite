---
id: collect-260926-mikrotik/mikrotik/mikrotik-the-linux-router-box-1
title: "Mikrotik - The Linux Router Box"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "licenses"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mikrotik-the-linux-router-box.md
source_anchor: ""
source_lines: [1, 167]
sha256: d9e617c88a5549765b3d63f2f25a8d5f50284b7e26b72c413d4e7c44e9f81e0c
---

# Mikrotik - The Linux Router Box

## Abstract

I'm sharing about Mikrotik, the vendor of the routers we use as routers for our servers at our new data center location. I describe some basics about Mikrotik and my experiences made.

## Decision Making

Not long ago, we had to move with all of our servers to a new data center. As our routers in place were not a really good solution and quite old, we decided to get something new.

The typical first choices would have been Cisco, Juniper, Hewlett Packard or any of the big players. But they are all rather costly. After talking to some routing specialists we decided to take the advice and go for Mikrotik. They are a lot cheaper and are favoured by the professionals we asked.

## Summary first

I decided to like Mikrotik. There have been some small quirks, and unintuitive behaviour in few places, but I got around them rather quickly. Being a Mikrotik-Beginner, I would say that many things can be improved. For example the backup and recovery process within HP Switches seem more convenient to me. But overall it does what it should and it's a pleasant user experience for me. The handling is easy for me even thought the routing feature set seems huge to me. I like it that you can remote execute any command via SSH. I also like that you can have some very little cheap boxes and they have the same feature set as the big router boxes. So you can try all scenarios with the tiny boxes in the lab and when it works deploy them to the big boxes. (Lab devices are the **hap lite** and the production devices are **CCR1036-8G-2S+**)

## What's Mikrotik?

Mikrotik is a latvian vendor of devices called **Mikrotik Router Board** with an own Linux based OS called **RouterOS**. The Devices are available from tiny to enterprise. You can even install RouterOS on an ordinary PC(You need a license for that. But the licenses are dead cheap). The devices are made to handle all what can be possibly done in the routing business, like:

- Routing Protokolls (rip, ospf, bgp, mpls, ...)
- firewalling
- IP (ipv4/v6/dhcp/ntp)
- failover(vrrp)
- secure networking(ipsec, openvpn, wireguard will be shortly available)
- authentication(radius,ppp,l2tp)

In comparison to an ordinary linux/unix system, mikrotik is a specialised os, focused on the area of routing. So you cannot just install any package and run it. It's a given feature set of application and configuration that you can make use of.

## Management

### Possible Configuration Methods

Mikrotik has 3 options to manage its devices:

- Webinterface (See demo here: Demo1: RouterOS Demo2: RouterOS)
- Winbox: A Windows Tool
- CLI via SSH

### Configuration Tree

In comparison to a linux system, where you install a package, edit its configuration files and start and stop sevices, mikrotik provides one configuration tree, which you use to configure every aspect of the operation. This configuration tree is accessed by any of the configuration methods mentioned above. Every leaf of the configuration tree has its own set of commands to configure the current context and some generic commands(like export) that work everywhere.

**Reduced Tree for demonstration**

```
+-- interface
|   +-- bonding
|   +-- bridge
|   +-- vrrp
+-- ip
|   +-- address
|   +-- dhcp-client
|   +-- dhcp-server
|   +-- firewall
+-- routing
|   +-- bgp
|   +-- osp
|   +-- rip
+-- system
    +-- clock
    +-- environment
    +-- script
```

### Documentation

There's a wiki and a forum in the web, where you find in detail explanation of everything and configuration examples.

## The Console

The console is available via all of the mentioned configuration methods. With ssh it's the primary config option. From the web interface and winbox it's available as additional tool.

### Safe Mode

As a router is often a very critical point of the infrastructure, one needs to be very careful when configuring it while online.

As every configuration change gets active with the hit of the [enter] key, it may happen that a configuration change may render a device unusable.

That's where the safe mode comes in. When you're at the console just hit [ctrl] + 

This concerns mainly network sessions, because the danger of getting disconnected while within a serial connection is less likely.

### Using the CLI

In the console you can either give complete commands absolute paths like...

```
/ip address add address=192.168.1.10 .....
```

or you can move around like you do with `cd` in unix.

```
ip
address
..
/
/ip address
```

and give a command for the current context

```
/ip address
add address=192.168.1.10 .....
```

#### Using export

You can use the `export` command at any level to print out the configuration of the configuration subtree of the current level, for example to review the current config before changing. If you're at the root level the complete configuration will be printed. The output can be directly copy and pasted to the cli at another router to repeat the configuration.

### Console Help

You can always enter **?** to get a short help of the current command you started to type or at the current level to get a short information of the available commands within the current context / configuration path.

Tab completion is available to show possible arguments to a current typed command, available commands at the current level and possible values for the currently typed in argument.

### Backup and Restore

There are two different options you have when it comes to to backup: Configuration Exports and Backups.

#### Export & Reset

Export prints the configuration or parts of it without security relevant information(passwords, ssh keys and certificates). It can be used to restore the current configuration when you did something wrong and do not know exactly what configuration setting it was.

I save exports daily via ssh on a linux box as text file.

To revert to a saved export you have to do the following:

- 
export the file to your linux box: `ssh mikrotik-ip /export >config.rsc`
- 
important: insert `/delay delay-time=15s` at line 1. (There's a bug that the config is not loaded when the initial pause is not there)
- 
reupload the config to the mikrotik: `scp config.rsc mikrotik:.`
- 
reset and reboot the router

`/system reset-configuration run-after-reset=config.rsc`

#### Backup & Restore

In opposite to the export a backup is not a configuration text file, but a binary blob holding every data of the current operation, even passwords, keys,... A backup can only be restored to the same hardware model.

To create a backup of the current machine, use this command within the cli:

`/system backup save name=current`

Then a file named current.backup will be saved on the Mikrotiks flash storage. You may download it via scp to another machine.

To restore the same backup, use this command:

`/system backup load name=current`

### Netwatch

Netwatch is a pingtool that pings some host and if that host gets up or down, you can fire up a script.

### Scripting

Mikrotik has its own Scripting language. You can execute them from the netwatch tool or from other events that special commands provide. E. g. if you have VRRP(IP-Failover) active you can set scripts on interfaces becoming master or backup.

I used the script feature on the backup device with VRRP. My scenario is to have VRRP on the external(ISP) side and on the internal side(LAN). So when one side goes up at the backup the other side must follow. The script checks the connections and lets the other VRRP-Connections follow to master role when the links are really up(used netwatch for that).

### Monitoring & API

