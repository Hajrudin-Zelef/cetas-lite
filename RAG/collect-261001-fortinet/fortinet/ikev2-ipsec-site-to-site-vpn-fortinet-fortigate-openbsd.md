---
id: collect-261001-fortinet/fortinet/ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-openbsd
title: "ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-openbsd"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-openbsd.md
source_anchor: ""
source_lines: [1, 197]
sha256: 1bbe653351efabb166deac3b2d680b6d0c8070d5d382d9b1859ae9d37e7e4aa7
---

# ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-openbsd

This is a step-by-step tutorial to set up a site-to-site VPN between a Fortinet FortiGate and a OpenBSD VPN-Gateway. The Key Exchange will be done using IKEv2 and both sites are using static ip-addresses on their wan interfaces.

**I typically use the strongest possible cryptographic algorithms between the two sites / vendors in my tutorials. This can result in degraded performance and higher ressource usage depending on the used hardware. You need to test the ressource usage and performance in your own environment.**

These are the devices I used for this tutorial:

Fortinet FortiGate 60F with FortiOS 7.2.3

OpenBSD 7.2 amd64 Virtual Machine

**Lab**

The following figure shows the lab environment I build for this tutorial:

**FortiGate Configuration using the WebUI**

First, we need to create a new custom tunnel in the FortiGate configuration, where we set the basic parts as the peer ip-address and the interface we want to use for our vpn connection:

Then we set our pre-shared key and change the IKE Version to „2“:

The encryption algorithms are set to AES256GCM with PRFSHA512 and Diffie-Hellmann-Group 31 which is also known as curve25519:

Then we create the Phase 2 Selector with the networks we want to connect.

For the encryption we use AES256GCM with Diffie-Hellmann-Group 31:

Now the vpn configuration is finished and can be saved.

As a last step a static route is needed, which tells the fortigate to route the destination network to the vpn tunnel interface:

**FortiGate Configuration using the CLI**

Using the FortiOS cli the configuration is done like this:

```
config vpn ipsec phase1-interface
    edit "vpn-to-obsd"
        set interface "wan1"
        set ike-version 2
        set peertype any
        set net-device disable
        set proposal aes256gcm-prfsha512
        set dhgrp 31
        set remote-gw 10.0.0.3
        set psksecret ENC <PRESHAREDKEY>
    next
end
config vpn ipsec phase2-interface
    edit "vpn-to-obsd"
        set phase1name "vpn-to-obsd"
        set proposal aes256gcm
        set auto-negotiate enable
        set src-subnet 192.168.100.0 255.255.255.0
        set dst-subnet 192.168.210.0 255.255.255.0
    next
end
config router static
    edit 0
        set dst 192.168.210.0 255.255.255.0
        set device "vpn-to-obsd"
    next
end
```
**OpenBSD Configuration**

I expect that the OpenBSD system is already configured / used as an vpn gateway.

If not, check the OpenBSD VPN FAQ first:

OpenBSD FAQ – Virtual Private Networks

First we need to add the configuration to the /etc/iked.conf file:

```
ikev2 "lab" esp \
	from 192.168.210.0/24 to 192.168.100.0/24 \
	local 10.0.0.3 peer 10.0.0.1 \
	ikesa enc aes-256 auth hmac-sha2-512 group curve25519 \
	childsa enc aes-256-gcm prf hmac-sha2-512 group curve25519 \
	ikelifetime 24h \
	lifetime 1h \
	psk "PRESHAREDKEY"
```
Reload the iked daemon to activate the new config:

`rcctl reload iked`
**Monitoring**

To view the currently established tunnels you can use the „IPsec Monitor“ in the FortiGates WebUI:

Or you can use the cli to view the established tunnel:

```
FG60F (root) # diagnose vpn tunnel list name vpn-to-obsd
list ipsec tunnel by names in vd 0
------------------------------------------------------
name=vpn-to-obsd ver=2 serial=4 10.0.0.1:500->10.0.0.3:500 tun_id=10.0.0.3 tun_id6=::10.0.0.3 dst_mtu=1500 dpd-link=on weight=1
bound_if=5 lgwy=static/1 tun=intf/0 mode=auto/1 encap=none/520 options[0208]=npu frag-rfc  run_state=0 role=primary accept_traffic=1 overlay_id=0
proxyid_num=1 child_num=0 refcnt=6 ilast=43916469 olast=43916469 ad=/0
stat: rxp=0 txp=0 rxb=0 txb=0
dpd: mode=on-demand on=1 idle=20000ms retry=3 count=0 seqno=0
natt: mode=keepalive draft=0 interval=10 remote_port=4500
fec: egress=0 ingress=0
proxyid=vpn-to-obsd proto=0 sa=1 ref=2 serial=2 auto-negotiate
  src: 0:192.168.100.0/255.255.255.0:0
  dst: 0:192.168.210.0/255.255.255.0:0
  SA:  ref=3 options=18227 type=00 soft=0 mtu=1438 expire=42133/0B replaywin=2048
       seqno=1 esn=0 replaywin_lastseq=00000000 qat=0 rekey=0 hash_search_len=1
  life: type=01 bytes=0/0 timeout=42902/43200
  dec: spi=83393171 esp=aes-gcm key=36 0f49113c4df6317f344adfff2f628dcd1f8ef3e75d1e370dfd83140d73ea5a70d0dd8ee4
       ah=null key=0 
  enc: spi=6fcc7d17 esp=aes-gcm key=36 fc33f9610679a3522b72def31af9aaa7fa2aa3841c6a991b90a225fc2bd67126e79b5865
       ah=null key=0 
  dec:pkts/bytes=0/0, enc:pkts/bytes=0/0
  npu_flag=00 npu_rgwy=10.0.0.3 npu_lgwy=10.0.0.1 npu_selid=3 dec_npuid=0 enc_npuid=0
run_tally=0
```
Or you can print the phase 1 and phase 2 part seperatly.

