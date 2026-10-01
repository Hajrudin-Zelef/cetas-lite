---
id: collect-261001-ia-llm/ia-llm/how-to-build-scalable-web-apps-with-openai-s-privacy-filter-2
title: "Model call → queued endpoint. Hit from the browser via"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "OpenAI"]
dates: []
keywords: ["compute", "gpu"]
source: docs/RAG/collect-261001-ia-llm/how-to-build-scalable-web-apps-with-openai-s-privacy-filter.md
source_anchor: ""
source_lines: [78, 119]
sha256: 13b04a877531417b21014f07c6d1c6c42a259ccb7400aaed21489e57c507c786
---

# Model call → queued endpoint. Hit from the browser via

```
# Model call → queued endpoint. Hit from the browser via
# client.predict("/create_paste", { text, ttl }).
@server.api(name="create_paste")
def create_paste(text: str, ttl: str = "never") -> dict:
    source_text, spans = run_privacy_filter(text)
    redacted = redact(source_text, spans)          # <CATEGORY> placeholders
    pid, reveal_token = secrets.token_urlsafe(6), secrets.token_urlsafe(22)
    PASTES[pid] = Paste(pid, reveal_token, source_text, redacted, spans,
                        expires_at=_ttl(ttl))      # see app.py
    return {
        "view_path":   f"/view/{pid}",
        "reveal_path": f"/view/{pid}?token={reveal_token}",
    }
# View page → plain FastAPI GET. No model, no queue needed, and we
# actually want the bespoke URL shape `/view/{pid}?token=...` that a
# queued endpoint couldn't give us.
@server.get("/view/{pid}", response_class=HTMLResponse)
async def view_paste(pid: str, token: str | None = None):
    p = _store_get(pid)                            # see app.py for store
    if p is None:
        return HTMLResponse(_not_found(), status_code=404)
    revealed = bool(token) and secrets.compare_digest(token, p.reveal_token)
    return HTMLResponse(_render_view(p, revealed))
```
A daemon thread evicts expired pastes every 30 seconds. The whole service, including storage, is about 200 lines of application code because everything lives in one process.

The split across all three apps is the same — anything that touches the model goes through `@server.api`, everything else stays on plain FastAPI routes:

| App | Queued compute ( `@server.api` ) | Plain FastAPI routes | 
|---|---|---|
| Document Privacy Explorer | `analyze_document` — extract, detect, stats | `GET /` serves the custom reader view | 
| Image Anonymizer | `anonymize_screenshot` — OCR, detect, spans → pixel boxes | `GET /` +`GET /examples/*` serve the canvas UI and preloaded examples | 
| SmartRedact Paste | `create_paste` — detect, redact, mint IDs | `GET /` compose page,`GET /view/{pid}?token=...` public + token-gated views,`GET /api/paste/{pid}` JSON lookup | 

`@server.api` gives you Gradio's queue (serialized requests, correct `@spaces.GPU` composition on ZeroGPU, progress events) and it's what the browser hits through `@gradio/client`. The same endpoint is also what `gradio_client` users hit from Python — one function, two SDKs, no duplicated code. Plain `@server.get`/`@server.post` are reserved for the static surfaces: HTML pages, file lookups, cheap dict reads. That's the rule of thumb from the gradio.Server intro post, and it's what makes these three apps feel consistent even though their UIs are very different.

Drop in a resume, a screenshot of a Slack thread, a log line with a token in it. The fun part is seeing what Privacy Filter catches (and occasionally misses) on text you actually care about.

- OpenAI's release post: Introducing OpenAI Privacy Filter
- Model card: openai/privacy-filter on Hugging Face
- Redaction examples and taxonomy on Model card
