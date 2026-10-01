---
id: collect-261001-fortinet/fortinet/questions-1073606-azure-run-command-for-fortinet-virtual-machine-with-serial-con-856a2a79
title: "questions-1073606-azure-run-command-for-fortinet-virtual-machine-with-serial-con-856a2a79"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-1073606-azure-run-command-for-fortinet-virtual-machine-with-serial-con-856a2a79.md
source_anchor: ""
source_lines: [1, 22]
sha256: c7b5834828aa1b50f00f227f18ea99073de1ca121ec3493adc4cabc46d7edf8c
---

# questions-1073606-azure-run-command-for-fortinet-virtual-machine-with-serial-con-856a2a79

I'm trying to execute with Azure a .conf file that has some command lines that are executable into the Fortinet's serial console in a virtual machine with FortiOS v6.4.3 (Linux based system).
I use AZ CLI to execute the script into the virtual machine. the command is the following:
az vm run-command invoke --subscription "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" --resource-group "myresourcegroup" --name "myfortinet" --command-id RunShellScript --scripts $scriptfgtvm
Where $scriptfgtvm contains this:
config system interface
edit port1
set allowaccess ping https ssh fgfm snmp
next
end
The problem is that we don't see the modifications desired, and thE feedback is "Not supported" but successfully Provisioned.
{
  "value": [
    {
      "code": "ProvisioningState/succeeded",
      "displayStatus": "Provisioning succeeded",
      "level": "Info",
      "message": "Not supported.",
      "time": "2021-08-05T08:24:01+00:00"
    }
  ]
}
When we connect to the vm remotely through SSH we get into the serial console, not into the shell console, the same with Azure Bastion
