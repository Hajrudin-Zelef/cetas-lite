---
id: collect-261001-huawei/huawei/security-technical-requirements-for-artificial-intelligence-54a86121-1
title: "security-technical-requirements-for-artificial-intelligence--54a86121"
domain: huawei
role: reference
task: reference
actors: ["China", "Google"]
dates: []
keywords: ["agent", "agents", "cybersecurity", "inference", "latency", "mcp", "model context protocol", "reasoning", "tool calling", "training"]
source: docs/RAG/collect-261001-huawei/security-technical-requirements-for-artificial-intelligence--54a86121.md
source_anchor: ""
source_lines: [1, 237]
sha256: bf151445d849925998ea67215331f788f2b6d00e3f4184b7532b7a3bc4cd6d3e
---

# security-technical-requirements-for-artificial-intelligence--54a86121

 ICS 35.030
 CCS L 80


                                                        CCIA
China Cybersecurity Industry Alliance Technical Specification
                                                             T/CCIA XXX-2026




              人工智能防火墙安全技术要求
 Security technical requirements for artificial intelligence firewall




Released on XXXX-XX-XX                           Effective since XXXX-XX-XX

              Released by China Cybersecurity Industry Alliance
             Security technical requirements for
                     artificial intelligence firewall

Introduction
As AI technologies continue to transform networks and information systems, security threats
targeting vulnerabilities specific to AI systems are becoming increasingly severe. Traditional
threats can also be enhanced and amplified by attackers using AI-generated tools.

As of 2026, general-purpose firewalls are inadequate for protecting AI systems. There is an
urgent need for firewalls specifically designed to provide security protection for AI systems,
while also incorporating AI-enhanced security capabilities.

Relevant industry demand and preliminary analyses for such AI firewalls have already
emerged around the world:

    •   According to Gartner, firewalls have evolved through several successive generations,
        including packet filtering firewalls, stateful inspection firewalls, next-generation firewalls
        (NGFWs), unified threat management (UTM) systems, and hybrid mesh firewalls (HMFs).
        More recently, industry demand analysis related to AI-Gateways have also emerged.

    •   Network security vendors such as Akamai, Fortinet, and Palo Alto Networks have also
        introduced related offerings, including AI WAF, FortiAI-Protect, and AI-Powered NGFW.

This document specifies additional security defense technical requirements for artificial
intelligence firewalls, on top of the Chinese National Firewall Standard GB/T 20281—2020,
Information Security Technology—Security Technical Requirements and Testing Assessment
Approaches for Firewall.



1. Scope
This document specifies the technical requirements for AI firewalls.

This document is applicable to the design, development, testing and deployment of AI
firewalls.



2. Normative references
The contents of the following documents constitute indispensable provisions of this
document through normative references in the text. For dated references, only the edition
corresponding to the cited date applies to this document. For undated references, the latest
edition of the referenced document (including any amendments) applies to this document.
GB/Z 185.1-2026 Artificial Intelligence—Agent Interconnection—Part 1:General Architecture

GB/T 25069-2022 Information Security Techniques—Terminology

GB/T 20281-2020 Information Security Technology—Security Technical Requirements and
Testing Assessment Approaches for Firewall



3. Terminology
The terms and definitions specified in GB/T 25069—2022, GB/T 20281-2020 and those given
below apply to this document.


3.1.     artificial intelligence firewall
A network security product that parses data flows associated with artificial intelligence systems,
implements access control and security protection functions, and provides artificial
intelligence-enhanced security protection capabilities.


3.2.     token
The smallest unit of measurement for information processed by a model.


3.3.     tool
A device, software, or system that provides specific functionality and can be used.


3.4.     skills
A modular capability extension encapsulated in a structured manner.

NOTE: Typically used in the fields of AI and AI agents.


3.5.     tool calling
The behavior of an AI model or agent invoking external tools, Application Programming
Interfaces (APIs), and services.


3.6.     caller
An access subject that invokes resources, including human users, agents, business
applications, and other entities.
3.7.     callee
Resources being invoked, including LLM services, agents, tools, APIs, and other resources.


3.8.     AI agent
An entity capable of perceiving its environment, making autonomous decisions, and
proactively taking actions to achieve specific goals.



4. Abbreviation
    The following abbreviations apply to this document:

    A2A: Agent to Agent Protocol

    ACP: Agent Communication Protocol

    AI: Artificial Intelligence

    ANP: Agent Network Protocol

    API: Application Programming Interface

    EDR: Endpoint Detection and Response

    gRPC: Google Remote Procedure Call

    HTTP: HyperText Transfer Protocol

    HTTPS: HyperText Transfer Protocol Secure

    IAM: Identity and Access Management

    IP: Internet Protocol

    IPS: Intrusion Prevention System

    IT: Information Technology

    LLM: Large Language Model

    MCP: Model Context Protocol

    NHI: Non-Human Identity

    OAuth 2.0: Open Authorization 2.0

    RAG: Retrieval Augmented Generation

    SLIM: Secure Low-Latency Interactive Messaging

    SSL: Secure Socket Layer
    SSRF: Server-Side Request Forgery

    SQL: Structured Query Language

    TLS: Transport Layer Security



5. Overview
An AI firewall is a new type of boundary security protection product that, on the basis of a
general-purpose firewall, can effectively address security risks specific to AI systems and
provides unified security management and control capabilities. The differences in
characteristics between AI firewalls and general-purpose firewalls include, but are not limited
to, the following:

——identification, classification, and control of non-human identities (NHIs), such as AI
agents;

——multidimensional perception of underlying security attack intents and semantics;

——AI-enhanced capabilities for identifying rapidly evolving variants and unknown threats;

——autonomous security reasoning, intelligent learning and generation of security policies,
automated policy upload and deployment, coordinated blocking;

——monitoring and observability capabilities for intelligent computing resource invocation,
such as AI agents and models and the consumption of computing power.

The security technical requirements for AI firewalls specified in this document refer to the
specific security requirements that products must additionally meet on the basis of satisfying
the relevant enhanced-level requirements specified in GB/T 20281-2020.



6. Technical Requirements

6.1.     Security functional requirements

6.1.1. Networking and deployment

6.1.1.1. Deployment mode
The product must meet the relevant requirements specified in 6.1.1.1 of GB/T 20281—2020.


6.1.1.2. Deployment location
The product must support the following deployment locations:
    1)   between callers (including users, AI agents, or business applications) and callees
         (including AI inference servers, public LLM service APIs, AI agents, tools, APIs, etc.),
         including:

         — in-domain callers and out-of-domain callees;

         — in-domain callers and in-domain callees.

    2)   between the network boundaries of internal zones within an AI information system,
         such as training zones, inference zones, and storage zones.


6.1.1.3. Deployment scenarios
The product must support the following deployment scenarios:

    1)   scenarios involving Internet egress of cloud services and internal network zone
         boundaries;

    2)   scenarios involving internal network zone boundaries in data centers (including
         intelligent computing, general-purpose computing, management, and other zones);

    3)   scenarios involving enterprise campus network boundaries and internal network
         zone boundaries.


6.1.1.4. Routing
In addition to meeting the relevant requirements specified in 6.1.1.2 of GB/T 20281—2020,
the product must also provide AI-oriented routing capabilities:

