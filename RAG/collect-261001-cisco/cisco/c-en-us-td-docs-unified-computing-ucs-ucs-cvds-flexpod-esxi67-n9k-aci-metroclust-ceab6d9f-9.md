---
id: collect-261001-cisco/cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f-9
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f"
domain: cisco
role: reference
task: reference
actors: ["California", "United States"]
dates: []
keywords: ["compute", "datacenter"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f.md
source_anchor: ""
source_lines: [287, 340]
sha256: 646c59e086e9d6518ba32b20eb90b6c741910bb8a9affdc53e543e96427a88f1
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f

FlexPod Datacenter with Cisco ACI Multi-Pod, NetApp MetroCluster IP and VMware vSphere 6.7 solution is validated for successful infrastructure configuration and availability across the two geographically dispersed datacenters by validating a wide variety of test cases and by simulating partial and complete failure scenarios. The types of tests executed on the system are listed below.
· Validate SAN boot configuration and vSphere Installation for iSCSI associated boot LUN
· Verify SAN access and boot with primary path to local storage array unavailable
· Verify single storage controller connectivity failure in a DC
· Verify complete storage system isolation in a DC
· Verify single Host failure in a DC
· Verify compute isolation in a DC (all hosts failed)
· Verify loss of a single Leaf Switch (one after another)
· Verify loss of a single Spine (one after another)
· Verify loss of an IPN device (one after another)
· Verify loss of a single path between the IPN devices (one after another)
· Path failure between a pair of Nexus 3Ks
· Single Nexus 3K device failure
· Failure of all links between the Nexus 3Ks
· Storage partition – Link between the two DCs goes down making two storage systems isolated
· Datacenter Isolation and recovery for compute and storage
· Datacenter maintenance use case validation – move all workloads to a single DC and back
· Datacenter failure and recovery – power out failure of a DC
FlexPod Datacenter with Cisco ACI Multi-Pod, NetApp MetroCluster IP and VMware vSphere 6.7 solution allows interconnecting and centrally managing two or more ACI fabrics deployed in separate, geographically dispersed datacenters. NetApp MetroCluster IP provides a synchronous replication solution between two NetApp controllers providing storage high availability and disaster recovery in a campus or metropolitan area. This validated design enables customers to quickly and reliably deploy VMware vSphere based private cloud on a distributed integrated infrastructure thereby delivering a unified solution which enables multiple sites to behave in much the same way as a single site.
· Campus-wide and metro-wide protection and provide WAN based disaster recovery
Cisco Unified Computing System:
http://www.cisco.com/en/US/products/ps10265/index.html
Cisco UCS 6200 Series Fabric Interconnects:
http://www.cisco.com/en/US/products/ps11544/index.html
Cisco UCS 5100 Series Blade Server Chassis:
http://www.cisco.com/en/US/products/ps10279/index.html
Cisco UCS B-Series Blade Servers:
Cisco UCS C-Series Rack Servers:
http://www.cisco.com/c/en/us/products/servers-unified-computing/ucs-c-series-rack-servers/index.html
Cisco UCS Adapters:
http://www.cisco.com/en/US/products/ps10277/prod_module_series_home.html
Cisco UCS Manager:
http://www.cisco.com/en/US/products/ps10281/index.html
Cisco Nexus 9000 Series Switches:
Cisco Application Centric Infrastructure:
VMware vCenter Server:
http://www.vmware.com/products/vcenter-server/overview.html
NetApp Data ONTAP:
http://www.netapp.com/us/products/platform-os/ontap/index.aspx
NetApp AFF A700:
https://www.netapp.com/us/products/storage-systems/all-flash-array/aff-a-series.aspx
https://www.netapp.com/us/products/backup-recovery/metrocluster-bcdr.aspx
Cisco UCS Hardware Compatibility Matrix:
https://ucshcltool.cloudapps.cisco.com/public/
VMware Compatibility Guide:
http://www.vmware.com/resources/compatibility
NetApp Interoperability Matric Tool:
http://mysupport.netapp.com/matrix/
Haseeb Niazi, Technical Marketing Engineer, Cisco Systems, Inc.
Haseeb Niazi has over 19 years of experience at Cisco in the Data Center, Enterprise and Service Provider Solutions and Technologies. As a member of various solution teams and Advanced Services, Haseeb has helped many enterprise and service provider customers evaluate and deploy a wide range of Cisco solutions. As a technical marking engineer at Cisco UCS Solutions group, Haseeb focuses on network, compute, virtualization, storage and orchestration aspects of various Compute Stacks. Haseeb holds a master's degree in Computer Engineering from the University of Southern California and is a Cisco Certified Internetwork Expert (CCIE 7848).
Arvind Ramakrishnan, Solutions Architect, NetApp, Inc.
Arvind Ramakrishnan works for the NetApp Infrastructure and Cloud Engineering team. He is focused on the development, validation and implementation of Cloud Infrastructure solutions that include NetApp products. Arvind has more than 10 years of experience in the IT industry specializing in Data Management, Security and Data Center technologies. Arvind holds a bachelor’s degree in Electronics and Communication.
· Ramesh Isaac, Technical Marketing Engineer, Cisco Systems, Inc.
· Aaron Kirk, Product Manager, NetApp
