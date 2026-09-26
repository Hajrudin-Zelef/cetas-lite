---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-26-5
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-26.md
source_anchor: ""
source_lines: [345, 439]
sha256: dc1304deec2f813a7e7204431b53e4792730380d932678129b7d50190b2779c8
---

# Overview

| **security** (*string* ; Default:**none** ) | Name of security configuration from **/caps-man security** | 
| **security.authentication-types** (*list of string* ; Default:**none** ) | Specify the type of Authentication from **wpa-psk** ,**wpa2-psk** ,**wpa-eap** or**wpa2-eap** | 
| **security.disable-pmkid** (; Default: ) |  | 
| **security.eap-methods** (*eap-tls \| passthrough* ; Default:**none** ) |   | 
| **security.eap-radius-accounting** (; Default: ) |  | 
| **security.encryption** (*aes-ccm \| tkip* ; Default:*aes-ccm* ) | Set type of unicast encryption algorithm used | 
| **security.group-encryption** (*aes-ccm \| tkip* ; Default:**aes-ccm** ) | Access Point advertises one of these ciphers, multiple values can be selected. Access Point uses it to encrypt all broadcast and multicast frames. Client attempts connection only to Access Points that use one of the specified group ciphers.  | 
| **security.group-key-update** (*time: 30s..1h* ; Default:**5m** ) | Controls how often Access Point updates the group key. This key is used to encrypt all broadcast and multicast frames. property only has effect for Access Points. | 
| **security.passphrase** (*string* ; Default: ) | WPA or WPA2 pre-shared key | 
| **security.tls-certificate** (*none \| name* ; Default: ) | Access Point always needs a certificate when configured when **security.tls-mode** is set to*verify-certificate* , or is set to*dont-verify-certificate* . | 
| **security.tls-mode** (*verify-certificate \| dont-verify-certificate \| no-certificates* ; Default: ) | This property has effect only when **security.eap-methods** contains*eap-tls* . | 
| **ssid** (*string (0..32 chars)* ; Default: ) | SSID (service set identifier) is a name broadcast in the beacons that identifies wireless network. | 
| **tx-chains** (*list of integer [0..2]* ; Default:**0** ) | Which antennas to use for transmit. | 

## **Channel Groups**

Channel group settings allows for the configuration of lists of radio channel related settings, such as radio band, frequency, Tx Power extension channel and width.

Channel group settings are configured in the Channels profile menu **/caps-man channels**

| Property | Description | 
|---|---|
| **band** (*2ghz-b \| 2ghz-b/g \| 2ghz-b/g/n \| 2ghz-onlyg \| 2ghz-onlyn \| 5ghz-a \| 5ghz-a/n \| 5ghz-onlyn* ; Default: ) | Define operational radio frequency band and mode taken from hardware capability of wireless card | 
| **comment** (*string* ; Default: ) | Short description of the Channel Group profile | 
| **extension-channel** (*Ce \| Ceee \| eC \| eCee \| eeCe \| eeeC \| disabled* ; Default: ) | Extension channel configuration. (E.g. Ce = extension channel is above Control channel, eC = extension channel is below Control channel) | 
| **frequency** (*integer [0..4294967295]* ; Default: ) | Channel frequency value in MHz on which AP will operate. | 
| **name** (*string* ; Default: ) | Descriptive name for the Channel Group Profile | 
| **tx-power** (*integer [-30..40]* ; Default: ) | TX Power for CAP interface (for the whole interface not for individual chains) in dBm. It is not possible to set higher than allowed by country regulations or interface. By default max allowed by country or interface is used. | 
| **width** (; Default: ) | Sets Channel Width in MHz. (E.g. 20, 40) | 
| **save-selected** (; Default:**yes** ) | Saves selected channel for the CAP Radio - will select this channel after the CAP reconnects to CAPsMAN and use it till the channel Re-optimize is done for this CAP. | 

## **Datapath Configuration**

Datapath settings control data forwarding related aspects. On CAPsMAN datapath settings are configured in datapath profile menu **/caps-man datapath** or directly in a configuration profile or interface menu as settings with **datapath.** prefix.

There are 2 major forwarding modes:

- local forwarding mode, where CAP is locally forwarding data to and from wireless interface
- manager forwarding mode, where CAP sends to CAPsMAN all data received over wireless and only sends out the wireless data received from CAPsMAN. In this mode even client-to-client forwarding is controlled and performed by CAPsMAN.

Forwarding mode is configured on a per-interface basis - so if one CAP provides 2 radio interfaces, one can be configured to operate in local forwarding mode and the other in manager forwarding mode. The same applies to Virtual-AP interfaces - each can have different forwarding mode from master interface or other Virtual-AP interfaces.

Most of the datapath settings are used only when in manager forwarding mode, because in local forwarding mode CAPsMAN does not have control over data forwarding.

There are the following datapath settings:

- bridge -- bridge interface to add interface to, as a bridge port, when enabled
- bridge-cost -- bridge port cost to use when adding as bridge port
- bridge-horizon -- bridge horizon to use when adding as bridge port
- client-to-client-forwarding -- controls if client-to-client forwarding between wireless clients connected to interface should be allowed, in local forwarding mode this function is performed by CAP, otherwise it is performed by CAPsMAN.
- local-forwarding -- controls forwarding mode
- openflow-switch -- OpenFlow switch to add interface to, as port when enabled
- vlan-id -- VLAN ID to assign to interface if vlan-mode enables use of VLAN tagging
- vlan-mode -- VLAN tagging mode specifies if VLAN tag should be assigned to interface (causes all received data to get tagged with VLAN tag and allows interface to only send out data tagged with given tag)

## **Local Forwarding Mode**

In this mode wireless interface on CAP behaves as a normal interface and takes part in normal data forwarding. Wireless interface will accept/pass data to networking stack on CAP. CAPsMAN will not participate in data forwarding and will not process any of data frames, it will only control interface configuration and client association process.

Wireless interface on CAP will change its configuration to 'enabled' and its state and some relevant parameters (e.g. mac-address, arp, mtu) will reflect that of the interface on CAPsMAN. Note that wireless related configuration **will not** reflect actual interface configuration as applied by CAPsMAN:

```
[admin@CAP] /interface wireless> pr
Flags: X - disabled, R - running 
 0  R ;;; managed by CAPsMAN
      ;;; channel: 5180/20-Ceee/ac, SSID: master, local forwarding
      name="wlan2" mtu=1500 mac-address=00:03:7F:48:CC:07 arp=enabled 
      interface-type=Atheros AR9888 mode=ap-bridge ssid="merlin" 
      frequency=5240 band=5ghz-a/n channel-width=20/40mhz-eC scan-list=default
      ...
```
Virtual-AP interfaces in local forwarding mode will appear as enabled and dynamic Virtual-AP interfaces:

```
[admin@CAP] /interface> pr
Flags: D - dynamic, X - disabled, R - running, S - slave 
 #     NAME                                TYPE         MTU L2MTU  MAX-L2MTU
 ...
 2  RS ;;; managed by CAPsMAN
       ;;; channel: 5180/20-Ceee/ac, SSID: master, local forwarding
       wlan2                               wlan        1500  1600
 3 DRS ;;; managed by CAPsMAN
       ;;; SSID: slave, local forwarding
       wlan6                               wlan        1500  1600
 ...
[admin@CAP] /interface> wireless pr   
Flags: X - disabled, R - running 
 ...
 2  R ;;; managed by CAPsMAN
      ;;; SSID: slave, local forwarding
      name="wlan6" mtu=1500 mac-address=00:00:00:00:00:00 arp=enabled 
      interface-type=virtual-AP master-interface=wlan2 
```
The fact that Virtual-AP interfaces are added as dynamic, somewhat limits static configuration possibilities on CAP for data forwarding, such as assigning addresses to Virtual-AP interface. This does not apply to master wireless interface.

