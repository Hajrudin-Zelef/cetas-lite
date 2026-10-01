---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-syst-mgmt-b-system-management-m-eem-poli-431c4665-10
title: "c-en-us-td-docs-routers-ios-config-17-x-syst-mgmt-b-system-management-m-eem-poli-431c4665"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-syst-mgmt-b-system-management-m-eem-poli-431c4665.md
source_anchor: ""
source_lines: [231, 256]
sha256: 3432d09c23213d4da868ecdb0c565fafd0068b2200c7941d036e99224dac6e4f
---

# c-en-us-td-docs-routers-ios-config-17-x-syst-mgmt-b-system-management-m-eem-poli-431c4665

                                    |  Watchdog System Monitor (IOSWDSysMon) mem_proc Subevents                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_diff                                               				                                               				                                                 					 _ioswd_sub2_diff                                               				                                               				                                         |  A percentage value of the difference that triggered the event.                                           				                                                                                                                                  | Note |  This variable is set only when the                                                        				  _ioswd_sub1_is_percent  or                                                        				  _ioswd_sub2_is_percent  variable contains a value of 1.                                                       				                                                     |  | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_is_percent                                               				                                               				                                                 					 _ioswd_sub2_is_percent                                               				                                               				                                         |  A number that identifies whether the value is a percentage. A value of 0 means that the value is not a percentage; a value                                           of 1 means that the value is a percentage.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_path                                               				                                               				                                                 					 _ioswd_sub2_path                                               				                                               				                                         |  The process name of subevents.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_pid                                               				                                               				                                                 					 _ioswd_sub2_pid                                               				                                               				                                         |  The process identifier of subevents.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_taskname                                               				                                               				                                                 					 _ioswd_sub2_taskname                                               				                                               				                                         |  The task name of subevents.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _ioswd_sub1_value                                               				                                               				                                                 					 _ioswd_sub2_value                                               				                                               				                                         |  The CPU utilization of subevents measured as a percentage.                                           				                                         | 
                                 
                                    |  Watchdog System Monitor (WDSysMon) Event Detector                                           				                                         | 
                                 
                                    |                                             				                                                 					 _wd_sub1_present                                               				                                               				                                                 					 _wd_sub2_present                                               				                                               				                                         |  A value to indicate whether subevent 1 or subevent 2 is present. A value of 1 means that the subevent is present; a value                                           of 0 means that the subevent is not present.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _wd_num_subs                                               				                                               				                                         |  The number of subevents present.                                           				                                         | 
                                 
                                    |                                             				                                                 					 _wd_sub1_type                                               				                                               				                                                 					 _wd_sub2_type                                               				                                               				                                         |  The event type: cpu_proc, cpu_tot, deadlock, dispatch_mgr, mem_proc, mem_tot_avail, or mem_tot_used.                                           				                                         | 
                                 
                                    |  Watchdog System Monitor (WDSysMon) cpu_proc Subevents                                           				                                         | 
                                 
                                    |                                             				                                                 					 _wd_sub1_node                                               				                                               				                                                 					 _wd_sub2_node                                               				                                               				                                         |  The slot number for the subevent RP reporting node.                                           				                                         | 
                                 
