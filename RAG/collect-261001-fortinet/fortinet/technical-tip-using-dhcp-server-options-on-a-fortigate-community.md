---
id: collect-261001-fortinet/fortinet/technical-tip-using-dhcp-server-options-on-a-fortigate-community
title: "technical-tip-using-dhcp-server-options-on-a-fortigate-community"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/technical-tip-using-dhcp-server-options-on-a-fortigate-community.md
source_anchor: ""
source_lines: [1, 21]
sha256: cb0e6368970ceb120ffdc3eabe52a9a366e28e1789301f480ab036f78e3e9cd1
---

# technical-tip-using-dhcp-server-options-on-a-fortigate-community

**Description**

It may be required to configure a FortiGate DHCP server that gives out a separate 'option' as well as IP information.  For example, in an environment that must support PXE boot with Windows images. **Solution**

After converting the required option code to HEX, the setup is done in CLI:

 config system dhcp server

edit <dhcpservername>

set option1 <option_code> [<option_hex>]

end

For example, to configure option 252 with value http://192.168.1.1/wpad.dat, the command line would be:

 ascii = ftpservers=192.168.21.30,country=1,language=1

hex = 667470736572766572733d3139322e3136382e32312e33302c636f756e7472793d312c6c616e67756167653d31

Additional information can be found in the FortiGate CLI Administration Guide.
