---
id: collect-260926-mikrotik/mikrotik/questions-657957-mikrotik-and-freeradius-user-time-limit-48219392
title: "questions-657957-mikrotik-and-freeradius-user-time-limit-48219392"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/qos/questions-657957-mikrotik-and-freeradius-user-time-limit-48219392.md
source_anchor: ""
source_lines: [1, 13]
sha256: cbf0765fa1f376fac0a4b541cbb0fa15b16fddf5fd4227673e409c109a4a780c
---

# questions-657957-mikrotik-and-freeradius-user-time-limit-48219392

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am using freeradius in order to authenticate users to access the internet over a mikrotik routerboard.
I would like to set up a maximum time limit in a default period of time in the mysql database. For instance I would like to accept a user to have a total connection time of 7 hours over a week period of time. What values and in which field of the database should I alternate.
The reference guide that I used for my implementation is here
I have tried playing around with various Variables in the group reply table of the database, and I got the reply on the test, but probably I am using wrong variables, for instance the Session-Time = 60 just resets the session every 60 seconds (just for testing) and not disconnecting the user.
A Session-Time will set the maximum time a session may take. This does not account for other rules like a maximum time over a day or a week. You need more logic for this than an attribute in your RADIUS reply.
A stateful storage is needed to keep track of the used time for a user. RADIUS accounting is sufficient for this purpose. For example, when storing accounting data in MySQL you can query the already used session time for a period to calculate a new Session-Time for the upcoming session.
FreeRADIUS has modules for this purpose: sqlcounter and counter. The documentation covers examples of implementation.
