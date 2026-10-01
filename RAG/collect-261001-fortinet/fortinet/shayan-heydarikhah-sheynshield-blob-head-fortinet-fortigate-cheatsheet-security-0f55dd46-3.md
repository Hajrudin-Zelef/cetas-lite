---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security-0f55dd46-3
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["memory", "throughput", "watermarking"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46.md
source_anchor: ""
source_lines: [663, 947]
sha256: 874ca29d39631ea072486c146367e8c7b138dd1b7cba34d9f7a6130ccfd686d9
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46

      │      ↓
      │   No Match
      │
      └── ≥ 40%
             ↓
          Match
             ↓
           Action
Common concepts:
Block
Ban
Quarantine
Restricts the matching DLP activity.
Violation
   ↓
Block
Can be used for stronger source restriction, particularly where repeated or abusive activity needs to be contained.
Repeated violation
       ↓
BAN
       ↓
Source restriction
Provides stronger restriction behavior for the violating source/activity according to the configured DLP policy.
Violation
   ↓
Quarantine
   ↓
Restricted activity
Exact timeout and scope behavior should be verified against the FortiOS release and DLP profile configuration.
             DOCUMENT REPOSITORY
                     │
                     ▼
              FortiGate scans
                     │
                     ▼
              Fingerprint DB
                     │
                     ▼
              Network Traffic
                     │
                     ▼
             DLP Fingerprint
                     │
                     ▼
                 Compare
                     │
              ┌──────┴──────┐
              ▼             ▼
            Match         No Match
              │             │
              ▼             ▼
           Action         Allow
Diagnostic command:
diagnose test application dlpfingerprint
Possible menu options include:
1   Show fingerprint daemon menu
2   Dump the database
3   Dump all files
5   Dump all chunks
6   Refresh all document sources in all VDOMs
7   Show database file size and limit
9   Display statistics
10  Clear statistics
99  Restart the daemon
Fingerprint not matching
        │
        ▼
Check Document Source
        │
        ▼
Refresh Database
        │
        ▼
Check Fingerprint DB
        │
        ▼
Check Statistics
        │
        ▼
Check DLP Profile
        │
        ▼
Test Traffic
[ ] DLP feature enabled
[ ] Correct DLP profile selected
[ ] Firewall policy uses proxy mode
[ ] Deep inspection enabled for HTTPS
[ ] Correct DLP sensor configured
[ ] Correct dictionary configured
[ ] Correct data type configured
[ ] Sensor match type verified
[ ] Rule protocol matches traffic
[ ] File type matches traffic
[ ] File size is within configured limit
[ ] Action is correctly configured
[ ] Logs enabled
Check:
HTTPS
  ↓
SSL Deep Inspection
  ↓
Decrypted Content
  ↓
DLP
  ↓
DLP Match?
If the content remains encrypted:
DLP
  ↓
Cannot inspect actual payload
Investigate:
Cloud Service
      ↓
Custom API
      ↓
Encoding / Metadata
      ↓
DLP Detection
      ↓
File blocked ✓
      │
      └── Filename logging may be inaccurate
Do not assume the DLP engine failed simply because the filename in the log is incorrect.
Recommended enterprise architecture:
                       INTERNET
                           │
                           ▼
                  ┌────────────────┐
                  │   FortiGate    │
                  │                │
                  │ SSL Inspection │
                  │      ↓         │
                  │     DLP        │
                  │      ↓         │
                  │     IPS        │
                  │      ↓         │
                  │ Web/App Control│
                  └───────┬────────┘
                          │
                          ▼
                         LAN
A strong DLP deployment can combine:
                DLP Security
                     │
       ┌─────────────┼─────────────┐
       ▼             ▼             ▼
   Data Types    Fingerprint    Watermark
       │             │             │
       └─────────────┼─────────────┘
                     ▼
                  Sensor
                     │
                     ▼
                DLP Profile
                     │
                     ▼
              Firewall Policy
                     │
                     ▼
            SSL Deep Inspection
DLP does not have to operate alone.
Example:
Traffic
   │
   ▼
SSL Deep Inspection
   │
   ▼
DLP
   │
   ├── Sensitive Data?
   │
   ▼
IPS
   │
   ▼
Application Control
   │
   ▼
Web Filter
   │
   ▼
Forward
This creates a layered security architecture.
DLP inspection can consume additional resources because FortiGate may need to:
Receive content
      ↓
Buffer / process content
      ↓
Inspect content
      ↓
Compare patterns
      ↓
Run additional security profiles
      ↓
Forward traffic
- Limiting unnecessary protocols
- Limiting file sizes
- Selecting relevant file types
- Using targeted sensors
- Avoiding unnecessary inspection
- Monitoring CPU/memory
- Testing cloud applications
- Validating throughput under real workloads
| Topic | Remember | 
|---|---|
| DLP | Prevents sensitive-data leakage | 
| Data Type | Defines detectable content | 
| Keyword | Matches words/phrases | 
| Regex | Matches structured patterns | 
| Dictionary | Groups matching entries | 
| Sensor | Defines detection logic | 
| DLP Profile | Defines inspection/action behavior | 
| Proxy Mode | Important for DLP processing | 
| Deep Inspection | Required to inspect decrypted HTTPS content | 
| Fingerprinting | Detects known files/documents | 
| Watermarking | Detects recognized watermark information | 
| Block | Blocks matching DLP activity | 
| Log Only | Detects/logs without blocking | 
| Quarantine | Stronger restriction behavior | 
| Ban | Strong source restriction for abusive activity | 
| Cloud DLP | API/encoding can affect filename logging | 
| File Type | Prefer when the requirement is based on file type | 
| Internal Storage | Required for document fingerprinting | 
| Resource Usage | DLP adds processing overhead | 
The easiest way to remember FortiGate DLP:
                 FORTIGATE DLP
                       │
                       ▼
                  WHAT TO FIND?
                       │
                       ▼
                  DATA TYPE
                       │
                       ▼
                HOW TO GROUP IT?
                       │
                       ▼
                  DICTIONARY
                       │
                       ▼
                HOW TO EVALUATE?
                       │
                       ▼
                    SENSOR
                       │
                       ▼
              WHAT SHOULD HAPPEN?
                       │
                       ▼
                 DLP PROFILE
                       │
                       ▼
               WHERE TO APPLY?
                       │
                       ▼
              FIREWALL POLICY
                  CLIENT
                     │
                     │ HTTPS
                     ▼
               ┌─────────────┐
               │  FortiGate  │
               └──────┬──────┘
                      │
                      ▼
              SSL Deep Inspection
                      │
                      ▼
                Decrypted Data
                      │
                      ▼
                     DLP
                      │
          ┌───────────┼───────────┐
          ▼           ▼           ▼
       Keyword      Regex     Fingerprint
          │           │           │
          └───────────┼───────────┘
                      ▼
                   Sensor
                      │
                      ▼
                 DLP Profile
                      │
