---
id: collect-261001-general-networking/general-networking/development-examples-helloworld-html-da9a1fe4-4
title: "Hello world module & pluginï"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/development-examples-helloworld-html-da9a1fe4.md
source_anchor: ""
source_lines: [542, 619]
sha256: 1d7ddc2ce0682ba15f0b7e366b5624438ace086a7146b953b0b0c91d74e9628e
---

# Hello world module & pluginï

If we want to authorize users to access this module, we can add an ACL to this module. Without it, only admin users can access it. Create an XML file in the model directory name ACL/ACL.xml and place the following content in it:

```
<acl>
    <!-- unique acl key, must be globally unique for all ACLs  -->
    <page-user-helloworld>
        <name>WebCfg - Users: Hello World! </name>
        <description>Allow access to the Hello World! module</description>
        <patterns>
            <pattern>ui/helloworld/*</pattern>
            <pattern>api/helloworld/*</pattern>
        </patterns>
    </page-user-helloworld>
</acl>
```
This creates an ACL key named âpage-user-helloworldâ which authorizes access to both the ui and API urls of this application. You can now grant access to this module from the system user manager.

The ACL system is subject to caching, so you may not see your changes
in the user manager page yet. Delete the
`/tmp/opnsense_acl_cache.json` file, if it exists. Now, when you
refresh the user manager page, you should see that the new ACL is
available to be assigned.

## Create an installable pluginï

All files are created in their original locations (on the OPNsense machine /usr/local/â¦), now we are ready to create a package from them. To fully use this process and create the actual package, itâs best to setup a full build environment (explained over here: https://github.com/opnsense/tools )

When everything is in place, we will create a new plugin directory. For this example we will use the following:

```
/usr/plugins/devel/helloworld/
```
Add a new Makefile, containing the information for our plugin:

```
     helloworld
PLUGIN_VERSION=        1.0
PLUGIN_COMMENT=        A sample framework application
#PLUGIN_DEPENDS=
PLUGIN_MAINTAINER= user@domain
.include "../../Mk/plugins.mk"
```
```
/usr/plugins/devel/helloworld/src/
```
Next copy all files created and located in /usr/local/ into this new src directory, which results in the following file listing:

```
src/opnsense/mvc/app/controllers/OPNsense/HelloWorld/Api/ServiceController.php
src/opnsense/mvc/app/controllers/OPNsense/HelloWorld/Api/SettingsController.php
src/opnsense/mvc/app/controllers/OPNsense/HelloWorld/IndexController.php
src/opnsense/mvc/app/controllers/OPNsense/HelloWorld/forms/general.xml
src/opnsense/mvc/app/models/OPNsense/HelloWorld/ACL/ACL.xml
src/opnsense/mvc/app/models/OPNsense/HelloWorld/HelloWorld.php
src/opnsense/mvc/app/models/OPNsense/HelloWorld/HelloWorld.xml
src/opnsense/mvc/app/models/OPNsense/HelloWorld/Menu/Menu.xml
src/opnsense/mvc/app/views/OPNsense/HelloWorld/index.volt
src/opnsense/scripts/helloworld/testConnection.py
src/opnsense/service/templates/OPNsense/HelloWorld/+TARGETS
src/opnsense/service/templates/OPNsense/HelloWorld/helloworld.conf
src/opnsense/service/conf/actions.d/actions_helloworld.conf
```
With everything in place, you could build the plugin package using the âmake pluginsâ command in the /usr/tools directory. The result of this will be a standard pkg package, which you can install on any OPNsense system and will be usable right after installing. All plugins are prefixed with os-, our new package file will be called:

```
os-helloworld-1.0.txz
```
(-1.0 comes from the version in the makefile)

Reference

- source of this exampleÂ : https://github.com/opnsense/plugins/tree/master/devel/helloworld
- build instructionsÂ : https://github.com/opnsense/tools
- practical frontend development : https://github.com/opnsense/ui_devtools
- frontend template language reference (Volt)Â : https://docs.phalcon.io/latest/volt/
- configuration template language reference (mostly the same as Volt)Â : https://jinja.palletsprojects.com/en/stable/
- OPNsense architecture Architecture
- OPNsense creating models Develop:Frontend/Creating_models
