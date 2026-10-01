---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-10-configuration-g-197f43df
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-10-configuration-g-197f43df"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-10-configuration-g-197f43df.md
source_anchor: ""
source_lines: [1, 60]
sha256: 7c291c037dbeaedb70390ed889381279ad5d724a8ff26499c44c5b41db05b60d
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-10-configuration-g-197f43df

Information About Configuring System Message Logs
System Messsage Logging
By default, a switch sends the output from system messages and debug privileged EXEC commands to a logging process. . The logging process controls the distribution of logging messages to various destinations, such as the logging buffer, terminal lines, or a UNIX syslog server, depending on your configuration. The process also sends messages to the console.
When the logging process is disabled, messages are sent only to the console. The messages are sent as they are generated, so message and debug output are interspersed with prompts or output from other commands. Messages appear on the active consoles after the process that generated them has finished.
You can set the severity level of the messages to control the type of messages displayed on the consoles and each of the destinations. You can time-stamp log messages or set the syslog source address to enhance real-time debugging and management. For information on possible messages, see the system message guide for this release.
You can access logged system messages by using the switch command-line interface (CLI) or by saving them to a properly configured syslog server. The switch software saves syslog messages in an internal buffer on a standalone switch. If a standalone switch , the log is lost unless you had saved it to flash memory.
You can remotely monitor system messages by viewing the logs on a syslog server or by accessing the switch through Telnet, through the console port, or through the Ethernet management port.
| Note | The syslog format is compatible with 4.3 BSD UNIX. | 
System Log Message Format
System log messages can contain up to 80 characters and a percent sign (%), which follows the optional sequence number or time-stamp information, if configured. Depending on the switch, messages appear in one of these formats:
- 
                                 
                                 seq no:timestamp: %facility-severity-MNEMONIC:description (hostname-n)
- 
                                 
                                 seq no:timestamp: %facility-severity-MNEMONIC:description
- 
                                    
                                    service sequence-numbers
- 
                                    
                                    service timestamps log datetime
- 
                                    
                                    service timestamps log datetime [localtime] [msec] [show-timezone]
- 
                                    
                                    service timestamps log uptime
| Table 1. System Log Message Elements |  | 
|---|---|
| Element | Description | 
|---|---|
| seq no: | Stamps log messages with a sequence number only if the service sequence-numbers global configuration command is configured. | 
| timestamp formats: mm/dd h h:mm:ss or hh:mm:ss (short uptime) or d h (long uptime) | Date and time of the message or event. This information appears only if the service timestamps log [datetime \| log] global configuration command is configured. | 
| facility | The facility to which the message refers (for example, SNMP, SYS, and so forth). | 
| severity | Single-digit code from 0 to 7 that is the severity of the message. | 
| MNEMONIC | Text string that uniquely describes the message. | 
| description | Text string containing detailed information about the event being reported. | 
| hostname-n | Hostname of a stack member and its switch number in the stack. Though the active switch is a stack member, it does not append its hostname to system messages. | 
Default System Message Logging Settings
| Table 2. Default System Message Logging Settings |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| System message logging to the console | Enabled. | 
| Console severity | Debugging. | 
| Logging file configuration | No filename specified. | 
| Logging buffer size | 102400 bytes. | 
| Logging history size | 1 message. | 
| Time stamps | Disabled. | 
| Synchronous logging | Disabled. | 
| Logging server | Disabled. | 
| Syslog server IP address | None configured. | 
| Server facility | Local7 | 
| Server severity | Informational. | 
Syslog Message Limits
If you enabled syslog message traps to be sent to an SNMP network management station by using the snmp-server enable trap global configuration command, you can change the level of messages sent and stored in the switch history table. You also can change the number of messages that are stored in the history table.
Messages are stored in the history table because SNMP traps are not guaranteed to reach their destination. By default, one message of the level warning and numerically lower levels are stored in the history table even if syslog traps are not enabled.
When the history table is full (it contains the maximum number of message entries specified with the logging history size global configuration command), the oldest message entry is deleted from the table to allow the new message entry to be stored.
The history table lists the level keywords and severity level. For SNMP usage, the severity level values increase by 1. For example, emergencies equal 1, not 0, and critical equals 3, not 2.
