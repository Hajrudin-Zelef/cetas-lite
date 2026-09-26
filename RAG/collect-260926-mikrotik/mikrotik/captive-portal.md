---
id: collect-260926-mikrotik/mikrotik/captive-portal
title: "captive-portal"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/captive-portal.md
source_anchor: ""
source_lines: [1, 27]
sha256: d8e963d1a9a0fd2328ce6167143012fb53014fd42f904a67036a06511d1526e0
---

# captive-portal

Good morning,

I have a MikroTik hAP lite and I use a captive portal to secure my wifi network. In the interfaces section of WinBox I have created a main network where I connect all my personal or home devices and another wlan interface as guest wifi. The problem is that I would like to put the captive portal only in the guest wifi not in the other interfaces as in the main wifi or in the LAN ports. How can I do? 

             
            
           
          
            
            
              Have you thought about creating a HotSpot for your guest Wi-Fi? You can attach that to a single interface (i.e. your guest one) and then configure your captive portal to the HotSpot service.

Example: /radius add service=hotspot address=radiusserveraddress secret=example

Would that be possible? I myself am new to MikroTik but from what I have read through in documentation, that might be a way to do it.

HotSpot: https://wiki.mikrotik.com/wiki/Manual:IP/Hotspot

RADIUS Client: https://wiki.mikrotik.com/wiki/Manual:RADIUS_Client

             
            
           
          
            
            
              Yes, I have created a single interface but since I am not very experienced with the mikrotik terminal, before doing some damage I would like to know if this thing can be done manually via WinBox and in which page I have to go.
