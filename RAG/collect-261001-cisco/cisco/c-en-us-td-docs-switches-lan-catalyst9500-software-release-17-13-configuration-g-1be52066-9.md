---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066-9
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066.md
source_anchor: ""
source_lines: [951, 1089]
sha256: 481cecba7017f96f90e1e15ea3a1194b2ef98579e09d42f320ea4ab1507d6582
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066

*Sep 29 09:38:20.256: %LINEPROTO-5-UPDOWN: Line protocol on Interface FiftyGigE1/6/0/41, changed state to up
Device#
Example: Displaying StackWise Virtual Link Information
Sample output of show stackwise-virtual link command
In this example, the output is displayed from a switch where SVL is configured using network modules.
Device# show stackwise-virtual link
Stackwise Virtual Link(SVL) Information:
----------------------------------------
Flags:
------
Link Status
-----------
U-Up D-Down
Protocol Status
---------------
S-Suspended P-Pending E-Error T-Timeout R-Ready
-----------------------------------------------
Switch SVL Ports Link-Status Protocol-Status
------ --- ----- ----------- ---------------
1 1 TenGigabitEthernet1/1/1 U R
2 1 TenGigabitEthernet2/1/1 U R
The following is a sample output from the C9500X-28C8D model of Cisco Catalyst 9500X Series Switches:
Device# show stackwise-virtual link
Stackwise Virtual Link(SVL) Information:
----------------------------------------
Flags:
------
Link Status
-----------
U-Up D-Down
Protocol Status
---------------
s-Suspended P-Bundled E-Error D-Down R-RLayer3 I-Indiv
------------------------------------------------------
Switch SVL Ports Link-Status Protocol-Status
------ --- ----- ----------- ---------------
1 1 HundredGigE1/0/7 U P
2 1 HundredGigE2/0/7 U P
By default in standalone mode, the switches are identified as Switch 1 unless explicitly changed to some other switch number.
During the conversion to StackWise Virtual, the switch numbers are changed automatically to reflect two switches in a StackWise
Virtual domain.
Example: Displaying StackWise Virtual Dual-Active-Detection Link Information
Sample output of show stackwise-virtual dual-active-detection command
StackWise Virtual DAD links configuration:
Device# show stackwise-virtual dual-active-detection
Recovery Reload for switch 1: Enabled
Recovery Reload for switch 2: Enabled
Dual-Active-Detection Configuration:
-------------------------------------
Switch Dad port Status
------ ------------ ---------
1 FortyGigabitEthernet1/0/3 up
2 FortyGigabitEthernet2/0/3 up
StackWise Virtual DAD links configuration after configuring the dual-active recovery-reload-disable command:
Device# show stackwise-virtual dual-active-detection
Recovery Reload for switch 1: Enabled
Recovery Reload for switch 2: Enabled
Dual-Active-Detection Configuration:
-------------------------------------
Switch Dad port Status
------ ------------ ---------
1 FortyGigabitEthernet1/0/3 up
2 FortyGigabitEthernet2/0/3 up
Sample output of show stackwise-virtual dual-active-detection epagp command
StackWise Virtual DAD ePAgP information:
Device# show stackwise-virtual dual-active-detection pagp
Pagp dual-active detection enabled: Yes
In dual-active recovery mode: No
Recovery Reload for switch 1: Enabled
Recovery Reload for switch 2: Enabled
Channel group 11
Dual-Active Partner Partner Partner
Port Detect Capable Name Port Version
Fo1/0/17 Yes SwitchA Hu2/0/1 1.1
Fo2/0/21 Yes SwitchA Hu1/0/4 1.1
Partner Name and Partner Port fields in the output represent the name and the ports of the peer switch to which the PagP port-channel is connected through
MEC.
Verifying Cisco StackWise Virtual Configuration
To verify your StackWise Virtual configuration, use the following show commands:
Table 3. show Commands to Verify Cisco StackWise Virtual Configuration
show stackwise-virtual switch number <1-2>
Displays information of a particular switch in the stack.
show stackwise-virtual link
Displays StackWise Virtual link information.
show secure-stackwise-virtual authorization-key
Displays the installed Secure StackWise Virtual authorization key.
show secure-stackwise-virtual status
Displays the Secure StackWise Virtual status.
show secure-stackwise-virtual interface
Displays the Secure StackWise Virtual interface statistics.
show stackwise-virtual bandwidth
Displays the bandwidth available for the Cisco StackWise Virtual.
(Optional)Assigns a new switch number. The default number is 1.
Additional References for StackWise Virtual
Table 4. Related Documents
Related Topic
Document Title
For complete syntax and usage information for the commands used in this chapter.
High Availability Command Reference for Catalyst 9500 Switches
Feature History for Cisco StackWise Virtual
This table provides release and related information for features explained in this module.
These features are available on all releases subsequent to the one they were introduced in, unless noted otherwise.
Release
Feature
Feature Information
Cisco IOS XE Everest 16.6.1
Cisco StackWise Virtual
Cisco StackWise Virtual is a network system virtualization technology that pairs two switches into one virtual switch to simplify
operational efficiency with a single control and management plane.
Support for this feature was introduced only on the C9500-12Q, C9500-16X, C9500-24Q, C9500-40X models of the Cisco Catalyst
9500 Series Switches.
Cisco IOS XE Gibraltar 16.10.1
Cisco StackWise Virtual
Support for this feature was introduced on the C9500-32C, C9500-32QC, C9500-48Y4C, and C9500-24Y4C models of the Cisco Catalyst
9500 Series Switches.
Cisco IOS XE Gibraltar 16.11.1
Recovery Reload
Support for disabling DAD recovery reload was introduced. Enter the dual-active recovery-reload-disable command in stackwise virtual mode (config-stackwise-virtual).
Support was introduced on all models of the Cisco Catalyst 9500 Series Switches.
Cisco IOS XE Gibraltar 16.12.1
Secure StackWise Virtual
Secure StackWise Virtual support was introduced for two node front-side stacking. Secure StackWise Virtual is FIPS 140-2 compliant
and encrypts control packets as well.
Support was introduced on all models of the Cisco Catalyst 9500 Series Switches.
BGP EVPN VXLAN on switches with Cisco StackWise Virtual
Support for the BGP EVPN VXLAN feature was introduced on switches with Cisco StackWise Virtual configured.
Support was introduced on all models of the Cisco Catalyst 9500 Series Switches.
Cisco IOS XE Amsterdam 17.2.1
BUM Traffic Optimization on switches with Cisco StackWise Virtual
Support for the BUM Traffic Optimization feature was introduced on switches with Cisco StackWise Virtual configured.
Support was introduced on all models of the Cisco Catalyst 9500 Series Switches.
Cisco IOS XE Dublin 17.10.1
Cisco StackWise Virtual
Support for this feature was introduced on the C9500X-28C8D model of Cisco Catalyst 9500 Series Switches.
Cisco IOS XE Dublin 17.10.1b
Cisco StackWise Virtual
Support for this feature was introduced on the C9500X-60L4D model of Cisco Catalyst 9500 Series switches.
Use the Cisco Feature Navigator to find information about platform and software image support. To access Cisco Feature Navigator,
go to Cisco Feature Navigator.
