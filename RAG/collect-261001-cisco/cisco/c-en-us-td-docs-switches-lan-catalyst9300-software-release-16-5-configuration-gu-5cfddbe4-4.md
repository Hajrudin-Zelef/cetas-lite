---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4-4
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "distribution", "ethernet", "memory", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4.md
source_anchor: ""
source_lines: [454, 712]
sha256: 531a18db2884ee33bbb0a9aa95cbe68a3ef53544a2c6b69a75b2c800d188444b
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4

the rate at which a specified IP SLA operation repeats. The range is from 1 to
604800 seconds; the default is 60 seconds.
Step 6
exit
Example:
Device(config-ip-sla-jitter)# exit
Exits UDP jitter
configuration mode, and returns to global configuration mode.
Step 7
ip sla scheduleoperation-number [life {forever |
seconds}] [start-time {hh:mm
[ :ss] [month day |
day month] |
pending |
now |
afterhh:mm:ss] [ageoutseconds] [recurring]
Example:
Device(config)# ip sla schedule 10 start-time now life forever
Configures the
scheduling parameters for an individual IP SLA operation.
operation-number: Enter the RTR entry number.
(Optional) life: Sets the operation to run indefinitely (forever) or for a specific number of seconds. The range is from 0 to 2147483647. The default is 3600 seconds (1 hour).
(Optional) start-time: Enters the time for the operation to begin collecting information:
To start at a
specific time, enter the hour, minute, second (in 24-hour notation), and day of
the month. If no month is entered, the default is the current month.
Enter
pending to select no information collection until
a start time is selected.
Enter
now to start the operation immediately.
Enter
afterhh:mm:ss to
show that the operation should start after the entered time has elapsed.
(Optional) ageoutseconds: Enter the number of seconds to keep the operation in memory when it is not actively collecting information. The range is
0 to 2073600 seconds, the default is 0 seconds (never ages out).
(Optional) recurring: Set the operation to automatically run every day.
Step 8
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
Step 9
show
running-config
Example:
Device# show running-config
Verifies your entries.
Step 10
copy running-config
startup-config
Example:
Device# copy running-config startup-config
(Optional) Saves
your entries in the configuration file.
Configuring a
UDP Jitter IP SLA Operation
This example shows how to
configure a UDP jitter IP SLA operation:
Device(config)# ip sla 10
Device(config-ip-sla)# udp-jitter 172.29.139.134 5000 source-ip 172.29.139.140 source-port 4000
Device(config-ip-sla-jitter)# frequency 30
Device(config-ip-sla-jitter)# exit
Device(config)# ip sla schedule 10 start-time now life forever
Device(config)# end
Device# show ip sla configuration 10
IP SLAs, Infrastructure Engine-II.
Entry number: 10
Owner:
Tag:
Type of operation to perform: udp-jitter
Target address/Source address: 10.0.0.10/10.0.0.1
Target port/Source port: 2/0
Request size (ARR data portion): 32
Operation timeout (milliseconds): 5000
Packet Interval (milliseconds)/Number of packets: 20/10
Type Of Service parameters: 0x0
Verify data: No
Vrf Name:
Control Packets: enabled
Schedule:
Operation frequency (seconds): 30
Next Scheduled Start Time: Pending trigger
Group Scheduled : FALSE
Randomly Scheduled : FALSE
Life (seconds): 3600
Entry Ageout (seconds): never
Recurring (Starting Everyday): FALSE
Status of entry (SNMP RowStatus): notInService
Threshold (milliseconds): 5000
Distribution Statistics:
Number of statistic hours kept: 2
Number of statistic distribution buckets kept: 1
Statistic distribution interval (milliseconds): 20
Enhanced History:
Configures the IP
SLA operation as an ICMP Echo operation and enters ICMP echo configuration
mode.
destination-ip-address |
destination-hostname—Specifies the destination IP
address or hostname.
(Optional)
source-ip {ip-address |
hostname} —Specifies the source IP address or
hostname. When a source IP address or hostname is not specified, IP SLA chooses
the IP address nearest to the destination.
(Optional)
source-interfaceinterface-id—Specifies the source interface for
the operation.
Step 5
frequencyseconds
Example:
Device(config-ip-sla-echo)# frequency 30
(Optional) Sets
the rate at which a specified IP SLA operation repeats. The range is from 1 to
604800 seconds; the default is 60 seconds.
Step 6
exit
Example:
Device(config-ip-sla-echo)# exit
Exits UDP echo
configuration mode, and returns to global configuration mode.
Step 7
ip sla scheduleoperation-number [life {forever |
seconds}] [start-time {hh:mm
[:ss] [month day |
day month] |
pending |
now |
afterhh:mm:ss] [ageoutseconds] [recurring]
Example:
Device(config)# ip sla schedule 5 start-time now life forever
Configures the
scheduling parameters for an individual IP SLA operation.
operation-number—Enter the RTR entry number.
(Optional)
life—Sets the operation to run indefinitely
(forever) or for a specific number of
seconds. The range is from 0 to 2147483647. The
default is 3600 seconds (1 hour)
(Optional)
start-time—Enter the time for the operation to
begin collecting information:
To start at a
specific time, enter the hour, minute, second (in 24-hour notation), and day of
the month. If no month is entered, the default is the current month.
Enter
pending to select no information collection until
a start time is selected.
Enter
now to start the operation immediately.
Enter
afterhh:mm:ss to
indicate that the operation should start after the entered time has elapsed.
(Optional)
ageoutseconds—Enter
the number of seconds to keep the operation in memory when it is not actively
collecting information. The range is 0 to 2073600 seconds; the default is 0
seconds (never ages out).
(Optional)
recurring—Sets the operation to automatically run
every day.
Step 8
end
Example:
Device(config)# end
Returns to
privileged EXEC mode.
Step 9
show running-config
Example:
Device# show running-config
Verifies your entries.
Step 10
copy running-config
startup-config
Example:
Device# copy running-config startup-config
(Optional) Saves your entries
in the configuration file.
Configuring an
ICMP Echo IP SLA Operation
This example shows how to
configure an ICMP echo IP SLA operation:
Device(config)# ip sla 12Device(config-ip-sla)# icmp-echo 172.29.139.134 Device(config-ip-sla-echo)# frequency 30Device(config-ip-sla-echo)# exitDevice(config)# ip sla schedule 5 start-time now life foreverDevice(config)# endDevice# show ip sla configuration 22
IP SLAs, Infrastructure Engine-II.
Entry number: 12
Owner:
Tag:
Type of operation to perform: echo
Target address: 2.2.2.2
Source address: 0.0.0.0
Request size (ARR data portion): 28
Operation timeout (milliseconds): 5000
Type Of Service parameters: 0x0
Verify data: No
Vrf Name:
Schedule:
Operation frequency (seconds): 60
Next Scheduled Start Time: Pending trigger
Group Scheduled : FALSE
Randomly Scheduled : FALSE
Life (seconds): 3600
Entry Ageout (seconds): never
Recurring (Starting Everyday): FALSE
Status of entry (SNMP RowStatus): notInService
Threshold (milliseconds): 5000
Distribution Statistics:
Number of statistic hours kept: 2
Number of statistic distribution buckets kept: 1
Statistic distribution interval (milliseconds): 20
History Statistics:
Number of history Lives kept: 0
Number of history Buckets kept: 15
History Filter Type: None
Enhanced History:
The following
table describes the commands used to display IP SLA operation configurations
and results:
Table 1. Monitoring IP SLA
Operations
show ip sla
application
Displays global information
about Cisco IOS IP SLAs.
show ip sla
authentication
Displays IP SLA
authentication information.
show ip sla
configuration [entry-number]
Displays configuration values
including all defaults for all IP SLA operations or a specific operation.
show ip sla
enhanced-history {collection-statistics |
distribution
statistics} [entry-number]
Displays enhanced history
statistics for collected history buckets or distribution statistics for all IP
SLA operations or a specific operation.
show ip sla ethernet-monitor
configuration [entry-number]
Displays IP SLA automatic
Ethernet configuration.
show ip sla group
schedule [schedule-entry-number]
Displays IP SLA group
scheduling configuration and details.
show ip sla history
[entry-number |
full |
tabular]
Displays history collected
for all IP SLA operations.
show ip sla mpls-lsp-monitor
{collection-statistics |
configuration |
ldp operational-state |
scan-queue |
