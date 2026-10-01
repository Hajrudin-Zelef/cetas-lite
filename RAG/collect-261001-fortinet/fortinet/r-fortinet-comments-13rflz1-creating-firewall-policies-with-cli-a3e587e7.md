---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-13rflz1-creating-firewall-policies-with-cli-a3e587e7
title: "r-fortinet-comments-13rflz1-creating-firewall-policies-with-cli-a3e587e7"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-13rflz1-creating-firewall-policies-with-cli-a3e587e7.md
source_anchor: ""
source_lines: [1, 28]
sha256: da82f74d7ea23a5cd303e718cddc1386243f214306e58bf6a169d3e6d85ca818
---

# r-fortinet-comments-13rflz1-creating-firewall-policies-with-cli-a3e587e7

Creating firewall policies with CLI 
        
    The only issue I have is the policy positioning. It always adds to the bottom of the sequence, I need it to be in the top 3.
Is there any way to set the position while creating the policy?
Using the move command after creating is not an option as I am creating the policy and on 14 appliances and don't want to login to check the created ID and the ID of policy 2 on every firewall.
Is there something similar to set sequence top or set position top something that can help me place the policy in the correct sequence order when I create it?
Section des commentaires
I have settled for adding the CLI to a script in each FMG ADOM, running it to create the policy, then manually going to each template and moving the policy to the correct location in the sequence and installing.
Still significantly quicker than logging into each device but certainly not min/maxed.
before creating the policy in FortiManager right click on the top policy and select Insert Above
A dummy policy will be created which is disabled. Configure the policy as you would normally in FMG, right click and select enable on the policy and do a policy push.
Then I need to manually create the same policy 12 times in the GUI.
Without move, no. There is no way to create a policy in the middle of the list. New policy, goes to the end of the line.
FortiManager ?
It aint CLI , I know but you would know what you did and where you did it, the way you want it.
I am using FMG, Devices are split across different ADOMS and the policy ID they get isn't the same on each device as they have different policies already.
Maybe its time to switch to Global if the rules are the same across ADOMs.
No, there is nothing that will do/facilitate what you requested.
I get around that by generally using specific IDs for devices I manage, so I can move them around via CLI, but that's not always going to be possible. (It is so far across over a dozen devices and a handful of customers, but that's not always going to be true)
You can use tcl on Fortimanager. That way you should be able to fetch the policies and issue the appropriate move command.
Disclaimer: I haven't tried it...
Theoretically it could work via API as well. Haven't tested it either.
http://docs.fortinet.com/document/fortigate/7.2.4/administration-guide/627485/subcommands
movemove <id> before <id>move <id> after <id>
With out using move after the fact. I don't want to login to 12 gates to check what the new policy ID is, then I may as well just log in and create the policy in the GUI.
Why not set the ID when you create the policy? That's what I do. When I create a new policy using "edit <new #>" then I can do exactly what you want predictably. Although I guess you might have the same problem with knowing what the policy number is you want to move it before/after. I have all mine documented and consistent in a spreadsheet -- probably better ways but it works for me and solves the problem you're trying to solve.
You didn't read OP's post. Move is pointless, because you don't know where to move it to or what ID the created policy gets.
Fair point, brain misfiring.
