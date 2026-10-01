---
id: collect-261001-automatisation-infra/automatisation-infra/t5-network-access-control-user-restrictions-via-tac-td-p-906223-8516386e
title: "t5-network-access-control-user-restrictions-via-tac-td-p-906223-8516386e"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/t5-network-access-control-user-restrictions-via-tac-td-p-906223-8516386e.md
source_anchor: ""
source_lines: [1, 116]
sha256: 03de3c3edd3068c01ca3cd6a31b764cb9748ea3e514f4688a2f394e5225b24e1
---

# t5-network-access-control-user-restrictions-via-tac-td-p-906223-8516386e

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-25-2007 07:16 AM - edited 03-10-2019 03:28 PM
hi @all,
i'm trying to restrict a user with tacacs+. the relevant router & tac.-config are as following:
ios:
aaa new-model
!
!
aaa authentication login console group tacacs+ local enable
aaa authentication login vty group tacacs+ local enable
aaa authentication enable default group tacacs+ enable
aaa authorization commands 0 en0 group tacacs+
aaa authorization commands 5 RESTRICT group tacacs+
!
aaa session-id common
!
line con 0
exec-timeout 0 0
logging synchronous
login authentication console
line aux 0
line vty 0 903
authorization commands 0 en0
authorization commands 5 RESTRICT
logging synchronous
login authentication vty
transport input ssh
tacacs:
user = guck {
login = cleartext guck
service = shell { priv_level = 5 }
cmd = enable { deny .* }
cmd = show { permit ver deny .* }
cmd = traceroute { permit .* }
cmd = exit { permit .* }
}
it's partially working, so i can't execute the enable command, but i can do a lot more than "show ver" as intended, and more than traceroute and exit. i can execute ping as well and various other commands. now i'd like to know if it's possible at all to restrict a user to the above mentioned commands in conjunction with tacacs, or doesn't this work that way?
tia
br
erik
Solved! Go to Solution.
- Labels:
- 
						
							
		
			AAA
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-25-2007 03:33 PM
I think the reason that you are able to use ping command is, because "ping" command level is not being authorized.
i.e. by default on IOS device we have three level, namely 0, 1 and 15.
At level zero you have command: disable, enable, exit, help, and logout
I think ping command is either at level 1 or 15, given you have not changed the level of command.
So I would suggest following,
aaa authorization commands 0 AUTHO-VTY group tacacs+
aaa authorization commands 1 AUTHO-VTY group tacacs+
aaa authorization commands 15 AUTHO-VTY group tacacs+
line vty 0 903
no authorization commands 0 en0
no authorization commands 5 RESTRICT
authorization commands 0 AUTHO-VTY
authorization commands 1 AUTHO-VTY
authorization commands 15 AUTHO-VTY
And then configure permitted or denied commands accordingly on Tacacs server user profile.
Regards,
Prem
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-25-2007 03:33 PM
I think the reason that you are able to use ping command is, because "ping" command level is not being authorized.
i.e. by default on IOS device we have three level, namely 0, 1 and 15.
At level zero you have command: disable, enable, exit, help, and logout
I think ping command is either at level 1 or 15, given you have not changed the level of command.
So I would suggest following,
aaa authorization commands 0 AUTHO-VTY group tacacs+
aaa authorization commands 1 AUTHO-VTY group tacacs+
aaa authorization commands 15 AUTHO-VTY group tacacs+
line vty 0 903
no authorization commands 0 en0
no authorization commands 5 RESTRICT
authorization commands 0 AUTHO-VTY
authorization commands 1 AUTHO-VTY
authorization commands 15 AUTHO-VTY
And then configure permitted or denied commands accordingly on Tacacs server user profile.
Regards,
Prem
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-26-2007 12:15 AM
hi Prem,
thank you very much for your help! you hit the bull's-eye!!!!
br
erik
