---
id: collect-260926-mikrotik/mikrotik/questions-908972-how-to-set-up-qos-on-mikrotik-routeros-696faa74
title: "questions-908972-how-to-set-up-qos-on-mikrotik-routeros-696faa74"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-908972-how-to-set-up-qos-on-mikrotik-routeros-696faa74.md
source_anchor: ""
source_lines: [1, 58]
sha256: 642315ecc0693bf3dcb8dfc754d1a7d43d50ad781d6ef45448d0950a141bc784
---

# questions-908972-how-to-set-up-qos-on-mikrotik-routeros-696faa74

I'm aware that this is a common question but I've invested two days now to learn this and still could not find a clear explanation...
Recently I bought a Miktotik hEX (RouterOS 6, Level 4), used as the main router for a 100MBit down / 30MBit up fiber connection (speed tests show actually 110/33 as my provider adds a 10% margin). I want to prioritize traffic mainly to avoid issues with VoIP calls (via Twilio).
I had successfully prioritized traffic before with FireQOS using a custom Linux box (that could not handle more than 20Mbit, though) using the following config:
DEVICE=ppp0
INPUT_SPEED=20mbit
OUTPUT_SPEED=10mbit
LINKTYPE="local pppoe-llc"
interface $DEVICE world-in input rate $INPUT_SPEED
  class voip commit 120kbit       # https://www.twilio.com/docs/api/client/regions
    match src 54.171.127.192/26   # Twilio: ie1
    match src 52.215.127.0/24     # Twilio: ie1
    match src 35.156.191.128/25   # Twilio: de1
    match src 185.187.132.64/26   # Twilio: ie1-tnx
    match udp port 5060           # SIP
    match udp dports 10000:10100  # RTP
    match sports 3478,5349        # STUN
  class interactive commit 20%
    match udp port 53             # DNS
    match tcp port 22             # SSH
    match icmp                    # ping
    match tcp sports 5222,5228    # gtalk
  class synacks      
    match tcp syn    
    match tcp ack    
  class web commit 5%
    match tcp sports 80,443 
  class mail
    match tcp sports 25,465,587
  class default
interface $DEVICE world-out output rate $OUTPUT_SPEED
  class voip commit 120kbit
    match dst 54.171.127.192/26   # Twilio: ie1
    match dst 52.215.127.0/24     # Twilio: ie1
    match dst 35.156.191.128/25   # Twilio: de1
    match dst 185.187.132.64/26   # Twilio: ie1-tnx
    match udp port 5060           # SIP
    match udp sports 10000:10100  # RTP
    match dports 3478,5349        # STUN
  class interactive commit 20%
    match udp port 53             # DNS
    match tcp port 22             # SSH
    match icmp                    # ping
    match tcp dports 5222,5228    # gtalk
  class synacks commit 2%      
    match tcp syn    
    match tcp ack    
  class web commit 5%
    match tcp dports 80,443 
  class mail
    match tcp dports 25,465,587
  class default
Meanwhile I tried to set up packet marks accordingly, in RouterOS:
I think I need to set up Queue Trees on the Mikrotik and create a main queue for the available up/down speeds and divide by some priorities.
Currently I don't know where to start as I can't see a way to create two distinct queues for incoming and for outgoing traffic (as I have an asymmetrical bandwidth).
I see there are predefined (parent) queues for each interface but the problem is that I am using VLAN to provide WAN access to three completely independent subnets. The single physical port I'm using is ether2-master:
I guess creating a queue for parent ether2-master won't work as packets won't be marked on that level (right?), since all ether2-master packets are VLAN-tagged.
I'd like to have the same QoS as I had with FireQOS but what's really important is the VoIP part.
So, how should I configure the Queue Tree?
