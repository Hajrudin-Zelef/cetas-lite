---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-25
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-25.md
source_anchor: ""
source_lines: [1, 43]
sha256: 196b781ff2d24401ea147a2197c5d63a1cc3891f4ef12790c6758a77d33bfbc7
---

# Introduction

The MikroTik RouterOS supports Universal Plug and Play architecture for transparent peer-to-peer network connectivity of personal computers and network-enabled intelligent devices or appliances.

UPnP enables data communication between any two devices under the command of any control device on the network. Universal Plug and Play is completely independent of any particular physical medium. It supports networking with automatic discovery without any initial configuration, whereby a device can dynamically join a network. DHCP and DNS servers are optional and will be used if available on the network. UPnP implements a simple yet powerful NAT traversal solution, that enables the client to get full two-way peer-to-peer network support from behind the NAT.

There are two interface types for UPnP: **internal** (the one local clients are connected to) and **external** (the one the Internet is connected to). **A router may only have one active external interface with a 'public' IP address on it**, and as many internal interfaces as needed, all with source-NATted 'internal' IP addresses. The protocol works by creating dynamic NAT entries.

UPnP **internal** interface can create NAT mapping for any subnet, not just the subnet present on the internal interface, so caution must be used when setting **internal** interfaces.

The UPnP protocol is used for many modern applications, like most DirectX games, as well as for various Windows Messenger features like remote assistance, application sharing, file transfer, voice, and video from behind a firewall.

# Configuration

## General properties

| Property | Description | 
|---|---|
| **allow-disable-external-interface** (*yes \| no*  ; Default:**yes** ) | whether or not the users are allowed to disable the router's external interface. This functionality (for users to be able to turn the router's external interface off without any authentication procedure) is required by the standard, but as it is sometimes not expected or unwanted in UPnP deployments which the standard was not designed for (it was designed mostly for home users to establish their own local networks), you can disable this behavior | 
| **enabled** (*yes \| no*  ; Default:**no** ) | Enable UPnP service | 
| **show-dummy-rule** (*yes \| no*  ; Default:**yes** ) | Enable a workaround for some broken implementations, which are handling the absence of UPnP rules incorrectly (for example, popping up error messages). This option will instruct the server to install a dummy (meaningless) UPnP rule that can be observed by the clients, which refuse to work correctly otherwise | 

If you do not disable the **allow-disable-external-interface**, any user from the local network will be able (without any authentication procedures) to disable the router's external interface

## UPnP Interfaces

| Property | Description | 
|---|---|
| **interface** (*string* ; Default: ) | Interface name on which uPnP will be running | 
| **type** (*external \| internal* ; Default:**no** ) | UPnP interface type:  | 
| **forced-external-ip** (*Ip* ; Default: ) | Allow specifying what public IP to use if the external interface has more than one IP available. | 

In more complex setups with VLANs, where the VLAN interface is considered as the LAN interface, the VLAN interface itself should be specified as the internal interface for UPnP to work properly.

# Configuration Example

We have masquerading already enabled on our router:

To enable the UPnP feature:

Now, all we have to do is to add interfaces:

Now once the client from the internal interface side sends UPnP request, dynamic NAT rules will be created on the router, example rules could look something similar to these:
