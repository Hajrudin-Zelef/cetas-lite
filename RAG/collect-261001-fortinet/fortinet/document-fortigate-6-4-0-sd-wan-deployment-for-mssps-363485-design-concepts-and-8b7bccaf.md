---
id: collect-261001-fortinet/fortinet/document-fortigate-6-4-0-sd-wan-deployment-for-mssps-363485-design-concepts-and-8b7bccaf
title: "document-fortigate-6-4-0-sd-wan-deployment-for-mssps-363485-design-concepts-and--8b7bccaf"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-6-4-0-sd-wan-deployment-for-mssps-363485-design-concepts-and--8b7bccaf.md
source_anchor: ""
source_lines: [1, 25]
sha256: 35a9f3130f3256d7322274133ee92a4b220ec04b6aa24fe1feba2f6c920e3f99
---

# document-fortigate-6-4-0-sd-wan-deployment-for-mssps-363485-design-concepts-and--8b7bccaf

Design concepts and considerations
Design concepts and considerations
The managed Fortinet Secure SD-WAN Solution described in this document consists of fully-functional FortiGate devices (FGT) deployed on every site, and centrally managed by FortiManager (FMG) and FortiAnalyzer (FAZ).
Multiple deployment methods are supported. One option could be the following:
- Manually deploy a new FGT with basic underlay configuration that is necessary to communicate with FMG.
- Onboard the FGT to FMG by using the Device Discovery method, which is initiated from FMG.
- Make the necessary configuration on FMG, and push it to the new FGT.
While this method can be satisfactory in some environments, we recommend following a different approach, which relies on the concept of a Model Device in FMG. The Model Device behaves similar to any managed FGT device, except that there is no real FGT device behind it. Therefore, the necessary steps will be as follows:
- In FMG, create a Model Device for the new FGT that needs to be deployed.
- In FMG, apply all the necessary configuration to the Model Device.
- Link the real FGT device to the Model Device.
The first two steps are done locally on FMG, without any communication with the real FGT device. Hence, they can be done during the preparation stage, when the real FGT device might not even exist yet.
The last step is where the configuration is pushed from FMG to the real device, making it part of the SD-WAN deployment. This step can also be done in several ways:
- Zero-Touch Provisioning (ZTP) using FortiDeploy  lets you link the real FGT device to the Model Device with minimal interaction with the FGT device itself:
			Once the new (unconfigured) FGT device is plugged in, it will call home to obtain the location of the FMG
  - It will then contact FMG, which will authorize it, and map it to the corresponding Model Device by using itss serial number.
  - FMG will then push the Model Device configuration to the FGT device.
  - Once this operation is complete, the FGT is fully deployed and managed by FMG.
- Zero-Touch Provisioning (ZTP) using DHCP Option is identical to the previous method, except for how the location of FMG is obtained. Instead of calling home, the new (unconfigured) FGT device will receive this information from the local DHCP server by using a special DHCP Option 240. But from the FMG perspective, the sequence of events remains the same.
- Low-Touch Provisioning requires a customer to manually configure the FMG details on the new FGT device:
  - Once the new (unconfigured) FGT device is plugged in, the user connects to it, and manually enters details of the FMG.
  - FGT then contacts FMG, and the rest remains identical to the previous methods.
 Here again from the FMG perspective, the sequence of events remains the same.
Detailed instructions for the above methods are outside the scope of this document.
This section contains the following topics:
