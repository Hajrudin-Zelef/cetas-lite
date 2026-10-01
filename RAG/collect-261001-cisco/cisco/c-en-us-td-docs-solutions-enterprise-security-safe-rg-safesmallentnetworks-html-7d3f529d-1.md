---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-1
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "alignment", "cost", "cyber"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [1, 43]
sha256: 243bff2d81b6345f6ae8c46d9a28504d8164bc421f1345b2339b74481636854f
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

Table Of Contents
Cisco SAFE for Small Enterprise Networks
Cisco SAFE Security Reference Architecture
Cisco Security Control Framework (SCF)
Cisco SAFE Architecture Principles
Small Enterprise Network Security Design
Internet Border Router Security Guidelines
Intrusion Prevention Guidelines
Network Access Security and Control
Catalyst Integrated Security Features
Cisco Unified Wireless Network (CUWN) Integrated Security Features
NAC Appliance Modes and Positioning
NAC Appliance Deployment in Small Enterprise Networks
Cisco Identity-Based Network Networking Services (IBNS)
Impacts of 802.1X on the Network
802.1X in Small Enterprise Networks
NAC 802.1X and CISF in Combination
Internet Border Router Deployment
Firewall Hardening and Monitoring
Network Address Translation (NAT)
Intrusion Prevention Deployment
Deploying IPS with the Cisco ASA
IPS Global Correlation Deployment
Basic Clean Access Switch Configuration
Basic Clean Access Out of Band Switch Configuration
Basic 802.1X Switch Configuration
Cisco SAFE for Small Enterprise Networks
Executive Summary
If there is something certain about Internet security, it is the fact that cyber crime and on-line threats affect all organizations, no matter how small or large. Internet-based organized crime and espionage, identity and data theft, botnet infections, and insider attacks are common threats affecting all types of businesses. Particularly attractive to small businesses, mobile access technologies and cloud-based services deliver great flexibility and cost-savings, but pose new challenges. Understanding the nature and diversity of the threats affecting the small business, and how they may evolve over time, is the first step towards a successful security strategy. While small businesses tend to have fewer locations and employees to protect, tighter budgets and limited resources require them to take an innovative and cost-effective approach to security. The security strategy should also help the small business achieve and maintain compliance with mandated standards and regulations.
This document explains how the proven design and implementation principles of the Cisco SAFE Reference Architecture help secure the small business by building a solid and reliable network infrastructure that is resilient to both well-known and new forms of attacks. The Cisco SAFE is a security reference architecture that provides detailed design and implementation guidelines for organizations looking to build highly-secure and reliable networks. The Cisco SAFE leverages Cisco's many years of design and deployment experience and is an architecture that is thoroughly tested and validated as part of the Cisco Validated Design (CVD) program. This document discusses the Cisco SAFE best practices and guidelines that ensure the confidentiality, integrity, and availability of data and system resources supporting key business functions. Design recommendations are based on an understanding of current and future needs and consider the technical and financial constraints often faced by small businesses.
The objective of this document is to present the Cisco SAFE best practices, designs, and configurations applicable to the small business and to provide network and security engineers with the necessary information to help them succeed in designing, implementing, and operating secure network infrastructures based on Cisco products and technologies. While the target audience is technical in nature, business decision makers, senior IT leaders, and systems architects can benefit from understanding the design principles and fundamental security concepts.
Cisco SAFE Security Reference Architecture
The Cisco SAFE consists of design blueprints based on CVDs and proven security best practices that provide the design guidelines for building secure and reliable network infrastructures. The Cisco SAFE design blueprints implement defense-in-depth by strategically positioning Cisco products and capabilities across the network and by leveraging cross-platform network intelligence and collaboration. To that end, multiple layers of security controls are implemented throughout the network, but under a common strategy and administration. The Cisco SAFE uses modular designs that accelerate deployment and that facilitate the implementation of new solutions and technologies as business needs evolve. This modularity extends the useful life of existing equipment, protecting capital investments. At the same time, the designs incorporate a set of tools to facilitate day-to-day operations, reducing overall operational expenditures.
The Cisco SAFE uses the Cisco Security Control Framework (SCF), a common framework that drives the selection of products and features that maximize visibility and control, the two most fundamental aspects driving security. Also used by Cisco's Continuous Improvement Lifecycle services, the framework facilitates the integration of Cisco's rich portfolio of security services designed to support the entire solution lifecycle.
Cisco Security Control Framework (SCF)
The Cisco SCF is a security framework aimed at ensuring network and service availability and business continuity. Security threats are an ever-moving target and the SCF is designed to address current threat vectors, as well as track new and evolving threats, through the use of best practices and comprehensive solutions. Cisco SAFE uses SCF to create network designs that ensure network and service availability and business continuity. Cisco SCF drives the selection of the security products and capabilities and guides their deployment throughout the network where they best enhance visibility and control.
SCF assumes the existence of security policies developed as a result of threat and risk assessments and in alignment to business goals and objectives. The security policies and guidelines are expected to define the acceptable and secure use of each service, device, and system in the environment. The security policies should also determine the processes and procedures needed to achieve the business goals and objectives. The collection of processes and procedures define security operations. It is crucial to business success that security policies, guidelines, and operations do not prevent, but rather empower the organization to achieve its goals and objectives.
The success of the security policies ultimately depends on the degree to which they enhance visibility and control. Simply put, security can be defined as a function of visibility and control. Without any visibility, there is no control, and without any control, there is no security. Therefore, SCF's main focus is on enhancing visibility and control. In the context of SAFE, SCF drives the selection and deployment of platforms and capabilities to achieve a desirable degree of visibility and control.
SCF defines six security actions that help enforce the security policies and improve visibility and control. Visibility is enhanced through the actions of identify, monitor, and correlate. Control is improved through the actions of harden, isolate, and enforce.
Figure 1 Cisco Security Control Framework Model
In an enterprise, there are various places in the network, such as data center, campus, and branch. The SAFE designs are derived from the application of SCF to each place in the network. The result is the identification of technologies and best practices that best satisfy each of the six key actions for visibility and control. In this way, SAFE designs incorporate a variety of technologies and capabilities throughout the network to gain visibility into network activity, enforce network policy, and address anomalous traffic. As a result, network infrastructure elements such as routers and switches are used as pervasive, proactive policy-monitoring and enforcement agents.
Cisco SAFE Architecture Principles
The Cisco SAFE design blueprints were created following these architecture principles:
