---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/art-of-wifi-unifi-api-client-91bb094d-1
title: "art-of-wifi-unifi-api-client-91bb094d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/art-of-wifi-unifi-api-client-91bb094d.md
source_anchor: ""
source_lines: [1, 117]
sha256: d992ef64dcd0c392d251e7c197fbee0e4b5e3666716097ab9661af8479e486f9
---

# art-of-wifi-unifi-api-client-91bb094d

A PHP class that provides access to Ubiquiti's UniFi Network Application API.
This class is used by our API Browser tool, which can be found here.
The package can be installed manually or by using composer/packagist for easy inclusion in your projects. See the installation instructions below for more details.
- Easy to use: clear docs, comprehensive method coverage, and helpful examples.
- Broad coverage: exposes many UniFi endpoints not (yet) available in the official APIs.
- Composer-friendly: installable via Composer and works with modern PHP projects.
- Lightweight and dependency-free: no external libraries required; uses cURL.
- Secure: communicates over TLS and supports optional SSL certificate validation.
- Flexible and extensible: includes custom_api_request() for calling any API endpoint.
- Robust error handling: throws named Exceptions for precise try/catch handling.
- Actively maintained: regular updates and compatibility with recent UniFi versions.
| Software | Versions | 
|---|---|
| UniFi Network Application/controller | 5.x, 6.x, 7.x, 8.x, 9.x, 10.x (10.5.62 is confirmed) | 
| UniFi OS | 3.x, 4.x, 5.x (5.1.26 is confirmed) | 
- a server or desktop with:
  - PHP 7.4.0 or higher (use version 1.1.83 for PHP 7.3.x and lower)
  - PHP cURL (php-curl ) module enabled
  - PHP JSON (php-json ) module enabled (always available on PHP 8.0 and later, but it can be
disabled at compile time on PHP 7.4)
  - direct network connectivity between this server/desktop and the host and port where the UniFi Network Application is running (usually TCP port 8443, port 11443 for UniFi OS Server, or port 443 for UniFi OS consoles)
- authentication — you need one of the following:
  - an admin account with local access permissions as explained here: https://artofwifi.net/blog/use-local-admin-account-unifi-api-captive-portal Do not use UniFi Cloud accounts and do not enable MFA/2FA for these accounts. (see Option 2: Username/Password below)
  - or an API key generated in your UniFi OS console or UniFi OS Server (see Option 1: API Key below)
  - or a Site Manager API key for routing requests through the Ubiquiti cloud proxy (see Option 3: Site Manager Proxy below)
Starting from version 1.1.47, this API client also supports UniFi OS-based controllers. These applications/devices/services have been verified to work:
- UniFi OS Server, announcement here
- UniFi Dream Router (UDR)
- UniFi Dream Machine (UDM)
- UniFi Dream Machine Pro (UDM PRO)
- UniFi Cloud Key Gen2 (UCK G2), firmware version 2.0.24 or higher
- UniFi Cloud Key Gen2 Plus (UCK G2 Plus), firmware version 2.0.24 or higher
- UniFi Express (UX)
- UniFi Dream Wall (UDW)
- UniFi Cloud Gateway Ultra (UCG-Ultra)
- UniFi CloudKey Enterprise (CK-Enterprise)
- UniFi Enterprise Fortress Gateway (EFG)
- Official UniFi Hosting, details here
- HostiFi UniFi Cloud Hosting, details here
The API Client automatically detects UniFi OS consoles/servers and adjusts URLs and several functions/methods accordingly.
UniFi OS-based consoles require you to connect using port 443, whereas port 8443 is used for the self-hosted/software-based controllers. When connecting to UniFi OS Server, you are required to use port 11443.
When connecting to a UniFi OS-based gateway through the WAN interface, you need to create a specific firewall rule to
allow this. See this blog post on the Art of WiFi website for detailed instructions when using the "Classic"
firewall:
https://artofwifi.net/blog/how-to-access-the-unifi-controller-by-wan-ip-or-hostname-on-a-udm-pro
See this blog post when using the Zone-Based firewall (ZBF):
https://artofwifi.net/blog/how-to-access-the-unifi-controller-by-wan-ip-or-hostname-on-a-udm-pro-using-zbf
When upgrading from a version before 2.0.0, please:
- change your code to use the new Exceptions that are thrown by the API Client class
- test the client with your code for any breaking changes
- make sure you are using Composer to install the API Client because the code is no longer held within a single file
- see the note here regarding the single file version (1.x.x) of the API client
The preferred installation method is through Composer. Follow these installation instructions if you don't have Composer installed already.
Once Composer is installed, execute this command from the shell in your project directory:
composer require art-of-wifi/unifi-api-client
Or manually add the package to your composer.json file:
{
    "require": {
        "art-of-wifi/unifi-api-client": "^2.0"
    }
}
Finally, be sure to include the composer autoloader in your code if your framework doesn't already do this for you:
/**
 * load the class using the composer autoloader
 */
