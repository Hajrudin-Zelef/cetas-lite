---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-5
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [706, 848]
sha256: e1cd9e4e6e99c280c4f247fd952a9f1349083e003f8ef60928dd2cf8594e9ad7
---

# Configure DeviceA to generate a local key pair.

In Figure 3-110, two HWTACACS servers are deployed. The enterprise requires that the two HWTACACS servers back up each other to ensure user authentication if one server is faulty.
Ensure that the shared key in the HWTACACS server template is the same as that configured on the HWTACACS server.
If the HWTACACS server does not support the user name containing the domain name, run the undo hwtacacs-server user-name domain-included command in the HWTACACS server template view to configure the device to send packets that do not contain the domain name to the HWTACACS server.
# Create an HWTACACS server template named ht.
[DeviceA] hwtacacs-server template ht
# Configure IP addresses and port numbers for the primary HWTACACS authentication, authorization, and accounting servers.
[DeviceA-hwtacacs-ht] hwtacacs-server authentication 10.7.66.66 49
[DeviceA-hwtacacs-ht] hwtacacs-server authorization 10.7.66.66 49
[DeviceA-hwtacacs-ht] hwtacacs-server accounting 10.7.66.66 49
# Configure IP addresses and port numbers for the secondary HWTACACS authentication, authorization, and accounting servers.
[DeviceA-hwtacacs-ht] hwtacacs-server authentication 10.7.66.67 49 secondary
[DeviceA-hwtacacs-ht] hwtacacs-server authorization 10.7.66.67 49 secondary
[DeviceA-hwtacacs-ht] hwtacacs-server accounting 10.7.66.67 49 secondary
# Configure a shared key for the HWTACACS server.
[DeviceA-hwtacacs-ht] hwtacacs-server shared-key cipher YsHsjx_202206139
[DeviceA-hwtacacs-ht] quit
# Create an authentication scheme named l-h, and set the authentication mode to HWTACACS and local authentication.
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme l-h
[DeviceA-aaa-authen-l-h] authentication-mode hwtacacs local
[DeviceA-aaa-authen-l-h] quit
# Create an authorization scheme named hwtacacs, and set the authorization mode to HWTACACS and local authorization.
[DeviceA-aaa] authorization-scheme hwtacacs
[DeviceA-aaa-author-hwtacacs] authorization-mode hwtacacs local
[DeviceA-aaa-author-hwtacacs] quit
# Create an accounting scheme named hwtacacs, and set the accounting mode to HWTACACS accounting. Configure a policy for the device to keep users online upon accounting-start failures.
[DeviceA-aaa] accounting-scheme hwtacacs
[DeviceA-aaa-accounting-hwtacacs] accounting-mode hwtacacs 
[DeviceA-aaa-accounting-hwtacacs] accounting start-fail online
# Set the real-time accounting interval to 3 minutes.
[DeviceA-aaa-accounting-hwtacacs] accounting realtime 3
[DeviceA-aaa-accounting-hwtacacs] quit
[DeviceA-aaa] domain huawei
[DeviceA-aaa-domain-huawei] authentication-scheme l-h
[DeviceA-aaa-domain-huawei] authorization-scheme hwtacacs
[DeviceA-aaa-domain-huawei] accounting-scheme hwtacacs
[DeviceA-aaa-domain-huawei] hwtacacs-server ht
[DeviceA-aaa-domain-huawei] quit
[DeviceA-aaa] quit
Run the display hwtacacs-server template command on DeviceA. The command output shows that the HWTACACS server template configuration meets the requirements.
[DeviceA] display hwtacacs-server template name ht
  ---------------------------------------------------------------------------
  HWTACACS-server template name        : ht
  HWTACACS-server template index       : 1
  Primary-authentication-server        : 10.7.66.66:49 Vrf:- Shared-key:- Status:UP
  Primary-authentication-ipv6-server   : -:0 Vrf:- Shared-key:- Status:-
  Primary-authorization-server         : 10.7.66.66:49 Vrf:- Shared-key:- Status:UP
  Primary-authorization-ipv6-server    : -:0 Vrf:- Shared-key:- Status:-
  Primary-accounting-server            : 10.7.66.66:49 Vrf:- Shared-key:- Status:UP
  Primary-accounting-ipv6-server       : -:0 Vrf:- Shared-key:- Status:-
  Secondary-authentication-server      : 10.7.66.66:49 Vrf:- Shared-key:- Status:-
  Secondary-authentication-ipv6-server : -:0 Vrf:- Shared-key:- Status:-
  Secondary-authorization-server       : 10.7.66.66:49 Vrf:- Shared-key:- Status:-
  Secondary-authorization-ipv6-server  : -:0 Vrf:- Shared-key:- Status:-
  Secondary-accounting-server          : 10.7.66.66:49 Vrf:- Shared-key:- Status:-
  Secondary-accounting-ipv6-server     : -:0 Vrf:- Shared-key:- Status:-
  Third-authentication-server          : -:0 Vrf:- Shared-key:- Status:-
  Third-authentication-ipv6-server     : -:0 Vrf:- Shared-key:- Status:-
  Third-authorization-server           : -:0 Vrf:- Shared-key:- Status:-
  Third-authorization-ipv6-server      : -:0 Vrf:- Shared-key:- Status:-
  Third-accounting-server              : -:0 Vrf:- Shared-key:- Status:-
  Third-accounting-ipv6-server         : -:0 Vrf:- Shared-key:- Status:-
  Fourth-authentication-server         : -:0 Vrf:- Shared-key:- Status:-
  Fourth-authentication-ipv6-server    : -:0 Vrf:- Shared-key:- Status:-
  Fourth-authorization-server          : -:0 Vrf:- Shared-key:- Status:-
  Fourth-authorization-ipv6-server     : -:0 Vrf:- Shared-key:- Status:-
  Fourth-accounting-server             : -:0 Vrf:- Shared-key:- Status:-
  Fourth-accounting-ipv6-server        : -:0 Vrf:- Shared-key:- Status:-
  Current-authentication-server        : 10.7.66.66:49 Vrf:- Shared-key:- Status:UP
  Current-authentication-ipv6-server   : -:0 Vrf:- Shared-key:- Status:-
  Current-authorization-server         : 10.7.66.66:49 Vrf:- Shared-key:- Status:UP
  Current-authorization-ipv6-server    : -:0 Vrf:- Shared-key:- Status:-
  Current-accounting-server            : 10.7.66.66:49 Vrf:- Shared-key:- Status:UP
  Current-accounting-ipv6-server       : -:0 Vrf:- Shared-key:- Status:-
  Source-IP-address                    : -
  Source-LoopBack                      : -
  Source-Vlanif                        : -
  Source-IPv6-address                  : -
  IPv6 Source-LoopBack                 : -
  IPv6 Source-Vlanif                   : -
  Shared-key                           : ****************
  Quiet-interval(min)                  : 5
  Response-timeout-Interval(sec)       : 5
  Domain-included                      : Original
  Traffic-unit                         : B
  User name in authen-start message    : No
  ---------------------------------------------------------------------------
