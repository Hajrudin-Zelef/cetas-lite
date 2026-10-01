---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-routers-ncs4200-configuration-guide-security-17-1-1-b-sec-usr-ra-6a5f8490-1
title: "c-en-us-td-docs-routers-ncs4200-configuration-guide-security-17-1-1-b-sec-usr-ra-6a5f8490"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-routers-ncs4200-configuration-guide-security-17-1-1-b-sec-usr-ra-6a5f8490.md
source_anchor: ""
source_lines: [1, 131]
sha256: 99d829b28578954eba6cf6e85bcc2e79c1b5da2ccdd5e561900c6803f1fcbeab
---

# c-en-us-td-docs-routers-ncs4200-configuration-guide-security-17-1-1-b-sec-usr-ra-6a5f8490

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
Configuring a device to use authentication, authorization, and accounting (AAA) server groups provides a way to group existing
server hosts. Grouping existing server hosts allows you to select a subset of the configured server hosts and use them for
a particular service. Configuring deadtime within a server group allows you to direct AAA traffic to separate groups of servers
that have different operational characteristics. This feature module describes how to configure AAA server groups and the
deadtimer.
Your software release may not support all the features documented in this module. For
the latest caveats and feature information, see Bug Search
Tool and the release notes for your platform and software release. To find
information about the features documented in this module, and to see a list of the
releases in which each feature is supported, see the feature information table.
Use Cisco Feature Navigator to find information about platform support and Cisco software
image support. To access Cisco Feature Navigator, go to https://cfnng.cisco.com/. An account on
Cisco.com is not required.
Configuring the device to use AAA server groups provides a way to group existing server hosts. Grouping existing server hosts
allows you to select a subset of the configured server hosts and use them for a particular service. A server group is used
with a global server-host list. The server group lists the IP addresses of the selected server hosts.
Server groups can also include multiple host entries for the same server, as long as each entry has a unique identifier.
The combination of an IP address and a UDP port number creates a unique identifier, allowing different ports to be individually
defined as RADIUS hosts providing a specific AAA service. This unique identifier enables RADIUS requests to be sent to different
UDP ports on a server at the same IP address. If two different host entries on the same RADIUS server are configured for the
same service—for example, accounting—the second host entry that is configured acts as a failover backup to the first one.
If the first host entry fails to provide accounting services, the network access server tries the second host entry configured
on the same device for accounting services. (The RADIUS host entries are tried in the order in which they are configured.)
AAA Server Groups with a Deadtimer
After you configure a server host with a server name, you can use the
deadtime command to configure each server per server group. Configuring deadtime within a server group allows you to direct AAA traffic
to separate groups of servers that have different operational characteristics.
Configuring deadtime is not limited to a global configuration. A separate timer is attached to each server host in every
server group. Therefore, when a server is found to be unresponsive after numerous retransmissions and timeouts, the server
is assumed to be dead. The timers attached to each server host in all server groups are triggered. In essence, the timers
are checked and subsequent requests to a server (once it is assumed to be dead) are directed to alternate timers, if configured.
When the network access server receives a reply from the server, it checks and stops all configured timers (if running) for
that server in all server groups.
If the timer has expired, the server to which the timer is attached is assumed to be alive. This becomes the only server
that can be tried for later AAA requests using the server groups to which the timer belongs.
Note
Because one server has different timers and might have different deadtime values configured in the server groups, the same
server might, in the future, have different states (dead and alive) at the same time.
Note
To change the state of a server, you must start and stop all configured timers in all server groups.
The size of the server group will be slightly increased because of the addition of new timers and the deadtime attribute.
The overall impact of the structure depends on the number and size of the server groups and how the servers are shared among
server groups in a specific configuration.
To define a server
host with a server group name, enter the following commands in global
configuration mode. The listed server must exist in global configuration mode.
Before you begin
Each server in the
group must be defined previously using the
radius-server
host command.
Device(config-sg-radius)# server 172.16.1.1 acct-port 1616
Associates a
particular RADIUS server with the defined server group.
Each
security server is identified by its IP address and UDP port number.
Repeat
this step for each RADIUS server in the AAA server group.
Step 6
end
Example:
Device(config-sg-radius)# end
Exits server
group RADIUS configuration mode and returns to privileged EXEC mode.
Configuring AAA Server Groups with a Deadtimer
SUMMARY STEPS
enable
configureterminal
aaagroupserverradiusgroup
deadtimeminutes
end
DETAILED STEPS
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password, if prompted.
Step 2
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
aaagroupserverradiusgroup
Example:
Device(config)# aaa group server radius group1
Defines a RADIUS type server group and enters server group RADIUS configuration mode.
Step 4
deadtimeminutes
Example:
Device(config-sg-radius)# deadtime 1
Configures and defines a deadtime value in minutes.
Note
Local server group deadtime overrides the global configuration. If the deadtime vlaue is omitted from the local server group
configuration, it is inherited from the primary list.
Step 5
end
Example:
Device(config-sg-radius)# end
Exits the server group RADIUS configuration mode and returns to the privileged EXEC mode.
The following example shows how to create server group radgroup1 with three different RADIUS server members, each using the
default authentication port (1645) and accounting port (1646):
aaa group server radius radgroup1
server 172.16.1.11
server 172.17.1.21
server 172.18.1.31
The following example shows how to create server group radgroup2 with three RADIUS server members, each with the same IP
address but with unique authentication and accounting ports:
aaa group server radius radgroup2
server 172.16.1.1 auth-port 1000 acct-port 1001
server 172.16.1.1 auth-port 2000 acct-port 2001
server 172.16.1.1 auth-port 3000 acct-port 3001
Example: Multiple RADIUS Server Entries Using AAA Server Groups
The following example shows how to configure the network access server to recognize two different RADIUS server groups. One
of these groups, group1, has two different host entries on the same RADIUS server configured for the same services. The second
host entry configured acts as failover backup to the first one. Each group is individually configured for the deadtime; the
deadtime for group 1 is one minute, and the deadtime for group 2 is two minutes.
Note
In cases where both global commands and
server commands are used, the
server command takes precedence over the global command.
! This command enables AAA.
aaa new-model
! The next command configures default RADIUS parameters.
aaa authentication ppp default group group1
! The following commands define the group1 RADIUS server group and associate servers
! with it and configures a deadtime of one minute.
