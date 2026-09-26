---
id: collect-260926-mikrotik/mikrotik/queue-tree-qos-with-simple-queue-bandwidth-limiting-help-2
title: "queue-tree-qos-with-simple-queue-bandwidth-limiting-help"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/qos/queue-tree-qos-with-simple-queue-bandwidth-limiting-help.md
source_anchor: ""
source_lines: [196, 279]
sha256: 50627cdc6bb32419cf5ecf4322b4dd403bdadb9992124fd9727c7e9d63508931
---

# queue-tree-qos-with-simple-queue-bandwidth-limiting-help

pcq-classifier=src-address,dst-address pcq-total-limit=2000

add name=“default-small” kind=pfifo pfifo-limit=10


i was told before this can’t work but for me it is working fine as a client is busy downloading it would go in red to limit them at ± 450kpps = 55kBps witch is fine for us now i have testing mangle as Janisk’s sample marking the packets as new and old





/ip firewall mangle add chain=forward protocol=tcp action=mark-connection new-connection-mark=new_conn passthrough=yes comment=“mark all new connections” disabled=no

/ip firewall mangle add chain=forward protocol=tcp connection-mark=new_conn connection-bytes=0-2000000 action=mark-packet new-packet-mark=new_packet passthrough=no comment=“mark packets” disabled=no

/ip firewall mangle add chain=forward protocol=tcp connection-mark=new_conn action=mark-packet new-packet-mark=old_packets passthrough=no comment=“marking old packets” disabled=no


hen the queue trees

queue tree

add name=“Main_Upload” parent=ether2 packet-mark=“” limit-at=0 queue=default 

priority=8 max-limit=4000000 burst-limit=0 burst-threshold=0 burst-time=0s 

disabled=no

add name=“Up First 2Mbit” parent=Main_Upload packet-mark=new_packet 

limit-at=2000000 queue=PCQ_Upload priority=5 max-limit=4000000 

burst-limit=0 burst-threshold=0 burst-time=0s disabled=no

add name=“Up Rest Mbits” parent=Main_Upload packet-mark=old_packets 

limit-at=1000000 queue=PCQ_Upload priority=8 max-limit=4000000 

burst-limit=0 burst-threshold=0 burst-time=0s disabled=no

add name=“Main_Download” parent=global-out packet-mark=“” limit-at=0 

queue=default priority=8 max-limit=4000000 burst-limit=0 burst-threshold=0 

burst-time=0s disabled=no

add name=“Down First 2Mbit” parent=Main_Download packet-mark=new_packet 

limit-at=2000000 queue=PCQ_Download priority=5 max-limit=4000000 

burst-limit=0 burst-threshold=0 burst-time=0s disabled=no

add name=“Down Rest Mbits” parent=Main_Download packet-mark=old_packets 

limit-at=1000000 queue=PCQ_Download priority=8 max-limit=4000000 

burst-limit=0 burst-threshold=0 burst-time=0s disabled=no


Now it is working fine as I’m being monitoring the clients i see in the queue tree the bandwidth is great and then in the simple queues as i left them as is it still locks the clients bandwidth as i had before

Now this is the only way we learn to use QOS by trial and error

I know is not what you want but try with setting up all your simple queues with per user source address and then setup your mangle with the queue trees as you need them because this is the only way to learn and i know every-ones setup is different so if swapping ideas it might help and please correct me if I’m wrong would always like to improve my self and just waiting for a mikrotik course to attend as soon as it near us

Sidney 

             
            
           
          
            
            
              let’s assume the following :

–packets----> sat modem ----> MT -----> user

—stage1----> sat modem ----> MT --stage2—> user

if we made a perfect QoS rules , very perfect lets say , this QoS will effect stage 2 of course .. but would it effect stage 1 ??

lets say , a user is downloading a big file , and another user begins a browsing , the MT will give the new packets of the browser the highest priority , of course that will be in stage 2 , but what about the stage 1 ??

does the browser user packets will get the highest priority also in stage 1 ?
