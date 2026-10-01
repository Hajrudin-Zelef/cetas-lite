---
id: collect-261001-meraki/meraki/dexterlabora-meraki-service-520ec68c
title: "dexterlabora-meraki-service-520ec68c"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "parameters", "sandbox"]
source: docs/RAG/collect-261001-meraki/dexterlabora-meraki-service-520ec68c.md
source_anchor: ""
source_lines: [1, 61]
sha256: 405b23d1115e972ad96291a0bdb91410afc80ebe119b9c71947f1280702c27da
---

# dexterlabora-meraki-service-520ec68c

A collection of methods to interact with the Meraki Dashboard API.
This is a convenient API wrapper built with Axios. The service saves time by handling common tasks when working with the API. It works for both backend and frontend code. Tested primarily with NodeJS w/ Express and VueJS.
-- In Active Development -- This service could have several breaking changes as it is being developed. You can clone this into your application if you want stability for now.
- Collection of the most common Dashboard API calls
- Handles URL redirects
- Handles Meraki error messages
- Custom scripts to traverse multiple API enpoints or enrich the response data
- The API rate limit is 5 calls per second (as of May, 2018)
- The Meraki API implements CORS, thus all API calls to Meraki must not come directly from the client browser. Instead, interact with the API from the server or create proxy server. You can then use this library on the client side, using your proxy address as the new endpoint. See the examples folder.
Requires Node >= v8
npm install https://github.com/dexterlabora/meraki-service.git
or (once published)
npm install meraki-service
An included test script will run an inventory of your organizations, networks, devices and group policies. Run this to confirm everything is working properly.
API_KEY=1234YourAPIKey node ./test/test
There are additional script samples in the file ./test/test-scripts. Open the file to uncomment out sample scripts, then run from the command.
API_KEY=1234YourAPIKey ORG_ID=123 node ./test/test-scripts
The meraki-service.js file contains a JavaScript class with all of the service methods. Explore this file to understand what each of the method names are and their required parameters. Use the Meraki Dashboard API docs to understand the required body parameters, etc.
// index.js
const Meraki = require('meraki-service');
const meraki = new Meraki('YourAPIKey','https://api.meraki.com/api/v0');
meraki.getOrganizations().then(res => {
    console.log('Organizations: ', res);
});
$ node index.js
$ Organizations:  [ { id: 549236, name: 'Meraki DevNet Sandbox' } ]... <node express> ...
app.use('/api', jsonParser, function (req, res){
    console.log('API request ', req.method, req.url, req.method != 'GET' ? req.body:'');
    var options = {
      url: req.url,
      method: req.method,
      body: JSON.stringify(req.body)
    };
    const apiKey = req.headers['x-cisco-meraki-api-key'] || configs.apiKey;
    const meraki = new Meraki(apiKey,configs.apiUrl);
    meraki.proxy(options).then((response) => {
      res.send(response);
      res.end();
    });
});$ node index.js
Server Running on:      http://localhost:5000
Meraki API Proxy:       http://localhost:5000/api
Meraki API Endpoint:    https://api.meraki.com/api/v0
Use the proxy address with desired API endpoint (uses server-side API key)
$ curl -X GET http://localhost:5000/api/organizations
[{"id":549236,"name":"Meraki DevNet Sandbox"}]
The proxy accepts an API key override. Just specify the new header in the reqest
$ curl -X GET \
>   http://localhost:5000/api/organizations \
>   -H 'X-Cisco-Meraki-API-Key: 2f301bccd61b6c6BoGuS3f76e5eb66ebd170f'
[{"id":549236,"name":"Meraki DevNet Sandbox"}]
- The /endpoints folder contains files for each of the endpoint groups.
- It is easy to duplicate any of the methods and modify them for new API endpoints.
- Send me a pull request if you want to contribute.
- improve redirect handling
- improve error handling
- complete JSDocs
  - Descriptions
  - Sample Code
- implement proper tests
Apache-v2.0
