---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-25656226682775-multi-link-operation-mlo-in-unifi-network-38d9abd6
title: "hc-en-us-articles-25656226682775-multi-link-operation-mlo-in-unifi-network-38d9abd6"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google", "Intel", "Samsung"]
dates: []
keywords: ["intel", "latency", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-25656226682775-multi-link-operation-mlo-in-unifi-network-38d9abd6.md
source_anchor: ""
source_lines: [1, 45]
sha256: e39c2fd86e75632e76517a23e5d66262943606ed9d4dabc117f8cfef4782f326
---

# hc-en-us-articles-25656226682775-multi-link-operation-mlo-in-unifi-network-38d9abd6

Multi-Link Operation (MLO) in UniFi Network
Multi-Link Operation (MLO) is a WiFi 7 feature that allows compatible client devices to maintain simultaneous connections to multiple bands—2.4, 5, or 6 GHz—on a single UniFi Access Point. Instead of switching between bands, MLO-enabled clients stay associated with multiple links and choose the best one(s) for transmitting and receiving data. This helps reduce latency, improve connection stability, and in some cases, increase throughput.
For more information on optimizing WiFi connectivity, click here.
Why Use MLO?
Enabling MLO can offer a number of practical benefits for modern networks:
- Reduced latency, especially in high-density environments.
- Improved reliability by allowing clients to maintain multiple links.
- Better performance in congested areas by distributing traffic across bands.
- Enhanced throughput for supported clients.
Requirements
To use MLO in UniFi Network, you'll need:
- UniFi Network version 8.2.93 or higher
- A WiFi 7 Access Point running firmware version 7.1.18 or higher or Wi-Fi Integrated Cloud Gateway supporting Wi-Fi 7
- A dual- or tri-band setup (e.g., 2.4/5/6 GHz or any combination)
MLO also requires MLO-compatible client devices, such as laptops with Intel BE200 wireless adapters, the Samsung Galaxy S24 Ultra, and the Google Pixel 8 and Pixel 8 Pro.
Enabling MLO
- Go to Settings > WiFi.
- Click the SSID you want to enable MLO on.
- Check the box next to Multi-Link Operation (MLO).
Types of Multi-Link Operation
The way a client uses MLO depends on its hardware. Here are the main types:
- 
MLSR (Multi-Link Single Radio)
  - The client connects on multiple channels but uses a single radio that can only operate on one band at a time (i.e., it switches between bands).
- 
EMLSR (Enhanced Multi-Link Single Radio)
  - The client still uses a single radio but can listen on multiple bands while transmitting on just one. This improves latency in crowded networks.
- 
Async MLMR (Asynchronous Multi-Link Multi-Radio)
  - A client uses multiple radios to transmit and receive simultaneously across different bands, with each link operating independently without tight synchronization.
- 
Sync MLMR (Synchronous Multi-Link Multi-Radio)
  - The client uses multiple radios to transmit and receive simultaneously across different bands, with transmissions coordinated and synchronized across the links.
These client types determine how MLO behaves in practice, especially in terms of latency and throughput.
What to Expect When Enabling MLO
The impact of MLO depends entirely on the client’s capabilities:
- 
Latency Reduction
  - MLO can improve performance in dense environments by allowing clients to quickly switch to less congested links or listen on multiple bands simultaneously.
- 
Connection Stability
  - Clients can select the most stable link in real time, improving reliability in areas with fluctuating signal quality. This makes MLO a great choice for optimizing call quality.
- 
Throughput Gains
  - While not guaranteed, clients with support for simultaneous transmission and reception with multi-radio setups may see better speeds when connected to an MLO capable UniFi AP.
