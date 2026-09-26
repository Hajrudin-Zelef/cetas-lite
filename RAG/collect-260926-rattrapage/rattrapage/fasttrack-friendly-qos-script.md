---
id: collect-260926-rattrapage/rattrapage/fasttrack-friendly-qos-script
title: "this is based on IntrusDave's QoS script, but modified"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "voice"]
source: docs/RAG/lot-rattrapage/servers-reviews/fasttrack-friendly-qos-script.md
source_anchor: ""
source_lines: [1, 111]
sha256: 0bec8c20c244ff998b5916f6d60b16212c920dc2dd5d279fc9dfb0ba6fee3db3
---

# this is based on IntrusDave's QoS script, but modified

Hi all,

Here is my FastTrack-friendly QoS script, based on one made by IntrusDave. It supports IPv4 and IPv6, and uses DSCP markings to match packets. It is fasttrack-friendly, but does not require fasttrack.

The mappings of IP precedence values to queues is in line with Cisco’s AVVID 802.1p UP-Based Traffic Types, used for wireless networks, and compatible with wireless priority from MikroTik and other vendors

To retain compatibility with FastTrack, it uses interface-attached HTB rather than global-attached (fasttracked packets bypass global HTB, but are still queued by interface HTB). The inbound WAN (download) rate queueing is implemented as outbound queueing on the internal interface, since interface-attached HTB only works for egress. This is fine for many home-use scenarios where there is only one LAN subnet and only one WAN subnet.

With this solution, any traffic that is left at DSCP 0 (Best Effort) will be fasttracked; other traffic will need to bypass fasttrack to get proper priority. Of course, in many cases, most traffic will not be tagged already. If you need to manually tag traffic, this can be done by bypassing fasttrack on only the packets you need to tag (by creating an ‘accept’ rule for that traffic (also matching connection-state=established,related) just above the fasttrack rule). Then you can add mangle rules to set DSCP tags on those packets, and place those mangle rules at the top of the list, above the other mangle rules created by this script.

You’ll need to copy and paste the script into scripts (system->scripts->new (+)->paste), change the upload and download bandwidth and inbound and outbound interface names at the top to match your settings, and run the script. (the bandwidths should be slightly less than what you normally receive as your maximum)

```
# this is based on IntrusDave's QoS script, but modified
# qosClasses are largely based on Cisco Wireless QoS mappings/guide
#Set outbound (WAN) interface here
:local outboundInterface "ether1"
#Set UPLOAD bandwidth of the outbound (WAN) interface
:local outInterfaceBandwidth 4900k
#Set inbound (LAN) interface here
:local inboundInterface "bridge"
#Set DOWNLOAD bandwidth of the outbound (WAN) interface
:local inInterfaceBandwidth 34500k
#Set type of queue here
:local queueType wireless-default
#Set where in the chain the packets should be mangled
:local mangleChain postrouting
#Don't mess with these. They set the parameters for what is to follow
:local queueName ("QoS_" . $outboundInterface)
:local inQueueName ("QoS_" . $inboundInterface)
# qosClasses from highest to lowest priority
:local qosClasses [:toarray "Network Control (Top Priority),Internetwork Control (High Priority),Voice (Medium-High Priority),Interactive Video (Medium Priority),Critical Data or Call Signaling (Medium-Low Priority),Best Effort (Low Priority),Background (Very Low Priority),Scavenger (Bottom Priority)"]
# maps queue priorities from highest to lowest to IP precedence values
:local priorityToIpPrecedenceMappings [:toarray "7,6,5,4,3,0,2,1"]
# queue priority used for best effort traffic (IP precedence 0)
:local beQueuePriority 6
/ip firewall mangle add action=set-priority \
   chain=postrouting new-priority=from-dscp-high-3-bits \
   passthrough=yes comment="Respect DSCP tagging"
/ip firewall mangle add action=set-priority \
   chain=postrouting new-priority=6 packet-size=0-123 \
   passthrough=yes protocol=tcp tcp-flags=ack comment="Prioritize ACKs"
/ip firewall mangle add action=accept \
   chain=postrouting priority=0 \
   comment="IP Precedence (aka Packet Priority) 0 - Best Effort (Low Priority) (default)"
:for indexA from 1 to 7 do={
    :local qosIndex (7-$indexA)
    # skip best effort in list
    :if ($indexA <= (8-$beQueuePriority)) do={ :set qosIndex (8-$indexA) }
    :local subClass ([:pick $qosClasses $qosIndex] )
    /ip firewall mangle add action=mark-packet chain=$mangleChain comment=("IP Precedence (aka Packet Priority) " . $indexA . " - " . $subClass . " (apply packet mark ip_precedence_" . $indexA . ")") \
         disabled=no priority=($indexA) new-packet-mark=("ip_precedence_" . $indexA) passthrough=no
}
:if ([/system package find name=ipv6 disabled=no] = "") do={
    :log info "IPv6 package is not installed - skipping IPv6 mangle rules";
} else={
   :for dscpValue from 0 to 7 do={
   /ipv6 firewall mangle add action=accept \
      chain=postrouting dscp=$dscpValue \
      comment="IP Precedence 0 (DSCP $dscpValue) - Best Effort (Low Priority) (default)"
   }
   :for indexA from 1 to 7 do={
       :local qosIndex (7-$indexA)
       # skip best effort in list
       :if ($indexA <= (8-$beQueuePriority)) do={ :set qosIndex (8-$indexA) }
       :local subClass ([:pick $qosClasses $qosIndex] )
       :for dscpValue from ($indexA*8) to (($indexA*8)+7) do={
       /ipv6 firewall mangle add action=mark-packet chain=$mangleChain comment=("IP Precedence " . $indexA . " (DSCP " . $dscpValue . ") - " . $subClass . " (apply packet mark ip_precedence_" . $indexA . ")") \
            disabled=no dscp=$dscpValue new-packet-mark=("ip_precedence_" . $indexA) passthrough=no
       }
   }
}
/queue tree add max-limit=$outInterfaceBandwidth name=$queueName parent=$outboundInterface comment="Uplink QoS" queue=$queueType
:for queuePriority from=1 to=8 do={
   :local qosIndex ($queuePriority-1)
   :local subClass ([:pick $qosClasses $qosIndex] )
   :local ipPrecedence ([:pick $priorityToIpPrecedenceMappings $qosIndex])
   :local ipPrecedenceMark ("ip_precedence_" . $ipPrecedence)
   :if ($ipPrecedence = "0") do={ :set ipPrecedenceMark ("no-mark") }
   /queue tree add \ 
      name=("IP Precedence " . $ipPrecedence . ". " . $subClass . " - " . $outboundInterface ) \
      parent=$queueName \
      priority=($queuePriority) \
      queue=$queueType \
      packet-mark=$ipPrecedenceMark \
      comment=("Queue Priority " . $queuePriority)
}
/queue tree add max-limit=$inInterfaceBandwidth name=$inQueueName parent=$inboundInterface comment="Downlink QoS" queue=$queueType
:for queuePriority from=1 to=8 do={
   :local qosIndex ($queuePriority-1)
   :local subClass ([:pick $qosClasses $qosIndex] )
   :local ipPrecedence ([:pick $priorityToIpPrecedenceMappings $qosIndex])
   :local ipPrecedenceMark ("ip_precedence_" . $ipPrecedence)
   :if ($ipPrecedence = "0") do={ :set ipPrecedenceMark ("no-mark") }
   /queue tree add \ 
      name=("IP Precedence " . $ipPrecedence . ". " . $subClass . " - " . $inboundInterface ) \
      parent=$inQueueName \
      priority=($queuePriority) \
      queue=$queueType \
      packet-mark=$ipPrecedenceMark \
      comment=("Queue Priority " . $queuePriority)
}
```

Then, if you are using fasttrack, to allow DSCP-tagged packets to be classified properly, bypass fast-track for those packets by adding the following:

```
/ip firewall filter add action=accept chain=forward comment="Bypass fasttrack for non-zero DSCP" connection-state=established,related dscp=!0
```

**Move that filter just *BEFORE* your fast-track rule!** All DSCP-marked packets will then be queued appropriately, and DSCP 0 packets fast-tracked. Fast-tracked packets are all placed in the best effort queue.
