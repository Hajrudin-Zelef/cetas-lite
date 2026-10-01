---
id: collect-261001-fortinet/fortinet/fortinetcloudcse-fortigate-automation-stitch-workshop-blob-head-content-04-5-tas-d7ef6db9
title: "fortinetcloudcse-fortigate-automation-stitch-workshop-blob-head-content-04-5-tas-d7ef6db9"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortinetcloudcse-fortigate-automation-stitch-workshop-blob-head-content-04-5-tas-d7ef6db9.md
source_anchor: ""
source_lines: [1, 62]
sha256: 992dad1adf3facd198a3967277dbe9ae107143348ea66e046ef8e9acd838d0a1
---

# fortinetcloudcse-fortigate-automation-stitch-workshop-blob-head-content-04-5-tas-d7ef6db9

| title | Task 4 - Automation Stitch | 
|---|---|
| chapter | true | 
| weight | 5 | 

A FortiGate Automation Stitch brings together a Trigger and one of more Actions.

Three FortiGate Automation Stitches need to be created, one for each trigger. All the stitches will utilize the same action. Each stitch is used to connect the trigger to one or more actions. The stitches created in this task each only have one action.

1. **Login** to the FortiGate using the IP address and credentials from the Terraform output.
2. **Click** through any opening screens for FortiGate setup actions, no changes are required.
3. **Click** the CLI Console
4. **Enter** the following CLI commands to create Automation Stitches to connect the triggers to the action.

- 
AppServers ```
config system automation-stitch
    edit "routetableupdate-AppServers"
        set description "Update route table for App Servers"
        set trigger "AppServer Existence"
        config actions
            edit 1
                set action "routetableupdate"
            next
        end
    next
end
```
- 
DbServers ```
config system automation-stitch
    edit "routetableupdate-DbServers"
        set description "Update route table for Db Servers"
        set trigger "DbServer Existence"
        config actions
            edit 1
                set action "routetableupdate"
            next
        end
    next
end
```
- 
WebServers ```
config system automation-stitch
    edit "routetableupdate-WebServers"
        set description "Update route table for Web Servers"
        set trigger "WebServer Existence"
        config actions
            edit 1
                set action "routetableupdate"
            next
        end
    next
end
```

1. View the configured Action in the FortiGate UI

1. View **routetableupdate-WebServers** configuration in the FortiGate UI

The configuration is very simple; a single trigger is added and multiple actions can be attached to the trigger.
