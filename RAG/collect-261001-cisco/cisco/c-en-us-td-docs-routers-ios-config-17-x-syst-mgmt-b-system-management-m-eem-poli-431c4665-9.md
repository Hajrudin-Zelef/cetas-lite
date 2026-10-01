---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-syst-mgmt-b-system-management-m-eem-poli-431c4665-9
title: "c-en-us-td-docs-routers-ios-config-17-x-syst-mgmt-b-system-management-m-eem-poli-431c4665"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-syst-mgmt-b-system-management-m-eem-poli-431c4665.md
source_anchor: ""
source_lines: [201, 230]
sha256: 3d564140f54a6b0514c66554eb5edea7e55af37513ee41d01828639555834eb0
---

# c-en-us-td-docs-routers-ios-config-17-x-syst-mgmt-b-system-management-m-eem-poli-431c4665

                                    |                                             				                                                 					 _timer_remain                                               				                                               				                                         |  The time available before the timer expires.                                            				                                                                                                                                  | Note |  This environment variable is not available for the CRON timer.                                                       				                                                     |  | 
                                 
                                    |                                             				                                                 					 _timer_time                                               				                                               				                                         |  The time at which the last event was triggered.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _timer_type                                               				                                               				                                         |  The type of timer.                                           				                                         | 
                                 
                                    |  Watchdog System Monitor (IOSWDSysMon) Event Detector                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_node                                               				                                               				                                         |  The slot number for the Route Processor (RP) reporting node.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_num_subs                                               				                                               				                                         |  The number of subevents present.                                           				                                         | 
                                 
                                    |  All Watchdog System Monitor (IOSWDSysMon) Subevents                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_present                                               				                                               				                                                 					 _ioswd_sub2_present                                               				                                               				                                         |  A value to indicate whether subevent 1 or subevent 2 is present. A value of 1 means that the subevent is present; a value                                           of 0 means that the subevent is not present.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_type                                               				                                               				                                                 					 _ioswd_sub2_type                                               				                                               				                                         |  The event type, either cpu_proc or mem_proc.                                           				                                         | 
                                 
                                    |  Watchdog System Monitor (IOSWDSysMon) cpu_proc Subevents                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_path                                               				                                               				                                                 					 _ioswd_sub2_path                                               				                                               				                                         |  A process name of subevents.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_period                                               				                                               				                                                 					 _ioswd_sub2_period                                               				                                               				                                         |  The time period, in seconds and optional milliseconds, used for measurement in subevents.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_pid                                               				                                               				                                                 					 _ioswd_sub2_pid                                               				                                               				                                         |  The process identifier of subevents.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_taskname                                               				                                               				                                                 					 _ioswd_sub2_taskname                                               				                                               				                                         |  The task name of subevents.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_value                                               				                                               				                                                 					 _ioswd_sub2_value                                               				                                               				                                         |  The CPU utilization of subevents measured as a percentage.                                           				                                         | 
                                 
