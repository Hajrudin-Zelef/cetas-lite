---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch-d1fca71b-1
title: "c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch--d1fca71b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch--d1fca71b.md
source_anchor: ""
source_lines: [1, 34]
sha256: 88863a56764232f463101468c1029ea13fa4411739bfa63983c654e56e693339
---

# c-en-us-td-docs-routers-asr920-configuration-guide-lanswitch-17-1-1-b-lanswitch--d1fca71b

|  Port-channel1                                           				  is up, line protocol is up                                            				                                         |  Indicates the                                           				  bundle interface is currently active and can transmit and receive or it has                                           				  been taken down by an administrator.                                            				                                         | 
                                 
                                    |  Hardware is                                            				                                         |  Hardware type                                           				  (Gigabit EtherChannel).                                            				                                         | 
                                 
                                    |  address is                                            				                                         |  Address being                                           				  used by the interface.                                            				                                         | 
                                 
                                    |  MTU                                            				                                         |  Maximum                                           				  transmission unit of the interface.                                            				                                         | 
                                 
                                    |  BW                                            				                                         |  Bandwidth of                                           				  the interface, in kilobits per second.                                            				                                         | 
                                 
                                    |  DLY                                            				                                         |  Delay of the                                           				  interface, in microseconds.                                            				                                         | 
                                 
                                    |  reliability                                            				                                         |  Reliability                                           				  of the interface as a fraction of 255 (255/255 is 100 percent reliability),                                           				  calculated as an exponential average over 5 minutes.                                            				                                         | 
                                 
                                    |  tx load                                           				  rxload                                            				                                         |  Transmit and                                           				  receive load on the interface as a fraction of 255 (255/255 is completely                                           				  saturated), calculated as an exponential average over 5 minutes. The                                           				  calculation uses the value from the                                            				  bandwidth                                            				  interface configuration command.                                            				                                         | 
                                 
                                    |  Encapsulation                                           				                                             				                                         |  Encapsulation                                           				  type assigned to the interface.                                            				                                         | 
                                 
                                    |  loopback                                            				                                         |  Indicates if                                           				  loopbacks are set.                                            				                                         | 
                                 
                                    |  keepalive                                            				                                         |  Indicates if                                           				  keepalives are set.                                            				                                         | 
                                 
                                    |  ARP type                                            				                                         |  Address                                           				  Resolution Protocol (ARP) type on the interface.                                            				                                         | 
                                 
                                    |  ARP Timeout                                            				                                         |  Number of                                           				  hours, minutes, and seconds an ARP cache entry stays in the cache.                                            				                                         | 
                                 
                                    |  No. of active                                           				  members in this channel                                            				                                         |  Number of                                           				  bundled ports (members) currently active and part of the port channel group.                                            				                                         | 
                                 
                                    |  Member                                           				  <no. > Gigabit                                           				  Ethernet: <no. /no. /no. >                                            				                                         |  Number of the                                           				  bundled port and associated Gigabit Ethernet port channel interface.                                            				                                         | 
                                 
                                    |  Last input                                            				                                         |  Number of                                           				  hours, minutes, and seconds since the last packet was successfully received by                                           				  an interface and processed locally on the Device. Useful for knowing when a                                           				  dead interface failed. This counter is updated only when packets are                                           				  process-switched, not when packets are fast-switched.                                            				                                         | 
                                 
                                    |  output                                            				                                         |  Number of                                           				  hours, minutes, and seconds since the last packet was successfully transmitted                                           				  by an interface. This counter is updated only when packets are                                           				  process-switched, not when packets are fast-switched.                                            				                                         | 
                                 
