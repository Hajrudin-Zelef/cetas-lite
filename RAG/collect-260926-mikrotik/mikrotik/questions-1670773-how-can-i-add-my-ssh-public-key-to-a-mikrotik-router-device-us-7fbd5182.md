---
id: collect-260926-mikrotik/mikrotik/questions-1670773-how-can-i-add-my-ssh-public-key-to-a-mikrotik-router-device-us-7fbd5182
title: "questions-1670773-how-can-i-add-my-ssh-public-key-to-a-mikrotik-router-device-us-7fbd5182"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1670773-how-can-i-add-my-ssh-public-key-to-a-mikrotik-router-device-us-7fbd5182.md
source_anchor: ""
source_lines: [1, 29]
sha256: 1e08d82b1666c837fd70557ea08d0bcc7ccf325782215897ce5b6bb30cb80f89
---

# questions-1670773-how-can-i-add-my-ssh-public-key-to-a-mikrotik-router-device-us-7fbd5182

Mikrotik doesn't support ssh-copy-id which copies ~/.ssh/id_rsa.pub into the .ssh/authorized_keys file of a host on almost all Linux devices.
Is there an equivalent way to authorize a key on a Mikrotik RouterOS device?
Mikrotik only allows you to import a key from a file that you copied over - but you can create this file from the command line. There's a one-liner that should work from any Linux host.
 ssh 192.168.88.1 "/file print file=mykey; file set mykey contents=\"`cat ~/.ssh/id_rsa.pub`\";/user ssh-keys import public-key-file=mykey.txt;/ip ssh set always-allow-password-login=yes"
That is all you need to do.
The line-by-line version of the above:
 $ ssh user@router
 [user@router] > /file print file=mykey;
 [user@router] > /file set mykey contents="copy and paste contents of ~/.ssh/id_rsa.pub here";
 [user@router] > /user ssh-keys import public-key-file=mykey.txt
 [user@router] > /ip ssh set always-allow-password-login=yes
This will allow you to simply ssh without being prompted for a password by the router. Note that if your private key is password protected, you will be prompted for your private key file password, but that password will not be sent to the router. You can use ssh-agent and ssh-add to cache your private key on the host, if it is password protected - which is highly recommended - otherwise anyone with a copy will be able to access all your routers that allow your key. Note that RouterOS automatically adds a .txt extension.
Another usability tip for ssh in general - if you add the following to /etc/ssh/ssh_config then connections will be kept open - and reused - which greatly speeds up logging in on devices where you log in a lot:
 ControlMaster auto
 ControlPath ~/.ssh/socket-%r@%h-%p
 ControlPersist 600
The only snag is that if the connection times out, connecting will hang if you reconnect within that timeout. You can of course manually delete the socket, which you can find with find .ssh/sock*
There's another way - although it requires an additional step, so it's not a "one-liner" solution - and that is to manually copy the key file over, before running the import command.
scp ~/.ssh/id_rsa.pub 192.168.88.1:/
ssh 192.168.88.1 "/user ssh-keys import public-key-file=id_rsa.pub; /ip ssh set always-allow-password-login=yes"
Note that the :/ is critically important, without it the file doesn't get saved - and like with a lot of Mikrotik things, there is no error condition created - also that this method the .txt extension does not get generated. There's also a potential race condition here with the file name - and more issues, so if you have a large team of engineers - or more than one key master... the following are probably better solutions:
scp ~/.ssh/id_rsa.pub 192.168.88.1:/`whoami`.pub
ssh 192.168.88.1 "/user ssh-keys import public-key-file=`whoami`.pub; /ip ssh set always-allow-password-login=yes"
Or
 (h=192.168.88.1; u=`whoami`; k="`cat ~/.ssh/id_rsa.pub`"; ssh $u@$h "/file print file=$u; file set $u contents=\"$k\";/user ssh-keys import public-key-file=$u.txt;/ip ssh set always-allow-password-login=yes;:set \$f [/file find name=$u.txt]; :if (\$f)  do={:put \"Duplicate key\"; [/file remove \$f]} else={:put \"Key added\"}")
The import command removes the file - unless there is already a key, then the file remains and needs to be manually removed. This last command does that too. If your router username is different from your local username, then change the u='whoami' to u=admin or whatever it may be. If you specify a different login username in your ssh_config file, then just remove the $u@ part after the ssh command.
/file).
                
                -o PubkeyAcceptedAlgorithms=+ssh-rsa on your ssh usr@host -i identify-file
