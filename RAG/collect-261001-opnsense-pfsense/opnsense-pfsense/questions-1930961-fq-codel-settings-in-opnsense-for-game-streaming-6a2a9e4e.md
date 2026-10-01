---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1930961-fq-codel-settings-in-opnsense-for-game-streaming-6a2a9e4e
title: "FQ_CoDel settings in OPNsense for Game Streaming"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1930961-fq-codel-settings-in-opnsense-for-game-streaming-6a2a9e4e.md
source_anchor: ""
source_lines: [1, 71]
sha256: c31278593a2a1e2ebbdb9761f1634ca15d5c4df2dbc401a61877004562dd3b9a
---

# FQ_CoDel settings in OPNsense for Game Streaming

*Score : 0 | Source : https://superuser.com/questions/1930961/fq-codel-settings-in-opnsense-for-game-streaming*

I'm having trouble understanding which values go where when trying to configure FQ_CoDel for game streaming in my OPNsense router. The Geforce Now documentation is clear about which ports they use and even that these should be used in traffic shaping but applying them is confusing me. TBH, traffic directions in the context of a router has never been obvious to me.
From Nvidia's documentation:
Here are the ports to add:
49003 – UDP Inbound AUDIO
49004 – UDP Outbound AUDIO
49005 – UDP Inbound VIDEO
49006 – TCP/UDP Outbound/Inbound Remote Input
Let's take the first as an example: When creating the shaping rules to match inbound port 49003, would I set that port as the source or destination port? Since it's inbound, I think this would be set on the WAN interface; would I use the LAN interface for an outbound port?
Here are the rules I've configured:
1: Interface: WAN
   Protocol: IP
   Source: any
   Src-port: any
   Destination: any
   Dst-port: 49006
   Direction: Out
   Target: UploadQueue_HighPriority
2: Interface: WAN
   Protocol: IP
   Source: any
   Src-port: 49006
   Destination: any
   Dst-port: any
   Direction: In
   Target: DownloadQueue_HighPriority
3: Interface: WAN
   Protocol: UDP
   Source: any
   Src-port: any
   Destination: any
   Dst-port: 49004
   Direction: Out
   Target: UploadQueue_HighPriority
4: Interface: WAN
   Protocol: UDP
   Source: any
   Src-port: 49003-49005
   Destination: any
   Dst-port: any
   Direction: In
   Target: DownloadQueue_HighPriority
The pipes and queues are set up mostly according to the OPNsense guide and a little bit of this one. I've just added high-priority upload and download queues with weights at 80 (the regular priority queues are set to 40).

---

### Reponse (acceptee) — score 2

Let me highlight a line from the documentation:
FQ_CoDel ignores the weight
The point of FQ_CoDel is that it provides fairness and balances flows using its internal algorithm; to this end, at least in the opnsense implementation, you cannot manually assign queue priorities/weights.
As an aside, the blog post you also linked looks heavily AI-formatted, so I'm not surprised it seems to have missed that this configuration would not work at all...
Thus you effectively need to make a choice of what you actually want here:
There is also the third option:
In your shoes, I would probably choose 1 in most situations. If game streaming is absolutely critical (say, you're streaming a tournament with thousands of viewers), then I would choose 3.
When we configure shaping for internet traffic, for simplicity it should all be configured on the WAN interface. If you follow the opnsense documentation, you'll notice it actually tells you how to configure the rules, notably you must (enable advanced mode via the toggle and then) set the direction.
Specifically, you want to configure:
| Queue | Interface | Direction | 
|---|---|---|
| Download | WAN | in | 
| Upload | WAN | out | 
If you still want to do specific rules for streaming, you should follow grawity's suggestion and perform a packet capture to confirm which ports are actually in use. GeForce documentation is unclear whether they are listing destination or source ports (though I would generally assume UDP ports to be destination...).
The rest of this is optional reading and a collection of what I discovered while trying to set up advanced shaping policies in a transition from EdgeOS to opnsense myself.
This quirk that you cannot apply prioritisation/weighting to FQ_CoDel queues/pipes is actually a bit of a opnsense/freebsd implementation limitation, but unfortunately not one you can easily get around. This is because opnsense uses freebsd dummynet for shaping. You are assigning the scheduler at the pipe level, and you cannot nest or otherwise relatively prioritise pipes (they must all be assigned a fixed bandwidth).
In theory I think pfsense's other shaping method with ALTQ is a bit more flexible in this sense, but I'm not certain since I've never run pfsense.
On the other hand, Linux's tc is far more flexible and fully capable of doing most of what you want here. By using the classful qdisc HTB, you can actually nest the equivalent of "pipe"s and assign relative priorities, and guaranteed minimum bandwidths. Then you can assign the classless FQ_CoDel qdisc as a nested inner level and it all works quite beautifully.
The problem we run into here is, well, currently there's no particularly good easy-to-use router distro for Linux. tc is a bit of a nightmare to configure manually. OpenWrt doesn't provide a GUI for this so you're left with scripting it. Ubiquiti EdgeOS had a quite nice Advanced Queues UI but that platform has been abandoned. Fortinet might(??) do it but is very expensive. VyOS can but has a questionable release model and is fully CLI, no GUI.
Long story short, Linux kernel/commandline has the flexibility to handle what you want but unless you want to dive deep into very advanced concepts and manual setup, you're probably better off just relaxing your requirements.
