---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/ansible-opnsense-docs-source-usage-2-basic-rst-at-5cec92ea25c8280a9913dfb958109f78ce930d00
title: "show difference"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/ansible-opnsense-docs-source-usage-2-basic-rst-at-5cec92ea25c8280a9913dfb958109f78ce930d00-o-x-l-ans.md
source_anchor: ""
source_lines: [1, 114]
sha256: f8dc3149c57416db05fb6bf683ce77d68701edee924d347eed149122b86f8bbe
---

# show difference

If you DO NOT want to use Ansible - this fork provides you with a raw Python3 interface.

You need to create API credentials as described in the OPNsense documentation.

**Menu**: System - Access - Users - Edit {admin user} - Add api key

Make sure to set `connection: local` in your OPNsense-Playbook, so the Module gets executed on your Ansible-Controller machine!

See also: How Ansible Works

If you are running the modules over hosts in your inventory - you would do it like that:

```
- hosts: firewalls
  connection: local  # execute modules on controller
  gather_facts: false
  tasks:
    - name: Example
      oxlorg.opnsense.alias:
        firewall: "{{ ansible_host }}"  # or use a per-host variable to store the FQDN..
```
If you use your firewall for non-testing purposes - you should **ALWAYS USE SSL VERIFICATION** for your connections!

`ssl_verify: true`
To make a connection trusted you need either:

- a valid public certificate for the DNS-Name your firewall has (*LetsEncrypt/ACME* )
- an internal certificate authority that is used to create signed certificates
  - you could create such internal certificates using OPNsense. See the OPNsense documentation for self-signed certificates.
  - if you do so - it is important that the IP-address and/or DNS-Name of your firewall is included in the 'Subject Alternative Name' (*SAN* ) for it to be valid

After you got a valid certificate - you need to import and activate it:

- Import: 'System - Trust - Certificates - Import'
- Make sure your DNS-Names are allowed: 'System - Settings - Administration - Alternate Hostnames'
- Activate: 'System - Settings - Administration - SSL Certificate'

If you are using an internal CA for your certificates - you have to provide its public key to the modules:

`ssl_ca_file: '/path/to/ca.pem'`
If some parameters will be the same every time - use 'module_defaults':

```
- hosts: firewalls
  connection: local
  gather_facts: false
  module_defaults:
    oxlorg.opnsense.alias:
        firewall: 'opnsense.template.opnsense.oxl.app'
        api_credential_file: '/home/guy/.secret/opn.key'
        # if you use an internal certificate:
        #   ssl_ca_file: '/etc/ssl/certs/custom/ca.crt'
        # else you COULD (but SHOULD NOT) use:
        #   ssl_verify: false
  tasks:
    - name: Example
      oxlorg.opnsense.alias:
        name: 'ANSIBLE_TEST1'
        content: ['1.1.1.1']
```
You may want to use '**ansible-vault**' to **encrypt** your 'api_secret'.

Vault-Encryption of the 'api_credential_file' is not yet supported.

`ansible-vault encrypt_string 'YOUR_API_SECRET'`
Then add it as variable to your inventory/config:

```
firewall:
  key: 'test'
  secret: !vault |
    $ANSIBLE_VAULT;1.1;AES256
    38303736393135366562396233353930366631396531613062366365363234363063626365656263
    6637646636323437333437353336316332663133316435650a366439336665383763376432653736
    32313332363032646436626230646461376532666366663265373663316331316664336134366338
    6531363362613039330a316436386533393636623837653163333564383232313363666361643730
    3132
```
And refer to it in the module calls or module-defaults:

```
- hosts: firewalls
  connection: local
  gather_facts: false
  module_defaults:
    oxlorg.opnsense.route:
      firewall: '...'
      api_key: "{{ firewall.key }}"
      api_secret: "{{ firewall.secret }}"
```
To decrypt those secrets at runtime, you need to supply the 'ask-vault-pass' argument:

`ansible-playbook -D opnsense.yml --ask-vault-pass`
These modules support check-mode and can show you the difference between existing and configured items:

```
# show difference
ansible-playbook opnsense.yml -D
# run in check-mode (no changes are made)
ansible-playbook opnsense.yml --check
```
The module's HTTP-Traffic can be forwarded over a forward-proxy like Squid by specifying the `HTTPS_PROXY` environmental variable!

```
- name: Example Playbook
  hosts: firewalls
  connection: local
  gather_facts: false
  environment:
    HTTPS_PROXY: 'http://user:password@squid.template.opnsense.oxl.app:3128'
    # to inherit it from the ansible-controller environment:
    # HTTPS_PROXY: "{{ lookup('ansible.builtin.env', 'HTTPS_PROXY') | default('') }}"
  ...
```
