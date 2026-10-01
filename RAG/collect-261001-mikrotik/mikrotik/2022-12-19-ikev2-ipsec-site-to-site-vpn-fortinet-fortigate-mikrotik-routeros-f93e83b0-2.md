---
id: collect-261001-mikrotik/mikrotik/2022-12-19-ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-mikrotik-routeros-f93e83b0-2
title: "IKEv2 / IPsec Site-to-Site VPN Fortinet FortiGate <-> Mikrotik RouterOS"
domain: mikrotik
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-mikrotik/2022-12-19-ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-mikrotik-routeros-f93e83b0.md
source_anchor: ""
source_lines: [5, 131]
sha256: 2bd5bffd691aa2d17816f5e58a0e86495fefba83391bb3dd42a9a5d2eb82a83a
---

# IKEv2 / IPsec Site-to-Site VPN Fortinet FortiGate <-> Mikrotik RouterOS

This is a step-by-step tutorial to set up a site-to-site VPN between a Fortinet FortiGate and a Mikrotik RouterOS. The Key Exchange will be done using IKEv2 and both sites are using static ip-addresses on their wan interfaces.
I typically use the strongest possible cryptographic algorithms between the two sites / vendors in my tutorials. This can result in degraded performance and higher ressource usage depending on the used hardware.
You need to test the ressource usage and performance in your own environment.
These are the devices I used for this tutorial:
Fortinet FortiGate 60F with FortiOS 7.2.3
Mikrotik hap Lite with RouterOS 7.6
Lab
The following figure shows the lab environment I build for this tutorial:
FortiGate Configuration using the WebUI
First, we need to create a new custom tunnel in the FortiGate configuration, where we set the basic parts as the peer ip-address and the interface we want to use for our vpn connection:
Then we set our pre-shared key and change the IKE Version to „2“:
As the other site can’t use GCM Ciphers in IKE / phase 1 right now, we need to stick with AES-CBC mode ciphers. Therefore, we use AES256 with SHA512 and the Diffie-Hellman Group 21 which is also known as 521-bit ECP:
Then we create the Phase 2 Selector with the networks we want to connect.
In the phase 2 the other site is able to use GCM ciphers, therefore we use AES256GCM and Diffie-Hellman Group 21:
Now just save the configuration.
FortiGate Configuration using the CLI
Using the FortiOS cli the configuration is done like this:
config vpn ipsec phase1-interface
    edit "vpn-to-mikrotik"
        set interface "wan2"
        set ike-version 2
        set peertype any
        set net-device disable
        set proposal aes256-sha512
        set dhgrp 21
        set remote-gw 10.0.0.2
        set psksecret <PRESHAREDKEY>
    next
end
config vpn ipsec phase2-interface
    edit "vpn-to-mikrotik"
        set phase1name "vpn-to-mikrotik"
        set proposal aes256gcm
        set dhgrp 21
        set auto-negotiate enable
        set keylifeseconds 3600
        set src-subnet 192.168.100.0 255.255.255.0
        set dst-subnet 192.168.200.0 255.255.255.0
    next
end
RouterOS Configuration using Winbox
All configuration is done in the „IP –> IPSec“ section using Winbox.
First we need to create the „IPsec Profile“ in which we define the IKE proposal:
In the next step, we create a new „IPSec Proposal“ for the phase 2 encryption.
No „Auth. Algorithms“ are needed, as we use aes-256-gcm as the encryption algorithm which already includes the authentication part:
For the peer configuration we only need to set the name, ip-address, ipsec profile and the „Exchange Mode“ to IKE2:
To set the authentication method, which is a „pre shared key“ in my case, we need to add a new „IPsec Identity“:
Then we define which networks need to talk to each other using the vpn tunnel:
In the last step we just need to select the „IPsec Proposal“ which we named „fortigate“, to use the correct encryption in phase 2 / esp:
RouterOS Configuration using the CLI
Using the RouterOS cli the configuration is done like this:
/ip ipsec profile
add dh-group=ecp521 dpd-interval=2m dpd-maximum-failures=5 enc-algorithm=aes-256 hash-algorithm=sha512 lifetime=1d name=fortigate \
    nat-traversal=no proposal-check=obey
/ip ipsec peer
add address=10.0.0.1/32 disabled=no exchange-mode=ike2 name=fortigate profile=fortigate send-initial-contact=yes
/ip ipsec proposal
add auth-algorithms="" disabled=no enc-algorithms=aes-256-gcm lifetime=1h name=fortigate pfs-group=ecp521
/ip ipsec identity
add auth-method=pre-shared-key disabled=no generate-policy=no peer=fortigate secret=<PRESHAREDKEY>
/ip ipsec policy
add action=encrypt disabled=no dst-address=192.168.100.0/24 dst-port=any ipsec-protocols=esp level=require peer=fortigate proposal=fortigate \
    protocol=all sa-dst-address=10.0.0.1 sa-src-address=10.0.0.2 src-address=192.168.200.0/24 src-port=any tunnel=yes