For the ike / phase 1 part you can print out the details using this command:

```
FG60F (root) # get vpn ike gateway vpn-to-obsd
vd: root/0
name: vpn-to-obsd
version: 2
interface: wan1 5
addr: 10.0.0.1:500 -> 10.0.0.3:500
created: 933s ago
peer-id: 10.0.0.3
peer-auth: no
IKE SA  created: 1/1  established: 1/1  time: 40/40/40 ms
IPsec SA  created: 1/1  established: 1/1  time: 40/40/40 ms
  id/spi: 8844 6c38e47ab0b977de/5b53bba8dd610f2c
  direction: initiator
  status: established 933-933s ago = 40ms
  proposal: aes256gcm
  SK_ei: fea20996c4834397-341f9180bf1e5232-89ba794f1f33ce9f-1aa9e23fdb8a68b0-9ecd78b7
  SK_er: 843c7787c4fba4f6-ac6b66f0ec42525f-215b87b53ab704e4-737303bfc0552eae-34cb8f7e
  SK_ai: 
  SK_ar: 
  lifetime/rekey: 86400/85166
  DPD sent/recv: 00000000/00000000
```
And for esp / phase 2 you can use this command:

```
FG60F (root) # get vpn ipsec tunnel name vpn-to-obsd
gateway
  name: 'vpn-to-obsd'
  local-gateway: 10.0.0.1:500 (static)
  remote-gateway: 10.0.0.3:500 (static)
  dpd-link: on
  mode: ike-v2
  interface: 'wan1' (5)
  rx  packets: 0  bytes: 0  errors: 0
  tx  packets: 0  bytes: 0  errors: 0
  dpd: on-demand/negotiated  idle: 20000ms  retry: 3  count: 0
  nat traversal mode: keep-alive      interval: 10
  selectors
    name: 'vpn-to-obsd'
    auto-negotiate: enable
    mode: tunnel
    src: 0:192.168.100.0/255.255.255.0:0
    dst: 0:192.168.210.0/255.255.255.0:0
    SA
      lifetime/rekey: 43200/41914   
      mtu: 1438
      tx-esp-seq: 1
      replay: enabled
      qat: 0
      inbound
        spi: 83393171
        enc:  aes-gc  0f49113c4df6317f344adfff2f628dcd1f8ef3e75d1e370dfd83140d73ea5a70d0dd8ee4
        auth:   null  
      outbound
        spi: 6fcc7d17
        enc:  aes-gc  fc33f9610679a3522b72def31af9aaa7fa2aa3841c6a991b90a225fc2bd67126e79b5865
        auth:   null  
      NPU acceleration: none
```
In OpenBSD you can use the „ikectl show sa“ to check if the ipsec flows are established:

```
obsd# ikectl show sa    
iked_sas: 0xfd98493780 rspi 0xb28ad1580eb6086f ispi 0x4d0c59adff475c62 10.0.0.3:500->10.0.0.1:500 ESTABLISHED r nexti 0x0 pol 0xfd9dbf9000
  sa_childsas: 0xfd984cf400 ESP 0xfb8cf58d in 10.0.0.1:4500 -> 10.0.0.3:500 (LA) B=0x0 P=0xfd9849ef00 @0xfd98493780
  sa_childsas: 0xfd9849ef00 ESP 0x83393172 out 10.0.0.3:500 -> 10.0.0.1:500 (L) B=0x0 P=0xfd984cf400 @0xfd98493780
  sa_flows: 0xfd984ae000 ESP out 192.168.210.0/24 -> 192.168.100.0/24 [0]@-1 (L) @0xfd98493780
  sa_flows: 0xfd984b0000 ESP in 192.168.100.0/24 -> 192.168.210.0/24 [0]@-1 (L) @0xfd98493780
iked_activesas: 0xfd9849ef00 ESP 0x83393172 out 10.0.0.3:500 -> 10.0.0.1:500 (L) B=0x0 P=0xfd984cf400 @0xfd98493780
iked_activesas: 0xfd984cf400 ESP 0xfb8cf58d in 10.0.0.1:500 -> 10.0.0.3:500 (LA) B=0x0 P=0xfd9849ef00 @0xfd98493780
iked_flows: 0xfd984b0000 ESP in 192.168.100.0/24 -> 192.168.210.0/24 [0]@-1 (L) @0xfd98493780
iked_flows: 0xfd984ae000 ESP out 192.168.210.0/24 -> 192.168.100.0/24 [0]@-1 (L) @0xfd98493780
iked_dstid_sas: 0xfd98493780 rspi 0xb28ad1580eb6086f ispi 0x4d0c59adff475c62 10.0.0.3:500->10.0.0.1:500 ESTABLISHED r nexti 0x0 pol 0xfd9dbf9000
```
## Schreibe einen Kommentar
