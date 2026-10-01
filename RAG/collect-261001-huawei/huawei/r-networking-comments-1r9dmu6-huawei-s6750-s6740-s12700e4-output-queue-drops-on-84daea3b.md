---
id: collect-261001-huawei/huawei/r-networking-comments-1r9dmu6-huawei-s6750-s6740-s12700e4-output-queue-drops-on-84daea3b
title: "r-networking-comments-1r9dmu6-huawei-s6750-s6740-s12700e4-output-queue-drops-on-84daea3b"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/r-networking-comments-1r9dmu6-huawei-s6750-s6740-s12700e4-output-queue-drops-on-84daea3b.md
source_anchor: ""
source_lines: [1, 27]
sha256: cda7226a8af69b8ef43e8ed35d7ae3aa7b3bacbe667e6f27903abe301b8e4ca0
---

# r-networking-comments-1r9dmu6-huawei-s6750-s6740-s12700e4-output-queue-drops-on-84daea3b

Huawei S6750 / S6740 / S12700E4 Output Queue Drops on Asymmetric Links? 
        
        
        
    
    
    I’d like to bring up a discussion regarding some Huawei switches (S6750, S6740, and S12700E4).
I’ve noticed output queue drops (packet discards due to output queue congestion) in several customer deployments. The issue seems to occur particularly in scenarios involving asymmetric links for example: devices with LAG with 2x100G and individual 100G or 10G interfaces connected backbones.
log messages:
%%01LDP/4/HOLDTMREXP(l)[244]: Sessions were deleted because the hello hold timer expired. (PeerId=x.x.x.x, SessionState=Operational)
%%01IFPDT/4/INT_OUTBRDR(l)[253]: The output rate change ratio exceeded the threshold. (IfIndex=9, InterfaceName=100GE3/0/1, ThresholdPercent=50%, CurrentStatisticalPeriodRate=6275734122, LastStatisticalPeriodRate=3853445142)
%%01IFPDT/4/INT_OUTBRDR(l)[261]: The output rate change ratio exceeded the threshold. (IfIndex=10, InterfaceName=100GE3/0/2, ThresholdPercent=50%, CurrentStatisticalPeriodRate=6241037609, LastStatisticalPeriodRate=2506630628)
%%01IFPDT/4/INT_OUTBRDR(l)[262]: The output rate change ratio exceeded the threshold. (IfIndex=38, InterfaceName=100GE4/0/6, ThresholdPercent=50%, CurrentStatisticalPeriodRate=1065582990, LastStatisticalPeriodRate=4629143363)
%%01IFPDT/4/INT_OUTBRDR(l)[263]: The output rate change ratio exceeded the threshold. (IfIndex=40, InterfaceName=100GE4/0/18, ThresholdPercent=50%, CurrentStatisticalPeriodRate=1392543154, LastStatisticalPeriodRate=4484749441)
%%01IFPDT/4/INT_OUTBRDR(l)[265]: The output rate change ratio exceeded the threshold. (IfIndex=39, InterfaceName=100GE4/0/23, ThresholdPercent=50%, CurrentStatisticalPeriodRate=1983168388, LastStatisticalPeriodRate=5008893731)
In these cases, the devices appear to experience packet drops when traffic flows from higher-capacity aggregated links toward lower-capacity interfaces and viceversa. In some situations, these discards have even affected keepalive and hello packets for protocols such as OSPF, LDP, and BGP.
Has anyone else observed this behavior? Also, is there any way to resize or tune the buffers or output queues on this platform to mitigate the issue? Or could this be related to the network architecture?
*In the past, I’ve seen this issue on Arista switches in a data center environment with streaming servers. In that situation, I resolved the issue by resizing the buffers and output queues. After that, the customer decided to purchase switches with deep buffers.
I’d appreciate any insights or recommendations.
Thanks for all
Section des commentaires
Yeah that is normal, going from a higher speed port to a lower speed port is like trying to send the water from a fire hydrant through a drinking straw.
The problem only gets worse when you go from really fast to really slow ports like 100G to 10G, I doubt any switch in the world has buffers big enough to fix that. if you constantly have 11G of traffic trying to go into a 10G port buffers won't save you, you will have to drop some traffic and that is where QoS comes in.
There are some buffer settings you can tune, but 100G -> 10G as I said buffers probably won't save you. https://support.huawei.com/enterprise/en/doc/EDOC1100366621/ce77b50f/configuring-a-burst-traffic-buffering-mode https://support.huawei.com/enterprise/en/doc/EDOC1100366621/b9708599/configuring-the-buffer-size
But I would just recommend reading the whole QoS section https://support.huawei.com/enterprise/en/doc/EDOC1100213121/7ac0d32a/overview-of-qos here they have an example config using PQ-WDRR and WRED to determine what traffic gets dropped when congestion happens https://support.huawei.com/enterprise/en/doc/EDOC1100366621/1f52b66b/example-for-configuring-congestion-management
This is for the V100 and 200 but it gives a real quick and dirty explanation of these outbound discards, it is mainly focused on micro bursts but it applies more generally as well to traffic on a congested link. https://support.huawei.com/enterprise/en/doc/EDOC1000091883/52827a28/checking-the-discard-count-in-the-outbound-direction
u/irefox005, thanks for the information. I’ll read more about workaroung QOS on Huawei through the links you provided. Thanks again!
