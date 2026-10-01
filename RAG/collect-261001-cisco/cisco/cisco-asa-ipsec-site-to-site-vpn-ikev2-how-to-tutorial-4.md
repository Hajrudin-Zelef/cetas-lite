---
id: collect-261001-cisco/cisco/cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial-4
title: "cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/collect-261001-cisco/cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial.md
source_anchor: ""
source_lines: [1011, 1536]
sha256: 1a9e1c757b33848225e1431582916c7843dd955d4c75d135ce04e8325ff8239c
---

# cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial

ca595485 26e52b8f 1b9febf5 a191c233 49d84363 6a524bd2 8fe87051 4dd18969

7bc770f6 b3dc1274 db7b5d4b 56d396bf 1577a1b0 f4a225f2 af1c9267 18e5f406

04ef90b9 e400e4dd 3ab519ff 02baf43c eee08beb 378becf4 d7acf2f6 f03dafdd

75913319 1d1c40cb 74241921 93d914fe ac2a52c7 8fd50449 e48d6347 883c6983

cbfe47bd 2b7e4fc5 95ae0e9d d4d143c0 6773e314 087ee53f 9f73b833 0acf5d3f

3487968a ee53e825 15020301 0001a381 b23081af 300f0603 551d1301 01ff0405

30030101 ff300e06 03551d0f 0101ff04 04030201 06306d06 082b0601 05050701

0c046130 5fa15da0 5b305930 57305516 09696d61 67652f67 69663021 301f3007

06052b0e 03021a04 148fe5d3 1a86ac8d 8e6bc3cf 806ad448 182c7b19 2e302516

23687474 703a2f2f 6c6f676f 2e766572 69736967 6e2e636f 6d2f7673 6c6f676f

2e676966 301d0603 551d0e04 1604147f d365a7c2 ddecbbf0 3009f343 39fa02af

33313330 0d06092a 864886f7 0d010105 05000382 01010093 244a305f 62cfd81a

982f3dea dc992dbd 77f6a579 2238ecc4 a7a07812 ad620e45 7064c5e7 97662d98

097e5faf d6cc2865 f201aa08 1a47def9 f97c925a 0869200d d93e6d6e 3c0d6ed8

e6069140 18b9f8c1 eddfdb41 aae09620 c9cd6415 3881c994 eea28429 0b136f8e

db0cdd25 02dba48b 1944d241 7a05694a 584f60ca 7e826a0b 02aa2517 39b5db7f

e784652a 958abd86 de5e8116 832d10cc defda882 2a6d281f 0d0bc4e5 e71a2619

e1f4116f 10b595fc e7420532 dbce9d51 5e28b69e 85d35bef a57d4540 728eb70e

6b0e06fb 33354871 b89d278b c4655f0d 86769c44 7af6955c f65d3208 33a454b6

183f685c f2424a85 3854835f d1e82cf2 ac11d6a8 ed636a

quit

crypto ikev2 policy 1

encryption aes-256

integrity sha

group 2

prf sha

lifetime seconds 86400

crypto ikev2 enable outside

telnet timeout 5

ssh stricthostkeycheck

ssh timeout 5

ssh key-exchange group dh-group1-sha1

console timeout 0

threat-detection basic-threat

threat-detection statistics access-list

no threat-detection statistics tcp-intercept

dynamic-access-policy-record DfltAccessPolicy

tunnel-group 192.0.2.6 type ipsec-l2l

tunnel-group 192.0.2.6 ipsec-attributes

ikev2 remote-authentication pre-shared-key MY_KEY

ikev2 local-authentication pre-shared-key MY_KEY

!

class-map inspection_default

match default-inspection-traffic

!

!

policy-map type inspect dns preset_dns_map

parameters

message-length maximum client auto

message-length maximum 512

no tcp-inspection

policy-map global_policy

class inspection_default

inspect ip-options

inspect netbios

inspect rtsp

inspect sunrpc

inspect tftp

inspect xdmcp

inspect dns preset_dns_map

inspect ftp

inspect h323 h225

inspect h323 ras

inspect rsh

inspect esmtp

inspect sqlnet

inspect sip

inspect skinny

policy-map type inspect dns migrated_dns_map_2

parameters

message-length maximum client auto

message-length maximum 512

no tcp-inspection

policy-map type inspect dns migrated_dns_map_1

parameters

message-length maximum client auto

message-length maximum 512

no tcp-inspection

!

service-policy global_policy global

prompt hostname context

no call-home reporting anonymous

call-home

profile CiscoTAC-1

no active

destination address http https://tools.cisco.com/its/service/oddce/services/DDCEService

destination address email callhome@cisco.com

destination transport-method http

subscribe-to-alert-group diagnostic

subscribe-to-alert-group environment

subscribe-to-alert-group inventory periodic monthly

subscribe-to-alert-group configuration periodic monthly

subscribe-to-alert-group telemetry periodic daily

profile License

destination address http https://tools.cisco.com/its/service/oddce/services/DDCEService

destination transport-method http

Cryptochecksum:a1d62628ba6d9e5af3ebffb781b6f968

: end

ASA2#

R1#sh run

Building configuration...

Current configuration : 3111 bytes

!

version 15.9

service timestamps debug datetime msec

service timestamps log datetime msec

no service password-encryption

