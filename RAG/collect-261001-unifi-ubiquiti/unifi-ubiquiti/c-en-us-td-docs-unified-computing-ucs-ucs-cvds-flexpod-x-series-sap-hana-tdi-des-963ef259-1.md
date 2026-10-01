---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-1
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "datacenter", "energy", "liquid cooling", "memory", "power delivery", "research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [1, 46]
sha256: ba86e7c400455a05355f452ca453d4382988c0b52d7db5df02d20304294331bf
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

In partnership with:
About the Cisco Validated Design Program
The Cisco Validated Design (CVD) program consists of systems and solutions designed, tested, and documented to facilitate faster, more reliable, and more predictable customer deployments. For more information, go to: http://www.cisco.com/go/designzone.
The FlexPod Datacenter solution is a validated approach for deploying Cisco and NetApp technologies and products to build shared private and public cloud infrastructure. Cisco and NetApp have partnered to deliver a series of FlexPod solutions that enable strategic data-center platforms. The success of the FlexPod solution is driven through its ability to evolve and incorporate both technology and product innovations in the areas of management, compute, storage, and networking.
This document explains the design details of incorporating the Cisco X-Series modular platform into the FlexPod Datacenter for SAP HANA Tailored Datacenter Integration (TDI) implementations and its ability to manage and orchestrate FlexPod components from the cloud using Cisco Intersight. Some of the key advantages of integrating Cisco UCS X-Series into the FlexPod infrastructure are:
The key benefits of this solution are:
● Simpler and programmable infrastructure: infrastructure as a code delivered through a single partner integrable open API
● Power and cooling innovations: higher power headroom and lower energy loss due to a 54V DC power delivery to the chassis
● Better airflow: midplane-free design with fewer barriers, therefore lower impedance
● Fabric innovations: PCIe/Compute Express Link (CXL) topology for heterogeneous compute and memory composability
● Innovative cloud operations: continuous feature delivery and no need for maintaining on-premise virtual machines supporting management functions
● Built for investment protections: design ready for future technologies such as liquid cooling and high-Wattage CPUs; CXL-ready
In addition to the compute-specific hardware and software innovations, the integration of the Cisco Intersight cloud platform with VMware vCenter and NetApp Active IQ Unified Manager delivers monitoring, orchestration, and workload optimization capabilities for different layers (virtualization and storage) of the FlexPod infrastructure.
If you’re interested in understanding the FlexPod design and deployment details, including the configuration of various elements of design and associated best practices, refer to Cisco Validated Designs for FlexPod, here: https://www.cisco.com/c/en/us/solutions/design-zone/data-center-design-guides/flexpod-design-guides.html.
This chapter contains the following:
● Audience
The Cisco Unified Computing System (Cisco UCS) X-Series is a brand-new modular compute system, configured and managed from the cloud. It is designed to meet the needs of modern applications and to improve operational efficiency, agility, and scale through an adaptable, future-ready, modular design. The Cisco Intersight platform is a Software-as-a-Service (SaaS) infrastructure lifecycle management platform that delivers simplified configuration, deployment, maintenance, and support.
SAP HANA in-memory database handles transactional and analytical workloads with any data type – on a single data copy. It breaks down the transactional and analytical silos in organizations, for quick decision-making, on premise and in the cloud. SAP HANA offers a multi-engine, query-processing environment that supports relational data (with both row- and column-oriented physical representations in a hybrid engine) as well as graph and text processing for semi-structured and unstructured data management within the same system. The SAP HANA TDI offers a more open and flexible way for the integration of SAP HANA into the data center with benefits like the virtualization of the SAP HANA platform or a flexible combination of multiple SAP HANA systems on the fully certified, converged infrastructure.
Powered by the Cisco Intersight cloud-operations platform, the Cisco UCS X-Series enables the next-generation cloud-operated FlexPod infrastructure that not only simplifies data-center management but also allows the infrastructure to adapt to the unpredictable needs of modern applications as well as traditional workloads. With the Cisco Intersight platform, you get all the benefits of SaaS delivery and the full lifecycle management of Intersight-connected distributed servers and integrated NetApp storage systems across data centers, remote sites, branch offices, and edge environments.
The intended audience of this document includes but is not limited to IT architects, sales engineers, field consultants, professional services, IT managers, partner engineering, and customers who are interested in learning about and deploying the FlexPod Datacenter for SAP and SAP HANA use cases, such as Scale-up system, SAP HANA scale out system with NFS or mixed protocol configurations with Linux bare metal as well as VMware ESXi installations.
This document compliments the FlexPod Datacenter with Cisco UCS X-Series, VMware 7.0 U2, and NetApp ONTAP 9.9 Design Guide: https://www.cisco.com/c/en/us/td/docs/unified_computing/ucs/UCS_CVDs/flexpod_xseries_esxi7u2_design.html and explains the design considerations and requirements for SAP HANA TDI deployments, both in virtualized as well as bare-metal scenarios.
The following design elements distinguish this version of FlexPod from previous models:
● Integration of Cisco UCS X-Series into FlexPod Datacenter for SAP HANA TDI
● Deploying and managing Cisco UCS X-Series from the cloud using Cisco Intersight
● Integration of Cisco Intersight with NetApp Active IQ Unified Manager for storage monitoring and orchestration
● Integration of Cisco Intersight with VMware vCenter for interacting with, monitoring, and orchestrating the virtualized SAP HANA environment
The FlexPod Datacenter solution with Cisco UCS X-Series and NetApp ONTAP 9.9.1 offers the following key benefits:
● Simplified cloud-based management of the solution components
● Hybrid-cloud-ready, policy-driven modular design
● Highly available and scalable platform with flexible architecture that supports various deployment models
● Easy to deploy, consume, and manage architecture, which saves time and resources required to research, procure, and integrate off-the-shelf components
● Support for component monitoring, solution automation and orchestration, and workload optimization
● Stateless architecture, providing the capability to expand and adapt to new business requirements
● Robust components capable of supporting high-performance and high bandwidth for virtualized and non-virtualized applications
● Risk reduction at each level of the design with resiliency built into each touch point
Like other FlexPod solution designs, FlexPod Datacenter with Cisco UCS X-Series, and Intersight for SAP HANA TDI is configurable according to demand and usage. You can purchase the exact infrastructure you need for your current application requirements and can then scale-up by adding more resources to the FlexPod system or scale-out by adding more FlexPod instances. By moving the management from the fabric interconnects into the cloud, the solution can respond to the speed and scale of your deployments with a constant stream of new capabilities delivered from Intersight software-as-a-service model at cloud-scale.
● Cisco Unified Computing System X-Series
● Cisco Nexus Switching Fabric
● Cisco MDS 9132T 32G Multilayer Fabric Switch
● Cisco Nexus Dashboard Fabric Controller
● Red Hat Ansible Automation Platform
● Red Hat Enterprise Linux for SAP Solutions
● SUSE Linux Enterprise Server for SAP Applications
● SAP Application Monitoring with AppDynamics
FlexPod Datacenter architecture is built using the following infrastructure components for compute, network, and storage:
● Cisco Unified Computing System (Cisco UCS)
