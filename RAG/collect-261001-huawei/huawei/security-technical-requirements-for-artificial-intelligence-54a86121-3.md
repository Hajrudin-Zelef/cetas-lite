---
id: collect-261001-huawei/huawei/security-technical-requirements-for-artificial-intelligence-54a86121-3
title: "security-technical-requirements-for-artificial-intelligence--54a86121"
domain: huawei
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "OpenAI"]
dates: []
keywords: ["agent", "agents", "chatgpt", "deepseek", "inference", "mcp", "multimodal", "parameters", "qwen", "tool calling", "tool use"]
source: docs/RAG/collect-261001-huawei/security-technical-requirements-for-artificial-intelligence--54a86121.md
source_anchor: ""
source_lines: [393, 540]
sha256: a5f55b06550ea864ead1af4f30d9767fa53bb4c662b8b4671bb0d119d2e05d3e
---

# security-technical-requirements-for-artificial-intelligence--54a86121

   3)   Intra-domain or inter-domain access control policies must be established based on
        session attributes, specifying the session types permitted to communicate and
        prohibiting unauthorized sessions.


6.1.2.4. Encrypted traffic parsing
   1)   The product must support Secure Socket Layer (SSL)/ Transport Layer Security (TLS)
        proxy decryption of HTTPS/HTTP traffic, including TLS 1.2 and TLS 1.3;

   2)   The product must support AI-enhanced feature analysis of encrypted traffic.
6.1.3. Application-layer control

6.1.3.1. Application protocol control
   1)   The product must support the parsing and control of general-purpose application-
        layer protocols, such as HTTP/HTTPS, WebSocket, and gRPC;

   2)   The product must support the parsing and control of AI agent communication
        protocols, such as Agent to Agent Protocol (A2A), including Message content
        scanning and Artifact content identification;

   3)   The product should support the parsing and control of emerging AI agent
        communication protocols, such as ACP, ANP, and SLIM;

   4)   The product must support the parsing of model context protocols, such as Model
        Context Protocol (MCP), and identify and control context information and tool call
        requests transmitted through such protocols;

   5)   The product must support the parsing and control of OpenAI-compatible protocols
        and interfaces, including Function Calling, Tool Use, and the Responses API;

   6)   The product must support semantic-level parsing of protocol messages, identify
        invocation intents such as tool calls, resource access, and prompt retrieval, and
        enforce controls based on the identified intents.


6.1.3.2. Identity and authorization management
   1)   The product must support authentication-based network access control;

   2)   The product must support the identification and authentication of external NHIs that
        invoke internal resources, using credentials such as X.509 certificates, OAuth 2.0
        tokens, and JWTs;

   3)   The product must support the identification of internal NHIs, such as AI agents,
        digital humans, embodied intelligent entities, workloads, machine identities,
        application accounts, and process IDs.

   4)   The product should support integrating NHIs into Identity and Access Management
        (IAM), establishing a unified identity management system, applying access control
        models such as Role-Based Access Control (RBAC), Policy-Based Access Control
        (PBAC), and Attribute-Based Access Control (ABAC), and integrating them with
        network-layer security policies;

   5)   The product must support identity traceability for enterprise users accessing
        sensitive resources, such as publicly available external LLM (for example including
        ChatGPT, DeepSeek, and Qwen), and must identify and record the identity of the
        user initiating the request, the department to which the user belongs, and terminal
        device information;
   6)   The product must support behavior and intent analysis for NHIs and human users,
        construct and train self-learning intelligent behavior models by collecting
        multidimensional behavioral features, and use such models to identify and block
        threat entities and behaviors.


6.1.3.3. Application-layer resource general access control
   1)   Application resources subject to access restrictions must perform unified
        authentication of callers using API keys, X.509 certificates, ID tokens, or other means;

   2)   The product must support dynamic authorization and access control, evaluate the
        trustworthiness of accessing entities based on information such as the invoking user’
        s identity, role, scenario, and context, grant resource accessors the minimum
        necessary privileges according to the evaluation results, and perform real-time
        blocking and/or privilege freezing when anomalies occur;

   3)   The product must intercept high-concurrency malicious requests targeting
        application-layer resources and support protection against application-layer
        denial-of-service (DoS) attacks.


6.1.3.3.1. Model resource control
   1)   Resources such as AI agents and LLM must be exposed as APIs when providing
        services externally;

   2)   The product must support registration and admission control, as well as fine-
        grained authorization, for MCP service resources;

   3)   The product must support monitoring the utilization status of computing power
        resources for business application models and regulating or denying abnormal
        behaviors, such as low-priority processes preempting high-priority resources and
        prolonged resource exhaustion;

   4)   The product must support control over token usage by business application model
        inference tasks or sessions and support a circuit breaker mechanism based on
        token-level budgets.


6.1.3.3.2. Tool calling control
   1)   For specific sensitive scenarios, an allowlist and denylist mechanism for tool calls
        must be established, and tool calls must comply with predefined authorization
        policies;

   2)   The product should support legitimacy checks on tool calls made through
        mechanisms such as MCP and Skills, and detect attacks initiated through tool
        parameters, including Structured Query Language (SQL) injection, command
        injection, path traversal, and Server-Side Request Forgery (SSRF);
   3)   The product should support the redirection of high-risk operations.


6.1.3.3.3. API call control
   1)   The product must support real-time monitoring of API access status. When the API
        access volume exceeds a specified threshold, access must be blocked to prevent
        access channels from becoming congested;

   2)   The product must support data leakage detection for API response content, identify
        whether responses contain sensitive data, and trigger alerts or blocking actions;

   3)   The product must support auditing of API call chains, record and generate complete
        call chain logs, and facilitate post-event auditing and leakage tracing;

   4)   The product must support risk assessment of third-party API calls and identify risky
        behaviors, such as excessive data sharing and unauthorized transmission, when
        enterprise applications call external LLM APIs.


6.1.3.4. Content Output control
   1)   The product must support protection against prompt injection attacks, covering
        both direct and indirect prompt injection, such as persistent inducement across
        multiple turns, context concatenation, and concealed content in front-end interfaces;

   2)   The product must support the detection of malicious instructions hidden in
        multimodal content, including images, audio, video, and files;

   3)   The product must support the configuration of policies that prohibit or restrict users
        from uploading sensitive internal enterprise data, such as code repositories,
        customer information, financial data, and internal documents, to external publicly
        available LLM;

   4)   The product should support content parsing and sensitive information scanning of
        network traffic and multimodal or multi-format files uploaded to LLM, such as PDF,
        Word, Excel, PowerPoint, and archive files, so as to prevent internal data leakage
        through attachments;

   5)   The product must support the detection and filtering of files downloaded by users
        from LLM responses, identify whether downloaded files contain prohibited or
        sensitive content, and support truncation during the output process;

