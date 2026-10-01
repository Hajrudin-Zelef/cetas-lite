---
id: collect-261001-general-networking/general-networking/uilibs-uiprotect-1b5cf075-1
title: "with the CLI:"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["decode"]
source: docs/RAG/collect-261001-general-networking/uilibs-uiprotect-1b5cf075.md
source_anchor: ""
source_lines: [1, 127]
sha256: c6e68952074c9a3fa9b07d9d64c4f723d8dd962adf06a0fac6a5006efac72ca6
---

# with the CLI:

Documentation: https://uiprotect.readthedocs.io
Source Code: https://github.com/uilibs/uiprotect
Python API and CLI for UniFi Protect (Unofficial).
This module communicates with UniFi Protect surveillance software installed on a UniFi OS Console such as a Ubiquiti CloudKey+ (Cloud Key Gen2 Plus), a UniFi Network Video Recorder (UNVR or UNVR Pro), or a UniFi Dream Machine Pro, SE, or Pro Max.
uiprotect is increasingly built on Ubiquiti's official, documented Public Integration API. Where a capability is not yet available there, it falls back to the older private API, which is undocumented and can change as Ubiquiti evolves the software — so those parts may have gaps or shift between firmware releases.
The module is primarily written for the purpose of being used in Home Assistant core integration for UniFi Protect but might be used for other purposes also.
Full documentation for the project is available at uiprotect.readthedocs.io.
If you want to install uiprotect natively, the below are the requirements:
- UniFi Protect version 7.2+
  - The library is generally tested against the latest stable version.
- Python 3.11+
- POSIX compatible system
- PyAV (av) - included as a dependency
  - PyAV is used for audio streaming to camera speakers (talkback feature)
Alternatively you can use the provided Docker container, in which case the only requirement is Docker or another OCI compatible orchestrator (such as Kubernetes or podman).
Windows is not supported. If you need to use uiprotect on Windows, use Docker Desktop and the provided docker container or WSL.
uiprotect is available on PyPI:
pip install uiprotect
To use the command-line interface, install the cli extra (it pulls in typer):
pip install "uiprotect[cli]"pip install git+https://github.com/uilibs/uiprotect.git#egg=uiprotect
# with the CLI:
pip install "uiprotect[cli] @ git+https://github.com/uilibs/uiprotect.git"
A Docker container is also provided, so you do not need to install/manage Python as well. You can add the following to your .bashrc or similar.
function uiprotect() {
    docker run --rm -it \
      -e UFP_USERNAME=YOUR_USERNAME_HERE \
      -e UFP_PASSWORD=YOUR_PASSWORD_HERE \
      -e UFP_ADDRESS=YOUR_IP_ADDRESS \
      -e UFP_PORT=443 \
      -e UFP_SSL_VERIFY=false \
      -e TZ=America/New_York \
      -v $PWD:/data ghcr.io/uilibs/uiprotect:latest "$@"
}
Some notes about the Docker version since it is running inside a container:
- You can update at any time using the command docker pull ghcr.io/uilibs/uiprotect:latest
- Your local current working directory ($PWD ) will automatically be mounted to/data inside of the container. For commands that output files, this is the only path you can write to and have the file persist.
- The container supports linux/amd64 andlinux/arm64 natively. This means it will also work well on macOS or Windows using Docker Desktop.
- TZ should be the Olson timezone name for the timezone your UniFi Protect instance is in.
- For more details on TZ and other environment variables, check the command line docs
Warning
Ubiquiti SSO accounts are not supported and actively discouraged from being used. There is no option to use MFA. You are expected to use local access user. uiprotect is not designed to allow you to use your owner account to access the console or to be used over the public internet as both pose a security risk.
Note
uiprotect is increasingly built on Ubiquiti's official Public Integration API, which authenticates with a console-scoped API key instead of a username/password — no SSO, MFA, or owner account involved. New functionality targets this path first, and it is expected to become the primary — and eventually the only — supported authentication method. See Public-only mode below.
export UFP_USERNAME=YOUR_USERNAME_HERE
export UFP_PASSWORD=YOUR_PASSWORD_HERE
export UFP_ADDRESS=YOUR_IP_ADDRESS
export UFP_PORT=443
# set to true if you have a valid HTTPS certificate for your instance
export UFP_SSL_VERIFY=false
# Alternatively, use an API key for authentication (required for public API operations)
export UFP_API_KEY=YOUR_API_KEY_HERE
uiprotect --help
uiprotect nvr
Top-level commands:
- uiprotect shell - Start an interactive Python shell with the API client
- uiprotect create-api-key <name> - Create a new API key for authentication
- uiprotect get-meta-info - Get metadata information
- uiprotect generate-sample-data - Generate sample data for testing
- uiprotect profile-ws - Profile WebSocket performance
- uiprotect decode-ws-msg - Decode WebSocket messages
Device management commands:
- uiprotect nvr - NVR information and settings
- uiprotect events - Event management and export
- uiprotect cameras - Camera management
- uiprotect lights - Light device management
- uiprotect sensors - Sensor management
- uiprotect viewers - Viewer management
- uiprotect liveviews - Live view configuration
- uiprotect chimes - Chime management
For more details on any command, use uiprotect <command> --help.
UniFi Protect itself is 100% async, so as such this library is primarily designed to be used in an async context.
The main interface for the library is the uiprotect.ProtectApiClient:
from uiprotect import ProtectApiClient
# Initialize with username/password
protect = ProtectApiClient(host, port, username, password, verify_ssl=True)
# Or with API key (required for public API operations)
protect = ProtectApiClient(host, port, username, password, api_key=api_key, verify_ssl=True)
await protect.update() # this will initialize the protect .bootstrap and open a Websocket connection for updates
# get names of your cameras
for camera in protect.bootstrap.cameras.values():
    print(camera.name)
# subscribe to Websocket for updates to UFP
def callback(msg: WSSubscriptionMessage):
    # do stuff
unsub = protect.subscribe_websocket(callback)
# remove subscription
unsub()
You can also build a client that does no private login at all — just an API
key. Private-session entry points (update(), authenticate(),
get_bootstrap()) raise PublicOnlyModeError; drive everything through
update_public(), subscribe_events(), subscribe_devices(), the
get_*_public() / update_*_public() methods, and get_meta_info(). A
revoked key surfaces as NotAuthorized.
from uiprotect import ProtectApiClient
protect = ProtectApiClient.public_only(host, port, api_key=api_key, verify_ssl=True)
await protect.update_public()
# work with the public-API device snapshots
for siren in await protect.get_sirens_public():
    print(siren.name)
# resolve the console identity; the mac comes straight off the public NVR.
# For new code, prefer the public-API primary key (nvr.id) as the device
# identity rather than the mac.
console_mac = await protect.resolve_nvr_mac()
ProtectApiClient exposes two parallel websocket contracts. The raw
subscribe_events_websocket continues to deliver WSSubscriptionMessage
frames for advanced callers, and the typed subscribe_events API
delivers (ProtectEvent, EventChange) pairs intended for application
code. The typed path goes through the Public Integration API, so the
ProtectApiClient must be configured with an API key and
update_public() must have been called at least once before calling
subscribe_events. Use subscribe_events_and_prime() to subscribe and prime
in one order-independent call.
import logging
from uiprotect import EventChange, ProtectApiClient, ProtectEvent
_LOGGER = logging.getLogger(__name__)
protect = ProtectApiClient(..., api_key="...")
def on_event(event: ProtectEvent, change: EventChange) -> None:
    if change is EventChange.STARTED:
        _LOGGER.info("%s on %s: %s", event.type, event.device_id, event.identity)
    elif change is EventChange.ENDED:
        _LOGGER.info("%s ended after %s", event.type, event.end - event.start)
# Order-independent: subscribes and primes in the correct order.
unsubscribe = await protect.subscribe_events_and_prime(on_event)
# ...
unsubscribe()
Notes:
- subscribe_events delivers only events whoseEventType maps to a
