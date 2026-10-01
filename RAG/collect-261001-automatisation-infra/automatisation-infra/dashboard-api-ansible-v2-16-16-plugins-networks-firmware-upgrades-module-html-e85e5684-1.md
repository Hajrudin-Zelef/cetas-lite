---
id: collect-261001-automatisation-infra/automatisation-infra/dashboard-api-ansible-v2-16-16-plugins-networks-firmware-upgrades-module-html-e85e5684-1
title: "dashboard-api-ansible-v2-16-16-plugins-networks-firmware-upgrades-module-html-e85e5684"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-automatisation-infra/dashboard-api-ansible-v2-16-16-plugins-networks-firmware-upgrades-module-html-e85e5684.md
source_anchor: ""
source_lines: [1, 152]
sha256: d87453c3416627f19d1bb2a9004158ee8fd070816f778c5961d9b86ed967b4d2
---

# dashboard-api-ansible-v2-16-16-plugins-networks-firmware-upgrades-module-html-e85e5684

cisco.meraki.networks_firmware_upgrades module – Resource module for networks _firmwareupgrades
Note
This module is part of the cisco.meraki collection (version 2.16.16).
To install it, use: ansible-galaxy collection install cisco.meraki.
You need further requirements to be able to use this module,
see Requirements for details.
To use it in a playbook, specify: cisco.meraki.networks_firmware_upgrades.
New in cisco.meraki 2.16.0
Synopsis
- Manage operation update of the resource networks _firmwareupgrades.
- Update firmware upgrade information for a network.
Note
This module has a corresponding action plugin.
Requirements
The below requirements are needed on the host that executes this module.
- meraki >= 2.4.9
- python >= 3.5
Parameters
| Parameter | Comments | 
|---|---|
|  | meraki_action_batch_retry_wait_time (integer), action batch concurrency error retry wait time Default: :ansible-option-default:`60` | 
|  | meraki_api_key (string), API key generated in dashboard; can also be set as an environment variable MERAKI_DASHBOARD_API_KEY | 
|  | meraki_base_url (string), preceding all endpoint resources Default: :ansible-option-default:`"https://api.meraki.com/api/v1"` | 
|  | meraki_be_geo_id (string), optional partner identifier for API usage tracking; can also be set as an environment variable BE_GEO_ID Default: :ansible-option-default:`""` | 
|  | meraki_caller (string), optional identifier for API usage tracking; can also be set as an environment variable MERAKI_PYTHON_SDK_CALLER Default: :ansible-option-default:`""` | 
|  | meraki_certificate_path (string), path for TLS/SSL certificate verification if behind local proxy Default: :ansible-option-default:`""` | 
|  | meraki_inherit_logging_config (boolean), Inherits your own logger instance Choices: | 
|  | meraki_log_file_prefix (string), log file name appended with date and timestamp | 
|  | log_path (string), path to output log; by default, working directory of script if not specified Default: :ansible-option-default:`""` | 
|  | meraki_maximum_retries (integer), retry up to this many times when encountering 429s or other server-side errors Default: :ansible-option-default:`2` | 
|  | meraki_nginx_429_retry_wait_time (integer), Nginx 429 retry wait time Default: :ansible-option-default:`60` | 
|  | meraki_output_log (boolean), create an output log file? Choices: | 
|  | meraki_print_console (boolean), print logging output to console? Choices: | 
|  | meraki_requests_proxy (string), proxy server and port, if needed, for HTTPS Default: :ansible-option-default:`""` | 
|  | meraki_retry_4xx_error (boolean), retry if encountering other 4XX error (besides 429)? Choices: | 
|  | meraki_retry_4xx_error_wait_time (integer), other 4XX error retry wait time Default: :ansible-option-default:`60` | 
|  | meraki_simulate (boolean), simulate POST/PUT/DELETE calls to prevent changes? Choices: | 
|  | meraki_single_request_timeout (integer), maximum number of seconds for each API call Default: :ansible-option-default:`60` | 
|  | meraki_suppress_logging (boolean), disable all logging? you’re on your own then! Choices: | 
|  | meraki_use_iterator_for_get_pages (boolean), list* methods will return an iterator with each object instead of a complete list with all items Choices: | 
|  | meraki_wait_on_rate_limit (boolean), retry if 429 rate limit error encountered? Choices: | 
|  | NetworkId path parameter. Network ID. | 
|  | Contains information about the network to update. | 
|  | The network device to be updated. | 
|  | The pending firmware upgrade if it exists. | 
|  | The time of the last successful upgrade. | 
|  | The version to be updated to. | 
|  | The version ID. | 
|  | Whether or not the network wants beta firmware. Choices: | 
|  | The network device to be updated. | 
|  | The pending firmware upgrade if it exists. | 
|  | The time of the last successful upgrade. | 
|  | The version to be updated to. | 
|  | The version ID. | 
|  | Whether or not the network wants beta firmware. Choices: | 
|  | The network device to be updated. | 
|  | The pending firmware upgrade if it exists. | 
|  | The time of the last successful upgrade. | 
|  | The version to be updated to. | 
|  | The version ID. | 
|  | Whether or not the network wants beta firmware. Choices: | 
|  | The network device to be updated. | 
|  | The pending firmware upgrade if it exists. | 
|  | The time of the last successful upgrade. | 
|  | The version to be updated to. | 
|  | The version ID. | 
|  | Whether or not the network wants beta firmware. Choices: | 
|  | The network device to be updated. | 
|  | The pending firmware upgrade if it exists. | 
|  | The time of the last successful upgrade. | 
|  | The version to be updated to. | 
|  | The version ID. | 
|  | Whether or not the network wants beta firmware. Choices: | 
|  | The network device to be updated. | 
|  | The pending firmware upgrade if it exists. | 
|  | The time of the last successful upgrade. | 
|  | The version to be updated to. | 
|  | The version ID. | 
|  | Whether or not the network wants beta firmware. Choices: | 
|  | The timezone for the network. | 
|  | Upgrade window for devices in network. | 
|  | Day of the week. | 
|  | Hour of the day. | 
Notes
Note
- SDK Method used are networks.Networks.update_network_firmware_upgrades,
- Paths used are put /networks/{networkId}/firmwareUpgrades,
- Does not support check_mode
- The plugin runs on the control node and does not use any ansible connection plugins, but instead the embedded connection manager from Cisco DNAC SDK
- The parameters starting with dnac_ are used by the Cisco DNAC Python SDK to establish the connection
See Also
See also
- Cisco Meraki documentation for networks updateNetworkFirmwareUpgrades
- Complete reference of the updateNetworkFirmwareUpgrades API.
Examples
- name: Update all
  cisco.meraki.networks_firmware_upgrades:
    meraki_api_key: "{{meraki_api_key}}"
    meraki_base_url: "{{meraki_base_url}}"
    meraki_single_request_timeout: "{{meraki_single_request_timeout}}"
    meraki_certificate_path: "{{meraki_certificate_path}}"
    meraki_requests_proxy: "{{meraki_requests_proxy}}"
    meraki_wait_on_rate_limit: "{{meraki_wait_on_rate_limit}}"
    meraki_nginx_429_retry_wait_time: "{{meraki_nginx_429_retry_wait_time}}"
    meraki_action_batch_retry_wait_time: "{{meraki_action_batch_retry_wait_time}}"
    meraki_retry_4xx_error: "{{meraki_retry_4xx_error}}"
    meraki_retry_4xx_error_wait_time: "{{meraki_retry_4xx_error_wait_time}}"
    meraki_maximum_retries: "{{meraki_maximum_retries}}"
    meraki_output_log: "{{meraki_output_log}}"
    meraki_log_file_prefix: "{{meraki_log_file_prefix}}"
    meraki_log_path: "{{meraki_log_path}}"
    meraki_print_console: "{{meraki_print_console}}"
    meraki_suppress_logging: "{{meraki_suppress_logging}}"
    meraki_simulate: "{{meraki_simulate}}"
    meraki_be_geo_id: "{{meraki_be_geo_id}}"
    meraki_use_iterator_for_get_pages: "{{meraki_use_iterator_for_get_pages}}"
    meraki_inherit_logging_config: "{{meraki_inherit_logging_config}}"
    state: present
    networkId: string
    products:
      appliance:
        nextUpgrade:
          time: '2019-03-17T17:22:52Z'
          toVersion:
            id: '1001'
        participateInNextBetaRelease: false
      camera:
        nextUpgrade:
          time: '2019-03-17T17:22:52Z'
          toVersion:
            id: '1003'
        participateInNextBetaRelease: false
      cellularGateway:
        nextUpgrade:
          time: '2019-03-17T17:22:52Z'
          toVersion:
            id: '1004'
        participateInNextBetaRelease: false
      sensor:
        nextUpgrade:
          time: '2019-03-17T17:22:52Z'
          toVersion:
            id: '1005'
        participateInNextBetaRelease: false
      switch:
        nextUpgrade:
          time: '2019-03-17T17:22:52Z'
          toVersion:
            id: '1002'
        participateInNextBetaRelease: false
      wireless:
        nextUpgrade:
