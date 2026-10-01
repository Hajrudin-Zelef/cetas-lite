---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security-0f55dd46-4
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46.md
source_anchor: ""
source_lines: [948, 1068]
sha256: e362b32f5c23068e4c42da5360b809f2327789f484bbcfb64684850be1f0955f
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-security--0f55dd46

          ┌───────────┼───────────┐
          ▼           ▼           ▼
         Allow       Block    Quarantine
                      │
                      ▼
                    Log
Flow
  ↓
DLP
  X
Use the appropriate proxy-based configuration when required by the DLP feature.
HTTPS
  ↓
Encrypted
  ↓
DLP
  X
Instead:
HTTPS
  ↓
SSL Deep Inspection
  ↓
DLP
  ✓
ALL TRAFFIC
     ↓
FULL DLP
     ↓
High Resource Usage
Prefer targeted inspection.
Especially with cloud applications:
Cloud API
   ↓
Custom Metadata
   ↓
Filename may not be reliable
When the goal is blocking a file type, use the appropriate file-type filter where possible.
Never commit:
username
password
API key
secret
private key
to GitHub.
Use placeholders:
<USERNAME>
<PASSWORD>
<SECRET>
DLP Deployment
│
├── [ ] Enable DLP feature
├── [ ] Enable proxy inspection
├── [ ] Configure SSL Deep Inspection where required
├── [ ] Define Data Types
├── [ ] Create Dictionaries
├── [ ] Configure Sensors
├── [ ] Create DLP Profile
├── [ ] Select protocols
├── [ ] Select file types
├── [ ] Define file-size limits
├── [ ] Select actions
├── [ ] Attach DLP to firewall policy
├── [ ] Enable required logging
├── [ ] Test HTTP
├── [ ] Test HTTPS
├── [ ] Test file upload
├── [ ] Test file download
├── [ ] Test cloud storage
├── [ ] Monitor CPU/memory
└── [ ] Validate false positives
Use a controlled test domain:
dlptest.com
Test flow:
Client
  │
  ▼
HTTPS
  │
  ▼
FortiGate
  │
  ├── SSL Inspection
  │
  ├── DLP
  │
  └── IPS
  │
  ▼
Test Destination
Test cases:
1. Keyword detection
2. Regex detection
3. Dictionary match
4. Sensor threshold
5. File-type filtering
6. File-size filtering
7. Block action
8. Log-only action
9. Fingerprint matching
10. HTTPS upload
- SSL Deep Inspection → Required when DLP must inspect encrypted HTTPS content
- Proxy-Based Inspection → Core processing model for advanced DLP
- IPS → Additional inspection layer
- Application Control → Control applications used for data transfer
- Web Filter → Control destination websites
- File Filtering → Restrict file types
- DLP Fingerprinting → Detect known sensitive documents
SheynShield Engineering Note
Don't think of FortiGate DLP as simply "block a keyword."
The real architecture is:
Data Type → Dictionary → Sensor → DLP Profile → Firewall Policy
And for encrypted web traffic:
HTTPS → SSL Deep Inspection → DLP → Action
For known sensitive documents:
Document Source → Fingerprint Database → Fingerprint Match → DLP Action
Finally, when troubleshooting cloud-storage DLP, separate detection accuracy from metadata/logging accuracy. A file can be successfully detected and blocked even when a cloud application's API causes the filename recorded in the log to be incomplete or inaccurate.
- 
YouTube — SheynShield 
  - Fortinet NSE content
  - FortiGate troubleshooting
  - Network Security Engineering
