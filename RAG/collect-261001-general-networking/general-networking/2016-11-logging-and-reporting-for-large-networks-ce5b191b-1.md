---
id: collect-261001-general-networking/general-networking/2016-11-logging-and-reporting-for-large-networks-ce5b191b-1
title: "2016-11-logging-and-reporting-for-large-networks-ce5b191b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/2016-11-logging-and-reporting-for-large-networks-ce5b191b.md
source_anchor: ""
source_lines: [1, 193]
sha256: 6aa3d198b253910cf94017149a63d09fbdf7cfc86cf4cf24a5b988cc05483ee4
---

# 2016-11-logging-and-reporting-for-large-networks-ce5b191b

**Logg****i****n****g and reporting for large networks**

This section explains how to configure the FortiGate unit for logging and reporting in a larger network, such as an enterprise network. To set up this type of network, you are modifying the default log settings, and you are also modifying the default report.

The following procedures are examples and can be used to help you when configuring your own network’s log topology.

Since some of these settings must be modified or enabled or disabled in the CLI, it is recommended to review the FortiGate CLI Reference for any additional information about the commands used herein, as well as any that you would need to use in your own newtork’s log topology.


**M****od****i****f****y****i****n****g default log device settings**

The default log device settings must be modified so that system performance is not compromised. The FortiGate unit, by default, has all logging of FortiGate features enabled and well as logging to either the FortiGate unit’s system memory or hard disk, depending on the model.


**M****od****i****f****y****i****n****g multiple FortiGate units’ system memory default settings**

When the FortiGate unit’s default log device is its system memory, you can modify it to fit your log network topology. In this topic, the following is an example of how you can modify these default settings.


**T****o modify the default system memory settings**

**1****.** Log in to the CLI.

**2****.** Enter the following command syntax to modify the logging settings:

config log memory setting set ips-archive disable set status enable

end

**3****.** Enter the following command syntax to modify the FortiGate features that are enabled for logging:

config log memory filter set attack enable

set forward-traffic enable set local-traffic enable set netscan enable

set email-log-imap enable

set multicast-traffic enable set scanerror enable

set app-ctrl enable end

**4****.** Repeat steps 2 and 3 for the other FortiGate units.

**5****.** Test the modified settings using the procedure below.


**M****od****i****f****y****i****n****g multiple FortiGate units’ hard disk default log settings**

You will have to modify each FortiGate unit’s hard disk default log settings. The following is an example of how to modify these default settings.


**T****o modify the default hard disk settings**

**1****.** Log in to the CLI.

**2****.** Enter the following command syntax to modify the logging settings:

config log disk setting

set ips-archive disable set status enable

set max-log-file-size 1000 set storage Internal

set log-quota 100

set report-quota 100 end

**3****.** In the CLI, enter the following to disable certain event log messages that you do not want logged:

config log disk filter

set sniffer-traffic disable set local-traffic enable

end

**4****.** Repeat the steps 2 to 4 for the other FortiGate units.

**5****.** Test the modified settings using the procedure below.


**T****es****t****i****n****g the modified log settings**

After modifying both the settings and the FortiGate features for logging, you can test that the modified settings are working properly. This test is done in the CLI.


**T****o test sending logs to the log device**

**1****.** In the CLI, enter the following command syntax:

diag log test

When you enter the command, the following appears:

generating a system event message with level – warning generating an infected virus message with level – warning generating a blocked virus message with level – warning generating a URL block message with level – warning generating a DLP message with level – warning

generating an IPS log message generating an anomaly log message

generating an application control IM message with level – information generating an IPv6 application control IM message with level – information generating deep application control logs with level – information generating an antispam message with level – notification

generating an allowed traffic message with level – notice generating a multicast traffic message with level – notice generating a ipv6 traffic message with level – notice

generating a wanopt traffic log message with level – notification

generating a HA event message with level – warning generating netscan log messages with level – notice generating a VOIP event message with level – information generating a DNS event message with level – information generating authentication event messages

generating a Forticlient message with level – information generating a NAC QUARANTINE message with level – notification generating a URL block message with level – warning

**2****.** In the web-based interface, go to **Lo****g & Report > Event Log > User**, and view the logs to see some of the recently generated test log messages.

You will be able to tell the test log messages from real log messages because they do not have “real” information;

for example, the test log messages for the vulnerability scan contain the destination IP address of 1.1.1.1 or 2.2.2.2.


**C****on****f****i****gu****r****i****n****g the backup solution**

Even though you are logging to multiple FortiAnalyzer units, this is more of a redundancy solution rather than a complete backup solution in this example.

The multiple FortiAnalyzer units act similar to a HA cluster, since if one FortiAnalyzer unit fails, the others continue storing the logs they receive. In a backup solution, the logs are backed up to another secure location if something happens to the log device.

A good alternate or redundant option is the FortiCloud service, which can provide secure online logging and management for multiple devices.


**C****on****f****i****gu****r****i****n****g logging to multiple FortiAnalyzer units**

The following example shows how to configure logging to multiple FortiAnalyzer units. Configuring multiple FortiAnalyzer units is quick and easy; however, you can only configure up to three FortiAnalyzer units per FortiGate unit.


**T****o configure multiple FortiAnalyzer units**

**1****.** In the CLI, enter the following command syntax to configure the first FortiAnalyzer unit:

config log fortianalyzer setting set status enable

set server 172.20.120.22 set max-buffer-size 1000 set buffer-max-send 2000 set address-mode static set conn-timeout 100

set monitor-keepalive-period 120

set monitor-failure-retry-period 2000

end

**2****.** Disable the features that you do not want logged, using the following example command syntax. You can view the

CLI Reference to see what commands are available.

config log fortianalyzer filter set traffic (enable | disable)

… end

**3****.** Enter the following commands for the second FortiAnalyzer unit:

config log fortianalyzer2 setting set status enable

set server 172.20.120.23 set max-buffer-size 1000 set buffer-max-send 2000 set address-mode static set conn-timeout 100

set monitor-keepalive-period 120

set monitor-failure-retry-period 2000

end

**4****.** Disable the features that you do not want logged, using the following example command syntax.

config log fortianalyzer filter set web (enable | disable)

… end

**5****.** Enter the following commands for the last FortiAnalyzer unit:

config log fortianalyzer3 setting set status enable

set server 172.20.120.23 set max-buffer-size 1000 set buffer-max-send 2000 set address-mode static set conn-timeout 100

set monitor-keepalive-period 120

set monitor-failure-retry-period 2000

end

**6****.** Disable the features that you do not want logged, using the following example command syntax.

config log fortianalyzer filter

set web-filter (enable | disable)

… end

**7****.** Test the configuration by using the procedure, “Testing the modified log settings”.

**8****.** On the other FortiGate units, configure steps 1 through 6, ensuring that logs are being sent to the FortiAnalyzer units.


