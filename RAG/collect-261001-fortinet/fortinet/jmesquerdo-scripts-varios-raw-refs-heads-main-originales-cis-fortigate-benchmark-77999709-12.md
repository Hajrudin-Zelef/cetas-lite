---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-12
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [1777, 1982]
sha256: 1b2b3b4abc19d9e82360fa37d1cc7e4325812bc523345f843458301d203e389f
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 63 
config firewall {local-in-policy | local-in-policy6} 
    edit <policy_number> 
        set intf <interface> 
        set srcaddr <source_address> [source_address] ... 
        set dstaddr <destination_address> [destination_address] ... 
        set action {accept | deny} 
        set service <service_name> [service_name] ... 
        set schedule <schedule_name> 
        set comments <string> 
    next 
end 
For example, to prevent the source subnet 10.10.10.0/24 from pinging port1, but allow 
administrative access for PING on port1: 
config firewall address 
    edit "10.10.10.0" 
        set subnet 10.10.10.0 255.255.255.0 
    next 
end 
config firewall local-in-policy 
    edit 1 
        set intf "port1" 
        set srcaddr "10.10.10.0" 
        set dstaddr "all" 
        set service "PING" 
        set schedule "always" 
    next 
end 
Default Value: 
There are no Local-in Policies in place by default

Page 64 
2.5 High Availability 
High Availability (HA) subsection includes configurations for High Availability between 
FortiGate devices

Page 65 
2.5.1 Ensure High Availability Configuration (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Ensure that FortiGate devices are configured for High Availability (HA). 
Rationale: 
Configuring High Availability (HA) increases system availability as well as decreases 
impact of routine maintenance (Firmware updates, cable moves, etc.) and the the 
impact of device failure. 
Impact: 
Not having High Availability (HA) configured correctly and synced properly impacts the 
availability of the FortiGate devices as well as any systems that require traversing the 
FortiGates. With properly configured HA in place outages can be minimized during 
firmware updates as well as if there are power outages or device failures. 
Audit: 
In GUI: 
Navigate to "System" and then "HA" 
Ensure "Mode" is set to proper setting "Active-Active" or "Active-Passive" 
Review Configuration settings 
 "Cluster Name" must match on devices 
 "Password" Must match on devices 
 "Heartbeat Interfaces" need to be defined on devices 
Click "OK" to save changes and exit 
In CLI: 
FGT1 # config system ha  
FGT1 (ha) # set mode a-p        ###(Active-Passive) 
FGT1 (ha) # set group-name "FGT-HA"            ###(Set cluster name) 
FGT1 (ha) # set password *******   ###(Set password)  
FGT1 (ha) # set hbdev port10 50   ###(Set Heartbeat 
Interface and priority) 
FGT1 (ha) # end 
To review configuration in CLI

Page 66 
FGT1 # config system ha 
FGT1 (ha) # show 
config system ha 
    set group-name "FGT-HA" 
    set mode a-p 
    set password ENC 
enrwD467hJmO6j6YW/l6FEOa1YNVYdo8Z5mCcTDEKUFpOVXcNYnPBmQDGX//ViXk6TkwNH0il5aJr
/fZY25lq+husndQHZVWp2LIlXmCv/n81U43nkZUWaIKvqkellGFbhv0/IHoOLzQPCsVcBbyrsgopr
YMvh6w7F06+nRriBtMNQxpOV5V+e388EcwsOOMsXBZOw== 
    set hbdev "port10" 50  
    set override disable 
end 
Remediation: 
In GUI: 
Navigate to "System" and then "HA" 
Ensure "Mode" is set to proper setting "Active-Active" or "Active-Passive" 
Review Configuration settings 
 "Cluster Name" must match on devices 
 "Password" Must match on devices 
 "Heartbeat Interfaces" need to be defined on devices 
Click "OK" to save changes and exit 
In CLI: 
FGT1 # config system ha  
FGT1 (ha) # set mode a-p        ###(Active-Passive) 
FGT1 (ha) # set group-name "FGT-HA"            ###(Set cluster name) 
FGT1 (ha) # set password *******   ###(Set password)  
FGT1 (ha) # set hbdev port10 50   ###(Set Heartbeat 
Interface and priority) 
FGT1 (ha) # end 
To review configuration in CLI 
FGT1 # config system ha 
FGT1 (ha) # show 
config system ha 
    set group-name "FGT-HA" 
    set mode a-p 
    set password ENC 
