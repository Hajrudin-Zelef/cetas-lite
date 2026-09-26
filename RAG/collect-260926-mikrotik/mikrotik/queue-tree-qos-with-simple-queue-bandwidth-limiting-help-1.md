---
id: collect-260926-mikrotik/mikrotik/queue-tree-qos-with-simple-queue-bandwidth-limiting-help-1
title: "queue-tree-qos-with-simple-queue-bandwidth-limiting-help"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["training"]
source: docs/RAG/lot-mikrotik/forum/qos/queue-tree-qos-with-simple-queue-bandwidth-limiting-help.md
source_anchor: ""
source_lines: [1, 195]
sha256: f75f8c5cd27869547b59c463732affce590280645b1f30408746447fec006b90
---

# queue-tree-qos-with-simple-queue-bandwidth-limiting-help

I need some help.  I am trying to perform QOS using mangle+queue trees and Bandwidth limiting per user source address using Simple Queues.  Below is a description and diagram of what I would like to achieve.




DESCRIPTION:

- 
I am planning to prioritize VoIP, HTTP, FTP, etc using mangling in prerouting. Then using Queue trees, I will create a main PCQ Queue with a 1M/1M upload/download limit (I need help in determining where to apply it).
- 
Next, I will create separate sub queues (Main PCQ is parent) for VoIP, HTTP, FTP, etc and assign specific bandwidth limits per category and set corresponding Priotities.
- 
I will then create a main a “Main Simple Q” of type PCQ and apply it to all interfaces.
- 
Next, I will create separate simple queues for each user with bandwidth limits with the parent being the “Main Simple Q”

DIAGRAM(underscore is spaces):

Local Interface ------Prioritize--------Limit---------------Shape-----------------Public interface

Pre Route Mangle __PCQ Tree____Simple Queues

VoIP------------- Priority 1|________10.0.0.10 --|

HTTP------------- Priority 2|–Main----10.0.0.11 --|

FTP -------------- Priority 3|  PCQ ___10.0.0.12 --|–Main Simple Q -------

P2P ------------- Priority 4|________10.0.0.13 --|  (public 1M/1M limit)

____________________________10.0.0.14 --|  (Applied to All INTs)



Questions:

1. Will this work?
2. In the “Prioritize” phase, what should be the parent of the “Main PCQ”?  Should it be the local Interface?
3. In the “Shape” phase, if I set the “Main Simple Q” type to PCQ,will it shape all traffic by users by IP?  Meaning, will it split the bandwidth evenly amongst the active users if congestions is occuring?
4. Anything that I am doing incorrectly?

 
            
           
          
          
            
            
              Yeah, I have read that but that doesn’t answer my implementation questions.  Thus the pictures and specific questions I was hoping someone could help me with.

             
            
           
          
            
            
              Well, I folow this forum for several years and I never saw descent help regarding QoS. What you asked now, I saw asked several times before, with no proper help.

             
            
           
          
            
            
              http://forum.mikrotik.com//viewtopic.php?t=12824&start=50

this is the best you can have last few posts is what you are asking at least it is the best solution for now.

Developers of MikroTik should try to help us a little bit  .

             
            
           
          
            
            
              Thank Titus.  Yeah, that is an interesting thread as well.   That would work well for load balancing all connections.  On top of that, I want to prioritze connection types though.  Everyone on the board indicates QOS is complicated.  I don’t think its that complicated.  I just am having an issue in trying to build the configuration that will support the example I noted at the top of the thread.   If we could get some complex QOS examples posted in the WIKI, it would be much easier to figure out the correct MT config.   I get the feeling the only way to get the answer is to pay someone to build the config and then you can look at it.   

             
            
           
          
            
            
              if you understand networking, and why and what is happening, and if you understand clearly what is happening in MT ROS packet flow you should not problems.

but, as usual, many users what to jump in QoS train when it is leaving the station 

in manual, you can get the basics. if that is not enough for succesful congiration i would suggest you to attend training in QoS

             
            
           
          
            
            
              Hi Janisk.    There really isn’t enough info to understand the packetflow with regard to the MT OS implementation.  That’s why there are so many questions.   The very basic examples in the manual are not enough.  Folks have been asking for more info on QOS for a long time.   It would be nice if we could get some.  

I have clearly asked some simple questions at the top of this thread.   No answers.  

             
            
           
          
            
            
              if you want to use queue tree you have to mangle traffic, if you can distinguish witch packets is what protocol you mark these packets - that is very basic.

you can implement as many packet marks you like to. as result you can have different queue tree elements to limit each traffic and prioritize it . as simple as that.

in titius post that link uses two packet marks and you can clearly see what have to be done.

just add different (what you need) markings and off you go  also, you have to read about what queue type you need to get results you want to.

toughest part is to mark everything as you imagined not as you got it 

scenario:

1. mark all the traffic add queue tree a rule that limits that traffic to size you want to
2. mark http traffic and give it priority you need and limitations. take in account all http downloads are downloads, you are browsing net - you download, you use filezilla to download from http - it is same type of download, just bigger is size. add corresponding queue tree branches
3. continue with marking packets so all traffic mark has less and less packets. in mangle when marking set to passthrough= no. you set mangle rules in manner that rule that marks all traffic stays last. and all the other rules are in front of it.
4. when you add any rle check if that is working as you wish.
 
 
EDIT:

in my interest is that you find your answers yourself, so you will have deeper knowledge what is happening in your router, your network, if i give solution, you will copy paste, and never know, what exactly you just copied over.

if you want complete solution, contact consultants

             
            
           
          
            
            
              Thanks Janisk.  Most of my question are specifically regardint the implmentation I noted in my first post above (see the picture with different phases).   I was hoping to get the answers to the questions I posed 

             
            
           
          
            
            
              Hi to All

Being following on everyone different ideas and using sum samples then changing to suit our needs as i was always using simple queues by ip address but using the total limits as sample

interface=all parent=“Master Queue” direction=both priority=4 

queue=PCQ_Upload/PCQ_Download limit-at=256000/512000 

max-limit=512000/767000 total-queue=“PCQ Total” total-limit-at=384000 

total-max-limit=512000 total-burst-limit=767000 

total-burst-threshold=448000 total-burst-time=4s disabled=no

add name=“Josef” target-addresses=10.0.50.1/32 dst-address=0.0.0.0/0 

interface=all parent=“Master Queue” direction=both priority=4 

queue=PCQ_Upload/PCQ_Download limit-at=256000/512000 

max-limit=512000/767000 total-queue=“PCQ Total” total-limit-at=384000 

total-max-limit=512000 total-burst-limit=767000 

total-burst-threshold=448000 total-burst-time=4s disabled=no

add name=“Prof Makinde” target-addresses=10.0.50.2/32 dst-address=0.0.0.0/0 

interface=all parent=“Master Queue” direction=both priority=4 

queue=PCQ_Upload/PCQ_Download limit-at=256000/512000 

max-limit=512000/767000 total-queue=“PCQ Total” total-limit-at=384000 

total-max-limit=512000 total-burst-limit=767000 

total-burst-threshold=448000 total-burst-time=4s disabled=no


also using the PCQ as queue type as

add name=“PCQ_Upload” kind=pcq pcq-rate=0 pcq-limit=50 

pcq-classifier=src-address pcq-total-limit=2000

add name=“PCQ_Download” kind=pcq pcq-rate=0 pcq-limit=50 

pcq-classifier=dst-address pcq-total-limit=2000

add name=“PCQ Total” kind=pcq pcq-rate=0 pcq-limit=50 

