---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/art-of-wifi-unifi-api-client-91bb094d-2
title: "art-of-wifi-unifi-api-client-91bb094d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "parameters"]
source: docs/RAG/collect-261001-unifi-ubiquiti/art-of-wifi-unifi-api-client-91bb094d.md
source_anchor: ""
source_lines: [118, 234]
sha256: 624bcf7c345b7e0aba6858c1704cf9beb23911f8ce9893d117de8aa75417ec04
---

# art-of-wifi-unifi-api-client-91bb094d

// ... make API calls through the proxy ...
$client->disable_site_manager_proxy();
// Client must be re-configured (login, set_api_key, etc.) before making direct calls
SSL verification:
In proxy mode, all requests go to api.ui.com which has a valid public CA certificate. SSL peer and host
verification is automatically enforced.
Performance note: Proxied requests add ~800ms of latency compared to direct access due to the cloud hop. Use direct access when available. The proxy is best suited for remote/headless deployments where a direct network path does not exist.
Rate limits:
The Site Manager API enforces rate limits (10,000 requests/minute for v1 stable). When exceeded,
the API returns HTTP 429 with a Retry-After header.
| Scenario | Method | 
|---|---|
| Controller is part of a UniFi Fabric | API key (required) | 
| UniFi OS console (UDM, UDR, UCG, etc.) | API key (recommended) or username/password | 
| UniFi OS Server | API key (recommended) or username/password | 
| Self-hosted Network Application (non-UniFi OS Server) | Username/password (API keys are not supported) | 
| Remote console via UI.com (no direct network path) | Site Manager proxy | 
Important: When your controller is a member of a UniFi Fabric, username/password authentication is not available. You must use API key authentication. UniFi Fabric uses centralized identity management which does not support local login sessions through the API.
If you have existing code using username/password authentication and want to switch to API keys, the changes are minimal:
Before (username/password):
$unifi_connection = new UniFi_API\Client($user, $password, $url, $site_id);
$unifi_connection->login();
$results = $unifi_connection->list_alarms();
$unifi_connection->logout();
After (API key):
$unifi_connection = new UniFi_API\Client('', '', $url, $site_id);
$unifi_connection->set_api_key($api_key);
$results = $unifi_connection->list_alarms();
What changes:
- Pass empty strings for the $user and$password constructor parameters
- Call set_api_key() instead oflogin()
- Remove logout() calls (or leave them — they will be silently ignored)
- In your exception handling, LoginFailedException andLoginRequiredException will no longer occur;
you can keep or remove those catch blocks as you prefer
- 
In the above example, $site_id is the short site "name" (usually 8 characters long) that is visible in the URL when
managing the site in the UniFi Network Controller. For example, with this URL:https://<controller IP address or FQDN>:8443/manage/site/jl3z2shm/dashboardjl3z2shm is the short site "name" and the value to assign to $site_id.
- 
The 6th optional parameter that is passed to the constructor in the above example ( true ), enables validation of
the controller's SSL certificate, which is otherwise disabled by default. It is highly recommended to enable
this feature in production environments where you have a valid SSL cert installed on the UniFi Controller that is
associated with the FQDN in thecontroller_url parameter. This option was added with API client version 1.1.16.
- 
Using an administrator account ( $controller_user in the above example) with read-only permissions can limit
visibility on certain collection/object properties. See this
issue and this
issue for an example where the WPA2 password isn't
visible for read-only administrator accounts.
More code examples are available in the examples/ directory.
The API Client class throws Exceptions for various error conditions instead of using PHP's trigger_error()
function. This allows for more granular error handling in your application code.
You can also choose to catch the UniFi_API\Exceptions\UnifiApiException Exception to catch all Exceptions that
might be thrown by the API Client class.
Here is an example of how to catch each of the Exceptions individually:
<?php
/**
 * PHP API usage example with Exception handling
 */
