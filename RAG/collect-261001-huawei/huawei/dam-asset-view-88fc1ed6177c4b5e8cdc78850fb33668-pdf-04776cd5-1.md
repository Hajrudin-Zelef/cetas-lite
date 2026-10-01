---
id: collect-261001-huawei/huawei/dam-asset-view-88fc1ed6177c4b5e8cdc78850fb33668-pdf-04776cd5-1
title: "dam-asset-view-88fc1ed6177c4b5e8cdc78850fb33668-pdf-04776cd5"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: []
keywords: ["copyright", "cyber", "embedding"]
source: docs/RAG/collect-261001-huawei/dam-asset-view-88fc1ed6177c4b5e8cdc78850fb33668-pdf-04776cd5.md
source_anchor: ""
source_lines: [1, 131]
sha256: d92703fc47cdca773957f639bc05dde2dcebc1c1b16ddc774f7ef42b7888d90a
---

# dam-asset-view-88fc1ed6177c4b5e8cdc78850fb33668-pdf-04776cd5

Huawei Product Security Baseline
Huawei Practices for Managing Cyber Security
HUAWEI TECHNOLOGIES CO., LTD.
Huawei Industrial Base
Bantian Longgang
Shenzhen 518129, P.R. China
Tel: +86-755-28780808
www.huawei.com
Copyright©Huawei Technologies Co., Ltd. 2021. All rights reserved.
No part of this document may be reproduced or transmitted in any form or by any means without prior written consent of Huawei Technologies Co., Ltd.
, HUAWEI, and         are trademarks or registered trademarks of Huawei Technologies Co., Ltd.
Trademark Notice
Other trademarks, product, service and company names mentioned are the property of their respective owners. 
General Disclaimer
The information in this document may contain predictive statements including, without limitation, statements regarding the future 
financial and operating results, future product portfolio, new technology, etc. There are a number of factors that could cause actual 
results and developments to differ materially from those expressed or implied in the predictive statements. Therefore, such information 
is provided for reference purpose only and constitutes neither an offer nor an acceptance. Huawei may change the information at any 
time without notice.

8. Summary
7. Baseline Overview
· Result-based
· Universal
· Apply to all
· Continuously optimized
1. Foreword
2. Executive Summary
3. Product Security Is Critical to Cyber Security
4. Security, a Fundamental Capability Built into Products
5. Baseline Application in Huawei Business Processes
6. Principles for Baseline Development and Update
Protecting users' communication content
Protecting user privacy
Backdoor prevention 
Prevention of malware and malicious behavior
Access channel control
System hardening 
Application security
Encryption
Sensitive data protection
Management and maintenance security
Secure boot and integrity protection
Security documentation
Secure coding
Secure compilation
Lifecycle management
01
02
03
04
04
08
08
08
08
08
09
10
10
10
10
10
10
11
11
11
11
11
11
12
12
12
12
Contents

1. Foreword
In 2020, the COVID-19 pandemic changed the way we live and how organizations operate. Many activities have gone 
online, and telecommuting, video conferencing, distance education, and telemedicine have become the new normal. In 
this context, digital technology has played an irreplaceable role in keeping our lives on track and our businesses open. At 
the same time, as digital transformation picks up speed, we see growing challenges relating to cyber security and privacy 
protection. We have witnessed a record number and scale of security vulnerabilities and cyber attacks around the world, 
with persistent occurrences of ransomware and data breaches. In a digital, intelligent world empowered by 5G, cloud, and 
AI, a secure and stable cyberspace is critical to securing people's livelihoods and protecting the vital public and economic 
functions of any society. It is clear that cyber security and privacy protection are becoming the inherent requirements and 
basic core capabilities in a digital world.
Cyber security is an opportunity for us to promote security and digitization through enhanced cooperation. All 
stakeholders in the digital space — including governments and regulators, industry and standards organizations, 
communication service and technology providers, and digital service providers — share a joint responsibility to address 
cyberspace challenges and improve the level of cyber security. In the case of a cyber attack, every product and service in 
the cyberspace is exposed to the attack, independent of the supplier. Only by strengthening cyber security protection for 
the end-to-end supply chain can we more effectively reduce the security risks across the entire network.
As a leading global provider of ICT infrastructure and smart devices, we offer a broad array of products and solutions that 
apply to diverse scenarios. Effectively managing the security of these products is a huge but surmountable challenge. In 
order to do so, Huawei has incorporated cyber security management requirements into all its business processes, 
especially the Integrated Product Development (IPD) process. By integrating product cyber security requirements 
throughout the planning, design, development, veriﬁcation, launch, and lifecycle management of the entire IPD process, 
Huawei is able to ensure that every product and version it releases delivers the expected level of quality in terms of cyber 
security. The applied management requirements and engineering and technical speciﬁcations include the Huawei Product 
Security Baseline ("the Baseline"), which contains mandatory requirements that Huawei products must meet. Drawing on 
our extensive experience over the past more than 10 years in product security quality, we continuously update and 
optimize the Baseline and share its value in ensuring the end-to-end security of the supply chain with our partners, 
suppliers, and others. Practices show that the Baseline applies not only to Huawei products but also to its entire supply 
chain. Our security practices over the past more than 10 years also demonstrate that the Baseline is an effective way to 
manage the quality of product security. The Baseline has ensured a stellar security record of Huawei products on 
customer networks.
2. Executive Summary
We believe that:
• Improving product security is key to mitigating risks of the cyber security incidents that occur frequently worldwide.
• Embedding security management into the product development process and making cyber security a fundamental 
capability of products is a fundamental approach to resolving cyber security issues.
• Developing and implementing a Baseline of common product security requirements ensures that all products meet the 
same fundamental requirements in terms of the security quality, and the security quality continuously improves as the 
Baseline is updated.
• Huawei's end-to-end cyber security framework integrates the Baseline into the product development process as a 
fundamental security requirement. The Baseline and various quality assurance activities are strictly implemented in order 
to ensure product security quality and prevent security incidents.
Huawei Product Security Baseline:
• Huawei has developed the result-based, universal, applicable-to-all, and continuously optimized Baseline, which is 
effective, implementable, and veriﬁable, and continuously improves the security quality of Huawei products.
• Huawei has developed the Baseline based on common and critical security requirements identiﬁed through its study of 
applicable laws and regulations as well as its deep understanding of legal and regulatory requirements, customers' 
business requirements, industry best practices, known issues, and more. The Baseline consists of 54 requirements under 
15 categories and 112 entries for implementation guidance and interpretation.
01 02
Foreword Executive Summary
NOTE: The Huawei Product Security Baseline describes the requirements that must be met in order to ensure Huawei products deliver the expected security 
capabilities. The requirements in the Baseline are a subset of all product security requirements.
Over the past more than 10 years, Huawei has been continuously updating the Baseline in order to address existing and emerging threats. Each update 
applies only to products developed after the update is released, unless otherwise speciﬁed.

