---
id: collect-261001-general-networking/general-networking/2018-12-firewall-policies-3-4c1b7ce8-4
title: "Firewall policies"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/2018-12-firewall-policies-3-4c1b7ce8.md
source_anchor: ""
source_lines: [105, 137]
sha256: a89325c241cc7c838a43545cf35718d656d2a9c7c841ebadb0a1916428cded70
---

# Firewall policies

When looking at the policy listing it can appear as if the policies are identified by the sequence number in the far left column. The problem is that this number changes as the position of the policy in the sequence changes. The column that correctly identifies the policy, and the value sticks with the policy is the “ID” column. This column is not shown by default in the listing but can be added to the displayed columns by right clicking on the column heading bar and selecting it from the list of possible columns.
When looking in the configuration file the sequence is based upon the order of the policies as they are in the file just as they are in the list in the GUI. However, if you need to edit the policy in the CLI you must use the ID number.
Universally Unique Identifier (UUID) attributes have been added to policies to improve functionality when working with FortiManager or FortiAnalyzer units. If required, the UUID can be set manually through the CLI.
CLI Syntax:
config firewall {policy/policy6/policy46/policy64} edit 1 set uuid <example uuid: 8289ef80-f879-51e2-20dd-fa62c5c51f44> next
NTurbo is used for IPSec+IPS case. The IPSec SA info is passed to NTurbo as part of VTAG for control packet and will be used for the xmit.
If the packets need to go through IPSec interface, the traffic will be always offloaded to NTurbo. But for the case that SA has not been installed to NP6 because of hardware limitation or SA offload disable, the packets will be sent out through raw socket by IPS instead of NTurbo, since the software encryption is needed in this case.
CLI :
Previously, NTurbo could only be enabled or disabled globally. The setting of np-acceleration has been added to the firewall policy context instead of just the global context.
CLI command in the firewall policy to enable/disable NTurbo acceleration.
config firewall policy edit 1 set np-accelation [enable|disable] end
When IPS is enabled for VPN IPsec traffic, the data can be accelerated by NTurbo.
The learning mode feature is a quick and easy method for setting a policy to allow everything but to log it all so that it can later be used to determine what restrictions and protections should be applied. The objective is to monitor the traffic not act upon it while in Learning mode.
Once the Learn action is enabled, functions produce hard coded profiles that will be enabled on the policy. The following profiles are set up:
Profiles that are not being used are:
The ability to allow policies to be set to a learning mode is enabled on a per VDOM basis.
config system settings set gui-policy-learning [enable | disable] end
Once the feature is enabled on the VDOM, Learn is an available Action option when editing a policy.
Once the Learning policy has been running for a sufficient time to collect needed information a report can be looked at by going to Log & Report > Learning Report.
The Report can be either a Full Report or a Report Summary The time frame of the report can be 5 minutes, 1 hour, or 24 hours.
The Learning Report includes: Deployment Methodology l Test Details l Start time l End time l Model
Policy modes
Executive Summary l Total Attacks Detected l Top Application Category l Top Web Category l Top Web Domain l Top Host by Bandwidth l Host with Highest Session Count Security and Threat Prevention l High Risk Applications l Application Vulnerability Exploits l Malware, botnets and Spyware/Adware l At-Risk Devices and Hosts User Productivity l Application Usage l Top Application Categories l Top Social Media Applications l Top Video/Audio Streaming Applications l Top Peer to Peer Applications l Top Gaming Applications
You can operate your FortiGate or individual VDOMs in Next Generation Firewall (NGFW) Policy Mode.
You can enable NGFW policy mode by going to System > Settings, setting the Inspection mode to Flowbased and setting the NGFW mode to Policy-based. When selecting NGFW policy-based mode you also select the SSL/SSH Inspection mode that is applied to all policies
Flow-based inspection with profile-based NGFW mode is the default in FortiOS 5.6.
Or use the following CLI command: config system settings
set inspection-mode flow
set ngfw-mode {profile-based | policy-based}
If your FortiGate is operating in NAT mode, rather than enabling source NAT in individual NGFW policies you go to Policy & Objects > Central SNAT and add source NAT policies that apply to all matching traffic. In many cases you may only need one SNAT policy for each interface pair. For example, if you allow users on the internal network (connected to port1) to browse the Internet (connected to port2) you can add a port1 to port2 Central SNAT policy similar to the following:
You configure Application Control simply by adding individual applications to security policies. You can set the action to accept or deny to allow or block the applications.
You configure Web Filter by adding URL categories to security policies. You can set the action to accept or deny to allow or block the applications.
You can also combine both application control and web filtering in the same NGFW policy mode policy. Also if the policy accepts applications or URL categories you can also apply Antivirus, DNS Filtering, and IPS profiles in NGFW mode policies as well a logging and policy learning mode.
