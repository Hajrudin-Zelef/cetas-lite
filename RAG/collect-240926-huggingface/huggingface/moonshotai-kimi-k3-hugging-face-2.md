---
id: collect-240926-huggingface/huggingface/moonshotai-kimi-k3-hugging-face-2
title: "moonshotai-kimi-k3-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Anthropic", "Google", "Moonshot", "OpenAI", "Z.ai"]
dates: ["2026-07-09", "2026-07-16", "2026-07-23"]
keywords: ["kimi", "agent", "agentic", "agents", "benchmark", "benchmarks", "claude", "context window", "cyber", "cybersecurity", "fable 5", "gemini"]
source: docs/RAG/clean_en/huggingface/moonshotai-kimi-k3-hugging-face.md
source_anchor: ""
source_lines: [91, 168]
sha256: 6d3f5957a49e6d8a45673829a9b24e6403bf3375ad80e3c6550068c89f13be20
---

# moonshotai-kimi-k3-hugging-face

1. **Reasoning & knowledge benchmarks**
  - **CritPt and AA-LCR.** Scores are cited from Artificial Analysis as of July 23, 2026.
2. **Coding benchmarks**
  - **DeepSWE.** Kimi K3 is evaluated with the Kimi Code harness. The GLM-5.2 score is taken from the GLM-5.2 release blog; all remaining scores are from the official DeepSWE leaderboard, under which Kimi K3 attains 67.3 with the mini-SWE-agent harness. We report the DeepSWE v1.1 tasks.
  - **Terminal-Bench 2.1.** Kimi K3 is evaluated with the Kimi Code harness. For all other models, we report the best score across harnesses: GLM-5.2 with Claude Code (GLM-5.2 release blog); Claude Opus 4.8 and Claude Fable 5 with Terminus 2 (Artificial Analysis); GPT-5.5 and GPT-5.6 Sol with Codex (OpenAI).
  - **ProgramBench.** Kimi K3 is evaluated with the Kimi Code harness. The GLM-5.2 score is from the GLM-5.2 release blog; all other scores are from Vals AI.
  - **SWE-Marathon.** Kimi K3, Claude Opus 4.8, and Claude Fable 5 are evaluated with the Claude Code harness; GPT-5.6 Sol is evaluated with the Codex harness. The GLM-5.2 score is from the GLM-5.2 release blog. Our evaluation is based on an H20-calibrated branch of the official tasks as of July 9, 2026, prior to the final v1.1 release: the Docker images, performance gates, and reference oracles for the GPU tasks have been recalibrated for H20, while the correctness and anti-cheat validators remain unchanged. Additionally, Claude Fable 5 hit fallbacks on 35% of the tasks in our evaluation, which may have negatively impacted its measured performance.
  - **FrontierSWE.** Kimi K3 is evaluated with the Kimi Code harness and GPT-5.6 Sol with the Codex harness; all other results are from FrontierSWE. Dominance scores are recomputed from the raw scores using the official evaluation script and are current as of July 16, 2026.
  - **PostTrainBench.** Scores for GLM-5.2, GPT-5.5, and Claude Opus 4.8 are adopted from the official PostTrainBench results. Kimi K3, Claude Fable 5, and GPT-5.6 Sol are evaluated with the official Harbor implementation at maximum reasoning effort, averaged over three runs on H20 GPUs (instead of H100 in the official setting) — Kimi K3 and Claude Fable 5 with the Claude Code harness, and GPT-5.6 Sol with the Codex harness.
  - **MLS-Bench-Lite.** Kimi K3 is evaluated with the Kimi Code harness; GLM-5.2 and the Claude models with the Claude Code harness; GPT-5.5 and GPT-5.6 Sol with the Codex harness.
  - **SciCode.** Scores are cited from Artificial Analysis as of July 23, 2026.
  - **Kimi Code Bench 2.0 (in-house).** Kimi K3 is evaluated with the Kimi Code harness (it attains 73.7 with the Claude Code harness); GLM-5.2, Claude Opus 4.8, and Claude Fable 5 with the Claude Code harness; GPT-5.5 and GPT-5.6 Sol with the Codex harness. All models are evaluated at maximum reasoning effort, except GPT-5.5, which uses the "xhigh" setting. As the benchmark includes cybersecurity and safety-related tasks, we also disclose the fraction of refused or fallback tasks: Claude Fable 5 hit 13 fallbacks and 1 refusal out of 80 tasks; 10 refusals out of 80 tasks entered GPT-5.6 Sol's cyber guard; GPT-5.5 had 3 refusals out of 80 tasks.
3. **Agentic benchmarks**
  - **OfficeQA Pro.** Each test case provides the agent with the entire PDF corpus, with all PDFs rendered as images and no machine-readable text available.
  - **OfficeQA Pro and SpreadsheetBench 2.** Kimi K3, GLM-5.2, Claude Opus 4.8, and Claude Fable 5 are evaluated with the Claude Code harness; GPT-5.5 and GPT-5.6 Sol are evaluated with the Codex harness.
  - **MCP-Atlas.** All models are evaluated on the 500-task public subset with a 100-turn limit, using Gemini 3.1 Pro as the judge.
  - **AutomationBench.** All models are evaluated on the 600-task public subset, following the official GitHub setup in all other respects.
  - **BrowseComp.** We adopt a context-compaction strategy triggered at 300K tokens. When evaluated with the full 1M-token context window and no context management, Kimi K3 achieves a score of 90.4. The results of Claude Fable 5, Claude Opus 4.8, GPT-5.6 Sol, and GPT-5.5 are cited from Anthropic and OpenAI.
  - **GDPval-AA v2, AA-Briefcase, τ³-Banking, Harvey Lab-AA, and APEX-Agents.** Scores are cited from Artificial Analysis and the APEX-Agents leaderboard as of July 23, 2026. For Harvey Lab-AA, we report the criterion pass rate.
  - **CorpFin v2, Finance Agent v2, and Legal Research Bench.** Scores are cited from Vals AI.
  - **Agents' Last Exam.** Scores are cited from the official leaderboard as of July 23, 2026; we report the leaderboard's primary pass-rate metric. On the leaderboard, each model is paired with a specific harness: Kimi K3 with Kimi Code; GPT-5.6 Sol and GPT-5.5 with Codex; Claude Fable 5, Claude Opus 4.8, and GLM-5.2 with Claude Code.<sup>†</sup> The Claude Fable 5 entry runs at xhigh effort with 40% of tasks annotated as downgraded.
4. **Multimodal benchmarks**
  - Except for ZeroBench, which follows the official setting and is run five times, all multimodal scores are averaged over three runs. MMMU-Pro is evaluated following the official protocol, preserving the original input order and prepending images to the text input.
  - **PerceptionBench** is an in-house benchmark that focuses on atomic visual perception capabilities.

Kimi K3 applies quantization-aware training from the SFT stage onward, using MXFP4 weights with MXFP8 activations for broad hardware compatibility.

You can access Kimi K3's API on https://platform.kimi.ai by selecting `kimi-k3`, and we provide OpenAI/Anthropic-compatible API for you. Currently, Kimi K3 is recommended to run on the following inference engines:


Kimi K3 always has thinking enabled, and will return `reasoning_content`. Thinking effort is configured with the top-level `reasoning_effort` request field, which supports `"low"`, `"high"`, and `"max"` (default `"max"`).

Kimi K3 was trained in the preserved thinking history mode. For multi-turn conversations and tool calls, Kimi K3 requires the complete assistant message returned by the API to be passed back to `messages` as-is — including `reasoning_content` and `tool_calls`, not just `content`:

```
import openai
def chat_with_preserved_thinking(client: openai.OpenAI, model_name: str):
    messages = [
        {
            "role": "user",
            "content": "Tell me three random numbers."
        },
        {
            "role": "assistant",
            "reasoning_content": "I'll start by listing five numbers: 473, 921, 235, 215, 222, and I'll tell you the first three.",
            "content": "473, 921, 235"
        },
        {
            "role": "user",
            "content": "What are the other two numbers you have in mind?"
        }
    ]
    response = client.chat.completions.create(
        model=model_name,
        messages=messages,
        stream=False,
        max_tokens=4096,
        reasoning_effort="max",
    )
    # the assistant should mention 215 and 222 that appear in the prior reasoning content
    print(f"response: {response.choices[0].message.reasoning}")
    return response.choices[0].message.content
```
For full guides and examples (vision input, structured output, partial mode, tool choice, dynamic tool loading, context caching), see the Kimi K3 Quickstart and Thinking Effort.

Kimi K3 works best with Kimi Code CLI as its agent framework. We warmly invite you to give it a try — run Kimi Code in your terminal and select Kimi K3 using the `/model` command. We hope you enjoy building with Kimi K3, and we would love to hear your feedback!

Both the code repository and the model weights are released under the Kimi K3 License.

If you have any questions, please reach out at support@moonshot.ai.

- Downloads last month
- 1,774,987

## Spaces using moonshotai/Kimi-K3 64

## Collection including moonshotai/Kimi-K3

