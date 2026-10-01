---
id: collect-261001-cisco/cisco/cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial-2
title: "cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial.md
source_anchor: ""
source_lines: [154, 552]
sha256: addfe85c9ae4a0546e4a24cdc8aabdeb5adb307fbeabda3811254941d9c8746b
---

# cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial

lifetime seconds 86400

**e)** enable crypto configuration on the outside interface:

crypto ikev2 enable outside

**f)** configure remote endpoint for the tunnel and pre-shared key:

tunnel-group 172.16.0.2 type ipsec-l2l

tunnel-group 172.16.0.2 ipsec-attributes

ikev2 remote-authentication pre-shared-key MY_KEY

ikev2 local-authentication pre-shared-key MY_KEY

**g)** define access lists for IPsec tunnel:

access-list OUTSIDE_CRYPTOMAP_10 remark ACL to encrypt traffic from Local to Remote

access-list OUTSIDE_CRYPTOMAP_10 extended permit ip 10.0.0.0 255.255.255.0 10.1.0.0 255.255.255.0

access-list OUTSIDE_CRYPTOMAP_10 extended permit icmp 10.0.0.0 255.255.255.0 10.1.0.0 255.255.255.0

**h)** configure access lists, allowed through this IPsec tunnel, remote peer ip, phase 2  proposal and outgoing interface:

crypto map outside_map 10 match address OUTSIDE_CRYPTOMAP_10

crypto map outside_map 10 set peer 172.16.0.2

crypto map outside_map 10 set ikev2 ipsec-proposal MY_PROPOSAL

crypto map outside_map interface outside

5. Cisco ASA2 configuration

**a)** Interface configuration:

!

interface GigabitEthernet0/0

nameif outside

security-level 0

ip address 172.16.0.2 255.255.255.252

!

interface GigabitEthernet0/1

nameif inside

security-level 100

ip address 10.1.0.254 255.255.255.0

!

**b)** routing configuration:

route outside 0.0.0.0 0.0.0.0 172.16.0.1 1

**c)** IPsec phase 2 configuration:

crypto ipsec ikev2 ipsec-proposal MY_PROPOSAL

protocol esp encryption aes-256

protocol esp integrity sha-1

**d)** IPsec phase 1 configuration:

crypto ikev2 policy 1

encryption aes-256

integrity sha

group 2

prf sha

lifetime seconds 86400

**e)** enable crypto configuration on the outside interface:

crypto ikev2 enable outside

**f)** configure remote endpoint for the tunnel and pre-shared key:

tunnel-group 192.0.2.6 type ipsec-l2l

tunnel-group 192.0.2.6 ipsec-attributes

ikev2 remote-authentication pre-shared-key MY_KEY

ikev2 local-authentication pre-shared-key MY_KEY

**g)** define access lists for IPsec tunnel:


**Note:** compared to ASA1 configuration where I defined networks directly in the access-lists, in ASA2 configuration I defined network objects for remote site and local site. Those objects are configured in cryptomap which allows in the future to simply add/delete networks from the IPsec tunnel definition.

object-group network clients

network-object 10.1.0.0 255.255.255.0

object-group network vpn-remote-site

network-object 10.0.0.0 255.255.255.0

access-list OUTSIDE_CRYPTOMAP_10 remark ACL to encrypt traffic from Local to Remote

access-list OUTSIDE_CRYPTOMAP_10 extended permit ip object-group clients object-group vpn-remote-site

access-list OUTSIDE_CRYPTOMAP_10 extended permit icmp object-group clients object-group vpn-remote-site

**h)** configure access lists, allowed through this IPsec tunnel, remote peer ip, phase 2  proposal and outgoing interface:

crypto map outside_map 10 match address OUTSIDE_CRYPTOMAP_10

crypto map outside_map 10 set peer 192.0.2.6

crypto map outside_map 10 set ikev2 ipsec-proposal MY_PROPOSAL

crypto map outside_map interface outside

Full ASA1, ASA2, R1 configuration:

ASA1# sh run

: Saved

:

: Serial Number: x

: Hardware: ASAv, 2048 MB RAM, CPU Pentium II 2808 MHz

:

ASA Version 9.8(1)

!

hostname ASA1

xlate per-session deny tcp any4 any4

xlate per-session deny tcp any4 any6

xlate per-session deny tcp any6 any4

xlate per-session deny tcp any6 any6

xlate per-session deny udp any4 any4 eq domain

xlate per-session deny udp any4 any6 eq domain

xlate per-session deny udp any6 any4 eq domain

xlate per-session deny udp any6 any6 eq domain

names

!

interface GigabitEthernet0/0

nameif outside

security-level 0

ip address 192.0.2.6 255.255.255.252

!

interface GigabitEthernet0/1

nameif inside

security-level 100

ip address 10.0.0.254 255.255.255.0

!

interface GigabitEthernet0/2

shutdown

no nameif

no security-level

no ip address

!

interface GigabitEthernet0/3

shutdown

no nameif

no security-level

no ip address

!

interface GigabitEthernet0/4

shutdown

no nameif

no security-level

no ip address

!

interface GigabitEthernet0/5

shutdown

no nameif

no security-level

no ip address

!

interface GigabitEthernet0/6

shutdown

no nameif

no security-level

no ip address

!

interface Management0/0

shutdown

no nameif

no security-level

no ip address

!

ftp mode passive

access-list OUTSIDE_CRYPTOMAP_10 remark ACL to encrypt traffic from Local to Remote

access-list OUTSIDE_CRYPTOMAP_10 extended permit ip 10.0.0.0 255.255.255.0 10.1.0.0 255.255.255.0

access-list OUTSIDE_CRYPTOMAP_10 extended permit icmp 10.0.0.0 255.255.255.0 10.1.0.0 255.255.255.0

pager lines 23

logging console debugging

mtu outside 1500

mtu inside 1500

no failover

no monitor-interface service-module

icmp unreachable rate-limit 1 burst-size 1

no asdm history enable

arp timeout 14400

no arp permit-nonconnected

arp rate-limit 8192

route outside 0.0.0.0 0.0.0.0 192.0.2.5 1

timeout xlate 3:00:00

timeout pat-xlate 0:00:30

timeout conn 1:00:00 half-closed 0:10:00 udp 0:02:00 sctp 0:02:00 icmp 0:00:02

timeout sunrpc 0:10:00 h323 0:05:00 h225 1:00:00 mgcp 0:05:00 mgcp-pat 0:05:00

timeout sip 0:30:00 sip_media 0:02:00 sip-invite 0:03:00 sip-disconnect 0:02:00

timeout sip-provisional-media 0:02:00 uauth 0:05:00 absolute

timeout tcp-proxy-reassembly 0:01:00

timeout floating-conn 0:00:00

timeout conn-holddown 0:00:15

timeout igp stale-route 0:01:10

user-identity default-domain LOCAL

aaa authentication login-history

no snmp-server location

no snmp-server contact

crypto ipsec ikev2 ipsec-proposal MY_PROPOSAL

protocol esp encryption aes-256

protocol esp integrity sha-1

crypto ipsec security-association pmtu-aging infinite

crypto map outside_map 10 match address OUTSIDE_CRYPTOMAP_10

crypto map outside_map 10 set peer 172.16.0.2

crypto map outside_map 10 set ikev2 ipsec-proposal MY_PROPOSAL

crypto map outside_map interface outside

crypto ca trustpoint _SmartCallHome_ServerCA

no validation-usage

crl configure

crypto ca trustpool policy

auto-import

crypto ca certificate chain _SmartCallHome_ServerCA

certificate ca xxxxxxxxxxxxxxxxxxxxxxxxxxxxx

xxxxxxxx xxxxxxxx xxxxxxxx 02021018 dad19e26 7de8bb4a 2158cdcc 6b3b4a30

0d06092a 864886f7 0d010105 05003081 ca310b30 09060355 04061302 55533117

30150603 55040a13 0e566572 69536967 6e2c2049 6e632e31 1f301d06 0355040b

13165665 72695369 676e2054 72757374 204e6574 776f726b 313a3038 06035504

0b133128 63292032 30303620 56657269 5369676e 2c20496e 632e202d 20466f72

20617574 686f7269 7a656420 75736520 6f6e6c79 31453043 06035504 03133c56

65726953 69676e20 436c6173 73203320 5075626c 69632050 72696d61 72792043

65727469 66696361 74696f6e 20417574 686f7269 7479202d 20473530 1e170d30

36313130 38303030 3030305a 170d3336 30373136 32333539 35395a30 81ca310b

30090603 55040613 02555331 17301506 0355040a 130e5665 72695369 676e2c20

496e632e 311f301d 06035504 0b131656 65726953 69676e20 54727573 74204e65

74776f72 6b313a30 38060355 040b1331 28632920 32303036 20566572 69536967

6e2c2049 6e632e20 2d20466f 72206175 74686f72 697a6564 20757365 206f6e6c

79314530 43060355 0403133c 56657269 5369676e 20436c61 73732033 20507562

6c696320 5072696d 61727920 43657274 69666963 6174696f 6e204175 74686f72

69747920 2d204735 30820122 300d0609 2a864886 f70d0101 01050003 82010f00

3082010a 02820101 00af2408 08297a35 9e600caa e74b3b4e dc7cbc3c 451cbb2b

e0fe2902 f95708a3 64851527 f5f1adc8 31895d22 e82aaaa6 42b38ff8 b955b7b1

b74bb3fe 8f7e0757 ecef43db 66621561 cf600da4 d8def8e0 c362083d 5413eb49

xxxxxxxx 26e52b8f 1b9febf5 a191c233 49d84363 6a524bd2 8fe87051 4dd18969

7bc770f6 b3dc1274 db7b5d4b 56d396bf 1577a1b0 f4a225f2 af1c9267 18e5f406

04ef90b9 e400e4dd 3ab519ff 02baf43c eee08beb 378becf4 d7acf2f6 f03dafdd

xxxxxxxx 1d1c40cb 74241921 93d914fe ac2a52c7 8fd50449 e48d6347 883c6983

cbfe47bd 2b7e4fc5 95ae0e9d d4d143c0 6773e314 087ee53f 9f73b833 0acf5d3f

