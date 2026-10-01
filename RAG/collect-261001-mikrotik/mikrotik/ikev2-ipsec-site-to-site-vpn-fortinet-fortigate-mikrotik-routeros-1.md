---
id: collect-261001-mikrotik/mikrotik/ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-mikrotik-routeros-1
title: "ID        STATE        UPTIME  PH2-TOTAL  REMOTE-ADDRESS"
domain: mikrotik
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-mikrotik/ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-mikrotik-routeros.md
source_anchor: ""
source_lines: [1, 135]
sha256: 171574ff3906403bdef36a7d3a41429f1619729069479f81e60f04974107062e
---

# ID        STATE        UPTIME  PH2-TOTAL  REMOTE-ADDRESS

This is a step-by-step tutorial to set up a site-to-site VPN between a Fortinet FortiGate and a Mikrotik RouterOS. The Key Exchange will be done using IKEv2 and both sites are using static ip-addresses on their wan interfaces.

**I typically use the strongest possible cryptographic algorithms between the two sites / vendors in my tutorials. This can result in degraded performance and higher ressource usage depending on the used hardware. You need to test the ressource usage and performance in your own environment.**

These are the devices I used for this tutorial:

Fortinet FortiGate 60F with FortiOS 7.2.3

Mikrotik hap Lite with RouterOS 7.6

**Lab**

The following figure shows the lab environment I build for this tutorial:

**FortiGate Configuration using the WebUI**

First, we need to create a new custom tunnel in the FortiGate configuration, where we set the basic parts as the peer ip-address and the interface we want to use for our vpn connection:

Then we set our pre-shared key and change the IKE Version to „2“:

As the other site can’t use GCM Ciphers in IKE / phase 1 right now, we need to stick with AES-CBC mode ciphers. Therefore, we use AES256 with SHA512 and the Diffie-Hellman Group 21 which is also known as 521-bit ECP:

Then we create the Phase 2 Selector with the networks we want to connect.

In the phase 2 the other site is able to use GCM ciphers, therefore we use AES256GCM and Diffie-Hellman Group 21:

Now just save the configuration.

**FortiGate Configuration using the CLI**

Using the FortiOS cli the configuration is done like this:

```
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
```
**RouterOS Configuration using Winbox**

All configuration is done in the „IP –> IPSec“ section using Winbox.

First we need to create the „IPsec Profile“ in which we define the IKE proposal:

In the next step, we create a new „IPSec Proposal“ for the phase 2 encryption.

No „Auth. Algorithms“ are needed, as we use aes-256-gcm as the encryption algorithm which already includes the authentication part:

For the peer configuration we only need to set the name, ip-address, ipsec profile and the „Exchange Mode“ to IKE2:

To set the authentication method, which is a „pre shared key“ in my case, we need to add a new „IPsec Identity“:

Then we define which networks need to talk to each other using the vpn tunnel:

In the last step we just need to select the „IPsec Proposal“ which we named „fortigate“, to use the correct encryption in phase 2 / esp:

**RouterOS Configuration using the CLI**

Using the RouterOS cli the configuration is done like this:

```
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
```
**Monitoring**

To view the currently established tunnels you can use the „IPsec Monitor“ in the FortiGates WebUI:

To my knowledge there are two ways to display the vpn tunnels using the commandline interface, the first one is this diagnose command which will display the phase 1 and phase 2 information with some details:

```
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
```
Or you can just print the ike / phase 1 part of your vpn:

