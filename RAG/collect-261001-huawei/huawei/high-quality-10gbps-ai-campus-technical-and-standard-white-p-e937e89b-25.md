---
id: collect-261001-huawei/huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b-25
title: "high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b.md
source_anchor: ""
source_lines: [923, 964]
sha256: 65c602d7c65a218f7d54b21e2c6f2c0ef2fc7f236ea25ba8531e166d02053186
---

# high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b

2.5 Deterministic Experience                                                                                     2.5.2.2 Deterministic Bandwidth                             Leveraging high-precision time synchronization,
                                                                                                                                                                             TSN identifies key service traffic based on flow
                                                                                                                 Deterministic bandwidth means that bandwidths               rules and performs deterministic traffic scheduling
                                                                                                                 to be occupied by different services can be isolated        at precisely defined time points. In this way, the
                                                                                                                 from each other, preventing interference during             single-hop forwarding latency of key service traffic
 2.5.1          Definition                                                                                       data transmission. On a network that carries                is controlled within microseconds. The involved
                                                                                                                 numerous services, network slicing can be used to           key technologies include high-precision time
                                                                                                                 guarantee bandwidth for key services to prevent             synchronization (such as IEEE 1588v2) and traffic
As the digital economy continues to evolve and              effort" data transmission, which is inherently
                                                                                                                 packet loss due to insufficient bandwidth during            scheduling mechanism (such as 802.1Qbv gate
industries undergo intelligent transformation,              unreliable and often results in packet loss and
                                                                                                                 service concurrency.                                        control list scheduling).
campuses have emerged as key hubs for economic              transmission jitter. The deterministic experience
activit y and innovation. Consequently, the                 capability improves user experience by delivering
                                                                                                                 Network slicing technology allocates resources of
demands on campus networks have shifted from                deterministic priority, bandwidth, and latency, as                                                               2.5.2.4 High Reliability
                                                                                                                 a physical network to multiple logical networks,
basic connectivity to delivering superior user              well as high reliability, meeting network quality
                                                                                                                 with each logical network serving a specific                High reliability prevents unreliable transmission
experiences. Traditional networks rely on "best-            requirements in different scenarios.
                                                                                                                 type of service or industry. The logical topology,          caused by factors such as packet loss and delay
                                                                                                                 SLA requirements, reliability, and security level           during data forwarding. As technology advances,
                                                                                                                 of each network slice can be flexibly defined,              more and more services (such as payment
                                                                                                                 thereby meeting diverse requirements of services,           service and PLC control information transmission
 2.5.2          Key Technologies                                                                                 industries, and users.                                      for industrial equipment) have become highly
                                                                                                                                                                             sensitive to packet loss and require higher network
                                                                                                                                                                             reliability. A common technology for reducing the
2.5.2.1 Deterministic Priority                              2.5.2.1.2 User Priority                              2.5.2.3 Deterministic Latency
                                                                                                                                                                             packet loss rate is dual fed and selective receiving.
Deterministic priority means that both application          Deterministic user priority means that key users     Deterministic latency refers to a definite time
forwarding and user network resource allocation             receive prioritized access to network services.      required for end-to-end data transmission.                  In dual fed and selective receiving, the ingress
follow predefined, deterministic priority rules.            Key users include important enterprise users and     In addition to the inherent physical latency,               device replicates received data, and sends the
                                                            important terminals on the network, such as          traditional networks have uncertain latency caused          original data and its copy to the egress device
                                                            video conferencing terminals in conference rooms.    by processes such as queue-based transmission               through different paths. The egress device then
2.5.2.1.1 Application Priority
                                                            Technologies such as preferential access, wireless   and table lookup-based forwarding. Moreover, the            forwards the data that arrives first. Dual fed and
Deterministic application packet priority ensures           signal enhancement, and wireless preferential        best-effort communication mechanism used on                 selective receiving mainly solves the problem of
that packets of key applications, such as video             forwarding are used to ensure the wireless           traditional networks lacks determinism and real-            packet loss caused by single points of failure on
conferencing and cloud desktop, are preferentially          network experience of key users.                     time performance — capabilities that are critical           forwarding paths, significantly improving network
forwarded by dynamically adjusting their priorities.                                                             in fields such as industrial manufacturing. To solve        reliability.
This involves the following key technologies:               . Preferential access guarantees the access          the problem of uncertain latency, technologies
                                                              experience of key users by allowing them to        such as time-sensitive networking (TSN) are
