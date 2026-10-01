---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-940102-how-to-reset-ubiquiti-unifi-controller-admin-password-f46eb8e5
title: "questions-940102-how-to-reset-ubiquiti-unifi-controller-admin-password-f46eb8e5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-940102-how-to-reset-ubiquiti-unifi-controller-admin-password-f46eb8e5.md
source_anchor: ""
source_lines: [1, 16]
sha256: 5f2b142db282e9828d615178b8f41080a3647cc58851d2065f0d987a7ce182dc
---

# questions-940102-how-to-reset-ubiquiti-unifi-controller-admin-password-f46eb8e5

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I need to get into a Ubiquiti UniFi Controller system that was setup by one of my predecessors. The admin password has long since been lost and the company now uses a different Wi-Fi system at the other offices, so there's no support contract (although I don't think Ubiquiti offers them anyhow). Also, it's a 32-bit Windows 7 laptop, so I can't use RoboMongo (64-bit only). This is at a remote office with no on-premise IT staffer, so I have to just take what little help I can get in the way of physical access.
The steps to resolve this for a UniFi controller hosted on 32-bit Windows were:
Install MongoDB (had to be 32-bit in my case, but most people will have this on a 64-bit OS)
Open elevated CMD prompt: CD "C:\Program Files\MongoDB\Server\3.2\bin"
(my version was 3.2 but yours may vary)
Start UniFi Controller
Run this to drill into the MongoDB database: .\mongo —-port 27117
Then change to the "ace" database: use ace
Run this query to find your admin, email, hashes, etc (outputs in JSON . db.admin.find().forEach(printjson);
Finally, run this command to change the new password to the SHA512 hashed & salted value of "password" (no quotes) db.admin.update( { name: "admin" }, {$set: { x_shadow: "$6$9Ter1EZ9$lSt6/tkoPguHqsDK0mXmUsZ1WE2qCM4m9AQ.x9/eVNJxws.hAxt2Pe8oA9TFB7LPBgzaHBcAfKFoLpRQlpBiX1" } } );
