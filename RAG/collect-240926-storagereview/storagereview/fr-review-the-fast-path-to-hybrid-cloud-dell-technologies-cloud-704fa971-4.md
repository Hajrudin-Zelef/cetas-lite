---
id: collect-240926-storagereview/storagereview/fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971-4
title: "fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971"
domain: storagereview
role: reference
task: reference
actors: ["United States"]
dates: ["2019-11"]
keywords: []
source: docs/RAG/clean_en/storagereview/fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971.md
source_anchor: ""
source_lines: [52, 68]
sha256: ea2c6086c0741dbc372d2ca9dfac4cc845937671e5d4ff582f11f1c68b4ddf53
---

# fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971

VMware Cloud Foundation Lifecycle Management Results on VxRail
When performing upgrades with VCF on VxRail, VMware SDDC Manager is your central source for upgrading the full stack. Dell Technologies co-designed with VMware the VxRail hardware upgrade from SDDC Manager. Dell Technologies uses a synchronous release schedule aligned with VMware updates. This means you only have to visit one place to find and install available updates. Another interesting feature is ordered installation. Often, you must be on a certain revision before you can update. Upgrading via SDDC Manager shows you only the current version you can install based on software versions; once that is met, you can proceed with the upgrade.
Compatibility and environment readiness for the upgrade are also performed. The ability to perform a pre-check of the environment beforehand makes upgrading a thing of the past. This also saves a lot of time by avoiding manually checking compatibility across the full stack. The engineering time and optimized integrations between Dell EMC VxRail and VMware systems are great benefits for businesses. This time savings for organizations could not be quantified during this test cycle.
As shown in the introductory graph, we found a 15% time savings with Dell EMC VxRail compared to our custom approach using the VMware Validated Design (VVD) platform. It is important to specify that this percentage takes into account the fact that the VVD LCM platform required no hardware updates. The VxRail deployment proved faster than updating the VVD LCM platform through a hardware firmware upgrade, designed for optimal performance and stability. Conversely, the VVD LCM platform only checks firmware compatibility and is not optimized for maximum performance or stability.
The VCF upgrade process on VxRail, both for the VMware SDDC stack and the VxRail hardware firmware, took a time of 11:16:24 to complete, for both the management domain and the workload domain deployed for testing.
DIY using VMware Validated Design Lifecycle Management Results
The update procedure for the deployed VMware Validated Design solution is not as simple as checking for updates and patches like VCF on the VxRail solution. Here, you must manually check the software revisions of each installed VMware SDDC component, then check the hardware of each host to ensure it is compatible with the versions you are upgrading to. For this test, we did not have to upgrade any hardware firmware because everything was compatible with the VMware software versions we were upgrading to.
A crucial step when upgrading software is to ensure that the underlying hardware is compatible with the new version of the VMware SDDC components. Once we verified and downloaded the upgrade bits, we then checked the compatibility of each host. This was a manual process of logging into the iDRAC of all 8 hosts and checking the firmware revisions on each of the installed drives, storage controllers, and network interface card adapters. Once we had our list, we then had to go to the VMware Hardware Compatibility List site and manually search and verify each of the components.
To perform the actual upgrade of the VMware SDDC components to go from VVD 5.1.1 to 5.1.2, the following software components are upgraded:
- Platform Services Controller appliances
- vCenter Server appliances
- vSphere Update Manager Download Service
- ESXi hosts
The update of the management and workload clusters took 12 h 59 min 16 s. Only the VMware software components were updated during this LCM test.
This report is sponsored by Dell Technologies. All views and opinions expressed in this report are based on our impartial view of the product(s) under study.
[1] Offer valid for certain preconfigured solutions. Contact your sales representative for more information. Offer not valid for orders of more than 1,000 instances, hybrid storage, certain vRealize components (vRA, vRO), and certain other features. Customer credit approval, site survey, and configuration specification must be completed before order validation. Product availability, delivery times, holidays, and other factors may impact deployment time. Deployment includes delivery, standardized installation, and hardware and software configuration. Offer valid only in the United States, United Kingdom, France, and Germany.
[2] Based on internal Dell Technologies analysis, November 2019
