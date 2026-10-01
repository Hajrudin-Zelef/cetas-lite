---
id: collect-261001-cisco/cisco/questions-1929-asa-ftp-special-character-in-password-issue-b4bdcef0
title: "questions-1929-asa-ftp-special-character-in-password-issue-b4bdcef0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-cisco/questions-1929-asa-ftp-special-character-in-password-issue-b4bdcef0.md
source_anchor: ""
source_lines: [1, 17]
sha256: 2ef8d78b4482cbe9968893371e742a9c45e0655b7794a0cb0e86d7cd5dce1fd3
---

# questions-1929-asa-ftp-special-character-in-password-issue-b4bdcef0

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
9
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am trying to copy a show tech from an ASA to a remote server via FTP.
My issue is that the password I am providing has a '@' and it is being interpreted incorrectly by the CLI. I have tried putting a '\@' but the command still fails.
Anyone have any luck with this?
This works
show tech-support file ftp://username:password@xxx.xxx.xxx.xxx/ASA-SHOW-TECH
This Does not work
show tech-support file ftp://username:p@ssword@xxx.xxx.xxx.xxx/ASA-SHOW-TECH
The problem here is that the device is interpreting the first @ as the delimiter between the credentials and the FTP server address.
Perhaps configuring the username and password in global-configuration mode and then redirecting the show tech without specifying the credentials would do the trick. Otherwise, and I know this is a wimpy solution, changing the password to not include that one character may be the easiest and perhaps most secure way of working around this issue.
On IOS you can enter control-v to escape the next character. This lets you enter a question mark without pulling up help. Don't if that works on ASA but might be worth a try.
It does not seem like it is supported, so the work around is to just change the password that does not have the @ special character if you need to automate the process.
