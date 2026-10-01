---
id: collect-261001-general-networking/general-networking/development-examples-helloworld-html-da9a1fe4-3
title: "Hello world module & pluginï"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "parameters"]
source: docs/RAG/collect-261001-general-networking/development-examples-helloworld-html-da9a1fe4.md
source_anchor: ""
source_lines: [364, 541]
sha256: eea06f778f6de536e6401474a53c49ddef207003920415c38bff0703241c5fa4
---

# Hello world module & pluginï

Our basic module provides a way to read and modify configuration data using the web interface (and in time also other consumers using the api). Next step is to add some activity to our system, all backend applications should use their own configuration, which in real life we would keep as standard as possible.

For our example we will follow the same process as for any other service and start writing some configuration data for our sample application. Which means, creating a template and hooking it into our save action.

Our example will write a simple configuration file, stored in /usr/local/etc/helloworld/helloworld.conf

The configd system is responsible for updating the contents of that file when requested, it does so by using a definition found in its template folder. This sample will use the following path to store the backend templates:

```
/usr/local/opnsense/service/templates/OPNsense/HelloWorld/
```
First we add a content definition, by creating a file named +TARGETS, which should hold the following information:

```
helloworld.conf:/usr/local/etc/helloworld/helloworld.conf
```
This basically tells the engine that there will be a file in the same folder named âhelloworld.confâ which provides, together with config.xml, data for the file in /usr/local/etc/helloworld/helloworld.conf

Next thing to do is create that helloworld.conf file in the templates directory. We will keep things very simple for this one and just copy in our data into an ini file structured configuration, when the module is enabled.

```
{% if not helpers.empty('OPNsense.helloworld.general.Enabled') %}
[general]
SMTPHost={{ OPNsense.helloworld.general.SMTPHost|default("") }}
FromEmail={{ OPNsense.helloworld.general.FromEmail|default("") }}
ToEmail={{ OPNsense.helloworld.general.ToEmail|default("") }}
Subject={{ OPNsense.helloworld.general.Description|default("") }}
{% endifÂ %}
```
Now we need to be able to reload this module (or in real life, this would probably be a service) by adding a service action into our ServiceController. Edit controllers/OPNsense/HelloWorld/Api/ServiceController.php and add the backend module to the use section, like this:

```
use \OPNsense\Core\Backend;
```
By doing this we can use the backend communication in this class. Next add a new action to the class called âreloadActionâ using this piece of code:

```
public function reloadAction()
{
    $status = "failed";
    if ($this->request->isPost()) {
        $status = strtolower(trim((new Backend())->configdRun('template reload OPNsense/HelloWorld')));
    }
    return ["status" => $status];
}
```
This validates the type of action (it should always be POST to enable CSRF protection) and adds a backend action for reloading the template. When successful the action will return âstatusâ:âokâ as json object back to the client.

Now we are able to refresh the template content, but the user interface doesnât know about it yet. To hook loading of the template into the save action, we will go back to the index.volt view and add the following jQuery / framework code between the braces of âsaveFormToEndPointâ.

```
ajaxCall(url="/api/helloworld/service/reload", sendData={},callback=function(data,status) {
    // action to run after reload
});
```
If you save the form now (when enabled), you should see a new file in

```
helloworld.conf:/usr/local/etc/helloworld/helloworld.conf
```
Containing something like this:

```
[general]
SMTPHost=127.0.0.1
FromEmail=sample@example.com
ToEmail=sample@example.com
Subject=test
```
What have we accomplished now, we can input data, validate it and save it to the corresponding format of the actual service or application, which uses this data. So if you have a third party application, which you want to integrate into the user interface. You should be able to generate what it needs now. (Thereâs more to learn, but these are the basics).

But how do should we control that third part program now? Thatâs the next step.

## Controlling the sampleï

Instead of running all kinds of shell commands directly from the PHP code, which very often need root access (starting/stopping services, etc.), we should always communicate to our backend process which holds templates of possible things to run and protects your system from executing arbitrary commands.

Another advantage of this approach is that all commands defined here, can also be ran from the command line of the firewall providing easier serviceability. For example, the command to refresh the helloworld configuration can be run from the command line by running:

```
configctl template reload OPNsense/HelloWorld
```
First thing to do when registering new actions to the system for a new application is to create a config template.

```
/usr/local/opnsense/service/conf/actions.d/actions_helloworld.conf
```
And add a command to the template like this:

```
[test]
command:/usr/local/opnsense/scripts/helloworld/testConnection.py
parameters:
type:script_output
message:hello world module test
```
Letâs test our new command by restarting configd from the command line:

```
service configd restart
```
And test our new command using:

```
configctl helloworld test
```
Which should return some response in json format.

Next step is to use this command in our controller (middleware), just like we did with the template action. For consistency we call our action testAction and let it pass json data to our clients when using a POST type request.

```
public function testAction()
{
    if ($this->request->isPost()) {
        $bckresult = json_decode(trim((new Backend())->configdRun("helloworld test")), true);
        if ($bckresult !== null) {
            // only return valid json type responses
            return $bckresult;
        }
    }
    return ["message" => "unable to run config action"];
}
```
And now we can make our user interface aware of the action, place a button and link an action in the index.volt. Using the following elements:

```
$("#testAct").SimpleActionButton({
    onAction: function(data) {
        $("#responseMsg").removeClass("hidden").html(data['message']);
    }
});
```
(in HTML section)

```
<div class="alert alert-info hidden" role="alert" id="responseMsg">
</div>
<button class="btn btn-primary" id="testAct" data-endpoint="/api/helloworld/service/test" data-label="{{ lang._('Test') }}"></button>
```
Tip

As you might have noticed the `testAct` button uses a different method to call an endpoint, it uses the SimpleActionButton
wrapper which eases the implementation of simple actions

Now go back to the page and save some data using the save button, next press test to see some results.

## Multi language / Translationsï

OPNsense is available in may different languages like english, german or japanese. This works because we are using the gettext library which is available to all GUI components. While the XML based user interfaces are supporting it automatically, there may still the need to call it manually (buttons, tabs etc.).

If you have a static string, you should add it like this into a classic PHP page:

```
<?= gettext('your string here') ?>
```
And this way into a volt template:

```
{{ lang._('your string here') }}
```
If your string is not only plaintext because it contains non-static words, HTML tags and other dynamic content, you need to use a format string. This way, you can use placeholders for such elements which should not land in the translation file.

For php it works this way:

```
<?= sprintf(gettext('your %s here'), $data) ?>
```
And for volt templates it works this way:

```
{{ lang._('your %s here') | format(data) }}
```
Note

You should NEVER split strings which should belong together like a sentence. This makes plugins hard to translate and will decrease the quality of OPNsense in other languages.

## Plugin to access control (ACL)ï

