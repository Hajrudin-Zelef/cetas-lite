---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-c1esvr-how-to-master-fortigate-cli-commands-c7627e4e
title: "r-fortinet-comments-c1esvr-how-to-master-fortigate-cli-commands-c7627e4e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-c1esvr-how-to-master-fortigate-cli-commands-c7627e4e.md
source_anchor: ""
source_lines: [1, 53]
sha256: 452179743c28c1cfc5f121cd3d8dd4bfbbf5a276377d05e206669b108380841d
---

# r-fortinet-comments-c1esvr-how-to-master-fortigate-cli-commands-c7627e4e

How to master Fortigate CLI commands 
        
    Hi!
I did search google but cannot find some good article to learn FortiGate Cli commands. I want to learn more in depth if someone knows some blog or some site which I cannot find. I did read somewhere that FortiGate show and get commands is different in a way that if configuration is default then you use either one of them and if configuration is changed that use either of them but cannot find that link.
Thanks
Section des commentaires
Thanks All. I know much of stuff like debug and edit ports etc. But they use a different command struture. I will go through Cli.fortinet.com and see if I can find good help.
Found this:
Looks helpful
https://www.reddit.com/r/fortinet/comments/9l6vmm/fortigate_cli_reference_sheet/
It is pretty easy to use, IMO the best CLI out there. With that being said just start using it. A few tips:
If you hit tab it will cycle between choices. For example if you did config VPN ipsec phase <tab> it would cycle between phase1, phase1-interface, phase2, phase2-interface.
Diagnose debug flow is a great tool to analyze traffic flow.
If you are adding an entry into a list "edit 0" will always put the new entry at the next sequence.
You have to "end" the configuration directory you are in for changes to take effect.
The CLI is more powerful and you have more options you can config than the GUI. For example you can set the source interface for your syslog, which you cant do in the GUI.
If you are lost in the directory looking for a command just follow the GUI tree, the CLI tree is extremely similar. For example the GUI tree to setup a ipv4 policy is Firewall > ipv4 policy, the CLI equivalent is config firewall policy.
Man you need to play with Junos then. That is one hell off a cli
Yah I think FortiGate is a superior product especially for the money, but hands down the best CLI on the market just has to be JunOS. I don’t even see how that’s a preference or opinion kind of thing. I mean I get being mainly exposed to one CLI or another and because of that having your personal preference, but nothing I’ve ever seen matches the shear logical layout of the config and built in power of JunOS.
I have, It is not bad, better than Palos that's for sure.
Best ever!
Agreed totally. I am coming from Junos/SRX.
show full | grep -f keyword
Will show you the section of the object around the line containing the keyword.
Just use it!
What I typically recommend is to watch the CLI commands that are being used when you are using the FortiGate WebGUI. You can do this by doing the following:
Open up putty (SSH)
Use the following commands:
diagnose debug cli 8diag debug enable
Now do things like create a firewall policy, create a route, objects, etc. You can learn quickly how to use the CLI to do common tasks you do in the WebGUI. I hope this helps.
Great tip, thank you
Already some good advice, but a couple things I'd add:
In addition to TAB to complete commands, you can use the ? mark to see available commands.
Using after an edit (such as in "config sys int" or "config firewall rule") will list the names of existing interfaces, rules, objects, etc.
It's a great way to see what is configured or possible to configure.
2) When editing a specific object (interface, vpn tunnel, rule) you can use the get command.
Show only displays the configured object, which is typically what you'll use.
Get displays all settings though, including default values.
It's another great command to see what is configurable and find default values you might not realize.
Bom dia. basta colocar um "?" no console ja aparece as variações dos principais comandos.
Show is for non-Default Config. Get or Show full for everything
Fortinet has a CLI guide on the docs site for various releases.
It’s probably helpful to download and read through a few Config files as well.
Cli.fortinet.com and navigate to the cli reference.
From the cli, tree will show the config tree.
And show full-configuration.
Show will reflect configured options but not necessarily all default settings.
Get in a config stanza will show all configured values including those with default settings.
Learn by scenarios. Use the cookbooks. Use the online manuals.
Setting up an IPSec tunnel between two devices will cover much of what you need to learn, including debug commands.
It may well also take you to levels of frustration that you didn’t think were humanly possible.
Which cookbook do you recommend? Thanks
If you’re in a bind, you can also use the tree command and search/skim through the output. Not the most effective way but in some cases it can be helpful.
