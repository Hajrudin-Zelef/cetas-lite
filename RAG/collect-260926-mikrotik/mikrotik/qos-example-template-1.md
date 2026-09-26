---
id: collect-260926-mikrotik/mikrotik/qos-example-template-1
title: "qos-example-template"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/qos/qos-example-template.md
source_anchor: ""
source_lines: [1, 145]
sha256: 8b29a6a3e03099f90e6909aa1002939dc36d686c0a5174e3aca2094bc6cd372b
---

# qos-example-template

Hello guys,

does anybody of you have some template for QoS settings? I found plenty of manuals but it’s still not clear how to set prioritization of data flow or why it’s so complicated.

My use case is easy, I just want to have QoS on my WAN port and to be able to use all my devices for: download/upload large files, watch Netflix, stream Spotify and browse pages.

So nothing special - Ideally to have QoS according to services which I’m currently using.

Thanks for any info.

Marek

             
            
           
          
            
            
              Hello, the qos you have to adapt to your needs, there are in the forum many related topics and there have also been presentations in the mum here I share a link to you go having some concepts and examples.

https://mum.mikrotik.com/presentations/CZ09/QoS_Megis.pdf

https://mum.mikrotik.com/presentations/US08/janism.pdf

https://mum.mikrotik.com/presentations/IT14/giordano.pdf

             
            
           
          
            
            
              Hello I already went through those links but still it’s not really clear how to do prioritization of servises.

Typicaly Youtube, HBO GO, Netflix, Spotify and other streaming services - do I really identify IP addresses in my connections? Why it’s so complicated?

If you imagine, you have 1 Mikrotik router and 1 PC connected to this router. You will start uploading huge files to fast servers and then you aren’t able to use this PC for browsing, because everything is consumed by upload.

There is no simple solution?

             
            
           
          
            
            
              The simple solution is to prioritize by DSCP (TOS) value.  There is a script on this forum that does it automatically.

See this topic: http://forum.mikrotik.com/t/fasttrack-friendly-qos-script/102401/1

This works OK for applications like VoIP because the writers of those applications usually set the right DSCP value in their packets.

When you have other applications that do not do that correctly, it is the fault of those applications, not of the mechanism.

However, when you merely setup a queue tree as done in the referenced topic, and you make sure your max speed is slightly below the real upload speed of your connection, you already improve matters a lot.

             
            
           
          
            
            
              
It isn’t complicated. You don’t understand the problem. A router makes decisions based on the contents of the IP header. You are wanting a device to make decisions based on websites and services which are not explicitly in IP headers. While some routers attempt to do what you want, they aren’t good at it (although their marketing departments think they are awesome at it).

The problem is that the Netflix show you want to stream might be hosted at cdn-server83432.thelocalisp.akamai.net. The router has no way to know that the DNS request your computer sent for cdn-server83432.thelocalisp.akamai.net should be part of a prioritized traffic flow.

The other problem is that you are trying to manage bandwidth utilization of inbound (to the router) traffic flows. It is difficult to effectively mange bandwidth utilization on a link in an inbound direction. You have to create queues which prioritize traffic on the egress interface (probably a bridge interface in this case), and that is only moderately effective because the sender can still easily overload the bandwidth-limited link, resulting in packet loss anyway.

The first question to ask is what problem are you trying to solve? Are you experiencing packet loss under certain conditions? What are those conditions? Perhaps rather than prioritizing streaming flows, you could lower the priority of other known flows.

             
            
           
          
            
            
              
That is a very important point.  One aspect of QoS is congestion management, when there is more data to be sent than the bandwidth can cope with.  Under these conditions QoS is essentially deciding what data should be sent and what is not sent.  It can’t directly do this at the receiving end.

             
            
           
          
            
            
              While that is true, the issue of “not being able to view streams while something is being uploaded” is actually not really an inbound QoS problem, it is more of an outbound problem.

Not really QoS usually, it is caused by buffer bloat (the outbound router doing way too much buffering, that buffer is entirely filled with upload packets, and the ACKs on downloads are sent with a lot of delay causing problems in the download direction).

Merely setting up a queue tree on the outbound interface with a max speed less than the actual max speed of the uplink will already help in that case.

Having different priorities on that queue tree can help when e.g. VoIP is in use, but even the local queueing in the router with different queues for different traffic helps most.

E.g. set a “per connection queue” type for the outgoing queue (pcq-upload-default), or “fair queue” (wireless-default).  That will mix the ACKs between the upload traffic.

             
            
           
          
            
            
              There’s is no simple solution because QoS is a complicated topic. People have been working on this for decades. Routers that promise easy one click solutions don’t work that well, otherwise everyone would include magic solutions on their hardware. CoDel and Cake have improved things for ease of use tremendously, but they aren’t perfect either. CoDel and Cake are being tested on MT (check the development thread).

On MikroTik, a single SFQ queue can mimic some of the benefits of CoDel with regards to fair bandwidth sharing and bufferbloat. Other queue types and combinations of prioritization strategies will require a little more time and effort. However, the more you try to control packets, the more resources will be needed. Many Internet users also forget that service providers have a say on how packets are sent to them. My ISP can throttle everything above 100M at night, regardless of the traffic type or the plan used, and once a buffer is full, packets are dropped.

It is possible to limit Netflix and other streaming sites on MT, but using DSCP is much more efficient. There’s good info on YT about both methods and much more. Note that prioritizing packets that you don’t want delayed or dropped is usually the best strategy (VOIP, conferencing apps), instead of trying to prioritize every packet type with dozens of needless rules.

             
            
           
          
            
            
              Hello guys, thanks for your reply, I appreciate and sorry for late reply.

I tried to set something as you recommend and result is:

1. I can use more devices on my network to download/upload streams or files
2. it didn’t solve my main issue which I’ll try to describe below:
 As many of you recommend to me to set Max limit - this will solve first point only - more devices could be connected with no issue while I’m uploading from my PC.
Issue I’m facing is, that while I’m uploading data I can’t use this PC, because again I will use whole my upload limit defined by my QoS settings in Mikrotik.

Please see my configuration in attachment.

Thank you, Marek

router_export.rsc (6.27 KB)

             
            
           
          
            
            
              Make sure your upload program in the PC uses a lower priority (DSCP 8 or 16) than the default traffic (normally DSCP 0).

Then use one of the QoS methods that use DSCP (high 3 bits) to determine priority.

             
            
           
          
            
            
