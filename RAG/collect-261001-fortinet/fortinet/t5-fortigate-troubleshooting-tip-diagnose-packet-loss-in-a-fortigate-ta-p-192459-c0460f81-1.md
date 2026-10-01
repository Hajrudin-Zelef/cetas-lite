---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-diagnose-packet-loss-in-a-fortigate-ta-p-192459-c0460f81-1
title: "t5-fortigate-troubleshooting-tip-diagnose-packet-loss-in-a-fortigate-ta-p-192459-c0460f81"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2025-07-12"]
keywords: ["ethernet", "memory"]
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-diagnose-packet-loss-in-a-fortigate-ta-p-192459-c0460f81.md
source_anchor: ""
source_lines: [1, 164]
sha256: a45cc4146ab47a72098a1fbf217e2281bf381d6efa985acb18b0411bca3c16db
---

# t5-fortigate-troubleshooting-tip-diagnose-packet-loss-in-a-fortigate-ta-p-192459-c0460f81

**Description**


This article describes how, in certain circumstances, a FortiGate deployment may experience higher packet loss than normal and some common reasons for this behavior. There are also recommendations on how to resolve common issues or test hardware for possible problems.


**Scope**

FortiGate.**Solution**

Several factors can cause packet loss on the FortiGate. Here are the most common reasons:


1. **Incorrect speed settings on the interface:**  Check the speed settings on each interface from the GUI by moving the mouse over the interface on**System -> Status -> Unit Operation,** or by running the following CLI command:


`diagnose hardware deviceinfo nic <interface name>`

Repeating the CLI command above several times (up to 3 outputs) will also help to check the drop counters that can help to narrow down the issue further. 

Further details on the drop counters can be referenced from these KB articles: 

- Technical Tip: Matching port counters with the various components of NP6-based hardware platforms
- Troubleshooting Tip: Network Interface Card NIC commands


To check if there are errors in the interface, use the following command:

```
fnsysctl ifconfig port1
port1 Link encap:Ethernet HWaddr D4:76:A0:1C:6D:B4
inet addr:192.168.10.1 Bcast:192.168.10.255 Mask:255.255.255.0
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
RX packets:10608 errors:0 dropped:0 overruns:0 frame:0
TX packets:5437 errors:0 dropped:0 overruns:0 carrier:0
collisions:0 txqueuelen:1000 
RX bytes:2232859 (2.1 MB) TX bytes:684968 (668.9 KB)
```

Users will be looking for a speed of 10half. This usually means that the FortiGate was not able to negotiate the speed correctly with the device on the other side.


To set the speed manually, use the commands:

```
config system interface
    edit <interface name>
        set speed 100full
end
```

**Warning:**

Some vendors will turn off the interface if auto-negotiate is turned off on the FortiGate. Make sure not to be connected through the same link being changed, or the connection to the FortiGate may be lost.


To check if there are errors or drops in an interface, use the above commands.


1. **High bandwidth usage:**  To generate bandwidth reports, make sure to have logging on firewall policies enabled. This is done by going to**Firewall -> Policy** and editing the policies. Enable logging by enabling 'log allowed traffic'.


On a FortiAnalyzer, go to **Report -> Config -> Layout -> Create New -> Add** charts as needed. Most users will need **Traffic Volume by Direction, Top Services by Volume, and Top Sources by Volume***.*


In **Report -> Schedule -> Create New ->** use the layout that was just created and select the devices (that is, FortiGates) on which to run the report. Select **OK**. Schedule the report or run it on demand using the **'Run now'** icon on the **Report -> Schedule** page.


1. **Hardware issues:** The problem can be caused by a hardware problem. Proceed to run the HQIP diagnostics, especially the network loopback test, to see if the physical ports are having a hardware problem or not. Refer to this article: Technical Tip: RMA - HQIP test (with built-in FortiOS diagnostic commands).

1. **Session stats:**  Use the session stat command to check for the increment trend of counters like 'extreme_low_mem', 'memory_tension_drop', 'ephemeral', and error stats by running this command multiple times and see if these counters increment continuously, and troubleshoot further accordingly based on which session stat counter is indicating a possible issue. If packets are being offloaded to NPU, use the command 'diagnose sys npu-session stat' to review session stats in the hardware.