require_once 'vendor/autoload.php';
The API client supports three authentication methods. Choose the one that fits your setup.
API key authentication is stateless — no login/logout flow is needed. This is the simplest way to get started with UniFi OS-based controllers.
require_once 'vendor/autoload.php';
$unifi_connection = new UniFi_API\Client('', '', 'https://unifi:443', 'default');
$unifi_connection->set_api_key('your-api-key-here');
$results = $unifi_connection->list_alarms(); // no login() needed
How to generate an API key:
- Open your UniFi OS console in a browser
- Navigate to Integrations in the sidebar menu
- Click Create New API Key
- Copy the generated key — it will not be shown again
The API key inherits the permissions from the admin user that created it.
Key points:
- API keys are only available on UniFi OS-based consoles (UDM, UDR, UCG, UX, UDW, UCG-Ultra, UniFi OS Server, etc.)
- No login() orlogout() calls are needed (calling them is harmless and will be ignored)
- The client automatically configures itself for UniFi OS when an API key is set
The traditional authentication method that works with both self-hosted controllers and UniFi OS consoles.
require_once 'vendor/autoload.php';
$unifi_connection = new UniFi_API\Client(
    $controller_user,
    $controller_password,
    $controller_url,
    $site_id
);
$login   = $unifi_connection->login();
$results = $unifi_connection->list_alarms();
Requirements for username/password authentication:
- You must use a local admin account with local access permissions as explained here: https://artofwifi.net/blog/use-local-admin-account-unifi-api-captive-portal
- Do not use UniFi Cloud accounts
- Do not enable MFA/2FA on accounts used with this client
If your console is managed through unifi.ui.com and you don't have a direct network path to it, you can route API requests through the Ubiquiti Site Manager cloud proxy. This uses the Site Manager API connector to reach the console via UI.com's cloud infrastructure.
require_once 'vendor/autoload.php';
$client = UniFi_API\Client::connect_via_site_manager(
    '245A4CA234150000000005F23204000000000638FE970000000061156371:48913759', // console ID
    'your-site-manager-api-key',                                              // Site Manager API key
    'default'                                                                  // site (optional)
);
// No login() needed — proxy mode is stateless
$stats = $client->stat_daily_site();
Finding the console ID:
The console ID (host ID) is visible in the URL when managing a console via unifi.ui.com:
https://unifi.ui.com/consoles/{console_id}/network/default/dashboard
Site Manager API key: Generate a Site Manager API key at https://unifi.ui.com under your account settings. This is not the same as a local controller API key generated in the UniFi OS console.
Requirements:
- Console firmware version must be >= 5.0.3
- Console must be online and connected to UI.com
- For non-organization API keys: limited to the API key owner's consoles only
- For organization API keys: can access any console within the organization
You can also enable/disable proxy mode on an existing client instance:
$client = new UniFi_API\Client('', '', 'https://127.0.0.1', 'default');
$client->enable_site_manager_proxy($console_id, $site_manager_api_key);
