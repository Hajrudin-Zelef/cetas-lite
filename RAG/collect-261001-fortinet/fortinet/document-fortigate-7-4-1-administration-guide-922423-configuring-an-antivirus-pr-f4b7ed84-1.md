---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-1-administration-guide-922423-configuring-an-antivirus-pr-f4b7ed84-1
title: "Configuring an antivirus profile"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "Microsoft"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-1-administration-guide-922423-configuring-an-antivirus-pr-f4b7ed84.md
source_anchor: ""
source_lines: [1, 60]
sha256: e5a43e7c60bc5f58e8ca8624335945ad28b7a3d2217a131746e9af6ffc0f28f7
---

# Configuring an antivirus profile

# Configuring an antivirus profile

In an antivirus profile, the FortiGate can be configured to apply antivirus protection to HTTP, FTP, IMAP, POP3, SMTP, CIFS, and NNTP sessions. Proxy-based profiles also support MAPI and SSH. Antivirus inspection prevents potentially unwanted and malicious files from entering the network. Antivirus profiles include multiple different functions, such as scanning files for virus signatures, scanning for advanced persistent threats, checking external malware hash lists and threat feeds, and others. Malicious files can be blocked or monitored, and can be quarantined. Some antivirus profile options require a license and/or other Fortinet products. Some antivirus profile options can only be configured in the CLI (refer to the FortiOS CLI Reference).

|  | The feature set setting (proxy or flow) in the antivirus profile must match the inspection mode setting (proxy or flow) in the associated firewall policy. For example, a flow-based antivirus profile must be used with a flow-based firewall policy. | 

###### To configure an antivirus profile:

1. 
                                                    Go to *Security Profiles > AntiVirus* and click*Create New* .
2. 
                                                    Configure the following settings: Name Enter a unique name for the profile. Comments Enter a comment (optional). AntiVirus scan Enable one or more protocols for inspection, then enable *AntiVirus scan* for the selected protocols with a specified action.
  - 
                                                                                *Block* : block the malicious traffic.
  - 
                                                                                *Monitor* : log malicious traffic and allow it to pass inspection.
 Feature set Select the feature set for the profile. The feature set mode must match the inspection mode used in the associated firewall policy. 
  - 
                                                                                *Flow-based*
  - 
                                                                                *Proxy-based*
 Additional options are available in proxy-based mode and are identified in the GUI with a *P* icon. See Inspection mode feature comparison for more details.If the *Feature set* option is not visible, enter the following in the CLI:```
config system settings
    set gui-proxy-inspection enable
end
```
Inspected Protocols Enable to inspect the protocol for session inspection: HTTP, SMTP, POP3, IMAP, FTP, and CIFS. Disabled protocols are not inspected. MAPI and SSH can be inspected in proxy-based mode. APT Protection Options This section includes options available with FortiGuard to mitigate advanced persistent threats (APT) in file-based attacks. Content Disarm and Reconstruction This option is available in proxy-based mode when at least one protocol is enabled for inspection and *AntiVirus scan* is enabled.Enable to allow the FortiGate to sanitize Microsoft Office documents and PDF files (including files in ZIP archives) by removing active content (disarm) without affecting the integrity of the textual content (reconstruction). See Content disarm and reconstruction for more details. Allow transmission when an error occurs Enable to allow traffic to pass when an inspection error occurs. Disable to block traffic when an inspection error occurs. Original File Destination Specify how to quarantine files processed by content disarm and reconstruction. 
  - 
                                                                                *FortiSandbox* : quarantine files on FortiSandbox. The FortiSandbox must be enabled. See Using FortiSandbox post-transfer scanning with antivirus for more details.
  - 
                                                                                *File Quarantine* : quarantine files on FortiGate models with a hard disk.
  - 
                                                                                *Discard* : discard suspicious files.
 Treat Windows executables in email attachments as viruses Enable to deem all Windows executable files located in email traffic as viruses. Send Files to FortiSandbox for Inspection Enable to send files to FortiSandbox for inspection. The FortiSandbox must be enabled. Scan strategy FortiSandbox scans files inline for flow-based mode ( *Inline* ) and after the file transfer is complete for proxy-based mode (*Post Transfer* ). See Using FortiSandbox inline scanning with antivirus and Using FortiSandbox post-transfer scanning with antivirus for more details.File types Specify which files to FortiSandbox for inspection. 
  - 
                                                                                *Suspicious Files Only* : only send suspicious files to FortiSandbox for inspection.
  - *All Supported Files* : send all supported files to FortiSandbox for inspection.
 Do not submit files matching types Click the *+* to exclude certain file types from being sent to FortiSandbox.Do not submit files matching file name patterns Click the *+* to enter a wildcard pattern to exclude files from being sent to FortiSandbox.Use FortiSandbox database Enable to use the signature database from FortiSandbox. The FortiSandbox must be enabled. Send files to FortiNDR for inspection This option is available in proxy-based mode when at least one protocol is enabled for inspection, *AntiVirus scan* is enabled, and FortiNDR is enabled. See Using FortiNDR inline scanning with antivirus for more details.Include mobile malware protection Enable to use the mobile malware protection database from FortiGuard for content scanning. Quarantine This option is available when at least one protocol is enabled for inspection and *AntiVirus scan* is enabled.Enable to quarantine infected files. See also Downloading quarantined files in archive format NEW. Virus Outbreak Prevention This section includes options available with the FortiGuard Virus Outbreak Protection Service. A license is required to use these options. See FortiGuard outbreak prevention for more details. Use FortiGuard outbreak prevention database Enable to use the outbreak prevention database that is available with Advanced Malware Protection on FortiGuard. A license is required. 
  - 
                                                                                *Block* : block the malicious traffic.
  - *Monitor* : log malicious traffic and allow it to pass inspection.
 Use external malware block list Enable to use one or more external blocklist file hashes. See Malware hash threat feed for more details. 
  - 
                                                                                *Block* : block the malicious traffic.
  - 
                                                                                *Monitor* : log malicious traffic and allow it to pass inspection.
  - 
                                                                                *All* : use all malware block lists.
  - 
                                                                                *Specify* : select specific malware block lists.
 Use EMS threat feed This option is available when at least one protocol is enabled for inspection and *AntiVirus scan* is enabled.Enable to use malware threat feeds from FortiClient EMS. A FortiClient EMS Fabric connector with EMS threat feed enabled is required. See Malware threat feed from EMS for more details.
3. 
                                                                                
4. 
                                                    Click *OK* .

## Protocol options

