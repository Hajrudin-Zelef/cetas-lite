---
id: collect-261001-general-networking/general-networking/bgp-oldest-path-it-s-not-what-you-think-practical-networking-net-2
title: "bgp-oldest-path-it-s-not-what-you-think-practical-networking-net"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/bgp-oldest-path-it-s-not-what-you-think-practical-networking-net.md
source_anchor: ""
source_lines: [235, 407]
sha256: 34be2835197d7acacc99285518dd5024b5de421b040693b77957ab75667d83bb
---

# bgp-oldest-path-it-s-not-what-you-think-practical-networking-net

Because of the order in which we brought the neighbor adjacencies back up, we know that R3 had the “next oldest” path. And indeed, this is the path R1 now selects:

R1# **show ip bgp**
...
     Network          Next Hop            Metric LocPrf Weight Path
 *   5.5.5.0/24       9.44.11.4                              0 44 55 i
 ***>                   9.33.11.3**                              0 33 55 i

R1# **show ip bgp 5.5.5.0/24**
BGP routing table entry for 5.5.5.0/24, version 8
Paths: (2 available, best #2, table default)
  Advertised to update-groups:
     3
  Refresh Epoch 2
  44 55
    9.44.11.4 from 9.44.11.4 (4.4.4.4)
      Origin IGP, localpref 100, valid, external
      rx pathid: 0, tx pathid: 0
  Refresh Epoch 2
  33 55
    **9.33.11.3 from 9.33.11.3 (3.3.3.3)**
      Origin IGP, localpref 100, valid, external, **best**
      rx pathid: 0, tx pathid: 0x0

We can then re-enable the adjacency to R2:

R1(config)# **router bgp 11**
R1(config-router)# **no neighbor 9.22.11.2 shutdown**
*May  9 19:27:22.811: %BGP-5-ADJCHANGE: neighbor 9.22.11.2 Up

And as expected, we see R1 learns of the new path through R2 but does not select it as best, as the current oldest (and best) path is still through R3:

R1# **show ip bgp**
...
     Network          Next Hop            Metric LocPrf Weight Path
 *   5.5.5.0/24       **9.22.11.2**                              0 22 55 i
 *                    9.44.11.4                              0 44 55 i
 ***>                   9.33.11.3**                              0 33 55 i

R1# **show ip bgp 5.5.5.0/24**
BGP routing table entry for 5.5.5.0/24, version 8
Paths: (3 available, best #3, table default)
  Advertised to update-groups:
     3
  Refresh Epoch 2
  22 55
    **9.22.11.2 from 9.22.11.2 (2.2.2.2)**
      Origin IGP, localpref 100, valid, external
      rx pathid: 0, tx pathid: 0
  Refresh Epoch 2
  44 55
    9.44.11.4 from 9.44.11.4 (4.4.4.4)
      Origin IGP, localpref 100, valid, external
      rx pathid: 0, tx pathid: 0
  Refresh Epoch 2
  33 55
    **9.33.11.3 from 9.33.11.3 (3.3.3.3)**
      Origin IGP, localpref 100, valid, external, **best**
      rx pathid: 0, tx pathid: 0x0

Notice also the output of the **`show ip bgp`** and **`show ip bgp 5.5.5.0/24`** commands are displaying the paths in the order they have learned them. R3 being the oldest (at the bottom), followed by R4, followed by R2 (at the top).

So far, everything is working as we expected it to. But maybe not exactly for the reason you thought it might. Continue reading to find out exactly why R3 was chosen and remained the best path after the path through R2 was lost and then reacquired.

### BGP Oldest Path – Trial 2

This is where it gets interesting.

At the moment, we have three paths to 5.5.5.0/24, learned in this order: R3, R4, R2. We will go ahead and shut down the R3 adjacency (the current oldest path) to see what R1 picks next for the best path.

R1(config)# **router bgp 11**
R1(config-router)# **neighbor 9.33.11.3 shutdown**
*May  9 19:35:16.041: %BGP-5-ADJCHANGE: neighbor 9.33.11.3 Down Admin. shutdown

If Step 10 was purely based upon a pathâs absolute age, then R4 should be picked as the best path, since it is the next oldest path. But youâll see that is not what happens:

R1# **show ip bgp**
...
     Network          Next Hop            Metric LocPrf Weight Path
 ***>**  5.5.5.0/24       **9.22.11.2**                              0 22 55 i
 *                    9.44.11.4                              0 44 55 i

R1# **show ip bgp 5.5.5.0/24**
BGP routing table entry for 5.5.5.0/24, version 9
Paths: (2 available, best #1, table default)
  Advertised to update-groups:
     3	
  Refresh Epoch 2
  22 55
    **9.22.11.2 from 9.22.11.2 (2.2.2.2)**
      Origin IGP, localpref 100, valid, external, **best**
      rx pathid: 0, tx pathid: 0x0
  Refresh Epoch 2
  44 55
    9.44.11.4 from 9.44.11.4 (4.4.4.4)
      Origin IGP, localpref 100, valid, external
      rx pathid: 0, tx pathid: 0

R2 was selected as the best path, even though we know that the absolute oldest path was the path through R4 (and the order of the output in the commands above also confirm this).

We will discuss the reason for this behavior, but before we do lets re-enable the R3 adjacency and ensure R1 knows of all three paths to 5.5.5.0/24:

R1(config)# **router bgp 11**
R1(config-router)# **no neighbor 9.33.11.3 shutdown**
*May  9 19:37:46.989: %BGP-5-ADJCHANGE: neighbor 9.33.11.3 Up

R1# **show ip bgp**
...
     Network          Next Hop            Metric LocPrf Weight Path
 *   5.5.5.0/24       9.33.11.3                              0 33 55 i
 ***>                   9.22.11.2**                              0 22 55 i
 *                    9.44.11.4                              0 44 55 i

R1# **show ip bgp 5.5.5.0/24**
BGP routing table entry for 5.5.5.0/24, version 9
Paths: (3 available, best #2, table default)
  Advertised to update-groups:
     3
  Refresh Epoch 2
  33 55
    9.33.11.3 from 9.33.11.3 (3.3.3.3)
      Origin IGP, localpref 100, valid, external
      rx pathid: 0, tx pathid: 0
  Refresh Epoch 2
  22 55
    **9.22.11.2 from 9.22.11.2 (2.2.2.2)**
      Origin IGP, localpref 100, valid, external, **best**
      rx pathid: 0, tx pathid: 0x0
  Refresh Epoch 2
  44 55
    9.44.11.4 from 9.44.11.4 (4.4.4.4)
      Origin IGP, localpref 100, valid, external
      rx pathid: 0, tx pathid: 0


### What Happened

Initially, R1 learned of three paths to the 5.5.5.0/24 prefix in the order of R2, then R3, then R4. When we disabled the R2 adjacency, R1âs next best path was through R3. It is generally accepted that this occurred because R3 was the next oldest path, but what happened next proved that this is not entirely true.

After R2 came back up, R1 had three paths to the 5.5.5.0/24 prefix in the order of R3, then R4, then R2. When we disabled R3, R1 did not choose R4 as the next best path. R1 instead chose R2 as the next best path. But why?

This happened because when the path through R3 was lost, BGP looked at its topology table and used the Path Selection process to pick the next best path. Steps 1-9 were all tied, which brought the BGP speaker to Step 10. At the time R1 lost the path through R3, R1Â *already had both paths through R2 and R4* in the topology table. Since they both already existed *when Step 10 was being processed*, from R1’s perspective neither one was older than the other. As such, Step 10 resulted in a tie.

BGP then used the *next step* in the Path Selection process to break the tie: Step 11 â Preferring the path with the lowest Router-ID. Between R2 and R4, R2 had a better Router-ID (`2.2.2.2`). In the output of **`show ip bgp 5.5.5.0/24`** we can see the Router-ID of each neighbor listed next to the next-hop IP address:

R1# **show ip bgp 5.5.5.0/24**
BGP routing table entry for 5.5.5.0/24, version 9
Paths: (3 available, best #2, table default)
  Advertised to update-groups:
     3
  Refresh Epoch 2
  33 55
    9.33.11.3 from 9.33.11.3 **(3.3.3.3)**
      Origin IGP, localpref 100, valid, external
      rx pathid: 0, tx pathid: 0
  Refresh Epoch 2
  22 55
    9.22.11.2 from 9.22.11.2 **(2.2.2.2)**
      Origin IGP, localpref 100, valid, external, **best**
      rx pathid: 0, tx pathid: 0x0
  Refresh Epoch 2
  44 55
    9.44.11.4 from 9.44.11.4 **(4.4.4.4)**
      Origin IGP, localpref 100, valid, external
      rx pathid: 0, tx pathid: 0

So what happened in the first example? R1 initially knew of the paths in this order R2, then R3, then R4. When R1 lost the path through R2, the next best path selected was R3. Not because R3 was the next oldest route, but because between the remaining paths R3 had a better Router-ID than R4. Step 11 preferred the path with the lower Router-ID.


### Synopsis

Ultimately, Step 10 exists to prevent route table instability.