```
diagnose sys session stat
misc info: session_count=21 setup_rate=0 exp_count=0 clash=0
memory_tension_drop=0 ephemeral=0/25231360 removeable=0 extreme_low_mem=0
npu_session_count=0
nturbo_session_count=0
delete=0, flush=1, dev_down=58/6610
session walkers: active=0, vf-40, dev-0, saddr-0, npu-0, wildcard-57
TCP sessions:
7 in ESTABLISHED state
firewall error stat:
error1=00000000
error2=00000000
error3=00000000
error4=00000000
tt=00000000
cont=00000000
ips_recv=00000000
policy_deny=00079e27
av_recv=00000000
fqdn_count=00000009
fqdn6_count=00000000
global: ses_limit=0 ses6_limit=0 rt_limit=0 rt6_limit=0
```

1. **Traffic Shapers:**  If a traffic shaper was applied, check the session list for possible drops.


Run 'diagnose sys session list' to see the session list details. 

```
diagnose sys session list
session info: proto=6 proto_state=11 duration=30 expire=3599 timeout=3600 flags=00000000 socktype=0 sockport=0 av_idx=0 use=4
origin-shaper=high-priority prio=2 guarantee 0Bps max 131072000Bps traffic 186Bps drops 0B 
reply-shaper=high-priority prio=2 guarantee 0Bps max 131072000Bps traffic 186Bps drops 0B 
per_ip_shaper=
class_id=0 shaping_policy_id=1 ha_id=0 policy_dir=0 tunnel=/ vlan_cos=0/255
state=log may_dirty ndr os rs f00 
statistic(bytes/packets/allow_err): org=1467/12/1 reply=1262/8/1 tuples=3
tx speed(Bps/kbps): 47/0 rx speed(Bps/kbps): 40/0
orgin->sink: org pre->post, reply pre->post dev=5->3/3->5 gwy=10.9.15.254/0.0.0.0
hook=post dir=org act=snat 192.168.1.2:56525->34.107.221.82:80(10.9.12.64:56525)
hook=pre dir=reply act=dnat 34.107.221.82:80->10.9.12.64:56525(192.168.1.2:56525)
hook=post dir=reply act=noop 34.107.221.82:80->192.168.1.2:56525(0.0.0.0:0)
pos/(before,after) 0/(0,0), 0/(0,0)
misc=0 policy_id=1 pol_uuid_idx=15747 auth_info=0 chk_client_info=0 vd=0
serial=0010ca5c tos=ff/ff app_list=0 app=0 url_cat=0
rpdb_link_id=00000000 ngfwid=n/a
npu_state=0x001108
no_ofld_reason: redir-to-ips denied-by-nturbo 
```

For more information about packets dropped by the traffic shaper, refer to Troubleshooting Tip: How to check packet drop by traffic shaper in NP6, NP6xlite and NP6lite unit.


**Run debug flow trace on the FortiGate and check the output:**

```
diagnose debug enable
diagnose debug flow filter addr X.X.X.X 
diagnose debug console timestamp enable
diagnose debug flow show iprope enable
diagnose debug flow show function-name enable
diagnose debug flow trace start 100 
diagnose debug enable
```
**In this example:** 

- X.X.X.X is the IP address of interesting traffic.
- 'diagnose debug flow trace start 100' will display 100 packets for this flow.

**The output will look like what is displayed below:**

```
2025-07-12 12:45:21 id=320 trace_id=18 func=__iprope_tree_check line=539 msg="gnum-100004, use addr/intf hash, len=10"
2025-07-12 12:45:21 id=320 trace_id=18 func=get_new_addr line=1231 msg="find SNAT: IP-168.8.168.250(from IPPOOL), port-60418"
2025-07-12 12:45:21 id=320 trace_id=18 func=fw_forward_handler line=990 msg="Allowed by Policy-614: SNAT"
2025-07-12 12:45:21 id=320 trace_id=18 func=shaper_handler line=884 msg="exceeded shaper limit, drop"
```

Check the traffic shaping policy, and adjust the shaping policy to accommodate more bandwidth or disable the traffic shaping policy.


1. **CPU and memory usage** : FortiGate may drop packets due to high memory or CPU usage.


Run 'get system performance status' to find the CPU and memory usage.

Note the top 3 lines of output in this example:


