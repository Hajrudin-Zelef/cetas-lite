---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3-5
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3.md
source_anchor: ""
source_lines: [139, 276]
sha256: 6a126bdacea70b5705c2ab27e7ddc9a3187065884fcd0e86c8b38157acb7bf23
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3

  The following example shows how to configure an ERSPAN destination session:
Device(config)# monitor session 2 type erspan-destination
Device(config-mon-erspan-dst)# destination interface GigabitEthernet1/3/2
Device(config-mon-erspan-dst)# destination interface GigabitEthernet2/2/0
Device(config-mon-erspan-dst)# source
Device(config-mon-erspan-dst-src)# erspan-id 100
Device(config-mon-erspan-dst-src)# ip address 10.1.0.2
The following example shows how to configure a source VRF for an ERSPAN destination session:
Device(config)# monitor session 2 type erspan-destination
Device(config-mon-erspan-dst)# destination interface GigabitEthernet1/3/2
Device(config-mon-erspan-dst)# destination interface GigabitEthernet2/2/0
Device(config-mon-erspan-dst)# source
Device(config-mon-erspan-dst-src)# erspan-id 100
Device(config-mon-erspan-dst-src)# ip address 10.1.0.2
Device(config-mon-erspan-dst-src)# vrf 1 
To verify the ERSPAN configuration, use the following commands:
The following is sample output from the show monitor session command:
Device# show monitor session 53
Session 53
----------
Type                     : ERSPAN Source Session
Status                   : Admin Enabled
Source Ports             : 
MTU                      : Fo1/0/2
The following is sample output from the show platform software monitor session command:
Device# show platform software monitor session 53
Span Session 53 (FED Session 0):
Type: ERSPAN Source
Prev type: Unknown
Ingress Src Ports:
Egress Src Ports: 
Ingress Local Src Ports: (null)
Egress Local Src Ports: (null)
Destination Ports: 
Ingress Src Vlans:
Egress Src Vlans: 
Ingress Up Src Vlans: (null)
Egress Up Src Vlans: (null)
Src Trunk filter Vlans:
RSPAN dst vlan: 0
RSPAN src vlan: 0
RSPAN src vlan sav: 0
Dest port encap = 0x0000
Dest port ingress encap = 0x0000
Dest port ingress vlan = 0x0
SrcSess: 1 DstSess: 0 DstPortCfgd: 0 RspnDstCfg: 0 RspnSrcVld: 0
DstCliCfg: 0 DstPrtInit: 0 PsLclCfgd: 0
Flags: 0x00000000
Remote dest port: 0 Dest port group: 0
FSPAN disabled
FSPAN not notified
ERSPAN Id : 0
ERSPAN Org Ip: 0.0.0.0
ERSPAN Dst Ip: 0.0.0.0
ERSPAN Ip Ttl: 255
ERSPAN DSCP : 0
ERSPAN MTU : 1500 >>>>
ERSPAN VRFID : 0
ERSPAN State : Disabled
ERSPAN Tun id: 61
ERSPAN header-type: 2
ERSPAN SGT : 
The following is sample output from the show monitor session erspan-source detail command:
Device# show monitor session erspan-source detail
Type                     : ERSPAN Source Session
Status                   : Admin Enabled
Description              : -
Source Ports             : 
    RX Only              : None
    TX Only              : None
    Both                 : None
Source Subinterfaces     : 
    RX Only              : None
    TX Only              : None
    Both                 : None
Source VLANs             :
    RX Only              : None
    TX Only              : None
    Both                 : None
Source Drop-cause        : None
Source EFPs              :
    RX Only              : None
    TX Only              : None
    Both                 : None
Source RSPAN VLAN        : None
Destination Ports        : None
Filter VLANs             : None
Filter SGT               : None
Dest RSPAN VLAN          : None
IP Access-group          : None
MAC Access-group         : None
IPv6 Access-group        : None
Filter access-group  :None
smac for wan interface   : None
dmac for wan interface   : None
Destination IP Address   : 192.0.2.1
Destination IPv6 Address : None
Destination IP VRF       : None
MTU                      : 1500
Destination ERSPAN ID    : 251
Origin IP Address        : 10.10.10.216
Origin IPv6 Address      : None
IP QOS PREC              : 0
IPv6 Flow Label          : None
IP TTL                   : 255
ERSPAN header-type       : 3
The following output from the show capability feature monitor erspan-source command displays information about the configured ERSPAN source sessions:
Device# show capability feature monitor erspan-source
ERSPAN Source Session:ERSPAN Source Session Supported: TRUE
No of Rx ERSPAN source session: 8
No of Tx ERSPAN source session: 8
ERSPAN Header Type supported: II and III
ACL filter Supported: TRUE
SGT filter Supported: TRUE
Fragmentation Supported: TRUE
Truncation Supported: FALSE
Sequence number Supported: FALSE
QOS Supported: TRUE
The following output from the show capability feature monitor erspan-destination command displays all the configured global built-in templates:
Device# show capability feature monitor erspan-destination
ERSPAN Destination Session:ERSPAN Destination Session Supported: TRUE
Maximum No of ERSPAN destination session: 8
ERSPAN Header Type supported: II and III
| Standard/RFC | Title | 
|---|---|
| RFC 2784 | Generic Routing Encapsulation (GRE) | 
| Description | Link | 
|---|---|
| The Cisco Support website provides extensive online resources, including documentation and tools for troubleshooting and resolving technical issues with Cisco products and technologies. To receive security and technical information about your products, you can subscribe to various services, such as the Product Alert Tool (accessed from Field Notices), the Cisco Technical Services Newsletter, and Really Simple Syndication (RSS) Feeds. Access to most tools on the Cisco Support website requires a Cisco.com user ID and password. | http://www.cisco.com/support | 
This table provides release and related information for the features explained in this module.
These features are available in all the releases subsequent to the one they were introduced in, unless noted otherwise.
| Release | Feature | Feature Information | 
|---|---|---|
| Cisco IOS XE Everest 16.5.1a | ERSPAN | The Cisco ERSPAN feature allows you to monitor traffic on ports or VLANs, and send the monitored traffic to destination ports. ERSPAN sends traffic to a network analyzer, such as a Switch Probe device or a Remote Monitoring (RMON) probe. | 
| Cisco IOS XE Gibraltar 16.11.1 | ERSPAN -support for Destination Sessions | The ERSPAN destination session defines the session configuration parameters and the ports that receive the monitored traffic. | 
| Cisco IOS XE Amsterdam 17.1.1 | ERSPAN IPv6 | IPv6 support was introduced for ERSPAN. This enables configuration of an IPv6 ERSPAN source and destination session. | 
| Cisco IOS XE Bengaluru 17.5.1 | ERSPAN over MPLS VPN | MPLS VPN support was introduced for ERSPAN. ERSPAN traffic can be transported over an MPLS VPN. Support for ERSPAN is limited to L3VPN IPV4 MPLS. | 
Use the Cisco Feature Navigator to find information about platform and software image support. To access Cisco Feature Navigator, go to https://cfnng.cisco.com/.
