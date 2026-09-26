---
id: collect-240926-storagereview/storagereview/fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971-2
title: "fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971"
domain: storagereview
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971.md
source_anchor: ""
source_lines: [3, 32]
sha256: 68c2461f3120858852afcb6c1ea65a236c6e0fdc9c223a4741ae92907fb83237
---

# fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971

86% faster to deploy VMware Cloud Foundation than a do-it-yourself model
VMware Cloud Foundation on Dell EMC VxRail (Dell Technologies APEX Hybrid Cloud) offers a simplified and rapid solution for deploying a hybrid cloud. Not only is the solution operational on-site in 14 days [1] , but commissioning VMware Cloud Foundation on VxRail is also much faster than an installation on dedicated servers. Indeed, during our evaluation, we found that deploying VCF on VxRail was 86% faster than the on-site installation model.
However, the benefits of VCF on VxRail continue with regard to lifecycle updates (hardware and software). VCF continues to show enormous value, being 15% faster, with fewer steps. Ultimately, the VCF on VxRail solution is easier than do-it-yourself to deploy and manage with hundreds of automated tasks. In this article, we detail the installation and configuration of VCF in these two modalities and highlight the technical advantages of Dell Technologies APEX Hybrid Cloud.
Introduction
Within Dell Technologies' enterprise portfolio is the company's flagship solution for the modern software-defined data center (SDDC). Dell EMC VxRail is an integrated hyperconverged infrastructure (HCI) appliance jointly designed by Dell Technologies and VMware. While VMware offers HCI software through partners who sell vSAN ReadyNodes™ (which Dell EMC also offers), VxRail goes even further. With VxRail, customers receive a fully integrated, preconfigured, and pretested system that delivers virtualization, compute, and storage in a single appliance. The fact that all elements (including VMware software and Dell EMC PowerEdge hardware and networking) are brought together in a single unit offers customers a smoother path to deploying VMware HCI.
However, Dell Technologies doesn't stop there with VxRail. Customers wishing to adopt a hybrid cloud vision, leverage containers for modern applications, or deploy a true software-defined data center (SDDC) can deploy VMware Cloud Foundation (VCF) on Dell EMC VxRail. VxRail is the first hyperconverged infrastructure system fully integrated with VMware Cloud Foundation (VCF) SDDC Manager [2] , offering a fully integrated set of VMware software components, including vSphere, vRealize, NSX, vSAN, and SDDC Manager.
The integration of VCF with VxRail offers customers a unified platform that constitutes a complete and automated experience across the entire hardware and software stack. Through this tight integration, customers will benefit from smooth and rapid deployment and a simplified management experience while enjoying infrastructure agility that can accelerate their organization's ability to deliver applications. Additionally, due to the deep integration between hardware and software, VxRail also offers crucial operational benefits in lifecycle management.
While the rapid bringing online of a VxRail cluster provides an immediate benefit in terms of business impact, the ongoing operational benefits offer the most impressive results. These range from the obvious, such as no longer searching for the latest supported drivers for items such as network cards, SSDs, and other installed components, to searching for VMware software patches/updates. But there is also the fact that Dell Technologies includes new VMware features within 30 days on VxRail and Dell Technologies serves as a single point of contact for all support issues.
Quickly bringing new features to customers is also a considerable advantage. For example, in mid-2020, VMware released several updates around its Tanzu software that allows customers to run Kubernetes from a single control plane. For VxRail customers adopting modern application delivery, VCF on VxRail offers a turnkey process to bring Tanzu online. Operationally, this offers customers a consistent way to deploy and manage traditional virtual machines alongside containers.
Given the breadth of VMware SDDC technologies and their adoption throughout the enterprise, customers have two distinct options when it comes to deployment. As such, we sought to compare the benefits of VCF on VxRail against the "Do-It-Yourself" alternative. We started by deploying VCF on VxRail, following the process of deploying the hosts, the cloud builder, and finally, hardware and software lifecycle updates.
At the end, we reallocated exactly the same hardware to vanilla PowerEdge servers, installing the individual components as an organization would if it were deploying vSphere, vSAN, NSX, and vRealize components. For the lifecycle management phase, we manually performed the upgrade from VVD 5.1.1 to 5.1.2.
The table below highlights these three definable segments, but it should be noted that lifecycle actions will be a perpetual task in which organizations will engage regularly.
While we find immediate deployment benefits (meaning customers will be online and delivering faster with VxRail), the benefits of ongoing lifecycle management will save time and allow the business to focus its efforts elsewhere throughout the cluster's lifecycle.
While these are high-level comparison figures, the following report describes these results in detail with the technical processes and the time required to complete each step. Although the cumulative accelerated results tell the story, the detail of the process and the accompanying task list clarify the differences between purchasing an appliance designed in VxRail from Dell EMC and standard x86 bare metal servers.
Finally, accounting for the benefits of VxRail neglects to take into account the time between order and delivery. Dell Technologies offers several fixed factory configurations that can be on-site and fully deployed in two weeks.
Technical Overview
The main objective was to quantify the value brought by the turnkey solution jointly designed with VCF on VxRail compared to building piece by piece in the DIY approach using VMware Validated Design. The tests were divided into three parts: building the management domain, the workload domain, and LCM. By doing so, we were able to gather a comparable time to complete the results on both resolution methods.
For testing purposes, we used an eight-node cluster; four nodes were configured for the management domain and four nodes for the workload domain.
Servers – 8 Dell EMC PowerEdge R640
- Memory - 576 GB of RAM
- Networking – 2 x Mellanox25GbE 2P ConnectX4LX
- Storage – 4 x 3.84 TB SSDs Total capacity 15.4 TB
- Storage controller – Dell EMC HBA330 Mini
All tests were performed starting with all hosts in a powered-off state. Dell Technologies and VMware installation and configuration documentations were followed.
VMware Software Versions
- VMware Cloud Foundation – Initial version 3.9.0 -Upgrade 3.10
- VMware Validated Design Initial Version 5.1.1 – Upgrade 5.1.2
Dell Technologies APEX Hybrid Cloud Deployment Results
Dell Technologies streamlined the initial build process with factory-preinstalled software. In less than two hours (01:54:44), we had our four management nodes powered on, ESXi installed and configured with Jumbo Frames, clustered, and the Cloud Builder appliance deployed, and we were ready to initialize the VMware Cloud Foundation deployment.
When installing and configuring VCF on VxRail, the process is primarily wizard-guided, entering crucial information that is extracted into a JSON file that the appliance uses to configure the specifics of networking to hostnames and IP addresses on the various virtual machines. This approach is similar to DIY using the VMware Validated Design method described later in the report. You fill out an environment workbook that is ingested into the wizard. You will still need to verify the information entered and enter passwords. We found this experience easy to use and it allowed for quick verification of information before the build.
