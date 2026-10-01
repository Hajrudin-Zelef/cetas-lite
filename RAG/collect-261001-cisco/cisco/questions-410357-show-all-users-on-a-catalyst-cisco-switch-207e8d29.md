---
id: collect-261001-cisco/cisco/questions-410357-show-all-users-on-a-catalyst-cisco-switch-207e8d29
title: "show all users on a Catalyst Cisco switch"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-410357-show-all-users-on-a-catalyst-cisco-switch-207e8d29.md
source_anchor: ""
source_lines: [1, 40]
sha256: 29d846496ca1f958b5534e0564f469db1c1eadd5b691f3c433eb423fff1b7fce
---

# show all users on a Catalyst Cisco switch

*Score : 8 | Source : https://serverfault.com/questions/410357/show-all-users-on-a-catalyst-cisco-switch*

I am new to Cisco, I am having some difficulty:
I'd like to list all user accounts. show users only displays currently logged in users.
I have no problem changing the enable password, but I'd like to see all available users so I can change specific user passwords as well.
Using 3750, 3560 switches and 55/10-20 ASAs.

---

### Reponse (acceptee) — score 11

Q1. How do I list all user accounts?
From the enable prompt, run show run | i username...
CORE01.PUB.DAL01#sh run | i user
username operator password 7 <someHashedPassword>
CORE01.PUB.DAL01#
  Q2. How would I reset the password for a specific user?
Change the password from configuration mode
CORE01.PUB.DAL01#conf t
CORE01.PUB.DAL01(config)#username <someuser> password 0 <somepassword>
Syntax is slightly different for an ASA...
mpenning-fw(config)# user <someuser> password <somepassword>

---

### Reponse — score 3

local users in Ciso IOS are listed in the running-config with the "username".
For your switches type "show run | b username" and look at the users listed there.
For the ASA it's a little bit easier, just type "show run username".
If the users are not local (radius, etc.) then you'll need to look on that server for the user list.

---

### Reponse — score 0

To list all users: "show running-config | include username"
If usernames are using secrets command to reset user secret: "username %username% secret 0 %secret%"
