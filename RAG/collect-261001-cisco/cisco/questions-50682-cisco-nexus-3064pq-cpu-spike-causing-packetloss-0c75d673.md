---
id: collect-261001-cisco/cisco/questions-50682-cisco-nexus-3064pq-cpu-spike-causing-packetloss-0c75d673
title: "questions-50682-cisco-nexus-3064pq-cpu-spike-causing-packetloss-0c75d673"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-50682-cisco-nexus-3064pq-cpu-spike-causing-packetloss-0c75d673.md
source_anchor: ""
source_lines: [1, 64]
sha256: 0c8cbed270ec91cad459e4dfb5305ae3d9820a9bad5829edacddecf79e87ca42
---

# questions-50682-cisco-nexus-3064pq-cpu-spike-causing-packetloss-0c75d673

Nexus 3064 CPU spike and packetloss I have L3 Cisco 3064PQ switch which is running latest "nxos.7.0.3.I4.7.bin" basically I am using for routing so my ISP 2x10Gbps link terminated (with LACP) on single switch (without STP,vPC) it and currently live traffic is 10Gbps on switch but i have notice periodic packetloss and not sure where its coming from, after digging i found when my CPU spike up to 70% i have seeing packet loss.
If you notice in below picture that 60% spike and at same time i have noticed 1 ping packet loss. How do i debug this issue and find out what is that pike for and why its happening periodically
    111113641111 11 111 11 11 1  11   1   1 11  1 11     11 12
    011245362136900971084282051974067736846931972921847680082588
100
 90
 80
 70
 60       #
 50       ##
 40      ###
 30      ###                                                 #
 20      ###   #    #                     #                  #
 10 ##################################### ########### ##########
    0....5....1....1....2....2....3....3....4....4....5....5....
              0    5    0    5    0    5    0    5    0    5
               CPU% per second (last 60 seconds)
                      # = average CPU%
    655647776676716562796267574676486635666266736656375765565666
    389694233637271934049753800212280577262865100886229818827948
100
 90                    *           *
 80                    *           *                   *
 70    * *** ****     *** ** *  *  * *   *  ***  * * * *     * *
 60 **** ******** *** *** **** *** *** **** *** **** ***********
 50 ************* *** *** **** *** *** **** *** **** ***********
 40 ************* *** *** ***************** *** **** ***********
 30 ************* *** ******************************************
 20 ***#**###**##******#***#*#*#***#******#**#**#****#**********
 10 ############################################################
    0....5....1....1....2....2....3....3....4....4....5....5....
              0    5    0    5    0    5    0    5    0    5
               CPU% per minute (last 60 minutes)
              * = maximum CPU%   # = average CPU%
UPDATE
In following sample you can see 50% spike t2usd eating more cpus
- Notes: I am seeing CPU spike every 1 minute and 30second, and its very accurate i did stop-watch test and spiking coming every 1.30second and same time ping drop. # show processes cpu sort | ex 0.00 PID    Runtime(ms)  Invoked   uSecs  1Sec    Process
-----  -----------  --------  -----  ------  -----------
12624    641746706  456010193   1407   7.00%  t2usd
   27   1262455737  1006811706   1253   4.00%  ksmd
11145    288596961  111352447   2591   2.00%  pfmclnt
11367          113       253    448   1.00%  arp
11402          200       349    575   1.00%  netstack
CPU util  :   51.33% user,    9.62% kernel,   39.03% idle
Please note that only processes from the requested vdc are shown above
second spike
 # show processes cpu sort | ex 0.00
   PID    Runtime(ms)  Invoked   uSecs  1Sec    Process
-----  -----------  --------  -----  ------  -----------
12624    641774351  456031502   1407  26.50%  t2usd
   27   1262516503  1006859899   1253   8.00%  ksmd
11367          113       253    448   2.00%  arp
11371          149       106   1406   2.00%  pktmgr
12764      5010346  18794402    266   2.00%  ipfib
11356           79        43   1838   1.00%  adjmgr
11402          200       349    575   1.00%  netstack
12261          116        65   1799   1.00%  rpm
12271      3321325  27299929    121   1.00%  ipfib
12334     23532716  29888867    787   1.00%  l2fm
CPU util  :   57.83% user,    5.40% kernel,   36.75% idle
Please note that only processes from the requested vdc are shown above
UPDATE 2
Does this related to CoPP, Default its enabled on Nexus switches https://supportforums.cisco.com/t5/data-center-documents/packet-loss-when-pinging-from-to-a-nexus-7000/ta-p/3110226
It shouldn't impact forwarding traffic right?
