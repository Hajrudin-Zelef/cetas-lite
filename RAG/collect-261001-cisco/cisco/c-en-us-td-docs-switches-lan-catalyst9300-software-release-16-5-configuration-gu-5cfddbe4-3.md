---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4-3
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "distribution", "memory", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4.md
source_anchor: ""
source_lines: [223, 453]
sha256: 27b916fe8730a50fa28547d8d847ef63e48ddae96b911381ece8f58b8f52b149
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4

configureterminal
ip sla responder { tcp-connect | udp-echo} ipaddressip-addressportport-number
end
show running-config
copy running-config
startup-config
DETAILED STEPS
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
ip sla responder { tcp-connect | udp-echo} ipaddressip-addressportport-number
Example:
Device(config)# ip sla responder udp-echo 172.29.139.134 5000
Configures the
device as an IP SLA responder.
The keywords have
these meanings:
tcp-connect—Enables the responder for TCP connect
operations.
udp-echo—Enables the responder for User Datagram
Protocol (UDP) echo or jitter operations.
ipaddressip-address—Enter the destination IP address.
portport-number—Enter the destination port number.
Note
The IP address and port
number must match those configured on the source device for the IP SLA
operation.
Step 4
end
Example:
Device(config)# end
Returns to
privileged EXEC mode.
Step 5
show running-config
Example:
Device# show running-config
Verifies your entries.
Step 6
copy running-config
startup-config
Example:
Device# copy running-config startup-config
(Optional) Saves your entries
in the configuration file.
Implementing IP SLA
Network Performance Measurement
Follow these steps
to implement IP SLA network performance measurement on your
device:
Before you begin
Use the
show ip sla application privileged EXEC command to
verify that the desired operation type is supported on your software image.
Configures the IP
SLA operation as the operation type of your choice (a UDP jitter operation is
used in the example), and enters its configuration mode (UDP jitter
configuration mode is used in the example).
destination-ip-address |
destination-hostname—Specifies the destination IP
address or hostname.
destination-port—Specifies the destination port
number in the range from 1 to 65535.
(Optional)
source-ip {ip-address |
hostname} —Specifies the source IP address or
hostname. When a source IP address or hostname is not specified, IP SLA chooses
the IP address nearest to the destination
(Optional)
source-portport-number—Specifies the source port number in
the range from 1 to 65535. When a port number is not specified, IP SLA chooses
an available port.
(Optional)
control—Enables or disables sending of IP SLA
control messages to the IP SLA responder. By default, IP SLA control messages
are sent to the destination device to establish a connection with the IP SLA
responder
(Optional)
num-packetsnumber-of-packets—Enters the number of packets to
be generated. The range is 1 to 6000; the default is 10.
(Optional)
intervalinter-packet-interval—Enters the interval between
sending packets in milliseconds. The range is 1 to 6000; the default value is
20 ms.
Step 5
frequencyseconds
Example:
Device(config-ip-sla-jitter)# frequency 45
(Optional)
Configures options for the SLA operation. This example sets the rate at which a
specified IP SLA operation repeats. The range is from 1 to 604800 seconds; the
default is 60 seconds.
Step 6
thresholdmilliseconds
Example:
Device(config-ip-sla-jitter)# threshold 200
(Optional) Configures
threshold conditions. This example sets the threshold of the specified IP SLA
operation to 200. The range is from 0 to 60000 milliseconds.
Step 7
exit
Example:
Device(config-ip-sla-jitter)# exit
Exits the SLA
operation configuration mode (UDP jitter configuration mode in this example),
and returns to global configuration mode.
Step 8
ip sla scheduleoperation-number
[life {forever |
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
operation-number—Enter the RTR entry number.
(Optional)
life—Sets the operation to run indefinitely
(forever) or for a specific number of
seconds. The range is from 0 to 2147483647. The
default is 3600 seconds (1 hour).
(Optional)
start-time—Enters the time for the operation to
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
show that the operation should start after the entered time has elapsed.
(Optional)
ageoutseconds—Enter
the number of seconds to keep the operation in memory when it is not actively
collecting information. The range is 0 to 2073600 seconds, the default is 0
seconds (never ages out).
(Optional)
recurring—Set the operation to automatically run
every day.
Step 9
end
Example:
Device(config)# end
Returns to
privileged EXEC mode.
Step 10
show running-config
Example:
Device# show running-config
Verifies your entries.
Step 11
copy running-config
startup-config
Example:
Device# copy running-config startup-config
(Optional) Saves your entries
in the configuration file.
UDP Jitter
Configuration
This example shows how to configure a UDP jitter IP SLA
operation:
Device(config)# ip sla 10Device(config-ip-sla)# udp-jitter 172.29.139.134 5000Device(config-ip-sla-jitter)# frequency 30Device(config-ip-sla-jitter)# exitDevice(config)# ip sla schedule 5 start-time now life foreverDevice(config)# endDevice# show ip sla configuration 10
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
SLA operation as a UDP jitter operation, and enters UDP jitter configuration
mode.
destination-ip-address | destination-hostname: Specifies the destination IP address or hostname.
destination-port: Specifies the destination port number in the range from 1 to 65535.
(Optional) source-ip {ip-address | hostname} : Specifies the source IP address or hostname. When a source IP address or hostname is not specified, IP SLA chooses the IP
address nearest to the destination.
(Optional) source-portport-number: Specifies the source port number in the range from 1 to 65535. When a port number is not specified, IP SLA chooses an available
port.
Note
If the udp-jitter command does not have the source port configured, UDP chooses any random port for control packets. In case UDP chooses the
reserved port 1967, it may result in high CPU utilisation by the IP SLA responder.
(Optional) control: Enables or disables sending of IP SLA control messages to the IP SLA responder. By default, IP SLA control messages are
sent to the destination device to establish a connection with the IP SLA responder.
(Optional) num-packetsnumber-of-packets: Enters the number of packets to be generated. The range is 1 to 6000; the default is 10.
(Optional) intervalinter-packet-interval: Enters the interval between sending packets in milliseconds. The range is 1 to 6000; the default value is 20 ms.
Step 5
frequencyseconds
Example:
Device(config-ip-sla-jitter)# frequency 45
(Optional) Sets
