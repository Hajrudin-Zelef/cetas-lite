---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-6
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["EU", "Huawei"]
dates: []
keywords: ["copyright", "energy"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [631, 781]
sha256: 748ad7622076b69e22b655f17d0fd2c945ab582256f030eeb963e09bbd079969
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

The architecture of the Intelligence Aware Engine (IAE) on the USG6000 
series is different from that of traditional threat detection engines. The attack 
detection engine of a conventional firewall matches each packet with the attack 
signature database. Attacks can easily evade such detection. The IAE 
reassembles packets based on sessions, parses protocols, and matches 
signatures for a more accurate detection of protocol-specific attacks. During 
attack detection, the IAE parses each packet only once over the multi-core CPU 
architecture and can perform multiple security inspection tasks at the same time. 
The hardware acceleration module identifies applications and matches 
signatures at a high speed. If all signatures for an attack are met, the IAE takes 
an appropriate action according to the configured policy. If the signatures are 
not met, the IAE automatically adjusts the tracing status to ensure the 
high-speed forwarding of secure traffic. This architecture ensures the minimum 
compromise of the overall performance with multiple security services 
enabled. 
The IAE uses a multi-core hardware platform for concurrent service processing. 
In addition, the IAE uses the hardware acceleration technology for application 
identification and signature matching, greatly improving attack detection 
efficiency. 
 Hardware co-processor acceleration 
Huawei NG_Security hardware platform has integrated the IPSec and SSL 
encryption and decryption, compression and decompression, pattern matching, 
and hard disk RAID hardware co-processors. These co-processors processes 
the services that may degrade CPU performance, such as encryption and 
decryption, compression and decompression, and pattern matching to reduce 
the consumption of CPU resources. 
Figure 3-3 Huawei NG_Security co-processor and CPU expansion capabilities 
+
Hardware co-processor acceleration Expandable processing capability
· The encryption and decryption computing, pattern matching, and compression and decompression 
consume over 50% of CPU resources, degrading the performance rapidly.
· Hardware co-processors provide such computing to protect CPU resources.
· The scalable capacity design doubles the hardware processing capability.
 
 High-speed Switch Fabric

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  18 
   
 
Huawei NG_Security hardware platform uses a 480 Gbit/s switching chip for 
the communications among the multi-core CPU, service processing module, 
and interface expansion module. Its high-speed switching bus provides 
sufficient bandwidths for all modules and ensures the smooth switching. 
Figure 3-4 Huawei NG_Security software platform 
CPU
CPU
CPU
CPU
CPU
。。。
CPU
CPU
CPU
CPU
CPU
。。。
20G
20G
20G
。
。
20G
20G
20G
。
。
Switch Fabric
 
 
 Storage module 
Huawei NG_Security hardware platform supports the 300 GB high-speed SAS 
hard disk to store real-time logs and reports. 
Two hard disks work in RAID1 mode to back up user data. 
The hot swap design of hard disks enables capacity expansion and upgrade. 
 Scalability 
1. The USG6000 series uses the flexible and scalable architecture to double 
security performance by adding more SPUs for different application 
scenarios. The combination of the intelligent awareness engine and the 
elastic hardware structure enable the USG6000 series to deliver 
10-Gigabit level threat prevention performance, meeting the security 
protection requirements of large enterprise data centers.  
2. The NGFW supports multiple slots for high-density expansion interface 
cards and diversified interface cards that provide the GE electrical and 
optical ports and 10GE ports. You can flexibly improve hardware 
forwarding capabilities and device performance according to actual 
conditions. 
3. Based on the virtual system function of the USG6000 series, you can 
divide a physical device into multiple virtual devices that are independent 
and locally isolated to implement system-level expansion and meet the 
requirements of device leasing and cloud computing. 
4. Hard disks are optional. You can choose hard disks as required.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  19 
   
 
The previous scalability features enable customers to configure only required 
modules in the initial phase and expand capacities when necessary to maximize 
customer investments. 
 High reliability 
1. Power supplies provide 1+1 redundancy, and hard disks work in RAID1 
mode. When one power supply or hard disk is faulty, the other one takes 
over all the services. The hardware design ensures service continuity. 
2. Fault detection: The system monitors the working statuses of the 
integrated device and key components on SPUs and LPUs and generates 
alarms (such as the fan failure, power failure, and over temperature alarms) 
when an anomaly is detected. 
3. Hot standby: The comprehensive hot standby mechanism ensures high 
availability. When an NGFW is faulty, services are smoothly switched to 
the other NGFW without affecting user services. Hot standby implements 
real-time data backup for key configurations and connection entries to 
ensure that firewall performance is not affected by the switchover. You 
can also manually back up data in batches. 
4. Hardware bypass: The built-in bypass card is supported. If the NGFW is 
faulty, traffic is bypassed to ensure service continuity. 
 Energy-saving and eco-friendly design 
Dynamic power consumption management: The NGFW has an architecture that 
uses low power consumption components and high efficiency power supplies 
to reduce power consumption. In addition, system software dynamically 
controls power consumption based on the device operating, function enabling, 
port connection, and temperature statuses, for example, dynamically closing 
idle ports and functional units and adjusting fan speeds. 
Intelligent heat dissipation: The NGFW uses PWM speed adjustment fans and 
reduces power consumption by 70% using the refined speed adjustment and 
area-specific heat dissipation technologies. The technologies also reduce 
noises. 
Eco-friendly manufacturing process: The design and production of the NGFW 
strictly comply with RoHS and WEEE laws and regulations, without any toxic 
substances. The design allows product disassembly and has high recyclability. 
Recyclable materials are widely used, and the product recycle ratio is above 
90%. The packing design complies with the EU requirements 94/62/EC. 
Eco-friendly and recyclable materials are used, and the types, quantity, and 
weight of required materials are reduced. 
Robust Software System 
The USG6000 series uses Huawei-proprietary VRP operating system as its core 
component. Therefore, the USG6000 series itself can prevent unreliable 
elements, such as security vulnerabilities in universal operating systems, 
viruses, and attacks. 
The VRP operating system is a dedicated platform for data communications. Its 
software architecture is customized for data communications devices and has 
taken the development of communications technologies into consideration. The 
USG6000 series not only ensures reliable and secure operating, but can also be

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  20 
   
 
