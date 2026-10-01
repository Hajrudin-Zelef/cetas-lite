---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-5
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "benchmarks", "reasoning", "research", "safeguards"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [312, 471]
sha256: b03b0ea5eafa7aed8135fa297a83780d90cf27414db463c7f0835bb581523ea5
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 10 
Internal Only - General 
Consensus Guidance 
This CIS Benchmark™ was created using a consensus review process comprised of a 
global community of subject matter experts. The process combines real world 
experience with data-based information to create technology specific guidance to assist 
users to secure their environments. Consensus participants provide perspective from a 
diverse set of backgrounds including consulting, software development, audit and 
compliance, security research, operations, government, and legal.  
Each CIS Benchmark undergoes two phases of consensus review. The first phase 
occurs during initial Benchmark development. During this phase, subject matter experts 
convene to discuss, create, and test working drafts of the Benchmark. This discussion 
occurs until consensus has been reached on Benchmark recommendations. The 
second phase begins after the Benchmark has been published. During this phase, all 
feedback provided by the Internet community is reviewed by the consensus team for 
incorporation in the Benchmark. If you are interested in participating in the consensus 
process, please visit https://workbench.cisecurity.org/.

Page 11 
Internal Only - General 
Typographical Conventions 
The following typographical conventions are used throughout this guide: 
Convention Meaning 
Stylized Monospace font 
Used for blocks of code, command, and 
script examples. Text should be interpreted 
exactly as presented. 
Monospace font 
Used for inline code, commands, UI/Menu 
selections or examples. Text should be 
interpreted exactly as presented. 
<Monospace font in brackets> Text set in angle brackets denote a variable 
requiring substitution for a real value. 
Italic font 
Used to reference other relevant settings, 
CIS Benchmarks and/or Benchmark 
Communities. Also, used to denote the title 
of a book, article, or other publication. 
Bold font 
Additional information or caveats things like 
Notes, Warnings, or Cautions (usually just 
the word itself and the rest of the text 
normal).

Page 12 
Internal Only - General 
Recommendation Definitions 
The following defines the various components included in a CIS recommendation as 
applicable.  If any of the components are not applicable it will be noted, or the 
component will not be included in the recommendation.    
Title 
Concise description for the recommendation's intended configuration.  
Assessment Status 
An assessment status is included for every recommendation. The assessment status 
indicates whether the given recommendation can be automated or requires manual 
steps to implement. Both statuses are equally important and are determined and 
supported as defined below:  
Automated 
Represents recommendations for which assessment of a technical control can be fully 
automated and validated to a pass/fail state. Recommendations will include the 
necessary information to implement automation. 
Manual 
Represents recommendations for which assessment of a technical control cannot be 
fully automated and requires all or some manual steps to validate that the configured 
state is set as expected. The expected state can vary depending on the environment. 
Profile 
A collection of recommendations for securing a technology or a supporting platform. 
Most benchmarks include at least a Level 1 and Level 2 Profile. Level 2 extends Level 1 
recommendations and is not a standalone profile. The Profile Definitions section in the 
benchmark provides the definitions as they pertain to the recommendations included for 
the technology.  
Description 
Detailed information pertaining to the setting with which the recommendation is 
concerned. In some cases, the description will include the recommended value. 
Rationale Statement 
Detailed reasoning for the recommendation to provide the user a clear and concise 
understanding on the importance of the recommendation.

Page 13 
Internal Only - General 
Impact Statement  
Any security, functionality, or operational consequences that can result from following 
the recommendation. 
Audit Procedure  
Systematic instructions for determining if the target system complies with the 
recommendation. 
Remediation Procedure 
Systematic instructions for applying recommendations to the target system to bring it 
into compliance according to the recommendation. 
Default Value 
Default value for the given setting in this recommendation, if known. If not known, either 
not configured or not defined will be applied.  
References 
Additional documentation relative to the recommendation.  
CIS Critical Security Controls® (CIS Controls®) 
The mapping between a recommendation and the CIS Controls is organized by CIS 
Controls version, Safeguard, and Implementation Group (IG). The Benchmark in its 
entirety addresses the CIS Controls safeguards of (v7) “5.1 - Establish Secure 
Configurations” and (v8) '4.1 - Establish and Maintain a Secure Configuration Process” 
so individual recommendations will not be mapped to these safeguards. 
Additional Information  
Supplementary information that does not correspond to any other field but may be 
useful to the user.

Page 14 
Internal Only - General 
Profile Definitions  
The following configuration profiles are defined by this Benchmark: 
• Level 1 
Items in this profile intend to: 
o be practical and prudent; 
o provide a clear security benefit; and 
o not inhibit the utility of the technology beyond acceptable means. 
• Level 2 
This profile extends the "Level 1" profile. Items in this profile exhibit one or more 
of the following characteristics: 
o are intended for environments or use cases where security is paramount. 
o acts as defense in depth measure. 
o may negatively inhibit the utility or performance of the technology.

Page 15 
Internal Only - General 
 
 
Acknowledgements 
This Benchmark exemplifies the great things a community of users, vendors, and 
subject matter experts can accomplish through consensus collaboration. The CIS 
community thanks the entire consensus team with special recognition to the following 
individuals who contributed greatly to the creation of this guide: 
 
Contributor 
Raphael Precigout 
Eric Pinnell  
Gavin Day  
Vikas Gothwal CISSP 
David Poirier  
Curtis Starnes  
Josh Eblin  
Darren Stevenson  
 
Editor 
Darren Stevenson

Page 16 
Internal Only - General 
Recommendations 
1 Management Plane 
Services, settings and data streams related to setting up and examining the static 
configuration of the firewall, and the authentication and authorization of firewall 
administrators. Examples of management plane services include: administrative device 
access (telnet, ssh, http, and https), SNMP, and security protocols like RADIUS and 
TACACS+.

Page 17 
Internal Only - General 
1.1 Local Authentication, Authorization and Accounting (AAA) Rules 
Rules in the Local authentication, authorization and accounting (AAA) configuration 
class enforce device access control, provide a mechanism for tracking configuration 
changes, and enforcing security policy.

