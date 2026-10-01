---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security-0f55dd46-1
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["context window", "watermarking"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46.md
source_anchor: ""
source_lines: [1, 319]
sha256: e22b181931ab6a25abfa7c84431aece98879d0aecbf4e2bf731780418c4f9842
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46

FortiOS Focus: DLP, Data Types, Dictionaries, Sensors, DLP Profiles, Fingerprinting, Proxy Inspection Audience: FortiGate / NSE4–NSE7 / Network & Security Engineers Primary Goal: Detect, log, block, quarantine, or restrict sensitive data leaving or entering the network Inspection Requirement: Proxy-based inspection for full DLP functionality Common Use Cases: Data leakage prevention, sensitive-file control, compliance, file inspection
| Component | Purpose | 
|---|---|
| DLP | Detects and controls sensitive information in network traffic | 
| Data Type | Defines what sensitive content should be detected | 
| Keyword | Matches specific words or phrases | 
| Regex | Matches data using regular expressions | 
| Hex | Matches binary/hexadecimal patterns | 
| Credit Card | Detects credit-card-like data | 
| SSN | Detects Social Security Number patterns | 
| Dictionary | Groups keywords/patterns into a logical detection set | 
| Sensor | Defines how DLP dictionaries/data types are evaluated | 
| DLP Profile | Defines protocol, file, action, and detection behavior | 
| Fingerprinting | Detects known files by comparing file fingerprints | 
| Watermarking | Detects data/files using embedded watermark information | 
| Proxy Inspection | Required for many advanced DLP inspection capabilities | 
| Deep Inspection | Allows inspection of encrypted HTTPS traffic where applicable | 
DLP — Data Leak Prevention is used to detect and control sensitive information moving through network traffic.
Typical use case:
Client
   │
   │ Upload / Download
   ▼
FortiGate
   │
   ▼
DLP Inspection
   │
   ├── Detect sensitive data
   ├── Log
   ├── Allow
   ├── Block
   └── Quarantine
   │
   ▼
Internet / External Service
Credentials
Personal information
Credit-card numbers
Social Security Numbers
Confidential documents
Source code
Internal keywords
Sensitive files
A simplified deployment:
                    INTERNET
                       │
                       ▼
                ┌──────────────┐
                │  FortiGate   │
                │              │
                │ Proxy        │
                │ Inspection   │
                │      │       │
                │      ▼       │
                │     DLP      │
                │      │       │
                │      ▼       │
                │ IPS / Other  │
                │ Profiles     │
                └──────┬───────┘
                       │
                       ▼
                     LAN
                    Clients
For HTTPS traffic:
HTTPS
  │
  ▼
SSL Deep Inspection
  │
  ▼
Decrypted Content
  │
  ▼
DLP Inspection
  │
  ▼
Other Security Profiles
  │
  ▼
Destination
Key idea: If FortiGate cannot inspect the actual content, DLP cannot reliably detect sensitive information inside that content.
One of the most important NSE concepts:
DLP
 │
 └── Proxy Inspection
For a firewall policy using DLP:
Firewall Policy
      │
      ├── Proxy Mode
      ├── Deep Inspection
      └── DLP Profile
LAN
 │
 ▼
Firewall Policy
 │
 ├── Proxy Inspection
 ├── SSL Deep Inspection
 ├── DLP Profile
 └── IPS / Other Security Profiles
 │
 ▼
Internet
Exam note: When using DLP features that require proxy processing, a flow-based policy is not sufficient.
If DLP options are hidden in the GUI:
System
  └── Feature Visibility
        └── DLP
Enable the DLP feature.
Depending on the FortiOS release, the GUI terminology may appear as:
Data Leak Prevention
or
DLP
Think of DLP as a hierarchy:
Data Types
    │
    ▼
Dictionaries
    │
    ▼
Sensors
    │
    ▼
DLP Profile
    │
    ▼
Firewall Policy
Or:
"What?"
   ↓
Data Type
"What group?"
   ↓
Dictionary
"How should it match?"
   ↓
Sensor
"What should FortiGate do?"
   ↓
DLP Profile
"Where should it apply?"
   ↓
Firewall Policy
DLP Data Types define the content that FortiGate should detect.
Common types:
Keyword
Regex
Hex
Credit Card
Social Security Number
Custom
Matches a specific word or phrase.
Example:
CONFIDENTIAL
INTERNAL
SECRET
Uses a regular expression to detect a pattern.
Example:
\b\d{3}-\d{2}-\d{4}\b
This can be used as part of a pattern for identifying SSN-like data.
Allows detection based on hexadecimal data patterns.
Useful when the sensitive content is represented in binary/hex form.
FortiGate can use built-in or customized patterns for credit-card-like information.
Can be detected using built-in or custom matching/verification patterns.
A DLP Data Type can use built-in patterns.
Example:
config dlp data-type
    edit "keyword"
        set pattern built-in
    next
    edit "regex"
        set pattern built-in
    next
    edit "hex"
        set pattern built-in
    next
end
Built-in patterns are supplied by FortiOS.
Custom patterns allow the administrator to define organization-specific detection logic.
Example conceptual configuration:
config dlp data-type
    edit "credit-card"
        set pattern "\\b([2-6]{1}\\d{3})[- ]?(\\d{4})[- ]?(\\d{2})[- ]?(\\d{2})[- ]?(\\d{2,4})\\b"
        set verify built-in
        set look-back 20
        set transform "\\b\\1[- ]?\\2[- ]?\\3[- ]?\\4[- ]?\\5\\b"
    next
end
| Parameter | Meaning | 
|---|---|
| pattern | Pattern used for detection | 
| verify | Additional validation/verification logic | 
| look-back | Context window around the detected value | 
| transform | Defines how matched capture groups are represented/transformed | 
Exact syntax and supported options should always be checked against the FortiOS version being used.
Example:
config dlp data-type
    edit "ssn-us"
        set pattern "\\b(\\d{3})-(\\d{2})-(\\d{4})\\b"
        set verify "(?<!-)\\b(?!666|000|9\\d{2})\\d{3}-(?!00)\\d{2}-(?!0{4})\\d{4}\\b(?!-)"
        set look-back 12
        set transform "\\b\\1-\\2-\\3\\b"
    next
end
Conceptually:
Raw Content
    │
    ▼
Pattern Match
    │
    ▼
Verification
    │
    ▼
Context Check
    │
    ▼
DLP Detection
A Dictionary groups patterns that belong to a logical category.
Example:
DLP Dictionary
      │
      ├── Keyword: CONFIDENTIAL
      ├── Keyword: INTERNAL
      ├── Keyword: SECRET
      └── Custom Pattern
GUI path:
Security Profiles
  └── Data Leak Prevention
        └── Dictionary
A dictionary can evaluate entries using logical relationships.
Keyword A
   OR
Keyword B
   OR
Keyword C
If one entry matches:
MATCH
Example:
Dictionary
 ├── PASSWORD
 ├── CONFIDENTIAL
 └── SECRET
Match Type = ANY
One matching entry can trigger the configured behavior.
Keyword A
   AND
Keyword B
   AND
Keyword C
All required entries must match.
Match Type = ALL
Dictionary: dlp-dic-test
Logical Relationship:
    ANY
Entries:
    ├── Type: Keyword
    │   Pattern: dlptest
    │
    └── Type: Keyword
        Pattern: confidential
Conceptually:
dlptest
   OR
confidential
   ↓
Dictionary Match
A DLP Sensor defines how dictionaries and detection entries are evaluated.
Architecture:
Data Type
   │
   ▼
Dictionary
   │
   ▼
Sensor
   │
   ▼
DLP Profile
Example:
config dlp sensor
    edit "sen-test"
        set match-type match-any
        config entries
            edit 1
                set dictionary "dlp-dic-test"
                set count 1
                set status enable
            next
        end
    next
end
Common logical behavior:
match-any
means:
Entry 1
   OR
Entry 2
   OR
Entry 3
Example:
Sensor
 │
 ├── Dictionary A
 ├── Dictionary B
 └── Dictionary C
match-any
A matching entry can trigger the sensor.
Example:
set count 1
Conceptually:
Dictionary Match Count >= 1
          ↓
Sensor Trigger
If the configured count is higher:
count 5
the configured detection must meet that threshold before the sensor triggers.
