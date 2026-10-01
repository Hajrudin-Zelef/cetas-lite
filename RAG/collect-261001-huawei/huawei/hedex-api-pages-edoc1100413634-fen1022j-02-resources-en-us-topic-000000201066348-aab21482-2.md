---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000201066348-aab21482-2
title: "hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000201066348-aab21482"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000201066348-aab21482.md
source_anchor: ""
source_lines: [72, 95]
sha256: 294e6fca33c7dedeb3b9ff2ff91fc42b2f44abe3a367f16a5fada21a99f60f5c
---

# hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000201066348-aab21482

| Cloning a customized application | You can quickly create a customized application by simply modifying an existing customized application. |  | 
| Deleting a customized application | You cannot delete a customized application that has been referenced by an application group. |  | 
| Table 9 Parameters on the Customized Application page |  |  |  |  |  | 
|---|---|---|---|---|---|
| Parameter |  |  |  | Description | Data Plan Required or Not | 
|---|---|---|---|---|---|
| Name |  |  |  | Name of a customized application. The value can contain only digits, letters, and underscores (_). The name of a customized application of the SA type can contain 1 to 29 characters, and that of a customized application of a non-SA type can contain 1 to 28 characters. | Y | 
| Description |  |  |  | Description of a customized application's function. | - | 
| Application Group |  |  |  | Application group to which the application is to be added. It must be an existing application group. You can also add an application to an application group after customizing it. | Y | 
| Application scope |  |  |  | The value can be Common or Extended. Extended is applicable to the scenario where the specifications of applications of the Advance Rule type need to be expanded. Currently, only AR6280 and AR6300 series devices support this function. | Y | 
| Application Type |  |  |  | You can select an application type to define a rule for identifying application packets. The application types include SA, Domain Name, and Advance Rule. The configurable rule parameters vary according to the application type. If Application scope is set to Extended, you can only create customized applications of the Advance Rule type. | Y | 
| Rule | Name |  |  | Name of a rule defined for identifying application packets. A rule contains the protocol number and port number used by an application, and other basic attributes. | Y | 
|  | Description |  |  | Description of a rule. | - | 
|  | When Application Type is set to SA | IP/Port |  | IP address and port number of application packets matching the rule. The system does not distinguish source and destination IP addresses of SA application packets, as well as source and destination port numbers. If an IP address and a port number are specified together in a rule, the relationship between IP and Port is And, that is, application packets only with the specified IP address and port number match this rule. For details about the value ranges of the IP address and port number, see the prompt message on the web UI.  NOTE:  When configuring a rule for V600 devices, you need to set this parameter to the destination IP address and destination port number of application packets. Otherwise, application packets cannot be matched. | Y | 
|  |  | Protocol |  | Transport layer protocol of a customized application rule. The options include All, TCP, and UDP. | Y | 
|  |  | Signature | Signature | Signature information. Data packets of some applications contain the same character string, which is regarded as a signature. | - | 
|  |  |  | Context | You can select the packet- or flow-based mode for signature identification: In the packet-based mode, the system checks every packet of applications. In flow-based mode, the system only checks the first packet in the application data flow and does not check the subsequent packets if the system detects that the subsequent packets belong to the same data flow based on the 5-tuple information. | - | 
|  |  |  | Direction | Direction of packets to be identified. You can configure a rule to identify the signatures only in request or response packets, or in both of them. | - | 
|  |  |  | Plain-text String | Signature string, which is case-sensitive. | - | 
|  | When Application Type is set to Domain Name | Domain Name |  | Domain name of a customized application. | Y | 
|  | When Application Type is set to Advance Rule | Source IP/Source Port |  | Source IP address and port number of application packets matching the rule. If the source IP address and source port number are specified together in a rule, the relationship between Source IP and Source Port is And, that is, application packets only with the specified source IP address and source port number match this rule. For details about the value ranges of the IP address and port number, see the prompt message on the web UI. | Y | 
|  |  | Destination IP/Destination Port |  | Destination IP address and destination port number of application packets matching the rule. If the destination IP address and destination port number are specified together in a rule, the relationship between Destination IP and Destination Port is And, that is, application packets only with the specified destination IP address and destination port number match this rule. In most cases, the IP address of an application server is a fixed public IP address. This allows the system to identify application packets based on the destination IP address. | Y | 
|  |  | DSCP |  | You can set the packet priority by specifying a DSCP value. The options include Default (which is 0), CS (low priority), AF (medium priority), EF (high priority), and User Defined. V600 devices do not support this function. | Y | 
|  |  | Protocol |  | Transport layer protocol of a customized application rule. The options include All, TCP, and UDP. | Y |
