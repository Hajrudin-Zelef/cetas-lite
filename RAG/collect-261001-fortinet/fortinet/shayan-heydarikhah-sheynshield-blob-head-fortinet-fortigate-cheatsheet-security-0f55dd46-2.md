---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security-0f55dd46-2
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["containment", "memory", "throughput", "watermarking"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46.md
source_anchor: ""
source_lines: [320, 662]
sha256: f2cf3aa0b131af422ed594b48e123ad8c410783dbe607f47fc9b3cb2d43e1d30
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46

The DLP Profile determines what FortiGate should do after detecting sensitive information.
GUI:
Security Profiles
  └── Data Leak Prevention
        └── Profile
Typical actions include:
Allow
Log Only
Block
Quarantine
Conceptual configuration:
config dlp profile
    edit "dlp-prof-test"
        set feature-set proxy
        config rule
            edit 1
                set proto http-get
                set filter-by none
                set file-type 1
                set file-size 500
                set action block
            next
        end
    next
end
set feature-set proxy
is important when the DLP feature requires proxy-based processing.
A DLP rule can control:
Protocol
Filter
File Type
File Size
Action
Example:
DLP Rule
 │
 ├── Protocol: HTTP
 ├── Filter: File Type
 ├── File Type: Built-in
 ├── File Size: 500 KB
 └── Action: Block
DLP can be restricted to specific file types.
Example concept:
File Type
   │
   ├── Documents
   ├── Archives
   ├── Images
   ├── Executables
   └── Other types
The exact numeric file-type IDs depend on the FortiOS implementation.
NSE note: Do not memorize numeric IDs without verifying them against the target FortiOS release.
Example:
set file-size 500
Conceptually:
Maximum / configured file size
        ↓
500 KB
This allows administrators to limit which files are inspected by a particular DLP rule.
Common DLP actions:
| Action | Behavior | 
|---|---|
| Allow | Permit the activity | 
| Log Only | Allow while recording the violation | 
| Block | Block the matching activity | 
| Quarantine | Restrict the violating source/activity according to the profile behavior | 
LOG ONLY
   ↓
Detect + Log
   ↓
Traffic continues
BLOCK
   ↓
Matching activity is denied
QUARANTINE
   ↓
Stronger restriction behavior
   ↓
Used when repeated/serious violations need containment
Exact quarantine behavior can vary by FortiOS version and DLP configuration.
A useful DLP concept:
Filters are ordered, but the possible actions do not have precedence over each other simply because they appear in a particular order.
Think of:
Filter 1
Filter 2
Filter 3
as ordered evaluation logic, while:
Allow
Block
Quarantine
Log
are configured actions rather than an inherent priority hierarchy.
Typical deployment:
LAN
 │
 ▼
Firewall Policy
 │
 ├── Proxy Inspection
 ├── Deep Inspection
 ├── DLP Profile
 └── IPS Profile
 │
 ▼
Internet
Source      : LAN
Destination : Internet
Service     : Required services
Inspection  : Proxy
SSL         : Deep Inspection
DLP         : dlp-prof-test
IPS         : Required IPS Sensor
NAT         : Enable
For encrypted web traffic:
Client
  │
  │ HTTPS
  ▼
FortiGate
  │
  ├── SSL Deep Inspection
  │
  ▼
Decrypted Content
  │
  ├── DLP
  ├── IPS
  ├── Application Control
  └── Web Filtering
  │
  ▼
Internet
Without decryption:
HTTPS
  │
  ▼
Encrypted Payload
  │
  X
DLP cannot inspect hidden content
Important: SSL inspection and DLP should be designed together when the goal is preventing leakage through HTTPS.
DLP introduces additional processing overhead.
Traffic
   ↓
Proxy
   ↓
Content Processing
   ↓
DLP
   ↓
Other Security Profiles
   ↓
Forwarding
Compared with simple firewall forwarding:
More inspection
      ↓
More CPU / Memory
      ↓
Potentially lower throughput
- Size FortiGate appropriately
- Test expected traffic volume
- Consider file size limits
- Avoid inspecting unnecessary content
- Use targeted DLP rules
- Monitor CPU and memory
- Test cloud applications separately
Cloud services can create special DLP challenges.
Examples:
Google Drive
SharePoint
Cloud Storage
Web-based SaaS
A browser upload does not always look like a simple:
HTTP POST
   ↓
file.pdf
Cloud applications may use:
Proprietary APIs
Custom encoding
Multiple HTTP requests
Metadata APIs
Chunked uploads
Dynamic endpoints
A practical issue can occur where:
File is blocked ✓
       │
       ▼
DLP detection works ✓
       │
       ▼
Filename in log
may be inaccurate / incomplete
This has been observed especially with some cloud-based services.
The application may transfer the file using:
Custom API
     +
Encoded content
     +
Separate metadata
instead of a traditional HTTP file-upload mechanism.
When blocking files, prefer file type when the requirement is actually based on file type.
Block:
secret-file-final-v2.pdf
using only a filename pattern.
Block:
*.pdf
using the appropriate DLP file-type classification.
Reason: Cloud applications may change how filenames and metadata are exchanged through their APIs.
FortiGate DLP can use several detection mechanisms:
                 DLP Detection
                       │
       ┌───────────────┼────────────────┐
       ▼               ▼                ▼
   Data Types      Dictionaries      Sensors
       │               │                │
       ├── Keyword     ├── Keyword      └── Logic
       ├── Regex       └── Patterns
       ├── Hex
       ├── CC
       └── SSN
                       │
                       ▼
                  Fingerprinting
                       │
                       ▼
                   Watermarking
DLP Fingerprinting can detect known sensitive documents by comparing fingerprints generated from files.
Conceptually:
Sensitive File
      │
      ▼
FortiGate
      │
      ▼
Generate Fingerprint
      │
      ▼
Store Fingerprint
      │
      ▼
Future Network Traffic
      │
      ▼
Generate Fingerprint
      │
      ▼
Compare
      │
   ┌──┴──┐
   ▼     ▼
 Match  No Match
   │       │
   ▼       ▼
 Action   Continue
| Method | Detects | 
|---|---|
| Keyword | Specific words/phrases | 
| Regex | Structured patterns | 
| Dictionary | Groups of matching terms | 
| Sensor | Matching logic | 
| Fingerprint | Known document/file content | 
| Watermark | Files/data containing recognized watermark information | 
Keyword
   ↓
"What does the content say?"
Fingerprint
   ↓
"Is this the known document?"
A document fingerprint source can be configured from a file repository.
Example source:
SMB Server
    │
    ▼
FortiGate
    │
    ▼
Document Source
    │
    ▼
Fingerprint Database
Important: The document fingerprint feature requires a FortiGate with internal storage.
Example configuration:
config dlp fp-doc-source
    edit "dlp-doc-test"
        set server-type smb
        set server 192.168.20.200
        set period daily
        set vdom root
        set scan-subdirectories enable
        set remove-deleted disable
        set keep-modified enable
        set username "<SERVICE_ACCOUNT>"
        set password "<SECRET>"
        set file-path "c:/w.pdf"
        set file-pattern "w.pdf"
        set sensitivity critical
        set tod-hour 00
        set tod-min 00
    next
end
Never publish real credentials:
❌ set username admin
❌ set password RealPassword123
Use:
<USER>
<SECRET>
in documentation and GitHub repositories.
Conceptually:
Sensitivity
    │
    ├── Warning
    ├── Low
    ├── Private
    ├── Medium
    └── Critical
The exact available values and mapping depend on the FortiOS implementation.
Example:
set sensitivity critical
A fingerprint rule can define a match percentage.
Example:
config dlp profile
    edit "dlp-prof-test"
        config filter
            edit 1
                set proto http-get
                set filter-by fingerprint
                set sensitivity critical
                set match-percentage 40
                set action block
            next
        end
    next
end
Conceptually:
Captured File
      │
      ▼
Fingerprint Comparison
      │
      ▼
Match Percentage
      │
      ├── < 40%
