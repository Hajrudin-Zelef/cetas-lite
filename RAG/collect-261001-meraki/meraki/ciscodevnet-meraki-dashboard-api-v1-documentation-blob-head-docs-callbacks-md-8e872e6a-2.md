---
id: collect-261001-meraki/meraki/ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-callbacks-md-8e872e6a-2
title: "ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-callbacks-md-8e872e6a"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-meraki/ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-callbacks-md-8e872e6a.md
source_anchor: ""
source_lines: [171, 223]
sha256: 4b320dfc3e39142d64c84a2cf06eca808a3dd107ff1460d6ed82b61ef58cc3c7
---

# ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-callbacks-md-8e872e6a

  "actionBatchId": "874399285760580000",
  "sentAt": "2023-10-01T09:30:00-07:00",
  "message": {
    "status": "completed",
    "results": {
      "actionsCompleted": 2,
      "actionsFailed": 0,
      "details": [
        {
          "networkId": "networkId1",
          "resource": "/networks/{networkId1}/ssids/0",
          "operation": "update",
          "status": "success"
        },
        {
          "networkId": "networkId2",
          "resource": "/networks/{networkId2}/ssids/1",
          "operation": "update",
          "status": "success"
        }
      ]
    }
  }
}
- callbackId: The unique identifier for the callback.
- organization: Details of the organization.
- actionBatchId: Identifier of the action batch.
- sentAt: Timestamp of when the callback was sent.
- message: Contains the status of the action batch and detailed results.
Demo
To aid in troubleshooting, use the following monitoring tools and endpoints:
- 
Callback Status Operation: Use the /callbacks/statuses endpoint to check the delivery status of callbacks.
- 
Webhook Logs: Review the webhook logs for detailed information on the delivery and any errors encountered.
- 
Postman: Setup a mock server to test your webhooks.
- 
Webhook.site: This free service provides a public URL for testing webhooks.
- Verify Operation Configuration: Ensure that the callback URL is correct and the server can receive POST requests.
- Check for Typos: Review the callback URL and payload for any typographical errors.
- Inspect Firewall Settings: Confirm that your network allows inbound connections on the port that is used by your webhook receiver.
- Validate Payload: Check if the payload structure matches the expected schema.
- Check for Payload Changes: Ensure that there haven't been changes in the payload template structure.
- Interpret Error Messages: Use the error messages received in the callbacks to understand what went wrong.
- Consult Documentation: Refer to the API documentation of the specific endpoint for detailed error descriptions.
- Network Latency: Test your network for latency issues.
- Service Overload: Verify if the receiving server is not overwhelmed with requests.
- Server Log Analysis: Check the server logs to identify processing issues.
- Code Review: Look over the code handling the callbacks to find logical errors or exceptions.
If after troubleshooting you are still facing issues:
- Meraki Community: Engage with the Meraki community for insights and advice.
- Meraki Support: Open a support ticket with Meraki, providing detailed information about the issue and the steps already taken.
