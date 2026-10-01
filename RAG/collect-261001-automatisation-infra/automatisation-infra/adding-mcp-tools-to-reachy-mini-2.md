---
id: collect-261001-automatisation-infra/automatisation-infra/adding-mcp-tools-to-reachy-mini-2
title: "profiles/default/tools.txt"
domain: automatisation-infra
role: reference
task: reference
actors: ["Google", "Hugging Face", "OpenAI"]
dates: []
keywords: ["agents", "gemini", "mcp", "robotics", "voice"]
source: docs/RAG/collect-261001-automatisation-infra/adding-mcp-tools-to-reachy-mini.md
source_anchor: ""
source_lines: [149, 206]
sha256: a57257e202aeb2c2a68f557d3fbc755e0829fc4046c7cf53f886cd2fa457c1ad
---

# profiles/default/tools.txt

```
[default_prompt]
## CANARY WEB SEARCH RULES
You have one remote tool for current web information.
Use it when the user asks for up-to-date facts, news, live availability, or anything else that may have changed recently.
When the search result already answers the question, answer directly in plain language.
Lead with the answer, not with tool chatter.
For remote lookups that may take a moment, you may give one very short English acknowledgment such as "Let me check that and I'll be right back," then continue.
Answer in English unless the user explicitly asks for another language.
Mention uncertainty briefly if the result snippet is incomplete or ambiguous.
Only mention links when they add value or the user asks for sources.
Keep responses short and spoken-style, as if read aloud by a voice assistant. One or two sentences is usually enough. Skip preamble, lists, headers, and filler. Give just the fact or direct answer the user needs.
```
```
[default_prompt]
## CANARY SEARCH AND WEATHER RULES
You have two remote tools:
- a weather brief tool for compact day weather at a location
- a web search tool for broader current web information
Use the weather tool for today's conditions, temperature, rain chance, sunrise, sunset, or simple advice like whether to bring a jacket.
Use web search for news, events, business hours, travel information, severe alerts, or broader current context.
When the user's question mixes a weather part and a current-info part (for example, "should I bring a jacket in Bordeaux today, and is there anything major happening downtown tonight?"), call both tools in parallel in the same turn. Do not wait for one result before starting the other unless the weather result is needed to narrow the search.
Then merge the results into a single short answer. Cover the weather part first, then the events or news part, in plain connected sentences. Do not label the sections or mention which tool gave which piece.
When the user asks about events, news, or what is happening, give them the actual answer from the search results: name specific events, venues, or headlines. Do not tell the user to check websites, visit listing sites, or look something up themselves. If the search returns nothing concrete, say plainly that you didn't find any notable events, rather than redirecting them elsewhere.
For remote lookups that may take a moment, you may give one very short English acknowledgment such as "Let me check that and I'll be right back," then continue.
Answer in English unless the user explicitly asks for another language.
Do not talk about tool usage unless the user asks.
Keep responses short and spoken-style, as if read aloud by a voice assistant. One or two sentences is usually enough. Skip preamble, lists, headers, and filler. Give just the fact or direct answer the user needs.
```
| Capability | Supported | 
|---|---|
| Install by slug for public, MCP-compatible Gradio Spaces (standard `/gradio_api/mcp/` endpoint) | ✅ | 
| Multiple Spaces at once | ✅ | 
| Per-profile enablement via `tools.txt` | ✅ | 
| Namespaced remote tool IDs | ✅ | 
| Backend-agnostic registration (OpenAI, Gemini, Hugging Face) | ✅ | 
| No arbitrary code downloaded into the local app | ✅ | 
| Private or authenticated Spaces | ❌ | 
| Non-Gradio Spaces | ❌ | 
| Arbitrary raw MCP URLs or non-Hugging Face MCP servers | ❌ | 
| Guaranteed parallel tool orchestration | ❌ | 

Two things are worth calling out. First, the Space has to actually behave like an MCP server; if tool discovery fails, the install fails. Second, prompt instructions can encourage parallel calls but cannot guarantee them. If deterministic orchestration matters for a use case, that logic should move from the prompt into code.

If you want others to use your tool, publish it as a public Gradio Space that exposes the standard MCP endpoint, and keep the tools stateless so they work well over the network. Whether a Space installs depends on this runtime behavior, not on tags.

Tags aren't required for installation, but they help people find compatible Spaces:

`reachy-mini-tool``mcp`
The app now has three kinds of tools sharing one registry: built-in, local custom, and remote MCP tools, and profiles still decide which of them a given assistant can reach. A small, trusted core stays at the center while the optional capabilities around it can be added, tested, and swapped without touching the app itself.

What we're most curious about now is what people build. If you publish a tool Space, tag it `reachy-mini-tool` and `mcp` so others can find it. We'd love to see what Reachy Mini ends up able to do!

*Acknowledgements: Many thanks to Fabien Danieau for proofreading this post and helping test the workflow, to Andres Marafioti for helping test it, and to Remi Fabre and the Pollen Robotics team for the ideas and feedback that shaped the remote tools workflow.*

The `tools.txt` approach makes the MCP integration especially clear. I like how remote tools are kept separate from the built-in robot tools while still being enabled per profile. The distinction between installing a Space and actually activating its tool IDs in `tools.txt` also seems important, particularly when multiple remote tools are being used together.

I’m also using an MCP server on my site, https://tinid.ph/, to work with AI agents, and it has been really useful for me. In particular, it has helped a lot with identifying and fixing technical issues on the site. The idea of keeping tools accessible to AI agents through MCP, rather than building every capability directly into the application, makes a lot of sense.
