---
id: collect-261001-meraki/meraki/ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-callbacks-md-8e872e6a-1
title: "ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-callbacks-md-8e872e6a"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-callbacks-md-8e872e6a.md
source_anchor: ""
source_lines: [1, 170]
sha256: 7ad5cf31d1691d54418abca76e07803a82dfced6878e0efb04cecd463b91fa87
---

# ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-callbacks-md-8e872e6a

A "Callback" is a programming feature that provides notifications upon the completion of a long-running operation. Unlike standard API calls that return immediately, callbacks issue results via a webhook when they become available. Aligned with the OpenAPI v3 specification, promoting standardized asynchronous operations.
A webhook is an outbound API request that Meraki will send to a third party service. Callbacks use the Meraki webhook system, supporting both pre-configured receivers and dynamic URLs. In addition, webhook templates can format the HTTP message body and headers for customized integrations.
Callbacks are a feature that is designed to handle operations that require time to process. They are useful in scenarios where immediate feedback is not essential, but timely notification upon completion is critical. By using callbacks, you optimize resource usage, reduce the need for continuous polling, and streamline workflows.
- 
Bulk Configuration: When pushing configuration changes to hundreds or thousands of devices, waiting for each action batch to respond synchronously can be time-consuming or rate-limit challenging. With callbacks, you initiate the asynchronous requests and receive updates as each job completes its configurations.
- 
Monitoring: In large networks, continuously polling each request for status can be inefficient and can generate significant API usage. Callbacks enable an event-driven approach, where results can be sent to a group messaging service or database as they are available.
- 
Event-Driven Automation: Use callbacks to trigger other processes or workflows within your system. For example, once a request has completed and a webhook has been sent, an automated system can continue its next operation using those results.
"Alerts" and "Callbacks" are distinct features within the Meraki API, each serving a different purpose, and following different data schemas.
Webhook Alerts are event-driven notifications for situations such as a device going offline, sensor detections or a security event.
{
  "version": "0.1",
  "sharedSecret": "secret",
  "sentAt": "2022-09-07T14:14:14.591941Z",
  "organizationId":"123",
  "organizationName":"Org in Company",
  "alertType": "APs went down",
  "alertLevel": "critical",
  "alertData": {
  ...
  }
}
Callbacks deliver the outcomes of specific asynchronous API operations such as with Live Tools or Action Batches.
{
  "callbackId": "643451796760560204",
  "organization": {
    "id": "123",
    "name": "Org in Company"
  },
  "message": {
    "pingId": "643451796835419513",
    "status": "complete",
	...
  }
}
- 
Alerts: Identified by the alertId and containing the dynamic data in thealertData property.
- 
Callbacks: Identified by the callbackId and containing the dynamic data in themessage property.
Differentiate between alerts and callbacks in your service, possibly using tailored webhook templates for each.
Here's an example Liquid template to handle the difference in wehbhook payload types:
{% if callbackId %}
  {# Handle as Callback #}
{% elsif alertId %}
  {# Handle as Webhook Alert #}
{% else %}
  {# Unknown payload type #}
{% endif %}
Learn more about customizing the shape and security of your callback webhooks.
Webhook Payload Templates Guide
This endpoint provides detailed status information about a specific callback, including its current state, any errors encountered, the initiating user, and the related webhook details.
Operation: GET /organizations/{organizationId}/webhooks/callbacks/statuses/{callbackId}
- callbackId: Unique identifier of the callback.
- status: Current status of the callback. Possible values are:
  - completed : Indicates the callback operation has successfully concluded.
  - failed : Signifies an error or failure in the callback process.
  - running : The callback is still in progress.
- errors: Array of error messages, if any, encountered during the callback process.
- createdBy: Information about the user who initiated the callback, including their admin ID.
- webhook: Detailed information about the webhook used for the callback, including:
  - url: The receiver URL for the callback results.
  - httpServer: Information about the HTTP server that receives the callback data.
  - payloadTemplate: Details about the payload template used for the callback.
  - sentAt: Timestamp indicating when the callback was dispatched to the webhook receiver.
An example JSON response from the /callbacks/statuses endpoint:
{
  "callbackId": "1284392014819",
  "status": "completed",
  "errors": ["Callback failed"],
  "createdBy": {
    "adminId": "212406"
  },
  "webhook": {
    "url": "https://webhook.site/28efa24e-f830-4d9f-a12b-fbb9e5035031",
    "httpServer": {
      "id": "aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vd2ViaG9va3M="
    },
    "payloadTemplate": {
      "id": "wpt_2100"
    },
    "sentAt": "2018-02-11T00:00:00.090210Z"
  }
}
The following endpoints are examples within the Meraki Dashboard API that support callbacks, enabling asynchronous operations:
- 
Ping: POST /devices/{serial}/liveTools/ping
- 
Ping: POST /devices/{serial}/liveTools/pingDevice
- 
Wake on LAN: POST /devices/{serial}/liveTools/wakeOnLan
- Action Batches: POST /organizations/{organizationId}/actionBatches
This example demonstrates initiating a ping operation with the result being sent to a specified callback URL:
Operation: POST /devices/{serial}/liveTools/ping
Request Body:
{
  "target": "8.8.8.8",
  "count": 4,
  "callback": {
    "url": "https://webhook.site/your-custom-url"
  }
}
An example JSON payload from a webhook callback following a ping operation:
{
  "callbackId": "643451796760560204",
  "organization": {
    "id": "321321",
    "name": "My Laboratory"
  },
  "network": {
    "id": "643451796760500000",
    "name": "Branch 1"
  },
  "sentAt": "2023-09-15T06:46:17-07:00",
  "message": {
    "pingId": "643451796835419513",
    "url": "/devices/Q2BX-9QRR-XXXX/liveTools/ping/64345179683500000",
    "request": {
      "serial": "Q2BX-9QRR-XXXX",
      "target": "8.8.8.8",
      "count": "3"
    },
    "status": "complete",
    "results": {
      // Results of the ping operation...
    }
  }
}
Demo
This example illustrates the use of callbacks in conjunction with an Action Batch operation. Action Batches in the Meraki Dashboard API allow for the execution of multiple actions in a single API call, which can be significantly enhanced with the use of callbacks for asynchronous processing and status updates.
Suppose we want to update the settings of multiple networks in an organization. Given the potentially long processing time, we'll use a callback to receive status updates on the batch operation.
Operation: POST /organizations/{organizationId}/actionBatches
{
  "organizationId": "123456",
  "confirmed": true,
  "synchronous": false,
  "actions": [
    {
      "resource": "/networks/{networkId1}/ssids/0",
      "operation": "update",
      "body": {
        "name": "New SSID Name 1",
        "enabled": true
      }
    },
    {
      "resource": "/networks/{networkId2}/ssids/1",
      "operation": "update",
      "body": {
        "name": "New SSID Name 2",
        "enabled": false
      }
    }
  ],
  "callback": {
    "url": "https://webhook.site/your-custom-callback-url"
  }
}
- organizationId: The identifier of the organization.
- confirmed: Indicates whether the action batch is confirmed for execution.
- synchronous: Set to false to indicate the operation is asynchronous and will utilize callbacks.
- actions: An array of actions to be executed. Each action specifies a resource to modify, an operation (like "update"), and the body with the new settings.
- callback: Object containing the callback URL where status updates will be sent.
When the action batch is processed, a callback is sent to the specified URL with the following example payload:
{
  "callbackId": "874399285760560205",
  "organization": {
    "id": "123456",
    "name": "Global Tech Inc."
  },
