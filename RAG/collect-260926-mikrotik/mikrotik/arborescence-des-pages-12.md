---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-12
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-12.md
source_anchor: ""
source_lines: [1, 29]
sha256: 128f067858b3ff9048f542cb16c74957cbcf56334c7c1fe57a46c520caae1b11
---

# Introduction

The Speed Test is an easy test tool for measuring ping, jitter, TCP and UDP throughput from one MikroTik device, to another. The "speed-test" command is based on the Ping Tool and Bandwidth Test. In order to use this command - Bandwidth test server needs to be accessible.

# General interface properties

The speed-test is based on five configurable properties:

- address - IP address of host;
- connection-count - If a device has more than 20 cores - core count will be used (default is 20);
- password - Password for the remote device;
- test-duration - Duration for each test (*By default: 5 tests * 10 sec duration + 1sec pause between each test = 55sec* );
- user - Remote device username;

# Configuration Example

Bandwidth and speed tests should be conducted through the devices, not on them to ensure real-life simulation and not to overload the CPU on the devices under testing,(DUT) due to the traffic generating process.

To run a simple test from device A (192.168.88.1) to device B (192.168.88.2):

If any of device CPU utilization during test reaches 100% warning message will appear:

"test-duration" parameter allows changing the duration of all of the 5 tests:

- ) Ping test with 50ms delay
- ) TCP receive
- ) TCP send
- ) UDP receive
- ) UDP send
