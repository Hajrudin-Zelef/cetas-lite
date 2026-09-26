---
id: collect-240926-storagereview/storagereview/fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a-3
title: "fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a"
domain: storagereview
role: reference
task: reference
actors: ["AWS", "Microsoft"]
dates: []
keywords: ["aws"]
source: docs/RAG/clean_en/storagereview/fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a.md
source_anchor: ""
source_lines: [44, 87]
sha256: da37f13302ab251d61d3a9e8725aabdd4ba3d6d9100a2639b59b90282c0a3977
---

# fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a

A catalog contains specific versions of Array OS, ESXi, HPE Storage Connection Manager, and Service Pack for ProLiant. The individual versions of the catalog are visible by hovering over the catalog version number displayed in the row of the stack you are examining.
To begin the upgrade process, select the cluster you want to upgrade and click Run Precheck, or browse this system and select Run Precheck from the actions menu. Then, select the destination catalog version to which you want to upgrade the cluster stack.
Once the upgrade has begun, it will be subjected to a comprehensive set of prechecks to ensure it is prepared and will not cause damage or data loss on the system.
Once completed, you will receive a notification and the main software view will also update its appearance to reflect the changes.
Health Checks
Maintaining and confirming the health of physical hardware is an imperative responsibility of every system administrator. The Health Checks application provides time-saving automation, ensuring the proper functioning of applications and infrastructure and compliance with best practices by the cloud environment. To validate this, HPE GreenLake tasks the system with running checks against approximately 75 rules or configuration guidelines to ensure the infrastructure is operating under optimal conditions.
As a concrete example, in the event of an access problem to a VM datastore, health checks will guide the administrator to the host concerned among all clusters exhibiting this problem. In this specific case, within the HPE GreenLake portal, the most efficient method to run the necessary health checks on Alletra dHCI infrastructure stacks is to access the portal page of the system concerned via the left menu and select "Systems." Once on this page, choose the specific dHCI stack on which you want to perform a health check.
Integrations with Cloud Data Services Console Automation
Understanding how HPE GreenLake for Private Cloud Business Edition integrates into the HPE Data Services architecture is essential. As shown in the diagram below, the Data Services Cloud Console (DSCC) has several layers. Starting from the bottom is the cloud-native data infrastructure layer. This is represented by on-premises clusters running virtual machines such as HPE Alletra dHCI clusters or HPE SimpliVity clusters, including cloud clusters running in one of the supported hyperscalers.
The middle layer consists of Cloud Infrastructure services, with which Private Cloud Business Edition integrates. It also includes services such as the configuration service to automate cluster deployments on sites to save time and Data Ops Manager which can be used to configure blocking services such as replication between clusters or performance analysis to detect storage performance issues.
The top layer consists of Cloud Data Services, with examples such as HPE GreenLake for backup and restore.
Backup and Recovery
Every system engineer knows that backup and recovery are among the most critical areas of administration. HPE understands this as well and dedicates an entire portal to this purpose.
The main dashboard provides an overview of the inventory summary, the number of tasks and protection jobs, warnings and issues, as well as data consumption from backup in the cloud and on-premises.
In the left menu, you have a wide range of options to manage backups and policies for all systems managed in HPE GreenLake, including:
- Protection Policies
- Amazon Web Services
  - Elastic Block Store
  - EC2
  - EKS Clusters
  - Relational Database Service
- Microsoft SQL Servers
  - Databases
  - Instances
  - Protection Groups
  - Application Hosts
- VMware
  - Virtual Machines
  - Datastores
  - Protection Groups
  - vCenter Servers
- HPE Array Volumes
- Reports
- On-Premises Configuration
  - Data Orchestrators
  - StoreOnce Gateways
  - StoreOnce Stores
  - StoreOnce
Final Thoughts
HPE GreenLake for Private Cloud Business Edition, building on HPE GreenLake Cloud Services, offers businesses a tailor-made hybrid cloud solution. Unlike traditional providers, namely Amazon Web Services and Microsoft Azure, HPE GreenLake combines the agility of cloud consumption with control of on-premises infrastructure, emphasizing data localization for businesses concerned about compliance and security. It relies on a cloud-native data infrastructure composed of HPE Alletra dHCI for business-critical workloads or HPE SimpliVity for distributed edge sites. It integrates seamlessly with Amazon's EC2 and Microsoft Azure instances. Users benefit from streamlined configuration, deployment, VM provisioning, comprehensive backup/recovery tools, and efficient system health checks, ensuring a robust and efficient cloud experience.
* HPE GreenLake deployed public cloud access to Azure virtual machines in December. This means you can now create and manage virtual machines in the Microsoft ecosystem in the same way you currently manage them with AWS EC2.
HPE GreenLake for Private Cloud Business Edition
Demo minute 8
Demo minute 3
