---
id: collect-261001-general-networking/general-networking/uilibs-uiprotect-1b5cf075-2
title: "with the CLI:"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "consumer", "cost", "license", "mit license"]
source: docs/RAG/collect-261001-general-networking/uilibs-uiprotect-1b5cf075.md
source_anchor: ""
source_lines: [128, 216]
sha256: 90f36febb0f149ba4f96cec9ae5accf89b2df338e00427418a21cbae2e48d758
---

# with the CLI:

non-OTHERProtectEventChannel (detection / sensor / alarm-hub /
access). Administrative events such asprovision ,factoryReset andfwUpdate are dropped. Callers that need the unfiltered stream
should usesubscribe_events_websocket .
- event.raw is a permanent escape hatch onto the underlying private-APIEvent model when the public contract does not expose the field you
need. In particular, smart-detect detected attributes (license-plate
text, face-match name) are not available over the public API today,
so consumers that need them must fall back to the private path viaevent.raw .
- EventChange.UPDATED may carry no public-visible delta — diffevent.raw if you need to know exactly what changed.
- protect.active_events(device_id=...) returns the in-flight set,
derived directly from the public bootstrap cache. Useful for restoring
binary-sensor state after a reload — it works before anysubscribe_events call as long asupdate_public() has primed the
cache.
- All runtime state is sourced from public_bootstrap : lifecycle/active
state frompublic_bootstrap.events , credential-event identity frompublic_bootstrap.ulp_users (UniFi Identity), andevent.device_mac from the bootstrap device stores. All are refreshed byupdate_public() — including automatically on websocket reconnect — and resolve with
eventual consistency: anidentity that resolves toUnknownIdentity(reason="ulp_user_not_cached") for a freshly-enrolled
ULP user, or adevice_mac ofNone for a device not yet in the
bootstrap, both fill in on the nextupdate_public() / reconnect resync.
subscribe_devices is the device-side analog of subscribe_events: it
delivers a typed ProtectDeviceChange for each ADDED / UPDATED /
REMOVED device over the Public Integration API. Together with
public_bootstrap (the device snapshot) and subscribe_events
(detection / sensor events), it gives a thin consumer the three
concern-separated primitives it needs without any model-type routing or
merge logic of its own.
Like subscribe_events, it delivers merged public models from the primed
cache, so subscribe_devices requires update_public() to have run first and
raises RuntimeError otherwise. To avoid encoding that ordering yourself, use
subscribe_devices_and_prime(), which connects the websocket and primes in the
correct order in one call so the prime/subscribe ordering is moot. With the
combined helper, updates that arrive while update_public() is priming are
buffered and replayed onto the fresh snapshot, so the already-connected
subscriber does not miss an update from the prime window; with the explicit
two-step form the subscriber is registered after priming and starts from the
refreshed cache.
import logging
from uiprotect import DeviceChange, ProtectApiClient, ProtectDeviceChange
_LOGGER = logging.getLogger(__name__)
protect = ProtectApiClient(..., api_key="...")
def on_device(change: ProtectDeviceChange) -> None:
    if change.change is DeviceChange.UPDATED and "state" in change.changed_fields:
        _LOGGER.info("%s -> %s", change.device_id, change.model.state)
# Order-independent: subscribes and primes in the correct order.
unsubscribe = await protect.subscribe_devices_and_prime(on_device)
# ...or the explicit two-step form, which must prime before subscribing:
# await protect.update_public()
# unsubscribe = protect.subscribe_devices(on_device)
unsubscribe()
Notes:
- Each change carries the merged Public* model inchange.model (None forREMOVED , where only an id /modelKey reference is
delivered).change.changed_fields is populated only forUPDATED .
- Single and bulk WS envelopes are expanded transparently to one change
per device, so consumers never see batched id arrays.
- Connection / state transitions surface as ordinaryUPDATED s withstate inchanged_fields — there is no separate side channel.
- The callback must not raise: an exception is caught and logged but otherwise swallowed.
- change.device_mac resolves with eventual consistency — a device not
yet in the bootstrap yieldsNone until the nextupdate_public() /
reconnect resync.
- This is device state only. It does not synthesize detection / motion
(use subscribe_events ) and there is no adoption concept folded in.
The library is moving from the legacy private API to Ubiquiti's official Public Integration API. The private API is considered legacy and is being phased out — new work targets the public API, implemented spec-conformantly: covering the features the spec exposes and staying as close to it as possible.
An explicit architectural goal is to keep the library shaped so the Home Assistant integration can stay thin — capabilities, device models, and events are surfaced here so the integration carries as little logic of its own as possible.
Please open an issue and agree on the approach before implementing anything — it avoids wasted effort on changes that don't fit the project's direction.
Important
This library does not accept new features built on the private API. uiprotect is migrating from the reverse-engineered private API to UniFi's official Public Integration API. If a capability is missing from the public API, the right path is to request it from Ubiquiti / wait for it to be exposed there — not to add it on the private path. Issues or PRs that introduce new private-API functionality will be closed.
Using AI? Fine — but you have to drive it. We use AI tooling ourselves, so an unreviewed AI-generated PR or issue doesn't save us anything; it just shifts the review and cleanup cost onto us. AI-assisted contributions are welcome only when you genuinely understand the architecture and the project's strategic direction, and the approach has been agreed in an issue first.
Where a contribution actually helps is the part AI can't supply — often because it
involves a device none of the maintainers happen to own. We have plenty of UniFi
hardware, just not every model, so testing and validation on a device we don't have,
sanitized payload captures from it, or first-hand knowledge of how it behaves in the
field are genuinely valuable — shaped to fit the architecture (see
AGENTS.md). Raw AI output that skips the prior discussion or ignores
these guidelines just creates review burden and will be closed.
The recommended way to develop is using the provided devcontainer with VS Code:
- Install VS Code and the Dev Containers extension
- Open the project in VS Code
- When prompted, click "Reopen in Container" (or use Command Palette: "Dev Containers: Reopen in Container")
- The devcontainer will automatically set up Python, Poetry, pre-commit hooks, and all dependencies
Alternatively, if you want to develop natively without devcontainer:
# Install dependencies (--all-extras installs the cli extra for CLI tests)
poetry install --with dev --all-extras
# Install pre-commit hooks
poetry run pre-commit install --install-hooks
# Run tests
poetry run pytest
# Run pre-commit checks manually
poetry run pre-commit run --all-files
This project was split off from pyunifiprotect because that project changed its license to one that would not be accepted in Home Assistant. This project is committed to keeping the MIT license.
- Bjarne Riis (@briis) for the original pyunifiprotect package
- Christopher Bailey (@AngellusMortis) for the maintaining the pyunifiprotect package
