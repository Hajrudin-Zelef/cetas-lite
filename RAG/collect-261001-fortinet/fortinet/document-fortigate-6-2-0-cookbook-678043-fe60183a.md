---
id: collect-261001-fortinet/fortinet/document-fortigate-6-2-0-cookbook-678043-fe60183a
title: "diagnose sys virtual-wan-link sla-log ping 1"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-6-2-0-cookbook-678043-fe60183a.md
source_anchor: ""
source_lines: [1, 83]
sha256: 2fe45c0dc422ed7df9f6fad55803561dcaf789b60aac2d009b1e905711c52cc6
---

# diagnose sys virtual-wan-link sla-log ping 1

SLA monitoring using the REST API
SLA monitoring using the REST API
SLA log information and interface SLA information can be monitored using the REST API. This feature is also be used by FortiManager as part of its detailed SLA monitoring and drill-down features.
Interface log command example:
https://172.172.172.9/api/v2/monitor/virtual-wan/interface-log
{
  "http_method":"GET",
  "results":[
    {
      "interface":"port13",
      "logs":[
        {
          "timestamp":1547087168,
          "tx_bandwidth":3447,
          "rx_bandwidth":3457,
          "bi_bandwidth":6904,
          "tx_bytes":748875,
          "rx_bytes":708799,
          "egress_queue":[
          ]
        },
        {
          "timestamp":1547087178,
          "tx_bandwidth":3364,
          "rx_bandwidth":3400,
          "bi_bandwidth":6764,
          "tx_bytes":753789,
          "rx_bytes":712835,
          "egress_queue":[
          ]
        },
....
....
                                            SLA log command example:
https://172.172.172.9/api/v2/monitor/virtual-wan/sla-log
{
  "http_method":"GET",
  "results":[
    {
      "name":"ping",
      "interface":"port13",
      "logs":[
        {
          "timestamp":1547087204,
          "link":"up",
          "latency":0.686433,
          "jitter":0.063400,
          "packetloss":0.000000
        },
        {
          "timestamp":1547087205,
          "link":"up",
          "latency":0.688433,
          "jitter":0.063133,
          "packetloss":0.000000
        },
        {
          "timestamp":1547087206,
          "link":"up",
          "latency":0.688300,
          "jitter":0.065267,
          "packetloss":0.000000
        },
....
....
                                            CLI diagnose commands:
# diagnose sys virtual-wan-link sla-log ping 1
    Timestamp: Wed Jan 9 18:35:11 2019, vdom root, health-check ping, interface: port13, status: up, latency: 0.698, jitter: 0.073, packet loss: 0.000%.
    Timestamp: Wed Jan 9 18:35:12 2019, vdom root, health-check ping, interface: port13, status: up, latency: 0.704, jitter: 0.073, packet loss: 0.000%.
    Timestamp: Wed Jan 9 18:35:13 2019, vdom root, health-check ping, interface: port13, status: up, latency: 0.709, jitter: 0.073, packet loss: 0.000%.
    Timestamp: Wed Jan 9 18:35:14 2019, vdom root, health-check ping, interface: port13, status: up, latency: 0.707, jitter: 0.066, packet loss: 0.000%.
    Timestamp: Wed Jan 9 18:35:15 2019, vdom root, health-check ping, interface: port13, status: up, latency: 0.710, jitter: 0.061, packet loss: 0.000%.
    Timestamp: Wed Jan 9 18:35:16 2019, vdom root, health-check ping, interface: port13, status: up, latency: 0.707, jitter: 0.055, packet loss: 0.000%.
    Timestamp: Wed Jan 9 18:35:17 2019, vdom root, health-check ping, interface: port13, status: up, latency: 0.703, jitter: 0.055, packet loss: 0.000%.
# diagnose sys virtual-wan-link intf-sla-log port13
    Timestamp: Wed Jan 9 18:33:49 2019, used inbandwidth: 3208bps, used outbandwidth: 3453bps, used bibandwidth: 6661bps, tx bytes: 947234bytes, rx bytes: 898622bytes.
    Timestamp: Wed Jan 9 18:33:59 2019, used inbandwidth: 3317bps, used outbandwidth: 3450bps, used bibandwidth: 6767bps, tx bytes: 951284bytes, rx bytes: 902937bytes.
    Timestamp: Wed Jan 9 18:34:09 2019, used inbandwidth: 3302bps, used outbandwidth: 3389bps, used bibandwidth: 6691bps, tx bytes: 956268bytes, rx bytes: 907114bytes.
    Timestamp: Wed Jan 9 18:34:19 2019, used inbandwidth: 3279bps, used outbandwidth: 3352bps, used bibandwidth: 6631bps, tx bytes: 958920bytes, rx bytes: 910793bytes.
    Timestamp: Wed Jan 9 18:34:29 2019, used inbandwidth: 3233bps, used outbandwidth: 3371bps, used bibandwidth: 6604bps, tx bytes: 964374bytes, rx bytes: 914854bytes.
    Timestamp: Wed Jan 9 18:34:39 2019, used inbandwidth: 3235bps, used outbandwidth: 3362bps, used bibandwidth: 6597bps, tx bytes: 968250bytes, rx bytes: 918846bytes.
    Timestamp: Wed Jan 9 18:34:49 2019, used inbandwidth: 3165bps, used outbandwidth: 3362bps, used bibandwidth: 6527bps, tx bytes: 972298bytes, rx bytes: 922724bytes.
    Timestamp: Wed Jan 9 18:34:59 2019, used inbandwidth: 3184bps, used outbandwidth: 3362bps, used bibandwidth: 6546bps, tx bytes: 977282bytes, rx bytes: 927019bytes.
