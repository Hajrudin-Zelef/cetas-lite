---
id: collect-261001-cisco/cisco/cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial-5
title: "cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial.md
source_anchor: ""
source_lines: [1537, 1754]
sha256: 88dcc317c7c2daaff9ac4b595b169e225c2a2e21cbc395bad2a029294b185f2f
---

# cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial

Anti replay bitmap:

0x00000000 0x0000001F

outbound esp sas:

spi: 0xE6376DE9 (3862392297)

SA State: active

transform: esp-aes-256 esp-sha-hmac no compression

in use settings ={L2L, Tunnel, IKEv2, }

slot: 0, conn_id: 233852928, crypto-map: outside_map

sa timing: remaining key lifetime (kB/sec): (4239359/28773)

IV size: 16 bytes

replay detection support: Y

Anti replay bitmap:

0x00000000 0x00000001

ASA1#

ASA1# show crypto ikev2 sa

IKEv2 SAs:

Session-id:57093, Status:UP-ACTIVE, IKE count:1, CHILD count:1

Tunnel-id Local Remote Status Role

9578947 192.0.2.6/500 172.16.0.2/500 READY INITIATOR

Encr: AES-CBC, keysize: 256, Hash: SHA96, DH Grp:2, Auth sign: PSK, Auth verify: PSK

Life/Active Time: 86400/33 sec

Child sa: local selector 10.0.0.0/0 - 10.0.0.255/65535

remote selector 10.1.0.0/0 - 10.1.0.255/65535

ESP spi in/out: 0x527873a4/0xe6376de9

ASA1#

ASA2# sh crypto ipsec sa

interface: outside

Crypto map tag: outside_map, seq num: 10, local addr: 172.16.0.2

access-list OUTSIDE_CRYPTOMAP_10 extended permit ip 10.1.0.0 255.255.255.0 10.0.0.0 255.255.255.0

local ident (addr/mask/prot/port): (10.1.0.0/255.255.255.0/0/0)

remote ident (addr/mask/prot/port): (10.0.0.0/255.255.255.0/0/0)

current_peer: 192.0.2.6

#pkts encaps: 4, #pkts encrypt: 4, #pkts digest: 4

#pkts decaps: 4, #pkts decrypt: 4, #pkts verify: 4

#pkts compressed: 0, #pkts decompressed: 0

#pkts not compressed: 4, #pkts comp failed: 0, #pkts decomp failed: 0

#pre-frag successes: 0, #pre-frag failures: 0, #fragments created: 0

#PMTUs sent: 0, #PMTUs rcvd: 0, #decapsulated frgs needing reassembly: 0

#TFC rcvd: 0, #TFC sent: 0

#Valid ICMP Errors rcvd: 0, #Invalid ICMP Errors rcvd: 0

#send errors: 0, #recv errors: 0

local crypto endpt.: 172.16.0.2/500, remote crypto endpt.: 192.0.2.6/500

path mtu 1500, ipsec overhead 74(44), media mtu 1500

PMTU time remaining (sec): 0, DF policy: copy-df

ICMP error validation: disabled, TFC packets: disabled

current outbound spi: 527873A4

current inbound spi : E6376DE9

inbound esp sas:

spi: 0xE6376DE9 (3862392297)

SA State: active

transform: esp-aes-256 esp-sha-hmac no compression

in use settings ={L2L, Tunnel, IKEv2, }

slot: 0, conn_id: 80543744, crypto-map: outside_map

sa timing: remaining key lifetime (kB/sec): (4239359/28761)

IV size: 16 bytes

replay detection support: Y

Anti replay bitmap:

0x00000000 0x0000001F

outbound esp sas:

spi: 0x527873A4 (1383625636)

SA State: active

transform: esp-aes-256 esp-sha-hmac no compression

in use settings ={L2L, Tunnel, IKEv2, }

slot: 0, conn_id: 80543744, crypto-map: outside_map

sa timing: remaining key lifetime (kB/sec): (4193279/28761)

IV size: 16 bytes

replay detection support: Y

Anti replay bitmap:

0x00000000 0x00000001

ASA2#

ASA2# show crypto ikev2 sa

IKEv2 SAs:

Session-id:19664, Status:UP-ACTIVE, IKE count:1, CHILD count:1

Tunnel-id Local Remote Status Role

10135907 172.16.0.2/500 192.0.2.6/500 READY RESPONDER

Encr: AES-CBC, keysize: 256, Hash: SHA96, DH Grp:2, Auth sign: PSK, Auth verify: PSK

Life/Active Time: 86400/43 sec

Child sa: local selector 10.1.0.0/0 - 10.1.0.255/65535

remote selector 10.0.0.0/0 - 10.0.0.255/65535

ESP spi in/out: 0xe6376de9/0x527873a4

ASA2#


##### **Knowledge base**

1. **interface: outside**
  1. interface on which the tunnel is established
2. **Crypto map tag: outside_map, seq num: 10, local addr: 192.0.2.6**
  1. crypto map for this tunnel
  2. ip address of local endpoint
3.  **access-list OUTSIDE_CRYPTOMAP_10 extended permit ip 10.0.0.0 255.255.255.0 10.1.0.0 255.255.255.0**
  1. access list allowed to communicate through the tunnel
4.  **local ident (addr/mask/prot/port): (10.0.0.0/255.255.255.0/0/0)****remote ident (addr/mask/prot/port): (10.1.0.0/255.255.255.0/0/0)****current_peer: 172.16.0.2**
  1. local encryption domain
  2. remote encryption domain
  3. remote peer ip
5.  **#pkts encaps: 4, #pkts encrypt: 4, #pkts digest: 4****#pkts decaps: 4, #pkts decrypt: 4, #pkts verify: 4****#pkts compressed: 0, #pkts decompressed: 0****#pkts not compressed: 4, #pkts comp failed: 0, #pkts decomp failed: 0****#pre-frag successes: 0, #pre-frag failures: 0, #fragments created: 0****#PMTUs sent: 0, #PMTUs rcvd: 0, #decapsulated frgs needing reassembly: 0****#TFC rcvd: 0, #TFC sent: 0****#Valid ICMP Errors rcvd: 0, #Invalid ICMP Errors rcvd: 0****#send errors: 0, #recv errors: 0**
  1. tunnel statistics, encrypted packets, decrypted packets etc.
6.  **local crypto endpt.: 192.0.2.6/500, remote crypto endpt.: 172.16.0.2/500**
  1. local peer ip and source port, remote peer ip and destination port
7.  **current outbound spi: E6376DE9****current inbound spi : 527873A4**
  1. spi for outbound and inbound tunnel. (The Security Parameter Index (SPI) is an identification tag added to the header while using IPsec for tunneling the IP traffic.)
8.  **inbound esp sas:**    **spi: 0x527873A4 (1383625636)**        **SA State: active**        **transform: esp-aes-256 esp-sha-hmac no compression**        **in use settings ={L2L, Tunnel, IKEv2, }**        **slot: 0, conn_id: 233852928, crypto-map: outside_map**        **sa timing: remaining key lifetime (kB/sec): (3916799/28773)**        **IV size: 16 bytes**        **replay detection support: Y**        **Anti replay bitmap:**        **0x00000000 0x0000001F****outbound esp sas:**     **spi: 0xE6376DE9 (3862392297)**        **SA State: active**        **transform: esp-aes-256 esp-sha-hmac no compression**        **in use settings ={L2L, Tunnel, IKEv2, }**        **slot: 0, conn_id: 233852928, crypto-map: outside_map**        **sa timing: remaining key lifetime (kB/sec): (4239359/28773)**        **IV size: 16 bytes**        **replay detection support: Y**        **Anti replay bitmap:**        **0x00000000 0x00000001**
  1. detailed parameter of the SPIs, SA State for the spi, lifetime, etc
9. **IKEv2 SAs:****Session-id:57093, Status:UP-ACTIVE, IKE count:1, CHILD count:1****Tunnel-id Local Remote Status Role****9578947 192.0.2.6/500 172.16.0.2/500 READY INITIATOR****Encr: AES-CBC, keysize: 256, Hash: SHA96, DH Grp:2, Auth sign: PSK, Auth verify: PSK****Life/Active Time: 86400/33 sec****Child sa: local selector 10.0.0.0/0 – 10.0.0.255/65535****remote selector 10.1.0.0/0 – 10.1.0.255/65535****ESP spi in/out: 0x527873a4/0xe6376de9**
  1. IKE version 2 parameter:
    1. Status, SA Count
    2. Local peer IP and Remote peer IP
    3. Node is Initiator or Responder
    4. Authentication method and encryption
    5. IKEv2 SA configured lifetime and active time
    6. local selector and remote selector (local/remote encryption domain parameter)
    7. responsible SPI sfor the phase 1
10. IKE version 2 parameter:

Commands for debugging and troubleshoot VPN tunnel problems:

The ASA debugs for tunnel negotiation are:

- **debug crypto ikev2 protocol**
- **debug crypto ikev2 platform**

The ASA debug for certificate authentication is:

- **debug crypto ca**

VPN Tunnel reset:

- **clear crypto ipsec sa peer <remote-peer-IP>**

Learn more on Cisco TechNotes:

8. Tunnel initiation – Packet Capture

In this section we will take a deeper look how the IPsec IKEv2 Site-to-Site VPN Tunnel initiation looks like on the packet level. I will make packet capture on ASA1 and ASA2 outside interfaces at the same time and then interesting traffic will be generated to bring up the tunnel.

##### **Knowledge base**

