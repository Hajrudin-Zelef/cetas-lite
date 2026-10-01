---
id: collect-261001-general-networking/general-networking/more-links-8-1
title: "Execute a CLI script based on CPU and memory thresholds"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/more-links-8.md
source_anchor: ""
source_lines: [1, 209]
sha256: 14fbc9f348c4b2d779ceb6fa0b947c3a8f24e0aca2411b11ea445c35b02a32b1
---

# Execute a CLI script based on CPU and memory thresholds

Automation stitches can be created to run a CLI script and send an email message when CPU or memory usage exceeds specified thresholds.

In this example, two automation stitches are created that run a CLI script to collect debug information, and then email the results of the script to a specified email address when the CPU usage threshold is exceeded, or memory usage causes the FortiGate to enter conserve mode.

|  | The maximum size of the CLI script action output is 16K characters. | 

###### To define CPU and memory usage thresholds:

```
config system global
    set cpu-use-threshold <percent>
    set memory-use-threshold-extreme <percent>
    set memory-use-threshold-green <percent>
    set memory-use-threshold-red <percent>
end
```
                                            Where:

| cpu-use-threshold | Threshold at which CPU usage is reported, in percent of total possible CPU utilization (default = 90). | 
| memory-use-threshold-extreme | Threshold at which memory usage is considered extreme, and new sessions are dropped, in percent of total RAM (default = 95). | 
| memory-use-threshold-green | Threshold at which memory usage forces the FortiGate to exit conserve mode, in percent of total RAM (default = 82). | 
| memory-use-threshold-red | Threshold at which memory usage forces the FortiGate to enter conserve mode, in percent of total RAM (default = 88). | 

## Configuring the automation stitches

### High CPU usage stitch

###### To create an automation stitch for high CPU usage in the GUI:

1. Go to *Security Fabric > Automation* and click*Create New* .
2. Enter the stitch name (*auto_high_cpu* ).
3. Configure the trigger:
  1. Click *Add Trigger* .
  2. Click *Create* and select*High CPU* .
  3. Enter the name, *auto_high_cpu* .
  4. Click *OK* .
  5. Select the trigger in the list and click *Apply* .
4. Click 
5. Configure the CLI Script action:
  1. Click *Add Action* .
  2. Click *Create* and select*CLI Script* .
  3. Enter the following:Name high_cpu_debug Script diagnose debug cli 8 diagnose debug console timestamp enable diagnose debug enable diagnose debug crashlog read get system performance status get system session status diagnose sys session full-stat diagnose firewall iprope state diagnose sys flash list diagnose hardware sysinfo memory diagnose hardware sysinfo slab diagnose hardware sysinfo shm diagnose hardware deviceinfo disk get system arp diagnose ip arp list diagnose ip address list get router info routing-table all get router info kernel diagnose ip rtcache list diagnose sys top-summary diagnose sys top 9 99 Administrator profile Select a profile
  4. Click *OK* .
  5. Select the action in the list and click *Apply* .
6. Click 
7. Configure the Email notification action:
  1. Click *Add Action* .
  2. Click *Create* and select*Email* .
  3. Enter the following:Name auto_high_cpu_email To Enter an email address Subject CSF stitch alert: high_cpu Body %%results%%
  4. Click *OK* .
  5. Select the action in the list and click *Apply* .
8. Click 
9. Click *OK* .

###### To create an automation stitch for high CPU usage in the CLI:

1. Create the automation trigger:```
config system automation-trigger
    edit "auto_high_cpu"
        set event-type high-cpu
    next
end
```
2. Create the automation actions:```
config system automation-action
    edit "high_cpu_debug"
        set action-type cli-script
        set script "diagnose debug cli 8
diagnose debug console timestamp enable
diagnose debug enable
diagnose debug crashlog read
get system performance status
get system session status
diagnose sys session full-stat
diagnose firewall iprope state
diagnose sys flash list
diagnose hardware sysinfo memory
diagnose hardware sysinfo slab
diagnose hardware sysinfo shm
diagnose hardware deviceinfo disk
get system arp
diagnose ip arp list
diagnose ip address list
get router info routing-table all
get router info kernel
diagnose ip rtcache list
diagnose sys top-summary
diagnose sys top 9 99"
        set accprofile "super_admin"
    next
    edit "auto_high_cpu_email"
        set action-type email
        set email-to "person@fortinet.com"
        set email-subject "CSF stitch alert: high_cpu"
        set message "%%results%%"
    next
end
```
3. Create the automation stitch:```
config system automation-stitch
    edit "auto_high_cpu"
        set trigger "auto_high_cpu"
        config actions
            edit 1
                set action "high_cpu_debug"
                set required enable
            next
            edit 2
                set action "auto_high_cpu_email"
                set required enable
            next
        end
    next
end
```

### High memory usage stitch

###### To create an automation stitch for high memory usage in the GUI:

1. Go to *Security Fabric > Automation* and click*Create New* .
2. Enter the stitch name (*auto_high_memory* ).
3. Configure the trigger:
  1. Click *Add Trigger* .
  2. Click *Create* and select*Conserve Mode* .
  3. Enter the name, *auto_high_memory* .
  4. Click *OK* .
  5. Select the trigger in the list and click *Apply* .
4. Click 
5. Configure the CLI Script action:
  1. Click *Add Action* .
  2. Click *Create* and select*CLI Script* .
  3. Enter the following:Name high_memory_debug Script diagnose debug cli 8 diagnose debug console timestamp enable diagnose debug enable diagnose debug crashlog read get system performance status get system session status diagnose sys session full-stat diagnose firewall iprope state diagnose sys flash list diagnose hardware sysinfo memory diagnose hardware sysinfo slab diagnose hardware sysinfo shm diagnose hardware deviceinfo disk get system arp diagnose ip arp list diagnose ip address list get router info routing-table all get router info kernel diagnose ip rtcache list diagnose sys top-summary diagnose sys top 9 99 Administrator profile Select a profile
  4. Click *OK* .
  5. Select the action in the list and click *Apply* .
6. Click 
7. Configure the Email notification action:
  1. Click *Add Action* .
  2. Click *Create* and select*Email* .
  3. Enter the following:Name auto_high_memory_email To Enter an email address Subject CSF stitch alert: high_memory Body %%results%%
  4. Click *OK* .
  5. Select the action in the list and click *Apply* .
8. Click 
9. Click *OK* .

###### To create an automation stitch for high memory usage in the CLI:

1. Create the automation trigger:```
config system automation-trigger
    edit "auto_high_memory"
        set event-type low-memory
    next
end
```
2. Create the automation actions:```
config system automation-action
    edit "high_memory_debug"
        set action-type cli-script
        set script "diagnose debug cli 8
diagnose debug console timestamp enable
diagnose debug enable
diagnose debug crashlog read
get system performance status
get system session status
diagnose sys session full-stat
diagnose firewall iprope state
diagnose sys flash list
diagnose hardware sysinfo memory
diagnose hardware sysinfo slab
diagnose hardware sysinfo shm
diagnose hardware deviceinfo disk
get system arp
diagnose ip arp list
diagnose ip address list
get router info routing-table all
get router info kernel
diagnose ip rtcache list
diagnose sys top-summary
diagnose sys top 9 99"
        set accprofile "super_admin"	
    next
    edit "auto_high_memory_email"
        set action-type email
        set email-to "person@fortinet.com"
        set email-subject "CSF stitch alert: high_memory"
        set message "%%results%%"
    next
end
```
3. Create the automation stitch:```
config system automation-stitch
    edit "auto_high_memory"
        set trigger "auto_high_memory"
        config actions
            edit 1
                set action "high_memory_debug"
                set required enable
            next
            edit 2
                set action "auto_high_memory_email"
                set required enable
            next
        end
    next
end
```

