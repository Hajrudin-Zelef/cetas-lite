---
id: collect-261001-fortinet/fortinet/github-isarmadfarooq-fortigate-ngfw-security-lab-a-complete-hands-on-lab-for-deploying-and-3
title: "Required"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/github-isarmadfarooq-fortigate-ngfw-security-lab-a-complete-hands-on-lab-for-deploying-and-configuri.md
source_anchor: ""
source_lines: [378, 518]
sha256: 833346542f25d2d69bd29f43f5d7cc857cd140535e84ce2e5834d5634b9fb36d
---

# Required

The **Wildcard** type ensures subdomains like `m.facebook.com`, `static.xx.fbcdn.net`, and `connect.facebook.net` are also blocked, preventing evasion.


**Policy & Objects → Firewall Policy → Create New**

```
Name:                         No-Facebook-Internet-Access
Incoming Interface:           port3 (LAN)
Outgoing Interface:           port2 (WAN)
Source:                       all
Destination:                  all
Service:                      ALL
Action:                       ACCEPT
NAT:                          ✅ Enable — Use Outgoing Interface Address
Security Profiles:
  └── Web Filter:             ✅ [your web filter profile with facebook block]
```
**Policy & Objects → Firewall Policy → View By Sequence**

Drag `No-Facebook-Internet-Access` to **position #1** (above `Internet-Traffic`).

```
Policy Evaluation Order:
  [1] No-Facebook-Internet-Access   ← Web filter applied here FIRST
  [2] Internet-Traffic               ← General internet access
  [3] (implicit deny all)
```
❗ **If this policy is below `Internet-Traffic`, the web filter will never trigger.** FortiGate evaluates policies top-to-bottom and stops at the first match.


**From any LAN client browser, navigate to:**

```
https://www.facebook.com
https://m.facebook.com
```
**Expected result:**

```
⛔ Access to the requested website is blocked.
   URL: facebook.com
   Policy: No-Facebook-Internet-Access
   Category: Social Media
```
Verify the block event in **Log & Report → Security Events → Web Filter**

Run the following from **Kali Linux** (on a network segment connected to FortiGate):

```
# SYN Flood Attack — targets FortiGate management IP
hping3 -c 15000 -d 120 -S -w 64 -p 80 --flood --rand-source 192.168.205.2
```
**Flag breakdown:**

| Flag | Value | Description | 
|---|---|---|
| `-c` | `15000` | Total packets to send | 
| `-d` | `120` | Data payload size in bytes | 
| `-S` | — | Set TCP SYN flag | 
| `-w` | `64` | TCP window size | 
| `-p` | `80` | Destination port (HTTP) | 
| `--flood` | — | Send at maximum rate (no delay) | 
| `--rand-source` | — | Randomize source IP (simulates DDoS spoofing) | 

**Expected result:**

```
# hping3 output — no responses received = traffic is being dropped
HPING 192.168.205.2 (eth0 192.168.205.2): S set, 40+120 headers+data bytes
[main] memlockall(): Success
# FortiGate DoS Policy Log:
Date/Time     | Anomaly        | Action  | Source         | Destination
--------------+----------------+---------+----------------+------------------
[timestamp]   | tcp_syn_flood  | Dropped | [rand IP]      | 192.168.205.2:80
```
Verify in **Log & Report → Security Events → DoS**

| Task | Test | Expected | Actual | Status | 
|---|---|---|---|---|
| Task 1 | FortiGate VM boots & CLI accessible | Login prompt visible | ✅ Login successful | **PASS** | 
| Task 1 | Management IP reachable | Ping replies from 192.168.139.131 | ✅ Ping successful | **PASS** | 
| Task 1 | LAN client internet access | Ping 8.8.8.8 from LAN | ✅ Internet reachable | **PASS** | 
| Task 2 | HTTPS GUI accessible | Dashboard loads | ✅ GUI accessible | **PASS** | 
| Task 2 | New credentials work | Login with updated password | ✅ Authentication success | **PASS** | 
| Task 3A | IPv4 policy active | Policy visible in table | ✅ Policy enabled | **PASS** | 
| Task 3B | DoS policy active | Policy visible in table | ✅ Policy enabled | **PASS** | 
| Task 3C | Web filter enabled | Feature visible in GUI | ✅ Feature active | **PASS** | 
| Task 4 | facebook.com blocked | Block page displayed | ✅ Access denied | **PASS** | 
| Task 4 | SYN flood blocked | hping3 receives no replies | ✅ Traffic dropped | **PASS** | 

## **❓ Cannot access FortiGate GUI after setting IP**

```
# Verify the IP was set correctly via CLI
show system interface port1
# Check that allowaccess includes https
# If not, re-run:
config system interface
    edit port1
        set allowaccess https ssh ping
    next
end
# Verify your host machine is on the same subnet
# Host IP should be in 192.168.139.0/24 range
```
## **❓ Browser shows certificate error for FortiGate GUI**

This is expected. FortiGate uses a self-signed certificate by default.

- **Chrome:** Click**Advanced** →**Proceed to 192.168.139.131 (unsafe)**
- **Firefox:** Click**Advanced** →**Accept the Risk and Continue**
- For production: Install a valid certificate via **System → Certificates**

## **❓ LAN clients cannot reach the internet**

Check in order:

1. **Default route exists:** Network → Static Routes → verify 0.0.0.0/0 route
2. **IPv4 policy exists:** Policy & Objects → Firewall Policy → verify Internet-Traffic policy
3. **NAT is enabled** on the firewall policy
4. **WAN interface IP is correct** (matches ISP configuration)
5. **DNS is configured:** Network → DNS → set DNS servers (8.8.8.8, 8.8.4.4)

## **❓ facebook.com is not being blocked**

1. Verify the web filter policy is **at the TOP** of the policy list (By Sequence view)
2. Confirm the web filter profile is **attached** to the policy (Security Profiles section)
3. Check that the URL entry is `facebook.com` , Type`Wildcard` , Action`Block` , Status`Enabled`
4. Clear browser cache and retry (browser may have cached the page)
5. Check FortiGate web filter logs: Log & Report → Security Events → Web Filter

## **❓ DoS policy is not triggering**

1. Verify the DoS policy interface is set to the **WAN interface** (where attack traffic arrives)
2. Check that the specific anomaly profile (e.g., `tcp_syn_flood` ) status is**Enabled**
3. Check that Action is set to **Block** (not Pass)
4. Verify the hping3 command is targeting the correct IP address
5. Check logs: Log & Report → Security Events → DoS

## **❓ VMware shows import error for FortiGate OVF**

