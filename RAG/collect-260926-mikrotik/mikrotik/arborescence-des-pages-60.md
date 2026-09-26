---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-60
title: "SSH Server"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-60.md
source_anchor: ""
source_lines: [1, 102]
sha256: 33331f1bc5432c356af43f2c214a621e4e9db484ae484434770045b61e27d71b
---

# SSH Server

RouterOS has built in SSH (SSH v2) server that is enabled by default and is listening for incoming connections on port TCP/22. It is possible to change the port and disable the server under Services menu.

**Sub-menu:** `/ip ssh`

| Property | Description | 
|---|---|
| **password-authentication** (*yes-if-no-key* \|*yes \| no* ; Default:**yes-if-no-key** ) | Whether to allow password login at the same time when public key authorization is configured for a user. | 
| **ciphers** (*3des-cbc *\| aes-cbc \| aes-ctr \| aes-gcm \| auto \| null*;* Default:**auto** ) | Allow to configure SSH ciphers. | 
| **forwarding-enabled** (*both \| local \| no \| remote* ; Default:**no** ) | Allows to control which SSH forwarding method to allow:  | 
| **host-key-size** (*1024 \| 1536 \| 2048 \| 4096 \| 8192* ; Default:**2048** ) | RSA key size when host key is being regenerated. | 
| **host-key-type**  (*ed25519* \|*rsa* ; Default:**rsa** ) | Select host key type | 
| **publickey-authentication-options** (*none* \|*touch-required* \|*verify-required* ; Default:**none** ) | Sets public key authentication options. The touch-required option causes public key authentication using a FIDO authenticator algorithm to always require the signature to attest that a physically present user explicitly confirmed the authentication (usually by touching the authenticator). The verify-required option requires a FIDO key signature attest that the user was verified, e.g. via a PIN. | 
| **strong-crypto** (*yes \| no* ; Default:**no** ) | Use stronger encryption, HMAC algorithms, use bigger DH primes and disallow weaker ones:  | 

**Commands**

| Property | Description | 
|---|---|
| **export-host-key** (*key-file-prefix* ) | Export public and private RSA/Ed25519 to files. Command takes two parameters:  Host keys are exported in PKCS#8 format. | 
| **import-host-key** (*private-key-file* ) | Import and replace private RSA/Ed25519 key from specified file. Command takes two parameters:  Private key is supported in PEM or PKCS#8 format. | 
| **regenerate-host-key** () | Generated new and replace current set of private keys (RSA/Ed25519) on the router. Be aware that previously imported keys might stop working. | 

Exporting the SSH host key requires "sensitive" user policy.

## Enabling PKI authentication

Example of importing public key for user *admin*

Get SSH key pair on the client device (the device you will connect from). Upload the public SSH key to the router and import it.

More information about supported SSH keys find here.

## SSH key pair generation

RouterOS does not support direct SSH key generation, which is available on Linux systems.

To obtain an SSH key pair (SSH key pair is automatically generated on the first SSH connection or when host key is exported), the device's SSH host key must be exported.

# SSH Client

**Sub-menu:** `/system ssh`

## **Simple log-in to remote host**

It is able to connect to remote host and initiate ssh session. IP address supports both IPv4 and IPv6.

In this case user name provided to remote host is one that has logged into the router. If other value is required, then *user=<username>* has to be used.

## **Log-in from certain IP address of the router**

For testing or security reasons it may be required to log in to other host using certain source address of the connection. In this case *src-address=<ip address>* argument has to be used. Note that IP address in this case supports both, IPv4 and IPv6.

in this case, ssh client will try to bind to address specified and then initiate ssh connection to remote host.

Example of importing RSA private key for user *admin.*

First, export currently generated SSH keys to a file:

Two files *admin_rsa* and *admin_rsa.pub* will be generated. The pub file needs to be trusted on the SSH server side (how to enable SSH PKI on RouterOS) The private key has to be added for the particular user.

Only user with full rights on the router can change 'user' attribute value under */user ssh-keys private*

After the public key is installed and trusted on the SSH server, a PKI SSH session can be created.

Watch how to:

Log in with an RSA key.

Log in with Ed25519.

## **Executing remote commands**

To execute remote command it has to be supplied at the end of log-in line

*If the server does not support pseudo-tty (ssh -T or ssh host command), like MikroTik ssh server, then it is not possible to send multiline commands via SSH*

For example, sending command `"/ip address \n add address=1.1.1.1/24"` to MikroTik router will fail.

If you wish to execute remote commands via **scripts** or **scheduler**, use command **ssh-exec**.

# SSH exec

**Sub-menu:** `/system ssh-exec`

Command *ssh-exec* is a non-interactive ssh command, thus allowing to execute commands remotely on a device via scripts and scheduler.

## **Retrieve information**

The command will return two values:

- **exit-code** : returns 0 if the command execution succeeded
- **output** : returns the output of remotely executed command

**Example:** Code below will retrieve interface status of ether1 from device 10.10.10.1 and output the result to "Log"

For security reasons you should not use plain text password with parameter "password" specified in the command line. To ensure safe execution of the command remotely, it is strongly recommended to use SSH PKI authentication for users on both sides.

The user group and script policy executing the command requires **test** permission

Watch how to execute commands through SSH.
