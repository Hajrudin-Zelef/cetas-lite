---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-large-file-downloads-fail-over-an-ipsec-tunnel-u-6f0cbeb5
title: "fortigate-3-troubleshooting-tip-large-file-downloads-fail-over-an-ipsec-tunnel-u-6f0cbeb5"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-large-file-downloads-fail-over-an-ipsec-tunnel-u-6f0cbeb5.md
source_anchor: ""
source_lines: [1, 32]
sha256: 6dc1901281b3948c522c4cd8bfa2be023132b2a4c1cd42e075acd42c6fa62472
---

# fortigate-3-troubleshooting-tip-large-file-downloads-fail-over-an-ipsec-tunnel-u-6f0cbeb5

**Description**

This article describes why downloading a file larger than the effective tunnel MTU fails over an IPsec VPN tunnel configured with vpn-id-ipip encapsulation when NAT-T is in use, and how to resolve it.**Scope**

FortiGate.**Solution**

Small file downloads over the tunnel complete normally, but a download of a file large enough to require fragmentation stalls after the initial exchange and does not progress further.

The tunnel interface MTU and the IPsec SA MTU can be checked with the following commands.

```
diagnose netlink interface list <phase1-interface> | grep mtu
diagnose vpn tunnel list name <phase1-interface> | grep mtu
```

When vpn-id-ipip encapsulation is used together with NAT-T, the tunnel interface MTU reported by the netlink interface list may be set higher than the effective MTU supported by the IPsec SA, which is shown as mtu 1402 in the SA details from the tunnel list. This mismatch causes large TCP packets to be sent at a size the tunnel cannot forward without fragmentation, and downloads of files above the effective MTU threshold fail or stall.

The tunnel interface MTU can be overridden to match the IPsec SA MTU on the phase1-interface, as shown below.

```
config system interface
    edit "<phase1-interface>"
        set mtu-override enable
        set mtu 1402
    next
end
```

After applying this change, the tunnel interface MTU can be confirmed against the SA MTU using the same commands, and the download of a large file across the tunnel can be re-tested.


Starting from FortiOS v8.0.1 and v7.6.7, the default MTU assigned to IPsec tunnel interfaces was lowered from 1420 to 1402 on all FortiGate models, so a manual override is no longer required on these releases to reach an MTU that matches the IPsec SA in this scenario.
