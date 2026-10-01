---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000201066348-aab21482-1
title: "hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000201066348-aab21482"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000201066348-aab21482.md
source_anchor: ""
source_lines: [1, 71]
sha256: 973a2b87ed38f722ab3a1a5c0eb0625f2d228c3e0fde3ac23ab385fa2ab2ba78
---

# hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000201066348-aab21482

iMaster NCE-Campus can identify common applications using the built-in application signature database. When predefined applications cannot meet your requirements, you can customize a new application according to its characteristics.
Currently, the following types of customized applications are supported: SA, Domain Name, and Advance Rule. Table 1 lists the methods of identifying customized applications of each type.
| Table 1 Methods of identifying a customized application |  | 
|---|---|
| Customized Application Type | Application Identification Method | 
|---|---|
| SA | Rules can be created through 3-tuple information, keywords, or both 3-tuple information and keywords. The 3-tuple information consists of the server address, protocol type, and port number. The keyword is the feature of the data packet or data flow of the corresponding application. The keyword uniquely identifies the application. If you create a customized application rule based on the 3-tuple information, the rule must contain at least one IP address or port number. | 
| Domain Name | FPI through association with the DNS function: Rules are created by specifying DNS domain names. If users access applications using domain names, you can configure DNS domain name rules to identify these application packets. For devices running V300, this function applies only to UDP packets, DNS packets cannot be truncated (Truncated cannot be set to 1), and a maximum of 10 IP addresses can be resolved in a packet. | 
|  | FPI based on advanced ACL rules: Rules are created based on 5-tuple and DSCP information. The 5-tuple information includes the protocol type, source IP address, destination IP address, source port number, and destination port number of a packet. If you know the 5-tuple and DSCP information of service packets, you can configure rules based on 5-tuple and DSCP information to identify these packets. When you use this method to configure a customized application rule, at least one of the source IP address, destination IP address, source port, destination port, and DSCP information must be configured. | 
| Advance Rule |  | 
For different models of ARs, refer to Table 2 to view the application specifications of the SA type and refer to Table 4 and Table 5 to view application specifications of the FPI types (Domain Name and Advance Rule).
For different models of firewall gateways, refer to Table 3 to view the application specifications of the SA type and refer to Table 6 to view the application specifications of the Advance Rule type.
The following tables describe the maximum specifications supported by the controller for applications of different types, and list device models whose application specifications are lower than the maximum specifications supported by the controller. If application specifications of a device model exceed those supported by the controller, the specifications supported by the controller take effect.
| Table 2 SA application specifications |  |  |  | 
|---|---|---|---|
| Device Model | Maximum Number of Customized Applications | Maximum Number of Rules for a Single Application | Maximum Number of Rules | 
|---|---|---|---|
| Device models are not distinguished. | 256 | 8 | 512 | 
| Table 3 SA specifications (firewall gateway) |  |  |  | 
|---|---|---|---|
| Device Type | Maximum Number of Customized Applications | Maximum Number of Rules for a Single Application | Maximum Number of Rules | 
|---|---|---|---|
| USG6510F-DK, USG6510F-DL, USG6000F-E06-DL, USG6510F-DPL, USG6530F-D, USG6000F-E12-D, USG6530F-DL, USG6560F-D, USG6530F-DPL, USG6510F-D, USG6000F-E06-D | 128 | 8 | 256 | 
| Other models | 256 | 8 | 512 | 
| Table 4 Specifications for applications of the Domain Name type |  |  |  | 
|---|---|---|---|
| Device Model | Maximum Number of Customized Applications | Maximum Number of Rules for a Single Application | Maximum Number of Rules | 
|---|---|---|---|
| AR611, AR611-S, AR611W, AR611W-S, AR611W-LTE4CN, AR611W-LTE6EA | 128 | 128 | 128 | 
| AR617, AR617-LTE4EA, AR617VW, AR617VW-LTE4, AR617VW-LTE4EA |  |  |  | 
| AR631I-LTE4EA, AR631I-LTE4CN |  |  |  | 
| AR651C and AR651F-Lite |  |  |  | 
| Other models | 256 |  | 256 | 
| Table 5 Application specifications of the Advance Rule type (AR) |  |  |  | 
|---|---|---|---|
| Device Model | Maximum Number of Customized Applications | Maximum Number of Rules for a Single Application | Maximum Number of Rules | 
|---|---|---|---|
| AR611, AR611-LTE4EA, AR611-S AR611W, AR611W-S, AR611W-LTE4CN, AR611W-LTE6EA | 128 | 8 | 256 | 
| AR617, AR617-LTE4EA, AR617VW, AR617VW-LTE4, AR617VW-LTE4EA |  |  |  | 
| AR631I-LTE4EA, AR631I-LTE4CN |  |  |  | 
| AR651C and AR651F-Lite |  |  |  | 
| AR6280/AR6300 series | Common application: 256 Extended application: 1792 | Common application: 8 Extended application: 1024 | Common application: 1024 Extended application: 15360 | 
| Other models | 256 | 8 | 1024 | 
| Table 6 Application specifications of the Advance Rule type (firewall gateway) |  |  |  | 
|---|---|---|---|
| Device Type | Maximum Number of Customized Applications | Maximum Number of Rules for a Single Application | Maximum Number of Rules | 
|---|---|---|---|
| USG6510F-DK, USG6510F-DL, USG6000F-E06-DL, USG6510F-DPL, USG6530F-D, USG6000F-E12-D, USG6530F-DL, USG6560F-D, USG6530F-DPL, USG6510F-D, USG6000F-E06-D | 128 | 8 | 256 | 
| USG6520F-K, USG6525F, USG6555F, USG6560F-K, USG6565F, USG6585F, USG6590F-K, USG6585F-B | 256 | 8 | 512 | 
| USG6600F series, USG6700F series | 256 | 8 | 1024 | 
Table 7 lists the specifications of customized applications supported by iMaster NCE-Campus.
| Table 7 Specifications of customized applications |  |  |  | 
|---|---|---|---|
| Application Type | Maximum Number of Customized Applications | Maximum Number of Rules for a Single Application | Maximum Number of Rules | 
|---|---|---|---|
| SA | 256 | 8 | 512 | 
| Domain Name | 256 (total number of applications of the Domain Name and Advance Rule types) | 128 | 256 | 
| Advance Rule (for common applications) |  | 8 | 1024 | 
| Advance Rule (for extended applications) | 1792 | 1024 | 15360 | 
Configuring extended configurations at sites where AR6280 or AR6300 series devices are deployed has the following restrictions:
The AR6700V-L supports this function starting from V600R023C10.
Customized applications of the Domain Name type are supported by devices running V600 since V600R023C00, and supported by AR6700V-L devices since V600R023C10.
Customized applications of the Advance Rule type are supported by devices running V600 since V600R022C00 and are supported by AR6700V-L devices since V600R023C10.
After a customized application is created, iMaster NCE-Campus delivers the customized application rule to its managed devices. If multiple customized applications are created, the corresponding rules are delivered based on their configuration sequence. That is, the rule of a customized application that is configured first is delivered first to devices. If an application packet matches rules of multiple customized applications, the rule of the customized application that is delivered first takes effect. Once devices find the matching application for traffic, they stop matching the traffic with other applications.
For example, application traffic traffic1 matches rules of both the customized applications APP1 and APP2. If the rule of APP1 is delivered to a device earlier than that of APP2, the device preferentially matches traffic1 with APP1 and does not continue to match the traffic with APP2.
| Table 8 Follow-up procedure of customized applications |  |  | 
|---|---|---|
| Function | Operation Scenario and Constraint | Procedure | 
|---|---|---|
| Viewing details about a customized application | You can view detailed information about a customized application. |  | 
| Modifying a customized application | You cannot modify Application Name of a customized application that has been referenced by an application group. |  | 
