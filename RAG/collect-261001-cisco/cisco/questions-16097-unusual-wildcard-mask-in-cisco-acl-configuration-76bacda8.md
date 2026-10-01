---
id: collect-261001-cisco/cisco/questions-16097-unusual-wildcard-mask-in-cisco-acl-configuration-76bacda8
title: "questions-16097-unusual-wildcard-mask-in-cisco-acl-configuration-76bacda8"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-16097-unusual-wildcard-mask-in-cisco-acl-configuration-76bacda8.md
source_anchor: ""
source_lines: [1, 11]
sha256: b067b4318537654476828a4102d7f46ed15e7b36df29411caacd7da9ee85c546
---

# questions-16097-unusual-wildcard-mask-in-cisco-acl-configuration-76bacda8

That is a valid wildcard mask.  It will match any IP with the format 10.(160-191).1.(0-255).  Whether or not this is intended, or good design, is unknown.
Wildcard masks are just "do-we-care" bitwise masks used when looking at an IP -- a value of 0 means "do-care", and 1 means "don't-care".
In this case, 0.31.0.255 translates to:
00000000 00011111 00000000 1111111
So the IP listed in the ACL (10.160.1.0) will have a binary value of 
00001010 10100000 00000001 00000000
Testing an IP (10.190.1.200):
00001010 10111110 00000001 11001000 (10.190.1.200)
00001010 10100000 00000001 00000000 (10.160.1.0)
^^^^^^^^ ^^^      ^^^^^^^^          (0.31.0.255)
The carets represent the wildcard mask, showing which bits MUST match.  Since they do, 10.190.1.200 would match this ACL statement.
