---
id: collect-261001-general-networking/general-networking/development-examples-helloworld-html-da9a1fe4-1
title: "Hello world module & pluginï"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/development-examples-helloworld-html-da9a1fe4.md
source_anchor: ""
source_lines: [1, 161]
sha256: 3412f641fcb546392602bdb7bb4f2751cc5a6bf7fb5934317f7e2d2625a66fea
---

# Hello world module & pluginï

## Goalï

Goal for this sample

The goal of the âHello worldâ module weâre creating in the example is to control a program on our system named âtestConnection.pyâ, which is part of the example package available on GitHub. It will try to send an email using plain smtp and will respond with a json text message about the result of that attempt.

Our application will need some settings to operate correctly, like an ip address and an email address and we need to be able to run that application. Because this application returns some valuable data for our users, we need to be able to fetch the response data back.

## Guidelinesï

Guidelines and coding style

For all OPNsense modules and applications there are some basic style and coding guides which you should use.

### Namingï

When creating modules for OPNsense, always name your components like this: VendorName/ModuleName

In our sample case this will be: OPNsense/HelloWorld

### Architectureï

Always make sure thereâs a clear separation of concerns, back-end calls (like shell scripts) should be implemented using the configd system, all communication to the client should be handled from an API endpoint. (the example provides more insights on how this works).

Back-end programs should not access the config.xml directly, if data is needed let the template system handle the desired output (most applications, daemons and tools deliver their own desired configuration format). Thereâs generally no good reason to avoid the standards that are already there.

If you follow this basic rules, youâre automatically building a command structure for the system administrators and provide a connector to third party tools to the API of your component.

## Skeletonï

Setup a skeleton for the frontend / middleware

First step for our project is to build a skeleton which holds the structure for our frontend/middleware. Do keep in mind to only build the structure, if you add empty files this will cause errors.

### Modelï

For our sample application we want to setup some configuration data, which for all new style projects should live in itâs own model.

First we start by creating two files inside the models/OPNsense/HelloWorld directory.

The first one is the boilerplate for the model class, which should contain model specific methods and (by deriving it from BaseModel) automatically understands the second file.

```
<?php
namespace OPNsense\HelloWorld;
use OPNsense\Base\BaseModel;
class HelloWorld extends BaseModel
{
}
```
Not all modules contain additional code in the PHP class, sometimes all the standard behaviour is already sufficient for your modules/application.

Note

When stored data has derivatives, the model is usually the place to build these. Good examples are a flag to figure out if a service is enabled in cases where it depends on structures inside the data. (such as a list of interfaces which can all be enabled and/or disabled)

Which is the model XML template, our skeleton starts with something like this:

```
<model>
    <mount>//OPNsense/helloworld</mount>
    <description>the OPNsense "Hello World" application</description>
    <items>
        <!-- container -->
    </items>
</model>
```
The content of the mount tag is very important, this is the location within the config.xml file where this model is responsible. Other models cannot write data into the same area. You should name this location with your vendor and module name to make sure others could easily identify it.

Use the description tag to identify your model, the last tag in place is the items tag, where the actual definition will live. We leave it empty for now as we proceed with the next step of creating our skeleton.

### Viewï

Page template (View)

We should add a (Volt) template to use for the index page of our module; we will use the same naming convention here.

Create a template named index.volt inside the views/OPNsense/HelloWorld directory containing the following data:

```
<h1>Hello World!</h1>
```
### Controllerï

Next step is to add controllers, which will be automatically picked up by the system routing. A controller connects the user interaction to logic and presentation.

Every OPNsense module should separate presentation from logic, thatâs why there should always be multiple controllers per module.

Our first controller handles the template rendering to the user and connects the user view we just created. We start by creating a PHP file in controllers/OPNsense/HelloWorld/ with the following name IndexController.php and contents:

```
<?php
namespace OPNsense\HelloWorld;
class IndexController extends \OPNsense\Base\IndexController
{
    public function indexAction()
    {
        // pick the template to serve to our users.
        $this->view->pick('OPNsense/HelloWorld/index');
    }
}
```
At this point you should be able to test if your work so far was successful, by going to the following location (after being logged in to the firewall as root user):

```
http[s]://<your ip>/ui/helloworld/
```
Which should serve you the âHello World!â text youâve added in the template.

Next two controllers we are going to create are to be used for the api to the system, they should take care of service actions and the retrieval/changing of configuration data.

They should live in a subdirectory of the controller called Api and extend the corresponding class.

For our modules we create two API controllers, one for controlling settings and one for performing service actions. (Named SettingsController.php and ServiceController.php)

```
<?php
namespace OPNsense\HelloWorld\Api;
use \OPNsense\Base\ApiMutableModelControllerBase;
class SettingsController extends ApiMutableModelControllerBase
{
}
```
```
<?php
namespace OPNsense\HelloWorld\Api;
use \OPNsense\Base\ApiMutableServiceControllerBase;
class ServiceController extends ApiMutableServiceControllerBase
{
}
```
Note

For the sake of simplicity we use `ApiMutableModelControllerBase` in our example, in practice this
class derives from `ApiControllerBase` which is the essential requirement for an api endpoint.
Using `ApiMutableModelControllerBase` prevents duplicating boilerplate code.

## First Input Formï

Building your first input form

The first step in building forms is to determine what information we should collect.

Our simple application will send an email using data in our configuration xml. For this very module we want to collect the following:

| Property | Default | Description | 
|---|---|---|
| General.Enabled | Enabled (1) | Should this module be enabled (Boolean) | 
| General.SMTPHost | <empty> | IP address for the remote smtp host | 
| General.FromEmail | sample@example.com | Email address of the sender | 
| General.ToEmail | <empty> | Email address to send our test email to | 
| General.Description | <empty> | Description, used as subject of the email. | 

### Adding Fieldsï

Adding fields to your model

When building the skeleton, we have created an empty model (XML), which we are going to fill with some attributes now. The items section of the model XML should contain the structure you want to use for your application, you can create trees to hold data in here. All leaves should contain a field type to identify and validate itâs content. The list of attributes for our application can be translated to this:

