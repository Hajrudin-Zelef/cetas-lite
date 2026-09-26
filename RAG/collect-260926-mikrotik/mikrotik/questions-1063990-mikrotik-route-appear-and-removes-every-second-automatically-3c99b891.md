---
id: collect-260926-mikrotik/mikrotik/questions-1063990-mikrotik-route-appear-and-removes-every-second-automatically-3c99b891
title: "ADDRESS            NETWORK         INTERFACE"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1063990-mikrotik-route-appear-and-removes-every-second-automatically-3c99b891.md
source_anchor: ""
source_lines: [1, 41]
sha256: e003956590ea60a4fe172d18ff4daa1baaff4a873d3daea491dc9bf633501091
---

# ADDRESS            NETWORK         INTERFACE

Mikrotik OS v6.48.2 has routes
 #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
 0 ADS  0.0.0.0/0                          193.168.1.1               1   <<<<<<<<<<<<<<<<<<<<<
 1   S  ;;; dyno
        0.0.0.0/0                          192.168.8.1               5
 2   S  ;;; dyno
        0.0.0.0/0                          pppoe-out1                6
 3   S  ;;; dyno
        0.0.0.0/0                          193.168.1.1              10
 4 ADC  10.32.181.1/32     x.x.x.147   pppoe-out1                0
 5 ADC  10.32.238.1/32     x.x.x.250   pppoe-out2                0
 6 A S  ;;; dyno
        x.x.x.18/32                 193.168.1.1               1
 7 ADC  192.168.8.0/24     192.168.8.100   lte1                      0
 8 ADC  192.168.88.0/24    192.168.88.1    bridge                    0
 9 ADC  193.168.0.0/16     193.168.0.177   ether1                    0
But route 0 dissapears and appear again every 3-5 second. So it becomes like
 #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
 0   S  ;;; dyno
        0.0.0.0/0                          192.168.8.1               5
 1 A S  ;;; dyno
        0.0.0.0/0                          pppoe-out1                6
 2   S  ;;; dyno
        0.0.0.0/0                          193.168.1.1              10
 3 ADC  10.32.181.1/32     x.x.x.147   pppoe-out1                0
 4 ADC  10.32.238.1/32     x.x.x.250   pppoe-out2                0
 5 A S  ;;; dyno
        x.x.x.18/32                 193.168.1.1               1
 6 ADC  192.168.8.0/24     192.168.8.100   lte1                      0
 7 ADC  192.168.88.0/24    192.168.88.1    bridge                    0
 8 ADC  193.168.0.0/16     193.168.0.177   ether1                    0
Nothing with LAN changed. No scripts are running. Every 3-5 seconds it again appears and disappears.
/ip address print
#   ADDRESS            NETWORK         INTERFACE                                                  
 0   ;;; defconf
     192.168.88.1/24    192.168.88.0    bridge                                                     
 1 D 193.168.0.177/16   193.168.0.0     ether1                                                     
 2 D 192.168.8.100/24   192.168.8.0     lte1                                                       
 3 D x.x.x.250/32   10.32.238.1     pppoe-out2                                                 
 4 D x.x.x.147/32   10.32.181.1     pppoe-out1 
How to solve this?
