---
id: collect-261001-cisco/cisco/gr-blogs-enterprise-networking-2023-smart-simplified-campus-network-solution-htm-c860aa29-1
title: "gr-blogs-enterprise-networking-2023-smart-simplified-campus-network-solution-htm-c860aa29"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["energy", "lean"]
source: docs/RAG/collect-261001-cisco/gr-blogs-enterprise-networking-2023-smart-simplified-campus-network-solution-htm-c860aa29.md
source_anchor: ""
source_lines: [1, 30]
sha256: 5bb73c98c3e9469570865f5a4186c7c7a8289df8e4ef881843aefcc89bc0d141
---

# gr-blogs-enterprise-networking-2023-smart-simplified-campus-network-solution-htm-c860aa29

Enterprise products, solutions & services
Cloud products, solutions & services
Products, Solutions and Services for Carrier
Phones, laptops, tablets, wearables & other devices
About Huawei, Press & Events , and More
Nowadays, diversified communication channels and devices provide consumers with unprecedented options for better interactions. In 2022, the volume of interactions through digital channels grew 68% year-on-year (YOY), whereas traditional channels such as SMS and email saw a YOY growth of 75% and 91% respectively. As shown by statistics, users' requirements for automation, intelligent operations, and interactive experience increase dramatically. As such, a more convenient and smooth user journey with faults being quickly resolved is urgently needed.
However, the reality is too harsh for many enterprise users. An advanced and premier digital office experience cannot be assured due to incomplete network capabilities. Worse yet, engineers responsible for network operations now suffer from increased workloads and user complaints.
A traditional network consists of the core layer, aggregation layer, and access layer. With the rapid rise of wireless and IoT services, as well as more complicated services, traditional three-layer networks face many challenges:
• The inflexible network architecture makes terminal capacity expansion costly and time-consuming.
• The bandwidth of network cables does not support long-term evolution, resulting in a waste of investment due to frequent re-cabling.
• The number of access devices increases exponentially, complicating management and causing high power consumption.
Huawei is a top vendor in the global enterprise wired and wireless LAN infrastructure market and has been named a leader in 2022 Gartner® Magic Quadrant™. To cope with the changes in enterprise office campus scenarios, Huawei launches the Intelligent Simplified Campus Network Solution. This solution features simplified architecture, energy efficiency, and always-on services, helping enterprises build a high-quality campus network.
Huawei's Intelligent Simplified Campus Network Solution simplifies the three-layer traditional network architecture into a two-layer one. In this solution, traditional switches are deployed at the core layer, and central switches and remote units (RUs) are deployed at the access layer. It is recommended that Huawei CloudEngine S series switches be used as central switches, and Huawei CloudEngine S5731-L series remote unit switches be deployed as RUs.
The core of the simplified architecture is to manage RUs in a unified manner through a central switch. Huawei's Intelligent Simplified Campus Network Solution leverages the distributed concept of managing APs through a WAC. In addition, it innovatively proposes the Huawei proprietary eXtremely Lean Discovery Protocol (XLDP) to implement status negotiation and maintenance between the central switch and RUs. Based on this protocol, the central switch can:
• Manage RU onboarding. After RUs are powered on, they exchange XLDP packets with the central switch for both onboarding and authentication negotiation. The RUs are configuration-free and support plug-and-play.
• View RU status, such as port status, port statistics, PoE power supply information, and optical module information.
• Deliver configurations to RUs based on interconnection interfaces. For example, the interface auto-negotiation mode and PoE power supply mode of RUs can be set.
• Upgrade the firmware of RUs for function expansion. RUs run only lightweight firmware such as BIOS firmware, APP firmware, and PoE firmware. In this way, their startup performance can be improved.
In actual application scenarios, the simplified architecture can bring many benefits to enterprise users. For example, as the port expansion modules of the central switch, RUs can be deployed on desktops, supporting the access of adjacent terminals. This helps enterprises slash the cabling from the extra low voltage (ELV) room to terminals. If 8-port RUs are used, cabling can be saved by 87.5%. When network access is required by new terminals, they can directly connect to RUs, implementing flexible network expansion. RUs are centrally managed by the central switch. Besides, they are planning-free, configuration-free, management-free, and support plug-and-play. This greatly simplifies campus network planning, deployment, and management. An 8-port RU, for example, can be used together with one 48-port access switch (as the central switch) to support the access of up to 384 wired terminals, equivalent to 8 traditional 48-port access switches. This equates to an 87.5% reduction in the number of managed nodes.
As carbon emission reduction becomes an increasingly important topic worldwide due to global warming, green and low-carbon networks are required. Huawei's Intelligent Simplified Campus Network Solution builds green and low-carbon networks based on network-level architecture optimization and device-level energy saving.
• Network-level architecture optimization: The simplified architecture uses central switches and RUs as combinations, simplifying the traditional campus network architecture. This greatly reduces cabling and the number of traditional access devices, avoiding unnecessary resource waste and energy consumption.
• Device-level energy saving: The energy-saving design is mainly reflected in heat dissipation, power modules, and chips. The whole series of RUs are noise-free and energy-saving thanks to their fanless design that achieves natural heat dissipation. They are also equipped with chips and components supporting ultra-low power consumption. The typical power consumption of each RU is less than 7 W, and the average power consumption of a single port is less than 1 W. This helps to reduce power consumption by about 30% compared with the industry average. Central switches function as the power supply center of the entire network. Therefore, higher requirements for heat dissipation and power supply are required on the switches.
Huawei CloudEngine S8700 series switches (the next-generation central switches), for example, provide the following features:
• Innovative Y+Psi (ψ)-type ventilation channel to support front-to-rear airflows, improving heat dissipation efficiency by 15%
• Huawei-developed ψ fans that adopt Huawei's patented mixed-flow single-rotor technology and high-performance airfoil algorithm, providing 10%+ lower fan speed, 30%+ lower power consumption, and 15%+ lower noise than the industry average in the same service scenario
• Power module: integrated magnetic core technology, reducing the wind resistance of power modules by 30%, improving the heat dissipation efficiency by 40%
In an office campus with 2000 users, for example, Huawei's Intelligent Simplified Campus Network Solution delivers the following benefits:
• 38% lower initial investments (including the procurement and deployment costs of devices, cables, and cabinets)
• 30% cut in electricity expenditure on the entire network, saving 25,000 kWh in a year, equivalent to reducing the use of 10.1 tons of standard coal or planting 1395 new trees
On a traditional campus network, PoE provides data access and power supply for terminals through network cables, making terminal deployment more flexible. However, due to the limitation of data transmission over network cables, the access distance for terminals is limited to 100 m. If, on the other hand, optical fibers are used, the data transmission issue is resolved, but power supply becomes a pain point. How to ensure both data transmission and long-distance power supply?
