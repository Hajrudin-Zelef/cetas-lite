---
id: collect-261001-huawei/huawei/resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd-3
title: "resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd"
domain: huawei
role: reference
task: reference
actors: ["CISA", "China", "Google"]
dates: ["2024-12-03"]
keywords: ["cost", "cyber", "cybersecurity", "exploit", "incident", "liability"]
source: docs/RAG/collect-261001-huawei/resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd.md
source_anchor: ""
source_lines: [80, 120]
sha256: 5d974dd848d86e6ee4de140b27c2bed276044baee1ca60e577c3046af3c447fb
---

# resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd

Organizations in the communications sector should be aware that the authoring agencies have observed Cisco-specific features often being targeted by, and associated with, these PRC cyber threat actors’ activity. To address the risk of exploitation by these specific threat actors, the authoring agencies urge organizations to apply the following hardening best practices to all Cisco operating systems. For additional information, see Cisco’s IOS XE Hardening Guide and Guide to Securing NX-OS Software Devices.
- Disable Cisco’s Smart Install service using no vstack .
- If not required, disable the guestshell access using guestshell disable for those versions which support the guestshell service.
- Disable all non-encrypted web management capabilities. If web management is required, configure servers in compliance with vendor recommended security settings and software images.
  - Always disable the underlying non-encrypted web server using no ip http server . If web management is not required, disable all of the underlying web servers usingno ip http server andno ip http secure-server .
- Always disable the underlying non-encrypted web server using 
- Disable telnet and ensure it is not available on any of the VTY lines by configuring all VTY stanzas with transport input ssh and transport output none .
- To securely store passwords on Cisco devices, organizations should:
  - Use Type-8 passwords when possible.
  - Avoid use of deprecated hashing or password types when storing passwords, such as Type-5 or Type-7.
  - If supported, secure the TACACS+ key as a Type-6 encrypted password.
Incident Reporting
- U.S. organizations: If suspicious activity is identified, contact your local FBI field office or the FBI’s Internet Crime Complaint Center (IC3). Cyber incidents can also be reported to CISA by calling 1-844-Say-CISA (1-844-729-2472), emailing contact@mail.cisa.dhs.gov, or reporting online at cisa.gov/report. For NSA client requirements or general cybersecurity inquiries, contact Cybersecurity_Requests@nsa.gov.
- Australian organizations: Visit cyber.gov.au or call 1300 292 371 (1300 CYBER 1) to report cybersecurity incidents and access alerts and advisories.
- Canadian organizations: Report incidents by emailing CCCS at contact@cyber.gc.ca.
- New Zealand organizations: Report cyber security incidents to incidents@ncsc.govt.nz or call 04 498 7654.
Secure by Design
The authoring agencies urge software manufacturers to incorporate secure by design principles into their software development lifecycle to strengthen the security posture of their customers. Software manufacturers should prioritize secure by design configurations to eliminate the need for customer implementation of hardening guidelines. Additionally, customers should demand that the software they purchase is secure by design. For more information on secure by design, see CISA’s Secure by Design webpage. Customers should refer to CISA’s Secure by Demand guidance for additional product security considerations.
Resources
- CISA: Cross-Sector Cybersecurity Performance Goals
- Joint Guide: Best Practices for Event Logging and Threat Detection
- NSA: Network Infrastructure Security Guide
- NSA, CISA, and FBI: People’s Republic of China State-Sponsored Cyber Actors Exploit Network Providers and Devices
- NSA: Hardening Network Devices
- NSA: Performing Out-of-Band Network Management
- NSA: Cisco Password Types: Best Practices
- NSA: Cisco Smart Install Protocol Misuse
- CCCS: Cryptographic Algorithms for UNCLASSIFIED, PROTECED A, and PROTECTED B Information – ITSP.40.111
- NIST: Special Publication 800-52: Guidelines for the Selection, Configuration, and Use of Transport Layer Security (TLS) Implementations
- NIST: Special Publication 800-77: Guide to IPsec VPNs
References
- CCCS: Guidance on Securely Configuring Network Protocols
- NSA: Network Infrastructure Security Guide
- CNSS: Committee on National Security Systems Policy (CNSSP)-15
Disclaimer
The authoring agencies do not endorse any commercial entity, product, company, or service, including any entities, products, or services linked within this document. Any reference to specific commercial entities, products, processes, or services by service mark, trademark, manufacturer, or otherwise, does not constitute or imply endorsement, recommendation, or favoring by the authoring agencies. Additionally, the information in this document is provided “as-is” and without warranties or representations of any kind. The users of this information shall have no recourse against the authoring parties for any loss, liability, damage or cost that may be suffered or incurred at any time arising from the use of information in this document, including but not limited to loss of data or interruption of business.
Acknowledgements
Cisco and Google Cloud Security contributed to this guidance.
Version History
December 3, 2024: Initial version.
Please share your thoughts with us via our anonymous product survey; we welcome your feedback.
