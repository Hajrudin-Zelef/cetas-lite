---
id: collect-261001-cisco/cisco/enterprise-tr-doc-edoc1000060766-e9dd7d04-how-do-i-troubleshoot-a-stack-setup-fa-7bf39302-1
title: "enterprise-tr-doc-edoc1000060766-e9dd7d04-how-do-i-troubleshoot-a-stack-setup-fa-7bf39302"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2012-11-23"]
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-tr-doc-edoc1000060766-e9dd7d04-how-do-i-troubleshoot-a-stack-setup-fa-7bf39302.md
source_anchor: ""
source_lines: [1, 81]
sha256: e2733b7816f8e27c4a6dbe487d80cf92bf511b37591691c4601abd2e0427c887
---

# enterprise-tr-doc-edoc1000060766-e9dd7d04-how-do-i-troubleshoot-a-stack-setup-fa-7bf39302

Enterprise
Check whether a stack can be set up between these member switches.
Run the display device command to view the member switch models and then determine whether a stack can be set up between these switches. If not, replace these switches.
You can refer to the product documentation to check whether these switches can set up a stack.
For example, if the switch model is CE5850-48T4S2Q-EI, you can confirm that this switch model can set up a stack with CE5850-48T4S2Q-EI or CE5850-24T4S2Q-EI.
<HUAWEI> display device
Device status:
-------------------------------------------------------------------------------------------
Slot  Card   Type                     Online   Power Register     Alarm     Primary
-------------------------------------------------------------------------------------------
1     -      CE5850-48T4S2Q-EI        Present  On    Registered   Normal    Master
      FAN2   FAN-40SA-B               Present  On    Registered   Normal    NA
      PWR2   PAC-150WA                Present  On    Registered   Normal    NA
-------------------------------------------------------------------------------------------
Check whether the stack configuration is correct.
Run the display stack configuration all command to check whether the stack configuration meets requirements.
The following uses the command output on a CE12800 switch as an example.
<HUAWEI> display stack configuration all
Oper : Operation
Conf : Configuration
*    : Offline configuration
Isolated Port: The port is in stack mode, but does not belong to any Stack-Port
Attribute Configuration:
---------------------------------------------------------------
 MemberID      Domain         Priority       Mode     Enable
Oper(Conf)   Oper(Conf)      Oper(Conf)   Oper(Conf)   Oper
---------------------------------------------------------------
1(1)         10(10)          150(150)     MB(MB)       Enable
---------------------------------------------------------------
Stack-Port Configuration:
--------------------------------------------------------------------------------
Stack-Port           Member Ports
--------------------------------------------------------------------------------
Stack-Port1/1        10GE1/1/0/1    10GE1/1/0/2
--------------------------------------------------------------------------------
Check whether all member switches use the same stack domain ID. All member switches must use the same stack domain ID. Otherwise, they cannot set up a stack.
If member switches use different stack domain IDs, run the stack member { member-id | all } domain domain-id command to change their stack domain IDs to the same.
(Applicable only to the CE12800 and CE12800E) Check whether all member switches use the same stack connection mode. All member switches must use the same stack connection mode. Otherwise, they cannot set up a stack. MB indicates the default MPU connection mode, and LC indicates the LPU connection mode.
If member switches use different stack connection modes, run the stack member { member-id | all } link-type { mainboard-direct | linecard-direct } command to change their stack connection modes to the same.
(Applicable only to the CE12800, CE12800E, and CE16800) Check whether the stacking function has been enabled on member switches. Member switches must have the stacking function enabled to set up a stack. Enable indicates that the stacking function is enabled, and Disable indicates that the stacking function is disabled.
If member switches have the stacking function disabled, run the stack enable command to enable this function.
Check whether member switches have offline configuration. The configuration marked with an asterisk (*) is offline configuration. If offline configuration exists, you must delete it. This is because offline configuration may case a stack configuration conflict, which will cause a failure to set up a stack.
Check whether stack connections are correct.
Check whether stack connections are consistent with the plan and configuration.
For details about how to connect switch members, see the corresponding product documentation.
CE12800 series switches support two stack connection modes: MPU connection and LPU connection. In MPU connection mode, both SIP ports on MPUs and service ports on LPUs need to be connected. In LPU connection mode, only service ports on LPUs need to be connected. The MPU connection mode is recommended because it separates management and forwarding links and ensures stack reliability.
Each CE8800, CE7800, CE6800, or CE5800 switch has two logical stack ports: Stack-Portn/1 and Stack-Portn/2, in which n indicates the stack member ID of a switch. One logical stack port can contain multiple physical member ports.
Stack member switches can be connected in a ring or chain topology. Logical stack ports on the member switches can be connected in any sequence.
Check whether ports used for stack connections are Up.
Run the display interface brief command to check whether the ports used for stack connections are physically Up, including stack member ports and SIP ports (applicable only to the CE12800, CE12800E, and CE16800). If these ports are physically Down, check whether optical modules and fibers are faulty.
<HUAWEI> display interface brief
......
Interface                  PHY      Protocol  InUti OutUti   inErrors  outErrors
...... 
10GE1/4/0/44               down     down         0%     0%          0          0 
10GE1/4/0/45               down     down         0%     0%          0          0 
MEth0/0/0/0                up       down      0.01%     0%          0          0 
NULL0                      up       up(s)        0%     0%          0          0 
Sip1/5/0/0                 down     down         0%     0%          0          0 
Sip1/5/0/1                 down     down         0%     0%          0          0 
Stack-Port1/1              down     down         0%     0%          0          0 
  10GE1/4/0/46             down     down         0%     0%          0          0 
  10GE1/4/0/47             down     down         0%     0%          0          
To set up a stack of CE12800 switches, ensure that SIP ports on MPUs are connected using network cables or using GE optical modules and LC fibers. Do not connect these SIP ports using 10GE optical modules.
You can also run the display stack link-state last-down-reason command to check the reason why the stack link protocol becomes Down. A possible cause is incorrect configuration or cable connection.
Check stack failure event information.
Run the display stack troubleshooting command to check whether stack failure events occur. This command can record some failures occurring during a stack setup, including configuration errors and connection errors. Troubleshoot these failures according to failure event description.
In V200R019C00 and earlier versions, run the display stack troubleshooting command to check faults in a stack.
In V200R005C20, V200R019C10, and later versions, run the display stack troubleshooting current command to check faults in a stack.
<HUAWEI> display stack troubleshooting current
Total :1      
----------------------------------------------------------------------------------------
Seq  Time                     Event Description
----------------------------------------------------------------------------------------
1    2012-11-23 19:28:23.889  The devices belong to different stack domains, 
                                    and stack cannot be established. (MemberID = 1, 
                                    DomainID = 10, PeerMemberID = 2, PeerDomainID = 20) 
----------------------------------------------------------------------------------------
| Fault | Description | Troubleshooting Procedure | 
|---|---|---|
| The devices belong to different stack domains. | The stack domain IDs of member switches are different. |  | 
