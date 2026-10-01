---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/andeman-opnsense-cli-fc216d2a-1
title: "update linked item with names"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "license"]
source: docs/RAG/collect-261001-opnsense-pfsense/andeman-opnsense-cli-fc216d2a.md
source_anchor: ""
source_lines: [1, 157]
sha256: 63dafadda5cc9dcd09f107a780ae5761f5c0f629bcbebcbf882972fdf6ee4f15
---

# update linked item with names

opn-cli - the OPNsense CLI written in python.

```
pip install opn-cli
```
1. 
Generate an api_key and api_secret. See: https://docs.opnsense.org/development/how-tos/api.html#creating-keys.
2. 
Create the default config file `~/.opn-cli/conf.yaml````
---
api_key: your_api_key
api_secret: your_api_secret
url: https://opnsense.example.com/api
timeout: 60
ssl_verify: true
ca: ~/.opn-cli/ca.pem
```
3. 
Install required opnsense plugins ```
opn-cli plugin install os-firewall
opn-cli plugin install os-haproxy
opn-cli plugin install os-node_exporter
```

Each command and subcommand support the `-h` or `--help` option to show help for the current command.

The config basedir is `~/.opn-cli/`. If the environment variable `XDG_CONFIG_HOME` is set, `~/.config/opn-cli` will be used instead.

```
$ opn-cli --help
Usage: opn-cli [OPTIONS] COMMAND [ARGS]...
  OPNsense CLI - interact with OPNsense via the CLI
  API key + secret:
  You need a valid API key and secret to interact with the API. Open your
  browser and go to System->Access->Users and generate or use an existing
  Api Key.
  See: https://docs.opnsense.org/development/how-tos/api.html#creating-keys.
  SSL verify / CA:
  If you use ssl verification (--ssl-verify), make sure to specify a valid
  ca or cert with --ca <path_to_bundle>.
  To download the default self-signed cert, open the OPNsense Web Gui and go to
  System->Trust->Certificates. Search for the Name: "Web GUI SSL certificate" and
  press the "export user cert" button.
  If you use a ca signed certificate, go to System->Trust->Authorities and
  press the "export CA cert" button to download the ca.
  Save the ca and pass the path to the --ca option.
  Configuration:
  The base directory for the config is ~/.opn-cli.
  If the environment variable XDG_CONFIG_HOME is set, ~/.config/opn-cli will be used instead.
  You can set the required options as environment variables. See --help
  "[env var: [...]"
  Or use a config file passed with -c option.
  The configuration cascade from the highest precedence to lowest:
  1. argument & options
  2. environment variables
  3. config file
  Happy automating!
Options:
  -c, --config FILE               path to the config file  [env var:
                                  OPN_CONFIG; default: ~/.opn-cli/conf.yaml]
  --ca FILE                       path to the ca bundle file for ssl
                                  verification  [env var: OPN_SSL_VERIFY_CA;
                                  default: ~/.opn-cli/ca.pem]
  -k, --api-key TEXT              Your API key for the OPNsense API  [env var:
                                  OPN_API_KEY]
  -s, --api-secret TEXT           Your API secret for the OPNsense API  [env
                                  var: OPN_API_SECRET]
  -u, --url TEXT                  The Base URL for the OPNsense API  [env var:
                                  OPN_API_URL]
  -t, --timeout INTEGER           Set timeout for API Calls in seconds.  [env
                                  var: OPN_API_TIMEOUT; default: 60]
  --ssl-verify / --no-ssl-verify  Enable or disable SSL verification for API
                                  communication.  [env var: OPN_SSL_VERIFY;
                                  default: True]
  -h, --help                      Show this message and exit.
Commands:
  completion  Output instructions for shell completion
  firewall    Execute firewall operations
  haproxy     Manage haproxy loadbalancer operations
  ipsec       Manage Ipsec  
  new         Generate scaffolding code
  openvpn     Manage OpenVPN
  plugin      Manage OPNsense plugins
  route       Manage routes
  version     Show the CLI version and exit.
```
There is a docker image available with the opn-cli installed. You can use it to run the cli without installing it on your system.

Just run the following command:

```
docker run andeman77/opn-cli --help
```
The default config for the container is located at ./docker/.opn-cli.

