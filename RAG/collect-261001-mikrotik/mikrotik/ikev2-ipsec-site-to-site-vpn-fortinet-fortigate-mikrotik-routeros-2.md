---
id: collect-261001-mikrotik/mikrotik/ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-mikrotik-routeros-2
title: "ID        STATE        UPTIME  PH2-TOTAL  REMOTE-ADDRESS"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-mikrotik/ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-mikrotik-routeros.md
source_anchor: ""
source_lines: [136, 309]
sha256: d84ca138db5fb9e5e07f5f9ba852307fb758e126b02f45779ae0314aa205f53a
---

# ID        STATE        UPTIME  PH2-TOTAL  REMOTE-ADDRESS

```
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
  DPD sent/recv: 00000000/00000000
```
And then print the details of each phase 2 configuration on their own:

```
FG60F (root) # get vpn ipsec tunnel name vpn-to-mikrotik
gateway
  name: 'vpn-to-mikrotik'
  local-gateway: 10.0.0.1:0 (static)
  remote-gateway: 10.0.0.2:0 (static)
  dpd-link: on
  mode: ike-v2
  interface: 'wan2' (6)
  rx  packets: 0  bytes: 0  errors: 0
  tx  packets: 0  bytes: 0  errors: 0
  dpd: on-demand/negotiated  idle: 20000ms  retry: 3  count: 0
  selectors
    name: 'vpn-to-mikrotik'
    auto-negotiate: enable
    mode: tunnel
    src: 0:192.168.100.0/255.255.255.0:0
    dst: 0:192.168.200.0/255.255.255.0:0
    SA
      lifetime/rekey: 3600/2815   
      mtu: 1280
      tx-esp-seq: 1
      replay: enabled
      qat: 0
      inbound
        spi: 83391146
        enc:  aes-gc  fe03aea03aa08dee28c9a4ad2022d55113ea72be2078e0d5c77c2eacafe5ccb7d87e3611
        auth:   null  
      outbound
        spi: 09387f45
        enc:  aes-gc  6cb5756673b296440c9a0e428662abd80e8e444e97eb2ba8777fd2f9c9ff6a6e033a5f77
        auth:   null  
      NPU acceleration: none
    SA
      lifetime/rekey: 3600/2844   
      mtu: 1280
      tx-esp-seq: 1
      replay: enabled
      qat: 0
      inbound
        spi: 83391147
        enc:  aes-gc  7e4d12c5363b3aec6c29a9872e50ad65f4e1969e8dbc2db753853c620fd70e0448fd3339
        auth:   null  
      outbound
        spi: 03b2b08e
        enc:  aes-gc  034b284ccfdd7bdfb7b6aeb81f9c8f9a94fc5eb03e595fc0966011c46aa88b6637b06854
        auth:   null  
      NPU acceleration: none
```
For the Mikrotik RouterOS site the monitoring is done in the „IP –> IPsec –> Active Peers“ using Winbox or the WebUI:

You can get the same output using the commandline:

```
[admin@MikroTik] > /ip/ipsec/active-peers/print
Columns: ID, STATE, UPTIME, PH2-TOTAL, REMOTE-ADDRESS
# ID        STATE        UPTIME  PH2-TOTAL  REMOTE-ADDRESS
0 10.0.0.1  established  1m32s           1  10.0.0.1      
[admin@MikroTik] > 
```
External IP address of the VPN gateway 195.xx.31.xx

Local Encryption domain/ Accessible networks 10.xx.16.0/24

Appliance in use for setting up the VPN tunnel FortiGate 100F

IKE version IKEv2

Authentication PSK (exchanged via SMSs)

IKE encryption AES-256

IKE integrity SHA-256

IKE Phase 1 Mode Main Mode

DH group Group 16 (4096-bit)

IKE SA timeout 28800 seconds

IPSec Authentication ESP

IPSec ESP encryption AES-256

IPSec ESP integrity SHA-256

IPSec SA timeout 28800 seconds

Perfect forward secrecy Yes

Allowed traffic between peers (IP and ports, or any):

10.xx.16.0/24

Om jag fått ovan uppgifter, hur bör jag då ställa in det i Mikrotiken? Får ingen uppkoppling hur jag än gör känns det som.

salve se volessi far passare tutto il traffico internet tramite firewall, in pratica quello nella sede dove installo il mikrotik cosa devo fare

Hello,

I dont know if I understand your question correctly.

I you want to use a Mikrotik RouterOS Device as the firewall for all of your devices, then you should view this video:

https://www.youtube.com/watch?v=6boYA7xdjZY

Why is the VPN not accelerated by the NP?

the „diag vpn tunnel list“ shows „npu_flag=00“

Proposals int this example should be able to be offloaded on a 60F, imho?

Hi Tom,

interesting find!

I dont have this test environment running right now, therefore I cant check on this topic.

As a SOC4 platform the 60F includes a CP9Xlite and the only unsupported value from my configuration could be the diffie-hellman group 21 (ECC P-521).

In the hardware acceleration documentation only ECC P-256 is listed for the CP9 platform:

https://docs.fortinet.com/document/fortigate/7.2.8/hardware-acceleration/340357/cp9-capabilities

Hi, is it possible to share internet over ipsec tunnel from mikrotik to fortigate lan devices?

Hi,

just set the Phase 2 Selector / IPsec Policy to 0.0.0.0/0.

Best Regards

Dominik

Thank you for your answer. That was my first thought, unfortunately after entering such settings I lose access to both the internet and the network on the other side. Of course, I also modify the NAT access rules and masqurade (I’m not sure if masq is needed at all in this solution)

Hi,

I faced the same problem, but I need to route traffic only to specify sites. after adding Phase2 for IP-address of the site to Mikrotik I lost access to it. But if talking about local traffic – it works fine, problem only with external addresses.

Did you solve the issue?
