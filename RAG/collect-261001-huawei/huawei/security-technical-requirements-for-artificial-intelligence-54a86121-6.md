---
id: collect-261001-huawei/huawei/security-technical-requirements-for-artificial-intelligence-54a86121-6
title: "security-technical-requirements-for-artificial-intelligence--54a86121"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "inference", "latency", "parameters"]
source: docs/RAG/collect-261001-huawei/security-technical-requirements-for-artificial-intelligence--54a86121.md
source_anchor: ""
source_lines: [920, 1072]
sha256: 8dbf0aa0618cb7265955b6f30f6477d8064c3e606ae88ae488dcfa9226a13362
---

# security-technical-requirements-for-artificial-intelligence--54a86121

    3)   The product must support categorized statistics on prompt injection attacks, tool
         abuse attacks, unauthorized model access attacks, tool call attacks, and anomalous
         AI agent behavior;

    4)   When multi-stage attacks or complex call-chain attacks occur, the product must
         support attack path tracing and call chain reconstruction to reconstruct and record
         the attack process, attack targets, and attack intent;

    5)   The product must support intelligent simplification of statistical results to prevent
         operational fatigue caused by massive volumes of logs.


6.2.     Security requirements of the Product Itself

6.2.1. Identification and authentication
The security requirements for identification and authentication of the product include but are
not limited to the following:

    1)   The product must identify and authenticate users, and user identities must be
         uniquely identifiable;

    2)   The product must securely protect user authentication information and ensure its
         confidentiality during storage and transmission;

    3)   The product must provide failed login attempt handling functions, such as limiting
         the times of consecutive invalid login attempts;

    4)   The product must provide login timeout handling functions and automatically
         terminate a login session when the session times out;

    5)   When password-based authentication is used, the product must check the
         complexity of user-defined passwords to ensure that user passwords meet specified
         complexity requirements;

    6)   Where default passwords exist in the product, users must be prompted to change
         them so as to reduce the risk of user identity impersonation;

    7)   Authorized administrators must be authenticated using a combination of two or
         more authentication techniques;
    8)   The product must identify and authenticate callers of AI model inference services to
         ensure that only authorized application systems or users can invoke AI model
         inference interfaces, thereby preventing unauthorized access and model theft;

    9)   The product must identify AI models themselves to ensure that each model can be
         uniquely identified during loading, updating, and invocation, thereby preventing
         model tampering or replacement.


6.2.2. Management capabilities
The security requirements for the management capabilities of the product include but are not
limited to the following:

    1)   The product must provide authorized administrators with functions for setting and
         modifying data parameters related to security management;

    2)   The product must provide authorized administrators with functions for setting,
         querying, and modifying various security policies;

    3)   The product must provide authorized administrators with functions for managing
         audit logs;

    4)   The product must support updates to its own system, including software system
         upgrades and upgrades to various signature databases;

    5)   The product must be capable of synchronizing system time with a Network Time
         Protocol (NTP) server;

    6)   The product must support the synchronization of information such as logs and alerts
         to a log server through the Syslog protocol;

    7)   The product must distinguish administrator roles and support the division of roles
         into system administrator, security operator, and security auditor, with the privileges
         of the three administrator roles mutually constrained;

    8)   The product must provide security policy validity checking functions, such as
         checking security policy matching status;

    9)   The product should provide authorized administrators with AI model lifecycle
         management functions, including model loading, unloading, version rollback,
         staged rollout, and full release, to ensure that the model update process is
         controllable and traceable;

    10) The product should provide authorized administrators with functions for configuring
        and adjusting AI inference policies, including model inference threshold settings,
        confidence level adjustment, and inference result handling rules, enabling
        administrators to flexibly control AI inference behavior according to actual security
        requirements;

    11) The product should provide AI model operational status monitoring functions,
         including real-time display of and alerting on key indicators such as model inference
         latency, resource utilization, and inference success rate.


6.2.3. Management auditing
The security requirements for management auditing of the product include but are not limited
to the following:

    1)   The product must log operations such as user account login and logout, system
         startup, significant configuration changes, addition, deletion, or modification of
         administrators, and saving or deletion of audit logs;

    2)   The product must generate alerts for abnormal states of the product and its modules
         and record such states in logs;

    3)   Log records must include the following information: date and time of occurrence of
         the event, event type, event subject, and operation result;

    4)   Only authorized administrators must be permitted to access logs;

    5)   The product should log lifecycle operations on AI models, such as loading, updating,
         unloading, and rollback, including information such as operation time, operator,
         model version, and operation result;

    6)   The product should maintain audit records of key decision-making behavior during
         AI model inference, including the source of the inference request, inference result,
         confidence level, and security policy actions triggered or generated, so as to ensure
         that the AI decision-making process is traceable and auditable;

    7)   The product should generate alerts and log AI model inference anomalies, including
         inference failures, inference timeouts, abnormal fluctuations in output results, and
         confidence levels that deviate significantly from the normal range.


6.2.4. Management methods
The security requirements for the management methods of the product include but are not
limited to the following:

    1)   The product must support local management through a console port;

    2)   The product must support remote management through network interfaces and
         must be capable of restricting the IP addresses and MAC addresses from which
         remote management is permitted;

    3)   During remote management, all communication data between the management
         terminal and the product must be transmitted in non-plaintext form;

    4)   The product must support monitoring and management through the Simple
         Network Management Protocol (SNMP);
    5)   The product must support separation of management interfaces from service
         interfaces;

    6)   The product must support centralized management through a centralized
         management platform, including monitoring operational status, distributing security
         policies, upgrading system versions, and upgrading signature database versions;

    7)   The product must support management of AI model services through RESTful APIs
         or other standardized interfaces, including operations such as model loading,
         unloading, status queries, and inference testing, and must enforce access control
         and rate limiting on API calls;