enrwD467hJmO6j6YW/l6FEOa1YNVYdo8Z5mCcTDEKUFpOVXcNYnPBmQDGX//ViXk6TkwNH0il5aJr
/fZY25lq+husndQHZVWp2LIlXmCv/n81U43nkZUWaIKvqkellGFbhv0/IHoOLzQPCsVcBbyrsgopr
YMvh6w7F06+nRriBtMNQxpOV5V+e388EcwsOOMsXBZOw== 
    set hbdev "port10" 50  
    set override disable 
end 
Default Value: 
N/A

Page 67 
References: 
1. https://docs.fortinet.com/document/fortigate/6.4.5/administration-
guide/489119/ha-cluster-setup-examples

Page 68 
2.5.2 Ensure "Monitor Interfaces" for High Availability Devices is 
Enabled (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Configure Interface Monitoring within High Availability settings, Interface Monitoring 
should be enabled on all critical interfaces. 
Rationale: 
With Interface Monitoring enabled on devices failover can occur if there are physical 
media issues or issues with the specific port that the FortiGate is connected to. 
Impact: 
Not configuring Interface Monitoring can directly impact services due to a failure to 
trigger a High Availability failover if an interface is impacted only on the primary device 
and it is not being monitored. Without the Interface monitoring enabled failover would be 
limited to hardware, system, or power faults. 
Audit: 
To Validate from GUI: 
go to System - > HA 
Under "Monitor Interfaces" validate all applicable interfaces are selected 
select "OK" 
To Validate from CLI: 
FGT1 # config system ha 
FGT1 (ha) # show 
config system ha 
    set group-name "FGT-HA" 
    set mode a-p 
    set password ENC 
enrwD467hJmO6j6YW/l6FEOa1YNVYdo8Z5mCcTDEKUFpOVXcNYnPBmQDGX//ViXk6TkwNH0il5aJr
/fZY25lq+husndQHZVWp2LIlXmCv/n81U43nkZUWaIKvqkellGFbhv0/IHoOLzQPCsVcBbyrsgopr
YMvh6w7F06+nRriBtMNQxpiTE+12xAHz7lA3EoYZzf8A== 
    set override disable 
    set monitor "port6" "port7"   ###Validate proper interfaces are present 
end 
Remediation: 
To Remediate from GUI:

Page 69 
go to System - > HA 
Under "Monitor Interfaces" select all applicable interfaces. 
select "OK" 
To Validate from CLI: 
FGT1 # config system ha 
FGT1 (ha) # set monitor "port6" "port7" 
FGT1 (ha) # show  ###To Review changes to monitored interfaces before 
applying 
config system ha 
    set group-name "FGT-HA" 
    set mode a-p 
    set password ENC 
enrwD467hJmO6j6YW/l6FEOa1YNVYdo8Z5mCcTDEKUFpOVXcNYnPBmQDGX//ViXk6TkwNH0il5aJr
/fZY25lq+husndQHZVWp2LIlXmCv/n81U43nkZUWaIKvqkellGFbhv0/IHoOLzQPCsVcBbyrsgopr
YMvh6w7F06+nRriBtMNQxpiTE+12xAHz7lA3EoYZzf8A== 
    set override disable 
    set monitor "port6" "port7"    
end 
Default Value: 
N/A 
References: 
1. https://docs.fortinet.com/document/fortigate/6.0.0/best-
practices/498515/interface-monitoring-port-monitoring

Page 70 
2.5.3 Ensure HA Reserved Management Interface is Configured 
(Manual) 
Profile Applicability: 
•  Level 1 
•  Level 2 
Description: 
Ensure Reserved Management Interfaces are configured on HA devices 
Rationale: 
To be able to access both the primary and secondary firewalls in an HA cluster 
Reserved Management Interfaces need to be configured to prevent them from syncing 
with HA and sharing a virtual MAC address 
Impact: 
Not configuring reserved Management Interfaces impacts the ability to access 
secondary devices directly due to the primary and secondary devices syncing 
configuration exactly and floating a virtualized mac address between them for failover 
Audit: 
Review through the GUI: 
go to System -> HA edit the "Master" device and verify that "Management 
Interface Reservation" is selected and there is an interface, and gateway 
defined 
Review through the CLI:

