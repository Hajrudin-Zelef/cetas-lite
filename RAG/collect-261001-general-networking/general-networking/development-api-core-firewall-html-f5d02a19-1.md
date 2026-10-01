---
id: collect-261001-general-networking/general-networking/development-api-core-firewall-html-f5d02a19-1
title: "Firewallï"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/development-api-core-firewall-html-f5d02a19.md
source_anchor: ""
source_lines: [1, 183]
sha256: 3cd65fc47069aada640e09571f0b5aec5daaf809cdfa059a572d901c1557357f
---

# Firewallï

The firewall API offers a way for machine to machine interaction between custom applications and OPNsense, it is part of the core system.

Although the module does contains a basic user interface (in ), itâs mirely intended as a reference and testbed. Thereâs no relation to any of the rules being managed via the core system.

Tip

Use your browsers âinspectâ feature to compare requests easily, the user interface in terms of communication is exactly the same as offered by the API . Rules not visible in the web interface () will not be returned by the API either.

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `POST` | firewall | alias | add_item |  | 
| `POST` | firewall | alias | del_item | $uuid | 
| `GET,POST` | firewall | alias | export |  | 
| `GET` | firewall | alias | get |  | 
| `GET` | firewall | alias | get_alias_u_u_i_d | $name | 
| `GET` | firewall | alias | get_geo_i_p |  | 
| `GET` | firewall | alias | get_item | $uuid=null | 
| `GET` | firewall | alias | get_table_size |  | 
| `POST` | firewall | alias | import |  | 
| `GET` | firewall | alias | list_categories |  | 
| `GET` | firewall | alias | list_countries |  | 
| `GET` | firewall | alias | list_network_aliases |  | 
| `GET` | firewall | alias | list_user_groups |  | 
| `POST` | firewall | alias | reconfigure |  | 
| `GET,POST` | firewall | alias | search_item |  | 
| `POST` | firewall | alias | set |  | 
| `POST` | firewall | alias | set_item | $uuid | 
| `POST` | firewall | alias | toggle_item | $uuid,$enabled=null | 
| `POST` | firewall | alias | update | $action=null | 
| `<<uses>>` |  |  |  | *model* Alias.xml | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `POST` | firewall | alias_util | add | $alias | 
| `GET` | firewall | alias_util | aliases |  | 
| `POST` | firewall | alias_util | delete | $alias | 
| `POST` | firewall | alias_util | find_references |  | 
| `POST` | firewall | alias_util | flush | $alias | 
| `GET` | firewall | alias_util | list | $alias | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `POST` | firewall | category | add_item |  | 
| `POST` | firewall | category | del_item | $uuid | 
| `GET` | firewall | category | download |  | 
| `GET` | firewall | category | get |  | 
| `GET` | firewall | category | get_item | $uuid=null | 
| `GET,POST` | firewall | category | search_item | $add_empty=0 | 
| `POST` | firewall | category | set |  | 
| `POST` | firewall | category | set_item | $uuid | 
| `POST` | firewall | category | upload |  | 
| `<<uses>>` |  |  |  | *model* Category.xml | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `POST` | firewall | d_nat | add_rule |  | 
| `POST` | firewall | d_nat | del_rule | $uuid | 
| `GET` | firewall | d_nat | get_rule | $uuid=null | 
| `GET` | firewall | d_nat | move_rule_before | $selected_uuid,$target_uuid | 
| `GET,POST` | firewall | d_nat | search_rule |  | 
| `POST` | firewall | d_nat | set_rule | $uuid | 
| `POST` | firewall | d_nat | toggle_rule | $uuid,$disabled=null | 
| `GET` | firewall | d_nat | toggle_rule_log | $uuid,$log | 
| `<<uses>>` |  |  |  | *model* DNat.xml | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `POST` | firewall | filter_base | apply |  | 
| `GET` | firewall | filter_base | get |  | 
| `GET` | firewall | filter_base | list_categories |  | 
| `GET` | firewall | filter_base | list_network_select_options |  | 
| `GET` | firewall | filter_base | list_port_select_options |  | 
| `POST` | firewall | filter_base | set |  | 
| `<<uses>>` |  |  |  | *model* Filter.xml | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `POST` | firewall | filter | add_rule |  | 
| `POST` | firewall | filter | del_rule | $uuid | 
| `GET` | firewall | filter | download_rules |  | 
| `POST` | firewall | filter | flush_inspect_cache |  | 
| `GET` | firewall | filter | get_interface_list |  | 
| `GET` | firewall | filter | get_rule | $uuid=null | 
| `POST` | firewall | filter | move_rule_before | $selected_uuid,$target_uuid | 
| `GET` | firewall | filter | search_rule |  | 
| `POST` | firewall | filter | set_rule | $uuid | 
| `POST` | firewall | filter | toggle_rule | $uuid,$enabled=null | 
| `GET` | firewall | filter | toggle_rule_log | $uuid,$log | 
| `POST` | firewall | filter | upload_rules |  | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `GET` | firewall | filter_util | rule_stats |  | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `POST` | firewall | group | add_item |  | 
| `POST` | firewall | group | del_item | $uuid | 
| `GET` | firewall | group | get |  | 
| `GET` | firewall | group | get_item | $uuid=null | 
| `POST` | firewall | group | reconfigure |  | 
| `GET,POST` | firewall | group | search_item |  | 
| `POST` | firewall | group | set |  | 
| `POST` | firewall | group | set_item | $uuid | 
| `<<uses>>` |  |  |  | *model* Group.xml | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `GET` | firewall | migration | download_rules |  | 
| `POST` | firewall | migration | flush |  | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `POST` | firewall | npt | add_rule |  | 
| `POST` | firewall | npt | del_rule | $uuid | 
| `GET` | firewall | npt | get_rule | $uuid=null | 
| `GET` | firewall | npt | move_rule_before | $selected_uuid,$target_uuid | 
| `GET,POST` | firewall | npt | search_rule |  | 
| `POST` | firewall | npt | set_rule | $uuid | 
| `POST` | firewall | npt | toggle_rule | $uuid,$enabled=null | 
| `GET` | firewall | npt | toggle_rule_log | $uuid,$log | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `POST` | firewall | one_to_one | add_rule |  | 
| `POST` | firewall | one_to_one | del_rule | $uuid | 
| `GET` | firewall | one_to_one | get_rule | $uuid=null | 
| `GET` | firewall | one_to_one | move_rule_before | $selected_uuid,$target_uuid | 
| `GET,POST` | firewall | one_to_one | search_rule |  | 
| `POST` | firewall | one_to_one | set_rule | $uuid | 
| `POST` | firewall | one_to_one | toggle_rule | $uuid,$enabled=null | 
| `GET` | firewall | one_to_one | toggle_rule_log | $uuid,$log | 

| ï |  |  |  |  | 
|---|---|---|---|---|
| Method | Module | Controller | Command | Parameters | 
|---|---|---|---|---|
| `POST` | firewall | source_nat | add_rule |  | 
| `POST` | firewall | source_nat | del_rule | $uuid | 
| `GET` | firewall | source_nat | get_rule | $uuid=null | 
| `GET` | firewall | source_nat | move_rule_before | $selected_uuid,$target_uuid | 
| `GET,POST` | firewall | source_nat | search_rule |  | 
| `POST` | firewall | source_nat | set_rule | $uuid | 
| `POST` | firewall | source_nat | toggle_rule | $uuid,$enabled=null | 
| `GET` | firewall | source_nat | toggle_rule_log | $uuid,$log | 

## Conceptï

The firewall plugin injects rules in the standard OPNsense firewall while maintaining visibility on them in the standard user interface.

We use our standard `ApiMutableModelControllerBase` to allow crud operations on rule entries and offer an
`apply` action to activate the new configuration.

The diagram above contains the basic steps to change rules and activate them.
Changes made through the administrative endpoints are staged in the configuration; calling `apply()` reloads
the firewall so the new ruleset becomes active.

Note

