---
id: collect-261001-huawei/huawei/questions-18315-which-qos-classification-should-be-applied-to-ntp-3935edfb
title: "questions-18315-which-qos-classification-should-be-applied-to-ntp-3935edfb"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-huawei/questions-18315-which-qos-classification-should-be-applied-to-ntp-3935edfb.md
source_anchor: ""
source_lines: [1, 15]
sha256: 115de7d867325f28684d6d32074870bdc15a93b197291f1151f6416d095309d2
---

# questions-18315-which-qos-classification-should-be-applied-to-ntp-3935edfb

NTP isn't particularly jitter-sensitive because it uses originate and transmit timestamps to keep track of delay.  Ntp.org explains in detail how it keeps delay in check, but here's a snippet:
  Synchronizing a client to a network server consists of several packet
  exchanges where each exchange is a pair of request and reply. When
  sending out a request, the client stores its own time (originate
  timestamp) into the packet being sent. When a server receives such a
  packet, it will in turn store its own time (receive timestamp) into
  the packet, and the packet will be returned after putting a transmit
  timestamp into the packet. When receiving the reply, the receiver will
  once more log its own receipt time to estimate the traveling time of
  the packet. The traveling time (delay) is estimated to be half of "the
  total delay minus remote processing time", assuming symmetrical
  delays.
The reason this isn't in the same category as network control is because this isn't directly responsible for the operation of routing/forwarding of packets.  All of the things in the network management category are not critical components of the networking system as a whole.  If you lost any packets related to SNMP, syslog, or NTP, you likely wouldn't even notice.  
SNMP would simply retransmit that information since it's TCP based.  Even if the connection dropped all together, nothing catastrophic would happen; you might just get a snmp agent not responding and then try again.  If you lost syslog traffic (UDP), you would simply lose a blip of logging information, which is probably still contained within the buffer or in a log file on the device.  Since NTP calculates delay based on previous packets, while also accounting for the maximum offset error, you really aren't running into any issues.  Worst case scenario, your time drifts by a few picoseconds…
If you lost a packet related to routing, even for a second, you may be facing the entire system going down; rendering any other markings worthless.  At that point, NTP would simply fall totally out of sync and rely on it's local ticker to keep time.
