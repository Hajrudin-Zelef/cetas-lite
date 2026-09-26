---
id: collect-260926-mikrotik/mikrotik/questions-1364065-mikrotik-hex-router-how-to-limit-bandwidth-on-ethernet-c83295d3
title: "questions-1364065-mikrotik-hex-router-how-to-limit-bandwidth-on-ethernet-c83295d3"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1364065-mikrotik-hex-router-how-to-limit-bandwidth-on-ethernet-c83295d3.md
source_anchor: ""
source_lines: [1, 72]
sha256: 662b866029dfd61b7a7560e2d170001e5899926d46eebc1a782ed0b377d212ee
---

# questions-1364065-mikrotik-hex-router-how-to-limit-bandwidth-on-ethernet-c83295d3

I have a MikroTik hEX router running RouterOS 6.4.  eth1 is my WAN port.
eth2 is connected to a TP-Link Wi-Fi router.  All the devices on this Wi-Fi network have addresses in the range 192.168.42.100 - 192.168.42.200.  The Wi-Fi gateway on this router is 192.168.42.254 and the eth1 on the TP-Link router is connected to the MikroTik with an address of 192.168.88.10.
eth3 is connected to a D-Link Wi-Fi router.  All the devices on this Wi-Fi network have addresses in the range 192.168.84.100 - 192.168.84.132.  The Wi-Fi gateway on this router is 192.168.84.254 and the eth1 on the D-Link router is connected to the MikroTik with an address of 192.168.88.11.
All the ports on the MikroTik are in a single bridge so that all the devices on the network (from the two different Wi-Fi routers) can see each other.
I have disabled FastForward on the bridge and I have enabled ip firewall on the bridge. I have also disabled the default FastForward firewall rule.
I have a simple queue which specified the bandwidth limit for IP addresses in the range 192.168.84.0/24.  This queue is not working.
I have tried to remove eth3 on the MikroTik from the bridge – but when I do this, the MikroTik nearly dies – the device becomes unresponsive as if the CPU (or something) is working extremely hard.
Here is a list of the relevant configurations.  All I want to do is limit the up/down load bandwidth on devices coming from eth3.
From what I have read on-line it probably has something to do with the bridge – but, as I have said – when I remove eth3 from the bridge – the device becomes un-responsive.
/interface print 
Flags: D - dynamic, X - disabled, R - running, S - slave 
 #     NAME                                TYPE       ACTUAL-MTU L2MTU  MAX-L2MTU MAC-ADDRESS      
 0  R  ether1-wan                          ether            1500  1596       2026 CC:2D:E0:68:E3:EE
 1  RS ether2-gw                           ether            1500  1596       2026 CC:2D:E0:68:E3:EF
 2  RS ether3-guest                        ether            1500  1596       2026 CC:2D:E0:68:E3:F0
 3   S ether4-vpn                          ether            1500  1596       2026 CC:2D:E0:68:E3:F1
 4   S ether5-direct                       ether            1500  1596       2026 CC:2D:E0:68:E3:F2
 5  R  ;;; defconf
       bridge                              bridge           1500  1596            CC:2D:E0:68:E3:EF
 6  R  pppoe-out1                          pppoe-out        1480
/interface bridge print 
Flags: X - disabled, R - running 
 0 R ;;; defconf
     name="bridge" mtu=auto actual-mtu=1500 l2mtu=1596 arp=enabled arp-timeout=auto mac-address=CC:2D:E0:68:E3:EF protocol-mode=rstp fast-forward=no igmp-snooping=no auto-mac=no admin-mac=CC:2D:E0:68:E3:EF 
     ageing-time=5m priority=0x8000 max-message-age=20s forward-delay=15s transmit-hold-count=6 vlan-filtering=no dhcp-snooping=no 
/ip address print 
Flags: X - disabled, I - invalid, D - dynamic 
 #   ADDRESS            NETWORK         INTERFACE
 0   192.168.88.2/24    192.168.88.0    ether2-gw
 1   192.168.88.5/24    192.168.88.0    ether5-direct
 2 D xxx.xxx.xx.x/32    xxx.xxx.xx.x    pppoe-out1
 3   192.168.88.3/24    192.168.88.0    ether3-guest
 4   192.168.88.4/24    192.168.88.0    ether4-vpn
/queue simple print 
Flags: X - disabled, I - invalid, D - dynamic 
 0    name="quest-limit" target=192.168.84.0/24 parent=none packet-marks="" priority=8/8 queue=default-small/default-small limit-at=0/0 max-limit=512k/2M burst-limit=0/0 burst-threshold=0/0 burst-time=0s/0s 
      bucket-size=0.1/0.1
/queue simple print 
Flags: X - disabled, I - invalid, D - dynamic 
 0    name="quest-limit" target=192.168.84.0/24 parent=none packet-marks="" priority=8/8 queue=default-small/default-small limit-at=0/0 max-limit=512k/2M burst-limit=0/0 burst-threshold=0/0 burst-time=0s/0s 
      bucket-size=0.1/0.1
/ip route print 
Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, B - blackhole, U - unreachable, P - prohibit 
 #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
 0 ADS  0.0.0.0/0                          pppoe-out1                1
 1   S  192.168.21.0/24                    ether4-vpn                1
 2 A S  192.168.42.0/24                    ether2-gw                 1
 3 A S  192.168.84.0/24                    ether3-guest              1
 4 ADC  192.168.88.0/24    192.168.88.2    bridge                    0
 5 ADC  xxx.xxx.xx.x/32    xxx.xxx.xx.x   pppoe-out1           0
/ip firewall> filter print 
Flags: X - disabled, I - invalid, D - dynamic 
 0    ;;; defconf: accept established,related,untracked
      chain=input action=accept connection-state=established,related,untracked 
 1    ;;; defconf: drop invalid
      chain=input action=drop connection-state=invalid 
 2    ;;; defconf: accept ICMP
      chain=input action=accept protocol=icmp 
 3    ;;; defconf: drop all not coming from LAN
      chain=input action=drop in-interface-list=!LAN 
 4    ;;; defconf: accept in ipsec policy
      chain=forward action=accept ipsec-policy=in,ipsec 
 5    ;;; defconf: accept out ipsec policy
      chain=forward action=accept ipsec-policy=out,ipsec 
 6 X  ;;; defconf: fasttrack
      chain=forward action=fasttrack-connection connection-state=established,related log=no log-prefix="" 
 7    ;;; defconf: accept established,related, untracked
      chain=forward action=accept connection-state=established,related,untracked 
 8    ;;; defconf: drop invalid
      chain=forward action=drop connection-state=invalid 
 9    ;;; defconf:  drop all from WAN not DSTNATed
      chain=forward action=drop connection-state=new connection-nat-state=!dstnat in-interface-list=WAN
