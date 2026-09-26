---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5-2
title: "fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["compute", "energy", "exploit", "gpu", "gpus", "inference", "intel", "latency", "memory", "research", "robotics"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5.md
source_anchor: ""
source_lines: [3, 52]
sha256: 5fcb26e8e90a47cba755047374c5b1e9111ebb58f74faee5b0a176bda7fdced5
---

# fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5

StorageReview has tested several ruggedized servers, focusing on their performance, capacity, and ease of use. We recently tested the Cheetah RAID solution equipped with Solidigm SSDs. Unfortunately, for obvious reasons, we couldn't test the resilience of these servers. However, Dell has raised the bar even higher by providing an Edge PowerEdge XR7620 server and a rugged transport container, capable of withstanding extreme conditions.
Life is hard at the edge
Rugged edge infrastructure is not a new model. Edge servers have existed for quite some time; they collect data and send it back to headquarters to exploit its full potential. Artificial intelligence has emphasized faster data processing for near-immediate use, which involves processing at the source. Edge AI has had a considerable impact, bringing more innovative, faster, and more secure processing to many concrete applications. A personal assistant always within reach, making real-time decisions, saving you time, money, and sometimes even lives!
Dell PowerEdge XR7620 Configuration
Dell introduced these specially designed PowerEdge servers last year. The model we are examining is the Dell PowerEdge XR7620, a rugged dual-socket server, optimized for the edge, short-depth, specially designed and compact, offering acceleration-focused solutions for the edge. The difference between the XR7620 and classic Edge servers is that this server addresses the rapid maturation of AI/ML with support for the most demanding workloads, including industrial automation, video analytics, point-of-sale analytics, AI inference, and Edge device aggregation.
Our XR7620 is configured with 2x Intel Xeon Gold 6426Y, 128 GB DDR5 4800, 480 GB Dell BOSS RAID1, and 4x Solidigm SSDs. Our tests will take place in Jordan's environmental research laboratory in Canada. So we purchased a 60 TB drive to send the data back to the Storagereview laboratory in Cincinnati.
Dell PowerEdge XR7620 Specifications
| Feature | Technical Specifications | 
|---|---|
| Processor | Two 4th Generation Intel® Xeon® Scalable processors with up to 32 cores per processor | 
| Memory | 16 DDR5 DIMM slots, supports RDIMM 1 TB maximum, speeds up to 4800 5 MT/s. Supports only registered DDRXNUMX ECC DIMMs | 
| Storage Controllers |  | 
| Drive Bays | Front bays: up to 4 x 2.5-inch SAS/SATA/NVMe SSDs, 61.44 TB maximum, up to 8 x E3.S direct NVMe drives, 51.2 TB maximum | 
| Power Supplies |  | 
| Cooling Options | air cooling | 
| Fans | Six hot-swappable cooling fans | 
| Dimensions |  | 
| Weight | Maximum 21.16 kg (46.64 lbs) | 
| Form Factor | 2U rack server | 
| Integrated Management | iDRAC9, iDRAC Direct, iDRAC RESTful API with Redfish, iDRAC Service Module | 
| Bezel | Optional security bezel with dust filter (dust sensor available only for front-access configuration systems) | 
| OpenManage Software |  | 
| Mobility | OpenManage Mobile | 
| OpenManage Integrations |  | 
| Security |  | 
| GPU Options | Up to 5 x 75 W (single-width, full-height/half-length, low-profile) GPUs or up to 2 x 300 W (double-width, full-height/full-length) | 
| Integrated Network Card | 2 x 1 GbE LOM | 
| Network Options | 1 x OCP 3.0 card (optional) | 
| Ports |  | 
| PCIe | 2-processor configuration: up to 5 PCIe slots (4 x16 Gen4/5, 1 x16 LP Gen4) | 
| Operating Systems and Hypervisors |  | 
Dell PowerEdge XR7620 – Critical Edge Features
Security
Like all other PowerEdge servers, the XR7620 is designed with security in mind at every stage of the server lifecycle. The security process begins with a secure supply chain before the server is built. It extends to delivered servers, with Secure Lifecycle Management and Silicon Root of Trust, then secures what is created/stored by the server in Data Protection.
This zero-trust security approach assumes privileged access authorizations requiring validation at every access and implementation point with features such as Identity and Access Management (IAM) and multi-factor authentication (MFA). This is essential for edge deployments where servers are typically installed in unsecured environments. This strict security policy provides the ability to detect and react to tampering or intrusion. Dell's silicon-based Root of Trust platform creates a secure environment that ensures firmware comes from a trusted source. Any detected change in firmware versions or configurations can force PowerEdge to lock the configuration and initiate a restore to the last known good environment.
Management
All Dell PowerEdge servers follow a standard three-tier systems management approach that includes the integrated Dell Remote Access Controller (iDRAC), OpenManage Enterprise (OME), and CloudIQ. This approach provides a unified, simple, automated, and secure systems management solution. It allows you to manage a single server using the iDRAC Baseboard Management Controller (BMC) console, manage thousands of servers simultaneously with OME, and use intelligent infrastructure insights and predictive analytics to maximize server productivity with CloudIQ.
Cooling
We discussed the importance of cooling on overall server performance and efficiency, especially regarding AI. It may not be sexy, but delivering optimal thermal performance is essential when designing resilient and rugged Edge servers. PowerEdge XR servers were designed with balanced and cooling-efficient airflow and, when necessary, comprehensive thermal management that optimizes airflow, regulates fan speeds, and reduces power consumption.
This level of detail allows PowerEdge XR servers to operate between -5 °C and 55 °C. Dell is working on solutions that could extend this operating range even further.
All PowerEdge XR servers are designed with multiple dual counter-rotating fans. This is equivalent to placing two fans in the same housing. It supports N+1 fan redundancy, allowing the server to continue operating with a single fan failure at a maximum temperature of 40 °C for at least four hours.
Scalability
The server's ability to host up to eight NVMe drives and support dense acceleration capabilities positions it well for future advancements in edge data applications and AI inference. Its rugged design and comprehensive cooling system also mean it will continue to operate efficiently as workloads and environmental conditions evolve.
Targeted Workloads
Targeted workloads include digital manufacturing workloads for machine aggregation, VDI, AI inference, OT/IT translation, industrial automation, ROBO, and military applications where rugged design is required. The XR7620 benefits the retail market and is designed for applications such as warehouse operations, point-of-sale aggregation, inventory management, robotics, and AI inference.
Servers like these are deployed by the military, energy exploration and extraction companies, telecommunications, and others needing edge nodes in the most demanding environments. The PowerEdge XR server line is a short-depth 2U, 2-socket server with data center-level compute that delivers high performance, high capacity, and reduced latency.
Design and Build
The Dell PowerEdge XR7620 stands out for its rugged and robust design. It is designed to withstand harsh environments and is often used in environments where reliability is crucial. Typical of all PowerEdge servers, the XR7620 chassis is well-built, with high-quality materials to meet various challenges.
In terms of design, it is a compact and efficient server, making the most of available space while maintaining excellent airflow for cooling.
A key feature of the XR7620 is the smart bezel. It plays an essential role in enhancing the functionality and security of the Dell PowerEdge XR7620 server. The smart bezel is not just an aesthetic addition but an essential component that contributes to the overall efficiency and reliability of the server.
