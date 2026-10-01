---
id: collect-261001-meraki/meraki/meraki-and-cisco-collaboration-teleworker-solution-for-businesses-1
title: "meraki-and-cisco-collaboration-teleworker-solution-for-businesses"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "license"]
source: docs/RAG/collect-261001-meraki/meraki-and-cisco-collaboration-teleworker-solution-for-businesses.md
source_anchor: ""
source_lines: [1, 132]
sha256: 056df58b789ceb781f82492736bc359603802ff809a787faf63d250daace1c97
---

# meraki-and-cisco-collaboration-teleworker-solution-for-businesses

## Introduction

In today’s digital transformation age businesses are looking for collaboration, networking, security and connectivity services that meet their needs. They are often challenged with a variety of complex solutions and providers to choose from.


The purpose of this solution guide is to provide best practices on how businesses can leverage Cisco Meraki’s cloud networking architecture and how Cisco Meraki networks should be configured and deployed taking full advantage of the benefits offered by Cisco’s on-premises and Cloud Collaboration platforms. This recommended network configuration is based on a Cisco tested and validated architecture optimized for the various Collaboration platforms such as Webex Meetings, Unified Communication Manager, Webex Calling, and Cisco BroadWorks. The solution utilizes the Webex Meetings client, the Webex Teams client, and/or a Cisco phone and it focuses on teleworker deployments.


Recommendations provided in this document can serve as the basis to create a comprehensive solution that allows for rapid and seamless Teleworker deployment, ease of management and troubleshooting, AutoVPN technology, and optionally service continuity with 4G LTE failover. The solution enhances the user experience for end-customers and providers by bringing together the best of breed in cloud-managed networking and collaboration.


The guide walks through some of the technical aspects of the solution and goes over the process of creating a network blueprint optimized for Cisco Collaboration platforms that can be used for new service creation allowing for consistency and standardization across Teleworker locations. Guidance on how to operationalize the use of the blueprint network at scale is also provided in this document.

## Key Benefits


There are several benefits that businesses can gain from leveraging Cisco Meraki’s cloud networking architecture in conjunction with Cisco’s Collaboration platforms for branch deployments. Below is a list of the key benefits of this solution.

- 
    Mitigate over-the-top (OTT) related challenges
- 
    Eliminate manual phone provisioning processes
- 
    Based on Cisco’s blueprint network optimized for Cisco Collaboration platforms
- 
    Rapid and consistent Teleworker deployment
- 
    Intuitive and centralized cloud-based management
- 
    Additional visibility of endpoints, networks and traffic usage for troubleshooting
- 
    Network intelligence and analytics
- 
    WAN monitoring and alerting for proactive response
- 
    Failover capabilities for service continuity

## Prerequisites


In order to deploy this solution, the following is required. Readers of this document should have...

- 
    **Prior experience working the Cisco Meraki dashboard and Cisco Collaboration platforms** in addition to being familiar with the concepts and terminology of local area networking, VPN and VoIP.
- 
    **Readers should have access to the Cisco Meraki dashboard.** Instructions on how to create a dashboard account and organization can be found in the__Creating a Dashboard Account and Organization article__ .
- 
    **Access to the management portal for the Cisco Collaboration platform** being used. Cisco Webex, Cisco Unified Communications Manager, Webex Calling or Cisco BroadWorks.
- 
    **Cisco Meraki devices along with their respective license keys.** Refer to**Table 1 of the Appendix** for the specific device models, technical specifications and sizing guide.

## Hardware Used and Options

There are many different designs and architectures that can be utilized in these teleworker solutions. Below is a list of hardware

- 
    **Security Gateway**
  - 
        Meraki MX 64/65W models (teleworker)
  - 
        Meraki MX 67/68W models (teleworker)
  - 
        Meraki Z1/Z3/Z4(c) models (teleworker)
  - 
        Meraki MX 84, 100, 250, 450 (VPN Hub)
- 
        
- 
    **Wireless Access Points**
  - 
        Meraki MR access points (teleworker)
- 
        
- 
    When we are referring to the configuration as Teleworker Gateway or SD-WAN & Security in the side menu in dashboard, these two can be used interchangeably.
- 
    When using a template “Security & SD-WAN” is what will be in the menu. These templates can be applied to Meraki WAN Appliances and Meraki Teleworker gateways with the same template.
- 
    When using the “Blueprint Network” or configuring networks on network by network basis, the Meraki WAN Appliance will show “Security & SD-WAN”. The Meraki Teleworker Gateway will show “Teleworker Gateway”

## **Overview of Deployment Architecture**

This section provides an overview on how the solution is to be deployed by businesses. It assumes that you already have a Meraki AutoVPN network configured and operational and provides details on the network infrastructure to be installed at Teleworker premises as well as general guidelines on the provisioning process.

### Teleworker Infrastructure

The network equipment to be installed at the Teleworker premises will consist of a Meraki Z3 Teleworker gateway and a Cisco Desk phone or a Cisco Headset to enable the remote worker to communicate effectively.



Specific device model information, technical specifications and a sizing guide is available in **Table 1 of the Appendix**.

### Setup and General Provisioning Options


There are two options to consider when deploying this solution.

1. 
    Creating a network template that is applied to all teleworker networks that are alike. This provides 100% consistency in deployment. Since this method is using a template, most configurations will be performed in the template and then are distributed to the individual networks upon saving of the template changes. This provides the greatest degree of consistency and single point of management. The trade-off for a template network deployment is that many local network configuration changes cannot be performed for one off use cases. Choose this option if you prefer less management overhead at the cost of individual Teleworker site customization.


1. 
    Creating a Blueprint Network that is then cloned for the creation of all other networks. This provides consistent network creation for deployment. Since this method is not using a template, following initial deployment, configuration changes to all teleworker networks will need to be made on an individual network by network basis. Configuration change overhead can be greatly reduced by using the API to configure all site changes simultaneously. Choose this option if you prefer a high degree of Teleworker site customization at the cost of more management overhead.

## **Setup and General Provisioning - Template Network Option** 

In order to deploy this solution using the best practices provided in this document, the following workflows must be performed in order by the administrator. The initial setup consists of...

1. 
    Configuring the general settings for the organization
2. 
    Creating the Template Network



The next sections walk through the process of configuring the base organization as well as the template network.

### Creating and Applying Configuration Templates

This section walks through the process of creating the template network that will serve as the base configuration for the network infrastructure. This will be used as a starting point for new customer deployments as well as the ongoing configuration point for all of the teleworker configurations. The template network represents the Cisco recommended settings for successful deployments.

#### New Template Creation

**After** a dashboard account and organization have been created:


**Step 1**. Log in to the Cisco Meraki dashboard as an organization administrator


**Step 2**. Select **Configuration templates** from the **Organization** tab 


