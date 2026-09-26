---
id: collect-240926-huggingface/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face-3
title: "Log in once; the token is cached at ~/.cache/huggingface/token"
domain: huggingface
role: reference
task: reference
actors: ["OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["benchmarks", "decode", "gpu", "inference", "latency", "llama", "llama.cpp", "memory", "nvfp4", "nvidia", "omni", "prefill"]
source: docs/RAG/clean_en/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face.md
source_anchor: ""
source_lines: [422, 616]
sha256: be8df564a680e00deaad86af934f1a5ac6d70f350509e2e820db60531baa612b
---

# Log in once; the token is cached at ~/.cache/huggingface/token

USER_PROMPT = (
    "Summarize this PDF page: main topic, section headings, important facts "
    "or bullets, and a brief note on each figure or table. "
    "Do not invent text you cannot read."
)
API_URL = "http://localhost:8000/v1/chat/completions"
MODEL = "nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-NVFP4"
MAX_TOKENS = 32000
DPI = 150
 
 
def page_to_b64(pdf_path: str, idx: int) -> str:
    doc = fitz.open(pdf_path)
    z = DPI / 72.0
    pix = doc.load_page(idx).get_pixmap(matrix=fitz.Matrix(z, z))
    img = Image.frombytes("RGB", [pix.width, pix.height], pix.samples)
    doc.close()
    buf = BytesIO()
    img.save(buf, format="PNG")
    return base64.b64encode(buf.getvalue()).decode("ascii")
 
 
def chat(url, model, b64, text, max_tokens):
    r = requests.post(url, json={
        "model": model,
        "messages": [{"role": "user", "content": [
            {"type": "text", "text": text},
            {"type": "image_url", "image_url": {"url": f"data:image/png;base64,{b64}"}},
        ]}],
        "max_tokens": max_tokens,
        "stream": False,
        "temperature": 0.2,
        "chat_template_kwargs": {"enable_thinking": False},
    }, timeout=120)
    r.raise_for_status()
    return r.json()["choices"][0]["message"]["content"]
 
 
def main():
    p = argparse.ArgumentParser()
    p.add_argument("pdf")
    p.add_argument("--page", type=int, default=0)
    p.add_argument("--all-pages", action="store_true")
    p.add_argument("-o", "--output")
    p.add_argument("--url", default=API_URL)
    p.add_argument("--model", default=MODEL)
    p.add_argument("--max-tokens", type=int, default=MAX_TOKENS)
    a = p.parse_args()
 
    doc = fitz.open(a.pdf); n = len(doc); doc.close()
    pages = range(n) if a.all_pages else [a.page]
    parts = [f"# Extracted: {Path(a.pdf).name}\n\n*Pages: {n}*\n"] if a.all_pages else []
 
    for i in pages:
        print(f"Page {i+1}/{n} ...", file=sys.stderr)
        b64 = page_to_b64(a.pdf, i)
        text = chat(a.url, a.model, b64, f"Page {i+1}.\n\n{USER_PROMPT}", a.max_tokens)
        parts.append(f"\n---\n\n## Page {i+1}\n\n{text.strip()}\n" if a.all_pages else text.strip())
 
    out = "\n".join(parts)
    if a.output:
        Path(a.output).write_text(out + "\n", encoding="utf-8")
    else:
        print(out)
 
if __name__ == "__main__":
    main()
```
**Single page:**

```
python3 pdf_vlm_chat.py /path/to/your_document.pdf --page 0
```
**All pages to markdown:**

```
python3 pdf_vlm_chat.py /path/to/your_document.pdf --all-pages -o extracted.md
```
Edit `USER_PROMPT` in the script for different tasks (detailed extraction, table parsing, etc.).

| Setting | Behavior | 
|---|---|
| **Default (omitted)** | Reasoning is **on** . The model emits chain-of-thought before the final answer, visible in`content` . | 
| `"chat_template_kwargs": {"enable_thinking": false}` | Reasoning is **off** . Only the final answer appears in`content` . | 

To disable reasoning on a request, add to the JSON body:

```
"chat_template_kwargs": {"enable_thinking": false}
```
In the Python heredoc pattern, use `False` (Python boolean), not `false` (invalid Python).

We recommend thinking mode for tasks that involve reasoning and complex understanding. For video, audio, and omni use cases, try both enabling and disabling thinking for best results.

## **Advanced: Budget-Controlled Reasoning**

```
from typing import Any, Dict, List
from openai import OpenAI
from transformers import AutoTokenizer
class ThinkingBudgetClient:
    def __init__(self, base_url: str, api_key: str, tokenizer_name_or_path: str):
        self.tokenizer = AutoTokenizer.from_pretrained(
            tokenizer_name_or_path, trust_remote_code=True
        )
        self.client = OpenAI(base_url=base_url, api_key=api_key)
    def chat_completion(        self,
        model: str,
        messages: List[Dict[str, Any]],
        reasoning_budget: int = 512,
        max_tokens: int = 1024,
        **kwargs,
    ) -> Dict[str, Any]:
        assert max_tokens > reasoning_budget, (
            f"reasoning_budget must be less than max_tokens. "
            f"Got {max_tokens=} and {reasoning_budget=}"
        )
        # Step 1: generate only the reasoning trace up to the requested budget.
        response = self.client.chat.completions.create(
            model=model,
            messages=messages,
            max_tokens=reasoning_budget,
            extra_body={
                "top_k": 1,
                "chat_template_kwargs": {
                    "enable_thinking": True,
                },
            },
            **kwargs,
        )
        reasoning_content = response.choices[0].message.content or ""
        if "</think>" not in reasoning_content:
            print("No </think> found in reasoning content")
            reasoning_content = f"{reasoning_content}</think>\n\n"
        reasoning_tokens_len = len(
            self.tokenizer.encode(reasoning_content, add_special_tokens=False)
        )
        remaining_tokens = max_tokens - reasoning_tokens_len
        assert remaining_tokens > 0, (
            f"No tokens remaining for response ({remaining_tokens=}). "
            "Increase max_tokens or lower reasoning_budget."
        )
        # Step 2: continue from the closed reasoning trace and ask for the final answer.
        continued_messages = messages + [
            {"role": "assistant", "content": reasoning_content}
        ]
        prompt = self.tokenizer.apply_chat_template(
            continued_messages,
            tokenize=False,
            continue_final_message=True,
        )
        response = self.client.completions.create(
            model=model,
            prompt=prompt,
            max_tokens=remaining_tokens,
            extra_body={"top_k": 1},
            **kwargs,
        )
        return {
            "reasoning_content": reasoning_content.strip(),
            "content": response.choices[0].text,
            "finish_reason": response.choices[0].finish_reason,
        }
```
Without explicit settings, vLLM may default to ~32 frames per video regardless of length. Always set `--media-io-kwargs` at server launch (already included in the General Invocation above):

```
--media-io-kwargs '{"video": {"fps": 2, "num_frames": 256}}'
```
Recommended `num_frames` ranges (at `fps=2`):

| GPU memory | Recommended `num_frames` range | 
|---|---|
| **80 GB** (A100/H100) | 128–512 | 
| **≤40 GB** | 64–256 | 

Higher values improve temporal coverage but increase VRAM and prefill time. Start at the low end of the range and increase as your workload and latency budget allow.

1. **Reasoning default:** Reasoning is on by default. If you omit`chat_template_kwargs` , the model will produce chain-of-thought traces in`content` . This is appropriate for text and image inputs.
2. **Video frame sampling:** The default (~32 frames) is too conservative for most real videos. Set`--media-io-kwargs` at server launch.
3. **PDF input format:** The API does not accept raw PDF uploads. Render pages to PNG and send as base64 (see PDF Example above).
4. **`max_tokens` vs `--max-model-len`:**`max_tokens` in the request caps only the completion (generated output). It cannot exceed the server's`--max-model-len` , which is the hard ceiling for prompt + completion combined. Increase the server flag if you need longer outputs.

For Jetson deployments, vLLM, SGLang, Ollama, llama.cpp, and TensorRT Edge-LLM are supported inference frameworks; see the Jetson AI Lab model page for more details.

TensorRT Edge-LLM support is only for Jetson Thor; TensorRT-LLM is not supported on Jetson.

**Total Size:** 354,587,705 data points (~717.0B tokens) 

**Total Number of Datasets:** 1395 dataset entries 

**Dataset partition:** Training [100%], Testing [N/A — evaluation benchmarks used separately], Validation [N/A — evaluation benchmarks used separately] 

**Time period for training data collection:** 2019–2025 

