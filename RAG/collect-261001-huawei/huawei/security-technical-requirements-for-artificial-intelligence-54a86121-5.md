---
id: collect-261001-huawei/huawei/security-technical-requirements-for-artificial-intelligence-54a86121-5
title: "security-technical-requirements-for-artificial-intelligence--54a86121"
domain: huawei
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["agent", "agents", "chatgpt", "inference", "mcp", "tool calling", "tool use"]
source: docs/RAG/collect-261001-huawei/security-technical-requirements-for-artificial-intelligence--54a86121.md
source_anchor: ""
source_lines: [701, 919]
sha256: b4451f688675c7eae10b6dcaa4fb8b4d99131e2dd25f701a95edd90095f70690
---

# security-technical-requirements-for-artificial-intelligence--54a86121

6.1.4.2.2. Tool calling attack protection
   1)   The product must support static security scanning during the registration of third-
        party Skills and MCP Servers, including source verification, signature verification,
        declaration of requested sensitive permissions, and matching against known
        vulnerabilities;

   2)   The product must be capable of protecting against tool description poisoning
        attacks, such as by blocking the poisoning of Skills repositories and prohibiting the
        automatic installation of Skills. All Skills must undergo security review before
        installation;

   3)   The product must support the detection of malicious Skills, identify requests to
        invoke malicious Skills, and block such requests in time;

   4)   The product must support the detection of attacks involving malicious registration
        of or tampering with tool calls (Function Calling/Tool Use), so as to prevent attackers
        from inducing models to call unauthorized internal APIs or perform dangerous
        operations;

   5)   The product must support call chain correlation analysis for auditing and traceability.


6.1.4.2.3. API attack protection
   1)   The product must support security protection for model service APIs, detect and
        block attacks such as SQL injection and cross-site request forgery (CSRF), and
        provide semantic analysis capabilities to improve the detection rate and accuracy;

   2)   The product must provide security protection capabilities for internal APIs, establish
        API baselines, identify shadow APIs and zombie APIs, and detect anomalous API
        attack behavior;

   3)   The product must support the detection and blocking of behavior that uses an LLM
        as a stepping stone to probe or attack internal networks, including SSRF, port
        scanning, and service discovery.


6.1.5. Security auditing, alerting, and statistics

6.1.5.1. Security auditing
   1)   The product must support security auditing functions, including but not limited to
        the following:

        a)   recording event types;

        b)   access requests matched by the product’s security policies;

        c)   detected attack behavior;
     d)   events involving autonomous access by AI agents to resources such as tools,
          MCP, APIs, and bare models;

     e)   events involving user access to sensitive resources, such as external LLMs
          including ChatGPT;

     f)   AI policy matching and response events.

2)   Log content:

     a)   the date and time of occurrence of the event;

     b)   the subject, object, and description of the event, where data packet logs include
          the protocol type, source address, destination address, source port, destination
          port, and other information;

     c)   a description of the attack event;

          ——In particular, for security attacks targeting inference and AI agents, the
          product must support comprehensive full-chain logging of user identity,
          prompt input, inference results, tool calls, execution results, data flows, and
          other relevant information, and must support chain reconstruction;

     d)   computing power resource access logs, including but not limited to the
          following:

          ——the identity of the human user or NHI caller;

          ——the name of the accessed model, AI agent, MCP service, tool, or knowledge
          base;

          ——the request type, request modality, and session identifier;

          ——the access result and policy enforcement result;

          ——session context association information, such as correlation indexes, to
          ensure that multi-turn requests within the same session are not recorded
          separately without correlation.

3)   Log management:

     a)   Only authorized administrators must be permitted to access logs, and functions
          such as log viewing and export must be provided;

     b)   The product must support intelligent aggregation and retrieval of session-level
          logs by authorized administrators based on dimensions such as session
          identifier, user identity, and time range;

     c)   Audit events must be searchable by criteria such as date, time, subject, and
          object;

     d)   Logs must be stored in non-volatile storage media that retain data in the event
          of power loss;
         e)   The log retention period must be set to no less than six months;

         f)   When the storage space reaches a specified threshold, the product must be
              capable of notifying authorized administrators and ensuring the normal
              operation of the auditing function;

         g)   Logs must support automated backup to other storage devices.


6.1.5.2. Security alerting
    1)   The product must support alerting on attack behavior defined in 7.1.5, variants of
         similar attacks, abnormal access behavior, policy violations, anomalies in AI
         capability invocation, cases where the confidence level of AI model inference results
         continuously falls below a threshold, and other abnormal conditions. Alert
         information must include at least the following:

         ——event subject;

         ——event object;

         ——event description;

         ——severity level;

         ——handling status;

         ——date and time of occurrence of the event.

    2)   The product must be capable of aggregating frequently recurring alert events of the
         same type. The product should use AI capabilities to perform correlation analysis on
         large volumes of logs, extract summaries, and aggregate risks, so as to reduce
         duplicate alerts and prevent alert storms.


6.1.5.3. Statistics

6.1.5.3.1. Network traffic statistics
The product must support the graphical display of network traffic information, including but
not limited to the following:

    1)   Network traffic statistics must be supported based on criteria such as IP address,
         time period, protocol type, NHI, AI service session, or any combination thereof.


6.1.5.3.2. Computing power usage statistics
The product must support statistics on model and computing power usage at the granularity
of resource callers and tasks, including but not limited to the following:

    1)   the number of times users, applications, and AI agents access model services;
    2)   the number of model inference requests and token usage.


6.1.5.3.3. Intelligent asset statistics
The product must support the discovery, statistics, and display of AI assets, including but not
limited to the following:

    1)   Model assets:

         ——model name;

         ——model type;

         ——service address;

         ——model version;

         ——last access time;

    2)   AI agent assets:

         ——AI agent name;

         ——AI agent version;

         ——associated user or organization;

         ——associated terminal information;

         ——associated model;

         ——times of accesses;；

         ——process information;

         ——last activity time;

    3)   MCP and tool assets:

         ——service name;

         ——protocol type;

         ——service address;

         ——times of calls;

         ——last access time;

    4)   Knowledge base and workflow assets:

         ——name;

         ——service type;

         ——call statistics;
         ——last update time.


6.1.5.3.4. Attack event statistics
    1)   When computing power resource usage exceeds a reasonable quota, the event
         must be recorded as a resource exhaustion event;

    2)   When the number of resource requests issued exceeds a reasonable quota, the
         event must be recorded as a DoS event;

