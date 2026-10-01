---
id: collect-261001-cisco/cisco/enterprise-en-ethernet-services-e-line-and-e-lan-thread-481187-875-104c1d5a
title: "enterprise-en-ethernet-services-e-line-and-e-lan-thread-481187-875-104c1d5a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/enterprise-en-ethernet-services-e-line-and-e-lan-thread-481187-875-104c1d5a.md
source_anchor: ""
source_lines: [1, 14]
sha256: d7c62f6d126d02e9770dad9219ccdd1edeacb8552b43ccf3670c718925726718
---

# enterprise-en-ethernet-services-e-line-and-e-lan-thread-481187-875-104c1d5a

Hi there!
This time, I will share with you about E-Line and E-LAN.
The OptiX OSN 550 supports point-to-point E-Line services and multipoint-to-multipoint E-LAN services.
Standardization organizations such as ITU-T, IETF, and MEF stipulate the model frames for L2 Ethernet services. Table 1 lists these model frames. In this document, the L2 Ethernet services are of the model frame stipulated by MEF.
lists the E-Line and E-LAN services supported by the OptiX OSN 550.
Figure 1 illustrates the E-Line service provided by the OptiX OSN equipment.
Company A has two branches in City 1 and City 3. Company B has two branches in City 2 and City 3. Company C has two branches in City 1 and City 2. The branches of Companies A, B, and C require data communication. The OptiX OSN equipment can separately provide a private line service for Companies A, B, and C to meet the communication requirement. In addition, the service data is completely isolated.
Figure 1 E-Line service 
Figure 2 illustrates the E-LAN service provided by the OptiX OSN equipment.
Company Z is headquartered in City 3. Branch A of the company is located in City 1 and City 2, and Branch B of the company is located in City 1, City 2, and City 3. Branch A and Branch B do not communicate with each other, and the data of them should be separated from each other. The headquarters, however, need to communicate with all the branches and need to access the Internet.
The OptiX OSN equipment can be used to provide the E-LAN service. Different VLAN tags are used to identify service data from different branches. In this manner, the headquarters can communicate with the branches and the data from different branches is isolated. In addition, the VLAN is used to isolate the Internet data accessed by the headquarters from the internal service data.
 
That's all, I welcome everyone to leave a message and exchange in the comment area!
Thank you！
