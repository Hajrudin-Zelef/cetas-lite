---
id: collect-240926-huggingface/huggingface/moonshotai-kimi-linear-48b-a3b-instruct-hugging-face
title: "moonshotai-kimi-linear-48b-a3b-instruct-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Hugging Face", "Moonshot", "OpenAI"]
dates: []
keywords: ["kimi", "attention", "benchmarks", "decode", "memory", "throughput", "training", "vllm"]
source: docs/RAG/clean_en/huggingface/moonshotai-kimi-linear-48b-a3b-instruct-hugging-face.md
source_anchor: ""
source_lines: [1, 84]
sha256: 6f9184cd424b48fa80d3cc58fc64376ba577a6893e54f83b461da7e8d300806b
---

# moonshotai-kimi-linear-48b-a3b-instruct-hugging-face

<!-- source: https://huggingface.co/moonshotai/Kimi-Linear-48B-A3B-Instruct -->

### 
  
    
     **Kimi Linear: An Expressive, Efficient Attention Architecture**
    
  

**(a)** On MMLU-Pro (4k context length), Kimi Linear achieves 51.0 performance with similar speed as full attention. On RULER (128k context length), it shows Pareto-optimal performance (84.3) and 3.98x speedup. **(b)** Kimi Linear achieves 6.3x faster TPOT compared to MLA, offering significant speedups at long sequence lengths (1M tokens).

Kimi Linear is a hybrid linear attention architecture that outperforms traditional full attention methods across various contexts, including short, long, and reinforcement learning (RL) scaling regimes. At its core is Kimi Delta Attention (KDA)—a refined version of Gated DeltaNet that introduces a more efficient gating mechanism to optimize the use of finite-state RNN memory.

Kimi Linear achieves superior performance and hardware efficiency, especially for long-context tasks. It reduces the need for large KV caches by up to 75% and boosts decoding throughput by up to $6\times$ for contexts as long as 1M tokens.

We open-source the KDA kernel in FLA, and release two versions model checkpoints trained with 5.7T tokens.

| **Model** | **#Total Params** | **#Activated Params** | **Context Length** | **Download Link** | 
|---|---|---|---|---|
| Kimi-Linear-Base | 48B | 3B | 1M | 🤗 Hugging Face | 
| Kimi-Linear-Instruct | 48B | 3B | 1M | 🤗 Hugging Face | 

- **Kimi Delta Attention (KDA):** A linear attention mechanism that refines the gated delta rule with finegrained gating.
- **Hybrid Architecture:** A 3:1 KDA-to-global MLA ratio reduces memory usage while maintaining or surpassing the quality of full attention.
- **Superior Performance:** Outperforms full attention in a variety of tasks, including long-context and RL-style benchmarks on 1.4T token training runs with fair comparisons.
- **High Throughput:** Achieves up to 6× faster decoding and significantly reduces time per output token (TPOT).

To use the Kimi Linear model, we recommend the following environment:

- `python` >= 3.10
- `torch` >= 2.6
- `fla-core` >= 0.4.0

```
pip install -U fla-core
```
Example Code:

```
from transformers import AutoModelForCausalLM, AutoTokenizer
model_name = "moonshotai/Kimi-Linear-48B-A3B-Instruct"
model = AutoModelForCausalLM.from_pretrained(
    model_name,
    torch_dtype="auto",
    device_map="auto",
    trust_remote_code=True
)
tokenizer = AutoTokenizer.from_pretrained(model_name, trust_remote_code=True)
messages = [
    {"role": "system", "content": "You are a helpful assistant provided by Moonshot-AI."},
    {"role": "user", "content": "Is 123 a prime?"}
]
input_ids = tokenizer.apply_chat_template(
    messages, 
    add_generation_prompt=True, 
    return_tensors="pt"
).to(model.device)
generated_ids = model.generate(inputs=input_ids, max_new_tokens=500)
response = tokenizer.batch_decode(generated_ids)[0]
print(response)
```
For deployment, you can use the latest vllm to create an OpenAI-compatible API endpoint.

```
vllm serve moonshotai/Kimi-Linear-48B-A3B-Instruct \
  --port 8000 \
  --tensor-parallel-size 4 \
  --max-model-len 1048576 \
  --trust-remote-code
```
If you found our work useful, please cite

```
@misc{team2025kimi,
    title         = {Kimi Linear: An Expressive, Efficient Attention Architecture},
    author        = {Zhang, Yu  and Lin, Zongyu  and Yao, Xingcheng  and Hu, Jiaxi  and Meng, Fanqing  and Liu, Chengyin  and Men, Xin  and Yang, Songlin  and Li, Zhiyuan  and Li, Wentao  and Lu, Enzhe  and Liu, Weizhou  and Chen, Yanru  and Xu, Weixin  and Yu, Longhui  and Wang, Yejie  and Fan, Yu  and Zhong, Longguang  and Yuan, Enming  and Zhang, Dehao  and Zhang, Yizhi  and T. Liu, Y.  and Wang, Haiming  and Fang, Shengjun  and He, Weiran  and Liu, Shaowei  and Li, Yiwei  and Su, Jianlin  and Qiu, Jiezhong  and Pang, Bo  and Yan, Junjie  and Jiang, Zhejun  and Huang, Weixiao  and Yin, Bohong  and You, Jiacheng  and Wei, Chu  and Wang, Zhengtao  and Hong, Chao  and Chen, Yutian  and Chen, Guanduo  and Wang, Yucheng  and Zheng, Huabin  and Wang, Feng  and Liu, Yibo  and Dong, Mengnan  and Zhang, Zheng  and Pan, Siyuan  and Wu, Wenhao  and Wu, Yuhao  and Guan, Longyu  and Tao, Jiawen  and Fu, Guohong  and Xu, Xinran  and Wang, Yuzhi  and Lai, Guokun  and Wu, Yuxin  and Zhou, Xinyu  and Yang, Zhilin  and Du, Yulun},
    year          = {2025},
    eprint        = {2510.26692},
    archivePrefix = {arXiv},
    primaryClass  = {cs.CL}
}
```
- Downloads last month
- 183,529
