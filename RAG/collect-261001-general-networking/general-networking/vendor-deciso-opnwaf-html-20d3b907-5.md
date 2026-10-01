---
id: collect-261001-general-networking/general-networking/vendor-deciso-opnwaf-html-20d3b907-5
title: "vendor-deciso-opnwaf-html-20d3b907"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "parameters"]
source: docs/RAG/collect-261001-general-networking/vendor-deciso-opnwaf-html-20d3b907.md
source_anchor: ""
source_lines: [322, 396]
sha256: a4fb9f1de80c8fc7511f11095ade4faa8bb8c58e1df2d2d6eea5a4e5a2996f68
---

# vendor-deciso-opnwaf-html-20d3b907

| Port | Port number to bind to, you can use Port forwarding to redirect traffic from standard ports to non standard ones when needed | 
| Certificate / Enable ACME | Either use an ACME certificate or define one yourself, this one should be trusted by the browser connecting to this host | 
| CA for client auth | select the Authority created earlier | 
Followed by a location, which maybe as simple as binding path / to a local machine without certificate at http://10.0.0.1.
Tip
You can use revocation lists to pull back access rights for selected clients, just make sure to restart the service in order to make the changes effective.
After this step, clients should not be able to access the virtual host, next you can create a certificate for the client and import it in the trust store. Usually browsers automatically pick these up when allowed by the client.
Protect a location with OpenID Connect
In the above virtual host and location configuration are a couple of parameters related to OpenID Connect. The advantage of using these is that you can prevent unauthenticated and unauthorized access to services using an identity provider.
First, add an identity provider for service OPNWAF in .
For more information refer to the OpenID Connect manual.
Next, add it to a virtual server in :
| Option | Description | 
|---|---|
| OpenID Connect |  | 
| OIDC Provider | Choose the identity provider created in | 
| OIDC Redirect URI | Leave default, this will create a URI that must be set with your identity provider. If the virtual server is example.com it will become https://example.com/oidc/callback if not specified otherwise. This location will be automatically removed from proxying. If you cannot use the default, choose an URI that does not collide with any path of your backend application. | 
As final step, ensure the following is set in each ProxyPass location of this virtual server:
| Option | Description | 
|---|---|
| OpenID Connect |  | 
| OIDC Auth Required | Select to enforce OIDC authentication with the below claim. | 
| OIDC Claims | Leave on default to allow any authenticated user in the OIDC scope access to the location. | 
After applying, the location will need authentication (user must log in).
OpenID Connect claims
A claim is a piece of information that can be used to identify a user. This means you can create a stricter policy which user has access to the location, not only enforcing authentication but also authorization.
As example, we only want to grant access to a location for all users with the first name John.
First, we add a claim in :
| Option | Description | 
|---|---|
| OpenID Connect |  | 
| Claim type | Most claim types are standardized via the OIDC spec. Some provider specific options are also offered (group). For our example case we choose name . | 
| Claim value | John | 
Next, we add the claim to an OpenID Connect enabled location in :
| Option | Description | 
|---|---|
| OpenID Connect |  | 
| OIDC Auth Required | Select to enforce OIDC authentication with the below claim. | 
| OIDC Claims | name John | 
After applying, the location will need authentication (user must log in) and authorization (user must be John).
Note
Multiple claims can be selected, they will be combined via or operator.
Tip
Authorizing unique users can be done with the preferred_username claim, which is the name a user authenticates with.
Some identity providers can send groups (non-standard) in their OIDC scope which simplifies authorization when you have a large amount of users.
Error Documents
By default, generic Apache documents will be served for HTTP response status codes. The most common client error responses can be styled OPNsense themed, or be branded with your own style.
To download the default error document templates, go to .
Select the Download command in the Default row. Afterwards you can unzip the archive, and change the individual error documents.
When you are done, select + to open the upload dialogue:
| Option | Description | 
|---|---|
| Name | Name for this template, e.g. MyErrorDocuments | 
| Uri | Uri used to serve error pages, when unspecified, /__waf_errors__/ will be used. Best to use the offered default. | 
| Content | Select the zip archive with the altered error documents. | 
After saving, the error documents can be added in :
| Option | Description | 
|---|---|
| Error Document | MyErrorDocuments will use your new template.Default will use the OPNsense styled template.None will use the unaltered default Apache documents. | 
To optionally overlay any error with only the template provided ones, you can set the following in a location:
| Option | Description | 
|---|---|
| Overlay error pages | Overlay common error pages with the ones specified in the virtual server. This means that all HTTP response status codes received from the Remote destinations will be stripped, and only matching HTTP response codes in the current selected error document template will be served. | 
Tip
When using OpenID Connect, it is a good idea to either use the Default or custom error documents, to ensure the Unauthorized
error pages have a more cohesive and user friendly style.
Server Status
The server exposes metrics that can be revealed in .
You can find tabs with general information, workers, current requests and ACME certificates.
If the server behaves unexpected, checking the current requests and server load can reveal potential issues:
- Detect backend connectivity issues Requests that remain open for unusually long periods or repeatedly fail may indicate that an upstream service is slow or unreachable. Inspecting the active request list can help determine whether the issue originates from the reverse proxy or the backend application.
- Validate configuration changes in real time After modifying the configuration, the status page can be used to verify that requests are handled as expected. Observing request paths and worker activity allows administrators to confirm that traffic reaches the intended backend services.
- Find malicious clients Malicious clients could DoS the hosted websites if the CPU load is very high and the number of concurrent requests is unexpected.
Tip
Server statistics are API enabled, meaning you could export them to an external monitoring system. Since they are not stored on disk, persistence requires external polling.
