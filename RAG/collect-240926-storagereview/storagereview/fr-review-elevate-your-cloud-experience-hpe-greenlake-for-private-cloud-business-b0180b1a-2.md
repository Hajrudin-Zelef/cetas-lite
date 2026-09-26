---
id: collect-240926-storagereview/storagereview/fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a-2
title: "fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a"
domain: storagereview
role: reference
task: reference
actors: ["AWS"]
dates: ["2020-06"]
keywords: ["aws", "compute", "disaggregated", "sovereignty"]
source: docs/RAG/clean_en/storagereview/fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a.md
source_anchor: ""
source_lines: [3, 43]
sha256: c9f70e88c2fedbc1fbbfab148c5a2b7759f715d274758665de6026b7a5736961
---

# fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a

In June 2020, Hewlett Packard Enterprise announced the launch of HPE GreenLake Cloud Services to meet the needs of hybrid and multicloud businesses. This year, the offering was further enriched with HPE GreenLake for Private Cloud Business Edition. Offering even greater flexibility and options, HPE GreenLake for Private Cloud Business Edition enables businesses to benefit from a highly personalized and efficient experience to choose the infrastructure and capacity best suited to their needs.
This unique experience and offering make HPE GreenLake more attractive as a hybrid cloud infrastructure option than the more ubiquitous cloud offerings, such as Amazon Web Services and Azure. It is specially designed for businesses seeking the agility of cloud consumption models combined with control of on-premises infrastructure. It offers simplified deployment with predefined configurations, enabling faster configuration and reduced complexity. Thanks to its optimization for specific workloads, businesses can ensure performance efficiency without the overhead associated with managing a vast range of services often present in large public clouds. Additionally, for businesses concerned about data sovereignty, regulatory compliance, or specific security measures, the Business Edition offers the advantage of data localization and increased control over their environment.
In short, HPE GreenLake for Private Cloud Business Edition is a self-managed private cloud solution with a unified interface to simplify management from VMs to infrastructure. It allows you to create your self-service cloud on demand where you need it, with a choice of predictable monthly billing or upfront payment.
Infrastructure
To get started with HPE GreenLake Private Cloud Business Edition, users must first provision the hardware to create their private cloud. For on-premises sites, HPE primarily uses a cloud-native infrastructure composed of HPE Alletra disaggregated hyperconverged infrastructure (dHCI) for business-critical workloads or HPE SimpliVity (true HCI) for distributed edge sites. It provides businesses with a complete private cloud package, true to the concept of HPE GreenLake flexibility for organizations with variable workloads.
HPE has partnered with Amazon Web Services to complement its on-premises offering by using its EC2 instances as a public cloud segment. The combination of these two solutions allows HPE GreenLake to present a unified hybrid cloud to businesses seeking the solution suited to their needs.
Once you have provisioned the hardware and prepared the cloud configuration via Amazon Web Services, you can begin provisioning virtual machines, policies, backups, health checks, and any other myriad of services offered by HPE GreenLake.
HPE SimpliVity
For a quick reminder about HPE SimpliVity and an update on its capabilities, it is a true hyperconverged infrastructure (HCI) appliance that combines all the core features and services necessary for server functionality (such as compute, networking, and storage) into a single large box. It significantly improves efficiency and data security as well as built-in resilience.
Previously, it was accessible and managed via a centralized portal. With the introduction of HPE GreenLake Private Cloud Business Enterprise, HPE SimpliVity can be managed within the HPE GreenLake platform.
In HPE GreenLake, the accessibility of an HPE SimpliVity system compared to an HPE Alletra system is displayed in a common view. It is located in the Systems portal and is only delineated by the system type.
VM Provisioning Policies
VM provisioning policies are essential to the success of an overall hybrid deployment strategy. To access the VM Provisioning Policy portal, select the horizontal bar menu button at the top left to open the list of portal access options. From there, choose VM Provisioning Policies, which will open the main page of all policies currently available in your environment.
To create a new VM provisioning profile, click the "+" icon at the top left and a window will open allowing you to begin filling in the required information. You will need to name the profile and provide a basic description. Select the appropriate radio button for Deduplication, Data Encryption, or All-Flash if these are the options you need. By hovering over the information icon to the right of each option, you get additional information about the use of these options. Finally, you can leave the default QoS performance settings, i.e., the limit of one million IOPS, or adjust them based on the environment.
VM Monitoring
Once you have created a provisioning policy, return to the Virtual Machines page to create the desired instances.
In the top menu, you can view your private cloud (physical) VMs or public cloud (AWS) VMs.
The main page gives you a brief overview of the virtual machines in your environment and other fundamental data about them, including details of their status and data protection. You can sort them by any of the listed characteristics and even have the ability to view their current performance and usage statistics from the buttons on the right side.
The Public Cloud view is no different, except that it does not have as much information to display due to the lack of physical hardware required for monitoring.
Creating a Virtual Machine – Private Cloud
To create a new virtual machine, click the (+) icon in the top left corner to bring the creation window to the foreground.
The first section contains general information: virtual machine names, number of virtual machines, and whether you want them to be powered on after creation.
In the next section, select the Hypervisor cluster where these VMs will be created. You can choose HPE SimpliVity or HPE Alletra dHCI clusters.
In section 3, choose the target datastore.
Next, select the desired operating system template.
Finally, the VM provisioning policy specifies the enterprise-class data protection policy and leverages HPE GreenLake for backup and restore.
Once all the different options and configurations are selected, click Create to submit the VM for creation to your infrastructure stack.
You can also check the VM provisioning status at the bottom, which will also bring up a window to view in-progress and pending tasks as well as the status and a brief basic log output of the process.
Creating a Virtual Machine – Public Cloud
Creating a VM in the Public Cloud is as simple as creating one in the Private Cloud. Click the (+) icon to display the creation options window. On the first page are the initial VM options:
- Virtual machine name
- Service provider (for this review, we tested AWS; Azure is also available)
- Account nickname – The AWS account to which this will be attached
- Region
- Key pair name – AWS will associate this name with the generated public key
After entering the basic information, the next screen will present over 2,200 images to select for your virtual machine instance.
The last screen consists of choosing the instance type. The options listed are the size, the architecture, and whether it is available in the AWS free tier.
As with the private cloud, during VM provisioning, you can view the status and any essential log messages generated during the process.
System Updates
The system update feature allows administrators to update, without interruption, the versions of HPE Alletra Storage Array, ESXi, and server firmware for an entire cluster. These updates are executed in the Systems section of the management console.
In this section of the management console, click the Software button in the right corner of the Systems list to display the software catalogs currently installed in the stack and the catalog available to update in the stack.
