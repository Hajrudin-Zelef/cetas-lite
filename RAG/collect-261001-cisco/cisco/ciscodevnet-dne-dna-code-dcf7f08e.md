---
id: collect-261001-cisco/cisco/ciscodevnet-dne-dna-code-dcf7f08e
title: "ciscodevnet-dne-dna-code-dcf7f08e"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/ciscodevnet-dne-dna-code-dcf7f08e.md
source_anchor: ""
source_lines: [1, 22]
sha256: 5f44b36df4053bb551bace9856951ae931c21232d2aa02da19f19350821c69b2
---

# ciscodevnet-dne-dna-code-dcf7f08e

This repository contains the sample code to go along with Cisco DevNet Learning Labs covering network programmability topics. During the setup steps of the labs, you'll be asked to clone this repository down to your workstation to get started.
Contributions are welcome, and we are glad to review changes through pull requests. See contributing.md for details.
Within this repository are several files and folders covering different topics and labs. This table provides details on what each is used for, and which labs they correspond to.
| File / Folder | Description | 
|---|---|
| env_lab.py | A Python file containing lab infrastructure details for routers, switches and appliances leveraged in the different labs.  This file provides a centralized  Python import that is used in  other code samples to retrieve IPs, usernames, and passwords for connections | 
| env_user.template | Similar to env_lab.py , this is a template for end users to copy within their own code repo asenv_user.py where they can provide unique details for their own accounts.  For example, their Webex authentication token.  Not all labs require this file, if one does it will be specified in setup. | 
| requirements.txt | Global Python requirements file containing the requirements for all labs within this repository.  Each folder also contains a local requirements.txt file. | 
| intro-python/ | Sample code and exercises for the Python Fundamentals Learning Labs | 
| rest-api/ | Sample code and exercises for the REST API Fundamentals Learning Labs | 
| intro-dnac/ | Sample code and exercises for the Cisco DNA Center Learning Labs | 
| intro-mdp/ | Sample code and exercises for the Introduction to Model Drive Programmability (NETCONF/RESTCONF/YANG) Learning Labs | 
| intro-guestshell/ | Sample code and exercises for the Introduction to Guest Shell on IOS XE Learning Labs | 
| intro-meraki/ | Sample code and exercises for the Introduction to Meraki Programmability Learning Labs | 
| intro-nfvis/ | Sample code and exercises for the Introduction to Network Function Virtualization with NFVIS APIs Learning Labs | 
| verify/ | A series of verification scripts primarily used during DevNet Express for DNA events to ensure the workshop environment is fully operational. | 
| dev/ | Resources and information for building code samples and labs. | 
| requirements-dev.txt | Python requirements file containing requirements only needed if developing new code samples. | 
Note: These code samples are also leveraged during DevNet Express for DNA events. If you are one of these events, your event proctors and hosts will walk you through event setup and verification steps as part of agenda.
These learning modules are for public consumption, so you must ensure that you have the rights to any content that you contribute.
- If you'd like to contribute to an existing lab, refer to contributing.md.
- If you're interested in creating a new Cisco DevNet Learning Lab, please contact a DevNet administrator for guidance.
