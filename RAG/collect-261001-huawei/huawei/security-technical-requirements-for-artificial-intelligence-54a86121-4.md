---
id: collect-261001-huawei/huawei/security-technical-requirements-for-artificial-intelligence-54a86121-4
title: "security-technical-requirements-for-artificial-intelligence--54a86121"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "context window", "distribution", "exploit", "jailbreak", "sandbox", "training"]
source: docs/RAG/collect-261001-huawei/security-technical-requirements-for-artificial-intelligence--54a86121.md
source_anchor: ""
source_lines: [541, 700]
sha256: fab3567174ca33ba1607585716513abd46a994151ff7e85979ef5e7b9c52b52e
---

# security-technical-requirements-for-artificial-intelligence--54a86121

   6)   The product must support the identification and control of attack-related anomalies,
        such as model topic deviation and access beyond authorized privileges by AI agents;

   7)   Before model invocation, user input must undergo anti-virus and intrusion
        prevention inspection, semantic analysis, sensitive instruction identification, and
        content compliance checks;

   8)   During the model output stage, sensitive information filtering and blocking,
        together with compliance checks, must be performed to prevent the model from
        generating content containing account credentials, private data (such as bank card
        numbers, personal IDs, and telephone numbers), commercially sensitive information,
        business secrets, or other sensitive content. Sensitive information outputs must be
        intercepted to prevent data leakage;

   9)   The product must support values-based content filtering, review input and output
        content, detect illegal or non-compliant content including politically sensitive topics,
        discrimination, harmful content, or other prohibited matters, and automatically
        block non-compliant inputs and outputs during interactions with LLM;

   10) The product must support real-time desensitization of content submitted by
       enterprise users to publicly available LLM. Before a request is sent, sensitive data,
       such as names, identity card numbers, customer information, and source code
       snippets, must be automatically identified and desensitized through measures such
       as replacement, masking, or blocking, so as to prevent sensitive data leakage.


6.1.4. Attack protection

6.1.4.1. Network-layer attack protection
   1)   The product must support the detection and blocking of various malformed packet
        attacks, including but not limited to:

        ——Local Area Network Denial (LAND) attacks;

        ——Smurf attacks;

        ——Teardrop attacks;

        ——Ping of Death attacks;

       ——Internet Control Message Protocol (ICMP) redirect or destination unreachable
   message attacks;

       — — invalid Transmission Control Protocol (TCP) packet flag settings, such as
   Acknowledgment (ACK), Synchronize (SYN), and Finish (FIN) flags;

   2)   The product must support protection against network-layer DoS attacks, including
        but not limited to:

        ——SYN Flood attacks;

        ——ACK Flood attacks;

        ——ICMP Flood attacks;

        ——User Datagram Protocol (UDP) Flood attacks;

   3)   The product must support protection against network scanning attacks and must
        detect and record network scanning behavior;
   4)   The product must support the detection and blocking of lateral attacks;

   5)   The product must support the identification and active termination of abnormal
        sessions, for example:

        ——zombie sessions that have no data transmission for longer than the specified
        timeout period;

        ——oscillating sessions for which the number of establishments and terminations
        per unit time exceeds a threshold;

        ——unauthorized sessions with abnormal transport-layer protocol states;

        ——sessions suspected of remote control via covert channels;

   6)   The product must support AI-enhanced detection and blocking of malformed
        network-layer packets, protocol anomalies, and unknown threats.


6.1.4.2. Application-layer attack protection
   1)   The product must support protection against application-layer DoS attacks;

   2)   The product should support semantic consistency validation of metadata across
        protocols to prevent cross-protocol attacks.


6.1.4.2.1. Model attack protection
   1)   The product must support the detection of malicious payloads concealed in model
        files, including but not limited to the following:

        ——identifying model files in network traffic and comparing them based on feature
        values, model fingerprints, or other characteristics;

        ——identifying model files in network traffic and detecting and parsing malicious
        Portable Executable (PE) headers, malicious scripts, and other malicious content
        through static file scanning;

        — — placing models in external isolated sandboxes for execution through
        coordinated sandbox-based isolation, and jointly detecting abnormal outbound
        network connections or API call behavior.

   2)   For potential vulnerabilities in models, the product must support the detection of
        plaintext traffic and SSL-encrypted traffic and detect potential triggers contained in
        prompts and payloads, so as to defend against attacks that exploit vulnerabilities
        through malicious traffic;

   3)   The product must detect and block high-risk operations on model weight files and
        training datasets, such as modification, overwriting, and deletion;

   4)   The product must support protection against model theft and detect potentially
        sensitive model files in traffic;
5)   The product must support the detection and blocking of resource-exhaustion
     attacks against model services;

6)   The product must support the detection and blocking of attacks in which malicious
     file uploads induce models to perform path traversal, command injection, or other
     malicious behavior;

7)   The product should support the detection of and protection against attacks that
     trigger model performance degradation or service disruption by constructing
     oversized Context Window Overflow;

8)   The product must support user-defined intrusion behavior patterns and allow
     corresponding detection and blocking rules to be customized according to user
     requirements;

9)   The product must be capable of blocking direct and indirect prompt injection attacks,
     including but not limited to injected instructions involving jailbreak attacks, role-
     playing, topic deviation, inducive instructions, and other techniques;

10) The device must include a built-in AI attack signature database covering
    vulnerability-related indicators, attack samples, and malicious behavior rules related
    to LLM and AI agents;

11) The product must support AI-enhanced detection and blocking of unknown threats
    and variant attacks and automatically generate protection policies based on
    detection results;

12) The AI attack signature database must be updated in coordination with the
    traditional Intrusion Prevention System (IPS) signature database. Information such
    as detection signatures and policies must support real-time upload to the cloud,
    multi-point synchronization, automated updates, and manually triggered updates;

13) The product must support centralized distribution of custom AI attack signature
    rules by the management platform, ensure rapid response to and coordinated
    defense against emerging AI attacks, and support coordinated response with
    firewalls at the network layer;

14) Secure images from enterprise private repositories and officially certified sources,
    with code signatures, must be used. Runtime environments and dependency
    versions must be standardized, and image signatures must be verified during image
    transfer to prevent untrusted images from entering the production environment;

15) The product should support protection against man-in-the-middle (MITM) attacks
    targeting enterprise users accessing external LLM, detect and alert on abnormal
    proxy behavior or certificate hijacking, and ensure the security of communication
    channels.

