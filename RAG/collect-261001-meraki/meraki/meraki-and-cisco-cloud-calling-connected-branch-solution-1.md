---
id: collect-261001-meraki/meraki/meraki-and-cisco-cloud-calling-connected-branch-solution-1
title: "Appendix"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-meraki/meraki-and-cisco-cloud-calling-connected-branch-solution.md
source_anchor: ""
source_lines: [1, 111]
sha256: 3125a309a4761bb41f5abcade5d32b6534992af9821355d231a0a6a35b4b85bd
---

# Appendix

## Introduction

In today’s digital transformation age, businesses are looking for collaboration, networking, security, and connectivity services that meet their needs. They are often challenged with a variety of complex solutions and providers to choose from. Managed Service Providers (MSPs) are in a unique position to deliver comprehensive solutions with managed networking and collaboration satisfying business needs and overcoming obstacles with over-the-top (OTT) services.

The purpose of this solution guide is to provide best practices on how MSPs can leverage Meraki’s cloud networking architecture and how Meraki networks should be configured and deployed, taking full advantage of the benefits offered by Cisco’s Cloud Calling platforms. This recommended network configuration is based on a Cisco-tested and -validated architecture optimized for the various Cloud Calling platforms such as Webex Calling and Cisco BroadWorks. The solution utilizes Cisco IP Phones with Multiplatform Firmware (MPP) as endpoints, and it focuses on branch deployments.

Recommendations provided in this document can serve as the basis to create a comprehensive solution that allows for rapid and seamless branch deployment, service continuity by leveraging SD-WAN technologies, as well as ease of management and troubleshooting. The solution enhances the user experience for end customers and providers by bringing together the best-of-breed in cloud-managed networking and collaboration.

The guide walks through some of the technical aspects of the solution and goes over the process of creating a network blueprint optimized for Cloud Calling platforms that can be used for new service creation allowing for consistency and standardization across end-customer sites. Guidance on how to operationalize the use of the blueprint network at scale is also provided in this document.

## Key Benefits

There are several benefits that MSPs can gain from leveraging Meraki’s cloud networking architecture in conjunction with Cisco’s Cloud Calling platforms for branch deployments. Below is a list of the key benefits of this solution.

- 
    Mitigate over-the-top (OTT) related challenges
- 
    Eliminate manual phone provisioning processes
- 
    Based on Cisco’s blueprint network which is optimized for Cisco Cloud Calling platforms
- 
    Rapid and consistent site deployment
- 
    Intuitive and centralized cloud-based management
- 
    Additional visibility of endpoints, networks, and traffic usage for troubleshooting
- 
    Network intelligence and analytics
- 
    WAN monitoring and alerting for proactive response
- 
    Failover capabilities for service continuity
- 
    Built-in security and SD-WAN capabilities

## Prerequisites

In order to deploy this solution, the following is required. Readers of this document should have:

- 
    **Prior experience working the Cisco Meraki dashboard and Cisco Cloud Calling platforms,** in addition to being familiar with the concepts and terminology of local area networking, SD-WAN, wireless, switching, and VoIP
- 
    **Readers should have access to the Cisco Meraki dashboard;** instructions on how to create a dashboard account and organization can be found in the__Creating a Dashboard Account and Organization article__
- 
    **Access to the management portal for the Cisco Cloud Calling platform** being used: Cisco Webex Calling or Cisco BroadWorks
- 
    **Cisco Meraki devices along with their respective license keys;** refer to table 1 (Appendix) for the specific device models, technical specifications, and sizing guide
- 
    **Cisco Multi-Platform Phones (MPP)** . Technical specifications for the different models are available in the__Cisco IP Phones with Multiplatform Firmware (MPP) article__
- 
    **Understanding of recommended dashboard structures for service providers** as outlined in the__Best Practices for Service Providers document__

## Overview of Deployment Architecture

This section provides an overview on how the solution is to be deployed by MSPs. It includes recommendations on how to structure the dashboard organizations, details on the network infrastructure to be installed at customer premises, as well as general guidelines on the customer provisioning process.

### Recommended Dashboard Structure

MSPs generally have multiple customers in the same dashboard organization when delivering a standard service. Utilizing the standard service model (organization per service, network per customer) provides several operational benefits, however this structure should not be used if SD-WAN is being utilized as part of the service. This is because the scope of an organization defines the connectivity domain for Meraki Auto VPN, and it is important to keep customer deployments independent from each other in terms of connectivity. Generally, organizations for MSPs are separated in a one-organization-per-SD-WAN manner.

Given that this solution leverages Meraki Auto VPN and SD-WAN, each customer will need to be assigned to its dedicated organization. The recommended structure for this solution specifies the use of a base organization containing a blueprint network. This organization will serve as the foundation when creating independent organizations for new customer deployments, and each site is represented as a network under that customer’s organization. The structure is illustrated below.

### Branch Network Infrastructure

The network equipment to be installed at customer premises will vary depending on the connectivity needs and amount of clients at each location. A general network diagram for **a single customer deployment with two branch locations** is available below. **This can be scaled to multiple customer organizations** each containing numerous branch networks.



Specific device model information, technical specifications, and a sizing guide is available in **table 1 of the Appendix**.

### Initial Setup and General Provisioning Workflow

In order to deploy this solution using the best practices provided in this document, the following workflows must be performed in order by the administrator.

The **initial setup** consists of... 

1. Creating the base organization for the service
2. Configuring the general settings for the organization
3. Creating the blueprint network



**After initial setup**, when the base organization containing the blueprint network is available, administrators can follow the below workflow for provisioning purposes.

1. Create a new organization by cloning the base organization
2. Create a new network by cloning the blueprint network
3. Add devices to the newly created network



A more detailed provisioning workflow is available in the “Automation of Network Provisioning” section of this document.

## Creating and Configuring the Base Organization

A new organization can be created by following the instructions available in the __Creating a Dashboard Account and Organization article__,

Once the organization is created, the next step is to configure the following from the **Organization > Settings** page on the Meraki dashboard:

- 
    Enable SAML SSO (optional); additional information on how to configure this is available in the __Configuring SAML Single Sign-on for Dashboard article__
- 
    Enable dashboard API access; more information on the dashboard API is available in the __The Cisco Meraki API document__
- 
    Add relevant administrator accounts that will manage customer organizations; the __Managing Dashboard Administrators and Permissions article__ describes the process of managing administrators


**Note:** It is important that all org settings that will be common to all customer organizations are configured on the base organization. The process of creating new customer organizations by cloning from the base org results in a child organization with all settings preconfigured.


## Creating the Blueprint Network

