---
id: collect-261001-meraki/meraki/questions-1059522-meraki-vpn-behind-nat-breaks-bitbucket-git-on-wsl2-ubuntu-20-o-b7f49644
title: "questions-1059522-meraki-vpn-behind-nat-breaks-bitbucket-git-on-wsl2-ubuntu-20-o-b7f49644"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-meraki/questions-1059522-meraki-vpn-behind-nat-breaks-bitbucket-git-on-wsl2-ubuntu-20-o-b7f49644.md
source_anchor: ""
source_lines: [1, 35]
sha256: cdf3586a0c5e017fcca6a8730509f9a56739c2b093f3ba0f8f7ed90d0287cba2
---

# questions-1059522-meraki-vpn-behind-nat-breaks-bitbucket-git-on-wsl2-ubuntu-20-o-b7f49644

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I access two Meraki VPNs from Windows 10 Pro 10.0.19042 Build 19042:
One which is not behind a NAT - when I switch this on, I can do git clone [...] or git fetch [...] just fine.
Second, which is behind
a NAT - when I connect and run git fetch, results in the error message: "fatal:
unable to access 'https://bitbucket.org/[project]/[project-name].git/':
gnutls_handshake() failed: Error in the pull function."
To make the second VPN work, I executed following commands:
The problem was in a mismatch between VPN MTU and Linux under WSL2 MTU sizes.
It can be identified via 2 commands:
Windows PowerShell (run as administrator)
netsh interface ipv4 show subinterfaces
Notice the first row - it shows how big MTU is allowed in your VPN.
Linux (inside WSL2) console
ip addr
Notice the row starting 'eth0' - its MTU must match or be lower that the one above.
In my case the MTU in Linux was higher.
Solution
The following command instantly solves the problem:
sudo ip link set dev eth0 mtu 1400 (update MTU value to fit your VPN)
I have put it inside my ~/.bashrc and put /usr/sbin/ip into sudoers NOPASSWD for my account.
Better solution
So far I haven't managed to use any of the standard Linux tools to change MTU on Linux startup inside WSL2 (and hence to avoid putting it into .bashrc).
rc-local doesn't work under WSL2
/etc/dhcp/dhcpclient.conf doesn't
propagate changes into default interface-mtu nor supersede interface-mtu
netsh interface ipv4 set subinterface "vEthernet (WSL)" mtu=1400 store=persistent doesn't affect Linux
/etc/netplan
doesn't run inside WSL2
If you find the way, I'd be more than happy to have it here!
