---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/thib3113-unifi-client-e41c6526
title: "thib3113-unifi-client-e41c6526"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/thib3113-unifi-client-e41c6526.md
source_anchor: ""
source_lines: [1, 132]
sha256: 71f46e55c44c241980d5ad21a2980e461627be8517bddcc81a1d5047086b8003
---

# thib3113-unifi-client-e41c6526

This library is a nodejs client to talk with unifi rest API .
install it
yarn add unifi-client
or via npm
npm i unifi-client
and import it in your code :
import Controller from 'unifi-client'
or with require :
const { Controller } = require('unifi-client')
- Installed UniFi-Controller version v6 or more, UnifiOs or not ( cloud keys / UDM or just unifi controller software ) ( not tested below v6 )
- A network connectivy between unifi-client and the controller ( env http_proxy and https_proxy can be handled automaticaly, but not tested )
- An account, cloud accounts, local accounts and 2FA are available
- Node.js version >= LTS (14)
Examples are located in examples folder
import Controller from 'unifi-client'
// works with local account, check examples for 2FA
const controller = new Controller({
  username: 'ubnt',
  password: 'ubnt',
  url: 'https://unifi',
  strictSSL: false
});
//login to the controller
await controller.login()
//retrieve sites
const sites = controller.getSites()
//work on one site
const site = sites[0];
//example request to get firewall rules
const rules = await site.firewall.getRules();
To work on multiples sites on the same time
const controller = new Controller({...});
const sites = controller.getSites();
//get firewallRules on two sites
const rules = await site[0].firewall.getRules();
const rules2 = await site[1].firewall.getRules();
How to use the HTTP instances
//use the controller instance directly . Authentication, url construction and other is already managed for you
const self = await controller.getInstance().get('/api/self');
//for a custom site :
const topology = await site.getInstance().get('/topology');
Get firewall rules, it return an array of FWRule
const rules = await site.firewall.getRules();
The instances returned by getInstance are basicaly some axios instances . With some additions :
//url params
// use them to manage a rest url for example
instance.get('/url/:uuid', {
    urlParams: {
        uuid: 'my-uuid'
    }
})
// will call the url /url/my-uuid
// apiVersion :
// some api call are on the v2 api, so you can set the api version when calling
console.log(site.name); // default
site.getInstance().get('/super/endpoint', {
    apiVersion: 2
});
// will call the url /v2/api/site/default/super/endpoint
The urls of unifi are like :
- /proxy/network/v2/api/site/default/notifications for unifiOs
or
- /v2/api/site/default/notifications for non unifiOs
the url is constructed like :
/[v<ApiVersion>/]api/site/<sitename>[/<urlToCall>]
- ApiVersion : the version of the API (for version > 1)
- sitename : the name of the site (already included in site instance)
To illustrate, to call the url I see on my browser network tabs :
https://192.168.1.1/proxy/network/v2/api/site/default/notifications
I just need to :
// configure my controller
const controller = new Controller({
    username: 'ubnt',
    password: 'ubnt',
    url: 'https://192.168.1.1',
    strictSSL: false
});
//login to the controller
await controller.login()
//retrieve sites
const sites = controller.getSites()
// select the site "default"
const site = sites.find((site) => site.name === 'default');
//call my url
const notifications = site.getInstance().get('/default/notifications', {
    apiVersion: 2
});
//do something with notifications
More examples in the folder examples
This library supports websockets, you can listen on them with :
// initWebSockets for controller :
await controller.initWebSockets();
// listen for controller websockets ( only for unifiOS )
controller.on() / controller.ws.on()
// listen for super site websockets
controller.superWS.on();
// initWebSockets for site :
await site.initWebSockets();
// listen for this site websockets
site.on() / site.ws.on()
// doesn't known the name of the event ? you can listen on joker :
controller.on('*', (eventName, ...args) => {
    console.log(eventName, ...args);
});
// want to listen on all websockets of a controller on the same listener
// all websockets registered :
// - controller websockets if unifiOS
// - super site WS
// - all sites where ws are init ( closed or not )
controller.globalWS.on('*', (eventName, ...args) => {
    console.log(eventName, ...args);
});
the support of websocket is experimental, and with a really bad coverage . Doesn't hesitate to open a PR, to add more websockets types
We add some optionnal dependencies to improve performances, more informations read here
they can be skipped by using
npm install --no-optional
All the technical documentation is available here
check technical documentation for available methods
This library is auto-tested on :
- 
UDM-pro : latest (7.3.83 )
- 
Unifi controller : latest ,7.3.83 ,7.2.95 ,7.1.68 ,7.0.25 ,6.5.55 ,6.4.54 ,6.2.26 ,6.0.45 (only some functions are tested, in an empty environment)
with the last node LTS
This nodejs package/class uses functionality/Know-How gathered from different third-party projects:
To debug, you can use the debug library, just set the DEBUG environnement variable (check this) .
- unifi-client : base debug namespace
- unifi-client:<className> : will show debug for this class (eg : unifi-client:Controller will log controller logs )
- unifi-client:axios : will logs axios request (can contain secrets informations)
- unifi-client:axios:verbose : will logs more things on axios, like send and receive payload (can contain secrets informations)
- unifi-client:axios:curl : will log curl cmd corresponding to the axios request (can contain secrets informations)
ubntwiki : https://ubntwiki.com/products/software/unifi-controller/api