Monitoring
To view the currently established tunnels you can use the „IPsec Monitor“ in the FortiGates WebUI:
To my knowledge there are two ways to display the vpn tunnels using the commandline interface, the first one is this diagnose command which will display the phase 1 and phase 2 information with some details:
FG60F (root) # diagnose vpn tunnel list name vpn-to-mikrotik
list all ipsec tunnel in vd 0
------------------------------------------------------
name=vpn-to-mikrotik ver=2 serial=2 10.0.0.1:0->10.0.0.2:0 tun_id=10.0.0.2 tun_id6=::10.0.0.2 dst_mtu=1500 dpd-link=on weight=1
bound_if=6 lgwy=static/1 tun=intf/0 mode=auto/1 encap=none/520 options[0208]=npu frag-rfc  run_state=0 role=primary accept_traffic=1 overlay_id=0
proxyid_num=1 child_num=0 refcnt=4 ilast=43746464 olast=43746464 ad=/0
stat: rxp=0 txp=0 rxb=0 txb=0
dpd: mode=on-demand on=1 idle=20000ms retry=3 count=0 seqno=0
natt: mode=none draft=0 interval=0 remote_port=0
fec: egress=0 ingress=0
proxyid=vpn-to-mikrotik proto=0 sa=2 ref=3 serial=1 auto-negotiate
  src: 0:192.168.100.0/255.255.255.0:0
  dst: 0:192.168.200.0/255.255.255.0:0
  SA:  ref=3 options=18227 type=00 soft=0 mtu=1280 expire=2953/0B replaywin=2048
       seqno=1 esn=0 replaywin_lastseq=00000000 qat=0 rekey=0 hash_search_len=1
  life: type=01 bytes=0/0 timeout=3301/3600
  dec: spi=83391146 esp=aes-gcm key=36 fe03aea03aa08dee28c9a4ad2022d55113ea72be2078e0d5c77c2eacafe5ccb7d87e3611
       ah=null key=0 
  enc: spi=09387f45 esp=aes-gcm key=36 6cb5756673b296440c9a0e428662abd80e8e444e97eb2ba8777fd2f9c9ff6a6e033a5f77
       ah=null key=0 
  dec:pkts/bytes=0/0, enc:pkts/bytes=0/0
  npu_flag=00 npu_rgwy=10.0.0.2 npu_lgwy=10.0.0.1 npu_selid=0 dec_npuid=0 enc_npuid=0
  SA:  ref=3 options=18227 type=00 soft=0 mtu=1280 expire=2982/0B replaywin=2048
       seqno=1 esn=0 replaywin_lastseq=00000000 qat=0 rekey=0 hash_search_len=1
  life: type=01 bytes=0/0 timeout=3333/3600
  dec: spi=83391147 esp=aes-gcm key=36 7e4d12c5363b3aec6c29a9872e50ad65f4e1969e8dbc2db753853c620fd70e0448fd3339
       ah=null key=0 
  enc: spi=03b2b08e esp=aes-gcm key=36 034b284ccfdd7bdfb7b6aeb81f9c8f9a94fc5eb03e595fc0966011c46aa88b6637b06854
       ah=null key=0 
  dec:pkts/bytes=0/0, enc:pkts/bytes=0/0
  npu_flag=00 npu_rgwy=10.0.0.2 npu_lgwy=10.0.0.1 npu_selid=0 dec_npuid=0 enc_npuid=0
run_tally=0
Or you can just print the ike / phase 1 part of your vpn:
FG60F (root) # get vpn ike gateway vpn-to-mikrotik
vd: root/0
name: vpn-to-mikrotik
version: 2
interface: wan2 6
addr: 10.0.0.1:500 -> 10.0.0.2:500
created: 459s ago
IKE SA  created: 2/2  established: 2/2  time: 300/10955/21610 ms
IPsec SA  created: 2/2  established: 2/2  time: 0/10805/21610 ms
  id/spi: 640 4dd64623860af1e1/16dec42f467670cd
  direction: responder
  status: established 441-441s ago = 300ms
  proposal: aes-256-sha512
  SK_ei: 552a9cfed6295473-2ee990b3fbfa3f52-0172646714a8ec84-010914fd7c314374
  SK_er: a70f8a21e6094b5a-80f379933f7b9b81-f9b4de02761cb4fc-e1869da813666032
  SK_ai: 8f117271dc7dce87-6461855320baf6ea-af038b787a3944d0-50192b83114d6d24-2a4a0a0c1e6eccf6-de847595def40637-1deff08961119347-aac53969384f69ab
  SK_ar: 0452b451c93c7dad-23c6d9f14e625710-a506d276c483c5a3-a0e07f0aa7fa0cd7-e33480c6f71c35e1-e5e8b0bdd0077f28-adda01d6138ce560-7e5cbca59e555b09
  lifetime/rekey: 86400/85688
  DPD sent/recv: 00000000/00000000
  id/spi: 639 0916e0d5e27adb99/6cd98939eb60ec00
  direction: initiator
  status: established 459-438s ago = 21610ms
  proposal: aes-256-sha512
  SK_ei: 9b897f8353c19c7a-ba177039b790a7c0-5d6ef153267b440b-a884f5ca77ec3e49
  SK_er: 613e2dfd3fd2e14e-1f7233a2875e7d93-625cd5899ac0af73-873d6a6fc0daaf3b
  SK_ai: 6d2ee04ff40b6b0d-0e1e1e9770d11b00-7cf0b8f9cf1534c3-ee239a0275df3bba-387bab3d4634952d-27acb5a3ff6fbdfb-636aac1e4de570b8-1e47c134d15ed371
  SK_ar: 595588a3506d553d-c5c161e696017b7f-a304939e3b93cc39-55a445ad1e57e966-4e3f5f24364f50d1-f4be78fe1299166a-aad1c6b08853e3ee-0a54d6a3cb3fe96a
  lifetime/rekey: 86400/85661
