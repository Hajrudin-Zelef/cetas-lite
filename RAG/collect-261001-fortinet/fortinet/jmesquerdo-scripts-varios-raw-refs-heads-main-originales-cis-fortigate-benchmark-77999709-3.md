---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-3
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["benchmark", "benchmarks", "cyber", "cybersecurity", "reasoning", "research"]
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [100, 205]
sha256: 9cbdb15a29b1d0fbcacb2c63516517d23a889f13b5552f9b065a04c95289ef31
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 4 
6 VPN ................................ ................................ ................................ ................................ ... 104 
6.1 SSL VPN ........................................................................................................................................ 105 
6.1.1 Apply a Trusted Signed Certificate for VPN Portal (Manual) ............................................................ 106 
6.1.2 Enable Limited TLS Versions for SSL VPN (Manual)....................................................................... 107 
7 Users and Authentication ................................ ................................ ...............................  108 
7.1 Configuring the maximum login attempts and lockout period (Automated) ......................................... 109 
8 Logs and Reports ................................ ................................ ................................ ............ 109 
8.1 Enable Logging ............................................................................................................................ 110 
8.1.1 Enable Event Logging (Automated) ................................................................................................. 111 
8.2 Encrypt Logs Sent to FortiAnalyzer / FortiManager ................................................................. 112 
8.2.1 Encrypt Log Transmission to FortiAnalyzer / FortiManager (Automated) ......................................... 113 
8.3 Centralized Logging and Reporting ........................................................................................... 114 
8.3.1 Centralized Logging and Reporting (Automated) ............................................................................. 115 
Appendix: Summary Table ....................................................................................... 116 
Appendix: Change History ....................................................................................... 131

Page 5 
Overview 
All CIS Benchmarks focus on technical configuration settings used to maintain and/or 
increase the security of the addressed technology, and they should be used in 
conjunction with other essential cyber hygiene tasks like: 
• Monitoring the base operating system for vulnerabilities and quickly updating with 
the latest security patches  
• Monitoring applications and libraries for vulnerabilities and quickly updating with 
the latest security patches 
In the end, the CIS Benchmarks are designed as a key component of a comprehensive 
cybersecurity program.  
 
This document provides prescriptive guidance for establishing a secure configuration 
posture for Fortinet FortiGate devices running the Forinet OS version 6.4 or above. This 
guide was tested against FortiOS 6.4.5. To obtain the latest version of this guide, please 
visit http://benchmarks.cisecurity.org. If you have questions, comments, or have 
identified ways to improve this guide, please write us at feedback@cisecurity.org. 
 
Intended Audience 
This benchmark is intended for security administrators, IT auditors, and platform 
deployment personnel who plan to develop, deploy, assess, or secure solutions that 
incorporate Fortinet OS on Fortinet network devices.

Page 6 
Consensus Guidance 
This CIS Benchmark was created using a consensus review process comprised of a 
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

Page 7 
Typographical Conventions 
The following typographical conventions are used throughout this guide: 
Convention Meaning 
Stylized Monospace font 
Used for blocks of code, command, and script 
examples. Text should be interpreted exactly as 
presented. 
Monospace font Used for inline code, commands, or examples. 
Text should be interpreted exactly as presented.  
<italic font in brackets> Italic texts set in angle brackets denote a variable 
requiring substitution for a real value. 
Italic font Used to denote the title of a book, article, or other 
publication. 
Note Additional information or caveats

Page 8 
Recommendation Definitions 
The following defines the various components included in a CIS recommendation as 
applicable.  If any of the components are not applicable it will be noted or the 
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

