---
id: collect-240926-storagereview/storagereview/fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a-4
title: "fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a"
domain: storagereview
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["latency"]
source: docs/RAG/clean_en/storagereview/fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a.md
source_anchor: ""
source_lines: [62, 69]
sha256: ac99ac66bb11bd438bc4df96ab7757561daa9cfb003d5b401b38e959dfeaf567
---

# fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a

vSphere 8 continues to add features and enhancements and recently increased supported namespaces to 256 and paths to 2K for NVMe-FC and TCP. Another feature, support for reservation commands for NVMe devices, was added to vSphere. Reservation commands allow customers to use clustered VMDK capability with Microsoft WSFC with NVMe-oF datastores.
Simple to configure and simple to manage!
Fibre Channel has another built-in efficiency: automatic discovery. When an FC device is connected to the network, it is automatically discovered and added to the fabric if it has the necessary credentials. The node map is updated and traffic can traverse the fiber. This is a simple process without administrator intervention.
There is more overhead when implementing NVMe/TCP. Because NVMe/TCP does not have an automatic discovery mechanism, ESXi added support for the NVMe discovery service. Advanced NVMe-oF discovery service support in ESXi enables dynamic discovery of the standards-compliant NVMe discovery service. ESXi uses the mDNS/DNS-SD service to obtain information such as the IP address and port number of active NVMe-oF discovery services on the network. ESXi sends a multicast DNS (mDNS) query requesting information from entities providing the discovery service (NVMe) (DNS-SD). If such an entity is active on the network (on which the query was sent), it will send a response (unicast) to the host with the requested information, i.e., the IP address and port number where the service is running.
Conclusion
Fibre Channel was specifically designed to carry block storage and, as stated in this article, it is a reliable, low-latency, lossless, and high-performance fabric. To be clear, progress is being made to improve the use of TCP for storage traffic on certain important high-speed networks. But the fact remains that TCP is not a lossless network, and data retransmission is always an issue.
Marvell FC Product Family
This report is sponsored by Marvell. All views and opinions expressed in this report are based on our impartial view of the product(s) under study.
