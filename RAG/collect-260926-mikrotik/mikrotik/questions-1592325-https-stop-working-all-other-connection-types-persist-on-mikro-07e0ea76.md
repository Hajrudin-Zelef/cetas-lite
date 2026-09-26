---
id: collect-260926-mikrotik/mikrotik/questions-1592325-https-stop-working-all-other-connection-types-persist-on-mikro-07e0ea76
title: "questions-1592325-https-stop-working-all-other-connection-types-persist-on-mikro-07e0ea76"
domain: mikrotik
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-1592325-https-stop-working-all-other-connection-types-persist-on-mikro-07e0ea76.md
source_anchor: ""
source_lines: [1, 132]
sha256: 518a879a7e24413f52e3ae5fc4df0da90005d0dace9b44b0ccd70f2b46f97f4e
---

# questions-1592325-https-stop-working-all-other-connection-types-persist-on-mikro-07e0ea76

I am looking for help in HOW to diagnose a problem in my network (or suggestions for a solution if something seems obvious).
My setup:
provider: mobile broadband service through AT&T SIM card (LTE)
modem: Mikrotik wAP LTE (US) // RBwAPR-2nD&R11e-LTE-US
router: Ubiquiti EdgeRouter X SFP // ER‑X‑SFP
mgmt: Unifi Cloud Key // UC‑CK
clients: wireless and wired; windows, mac and linux (just me)
I'm almost always on a VPN (wireguard) throughout the day and am the only one on the LAN doing so.
My problem:
This network setup has been running fine for the last 6+ months. Over the last three weeks, all clients suddenly lose the ability to connect to the internet. This includes wired and wireless clients and all operating systems. The clients maintain their leases to the LAN.
Observations when it happens:
- all clients on the LAN cannot browse the internet except my machine, which is on a wireguard VPN
- Email ports/protocols appear to work fine as a Mac user used the Mail program's "troubleshooting" menu and connections appeared normal
- I can continue to browse the internet while I'm connected to the VPN
- However, when I disconnect from the VPN:
  - I then am unable to browse the internet as well
  - I can still ping google.com so DNS resolution does not appear to be the problem
  - curl seems to indicate that it is trying to do something with IPv6 (see below) however I have never enabled or configured the modem or router to use IPv6
curl when issue is active and I disconnect from VPN to experience it:
-> % curl -v google.com
*   Trying 216.58.195.78...
* TCP_NODELAY set
*   Trying 2607:f8b0:4005:807::200e...
* TCP_NODELAY set
* Immediate connect fail for 2607:f8b0:4005:807::200e: Network is unreachable
*   Trying 2607:f8b0:4005:807::200e...
* TCP_NODELAY set
* Immediate connect fail for 2607:f8b0:4005:807::200e: Network is unreachable
^C
...however I can still ping google.com
-> % ping google.com
PING google.com (216.58.195.78) 56(84) bytes of data.
64 bytes from sfo07s16-in-f78.1e100.net (216.58.195.78): icmp_seq=1 ttl=111 time=32.2 ms
curl normally (when not having issue):
-> % curl -v google.com
*   Trying 216.58.194.174...
* TCP_NODELAY set
* Connected to google.com (216.58.194.174) port 80 (#0)
> GET / HTTP/1.1
> Host: google.com
> User-Agent: curl/7.64.0
> Accept: */*
>
< HTTP/1.1 301 Moved Permanently
< Location: http://www.google.com/
< Content-Type: text/html; charset=UTF-8
< Date: Wed, 23 Sep 2020 18:11:24 GMT
< Expires: Fri, 23 Oct 2020 18:11:24 GMT
< Cache-Control: public, max-age=2592000
< Server: gws
< Content-Length: 219
< X-XSS-Protection: 0
< X-Frame-Options: SAMEORIGIN
<
<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>301 Moved</TITLE></HEAD><BODY>
<H1>301 Moved</H1>
The document has moved
<A HREF="http://www.google.com/">here</A>.
</BODY></HTML>
* Connection #0 to host google.com left intact
My Immediate Fix:
When this happens, I ssh onto the Mikrotik modem and restart it, which seems to restore things back to normal, but the problem can pop up again in a few hours or days (if lucky).
Obviously, constantly restarting the modem is not a long-term strategy.
I'm not exactly enamored with the Mikrotik RouterOS so I'm trying to determine the best way to attempt to capture more information in logs about what could be causing this issue so that I can figure out a better long-term fix.
My question(s):
If anyone has any immediate ideas as to why this might be happening, I'd love to hear them.
Otherwise I would love advice on where else to look on the modem and router to get some log information about these events.
Extra details on the router config. Pretty much all defaults.
$ ip route
default via 192.168.88.1 dev eth4 proto zebra 
192.168.1.0/24 dev switch0 proto kernel scope link src 192.168.1.1 
192.168.88.0/24 dev eth4 proto kernel scope link src 192.168.88.252
$ show configuration | grep -i ipv6
    ipv6-receive-redirects disable
    ipv6-src-route disable
Extra details on the modem config. Also all default (other than APN config).
The following default configuration has been installed on your router:
-------------------------------------------------------------------------------
LTE CPE Router with wireless AP:
 * lte interface connected to providers network (WAN port);
 * WAN port is protected by firewall and enabled DHCP client
LAN Configuration:
    IP address 192.168.88.1/24 is set on bridge (LAN port)
    DHCP Server: enabled;
    DNS: enabled;
wlan1 Configuration:
    mode:                ap-bridge;
    band:                2ghz-b/g/n;
    tx-chains:           0;1;
    rx-chains:           0;1;
    installation:        outdoor;
    ht-extension:        20/40mhz-XX;
WAN (gateway) Configuration:
    gateway:  lte1 ;
    ip4 firewall:  enabled;
    NAT:   enabled;
[admin@MikroTik] /interface lte> info lte1 once
           pin-status: ok
  registration-status: registered
        functionality: full
         manufacturer: MikroTik
                model: R11e-LTE-US
             revision: MPSS: R11eL_v12.09.174661 APSS: R11eL_v02.14.174662 CUSTAPP:
     current-operator: AT&T
    access-technology: Evolved 3G (LTE)
                 rssi: -71dBm
                 rsrp: -108dBm
                 rsrq: -14dB
[admin@MikroTik] /system routerboard> print
       routerboard: yes
        board-name: wAP R
             model: RBwAPR-2nD
     serial-number: ************
     firmware-type: qca9531L
  factory-firmware: 3.41
  current-firmware: 6.46.4
  upgrade-firmware: 6.46.4
[admin@MikroTik] /interface lte> print
Flags: X - disabled, R - running
 0  R name="lte1" mtu=1480 mac-address=**:**:**:**:**:** apn-profiles=att  network-mode=gsm,3g,lte
[admin@MikroTik] /interface lte apn> print
Flags: * - default
 0 * name="att" apn="broadband" use-peer-dns=yes add-default-route=yes default-route-distance=2
[admin@MikroTik] /ip address> print
Flags: X - disabled, I - invalid, D - dynamic
 #   ADDRESS            NETWORK         INTERFACE
 0   ;;; defconf
     192.168.88.1/24    192.168.88.0    bridge
 1 D **.**.**.**/32     **.**.**.**     lte1
edit:
This NetworkEngineering question appears to describe a similar situation, but uses Fortigate hardware.
