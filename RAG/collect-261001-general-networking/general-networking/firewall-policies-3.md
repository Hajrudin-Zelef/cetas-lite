---
id: collect-261001-general-networking/general-networking/firewall-policies-3
title: "firewall-policies"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/firewall-policies.md
source_anchor: ""
source_lines: [221, 342]
sha256: 4f1bcce8070609b85670c1a0348226b317e733a62b545d3233d91f6fbcb73dcb
---

# firewall-policies

A logical best practice that comes from the knowledge of how this process works is to make sure that the more specific or specialized a policy is, the closer to the beginning of the sequence it should be. The more general a policy is the higher the likelihood that it could include in its range of parameters a more specifically targeted policy. The more specific a policy is, the higher the probability that there is a requirement for treating that traffic in a specific way.

**Example**

For security reasons there is no FTP traffic allowed out of a specific subnet so there is a policy that states that any traffic coming from that subnet is denied if the service is FTP, so the following policy was created:

**Policy #1**

| **Source Interface** | Internal1 | 
| **Source Address** | 192.168.1.0/24 | 
| **Source User(s)** | <left at default setting> | 
| **Source Device Type** | <left at default setting> | 
| **Outgoing****Interface** | WAN1 | 
| **Destination Address** | 0.0.0.0/0.0.0.0 | 
| **Service** | FTP | 
| **Schedule** | always | 
| **Action** | deny | 

Now as these things usually go it turns out that there has to be an exception to the rule. There is one very secure computer on the subnet that is allowed to use FTP and once the content has been checked it can them be distributed to the other computer on the subnet. So a second firewall policy is created.

**Policy #2**

| **Source Interface** | Internal1 | 
| **Source Address** | 192.168.1.38/32 | 
| **Source User(s)** | <left at default setting> | 
| **Source Device Type** | <left at default setting> | 
| **Outgoing****Interface** | WAN1 | 
| **Destination Address** | 0.0.0.0/0.0.0.0 | 
| **Service** | FTP | 
| **Schedule** | always | 
| **Action** | Allow | 

By default, a policy that has just been created will be placed last in the sequence so that it is less likely to interfere with existing policies before it can be moved to its intended position. If you look at Policy #2 you will notice that it is essentially the same as Policy #1 exempt for the Source Address and the Action. You will also notice that the Source Address of the Policy #2 is a subset of the Source address in policy #1. This means that if nothing further is done, Policy #2 will never see any traffic because the traffic will always be matched by Policy #1 and processed before it has a chance to reach the second policy in the sequence. For both policies to work as intended Policy #2 needs to be moved to before Policy #1 in the sequence.

### Policy identification

There are two ways to identify a policy. The most obvious is the policy name and this is easily read by humans, but with a little effort it is possible to have a policy without a name, therefore every policy has an ID number.

When looking at the policy listing it can appear as if the policies are identified by the sequence number in the far left column. The problem is that this number changes as the position of the policy in the sequence changes. The column that correctly identifies the policy, and the value sticks with the policy is the “ID” column. This column is not shown by default in the listing but can be added to the displayed columns by right clicking on the column heading bar and selecting it from the list of possible columns.

When looking in the configuration file the sequence is based upon the order of the policies as they are in the file just as they are in the list in the GUI. However, if you need to edit the policy in the CLI you must use the ID number.

### UUID support

Universally Unique Identifier (UUID) attributes have been added to policies to improve functionality when working with FortiManager or FortiAnalyzer units. If required, the UUID can be set manually through the CLI.

CLI Syntax:

config firewall {policy/policy6/policy46/policy64} edit 1 set uuid <example uuid: 8289ef80-f879-51e2-20dd-fa62c5c51f44> next

end

### NTurbo support CAPWAP traffic

NTurbo is used for IPSec+IPS case. The IPSec SA info is passed to NTurbo as part of VTAG for control packet and will be used for the xmit.

If the packets need to go through IPSec interface, the traffic will be always offloaded to NTurbo. But for the case that SA has not been installed to NP6 because of hardware limitation or SA offload disable, the packets will be sent out through raw socket by IPS instead of NTurbo, since the software encryption is needed in this case.

**CLI :**

Previously, NTurbo could only be enabled or disabled globally. The setting of np-acceleration has been added to the firewall policy context instead of just the global context.

CLI command in the firewall policy to enable/disable NTurbo acceleration.

config firewall policy edit 1 set np-accelation [enable|disable] end

When IPS is enabled for VPN IPsec traffic, the data can be accelerated by NTurbo.

### Learning mode for policies

The learning mode feature is a quick and easy method for setting a policy to allow everything but to log it all so that it can later be used to determine what restrictions and protections should be applied. The objective is to monitor the traffic not act upon it while in Learning mode.

Once the **Learn** action is enabled, functions produce hard coded profiles that will be enabled on the policy. The following profiles are set up:

- AntiVirus (av-profile) l Web Filter ( webfilter-profile) l Anti Spam( spamfilter-profile ) l Data Leak Prevention (dlp-sensor ) l Intrusion Prevention (ips-sensor ) l Application Control (application-list ) l Proxy Options (profile-protocol-options)

Profiles that are not being used are:

- DNS Filter (Does not have a Flow mode) l Web Application Firewall(Does not have a Flow mode) l CASI(Almost all signatures in CASI require SSL deep inspection. Without SSL inspection, turning on CASI serves little purpose)

The ability to allow policies to be set to a learning mode is enabled on a per VDOM basis.

config system settings set gui-policy-learning [enable | disable] end

Once the feature is enabled on the VDOM, Learn is an available **Action** option when editing a policy.

Once the Learning policy has been running for a sufficient time to collect needed information a report can be looked at by going to **Log & Report > Learning Report**.

The Report can be either a **Full Report** or a **Report Summary** The time frame of the report can be **5 minutes**, **1 hour**, or **24 hours**.

The Learning Report includes: **Deployment Methodology** l Test Details l Start time l End time l Model


Policy modes

- Firmware
- Policy List

**Executive Summary** l Total Attacks Detected l Top Application Category l Top Web Category l Top Web Domain l Top Host by Bandwidth l Host with Highest Session Count **Security and Threat Prevention** l High Risk Applications l Application Vulnerability Exploits l Malware, botnets and Spyware/Adware l At-Risk Devices and Hosts **User Productivity** l Application Usage l Top Application Categories l Top Social Media Applications l Top Video/Audio Streaming Applications l Top Peer to Peer Applications l Top Gaming Applications

- Web Usage l Top Web Categories l Top Web Applications l Top Web Domains

## Policy modes

You can operate your FortiGate or individual VDOMs in **Next Generation Firewall (NGFW) Policy Mode**.

You can enable NGFW policy mode by going to **System > Settings**, setting the **Inspection mode** to **Flowbased** and setting the NGFW mode to **Policy-based**. When selecting **NGFW policy-based** mode you also select the SSL/SSH Inspection mode that is applied to all policies

**Flow-based** inspection with profile-based **NGFW mode** is the default in FortiOS 5.6.

Or use the following CLI command: config system settings

Policy modes

set inspection-mode flow

set ngfw-mode {profile-based | policy-based}

end

### NGFW policy mode and NAT

