---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch-d1fca71b-2
title: "c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch--d1fca71b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch--d1fca71b.md
source_anchor: ""
source_lines: [35, 60]
sha256: 53a26d9b67b24ac812f56f3a46a811b15f067ba0caf05c94d57b4a38ec137720
---

# c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch--d1fca71b

                                    |  output hang                                           				                                             				                                         |  Number of                                           				  hours, minutes, and seconds since the interface was last reset because of a                                           				  transmission that took too long. When the number of hours in any of the “last”                                           				  fields exceeds 24 hours, the number of days and hours is printed. If that field                                           				  overflows, asterisks are printed.                                            				                                         | 
                                 
                                    |  last                                           				  clearing                                            				                                         |  Time at                                           				  which the counters that measure cumulative statistics (such as number of bytes                                           				  transmitted and received) shown in this report were last reset to zero.                                           				  Variables that might affect routing (for example, load and reliability) are not                                           				  cleared when the counters are cleared.                                            				                                          ***                                           				  indicates that the elapsed time is too long to be displayed.                                            				                                          0:00:00                                           				  indicates that the counters were cleared more than 231 ms and less than 232 ms                                           				  ago.                                            				                                         | 
                                 
                                    |  Input queue                                           				                                             				                                         |  Number of                                           				  packets in the input queue and the maximum size of the queue.                                            				                                         | 
                                 
                                    |  Queueing                                           				  strategy                                            				                                         |  First-in,                                           				  first-out queueing strategy (other queueing strategies you might see are                                           				  priority-list, custom-list, and weighted fair).                                            				                                         | 
                                 
                                    |  Output                                           				  queue                                            				                                         |  Number of                                           				  packets in the output queue and the maximum size of the queue.                                            				                                         | 
                                 
                                    |  5 minute                                           				  input rate 5 minute output rate                                            				                                         |  Average                                           				  number of bits and packets received or transmitted per second in the last 5                                           				  minutes.                                            				                                         | 
                                 
                                    |  packets                                           				  input                                            				                                         |  Total                                           				  number of error-free packets received by the system.                                            				                                         | 
                                 
                                    |  bytes                                           				  (input)                                            				                                         |  Total                                           				  number of bytes, including data and MAC encapsulation, in the error-free                                           				  packets received by the system.                                            				                                         | 
                                 
                                    |  no buffer                                            				                                         |  Number of                                           				  received packets discarded because there was no buffer space in the main                                           				  system. Broadcast storms on Ethernet lines and bursts of noise on serial lines                                           				  are often responsible for no input buffer events.                                            				                                         | 
                                 
                                    |  broadcasts                                            				                                         |  Total                                           				  number of broadcast or multicast packets received by the interface.                                            				                                         | 
                                 
                                    |  runts                                            				                                         |  Number of                                           				  packets that are discarded because they are smaller than the minimum packet                                           				  size for the medium.                                            				                                         | 
                                 
                                    |  giants                                            				                                         |  Number of                                           				  packets that are discarded because they exceed the maximum packet size for the                                           				  medium.                                            				                                         | 
                                 
                                    |  input                                           				  errors                                            				                                         |  Total number of no buffer, runts, giants, cyclic redundancy checks (CRCs), frame, overrun, ignored, and terminated counts.                                           Other input-related errors can also increment the count, so that this sum might not balance with the other counts.                                          | 
                                 
