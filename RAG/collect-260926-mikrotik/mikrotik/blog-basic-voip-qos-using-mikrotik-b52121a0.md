---
id: collect-260926-mikrotik/mikrotik/blog-basic-voip-qos-using-mikrotik-b52121a0
title: "blog-basic-voip-qos-using-mikrotik-b52121a0"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2020-07-02"]
keywords: ["latency", "training", "voice"]
source: docs/RAG/lot-mikrotik/forum/qos/blog-basic-voip-qos-using-mikrotik-b52121a0.md
source_anchor: ""
source_lines: [1, 62]
sha256: cbe5789ec12e0481de9958fc49e1783c539de1fe2ffc412e24418d499652fc7b
---

# blog-basic-voip-qos-using-mikrotik-b52121a0

July 2, 2020            
                                                    
                    4                     min read                
                    
                    
                
  
VoIP is fast becoming the industry standard for office telecommunications and as more people work from home, it makes sense to ensure traffic is prioritised adequately to deliver voice packets smoothly. In this article, we have included a useful configuration guide to assist with Quality of Service best practices using MikroTik. VoIP does not require a large amount of bandwidth but it relies on low latency and jitter as voice communication is happening in real-time. Although QoS will help in many situations, it is not a silver bullet to solve VoIP quality problems on poor connections and is therefore not recommended to be used with high latency bandwidth types such as DSL or satellite.
Identify and Mark VoIP Traffic
Using MikroTik's 'Firewall Mangle', you will need to tell the router which types of connections are being used for VoIP, then mark each packet for later processing in Queues. In order to preserve resources on the router, it is best practice to mark each connection first, then mark packets within those connections. VoIP, in this case, makes use of UDP port 5060 (default) for registration along with random RTP ports for speech. It is important to check exactly which ports are being used in your situation before applying your configuration.
Step 1 - Identify SIP Connections
Step 2 - Mark All SIP Packets
Step 3 - Identify RTP Connections
Step 4 - Mark All RTP Packets
Differentiated Service Code Point (DSCP)
'Differentiated Services' is a default network protocol designed to classify and manage network traffic. It makes use of DSCP values to prioritise latency for certain traffic types over others.
| DSCP Name | DS Field Value (Dec) | IP Precedence (Description) | 
| CS0 | 0 | 0: Best Effort | 
| LE | 1 | n/a | 
| CS1, AF11-13 | 8, 10, 12, 14 | 1: Priority | 
| CS2, AF21-23 | 16, 18, 20, 22 | 2: Immediate | 
| CS3, AF31-33 | 24, 26, 28, 30 | 3: Flash - mainly used for voice signaling | 
| CS4, AF41-43 | 32, 34, 36, 38 | 4: Flash Override | 
| CS5, EF | 40, 46 | 5: Critical - mainly used for voice RTP | 
| CS6 | 48 | 6: Internetwork Control | 
| CS7 | 56 | 7: Network Control | 
In most cases, VoIP vendor equipment will use DSCP-46 (Expedited Forwarding) by default to prioritise RTP traffic. It is also possible to ensure that all RTP communications to your VoIP server are prioritised correctly with a Mangle rule.
Command Line Interface
/ip firewall mangle
add action=mark-connection chain=forward comment=SIP dst-address=1.2.3.4 dst-port=5060 \
new-connection-mark=SIP_Connection passthrough=yes protocol=udp
add action=mark-packet chain=forward connection-mark=SIP_Connection new-packet-mark=SIP_Packet \
passthrough=yes
add action=mark-connection chain=forward comment=RTP dst-address=1.2.3.4 new-connection-mark=\
RTP_Connection passthrough=yes port=10000-12000 protocol=udp
add action=mark-packet chain=forward connection-mark=RTP_Connection new-packet-mark=RTP_Packet \
passthrough=yes protocol=udp
add action=change-dscp chain=postrouting comment="DSCP Priority" dst-address=1.2.3.4 new-dscp=46 \
packet-mark=RTP_Packet passthrough=yes
Queues
Queues deal with network traffic by dropping packets once the pre-set maximum limits are reached. We recommend using 'Queue Tree'. 'Queue Tree' relies on packet marks as all traffic passes through at the same time. It will also allow control of both upstream and downstream traffic types independently. The idea here is to create a Parent queue which distributes bandwidth resources to subordinate or child queues based on limits and priority. How much bandwidth to allocate to these services really depends on the size of the codec and how many concurrent VoIP calls there are.
TIP: Check your connection tracking table for RTP marked connections during a VoIP call to get an idea of how much bandwidth you will need. We found when using the G729 codec, including overhead is approximately 32kbps up/down per call.
Create the Parent Queue with the LAN interface set as download, and the WAN interface set for upload. The priority for the Parent Queue is irrelevant.
Create Child Queues and attach to the relevant Parent with packet marks and priority. The lowest priority queue will not require a limit and will use all available bandwidth from the parent if available.
Command Line Interface
/queue tree
add limit-at=20M max-limit=20M name=Upload parent=pppoe-out1 queue=default
add limit-at=128k max-limit=128k name=Priority1-Up packet-mark=RTP_Packet parent=Upload priority=1 queue=\
default
add limit-at=20k max-limit=20k name=Priority2-Up packet-mark=SIP_Packet parent=Upload priority=2 queue=\
default
add name=Priority3-Up packet-mark=no-mark parent=Upload priority=3 queue=default
add limit-at=20M max-limit=20M name=Download parent=bridge queue=default
add limit-at=128k max-limit=128k name=Priority1-Down packet-mark=RTP_Packet parent=Download priority=1 \
queue=default
add limit-at=20k max-limit=20k name=Priority2-Down packet-mark=SIP_Packet parent=Download priority=2 \
queue=default
add name=Priority3-Down packet-mark=no-mark parent=Download priority=3 queue=default
Queue Tree List with Complete Configuration
TIP: Show the 'Dropped' column to display any dropped packets which will assist when troubleshooting problems. When troubleshooting, always check latency and jitter to the VoIP server.
Although this QoS technique is currently only dealing with VoIP, the same could be applied to other types of traffic. We trust this article provides a useful reference for you to improve and solve VoIP quality problems you may have been experiencing while adapting to the new reality of having to incorporate remote locations as part of your business or company.
If you are interested in learning more about MikroTik, consider attending certified training with us.
