---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-48
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-48.md
source_anchor: ""
source_lines: [1, 63]
sha256: 789fed7cb8979e593ddef81591df42d863bc599495b4a441578e68cc05bdb0dd
---

# Summary

Fetch is one of the console tools in MikroTik RouterOS. It is used to copy files to/from a network device via HTTP, HTTPS, FTP or SFTP. It can also be used to send POST/GET requests and send any kind of data to a remote server. In HTTPS mode by default, no certificate checks are made, setting *check-certificate* to *yes* enables trust chain validation from the local certificate store (can be used only in HTTPS mode).

# Properties

| Property | Description | 
|---|---|
| **address** (*string* ; Default: ) | IP address of the device to copy file from. Also at the end of the address you can specify "@vrf_name" in order to run fetch on particular VRF. You can skip specifying address and specify only VRF on this parameter, if you use URL parameter. | 
| **as-value** (*set \| not-set* ; Default:**not-set** ) | Store the output in a variable, should be used with the output property. | 
| **ascii** (*yes \| no* ; Default:**no** ) | Can be used with FTP and TFTP | 
| **certificate**  (*string* ; Default: ) | Certificate that should be used for host verification. Can be used only in HTTPS mode. | 
| **check-certificate** (*yes \| yes-without-crl \| no* ; Default:**no** ) | Enables trust chain validation from local certificate store. *yes-without-crl, validates a certificate, not performing CRL check (certificate revocation list).*   Can be used only in HTTPS mode. | 
| **dst-path** (*string* ; Default: ) | Destination path. Can be used to download file directly into an external disk, for example. | 
| **duration** (*time;* Default: ) | Time how long fetch should run. | 
| **host** (*string* ; Default: ) | A domain name or virtual domain name (if used on a website, from which you want to copy information). For example, address=wiki.mikrotik.com host=forum.mikrotik.com In this example the resolved ip address is the same (66.228.113.27), but hosts are different. | 
| **http-auth-scheme** (*basic\|digest* ; Default:**basic** ) | HTTP authentication scheme | 
| **http-method** (*delete\|get\|head\|post\|put\|patch* ; Default:**get** ) | HTTP method to use | 
| **http-data** (*string* ; Default: ) | The data, that is going to be sent. Data limit is 64Kb. | 
| **http-header-field** (*string* ; Default:***empty*** ) | List of all header fields and their values, in the form of `http-header-field="h1:fff,h2:yyy"` or`http-header-field="h:fff\\,yyy"` (within a single header multiple values need to be "escaped" using two backlashes). | 
| **http-content-encoding** (*deflate\|gzip* ; Default: ***empty*** ) | Encodes the payload using **gzip** or**deflate** compression and adds a corresponding Content-Encoding header. Usable for HTTP POST and PUT only. | 
| **http-max-redirect-count** (*integer* ; Default: 2) | Allows to follow redirects the specified number of times. | 
| **http-percent-encoding**  (*yes \| no* ; Default:**no** ) | Enables HTTP percent encoding | 
| **http-version** (*http1_1* \|*http2* ; Default:**http1_1** ) | Specifies which HTTP version to use HTTP2 supported only on ARM64 and x86/CHR devices | 
| **idle-timeout** (*time;* Default: 10s) | Idle timeout since last read/write action. | 
| **keep-result** (*yes \| no* ; Default:**yes** ) | If yes, creates an input file. | 
| **mode** (*ftp\|http\|https\|sftp\|tftp* ; Default:**http** ) | Choose the protocol of connection - http, https , ftp, sftp or tftp. Mode option is deprecated. To specify a protocol that you wish to use, we advise using "url" parameter instead (for example, like this "url=sftp://your_IP_address"). | 
| **output** (*none\|file\|user\|user-with-headers* ; Default:**file** ) | Sets where to store the downloaded data.  | 
| **password** (*string* ; Default:**anonymous** ) | Password, which is needed for authentication to the remote device. | 
| **port** (*integer* ; Default: ) | Connection port. | 
| **src-address** (ip address; Default: ) | Source address that is used to establish connection. Can be used only HTTP/S and SFTP modes. | 
| **src-path** (*string* ; Default: ) | Title of the remote file you need to copy. | 
| **upload** (*yes \| no* ; Default:**no** ) | Only (S)FTP modes support upload. If enabled then fetch will be used to upload files to a remote server. Requires *src-path* and*dst-path* parameters to be set. | 
| **url** (*string* ; Default: ) | URL pointing to file. Can be used instead of **address** and**src-path** parameters. | 
| **user** (*string* ; Default:**anonymous** ) | Username, which is needed for authentication to the remote device. | 

# Configuration Examples

The following example shows how to copy the file with filename "conf.rsc" from a device with ip address 192.168.88.2 by FTP protocol and save it as file with filename "123.rsc". User and password are needed to login into the device.

Example to upload file to another router:

Another file download example that demonstrates the usage of url property.

It is also possible to transfer files over some specific VRF. You can specify VRF at the end of the URL address part (url="http://192.168.88.2@vrf1/..." - will not work for SFTP), address property combined with address and separated with "@" (address=192.168.88.2@vrf1) or simply as address parameter with "@" symbol before the VRF name as in the example (uses address from URL and combines it with VRF name from address parameter):

## Sending information to a remote host

It is possible to use an HTTP POST request to send information to a remote server, that is prepared to accept it. In the following example, we send geographic coordinates to a PHP page:

In this example, the data is uploaded as a file. Important note, since variable data comes from a file, a file can only be in size up to 4KB. This is a limitation of RouterOS variables.

## Return value to a variable

It is possible to save the result of the fetch command to a variable.

### Example 1

It's possible to trigger a certain action based on the result that an HTTP page returns. You can find a very simple example below that disables **ether2** whenever a PHP page returns "0": 

### Example 2

In case fetch fails, it is possible to access error code ("code") and returned HTTP headers ("http-headers").
