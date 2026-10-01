---
id: collect-261001-cisco/cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f-1
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f"
domain: cisco
role: reference
task: reference
actors: ["Apple", "United States"]
dates: ["2018-09-05"]
keywords: ["compute", "cost", "datacenter"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f.md
source_anchor: ""
source_lines: [1, 38]
sha256: 8a3afe4ede2e9682caf3485700d283795b219bfd8dda2d0d13a11c71c43f549e
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f

Last Updated: September 5, 2018
About the Cisco Validated Design Program
The Cisco Validated Design (CVD) program consists of systems and solutions designed, tested, and documented to facilitate faster, more reliable, and more predictable customer deployments. For more information, visit:
http://www.cisco.com/go/designzone.
ALL DESIGNS, SPECIFICATIONS, STATEMENTS, INFORMATION, AND RECOMMENDATIONS (COLLECTIVELY, "DESIGNS") IN THIS MANUAL ARE PRESENTED "AS IS," WITH ALL FAULTS. CISCO AND ITS SUPPLIERS DISCLAIM ALL WARRANTIES, INCLUDING, WITHOUT LIMITATION, THE WARRANTY OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT OR ARISING FROM A COURSE OF DEALING, USAGE, OR TRADE PRACTICE. IN NO EVENT SHALL CISCO OR ITS SUPPLIERS BE LIABLE FOR ANY INDIRECT, SPECIAL, CONSEQUENTIAL, OR INCIDENTAL DAMAGES, INCLUDING, WITHOUT LIMITATION, LOST PROFITS OR LOSS OR DAMAGE TO DATA ARISING OUT OF THE USE OR INABILITY TO USE THE DESIGNS, EVEN IF CISCO OR ITS SUPPLIERS HAVE BEEN ADVISED OF THE POSSIBILITY OF SUCH DAMAGES.
THE DESIGNS ARE SUBJECT TO CHANGE WITHOUT NOTICE. USERS ARE SOLELY RESPONSIBLE FOR THEIR APPLICATION OF THE DESIGNS. THE DESIGNS DO NOT CONSTITUTE THE TECHNICAL OR OTHER PROFESSIONAL ADVICE OF CISCO, ITS SUPPLIERS OR PARTNERS. USERS SHOULD CONSULT THEIR OWN TECHNICAL ADVISORS BEFORE IMPLEMENTING THE DESIGNS. RESULTS MAY VARY DEPENDING ON FACTORS NOT TESTED BY CISCO.
CCDE, CCENT, Cisco Eos, Cisco Lumin, Cisco Nexus, Cisco StadiumVision, Cisco TelePresence, Cisco WebEx, the Cisco logo, DCE, and Welcome to the Human Network are trademarks; Changing the Way We Work, Live, Play, and Learn and Cisco Store are service marks; and Access Registrar, Aironet, AsyncOS, Bringing the Meeting To You, Catalyst, CCDA, CCDP, CCIE, CCIP, CCNA, CCNP, CCSP, CCVP, Cisco, the Cisco Certified Internetwork Expert logo, Cisco IOS, Cisco Press, Cisco Systems, Cisco Systems Capital, the Cisco Systems logo, Cisco Unified Computing System (Cisco UCS), Cisco UCS B-Series Blade Servers, Cisco UCS C-Series Rack Servers, Cisco UCS S-Series Storage Servers, Cisco UCS Manager, Cisco UCS Management Software, Cisco Unified Fabric, Cisco Application Centric Infrastructure, Cisco Nexus 9000 Series, Cisco Nexus 7000 Series. Cisco Prime Data Center Network Manager, Cisco NX-OS Software, Cisco MDS Series, Cisco Unity, Collaboration Without Limitation, EtherFast, EtherSwitch, Event Center, Fast Step, Follow Me Browsing, FormShare, GigaDrive, HomeLink, Internet Quotient, IOS, iPhone, iQuick Study, LightStream, Linksys, MediaTone, MeetingPlace, MeetingPlace Chime Sound, MGX, Networkers, Networking Academy, Network Registrar, PCNow, PIX, PowerPanels, ProConnect, ScriptShare, SenderBase, SMARTnet, Spectrum Expert, StackWise, The Fastest Way to Increase Your Internet Quotient, TransPath, WebEx, and the WebEx logo are registered trademarks of Cisco Systems, Inc. and/or its affiliates in the United States and certain other countries.
All other trademarks mentioned in this document or website are the property of their respective owners. The use of the word partner does not imply a partnership relationship between Cisco and any other company. (0809R)
© 2018 Cisco Systems, Inc. All rights reserved.
Table of Contents
VMware vSphere Metro Storage Cluster (vMSC)
Inter Pod Network (IPN) Connectivity
IPN Configuration Requirements
APIC Controller Considerations
Foundation Tenant for FlexPod Connectivity
Compute-to-Storage Connectivity for iSCSI and NFS
Recovery Point Objective and Recovery Time Objective Requirements
Deployment Hardware and Software
Hardware and Software Revisions
Compute to Storage Connectivity
Cisco Validated Designs deliver systems and solutions that are designed, tested, and documented to facilitate and improve customer deployments. These designs incorporate a wide range of technologies and products into a portfolio of solutions that have been developed to address the business needs of the customers and to guide them from design to deployment.
Customers looking to deploy applications using a shared data center infrastructure face a number of challenges. A recurrent infrastructure challenge is to achieve the required levels of IT agility and efficiency that can effectively meet the company’s business objectives. Addressing these challenges requires having an optimal solution with the following key characteristics:
· Availability: Help ensure applications and services availability at all times with no single point of failure
· Flexibility: Ability to support new services without requiring underlying infrastructure modifications
· Efficiency: Facilitate efficient operation of the infrastructure through re-usable policies
· Manageability: Ease of deployment and ongoing management to minimize operating costs
· Scalability: Ability to expand and grow with significant investment protection
· Compatibility: Minimize risk by ensuring compatibility of integrated components
Cisco and NetApp have partnered to deliver a series of FlexPod solutions that enable strategic data center platforms with the above characteristics. FlexPod solution delivers an integrated architecture that incorporates compute, storage, and network design best practices thereby minimizing IT risks by validating the integrated architecture to ensure compatibility between various components. The solution also addresses IT pain points by providing documented design guidance, deployment guidance and support that can be used in various stages (planning, designing and implementation) of a deployment.
In today’s world, business continuity and reliable IT infrastructure are a crucial part of every successful company. Businesses rely on their information systems to run their operations successfully and therefore require their systems, especially their datacenters, to be available with near zero downtime. To support these organizational goals, datacenter architects have been exploring various high availability solutions to improve the availability and resiliency of datacenter services. In traditional single site datacenter architectures, high availability solutions mostly comprise of technologies such as application clustering and active-standby services architectures designed to improve the robustness of local systems i.e. systems within the single datacenter. To safeguard against site failures due to power or infrastructure outages, datacenter solutions that span multiple sites have become exceedingly important.
This document describes integration of the Cisco ACI Multi-Pod and NetApp MetroCluster IP solution into the FlexPod Datacenter to provide a highly available multi-datacenter solution. The multi-datacenter architecture offers the ability to balance workloads between two datacenters utilizing non-disruptive workload mobility thereby enabling migration of services between sites without the need for sustaining an outage.
The FlexPod with ACI Multi-Pod and NetApp MetroCluster IP solution provides the following benefits:
· Seamless workload mobility across Data Centers
· Consistent policies across the sites
· Layer-2 extension across geographically dispersed DCs
· Enhanced downtime avoidance during maintenance
· Disaster avoidance and recovery
FlexPod solution is a pre-designed, integrated and validated architecture for data center that combines Cisco UCS servers, Cisco Nexus family of switches, Cisco MDS fabric switches and NetApp Storage Arrays into a single, flexible architecture. FlexPod is designed for high availability, with no single points of failure, while maintaining cost-effectiveness and flexibility in the design to support a wide variety of workloads.
