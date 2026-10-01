---
id: collect-261001-huawei/huawei/security-technical-requirements-for-artificial-intelligence-54a86121-8
title: "security-technical-requirements-for-artificial-intelligence--54a86121"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "distribution", "reasoning"]
source: docs/RAG/collect-261001-huawei/security-technical-requirements-for-artificial-intelligence--54a86121.md
source_anchor: ""
source_lines: [1211, 1265]
sha256: 317a6cb87e5ca4cb43743f1048e287507c16f04be6cfaf43102d7c30a7b981b7
---

# security-technical-requirements-for-artificial-intelligence--54a86121

6.3.3. Centralized policy management and control
    1)   The product must support centralized orchestration and batch distribution of AI-
         specific security policies by the management platform, thereby uniformly
         implementing AI security protection capabilities such as prompt injection protection,
         sensitive content filtering, model access control, and rate limiting and blocking of
         abnormal traffic;

    2)   The product must support policy version management, conflict detection, staged
         rollout, and global withdrawal capabilities to ensure the consistency and
         effectiveness of protection policies for multi-node AI firewalls across the network.


6.3.4. Collaborative security response capabilities
    1)   The product must support receiving dynamic security response instructions issued
         by the security management platform and enable coordinated response across all
         domains to perform real-time blocking of malicious IP addresses, risky accounts,
         and anomalous model invocation behavior, thereby establishing a closed-loop
         coordinated security response system.

    2)   The product must support the use of threat information received from endpoint-
         side threat detection and response products, other firewalls, and other sources to
         perform collaborative security event analysis, reasoning, and protection.


6.4.     Performance requirements
The relevant requirements specified in 6.4 of GB/T 20281—2020 should apply.


6.5.     Security Assurance requirements
The relevant requirements specified in 6.5 of GB/T 20281—2020 should apply.
Appendix: AI Firewall Reference Deployment Architecture Examples

Deployment scenarios for AI firewalls include two typical use cases: data centers (including
general data centers, AI data centers, etc.) and campus branch interconnections. The
specific networking topology is shown in the figure below.

As shown from the typical networking topology, with the widespread deployment of AI
systems, typical east-west traffic interactions have emerged within data centers and
between campus branches, in addition to traditional north-south traffic. By deploying AI-
agent firewalls at the boundaries of network zones within data centers and at the
boundaries of enterprise campus branches, organizations can effectively isolate not only
north-south security risks but also east-west security risks (e.g., when an intelligent agent
calls resources within a domain), thereby enhancing the overall security posture.




        Figure 1. Reference Deployment Architecture of AI Firewalls in Data Centers




 Figure 2. Reference Deployment Architecture of AI Firewalls for Campus Branch Networks

