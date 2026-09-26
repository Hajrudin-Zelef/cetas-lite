---
id: collect-260926-mikrotik/mikrotik/index-php-stc-article-download-369-398-fc120320-1
title: "QoS aware traffic shaping for TikTok live streaming over congested LAN using MikroTik queue tree"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2026-06"]
keywords: ["consumer", "distribution", "latency", "parameters", "research", "throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/index-php-stc-article-download-369-398-fc120320.md
source_anchor: ""
source_lines: [1, 94]
sha256: 17a24b46982be8ddd86a4131f495849a9c328ba10e3d3362ada484ea68a84e26
---

# QoS aware traffic shaping for TikTok live streaming over congested LAN using MikroTik queue tree

*Source : https://sintechcomjournal.com/index.php/stc/article/download/369/398*
*SINTECHCOM: Science, Technology and Communication Journal, Vol. 6, No. 3, June 2026, pp. 199-206, DOI: 10.59190/stc.v6i3.369*
*Diah Risqiwati, Wiryawan Ananta Pratama Panigoro, Hanugra Aulia Sidharta*

## Abstract

Real time video streaming applications such as TikTok Live are highly sensitive to network instability, especially under bandwidth contention on shared access networks. This paper evaluates a Quality of Service (QoS)-aware traffic shaping scheme for TikTok live streaming in a congested local area network using a MikroTik router with a queue tree configuration. TikTok traffic is identified and assigned minimum bandwidth guarantees and highest priority, while non streaming traffic is treated as best effort. Network performance is assessed under two scenarios: baseline (without shaping) and experiment (with shaping), using the TIPHON standard and four QoS parameters: throughput, delay, jitter, and packet loss.

The experimental results show that although the average throughput only increases slightly from 70.67 Kbps to 80 Kbps and remains in the "Poor" category, the proposed scheme significantly improves temporal QoS metrics: average delay is reduced from 88.65 ms to 45.16 ms ("Very Good"), jitter decreases from 95.61 ms to 89.80 ms ("Good"), and packet loss drops from 7.66% to 3.78% ("Good"). These findings indicate that priority-based traffic shaping using a queue tree can effectively stabilize latency and data delivery for TikTok live streaming on bandwidth-limited networks without requiring capacity upgrades.

Keywords: Limited Resource, MikroTik, QoS, Queue Tree, Traffic Shaping

## 1. Introduction

Short-form video platforms such as TikTok have rapidly become major consumer of IP traffic, particularly due to their highly engaging live streaming features that blend real-time audio visual content with interactive viewer participation. TikTok Live is increasingly used not only for entertainment but also for social commerce, where live product demonstrations and real time promotions significantly influence purchase decisions.

At the same time, many organizations, campuses, and small internet service providers in emerging markets operate under strict bandwidth with infrastructure constraints. Upgrading last mile capacity is often limited by budget and contractual conditions, so multiple users and applications must share the same constrained access link. In such environments, real time video traffic competes with web browsing, file downloads, and other best effort services, while inadequate Quality of Service (QoS) control leads to buffering, reduced video quality, and audio visual de-synchronization.

Video QoS is typically characterized by throughput, delay, jitter, and packet loss. TIPHON provides widely used threshold values as standard, such as delay below 150 ms as "Very Good", 150–300 ms as "Good", and above 450 ms as "Poor". Similar qualitative categories are defined for jitter and packet loss, with packet loss below 3% generally considered "Good" or better for real-time applications.

In parallel, a pragmatic engineering practice has evolved around controlling network traffic using MikroTik RouterOS. Official documentation and expert presentations describe how queue trees, Per Connection Queuing (PCQ), and Stochastic Fairness Queuing (SFQ) can be combined with firewall mangle rules and DSCP markings to prioritize latency sensitive flows (VoIP, streaming video) while deprioritizing bulk traffic.

Research gaps identified: (1) most QoS studies focus on generic services, not TikTok Live specifically; (2) many prior works are purely observational, without experimentally validating concrete application-aware traffic shaping strategies; (3) little peer reviewed evidence on how a specific MikroTik queue tree configuration affects TIPHON-based QoS metrics for TikTok Live under congested LAN conditions.

This paper proposes and evaluates an application aware traffic shaping scheme for TikTok live streaming on a bandwidth limited LAN using MikroTik RouterOS. Contributions: extends QoS analysis to TikTok Live; moves from passive measurement to active control with a concrete queue tree configuration; provides a TIPHON-based empirical evaluation offering a reproducible configuration pattern for network administrators.

## 2. Research Methods

### 2.1 Research Design

Experimental research design comparing two configurations: baseline (no traffic shaping) vs experimental (TikTok traffic prioritized using a MikroTik queue tree), under identical hardware, topology, and traffic load.

### 2.2 Network Testbed

A MikroTik router acting as Internet gateway, an access switch, and several client devices on wired/wireless links. One client runs TikTok Live; others generate background traffic (web browsing, file downloads) to emulate congestion. Upstream/downstream bandwidth set to low fixed values. Star topology with the MikroTik router at the center behind a single bottleneck link.

### 2.3 Traffic Shaping Configuration

Implemented via address lists, firewall mangle rules, and a hierarchical queue tree:

1. Identify TikTok server IPs, add them to an address list.
2. For each new connection, if destination IP is in the TikTok list, mark connection as `conn_tiktok`.
3. For each packet, if connection mark is `conn_tiktok`, mark packet as `pkt_tiktok`; otherwise `pkt_general`.
4. Queue tree on the bottleneck interface: child queue Q_tiktok mapped to `pkt_tiktok` with highest priority (Priority 1) and minimum bandwidth guarantee.
5. Child queue Q_general mapped to `pkt_general` with lower priority, best-effort.

PCQ or SFQ can be used as underlying queue types for fair bandwidth distribution within each class.

### 2.4 Measurement Scenarios

1. **Baseline** (no shaping): all flows best-effort, TikTok Live + background traffic saturating the bottleneck.
2. **Experimental** (shaping enabled): same traffic, TikTok packets prioritized with bandwidth guarantees.

Each scenario repeated three times; experiments conducted during similar time periods.

### 2.5 Data Collection and QoS Metrics

Packet traces captured at the gateway; four TIPHON metrics derived:

| Metric | Poor | Fair | Good | Very good |
|---|---|---|---|---|
| Throughput (%) | 0–25 | 26–50 | 51–75 | 76–100 |
| Delay (ms) | > 450 | 300–450 | 150–300 | < 150 |
| Jitter (ms) | 125–225 | 75–125 | 0–75 | 0 |
| Packet loss (%) | > 25 | 16–25 | 3–15 | 0–2 |

## 3. Results and Discussions

### 3.1 Baseline: QoS without Traffic Shaping

Average throughput 70.67 Kbps (< 25% of allocated bandwidth → "Poor"). Average one-way delay 88.65 ms ("Very Good"), jitter 95.61 ms ("Good"), packet loss 7.66% ("Good"). Jitter near 100 ms and loss above 7% increase risk of buffering, frame drops, and A/V desynchronization.

| QoS metric | Run 1 | Run 2 | Run 3 | Average | TIPHON |
|---|---|---|---|---|---|
| Throughput (kbps) | 73 | 76 | 63 | 70.67 | Poor |
| Delay (ms) | 104.60 | 101.81 | 59.55 | 88.65 | Very good |
| Jitter (ms) | 96.10 | 93.39 | 97.33 | 95.61 | Good |
| Packet loss (%) | 7.71 | 10.83 | 4.43 | 7.66 | Good |

### 3.2 Experimental: QoS with Queue Tree Traffic Shaping

Average throughput rises slightly to 80 Kbps (still "Poor"). Delay drops from 88.65 ms to 45.16 ms (halved, "Very Good"). Jitter reduced from 95.61 ms to 89.80 ms ("Good"). Packet loss drops from 7.66% to 3.78% ("Good") — more than 50% fewer dropped packets.

| QoS metric | Run 1 | Run 2 | Run 3 | Average | TIPHON |
|---|---|---|---|---|---|
| Throughput (kbps) | 127 | 59 | 54 | 80 | Poor |
| Delay (ms) | 41.70 | 51.64 | 45.13 | 46.16 | Very good |
| Jitter (ms) | 83.42 | 92.80 | 93.17 | 89.80 | Good |
| Packet loss (%) | 3.15 | 4.02 | 4.16 | 3.78 | Good |

### 3.3 Detailed Comparative Analysis

