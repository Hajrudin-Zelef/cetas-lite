---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-04-resources-dc-display-controller-cpos-0133c52e
title: "Display the physical layer configuration of CPOS1/0/0 and all E1/T1 channels."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-04-resources-dc-display-controller-cpos--0133c52e.md
source_anchor: ""
source_lines: [1, 43]
sha256: c45ef10a70444c5e76fd0331b37ed259cf1678a5cedf069e75c1f4162ef23714
---

# Display the physical layer configuration of CPOS1/0/0 and all E1/T1 channels.

The display controller cpos command displays the physical layer configuration of the CPOS interface and all E1/T1 channels.
| Parameter | Description | Value | 
|---|---|---|
| cpos-number | Displays the physical layer configuration of the CPOS interface with a specified number and all E1/T1 channels. The value is in the format of slot ID/subcard ID/interface sequence number. If the cpos-number parameter is not specified, physical layer configurations of all CPOS interfaces are displayed. | - | 
When monitoring the interface status or identifying the cause of the interface faults, you can run the display controller cpos command to obtain the status and statistics of the interface. According to the information, you can measure the traffic and identify the fault.
# Display the physical layer configuration of CPOS1/0/0 and all E1/T1 channels.
<Huawei> display controller cpos 1/0/0
Cpos1/0/0 current state : UP
Description : HUAWEI, AR Series, Cpos1/0/0 Interface
  Frame-format SDH, multiplex AU-3, clock master, loopback not set
  Tx: J0: 0x1, J1: "NetEngine", C2: 0x2
  Rx: J0: 0x1, J1: "NetEngine", C2: 0x2
Regenerator section:
  Alarm: none
  Error: 0 BIP
Multiplex section:
  Error: 0 BIP, 0 REI
Higher order path (VC-3-1):
Higher order path (VC-3-2):
Higher order path (VC-3-3):
Cpos1/0/0  CT1 1  is up
  Frame-format ESF, clock master, loopback not set
Cpos1/0/0  CT1 2  is up
Cpos1/0/0  CT1 3  is up
 (here omitted some output information)
Cpos1/0/0  CT1 83  is up
Cpos1/0/0  CT1 84  is up
| Table 1 Description of the display controller cpos command output |  | 
|---|---|
| Item | Description | 
|---|---|
| Cpos1/0/0 current state: | Current physical status of the CPOS interface. | 
| Description | Description of the CPOS interface. The value is a maximum of 242 case-sensitive characters with spaces. The description helps users learn about the functions of the CPOS interface. | 
| Frame-format SDH, multiplex AU-3, clock master, loopback not set | Physical layer information of the CPOS interface:  | 
| Tx: J0: 0x1, J1: "NetEngine", C2: 0x2 | Sent overhead bytes. To set the overhead byte, run the flag command. | 
| Rx: J0: 0x1, J1: "NetEngine", C2: 0x2 | Received overhead bytes. | 
| Regenerator section | Alarm and error statistics of the regenerator section. | 
| Multiplex section | Alarm and error statistics of the multiplexing section. | 
| Higher order path (VC-3-x) | Alarm and error statistics of the higher-order path. VC-3-x indicates VC-3 numbered x. When AU-3 multiplexing path is adopted, three VC-3s are multiplexed into one STM-1. Therefore, there are three higher-order paths. When AU-4 multiplexing path is adopted, there is only one higher-order path VC-4. | 
| Alarm | Alarm statistics.  | 
| Error | Error statistics. Possible error types are as follows:  | 
| Cpos1/0/0 CT1 1 is up | Status of T1 channel 1 of CPOS 1/0/0. | 
| Frame-format ESF, clock master, loopback not set | Physical layer information of the T1 channel:  |