#
sysname DeviceA
#
hwtacacs-server template ht
 hwtacacs-server authentication 10.7.66.66
 hwtacacs-server authentication 10.7.66.67 secondary
 hwtacacs-server authorization 10.7.66.66
 hwtacacs-server authorization 10.7.66.67 secondary
 hwtacacs-server accounting 10.7.66.66
 hwtacacs-server accounting 10.7.66.67 secondary
 hwtacacs-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs."u,S-6a-X1'[X=L"cpF!5Oz`1!!!!!2jp5!!!!!!A!!!!Ix>cM8i{y6!);(8Dr9:dK`&BHfE(H2=.:SH{@pT%+%#
#
aaa
 authentication-scheme l-h
  authentication-mode hwtacacs local
 authorization-scheme hwtacacs
  authorization-mode hwtacacs local
 accounting-scheme hwtacacs
  accounting-mode hwtacacs
  accounting realtime 3
 domain huawei
  authentication-scheme l-h
  accounting-scheme hwtacacs
  authorization-scheme hwtacacs
  hwtacacs-server ht
 local-user user1-huawei password irreversible-cipher $1d$OwseVRh@LH}ZeTBm$1nH4$ab>d(N{-%0!ab48y=Ic*xEUR4pVhR2"9-~,$
 local-user user1-huawei privilege level 3
 local-user user1-huawei service-type ssh   
#
domain huawei admin
#
vlan batch 10 20 30
#
interface Vlanif10
 ip address 10.1.1.2 255.255.255.0
#
interface Vlanif20
 ip address 10.1.6.2 255.255.255.0
#
interface Vlanif30
 ip address 10.1.5.2 255.255.255.0
#
interface 10GE1/0/1 
 port link-type trunk
 port trunk allow-pass vlan 10
#
interface 10GE1/0/2 
 port link-type trunk
 port trunk allow-pass vlan 20
#
interface 10GE1/0/3 
 port link-type trunk
 port trunk allow-pass vlan 30
#
return 
As shown in Figure 3-111, a RADIUS server is deployed on the network. To ensure the confidentiality and integrity of data transmitted between the RADIUS server and client, the administrator configures DTLS for data transmission between the server and client.