use UniFi_API\Exceptions\ConsoleOfflineException;
use UniFi_API\Exceptions\CurlExtensionNotLoadedException;
use UniFi_API\Exceptions\CurlGeneralErrorException;
use UniFi_API\Exceptions\CurlTimeoutException;
use UniFi_API\Exceptions\InvalidBaseUrlException;
use UniFi_API\Exceptions\InvalidSiteNameException;
use UniFi_API\Exceptions\JsonDecodeException;
use UniFi_API\Exceptions\LoginFailedException;
use UniFi_API\Exceptions\LoginRequiredException;
/**
 * load the class using the composer autoloader
 */
require_once 'vendor/autoload.php';
/**
 * include the config file (place your credentials etc. there if not already present)
 */
require_once 'config.php';
/**
 * initialize the UniFi API connection class, log in to the controller and request the alarms collection
 * (this example assumes you have already assigned the correct values to the variables in config.php)
 */
try {
    $unifi_connection = new UniFi_API\Client($controller_user, $controller_password, $controller_url, $site_id, $controller_version, true);
    $login            = $unifi_connection->login();
    $results          = $unifi_connection->list_alarms(); // returns a PHP array containing alarm objects
} catch (CurlExtensionNotLoadedException $e) {
    echo 'CurlExtensionNotLoadedException: ' . $e->getMessage(). PHP_EOL;
} catch (InvalidBaseUrlException $e) {
    echo 'InvalidBaseUrlException: ' . $e->getMessage(). PHP_EOL;
} catch (InvalidSiteNameException $e) {
    echo 'InvalidSiteNameException: ' . $e->getMessage(). PHP_EOL;
} catch (JsonDecodeException $e) {
    echo 'JsonDecodeException: ' . $e->getMessage(). PHP_EOL;
} catch (LoginRequiredException $e) {
    echo 'LoginRequiredException: ' . $e->getMessage(). PHP_EOL;
} catch (ConsoleOfflineException $e) {
    echo 'ConsoleOfflineException: ' . $e->getMessage(). PHP_EOL;
} catch (CurlGeneralErrorException $e) {
    echo 'CurlGeneralErrorException: ' . $e->getMessage(). PHP_EOL;
} catch (CurlTimeoutException $e) {
    echo 'CurlTimeoutException: ' . $e->getMessage(). PHP_EOL;
} catch (LoginFailedException $e) {
    echo 'LoginFailedException: ' . $e->getMessage(). PHP_EOL;
} catch (Exception $e) {
    /** catch any other Exceptions that might be thrown */
    echo 'General Exception: ' . $e->getMessage(). PHP_EOL;
}
Although the PHP DocBlocks for most public methods/functions contain @throws Exception, it is recommended to catch
specific Exceptions that can be thrown by the API Client class to provide more detailed error messages to your
application code.
In most cases, the class will let Exceptions bubble up to the calling code, but in some cases it will catch them and throw a new Exception with a more specific message.
The list_alarms.php example in the examples/ directory is a good starting point to see how you can implement
Exception handling.
The API Client class currently supports a large and growing number of functions/methods to access the UniFi Controller API. Please refer to the comments/PHP DocBlocks in the source code for more details on each of the functions/methods, their purpose, and their respective parameters.
If you are using an advanced IDE such as PHPStorm or VS Code, you can use its code completion and other features to explore the available functions/methods thanks to the extensive PHP DocBlocks throughout the code.
For a quick overview of the available functions/methods, you can also check the API Reference here:
API Reference
There is still work to be done to add functionality and further improve the usability of this API Client class, so all suggestions/comments are welcome. Please use the GitHub Issues section or the Ubiquiti Community forums (https://community.ui.com/questions/PHP-client-class-to-access-the-UniFi-controller-API-updates-and-discussion-part-2/a793904e-6023-4a7f-bcae-340db2a03fc1) to share your suggestions and questions.
When encountering issues with the UniFi API using other libraries, cURL or Postman, please do not open an Issue. Such issues will be closed immediately. Please use the Discussions section instead.
