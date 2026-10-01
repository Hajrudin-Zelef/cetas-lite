---
id: collect-261001-general-networking/general-networking/c-en-us-support-docs-security-vpn-ipsec-negotiation-ike-protocols-117337-config-fb1867ae-4
title: "c-en-us-support-docs-security-vpn-ipsec-negotiation-ike-protocols-117337-config--fb1867ae"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "voice"]
source: docs/RAG/collect-261001-general-networking/c-en-us-support-docs-security-vpn-ipsec-negotiation-ike-protocols-117337-config--fb1867ae.md
source_anchor: ""
source_lines: [421, 525]
sha256: de62d5bc71fb59775820bbf4b8b4ed48c04fced1b34e4d86a99e503eedcc5da5
---

# c-en-us-support-docs-security-vpn-ipsec-negotiation-ike-protocols-117337-config--fb1867ae

AA301D06 03551D0E 04160414 082896B9 F4AF2075 514321D0 72F161D0 9D2EC8AA
300D0609 2A864886 F70D0101 04050003 81810087 A06D354A F7423E0E 64A7C5EC
6006FBDE 914D7BFD F86ADA50 B1A00D17 0BF06EC1 5423D514 FBEB0A76 986EB63F
F7FCE99A 81C4B112 61FD69CE A2CE750E B1B3A6F9 84E92490 8F213613 451DD9A8
3FC3406A 854B20ED 27E4DDD8 62F6DEA5 DD8B4396 1879B3E7 651CB9D1 3DD46B8B
32796963 9F6854F1 389F0060 AA0D1B8D F83E09
quit
voice-card 0
!
!
interface Loopback0
ip address 192.168.254.254 255.255.255.255
!
interface GigabitEthernet0/0
ip address 192.168.0.254 255.255.255.0
duplex auto
speed auto
!
interface GigabitEthernet0/1
ip address 192.168.1.254 255.255.255.0
duplex auto
speed auto
!
! http-server needs to be enabeld for SCEP
!
ip http server
no ip http secure-server
!
ip route 0.0.0.0 0.0.0.0 10.122.162.129
ip route 172.18.108.26 255.255.255.255 10.122.162.129
!
! ntp configuration
!
ntp trusted-key 1
ntp master 1
!
end
Use this section in order to confirm that your configuration works properly.
These commands work on both ASAs and routers:
show crypto ikev2 sa - Displays the state of the phase 1 Security Association (SA).show crypto ipsec sa - Displays the state of the phase 2 SA.
Note: In this output, unlike in IKEv1, the Perfect Forwarding Secrecy (PFS) Diffie-Hellman (DH) group value displays as 'PFS (Y/N): N, DH group: none' during the first tunnel negotiation; after a rekey occurs, the correct values appear. This is not a bug, but is expected behavior.
The difference between IKEv1 and IKEv2 is that, in IKEv2, the Child SAs are created as part of the AUTH exchange itself. The DH Group configured under the crypto map is used only during a rekey. Thus, you see 'PFS (Y/N): N, DH group: none' until the first rekey. With IKEv1, you see a different behavior because Child SA creation happens during Quick Mode, and the CREATE_CHILD_SA message has the provision to carry the Key Exchange payload, which specifies the DH parameters to derive the new shared secret.
This procedure verifies phase 1 activity:
R1#show crypto ikev2 sa
IPv4 Crypto IKEv2 SA
Tunnel-id Local Remote fvrf/ivrf Status
1 172.16.1.1/500 172.16.1.2/500 none/none READY
Encr: AES-CBC, keysize: 256, Hash: SHA96, DH Grp:14, Auth sign: RSA,
Auth verify: RSA
Life/Active Time: 86400/53 sec
IPv6 Crypto IKEv2 SA
ciscoasa/vpn(config)# show crypto ikev2 sa
IKEv2 SAs:
Session-id:138, Status:UP-ACTIVE, IKE count:1, CHILD count:1
Tunnel-id Local Remote Status Role
45926289 172.16.1.2/500 172.16.1.1/500 READY INITIATOR
Encr: AES-CBC, keysize: 256, Hash: SHA96, DH Grp:14, Auth sign: RSA,
Auth verify: RSA
Life/Active Time: 86400/4 sec
Child sa: local selector 192.168.0.0/0 - 192.168.0.255/65535
remote selector 172.16.2.0/0 - 172.16.2.255/65535
ESP spi in/out: 0xa84caabb/0xf18dce57
This procedure describes how to verify if the Security Parameter Index (SPI) has been negotiated correctly on the two peers:
show crypto ipsec sa | i spi command on the router:
R1#show crypto ipsec sa | i spi 
current outbound spi: 0xA84CAABB(2823596731)
spi: 0xF18DCE57(4052602455)
spi: 0xA84CAABB(2823596731)
ciscoasa/vpn(config)# show crypto ipsec sa | i spi
current outbound spi: F18DCE57
current inbound spi : A84CAABB
spi: 0xA84CAABB (2823596731)
spi: 0xF18DCE57 (4052602455)
This procedure describes how to confirm whether traffic flows across the tunnel:
show crypto ipsec sa | i pkts command on the router:
R1#show crypto ipsec sa | i pkts
#pkts encaps: 21, #pkts encrypt: 21, #pkts digest: 21
#pkts decaps: 30, #pkts decrypt: 30, #pkts verify: 30
#pkts compressed: 0, #pkts decompressed: 0
#pkts not compressed: 0, #pkts compr. failed: 0
#pkts not decompressed: 0, #pkts decompress failed: 0
ciscoasa/vpn(config)# show crypto ipsec sa | i pkts
#pkts encaps: 4, #pkts encrypt: 4, #pkts digest: 4
#pkts decaps: 4, #pkts decrypt: 4, #pkts verify: 4
#pkts compressed: 0, #pkts decompressed: 0
#pkts not compressed: 4, #pkts comp failed: 0, #pkts decomp
failed: 0
This section provides information you can use in order to troubleshoot your configuration.
Note: Refer to Important Information on Debug Commands before you use debug commands.
Caution: On the ASA, you can set various debug levels; by default, level 1 is used. If you change the debug level, the verbosity of the debugs can increase. Do this with caution, especially in production environments!
The ASA debugs for tunnel negotiation are:
debug crypto ikev2 protocoldebug crypto ikev2 platform
The ASA debug for certificate authentication is:
debug crypto ca
The router debugs for tunnel negotiation are:
debug crypto ikev2debug crypto ikev2 errordebug crypto ikev2 internal
The router debugs for certificate authentication are:
debug cry pki validationdebug cry pki transactiondebug cry pki messages
| Revision | Publish Date | Comments | 
|---|---|---|
| 3.0 |                                                                                               29-Aug-2023                                                                                       | Edited the title. Updated device and software under Components Used. Deleted or updated broken links. Edited for clarity.  Updated description. | 
| 2.0 |                                                                                               21-Jul-2022                                                                                       | Edited the title. Updated device and software under Components Used. Deleted or updated broken links. Edited for clarity. | 
| 1.0 |                                                                                               13-Feb-2014                                                                                       | Initial Release | 
Lorsque vous consultez un site Web, celui-ci peut enregistrer ou récupérer des informations présentes dans votre navigateur, le plus souvent sous la forme de témoins. Ces informations peuvent vous concerner, ainsi que vos préférences ou votre périphérique, et servent principalement à faire fonctionner le site selon vos attentes. Généralement, ces informations ne permettent pas de vous identifier directement, mais elles sont utilisées afin de vous offrir une expérience Web personnalisée. Parce que nous respectons votre droit à la vie privée, vous pouvez choisir de ne pas autoriser certains types de témoins. Cliquez sur les différents en-têtes de catégorie pour en apprendre davantage et modifier nos paramètres par défaut. Toutefois, bloquer certains types de témoins peut avoir une incidence sur votre expérience du site et les services que nous proposons.
Ces cookies sont indispensables au bon fonctionnement du site web et ne peuvent pas être désactivés au niveau de nos systèmes. Ils ne résultent généralement que d'actions que vous avez effectuées et qui correspondent à une demande de service, comme lorsque vous définissez vos préférences en matière de confidentialité, que vous vous connectez ou que vous remplissez des formulaires. Vous pouvez configurer votre navigateur pour qu'il bloque ces cookies ou vous avertisse de leur présence, mais certaines parties du site ne seront alors plus opérationnelles. Ces cookies n'enregistrent aucune information d'identification personnelle.
