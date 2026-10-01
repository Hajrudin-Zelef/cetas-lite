---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-administration-guide-639044-cli-script-action-b990f18d
title: "CLI script action"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-administration-guide-639044-cli-script-action-b990f18d.md
source_anchor: ""
source_lines: [1, 127]
sha256: e03d2649e09c0971a91500830f67d99c5e7f351c3282b423ac8fe092294f3a8a
---

# CLI script action

The CLI script action can run when an automation stitch is triggered. It executes a series of commands in the CLI, as defined by the administrator. The scripts commands can be entered manually, uploaded as a file, or recorded in the CLI console. The output of the script can be fed as a variable (`%%results%%`) into the next action in the stitch. This could then be sent as an email using the email action, for example.

|  | The CLI script action utilizes the *auto-script* feature to perform the execution of the script commands. The output size of the auto-script feature controls the size of the output for the script execution (10MB by default). This output is read into a buffer for use by the automation stitch action, and that buffer is limited to 192K characters for the email action. This means that the total allowable limit for CLI script output that is used in an email action is 192K characters. The buffer sizes for other actions may vary, for example the buffer to use the results in a webhook action could be a different size. | 

```
config system automation-action
    edit <name>
        set action-type cli-script
        set output-size <integer>
        set timeout <integer>
    next
end
```
                                            | output-size <integer> | Set the size to limit the script output, in megabytes (1 - 1024, default = 10). | 
| timeout <integer> | Set the maximum running time for this script, in seconds (0 - 300, 0 = no timeout). | 

|  | Certain diagnostic commands may not function as expected with CLI scripts and result in no output. For example, when used in a CLI script, the diagnostic command `dia test application dnsproxy 6` fails to produce any output because the cli-script feature does not support`daemon message()` prints. | 

In this example, the script sets the idle timeout value to 479 minutes, and sends an email with the script output.

###### To configure a stitch with a CLI script action in the GUI:

1. 
                                                    Go to *Security Fabric > Automation* and click*Create New* .
2. 
                                                    Enter the stitch name ( *auto-cli-1* ).
3. 
                                                    Configure the trigger: 
  1. 
                                                            Click *Add Trigger* .
  2. 
                                                            Select *Any Security Rating Notification* .
  3. 
                                                            Click *Apply* .
4. 
                                                            
5. 
                                                    Configure the CLI Script action: 
  1. 
                                                            Click *Add Action* .
  2. 
                                                            Click *Create* and select*CLI Script* .
  3. 
                                                            Enter the following: Name admintimeout Script ```
config system global
    set admintimeout 479
end
```
Alternatively, click *Upload* to upload a file, or click*>_Record in CLI console* and enter the CLI commands.Administrator profile Select a profile
  4. 
                                                            Click *OK* .
  5. 
                                                            Select the action in the list and click *Apply* .
6. 
                                                            
7. 
                                                    Configure the Email notification action: 
  1. 
                                                            Click *Add Action* .
  2. 
                                                            Click *Create* and select*Email* .
  3. 
                                                            Enter the following: Name auto-cli-1_email To Enter an email address Subject CSF stitch alert Body %%results%%
  4. 
                                                            Click *OK* .
  5. 
                                                            Select the action in the list and click *Apply* .
8. 
                                                            
9. 
                                                    Click *OK* .

###### To configure a stitch with a CLI script action in the CLI:

1. 
                                                    Create the automation trigger: ```
config system automation-trigger
    edit "Any Security Rating Notification"
        set event-type security-rating-summary
        set report-type any			
    next
end
```
2. 
                                                    Create the automation actions: ```
config system automation-action
    edit "admintimeout"
        set action-type cli-script
        set script "config system global
            set admintimeout 479
            end"
        set output-size 10
        set timeout 0	
        set accprofile "super_admin"
    next
    edit "auto-cli-1_email"
        set action-type email
        set email-to "admin@fortinet.com"
        set email-subject "CSF stitch alert"
        set message "%%results%%"
    next
end
```
3. 
                                                    Create the automation stitch: ```
config system automation-stitch
    edit "auto-cli-1"
        set trigger "Any Security Rating Notification"
        config actions
            edit 1
                set action "admintimeout"
                set required enable
            next
            edit 2
                set action "auto-cli-1_email"
                set required enable
            next
        end
    next
end
```

###### Sample email

The email sent by the action will look similar to the following:
