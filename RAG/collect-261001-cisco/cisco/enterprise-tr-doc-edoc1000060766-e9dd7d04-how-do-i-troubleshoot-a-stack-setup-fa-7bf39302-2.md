---
id: collect-261001-cisco/cisco/enterprise-tr-doc-edoc1000060766-e9dd7d04-how-do-i-troubleshoot-a-stack-setup-fa-7bf39302-2
title: "enterprise-tr-doc-edoc1000060766-e9dd7d04-how-do-i-troubleshoot-a-stack-setup-fa-7bf39302"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-tr-doc-edoc1000060766-e9dd7d04-how-do-i-troubleshoot-a-stack-setup-fa-7bf39302.md
source_anchor: ""
source_lines: [82, 107]
sha256: 58f126d71879944c98efbea2a39e01816da5fa93b62f64fe271387cd6ed6184f
---

# enterprise-tr-doc-edoc1000060766-e9dd7d04-how-do-i-troubleshoot-a-stack-setup-fa-7bf39302

| The type of devices is different. | The device types of member switches are different. | Set up a stack using switches of the same device type. | 
| Stack-Port link invalid. | Logical stack ports are incorrectly connected. |  | 
| Switches working in different forward modes cannot set up a CSS. | The card interoperability modes of member switches are different. |  | 
| The interface Stack-Port is down. | The protocol status of a logical stack port is Down. |  | 
| The physical status of the stack member port is up, but the protocol status is down. | The protocol status of a physical stack member port is Down, but its physical status is Up. |  | 
| Configuration conflict. | The configuration of a member switch conflicts with that of the master switch. | In most cases, this fault will occur if the master switch has member switch's stack port configurations that do not take effect. Run the display current-configuration all command on the master switch to check whether there are member switch's stack port configurations that do not take effect. If so, delete the configurations from the master switch. | 
| The port on CE-FWA board did not support configured as stack port. | Ports on the card cannot be used for stacking. | Use ports on stack-supporting cards to set up a stack. | 
Run the display diagnostic-information file-name command in the user view to collect diagnostic information and save the information to a file.
<HUAWEI> display diagnostic-information dia-info.txt 
Now saving the diagnostic information to the device 
 100% 
Info: The diagnostic information was saved to the device successfully.
The text file is stored in the flash:/ directory by default. You can run the dir command in the user view to check whether the file is generated.
After diagnostic information is saved to a file, you can export it from the device through FTP, SFTP, or SCP. For details, see Local File Management.
You can run the display diagnostic-information command to display diagnostic information and save terminal logs in a diagnostic file on a disk. For details, see Diagnostic File Obtaining Guide.
Run the save logfile command to save the log and trap information in the log buffer to files.
<HUAWEI> save logfile //Collect common user logs.
<HUAWEI> system-view 
[~HUAWEI] diagnose 
[~HUAWEI-diagnose] save logfile diagnose-log             //Collect diagnostic logs.
[~HUAWEI-diagnose] collect diagnostic information       //Collect diagnostic information of the operating system.
When the log files are generated, you can export the files from the device using FTP, SFTP, or SCP. For details, see Local File Management.
You can also run the display logbuffer and display trapbuffer commands to view the log and trap information on the device, and save terminal logs in a diagnostic file on a disk. For details, see Diagnostic File Obtaining Guide.
Contact us for technical support.
Technical support personnel will provide instructions for you to submit all the collected information and files, so that they can locate faults.
Select the content with the mouse pointer to quickly report the problem.
