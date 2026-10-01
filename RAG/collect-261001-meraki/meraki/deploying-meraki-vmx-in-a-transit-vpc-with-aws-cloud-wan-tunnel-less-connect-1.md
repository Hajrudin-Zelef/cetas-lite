---
id: collect-261001-meraki/meraki/deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect-1
title: "deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect"
domain: meraki
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "compute", "latency", "throughput"]
source: docs/RAG/collect-261001-meraki/deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect.md
source_anchor: ""
source_lines: [1, 66]
sha256: 0af670a7c7ea5632321a438f0f7e540c63947f68608e6e97c4cc3b72ac4623a5
---

# deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect

## Overview

This document goes over the step-by-step configuration and provides a reference architecture for deploying the Meraki vMXs with AWS Cloud WAN. It helps organizations to extend their Meraki SD-WAN fabric to the application running on AWS.

## Why AWS Cloud WAN

Customers are increasingly moving towards multi-region deployments in the cloud and require a more dynamic, secure and a reliable way to connect from branch sites to cloud workloads across different regions. While this can be done using VPC peering across different regions and using static routes, this can quickly become cumbersome and difficult to manage.

AWS Cloud WAN provides a central dashboard for making connections between your branch offices, data centers, and Amazon Virtual Private Clouds (VPCs) in just a few clicks. With Cloud WAN, you use network policies to automate network management and security tasks in one location. Cloud WAN generates a complete view of your on-premises and AWS networks to help you monitor network health, security, and performance.

Meraki vMX integrates with AWS Cloud WAN to allow admins to define a multi-region, segmented, dynamically routed global network with intent-driven policies. This allows organizations to scale across different regions without worrying about managing the complexity of peering across different regions.

## Tunnel-Less Connect

Tunnel-less connect provides a simple and high-performance way to build global SD-WANs using the AWS Global network as a middle-mile transport network. Tunnel-less connect allows SD-WAN appliances to peer natively with Cloud WAN using BGP without any sort of tunneling technology like GRE or IPsec. The key benefits of this are:

- 
    **Cloud-enabled SD-WAN:** Simplify branch connectivity between on premises resources and the AWS cloud. With tunnel-less connect, these SD-WAN appliances can be onboarded into AWS with less overhead and higher throughput.
- 
    **AWS Global Network as a middle-mile for inter-office connectivity:** AWS Cloud WAN can provide high throughput and low latency connectivity for geographically disperse sites, using the AWS cloud as a backbone network.
- 
    **Improved throughput performance:** Without GRE or IPsec adding overhead and reducing overall effective throughput, tunnel-less connect provides much higher performance and simplicity, with up to 100Gbps of bandwidth per attachment.

## Cloud WAN Components

The main components of Cloud WAN with Tunnel-less connect are:

- 
    **Global network:** Private network that acts as the high-level container for your network objects. Global networks can contain AWS Transit Gateways and Cloud WAN core networks.
- 
    **Cloud WAN Core Network:** Part of your global network managed by AWS. Includes regional connection points and attachments, such as VPNs, VPCs, and Transit Gateway Connects. Your core network operates in the regions that are defined in your core network policy document.
- 
    **Core Network Edge (CNE):** Regional connection points managed by AWS in each Region, as defined in the core network policy. Under the hood, CNEs are AWS Transit Gateways, and they inherit many of the same properties.
- 
    **Attachments:** Connections or resources that you want to add to your core network. Supported attachments include VPCs, VPNs, Transit Gateway route table attachments and Connect attachments.
- 
    **Core network policy:** Single document applied to your core network that captures your intent and deploys it for you. The policy uses a declarative language defining segments, AWS region routing and attachment to segment mappings.
- 
    **Network segment:** Segments are dedicated routing domains, similar to globally consistent Virtual Routing and Forwarding (VRF) tables. By default, only attachments within the same segment can communicate. Segment actions can be defined to share routes across segments in the core network policy.
- 
    **Connect Attachment (Tunnel-less):** Peering point for your SD-WAN appliances that functions in a tunnel-less manner and uses native BGP to dynamically exchange routing and reachability information between SD-WAN appliances in the VPC and the CNE.
- 
    **VPC Attachments:** VPC attachments act as transport attachments and carry data-plane traffic between SD-WAN appliances and the CNE.
- 
    **Transport VPCs:** These are dedicated VPCs for hosting your SD-WAN appliances, in this case Cisco Meraki virtual MXs (vMX). These VPCs hold only the necessary resources to provide connectivity to your vMXs, and do not host any additional workloads or services. One Transport VPC is recommended for every geographic area where you have SD-WAN services that need connectivity to your AWS resources. It is also recommended to host your vMXs in pairs, in separate Availability Zones (AZs) for maximum availability.
- 
    **Workload VPCs:** Workload VPCs host any of the compute applications in your AWS environment. In practice, it is common to have multiple types of workload VPCs, such as Production, Development and Testing, but in this guide we will just deploy Production VPCs (the same steps can be repeated for any number of additional workload types). For high availability of workloads, it is common to have dedicated Workload VPCs in every region where you need AWS services.

## Solution Architecture

In this guide we will guide you through the steps to set up an environment with a core network interconnecting two separate AWS regions, each with their own Transit and Workload VPCs. These Transit and Workload VPCs will be assigned to two separate segments, SD-WAN and Production, and we will enable segment sharing between these for end-to-end connectivity.

The sample topology is as below:

For every region in your deployment, you will want to have a dedicated SD-WAN or Transit VPC to host your vMXs. In this guide, we will deploy these VPCs in the us-east-1 and us-west-1 regions. In each of these VPCs, you will want to deploy a pair of subnets, each in a separate Availability Zone (AZ) within the region for maximum availability. Each of these subnets will host one vMX only, so /28 addressing is fine for these (5 addresses of every subnet are reserved for AWS usage). All these regional SD-WAN VPCs will be associated with the SD-WAN segment in Cloud WAN.

You can deploy any number of workload VPCs in each of these regions. If your environment has multiple workload environments (Prod, Dev, Test), it is recommended to have a dedicated Cloud WAN segment for each. In this case, we will only be using a single Production segment for all the workload VPCs.

A single Cloud WAN core network will be provisioned, with edge locations in us-east-1 and us-west-1. Each of these edge locations will have a Core Network Edge (CNE) associated with it, and the core network edge will have VPC attachments to each of the SD-WAN and Production VPCs in its own region. Additionally, over the SD-WAN VPC attachment, a Connect attachment will be deployed for each of the subnets and AZs in the VPC (two per VPC). These Connect attachments will be provisioned in Tunnel-less Connect or no encapsulation mode, and these will be the endpoints that will peer via EBGP with your vMXs in the VPC.

The Meraki SD-WAN fabric will have a single ASN associated with it, while the Core Network will deploy one ASN for every CNE deployed (two in this case, due to operating in two regions). The BGP ASNs will be assigned as:

We will also enable bidirectional segment sharing between the two segments (SD-WAN and Production), which will allow them to advertise all their prefixes to each other.

By default, Meraki AutoVPN hubs form tunnels to each other. For large scale deployments, it is often recommended to disable hub-to hub-tunnels in the Meraki SD-WAN AutoVPN to achieve routing that is as deterministic as possible. This is because:

