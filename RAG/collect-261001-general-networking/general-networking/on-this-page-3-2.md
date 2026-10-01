---
id: collect-261001-general-networking/general-networking/on-this-page-3-2
title: "Interface-based traffic shaping profile"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/on-this-page-3.md
source_anchor: ""
source_lines: [120, 275]
sha256: f07885859260e97d9c08012a562b8debeae0b3d5eecb5620b064a6417c291b76
---

# Interface-based traffic shaping profile

1. 
                                                    Create a recurring schedule: ```
config firewall schedule recurring
    edit "Day_Hours"
        set start 08:00
        set end 18:00
        set day sunday monday tuesday wednesday thursday friday saturday
    next
end
```
2. 
                                                    Create the traffic class IDs: ```
config firewall traffic-class
    edit 3
        set class-name "Web Access"
    next
    edit 4
        set class-name "File Access"
    next
end
```
3. 
                                                    Create the web and file accessing traffic shaping policies: ```
config firewall shaping-policy
    edit 2
        set name "web_access_day_hours"
        set comment "Limit web accessing traffic to 8Mb/s in day time"
        set service "HTTP" "HTTPS"
        set schedule "Day_Hours"
        set dstintf "wan1"
        set class-id 3
        set srcaddr "all"
        set dstaddr "all"
    next
    edit 3
        set name "file_access_day_hours"
        set comment "Limit file accessing traffic to 2Mb/s during the day"
        set service "AFS3" "FTP" "FTP_GET" "FTP_PUT" "NFS" "SAMBA" "SMB" "TFTP"
        set schedule "Day_Hours"
        set dstintf "wan1"
        set class-id 4
        set srcaddr "all"
        set dstaddr "all"
    next
end
```

### Allocating bandwidth to the shaping classes

A traffic shaping profile defines the guaranteed and maximum bandwidths each class receives. In this example, file access can use up to 2 Mb/s and web access can use 8 Mb/s from 8:00 AM to 6:00 PM.

###### To create a traffic shaping profile using the GUI:

1. 
                                                    Go to *Policy & Objects > Traffic Shaping* , select the*Traffic Shaping Profile* tab, and click*Create New* .
2. 
                                                    Enter a name for the profile, such as *Day_Hours_Profile* .
3. 
                                                    Configure a default traffic shaping class: This class has a high priority, meaning that when the other classes have reached their guaranteed bandwidths, this default class will use the rest of the available bandwidth. 
  1. 
                                                            In the *Traffic Shaping Classes* table click*Create New* .
  2. 
                                                            Click the *Traffic shaping class ID* drop down then click*Create* .
  3. 
                                                            Enter a name for the class, such as *Default Access* .
  4. 
                                                            Click *OK* .
  5. 
                                                            Select the class ID you just created for *Traffic shaping class ID* .
  6. 
                                                            Configure the following settings, then click *OK* :Guaranteed bandwidth 30 Maximum bandwidth 100 Priority High
4. 
                                                            
5. 
                                                    Configure a web accessing traffic shaping class: When other types of traffic are competing for bandwidth, this class is guaranteed to 6 Mb/s, or 60% of the bandwidth. 
  1. 
                                                            In the *Traffic Shaping Classes* table click*Create New* .
  2. 
                                                            Configure the following settings, then click *OK* :Traffic shaping class ID Web Access Guaranteed bandwidth 60 Maximum bandwidth 80 Priority Medium
6. 
                                                            
7. 
                                                    Configure a file accessing traffic shaping class: When other types of traffic are competing for bandwidth, this group is guaranteed to 1 Mb/s, or 10% of the bandwidth. 
  1. 
                                                            In the *Traffic Shaping Classes* table click*Create New* .
  2. 
                                                            Configure the following settings, then click *OK* :Traffic shaping class ID File Access Guaranteed bandwidth 10 Maximum bandwidth 20 Priority Medium
8. 
                                                            
9. 
                                                    Click *OK* .

###### To create a traffic shaping profile using the CLI:

```
config firewall shaping-profile
    edit "Day_Hours_Profile"
        set default-class-id 2
        config shaping-entries
            edit 1
                set class-id 2
                set guaranteed-bandwidth-percentage 30
                set maximum-bandwidth-percentage 100
            next
            edit 2
                set class-id 3
                set priority medium
                set guaranteed-bandwidth-percentage 60
                set maximum-bandwidth-percentage 80
            next
            edit 3
                set class-id 4
                set priority medium
                set guaranteed-bandwidth-percentage 10
                set maximum-bandwidth-percentage 20
            next
        end
    next
end
```
                                            ### Defining the available bandwidth on an interface

In this example, the link speed of the wan1 interface is 10 Mb/s.

###### To set the bandwidth of the wan1 interface in the GUI:

1. 
                                                    Go to *Network > Interfaces* .
2. 
                                                    Edit the wan1 interface.
3. 
                                                    Under Traffic Shaping, enable *Outbound shaping profile* and select the profile that you just created,*Day_Hours_Profile* .
4. 
                                                    Enable *Outbound Bandwidth* and set it to*10000* Kbps.
5. 
                                                    Click *OK* .

###### To set the bandwidth of the wan1 interface in the CLI:

```
config system interface
    edit "wan1"
        set egress-shaping-profile "Day_Hours_Profile"
        set outbandwidth 10000
    next
end
```
                                            ### Diagnose commands

###### To check that the specific traffic is put into the correct shaping group or class ID:

# diagnose firewall iprope list 100015

###### To check the speed limit for each class ID on an interface:

# diagnose netlink interface list wan1