To configure opn-cli bind mount the opn-cli config to the container. The config file should be named `conf.yaml` and the ca.pem
file should be named `ca.pem`. The config file should be located in the `~/.opn-cli` directory.

Example using the default config file (see: Configure):

```
docker run -v$(realpath $HOME)/.opn-cli:/home/appuser/.opn-cli andeman77/opn-cli firewall alias list
```
```
$ opn-cli completion
Instructions for shell completion:
See: https://click.palletsprojects.com/en/latest/shell-completion/
Bash (invoked every time a shell is started):
echo '# shell completion for opn-cli' >> ~/.bashrc
echo 'eval "$(_OPN_CLI_COMPLETE=bash_source opn-cli)"' >> ~/.bashrc
Bash (current shell):
_OPN_CLI_COMPLETE=bash_source opn-cli > ~/.opn-cli/opn-cli-complete.bash
source ~/.opn-cli/opn-cli-complete.bash
Zsh (invoked every time a shell is started):
echo '# shell completion for opn-cli' >> ~/.zshrc
echo 'eval "$(_OPN_CLI_COMPLETE=zsh_source opn-cli)"' >> ~/.zshrc
Zsh (current shell):
_OPN_CLI_COMPLETE=zsh_source opn-cli >! ~/.opn-cli/opn-cli-complete.zsh
source ~/.opn-cli/opn-cli-complete.zsh
```
Each command has a default output format. For lists and details the table output format and for create / update / delete the plain output formatis used.

You can always specify the output format with the `-o` output and show or hide columns with the `-c` output.

Show which default columns will be shown

```
$ opn-cli plugin installed -o cols
name,version,comment,locked
```
Show which columns are available. You could always pass an empty string to show all columns.

```
$ opn-cli plugin installed -o cols -c ''
name,version,comment,flatsize,locked,automatic,license,repository,origin,provided,installed,path,configured
```
Show output as pretty table.

```
$ opn-cli plugin installed -o table
+-----------------+---------+-----------------------------------+--------+
|       name      | version |              comment              | locked |
+-----------------+---------+-----------------------------------+--------+
|   os-firewall   |  1.0_2  | Firewall API supplemental package |  N/A   |
|     os-iperf    |  1.0_1  |      Connection speed tester      |  N/A   |
|  os-virtualbox  |  1.0_1  |     VirtualBox guest additions    |  N/A   |
| os-zabbix-agent |  1.8_2  |      Zabbix monitoring agent      |  N/A   |
+-----------------+---------+-----------------------------------+--------+
```
Always returns the complete json output. The `-c` output will be ignored.

```
$ opn-cli plugin installed -o json 
[{"name": "os-firewall", "version": "1.0_2", "comment": "Firewall API supplemental package", "flatsize": "56.0KiB", "locked": "N/A", "automatic": "N/A", "license": "BSD2CLAUSE", "repository": "OPNsense", "origin": "opnsense/os-firewall", "provided": "1", "installed": "1", "path": "OPNsense/opnsense/os-firewall", "configured": "1"}, {"name": "os-iperf", "version": "1.0_1", "comment": "Connection speed tester", "flatsize": "24.6KiB", "locked": "N/A", "automatic": "N/A", "license": "BSD2CLAUSE", "repository": "OPNsense", "origin": "opnsense/os-iperf", "provided": "1", "installed": "1", "path": "OPNsense/opnsense/os-iperf", "configured": "1"}, {"name": "os-virtualbox", "version": "1.0_1", "comment": "VirtualBox guest additions", "flatsize": "525B", "locked": "N/A", "automatic": "N/A", "license": "BSD2CLAUSE", "repository": "OPNsense", "origin": "opnsense/os-virtualbox", "provided": "1", "installed": "1", "path": "OPNsense/opnsense/os-virtualbox", "configured": "1"}, {"name": "os-zabbix-agent", "version": "1.8_2", "comment": "Zabbix monitoring agent", "flatsize": "49.2KiB", "locked": "N/A", "automatic": "N/A", "license": "BSD2CLAUSE", "repository": "OPNsense", "origin": "opnsense/os-zabbix-agent", "provided": "1", "installed": "1", "path": "OPNsense/opnsense/os-zabbix-agent", "configured": "1"}]
```
Filter the json output and return the columns specified with the `-c` output.

