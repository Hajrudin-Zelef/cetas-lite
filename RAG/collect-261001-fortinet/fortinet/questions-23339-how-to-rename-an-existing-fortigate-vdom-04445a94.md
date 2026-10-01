---
id: collect-261001-fortinet/fortinet/questions-23339-how-to-rename-an-existing-fortigate-vdom-04445a94
title: "questions-23339-how-to-rename-an-existing-fortigate-vdom-04445a94"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2017-07-21"]
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-23339-how-to-rename-an-existing-fortigate-vdom-04445a94.md
source_anchor: ""
source_lines: [1, 11]
sha256: 2c10f78c182087efd0a6bdc113caf8849091f6e5277342fed1a674f2c4de6f62
---

# questions-23339-how-to-rename-an-existing-fortigate-vdom-04445a94

I need to rename an existing Fortigate 100D VDOM. I cannot find the command unless it's really counter-intuitive. Furthermore, if this is not possible, can I copy the existing VDOM to a new one with the correct name?
- 
        Did the answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer.Ron Maupin– Ron Maupin ♦2017-07-21 13:01:13 +00:00Commented Jul 21, 2017 at 13:01
1 Answer 1
Be forewarned, you will probably need to reboot the Fortigate for the changes to take place.
- Backup your entire config (not just the vdom)
- Open the config in a text editor
- Do a Search and replace of the old vdom name with the new name.
- Save the config under a new name, so that you have a backup of the old config.
- Restore the new config.
- Your vdoms should now be renamed.
