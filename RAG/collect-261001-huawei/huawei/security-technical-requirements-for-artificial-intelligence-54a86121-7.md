---
id: collect-261001-huawei/huawei/security-technical-requirements-for-artificial-intelligence-54a86121-7
title: "security-technical-requirements-for-artificial-intelligence--54a86121"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["inference", "parameters", "training"]
source: docs/RAG/collect-261001-huawei/security-technical-requirements-for-artificial-intelligence--54a86121.md
source_anchor: ""
source_lines: [1073, 1210]
sha256: e815424d1c129891df627fc444a448c6ba40d3f1c51d5de026b37c34c1dbee04
---

# security-technical-requirements-for-artificial-intelligence--54a86121

    8)   The product must support version updates and configuration changes of AI models
         through secure remote management methods. During the update process, integrity
         verification of model files must be ensured to prevent tampering with or
         replacement of model files during transmission;

    9)   The product must provide health check interfaces for AI model services, enabling
         external monitoring systems to obtain information such as model operational status,
         inference performance metrics, and resource utilization in real time.


6.2.5. Security-supporting systems
The security requirements for the support system of the product include but are not limited
to the following:

    1)   Necessary tailoring must be performed so that no redundant components or
         network services are provided;

    2)   Security policies and log information must not be lost during system restart;

    3)   The product must contain no known medium- or high-risk security vulnerabilities;

    4)   The AI model runtime environment should be securely isolated from the business
         processing environment to prevent abnormalities in the model runtime environment
         from affecting the normal operation of business processing and to prevent
         unauthorized interference by business processing processes with the model runtime
         environment;

    5)   The AI model runtime environment must be configured in accordance with the
         principle of least privilege. Only system resources and network access permissions
         necessary for model inference must be enabled, and the file system access, network
         connection, and system call capabilities of the model runtime environment must be
         restricted;

    6)   The libraries and frameworks on which the AI model runtime environment depends
         must contain no known medium- or high-risk security vulnerabilities. An inventory
         of dependencies and vulnerability scanning mechanism must be established to
         regularly inspect and update components with potential security risks.
6.2.6. AI model security
The security requirements for AI models of the product include but are not limited to the
following:

    1)   The product must perform integrity verification of AI model files. Before a model is
         loaded, hash verification, digital signatures, or other means must be used to verify
         that the model files have not been tampered with, replaced, or corrupted, so as to
         prevent the loading of malicious models from rendering security policies ineffective;

    2)   The product must be capable of resisting adversarial example attacks. Anomaly
         detection and filtering must be performed on input data to identify and reject
         carefully crafted adversarial example inputs, thereby preventing attackers from
         misleading model decisions through minor perturbations to input data;

    3)   The product should be capable of resisting model extraction attacks and must take
         effective measures to prevent attackers from extracting model parameters, model
         architecture, or characteristics of training data through means such as issuing large
         numbers of queries to interfaces and observing output results. Such measures may
         include limiting the precision of confidence scores in model outputs and adding
         perturbations to outputs;

    4)   The product should establish security boundaries for AI model inference behavior,
         including restrictions on the data range of inference inputs, the types of decisions in
         inference outputs, and the frequency of inference invocations, so as to prevent
         models from being used for unintended purposes or in scenarios beyond their
         designed scope;

    5)   The product should establish an AI model security testing mechanism. Security
         performance testing must be conducted before model deployment and periodically
         during operation, including adversarial robustness testing, model extraction defense
         testing, and data poisoning defense testing, so as to ensure model reliability under
         security threat scenarios.


6.2.7. Data security and privacy protection
The data security and privacy protection requirements for the product include but are not
limited to the following:

    1)   Temporary data and intermediate results generated during AI model inference
         should be securely managed. Upon completion of inference, temporary data must
         be promptly cleared to prevent intermediate results from being accessed or used
         without authorization;

    2)   Storage protection should be provided for AI model training datasets and feature
         data. Measures such as encrypted storage and access control should be adopted to
         prevent data leakage, tampering, or unauthorized copying.
6.2.8. AI algorithm security
The AI algorithm security requirements for the product include but are not limited to the
following:

    1)   The product must establish human review or secondary confirmation mechanisms
         for significant decisions made by AI models. Where AI model inference results would
         lead to major changes to security policies or affect critical business operations,
         execution must proceed only after confirmation by an authorized administrator, so
         as to prevent uncontrollable risks arising from autonomous AI decision-making;

    2)   The product should conduct security assessments of the design and implementation
         of AI model algorithms, including algorithm security analysis, identification of
         potential risks, and detection of security defects, to ensure that the algorithms
         themselves contain no exploitable security vulnerabilities or backdoors;

    3)   The product should record and manage complete version information of AI models,
         including metadata such as model version numbers, training dataset versions,
         feature repository versions, training parameter configurations, and training time, to
         ensure model traceability and reproducibility;

    4)   The product may establish continuous monitoring and model degradation detection
         mechanisms for AI models, periodically evaluate model performance in real-world
         operating environments, and generate timely alerts and trigger model update
         processes when significant model performance degradation is detected.


6.3.     Centralized management requirements

6.3.1. Platform integration
    1)   The product must support standardized integration with unified firewall
         management platforms and integrated security management platforms, as well as
         centralized management, and must support unified management and control of
         multiple distributed devices to eliminate isolated silos of standalone device
         management;

    2)   The product must provide standardized external interfaces through standard
         protocols such as NETCONF, RESTful API, Syslog over TLS, and SNMPv3, to ensure
         compliant and stable interoperability with management platforms.


6.3.2. Centralized device management
    1)   The product must support centralized lifecycle management of devices by the
         management platform, including batch registration, remote operations and
         maintenance, firmware version upgrades, security signature database updates,
         centralized configuration backup, and one-click rollback;

    2)   The product must support normalized reporting to the management platform of
         data such as device operational status, computing power load, network-wide traffic
         logs, AI security attack events, and protection statistics metrics, so as to support
         unified monitoring, statistics, and analysis by the platform.