!

hostname R1

!

boot-start-marker

boot-end-marker

!

!

!

no aaa new-model

!

!

!

mmi polling-interval 60

no mmi auto-configure

no mmi pvc

mmi snmp-timeout 180

!

!

!

!

!

no ip icmp rate-limit unreachable

!

!

!

!

!

!

no ip domain lookup

ip cef

no ipv6 cef

!

multilink bundle-name authenticated

!

!

!

redundancy

!

no cdp log mismatch duplex

!

ip tcp synwait-time 5

!

!

!

!

interface GigabitEthernet0/0

ip address 172.16.0.1 255.255.255.252

duplex auto

speed auto

media-type rj45

!

interface GigabitEthernet0/1

ip address 192.0.2.5 255.255.255.252

duplex auto

speed auto

media-type rj45

!

interface GigabitEthernet0/2

no ip address

shutdown

duplex auto

speed auto

media-type rj45

!

interface GigabitEthernet0/3

no ip address

shutdown

duplex auto

speed auto

media-type rj45

!

ip forward-protocol nd

!

!

no ip http server

!

ipv6 ioam timestamp

!

!

!

control-plane

!

!

line con 0

exec-timeout 0 0

privilege level 15

logging synchronous

line aux 0

exec-timeout 0 0

privilege level 15

logging synchronous

line vty 0 4

login

transport input none

!

no scheduler allocate

!

end

R1#

6. IPsec Tunnel bring up

- To bring up IPsec tunnel it is needed to generate interesting traffic from one side (except permanent IPsec tunnels).
- Interesting traffic is the traffic that is allowed in the encryption domain.
- The encryption domain represents the traffic that participates in VPN Tunnel.

In our case encryption domain on ASA1 contains subnet 10.0.0.0/24 and encryption domain on ASA2 contains 10.1.0.0/24. If we want to bring up the IPsec tunnel from ASA1 side for example, it is needed to generate traffic for other side of the tunnel with destination in network 10.1.0.0/24.

**a)** verify that no VPN tunnel exists before we will bring the VPN tunnel up:

ASA1# show crypto ipsec sa

There are no ipsec sas

ASA1#

ASA2# sh crypto ipsec sa

There are no ipsec sas

ASA2#

**b)** bring up the IPsec VPN tunnel up:


**Note:** I will generate interesting traffic using icmp ping test from PC1 (10.0.0.1) to PC2 (10.1.0.1)

PC1> ping 10.1.0.1

10.1.0.1 icmp_seq=1 timeout

84 bytes from 10.1.0.1 icmp_seq=2 ttl=64 time=4.949 ms

84 bytes from 10.1.0.1 icmp_seq=3 ttl=64 time=2.930 ms

84 bytes from 10.1.0.1 icmp_seq=4 ttl=64 time=2.192 ms

84 bytes from 10.1.0.1 icmp_seq=5 ttl=64 time=3.235 ms

PC1>

IPsec Tunnel is up and functional.

7. Verify IPsec tunnel status

There are following major ways to check IPsec tunnel status:

- show crypto ipsec sa (Show IPsec SAs for Phase 1)
- show crypto ikev2 sa (Show IKEv2 sas for Phase 2)

Let’s verify IPsec tunnel status on ASA1 and ASA2 and take a look on IPsec SA parameter and IKEv2 parameter.

ASA1# show crypto ipsec sa

interface: outside

Crypto map tag: outside_map, seq num: 10, local addr: 192.0.2.6

access-list OUTSIDE_CRYPTOMAP_10 extended permit ip 10.0.0.0 255.255.255.0 10.1.0.0 255.255.255.0

local ident (addr/mask/prot/port): (10.0.0.0/255.255.255.0/0/0)

remote ident (addr/mask/prot/port): (10.1.0.0/255.255.255.0/0/0)

current_peer: 172.16.0.2

#pkts encaps: 4, #pkts encrypt: 4, #pkts digest: 4

#pkts decaps: 4, #pkts decrypt: 4, #pkts verify: 4

#pkts compressed: 0, #pkts decompressed: 0

#pkts not compressed: 4, #pkts comp failed: 0, #pkts decomp failed: 0

#pre-frag successes: 0, #pre-frag failures: 0, #fragments created: 0

#PMTUs sent: 0, #PMTUs rcvd: 0, #decapsulated frgs needing reassembly: 0

#TFC rcvd: 0, #TFC sent: 0

#Valid ICMP Errors rcvd: 0, #Invalid ICMP Errors rcvd: 0

#send errors: 0, #recv errors: 0

local crypto endpt.: 192.0.2.6/500, remote crypto endpt.: 172.16.0.2/500

path mtu 1500, ipsec overhead 74(44), media mtu 1500

PMTU time remaining (sec): 0, DF policy: copy-df

ICMP error validation: disabled, TFC packets: disabled

current outbound spi: E6376DE9

current inbound spi : 527873A4

inbound esp sas:

spi: 0x527873A4 (1383625636)

SA State: active

transform: esp-aes-256 esp-sha-hmac no compression

in use settings ={L2L, Tunnel, IKEv2, }

slot: 0, conn_id: 233852928, crypto-map: outside_map

sa timing: remaining key lifetime (kB/sec): (3916799/28773)

IV size: 16 bytes

replay detection support: Y

