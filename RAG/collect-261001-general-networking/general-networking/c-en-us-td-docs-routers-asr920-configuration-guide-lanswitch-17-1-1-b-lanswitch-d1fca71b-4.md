---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch-d1fca71b-4
title: "c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch--d1fca71b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory"]
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch--d1fca71b.md
source_anchor: ""
source_lines: [87, 99]
sha256: 826bc2aee9e01a404f52fb0d8c0f745dbb8dac8459bda83a3fdf8640c5d74ed7
---

# c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch--d1fca71b

                                    |  late                                           				  collision                                            				                                         |  Number of                                           				  late collisions. Late collision happens when a collision occurs after                                           				  transmitting the preamble. The most common cause of late collisions is that                                           				  your Ethernet cable segments are too long for the speed at which you are                                           				  transmitting.                                            				                                         | 
                                 
                                    |  deferred                                            				                                         |  Indicates                                           				  that the chip had to defer while ready to transmit a frame because the carrier                                           				  was asserted.                                            				                                         | 
                                 
                                    |  lost                                           				  carrier                                            				                                         |  Number of                                           				  times the carrier was lost during transmission.                                            				                                         | 
                                 
                                    |  no carrier                                            				                                         |  Number of                                           				  times the carrier was not present during the transmission.                                            				                                         | 
                                 
                                    |  PAUSE                                           				  output                                            				                                         |  Not                                           				  supported.                                            				                                         | 
                                 
                                    |  output                                           				  buffer failures                                            				                                         |  Number of                                           				  times that a packet was not output from the output hold queue because of a                                           				  shortage of shared memory.                                            				                                         | 
                                 
                                    |  output                                           				  buffers swapped out                                            				                                         |  Number of                                           				  packets stored in main memory when the output queue is full; swapping buffers                                           				  to main memory prevents packets from being dropped when output is congested.                                           				  The number is high when traffic is bursty.                                            				                                         |
