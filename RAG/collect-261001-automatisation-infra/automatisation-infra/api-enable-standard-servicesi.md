---
id: collect-261001-automatisation-infra/automatisation-infra/api-enable-standard-servicesi
title: "API enable standard servicesï"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-automatisation-infra/api-enable-standard-servicesi.md
source_anchor: ""
source_lines: [1, 104]
sha256: 7324b0066bb5f1635844dd61674bd28e59acab87fa3d20fcce93b4ee81116cf9
---

# API enable standard servicesï

OPNsense contains a simple wrapper which handles standard service actions like starting and stopping services. For this example, we assume the HelloWorld example is created and the model exists.

## Prerequisitesï

The HelloWorld example forms the basis for this one, please make sure you read and try it first before starting with this one.

## Configd actionsï

Since the example didnât contain a service, weâre going to extend the configd template first.
Edit `/usr/local/opnsense/service/conf/actions.d/actions_helloworld.conf` and expand with the following sections:

```
[start]
command:exit 0
parameters:
type:script
message:hello world service start
[stop]
command:exit 0
parameters:
type:script
message:hello world service stop
[restart]
command:exit 0
parameters:
type:script
message:hello world service restart
[reload]
command:exit 0
parameters:
type:script
message:hello world service restart
[status]
command: echo "hello world is running"
parameters:
type:script_output
message:hello world service status
```
Next restart configd using `service configd restart` and test these new calls using the following commands

```
# configctl helloworld start
OK
# configctl helloworld stop
OK
# configctl helloworld status
hello world is running
# configctl helloworld restart
OK
```
Our template only simulates a service, it doesnât actually do anything (`exit 0`).

## Update the service controllerï

Next we change the existing controller to use `ApiMutableServiceControllerBase`, which links the existing model to the service callouts defined.
The `testAction` used in the HelloWorld example is left out to avoid confusion.

```
use OPNsense\Base\ApiMutableServiceControllerBase;
class ServiceController extends ApiMutableServiceControllerBase
{
    protected static $internalServiceClass = '\OPNsense\HelloWorld\HelloWorld';
    protected static $internalServiceTemplate = 'OPNsense/HelloWorld';
    protected static $internalServiceEnabled = 'general.enabled';
    protected static $internalServiceName = 'helloworld';
    protected function reconfigureForceRestart()
    {
        return 0;
    }
}
```
The service above defines the following static variables:

  - $internalServiceClass
    - reference the model class, which is used to determine if this service is enabled (links the model to the service)
  - $internalServiceTemplate
    - before starting the service it will call configd to generate configuration data, in this case it would execute the equivalent of `configctl template reload OPNsense/HelloWorld` on the console
  - $internalServiceEnabled
    - Which section of the model contains a boolean defining if the service is enabled (general.enabled)
  - $internalServiceName
    - refers to the action template, where it can find start/stop/restart/status/reload actions (actions_helloworld.conf)

The `reconfigureForceRestart` overwrite tells the controller if it should always stop the service before trying a start, some
services can be signaled to do a reconfigure without stopping.

Tip

When a service may register new  interfaces, one could implement the protected method
`invokeInterfaceRegistration` and make it return `true`, after which the plugin system will detect changes to the
interface registrations automatically.

## Endpointsï

When deploying this controller into the HelloWorld module it would expose the following endpoints

| Method | Endpoint | 
|---|---|
| `POST` | /api/helloworld/service/stop | 
| `POST` | /api/helloworld/service/start | 
| `POST` | /api/helloworld/service/restart | 
| `POST` | /api/helloworld/service/reconfigure | 
| `GET` | /api/helloworld/service/status |
