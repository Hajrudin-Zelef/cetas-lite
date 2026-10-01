---
id: collect-261001-cisco/cisco/questions-33910-how-can-i-add-a-host-to-an-already-existing-object-group-in-asa-5aea6ba3
title: "questions-33910-how-can-i-add-a-host-to-an-already-existing-object-group-in-asa-5aea6ba3"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-33910-how-can-i-add-a-host-to-an-already-existing-object-group-in-asa-5aea6ba3.md
source_anchor: ""
source_lines: [1, 28]
sha256: c01ee54ab9f62755d7f643b1eed1f1ec2375bad269cbb1f22de9e47e52fe84ed
---

# questions-33910-how-can-i-add-a-host-to-an-already-existing-object-group-in-asa-5aea6ba3

A better way to create object groups is to fully embrace object orientated programing:
Create objects
object network test
 host 1.1.1.1
object network test2
 host 2.2.2.2
Add the object to a object-group
object-group network testing
 network-object object test
Use the object-group in a access list
access-list test_acl extended permit ip object test2 object-group testing
The access list will look like this:
ciscoasa# sh access-list test_acl
access-list test_acl; 1 elements; name hash: 0x71b1e4a4
access-list test_acl line 1 extended permit ip object test2 object-group testing (hitcnt=0) 0x4398ab6a
  access-list testing line 1 extended permit ip host 2.2.2.2 host 1.1.1.1 (hitcnt=0) 0xf5676f24
If you ever need to change the ip of the object "test" you only have to change it in one place instead of digging through object groups
Adding additional objects is easy, just create another object and add it to the group using the same steps outlined above.
ciscoasa(config)# object network test4
ciscoasa(config-network-object)# host 3.3.3.3
ciscoasa(config-network-object)# object-group network testing
ciscoasa(config-network-object-group)# network-object object test4
ciscoasa(config-network-object-group)# end
ciscoasa# sh access-list test_acl
access-list test_acl; 2 elements; name hash: 0x71b1e4a4
access-list test_acl line 1 extended permit ip object test2 object-group testing (hitcnt=0) 0x4398ab6a
  access-list test_acl line 1 extended permit ip host 2.2.2.2 host 1.1.1.1 (hitcnt=0) 0xf5676f24
  access-list test_acl line 1 extended permit ip host 2.2.2.2 host 3.3.3.3 (hitcnt=0) 0x0be1c62a
