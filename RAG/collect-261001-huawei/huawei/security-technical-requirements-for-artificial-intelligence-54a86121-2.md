---
id: collect-261001-huawei/huawei/security-technical-requirements-for-artificial-intelligence-54a86121-2
title: "security-technical-requirements-for-artificial-intelligence--54a86121"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "distribution", "inference", "training"]
source: docs/RAG/collect-261001-huawei/security-technical-requirements-for-artificial-intelligence--54a86121.md
source_anchor: ""
source_lines: [238, 392]
sha256: 01dac7702d69a7cbabc70ce058931af1ceba012c25306c7c886fd73dad977670
---

# security-technical-requirements-for-artificial-intelligence--54a86121

    1)   The product must support interconnection with multiple AI services and be capable
         of forwarding requests to corresponding capability endpoints based on access
         destinations, request types, request modalities, security policies or business policies;

    2)   The product must support capability routing based on the identity, Internet Protocol
         (IP) address, network zone, data sensitivity level, and request modality of the request
         caller, as well as security policies or business policies, and must support the
         configuration of static routing, policy-based routing, or dynamic routing;

    3)   The product must support parsing application-layer authentication information in
         traffic, such as API keys, user tokens, and service identity identifiers, and dynamically
         allowing or denying requests based on predefined routing policies, as a supplement
         to network-layer information.


6.1.1.5. High availability
In addition to meeting the relevant requirements specified in 6.1.1.3 of GB/T 20281—2020,
the product must also provide AI-oriented load balancing functions:

    1)   The product must support dynamic load sharing across model services and
         computing power zones based on computing power resource utilization, for
         example, through token usage and coordinated sensing with computing servers;

    2)   In the event of an error, the product must automatically retry using other models or
         computing resource clusters to achieve automatic failover.


6.1.2. Network-layer control

6.1.2.1. Access control

6.1.2.1.1. Packet filtering
In addition to meeting the relevant requirements specified in 6.1.2.1.1 of GB/T 20281—2020,
the product must also meet the following requirements for the packet filtering function:

    1)   User-defined security policies must be supported. Such security policies must
         support combinations of some or all of the following: MAC address, IP address, port,
         user identity, AI agent identity, terminal type, application type, protocol type, and
         time.


6.1.2.1.2. Network address translation
The product must meet the relevant requirements specified in 6.1.2.1.2 of GB/T 20281—2020.


6.1.2.1.3. Stateful inspection
The product must meet the relevant requirements specified in 6.1.2.1.3 of GB/T 20281—2020.


6.1.2.1.4. Dynamic open ports
The product must meet the relevant requirements specified in 6.1.2.1.4 of GB/T 20281—2020.


6.1.2.1.5. IP/MAC address binding
In addition to meeting the relevant requirements specified in 6.1.2.1.5 of GB/T 20281—2020,
the product must also support intelligent learning of IP/MAC address binding relationships
and anomaly detection.


6.1.2.1.6. IP Blacklisting and Whitelisting
The product must support IP address whitelist and blacklist functions, including but not
limited to the following:

    1)   The product must support the configuration of source/destination IP address
         whitelists and blacklists;
         — For example, for sensitive sessions initiated by NHI, an whitelists and blacklists -
         based access control mode must be supported, whereby sessions are permitted to
         be established only with the IP addresses and ports of pre-authorized destinations,
         such as public network resources, internal network models, and necessary services.

    2)   Batch distribution of whitelists and blacklists must be supported.


6.1.2.2. Traffic management

6.1.2.2.1. Bandwidth management
In addition to meeting the relevant requirements specified in 6.1.2.2.1 of GB/T 20281—2020,
the product must also support the following bandwidth management functions:

    1)   Bandwidth must be configurable separately for upstream and downstream traffic,
         and upstream and downstream traffic must be limited separately;

    2)   Bandwidth allocation and limitation based on AI service session attributes must be
         supported, and independent bandwidth quotas must be configurable for different
         unique AI agent identifiers and session service types;

    3)   Limits on the number of concurrent sessions based on unique AI agent identifiers
         and session service types must be supported;

    4)   Session-based traffic rate limiting must be supported, enabling the data
         transmission rate of an individual AI service session to be limited;

    5)   Identification and rate limiting of abnormal traffic sessions must be supported,
         including sessions with sudden traffic surges, sessions carrying unusually large
         volumes of traffic, and sessions suspected of remote control via covert channels.


6.1.2.2.2. Connection control
The product must support limiting the maximum number of concurrent sessions and the new
connection establishment rate for an individual AI agent, so as to prevent network
performance from being affected by a large number of unauthorized connections.


6.1.2.2.3. Session management
  The product must support the following session management functions:

    1)   The product must terminate a session when the session remains inactive for a
         specified period, when an abnormal condition occurs (e.g. the number of session
         establishments and terminations per unit time exceeds a threshold), or after the
         session ends;

    2)   The product must provide session state monitoring and attribute management
         capabilities to support AI service scenarios, establish standard five-tuple sessions for
        network traffic traversing the boundary, and perform full-lifecycle management of
        such sessions;

   3)   The product must support the configuration of independent session timeout periods
        for different types of AI service sessions. The configuration granularity must reach
        the session service type level so as to meet the requirements of long-duration
        sessions for services such as model training and batch inference;

   4)   The product must support the configuration of access control policies based on
        extended attributes of AI service sessions, including session initiator type, unique AI
        agent identifier, and session service type;

   5)   The product should extend AI service sessions with the following verifiable and
        configurable dedicated session attributes. All such attributes must be directly
        accessible to modules such as access control and traffic management:

        ——Session initiator type: identifies the identity type of the session initiator, with
        values including “human user”, “NHI”, “process”, and others;

        ——Unique AI agent identifier: an identifier used to identify an AI agent at the
        network layer. It may be derived from IP address bindings, preconfigured mapping
        tables, fixed port ranges, information carried in packets, or other means;

        ——Session service type: identifies the category of AI service corresponding to the
        session, with values including “model invocation”, “training data”, “inter-agent
        communication”, “model synchronization”, and others.


6.1.2.3. Network isolation
   1)   The product must support security domain partitioning, enabling AI service
        components, such as AI agents, model services, and training clusters, to be deployed
        in independent security domains and isolated at the network layer from the office
        network, business system zone, database zone, Internet egress zone, and operations
        and maintenance management zone;

   2)   All network sessions between different security domains, including AI service
        sessions, must pass through the firewall. The firewall must serve as the sole ingress
        and egress point for inter-domain communication;

