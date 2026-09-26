---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-26-6
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-26.md
source_anchor: ""
source_lines: [440, 569]
sha256: 83f17b93430bae8f7276030f7652799c7604612cd3eea830864bd289d4464c63
---

# Overview

To overcome this it is possible to use the static-virtual setting on the CAP which will create Static Virtual Interfaces instead of Dynamic and allows the possibility to assign IP configuration to those interfaces. MAC address is used to remember each static-interface when applying the configuration from the CAPsMAN. If two or more static interfaces will have the same MAC address the configuration could be applied in random order.

To facilitate data forwarding configuration, CAP can be configured with bridge to which interfaces are automatically added as ports when interfaces are enabled by CAPsMAN. This can be done in **/interface wireless cap** menu.

## **Manager Forwarding Mode**

In this mode CAP sends all data received over wireless to CAPsMAN and only sends out over wireless, data received from CAPsMAN. CAPsMAN has full control over data forwarding including client-to-client forwarding. Wireless interface on CAP is disabled and does not participate in networking:

```
 ...
 1 X  ;;; managed by CAPsMAN
      ;;; channel: 5180/20-Ceee/ac, SSID: master, manager forwarding
      name="wlan2" mtu=1500 mac-address=00:03:7F:48:CC:07 arp=enabled 
      interface-type=Atheros AR9888 mode=ap-bridge ssid="merlin" 
 ...
```
Virtual-AP interfaces are also created as 'disabled' and do not take part in data forwarding on CAP.

## Access List

Access list on CAPsMAN is an ordered list of rules that is used to allow/deny clients to connect to any CAP under CAPsMAN control. When client attempts to connect to a CAP that is controlled by CAPsMAN, CAP forwards that request to CAPsMAN. As a part of registration process, CAPsMAN consults access list to determine if client should be allowed to connect. The default behaviour of the access list is to allow connection.

Access list rules are processed one by one until matching rule is found. Then the action in the matching rule is executed. If action specifies that client should be accepted, client is accepted, potentially overriding it's default connection parameters with ones specified in access list rule.

Access list is configured in **/caps-man access-list** menu. There are the following parameters for access list rules:

- client matching parameters:
  - address - MAC address of client (or, if mask is specified, only those parts will be checked as per the mask, so to match vendor D8 from "D8:1C:79:6E:1E:FE", simply enter a bogus entry, such as "D8:00:00:00:00" and then use the mask as per next line)
  - mask - MAC address mask to apply when comparing client address. For example, use FF:00:00:00:00:00 to match only the first octet of the specified MAC address. In above example, regardless of entered MAC, it will match only first octet. Similarly, entering 00:00:00:00:FF will only match the last octet (FE) of a hypotetical MAC "D8:1C:79:6E:1E:FE"). So in the mac line, you could just enter 00:00:00:00:00:FE, if you would use such a mask.
  - interface - optional interface to compare with interface to which client actually connects to
  - time - time of day and days when rule matches
  - signal-range - range in which client signal must fit for rule to match
- action parameter - specifies action to take when client matches:
  - accept - accept client
  - reject - reject client
  - query-radius - query RADIUS server if particular client is allowed to connect
- connection parameters:
  - ap-tx-limit - tx speed limit in direction to client
  - client-tx-limit - tx speed limit in direction to AP (applies to RouterOS clients only)
  - client-to-client-forwarding - specifies whether to allow forwarding data received from this client to other clients connected to the same interface
  - private-passphrase - PSK passphrase to use for this client if some PSK authentication algorithm is used
  - radius-accounting - specifies if RADIUS traffic accounting should be used if RADIUS authentication gets done for this client
  - vlan-mode - VLAN tagging mode specifies if traffic coming from client should get tagged (and untagged when going to client).
  - vlan-id - VLAN ID to use if doing VLAN tagging.

## Registration Table

Registration table contains a list of clients that are connected to radios controlled by CAPsMAN and is available in **/caps-man registration-table** menu:

[admin@CM] /caps-man> registration-table print
 # INTERFACE                   MAC-ADDRESS       UPTIME                RX-SIGNAL
 0 cap1                        00:03:7F:48:CC:0B 1h38m9s210ms                -36

# **Examples**

## **Basic configuration with master and slave interface**

Create security profile for WPA2 PSK, without specifying passphrase:

[admin@CM] /caps-man security>add name="wpa2psk" authentication-types=wpa2-psk encryption=aes-ccm

Create configuration profile to be used by master interface

- specify WPA2 passphrase in configuration
- specify channel settings in configuration:

[admin@CM] /caps-man configuration> add name=master-cfg ssid=master security=wpa2psk
security.passphrase=12345678 channel.frequency=5180 channel.width=20 channel.band=5ghz-a

Create configuration profile to be used by virtual AP interface

- specify different WPA2 passphrase in configuration:

[admin@CM] /caps-man configuration> add name=slave-cfg ssid=slave security=wpa2psk
security.passphrase=87654321

Create provisioning rule that matches any radio and creates dynamic interfaces using master-cfg and slave-cfg:

[admin@CM] /caps-man provisioning> add action=create-dynamic-enabled master-configuration=master-cfg
slave-configurations=slave-cfg

Now when AP connects and is provisioned 2 dynamic interfaces (one master and one slave) will get created:

```
[admin@CM] /caps-man interface> print detail 
Flags: M - master, D - dynamic, B - bound, X - disabled, I - inactive, R - running 
 0 MDB  name="cap1" mtu=1500 l2mtu=2300 radio-mac=00:0C:42:1B:4E:F5 master-interface=none 
        configuration=master-cfg 
 1  DB  name="cap2" mtu=1500 l2mtu=2300 radio-mac=00:00:00:00:00:00 master-interface=cap1 
        configuration=slave-cfg 
```
Consider an AP, that does not support configured frequency connects and can not become operational:

```
[admin@CM] /caps-man interface> pr
Flags: M - master, D - dynamic, B - bound, X - disabled, I - inactive, R - running 
 #      NAME                                 RADIO-MAC         MASTER-INTERFACE                               
 0 MDB  ;;; unsupported band or channel
        cap3                                 00:0C:42:1B:4E:FF none    
 ...
```
We can override channel settings for this particular radio in interface settings, without affecting master-cfg profile:

[admin@CM] /caps-man interface> set cap3 channel.frequency=2142 channel.band=2ghz-b/g

Allow Specific MAC address range to match the Access-list, for example, match all the Apple devices:

[admin@CM] /caps-man access-list> add mac-address=18:34:51:00:00:00 mac-address-mask=FF:FF:FF:00:00:00 action=accept

Configuring DHCP Server Option 138 for setting the CAPsMAN address on the CAP boards

[admin@CM] /ip dhcp-server network set <network-id> caps-manager=<capsman-server-ip>

DHCP client this CAPsMAN IP will see in "/ip dhcp-client print detail"

## **Configuration with certificates**

You would want to configure certificates in your CAPsMAN to use options as *Require Peer Certificate* and *Lock To Caps Man*. These options increase security and in some cases stability of your CAPsMAN network. CAPs won't connect to CAPsMAN without a specific certificate and vice versa.

### Fast and easy configuration

This is a basic configuration for using certificates in your CAPsMAN setup. This example assumes that you already have basic configuration on your CAPsMAN and CAP. It is best to use this configuration in CAPsMAN networks which are not constantly growing. For more details read about CAP to CAPsMAN Connection.

**CAPsMAN device:**

In CAPsMAN Manager menu set *Certificate* and *CA Certificate* to *auto*:

/caps-man manager
set ca-certificate=auto certificate=auto

