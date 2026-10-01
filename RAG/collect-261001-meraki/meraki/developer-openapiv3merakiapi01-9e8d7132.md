---
id: collect-261001-meraki/meraki/developer-openapiv3merakiapi01-9e8d7132
title: "developer-openapiv3merakiapi01-9e8d7132"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["governance"]
source: docs/RAG/collect-261001-meraki/developer-openapiv3merakiapi01-9e8d7132.md
source_anchor: ""
source_lines: [1, 23]
sha256: 95cfcf2c256292a1121003308105045668b5a65af70e53af05f5ff8f99ae372f
---

# developer-openapiv3merakiapi01-9e8d7132

Meraki’s API has transformed significantly starting with the adoption of Swagger, now known as OpenAPI v2. The integration of the OpenAPI Specification has been crucial in streamlining our API development process and crafting a seamless experience for developers.
Now, we’re taking a significant leap forward by adopting the OpenAPI v3 Specification. This update is a testimony to our commitment to fostering a high-quality API contract, and equipping our developer community with the most up-to-date tools and technologies. With that said, we are excited to announce the rollout of OpenAPI v3 support for the Meraki Dashboard API!
So, what’s it used for?
Use Cases:
- API Design and Development: Design APIs in a standard, language-agnostic manner to reduce misunderstandings and rework.
- API Documentation: Generate interactive, up-to-date API documentation that’s easy for developers to understand and use.
- Code Generation: Automatically generate client and server code in many languages, ensuring consistency between the API’s implementation and its documentation.
- API Testing and Validation: Create test cases and validate that the API’s implementation matches its design.
- API Management and Governance: Standardize APIs for easier management and governance within organizations.
New Use Cases Enabled by v3:
- Content Negotiation: Define different schema and examples for different media types, supporting multiple formats like JSON and XML.
- Linking and Callbacks: Describe complex, multi-step operations and outgoing ‘callback’ webhooks for asynchronous APIs and push notifications.
- Security Enhancements: Support complex OAuth 2.0 flows and OpenID Connect Discovery with flexible security definitions.
- Better Examples: Provide multiple examples for request bodies and responses to improve API documentation.
Key Improvements
The transition to OpenAPI v3 brings a lot of great stuff and new possibilities for developers. Here are some of the initial improvements we’re implementing out of the gate:
- Callback Support for Async Operations: The introduction of callbacks in v3 paves the way for better handling of complex, multi-step operations and outgoing ‘callback’ webhooks for asynchronous APIs and push notifications. This empowers developers to manage intricate, time-consuming tasks more efficiently, improving the overall performance. Stay tuned as we adopt some of these new features into the Dashboard API (hint: liveTools & actionBatches).
- Enhanced API Definitions: Providing API response schemas are critical for strongly typed languages to ensure the data and types are as expected. OASv3 has enhanced these capabilities by providing flexible options to define dynamic, additive and conditional properties which the Dashboard API is taking advantage of. Creating SDKs for your favorite language should be easier and more reliable than ever before.
- Improved Tooling and Validations: v3 introduces improved tools and validation mechanisms that streamline the API design, development, and testing processes. These enhancements ensure that the API’s implementation aligns with its design, leading to a robust and reliable API experience.
Get Started with OpenAPI v3
Starting with OpenAPI v3 support for the Meraki Dashboard API is straightforward. Visit our Meraki OpenAPI GitHub repository. There you can find OpenAPI specifications for the Meraki Dashboard API. Here are shortcuts to the v1 & v1-beta specifications in JSON format if you want to have a quick peak under the hood.
Alternatively, you can explore the Developer Hub at Meraki.io for comprehensive documentation, sample code, and integration guides to start taking advantage of the Meraki Dashboard API.
We look forward to hearing your OpenAPI v3 creations and stories with us – so please provide comments and feedback below!
