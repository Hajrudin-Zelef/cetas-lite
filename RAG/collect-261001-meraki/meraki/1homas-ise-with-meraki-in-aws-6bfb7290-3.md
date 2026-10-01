---
id: collect-261001-meraki/meraki/1homas-ise-with-meraki-in-aws-6bfb7290-3
title: "1homas-ise-with-meraki-in-aws-6bfb7290"
domain: meraki
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "license", "mit license"]
source: docs/RAG/collect-261001-meraki/1homas-ise-with-meraki-in-aws-6bfb7290.md
source_anchor: ""
source_lines: [286, 301]
sha256: 5365e29ef4447402c59d9594f727f81d2ec0f0e89d24f8a99eef284b386fa5be
---

# 1homas-ise-with-meraki-in-aws-6bfb7290

1. In the AWS Console, navigate to **CloudFormation > Stacks**
2. Select `ISE-3-1-518`
3. Choose Create Stack > With New Resources (standard)
4. For Create Stack, choose `Upload a template file` , choose your file (`ISE-3-1-518.CFT.yaml` ) and click`Next`
5. choose a unique Stack Name (`ISE-3-1-518` )

SSH to the ISE console

`ssh -i "~/.ssh/ISEinAWS.pem" admin@{ hostname | IP }`
- Installing Ansible for all platforms
- Documentation for Ansible collections:
- Cisco Meraki vMX en AWS - YouTube video for vMX in AWS (Spanish)
- AWS re:Invent 2019: AWS Networking Fundamentals (NET201-R2) provides an excellent overview of the networking elements that are automatically created when using the Launch Instance wizard
- Best practices for managing AWS access keys

This repository is licensed under the MIT License.
