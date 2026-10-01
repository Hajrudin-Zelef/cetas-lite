---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-b0e646e9-1
title: "c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-b0e646e9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["incident", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-b0e646e9.md
source_anchor: ""
source_lines: [1, 112]
sha256: 8a03fb6116d573a114aa921953f2e3758a500a5238c01dcfc6c19964c0db4275
---

# c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-b0e646e9

About Logging
System logging is a method of collecting messages from devices to a server running a syslog daemon. Logging to a central syslog server helps in aggregation of logs and alerts. Cisco devices can send their log messages to a UNIX-style syslog service. A syslog service accepts messages and stores them in files, or prints them according to a simple configuration file. This form of logging provides protected long-term storage for logs. Logs are useful both in routine troubleshooting and in incident handling.
The ASA system logs provide you with information for monitoring and troubleshooting the ASA. With the logging feature, you can do the following:
-  
                              		  
                              Specify which syslog messages should be logged.
-  
                              		  
                              Disable or change the severity level of a syslog message.
-  
                              		  
                              Specify one or more locations where syslog messages should be sent, including: 
  -  
                                       				
                                       An internal buffer
  -  
                                       				
                                       One or more syslog servers
  -  
                                       				
                                       ASDM
  -  
                                       				
                                       An SNMP management station
  -  
                                       				
                                       Specified e-mail addresses
  - 
                                       				
                                       Console
  -  
                                       				
                                       Telnet and SSH sessions.
-  
                                       				
                                       
-  
                              		  
                              Configure and manage syslog messages in groups, such as by severity level or class of message.
-  
                              		  
                              Specify whether or not a rate-limit is applied to syslog generation.
-  
                              		  
                              Specify what happens to the contents of the internal log buffer when it becomes full: overwrite the buffer, send the buffer contents to an FTP server, or save the contents to internal flash memory.
-  
                              		  
                              Filter syslog messages by locations, severity level, class, or a custom message list.
Logging in Multiple Context Mode
Each security context includes its own logging configuration and generates its own messages. If you log in to the system or admin context, and then change to another context, messages you view in your session are only those messages that are related to the current context.
Syslog messages that are generated in the system execution space, including failover messages, are viewed in the admin context along with messages generated in the admin context. You cannot configure logging or view any logging information in the system execution space.
You can configure the ASA to include the context name with each message, which helps you differentiate context messages that are sent to a single syslog server. This feature also helps you to determine which messages are from the admin context and which are from the system; messages that originate in the system execution space use a device ID of system, and messages that originate in the admin context use the name of the admin context as the device ID.
Syslog Message Analysis
The following are some examples of the type of information you can obtain from a review of various syslog messages:
-  
                                 		  
                                 Connections that are allowed by ASA security policies. These messages help you spot holes that remain open in your security policies.
-  
                                 		  
                                 Connections that are denied by ASA security policies. These messages show what types of activity are being directed toward your secured inside network.
- 
                                 				
                                 Using the ACE deny rate logging feature shows attacks that are occurring on your ASA.
-  
                                 		  
                                 IDS activity messages can show attacks that have occurred.
-  
                                 		  
                                 User authentication and command usage provide an audit trail of security policy changes.
-  
                                 		  
                                 Bandwidth usage messages show each connection that was built and torn down as well as the duration and traffic volume used.
-  
                                 		  
                                 Protocol usage messages show the protocols and port numbers used for each connection.
-  
                                 		  
                                 Address translation audit trail messages record NAT or PAT connections being built or torn down, which are useful if you receive a report of malicious activity coming from inside your network to the outside world.
Syslog Message Format
Syslog messages begin with a percent sign (%) and are structured as follows:
%ASA Level Message_number: Message_text
Field descriptions are as follows:
| ASA | The syslog message facility code for messages that are generated by the ASA. This value is always ASA. | 
| Level | 1 through 7. The level reflects the severity of the condition described by the syslog message—the lower the number, the more severe the condition. | 
| Message_number | A unique six-digit number that identifies the syslog message. | 
| Message_text | A text string that describes the condition. This portion of the syslog message sometimes includes IP addresses, port numbers, or usernames. | 
Severity Levels
The following table lists the syslog message severity levels. You can assign custom colors to each of the severity levels to make it easier to distinguish them in the ASDM log viewers. To configure syslog message color settings, either choose the Tools > Preferences > Syslog tab or, in the log viewer itself, click Color Settings on the toolbar.
| Table 1. Syslog Message Severity Levels |  |  | 
|---|---|---|
| Level Number | Severity Level | Description | 
|---|---|---|
| 0 | emergencies | System is unusable. | 
| 1 | alert | Immediate action is needed. | 
| 2 | critical | Critical conditions. | 
| 3 | error | Error conditions. | 
| 4 | warning | Warning conditions. | 
| 5 | notification | Normal but significant conditions. | 
| 6 | informational | Informational messages only. | 
| 7 | debugging | Debugging messages only. Log at this level only temporarily, when debugging issues. This log level can potentially generate so many messages that system performance can be affected. | 
| Note | ASA does not generate syslog messages with a severity level of zero (emergencies). | 
Syslog Message Filtering
You can filter generated syslog messages so that only certain syslog messages are sent to a particular output destination. For example, you could configure the ASA to send all syslog messages to one output destination and to send a subset of those syslog messages to a different output destination.
Specifically, you can direct syslog messages to an output destination according to the following criteria:
-  
                                 		  
                                 Syslog message ID number
-  
                                 		  
                                 Syslog message severity level
-  
                                 		  
