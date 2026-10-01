---
id: collect-261001-meraki/meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9-5
title: "fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9"
domain: meraki
role: reference
task: reference
actors: ["China", "United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9.md
source_anchor: ""
source_lines: [424, 491]
sha256: 023c5a71c8e56f776b7741f2debd147da2d40faab4c50c1400ccb79b5f1e1482
---

# fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9

| --extend-perfdata-group | Add new aggregated metrics (min, max, average or sum) for groups of metrics defined by a regex match on the metrics' names. Syntax: --extend-perfdata-group=regex,<names-of-new-metrics>,calculation[,[<new-unit-of-mesure>],[min],[max]] regex: regular expression <names-of-new-metrics>: how the new metrics' names are composed (can use $1, $2... for groups defined by () in regex). calculation: how the values of the new metrics should be calculated <new-unit-of-mesure> (optional): unit of measure for the new metrics min (optional): lowest value the metrics can reach max (optional): highest value the metrics can reach Common examples: um wrong packets from all interfaces (with interface need --units-errors=absolute): --extend-perfdata-group=',packets_wrong,sum(packets_(discard\|error)_(in\|out))' Sum traffic by interface: --extend-perfdata-group='traffic_in_(.*),traffic_$1,sum(traffic_(in\|out)_$1)' =back | 
| --change-short-output --change-long-output | Modify the short/long output that is returned by the plugin. Syntax: --change-short-output=pattern replacement modifier Most commonly used modifiers are i (case insensitive) and g (replace all occurrences). Example: adding --change-short-output='OKUp gi' will replace all occurrences of 'OK', 'ok', 'Ok' or 'oK' with 'Up' | 
| --change-short-output | Modify the short/long output that is returned by the plugin. Syntax: --change-short-output=pattern replacement modifier Most commonly used modifiers are i (case insensitive) and g (replace all occurrences). Example: adding --change-short-output='OKUp gi' will replace all occurrences of 'OK', 'ok', 'Ok' or 'oK' with 'Up' | 
| --change-long-output | Modify the short/long output that is returned by the plugin. Syntax: --change-short-output=pattern replacement modifier Most commonly used modifiers are i (case insensitive) and g (replace all occurrences). Example: adding --change-short-output='OKUp gi' will replace all occurrences of 'OK', 'ok', 'Ok' or 'oK' with 'Up' | 
| --change-exit | Replace an exit code with one of your choice. Example: adding --change-exit=unknown=critical will result in a CRITICAL state instead of an UNKNOWN state. | 
| --change-output-adv | Replace short output and exit code based on a "if" condition using the following variables: short_output, exit_code. Variables must be written either %{variable} or %(variable). Example: adding --change-output-adv='%(short_ouput) =~ /UNKNOWN: No daemon/,OK: No daemon,OK' will change the following specific UNKNOWN result to an OK result. | 
| --range-perfdata | Rewrite the ranges displayed in the perfdata. Accepted values: 0: nothing is changed. 1: if the lower value of the range is equal to 0, it is removed. 2: remove the thresholds from the perfdata. | 
| --filter-uom | Mask the units when they don't match the given regular expression. | 
| --opt-exit | Replace the exit code in case of an execution error (i.e. wrong option provided, SSH connection refused, timeout, etc). Default: unknown. | 
| --output-ignore-perfdata | Remove all the metrics from the service. The service will still have a status and an output. | 
| --output-ignore-label | Remove the status label ("OK:", "WARNING:", "UNKNOWN:", CRITICAL:") from the beginning of the output. Example: 'OK: Ram Total:...' will become 'Ram Total:...' | 
| --output-xml | Return the output in XML format (to send to an XML API). | 
| --output-json | Return the output in JSON format (to send to a JSON API). | 
| --output-openmetrics | Return the output in OpenMetrics format (to send to a tool expecting this format). | 
| --output-file | Write output in file (can be combined with JSON, XML and OpenMetrics options). Example: --output-file=/tmp/output.txt will write the output in /tmp/output.txt. | 
| --disco-format | Applies only to modes beginning with 'list-'. Returns the list of available macros to configure a service discovery rule (formatted in XML). | 
| --disco-show | Applies only to modes beginning with 'list-'. Returns the list of discovered objects (formatted in XML) for service discovery. | 
| --float-precision | Define the float precision for thresholds (default: 8). | 
| --source-encoding | Define the character encoding of the response sent by the monitored resource Default: 'UTF-8'. <output>. | 
| --http-peer-addr | Set the address you want to connect to. Useful if hostname is only a vhost, to avoid IP resolution. | 
| --proxyurl | Proxy URL. Example: http://my.proxy:3128 | 
| --proxypac | Proxy PAC file (can be a URL or a local file). | 
| --insecure | Accept insecure SSL connections. | 
| --http-backend | Perl library to use for HTTP transactions. Possible values are: lwp (default) and curl. | 
| --memcached | Memcached server to use (only one server). | 
| --redis-server | Redis server to use (only one server). Syntax: address[:port] | 
| --redis-attribute | Set Redis Options (--redis-attribute="cnx_timeout=5"). | 
| --redis-db | Set Redis database index. | 
| --failback-file | Fall back on a local file if Redis connection fails. | 
| --memexpiration | Time to keep data in seconds (default: 86400). | 
| --statefile-dir | Define the cache directory (default: '/var/lib/centreon/centplugins'). | 
| --statefile-suffix | Define a suffix to customize the statefile name (default: ''). | 
| --statefile-concat-cwd | If used with the '--statefile-dir' option, the latter's value will be used as a sub-directory of the current working directory. Useful on Windows when the plugin is compiled, as the file system and permissions are different from Linux. | 
| --statefile-format | Define the format used to store the cache. Available formats: 'dumper', 'storable', 'json' (default). | 
| --statefile-key | Define the key to encrypt/decrypt the cache. | 
| --statefile-cipher | Define the cipher algorithm to encrypt the cache (default: 'AES'). | 
| --hostname | Meraki API hostname (default: 'api.meraki.com') The default value 'api.meraki.com' will work for most of the world. However, for organizations hosted in the following country dashboards, you need to override this value and specify the respective base URI instead: Canada: https://api.meraki.ca/api/v1 China: https://api.meraki.cn/api/v1 India: https://api.meraki.in/api/v1 United States FedRAMP: https://api.gov-meraki.com/api/v1 Please refer to Meraki API documentation https://developer.cisco.com/meraki/api-v1/getting-started/#base-uri for more details. | 
| --port | Define the TCP port to use to reach the API (default: 443). | 
| --proto | Define the protocol to reach the API (default: 'https'). | 
| --api-token | Meraki API token. | 
| --timespan | Define the duration, in seconds, of the historical data to fetch. Can be 300, 600, 1200, 3600, 14400, 86400 (default: 300). | 
| --timeout | Define the timeout for HTTP requests. | 
| --ignore-permission-errors | Ignore permission errors (403 status code). | 
| --ignore-orgs-api-disabled | Ignore organizations where the API is disabled. | 
| --api-filter-orgs | Define the organizations to monitor (regular expression). | 
| --cache-use | Use the cache file instead of requesting the API (the cache file can be created with the cache mode). | 
Options des modes
Les options disponibles pour chaque modèle de services sont listées ci-dessous :
- Api-Requests
- Device
- Devices
- Network
- Networks
- Vpn-Tunnels
| Option | Description | 
|---|---|
| --filter-organization-name | Filter organization name (can be a regexp). | 
| --warning-api-requests-200 | Threshold. | 
| --critical-api-requests-200 | Threshold. | 
| --warning-api-requests-404 | Threshold. | 
| --critical-api-requests-404 | Threshold. | 
| --warning-api-requests-429 | Threshold. | 
| --critical-api-requests-429 | Threshold. | 
| Option | Description | 
|---|---|
| --filter-device-name | Filter devices by name (can be a regexp). | 
| --filter-link-name | Filter VPN links by name (can be a regexp). | 
| --filter-network-id | Filter devices by network ID (can be a regexp). | 
