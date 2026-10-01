---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-4
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["benchmark", "benchmarks", "cost", "cyber", "cybersecurity"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [180, 311]
sha256: bbeed7752700f6bece083a590b2deef7c2e702c618efb78a70dac55dbce018b3
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 6 
Internal Only - General 
Overview 
All CIS Benchmarks™ (Benchmarks) focus on technical configuration settings used to 
maintain and/or increase the security of the addressed technology, and they should be 
used in conjunction with other essential cyber hygiene tasks like: 
• Monitoring the base operating system and applications for vulnerabilities and 
quickly updating with the latest security patches. 
• End-point protection (Antivirus software, Endpoint Detection and Response 
(EDR), etc.). 
• Logging and monitoring user and system activity. 
In the end, the Benchmarks are designed to be a key component of a comprehensive 
cybersecurity program.  
Important Usage Information 
All Benchmarks are available free for non-commercial use from the CIS Website. They 
can be used to manually assess and remediate systems and applications. In lieu of 
manual assessment and remediation, there are several tools available to assist with 
assessment: 
• CIS Configuration Assessment Tool (CIS-CAT® Pro Assessor) 
• CIS Benchmarks™ Certified 3rd Party Tooling 
These tools make the hardening process much more scalable for large numbers of 
systems and applications.  
NOTE:  Some tooling focuses only on the Benchmark Recommendations that can 
be fully automated (skipping ones marked Manual). It is important that ALL 
Recommendations (Automated and Manual) be addressed since all are 
important for properly securing systems and are typically in scope for 
audits.  
Key Stakeholders 
Cybersecurity is a collaborative effort, and cross functional cooperation is imperative 
within an organization to discuss, test, and deploy Benchmarks in an effective and 
efficient way. The Benchmarks are developed to be best practice configuration 
guidelines applicable to a wide range of use cases. In some organizations, exceptions 
to specific Recommendations will be needed, and this team should work to prioritize the 
problematic Recommendations based on several factors like risk, time, cost, and labor. 
These exceptions should be properly categorized and documented for auditing 
purposes.

Page 7 
Internal Only - General 
 
Apply the Correct Version of a Benchmark 
Benchmarks are developed and tested for a specific set of products and versions and 
applying an incorrect Benchmark to a system can cause the resulting pass/fail score to 
be incorrect. This is due to the assessment of settings that do not apply to the target 
systems. To assure the correct Benchmark is being assessed:  
• Deploy the Benchmark applicable to the way settings are managed in the 
environment: An example of this is the Microsoft Windows family of 
Benchmarks, which have separate Benchmarks for Group Policy, Intune, and 
Stand-alone systems based upon how system management is deployed. 
Applying the wrong Benchmark in this case will give invalid results. 
• Use the most recent version of a Benchmark: This is true for all Benchmarks, 
but especially true for cloud technologies. Cloud technologies change frequently 
and using an older version of a Benchmark may have invalid methods for 
auditing and remediation. 
Exceptions 
The guidance items in the Benchmarks are called recommendations and not 
requirements, and exceptions to some of them are expected and acceptable. The 
Benchmarks strive to be a secure baseline, or starting point, for a specific technology, 
with known issues identified during Benchmark development are documented in the 
Impact section of each Recommendation. In addition, organizational, system specific 
requirements, or local site policy may require changes as well, or an exception to a 
Recommendation or group of Recommendations (e.g. A Benchmark could Recommend 
that a Web server not be installed on the system, but if a system's primary purpose is to 
function as a Webserver, there should be a documented exception to this 
Recommendation for that specific server). 
In the end, exceptions to some Benchmark Recommendations are common and 
acceptable, and should be handled as follows: 
• The reasons for the exception should be reviewed cross-functionally and be well 
documented for audit purposes. 
• A plan should be developed for mitigating, or eliminating, the exception in the 
future, if applicable. 
• If the organization decides to accept the risk of this exception (not work toward 
mitigation or elimination), this should be documented for audit purposes. 
It is the responsibility of the organization to determine their overall security policy, and 
which settings are applicable to their unique needs based on the overall risk profile for 
the organization.

Page 8 
Internal Only - General 
Remediation 
CIS has developed Build Kits for many technologies to assist in the automation of 
hardening systems. Build Kits are designed to correspond to Benchmark's 
“Remediation” section, which provides the manual remediation steps necessary to make 
that Recommendation compliant to the Benchmark. 
When remediating systems (changing configuration settings on 
deployed systems as per the Benchmark's Recommendations), 
please approach this with caution and test thoroughly. 
The following is a reasonable remediation approach to follow: 
• CIS Build Kits, or internally developed remediation methods should never be 
applied to production systems without proper testing. 
• Proper testing consists of the following: 
o Understand the configuration (including installed applications) of the targeted 
systems. Various parts of the organization may need different configurations 
(e.g., software developers vs standard office workers). 
o Read the Impact section of the given Recommendation to help determine if 
there might be an issue with the targeted systems. 
o Test the configuration changes with representative lab system(s). If issues 
arise during testing, they can be resolved prior to deploying to any production 
systems. 
o When testing is complete, initially deploy to a small sub-set of production 
systems and monitor closely for issues. If there are issues, they can be 
resolved prior to deploying more broadly. 
o When the initial deployment above is completes successfully, iteratively 
deploy to additional systems and monitor closely for issues. Repeat this 
process until the full deployment is complete.  
Summary 
Using the Benchmarks Certified tools, working as a team with key stakeholders, being 
selective with exceptions, and being careful with remediation deployment, it is possible 
to harden large numbers of deployed systems in a cost effective, efficient, and safe 
manner. 
NOTE: As previously stated, the PDF versions of the CIS Benchmarks™ are 
available for free, non-commercial use on the CIS Website. All other formats 
of the CIS Benchmarks™ (MS Word, Excel, and Build Kits) are available for 
CIS SecureSuite® members. 
CIS-CAT® Pro is also available to CIS SecureSuite® members.

Page 9 
Internal Only - General 
Target Technology Details 
This document, Security Configuration Benchmark for Cisco IOS, provides prescriptive 
guidance for establishing a secure configuration posture for Cisco Router running Cisco 
IOS version 17.06. This guide was tested against Cisco IOS 17 XE. To obtain the latest 
version of this guide, please visit http://benchmarks.cisecurity.org. If you have 
questions, comments, or have identified ways to improve this guide, please write us at 
benchmarkinfo@cisecurity.org. 
 
Intended Audience 
This benchmark is intended for system and application administrators, security 
specialists, auditors, help desk, and platform deployment personnel who plan to 
develop, deploy, assess, or secure solutions that incorporate Cisco IOS on a Cisco 
routing and switching platforms.

