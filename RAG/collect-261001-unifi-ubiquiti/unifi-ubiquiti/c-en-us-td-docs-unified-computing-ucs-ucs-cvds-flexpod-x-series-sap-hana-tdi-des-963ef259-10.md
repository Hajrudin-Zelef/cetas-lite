---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-10
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["agent", "agents", "compute", "cost", "datacenter", "ethernet", "intel", "memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [319, 369]
sha256: bf199ec002a03e70e7830bfd2898fffe4e69d03a2e722c96d51e91acf6cf220d
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

The AppDynamics APM Platform enables you to monitor and manage your entire application-delivery ecosystem, from the mobile app or browser client request through your network, backend databases and application servers and more. AppDynamics APM gives you a single view across your application landscape, letting you quickly navigate from the global perspective of your distributed application right down to the call graphs or exception reports generated on individual hosts.
AppDynamics has an agent-based architecture. Once the agents are installed you receive a dynamic flow map or topography of your application. It uses the concept of traffic lights to indicate the health of your application (green is good, yellow is slow, and red indicates potential issues) with dynamics baselining. AppDynamics measures application performance based on business transactions which essentially are the key functionality of the application. When the application deviates from the baseline AppDynamics captures and provides deeper diagnostic information to help be more proactive in troubleshooting and reduce the MTTR (Mean Time To Repair).
For more information about SAP monitoring using AppDynamics, see: https://docs.appdynamics.com/display/SAP/SAP+Monitoring+Using+AppDynamics
● SAP HANA System Implementation Options
● Interoperability and Feature Compatibility
● Network and SAN Design Considerations
● Cisco Nexus Ethernet Connectivity
● Cisco MDS SAN Connectivity – Fibre Channel Only Design
● NetApp AFF A400 – Storage Virtual Machine (SVM) Design
● VMware vSphere – ESXi Design
● Cisco Intersight Integration with VMware vCenter and NetApp Storage
The FlexPod Datacenter with Cisco UCS X-Series and Intersight solution delivers a cloud-managed infrastructure solution on the latest Cisco UCS hardware for both bare-metal as well as virtualized implementations of SAP HANA. For the virtualized deployments, VMware vSphere 7.0 U3i hypervisor is installed on the Cisco UCS X210c M6 Compute Nodes configured for stateless compute design using boot from SAN. NetApp AFF A400 provides the required storage infrastructure. The Cisco Intersight cloud-management platform is utilized to configure and manage the infrastructure. The solution requirements and design details are covered in this section.
The section explains the SAP HANA system requirements defined by SAP followed by the reference architecture of FlexPod Datacenter Solution providing the platform for SAP and SAP HANA.
The FlexPod Datacenter with Cisco UCS X-Series for SAP HANA TDI design meets the following general design requirements:
● Resilient design across all layers of the infrastructure with no single point of failure.
● Scalable design with the flexibility to add compute capacity, storage, or network bandwidth as needed.
● Modular design that can be replicated to expand and grow as the needs of the business grow.
● Simplified design with ability to integrate and automate with external automation tools.
● Cloud-enabled design which can be configured, managed, and orchestrated from the cloud using GUI or APIs.
● The FlexPod solution is SAP HANA TDI certified to provide organizations with the flexibility to choose the best, cost-effective, and appropriate solution that meets their needs.
SAP HANA System Implementation Options
This section defines the basic requirements for available implementation options with Cisco X210C M6 based FlexPod DC.
Single SAP HANA System on a Single Node: Scale-Up (Bare Metal or Virtualized)
A scale-up TDI solution is the simplest of the installation types. All data and processes are located on the same server in this single-node solution. SAP HANA scale-up TDI solutions are based on X-series compute node and use the intended external storage.
The network requirements for this option depend on the client and application server connectivity, backup/storage connectivity, and optional system replication services access needs. At a minimum application server access network and bandwidth factored for the data, log and shared filesystem access storage networks are required to run SAP HANA in a scale-up configuration.
Co-existing SAP HANA and SAP Application Workloads
Scenarios where SAP HANA database bare metal installation along with virtualized SAP application workloads are common in the datacenter. With SAP HANA TDI it is possible to run SAP HANA on shared infrastructure that also hosts non-HANA workloads like standard SAP applications. It is important to ensure appropriate storage IO and network bandwidth segregation so that HANA systems get their due to comfortably satisfy the storage and network KPIs for production support.
Interoperability and Feature Compatibility
The different hardware and software compatibility tools are available at the following links:
● Cisco UCS Hardware and Software Interoperability Matrix
● Cisco MDS and Nexus Interoperability Matrix
● NetApp Interoperability Matrix Tool
In addition to the hardware components the software product features need to fully integrate with SAP solutions which is confirmed with SAP certifications and SAP notes accordingly:
● Certified and supported SAP HANA hardware
● SAP note 2235581 – SAP HANA: Supported Operating Systems
● SAP note 2937606 – SAP HANA on VMware vSphere 7.0 in production
To achieve the performance and reliability requirements for SAP HANA it is vital to select the correct components and configuration for the SAP landscape.
Bare-metal Installation
The existing core-to-memory ratios for SAP HANA bare-metal environments are dependent on the Intel CPU architecture and the type of SAP data processing: online analytical processing (OLAP), online transaction processing (OLTP), or a mixed data processing system like with SAP Suite on/for HANA (SoH/S4H).
With these dependencies the 2-socket, Intel Ice Lake CPU architecture-based Cisco UCS X210c M6 compute node can scale up to 2 TB DDR main memory for SAP Business Warehouse (BW) systems or 4 TB DDR main memory for SAP Suite systems.
With SAP expert sizing mixed memory configurations of DDR memory and Intel Persistent Memory (PMem) using the AppDirect mode of the Intel PMem modules can increase the amount of available memory for the SAP HANA in-memory database further.
Virtualized installation
Since SAP HANA TDI Phase 5, it is possible to perform a workload-based sizing (SAP note 2779240) which can deviate from the existing core-to-memory ratio if the following conditions are met:
● Certified SAP HANA hardware
● Validated hypervisor
● Deviations are within the upper and lower limits of the hypervisor
VMware vSphere and Intel Ice Lake CPUs are validated for SAP HANA starting with VMware vSphere 7.0 U3c. Sizing of the virtualized SAP HANA machines (vHANA) depends on the CPU model, number of cores and number of CPU sockets.
The minimum requirement is a 2-CPU socket node, 0.5-CPU socket reserved with 8vCPUs based on 8 physical cores and 128 GB main memory. The upper limits dependent on the number of sockets, CPU models and cores, and VMware vSphere versions:
● Cisco UCS M6 X-Series (Ice Lake)
● Cisco UCS X210c M6: 160 vCPUs and 2-CPU socket wide VMs
The recommended approach to configure vHANA machines is to match the actual hardware configuration in regards of the number of cores per socket and available total amount of memory. For example, a 0.5-CPU socket configuration for the Intel Xeon Platinum 8380 processor with 40 cores per socket, configure 20 physical cores and ¼ of the available main memory. If SAP HANA requires more memory double the physical cores and memory. Odd VM configurations like 1.5 or 2.5-CPU sockets are not allowed. It is possible to run up to 4 individual vHANA production machines on a single Cisco UCS X210c M6 compute node.
