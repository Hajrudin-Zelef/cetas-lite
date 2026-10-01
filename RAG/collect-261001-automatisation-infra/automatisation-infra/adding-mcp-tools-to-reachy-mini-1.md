---
id: collect-261001-automatisation-infra/automatisation-infra/adding-mcp-tools-to-reachy-mini-1
title: "profiles/default/tools.txt"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "mcp", "robotics", "voice"]
source: docs/RAG/collect-261001-automatisation-infra/adding-mcp-tools-to-reachy-mini.md
source_anchor: ""
source_lines: [1, 148]
sha256: 40aab8f8420599cd7a9967df12bbdba593dc1263aed9cea5dde1cee17f8cf43b
---

# profiles/default/tools.txt

#### Reachy Mini Search Tool

MCP Space for Reachy Mini web search.

Adding a tool takes one command:

```
reachy-mini-conversation-app tool-spaces add pollen-robotics/reachy-mini-weather-tool
```
Then start the app as usual:

```
reachy-mini-conversation-app
```
Now you can just ask:

```
What's the weather in Paris today?
```
Below, we look at what a tool is, how profiles control what the robot can use, and the current limits of the remote path.

When you talk to the robot, what you get back isn't only a voice, it's a system that reacts to the conversation: the robot can move and respond non-verbally, when it's applicable. The part we want to focus on here is the tools that make that possible. A tool is something the model can do during a conversation: play an emotion, move the head, look through the camera. Each tool has a name and a short description. The model reads those, decides when one is useful, calls it, and uses what comes back.

Today every tool is local and ships inside the app, and most of them are about the robot's body:

| Tool | What it does | 
|---|---|
| `move_head` | Queue a head pose change | 
| `dance` /`stop_dance` | Play or clear a dance from the dances library | 
| `play_emotion` /`stop_emotion` | Play or clear a recorded emotion clip | 
| `head_tracking` | Toggle head-tracking offsets | 
| `camera` | Capture a frame and analyze it | 
| `idle_do_nothing` | Explicitly stay idle on an idle turn | 

A tool in the code isn't usable until it's enabled in **a profile**, a folder with two files that matter here: `instructions.txt` (the prompt) and `tools.txt` (the tools that are turned on).

The `default` profile enables the full set:

```
# profiles/default/tools.txt
dance
stop_dance
play_emotion
stop_emotion
camera
idle_do_nothing
head_tracking
move_head
```
If a name isn't in `tools.txt`, the model can't call it.

You can also write your own tool: add a Python file to the profile (or `external_tools/`), give it a name and description, and list that name in `tools.txt`.

Today there are built-in tools and custom local tools, and `tools.txt` decides which are active. This works well for the robot's body and keeps the trusted core small.

The constraint here is that every tool has to be local Python. For `move_head` or `play_emotion` that's right: they talk to the hardware and belong in the app but a lot of useful things have nothing to do with the body, like web search, weather, or lookups. For those, keeping everything local is mostly friction:

Remote tools add a third kind, alongside the built-in and custom local tools you already have, for capabilities that are easier to publish, share, and update on their own:

`external_tools/`
It's a good fit for stateless capabilities like search, weather, and lookups: anything you want to iterate on without touching the app itself. And because anyone can publish a compatible Space, it's easy to share tools and build on each other's work.
We started with two canary tools, small test tools to exercise the new flow:

They're enough to exercise the whole feature: install from the Hub, discover the remote tools, enable them per profile, and let the realtime backend call them exactly like built-in tools.

To use both at once, add each Space and their tools stack in the same profile:

```
reachy-mini-conversation-app tool-spaces add pollen-robotics/reachy-mini-search-tool
reachy-mini-conversation-app tool-spaces add pollen-robotics/reachy-mini-weather-tool
```
Now the robot can search the web and check the weather in the same conversation, which is exactly what the `canary_web_search_weather` profile below does.

```
# install + enable in active profile
reachy-mini-conversation-app tool-spaces add <owner/space-name>
 
# enable in a specific profile
reachy-mini-conversation-app tool-spaces add <owner/space-name> --profile <NAME>
 
# install without enabling
reachy-mini-conversation-app tool-spaces add <owner/space-name> --install-only
 
# list installed spaces
reachy-mini-conversation-app tool-spaces list
 
# remove an installed space
reachy-mini-conversation-app tool-spaces remove <owner/space-name>
```
`add` validates the Space on the Hub, probes the MCP endpoint, discovers its tools, and by default appends the tool IDs to the active profile's `tools.txt`. The active profile is `default` unless you've set `REACHY_MINI_CUSTOM_PROFILE`. Use `--install-only` to skip that step.

`tools.txt` is the gatekeeper: a remote tool is only active if its ID appears in the profile's `tools.txt`, alongside whatever built-in tools you want.


Installed sources are persisted in:

`installed_tool_spaces.json` in managed app mode`external_content/installed_tool_spaces.json` in terminal mode
Each installed Space gets a local alias derived from its slug, with hyphens, dots, and slashes collapsing to underscores:

```
pollen-robotics/reachy-mini-search-tool → pollen_robotics_reachy_mini_search_tool
```
Remote tools are then namespaced with a double underscore:

```
pollen_robotics_reachy_mini_search_tool__search_web
pollen_robotics_reachy_mini_weather_tool__get_day_brief
```
This keeps remote tool names from colliding with built-in ones and lets multiple Spaces coexist in the same profile.

The implementation also strips redundant Space-name prefixes when possible, so a verbose remote tool name becomes a cleaner local ID. If stripping would cause a collision between two tools from the same Space, the code falls back to the fully namespaced name.

There is also a duplicate safety check at registry level: `Tool.name` values must be unique across the entire merged tool set. The app fails fast if two sources claim the same name.

For this work we created two focused canary profiles to isolate the MCP experiment from the full embodied tool set.

The first keeps a few expressive tools (emotions, head movement) and adds web search on top:

```
# profiles/canary_web_search/tools.txt
play_emotion
stop_emotion
idle_do_nothing
move_head
pollen_robotics_reachy_mini_search_tool__search_web
```
The second is the same, plus the weather tool alongside search:

```
# profiles/canary_web_search_weather/tools.txt
play_emotion
stop_emotion
idle_do_nothing
move_head
pollen_robotics_reachy_mini_search_tool__search_web
pollen_robotics_reachy_mini_weather_tool__get_day_brief
```
The small physical tool set means Reachy Mini can still react expressively while answering current questions from the web.

The remote-tool plumbing gets the tools into the model. The prompts decide how the model uses them.

That was especially visible in the search-plus-weather canary. A combined question like:

```
Should I bring a jacket in Bordeaux today, and is there anything major happening downtown tonight?
```
can be handled in at least three ways: weather first then search, search first then weather, or both in the same turn. If the prompt is vague, the model serialises the calls and creates unnecessary latency. So the canary prompts became part of the feature, not just incidental configuration.

