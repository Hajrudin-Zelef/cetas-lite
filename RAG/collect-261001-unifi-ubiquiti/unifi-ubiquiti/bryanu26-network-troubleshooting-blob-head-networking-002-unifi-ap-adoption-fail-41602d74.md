---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/bryanu26-network-troubleshooting-blob-head-networking-002-unifi-ap-adoption-fail-41602d74
title: "bryanu26-network-troubleshooting-blob-head-networking-002-unifi-ap-adoption-fail-41602d74"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/bryanu26-network-troubleshooting-blob-head-networking-002-unifi-ap-adoption-fail-41602d74.md
source_anchor: ""
source_lines: [1, 16]
sha256: 5c3dfa82256f027848942b96d33ac45e0b48cde265bf13cedaf177573a9c376a
---

# bryanu26-network-troubleshooting-blob-head-networking-002-unifi-ap-adoption-fail-41602d74

Residential network using UniFi infrastructure.
- Access Point appeared in the UniFi Controller.
- Device remained in a pending or failed adoption state.
- AP did not become manageable through the controller.
- Verified PoE power to the AP.
- Confirmed network connectivity.
- Checked DHCP assignment.
- Verified the controller could reach the AP.
- Performed a factory reset by holding the reset button for 10+ seconds.
- Re-attempted adoption through the UniFi Controller.
- Verified firmware after successful adoption.
The Access Point had not been properly factory reset, preventing it from completing the adoption process.
Performed a complete factory reset and adopted the AP again through the UniFi Controller. Once adopted, updated the firmware and verified normal operation.
- Always perform a full factory reset before troubleshooting adoption issues.
- Verify power and DHCP before assuming a controller problem.
- Keep firmware updated after a successful adoption to avoid compatibility issues.
