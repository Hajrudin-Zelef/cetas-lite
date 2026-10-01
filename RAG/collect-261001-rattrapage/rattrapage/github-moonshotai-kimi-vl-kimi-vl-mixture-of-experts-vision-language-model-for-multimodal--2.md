---
id: collect-261001-rattrapage/rattrapage/github-moonshotai-kimi-vl-kimi-vl-mixture-of-experts-vision-language-model-for-multimodal--2
title: "If flash-attn has been installed, it is recommended to set torch_dtype=torch.bfloat16 and attn_implementation=\"flash_attention_2\""
domain: rattrapage
role: reference
task: reference
actors: ["Meta", "Moonshot", "OpenAI", "vLLM"]
dates: []
keywords: ["attention", "context window", "decode", "fine-tuning", "gpu", "inference", "kimi", "llama", "lora", "memory", "vllm"]
source: docs/RAG/collect-261001-rattrapage/github-moonshotai-kimi-vl-kimi-vl-mixture-of-experts-vision-language-model-for-multimodal-reasoning-.md
source_anchor: ""
source_lines: [58, 212]
sha256: e0c581f3da0665c8afed162244f3c87da51a66d7fe4f63716d52a8229b31c1b1
---

# If flash-attn has been installed, it is recommended to set torch_dtype=torch.bfloat16 and attn_implementation="flash_attention_2"

```
import torch
from PIL import Image
from transformers import AutoModelForCausalLM, AutoProcessor
model_path = "moonshotai/Kimi-VL-A3B-Instruct"
model = AutoModelForCausalLM.from_pretrained(
    model_path,
    torch_dtype="auto",
    device_map="auto",
    trust_remote_code=True,
)
# If flash-attn has been installed, it is recommended to set torch_dtype=torch.bfloat16 and attn_implementation="flash_attention_2"
# to save memory and speed up inference
# model = AutoModelForCausalLM.from_pretrained(
#     model_path,
#     torch_dtype=torch.bfloat16,
#     device_map="auto",
#     trust_remote_code=True,
#     attn_implementation="flash_attention_2"
# )
processor = AutoProcessor.from_pretrained(model_path, trust_remote_code=True)
image_path = "./figures/demo.png"
image = Image.open(image_path)
messages = [
    {"role": "user", "content": [{"type": "image", "image": image_path}, {"type": "text", "text": "What is the dome building in the picture? Think step by step."}]}
]
text = processor.apply_chat_template(messages, add_generation_prompt=True, return_tensors="pt")
inputs = processor(images=image, text=text, return_tensors="pt", padding=True, truncation=True).to(model.device)
generated_ids = model.generate(**inputs, max_new_tokens=512)
generated_ids_trimmed = [
    out_ids[len(in_ids) :] for in_ids, out_ids in zip(inputs.input_ids, generated_ids)
]
response = processor.batch_decode(
    generated_ids_trimmed, skip_special_tokens=True, clean_up_tokenization_spaces=False
)[0]
print(response)
```
```
import torch
from PIL import Image
from transformers import AutoModelForCausalLM, AutoProcessor
model_path = "moonshotai/Kimi-VL-A3B-Thinking-2506"
model = AutoModelForCausalLM.from_pretrained(
    model_path,
    torch_dtype="auto",
    device_map="auto",
    trust_remote_code=True,
)
# If flash-attn has been installed, it is recommended to set torch_dtype=torch.bfloat16 and attn_implementation="flash_attention_2"
# to save memory and speed up inference
# model = AutoModelForCausalLM.from_pretrained(
#     model_path,
#     torch_dtype=torch.bfloat16,
#     device_map="auto",
#     trust_remote_code=True,
#     attn_implementation="flash_attention_2"
# )
processor = AutoProcessor.from_pretrained(model_path, trust_remote_code=True)
image_paths = ["./figures/demo1.png", "./figures/demo2.png"]
images = [Image.open(path) for path in image_paths]
messages = [
    {
        "role": "user",
        "content": [
            {"type": "image", "image": image_path} for image_path in image_paths
        ] + [{"type": "text", "text": "Please infer step by step who this manuscript belongs to and what it records"}],
    },
]
text = processor.apply_chat_template(messages, add_generation_prompt=True, return_tensors="pt")
inputs = processor(images=images, text=text, return_tensors="pt", padding=True, truncation=True).to(model.device)
generated_ids = model.generate(**inputs, max_new_tokens=32768, temperature=0.8)
generated_ids_trimmed = [
    out_ids[len(in_ids) :] for in_ids, out_ids in zip(inputs.input_ids, generated_ids)
]
response = processor.batch_decode(
    generated_ids_trimmed, skip_special_tokens=True, clean_up_tokenization_spaces=False
)[0]
print(response)
```
Collaborating closely with the open-source community, Kimi-VL now offers seamless support for efficient fine-tuning through the latest version of LLaMA-Factory.

The framework enables Single-GPU LoRA fine-tuning with 50GB of VRAM, as well as Multi-GPU full/lora fine-tuning using DeepSpeed ZeRO-2. For more detailed configuration instructions, check out this PR.

The vLLM main branch has supported Kimi-VL deployment. You are welcome to deploy Kimi-VL using vLLM.

```
from PIL import Image
from transformers import AutoProcessor
from vllm import LLM, SamplingParams
model_path = "moonshotai/Kimi-VL-A3B-Instruct"  # or "moonshotai/Kimi-VL-A3B-Thinking-2506"
llm = LLM(
    model_path,
    trust_remote_code=True,
)
processor = AutoProcessor.from_pretrained(model_path, trust_remote_code=True)
image_path = "./figures/demo.png"
image = Image.open(image_path)
messages = [
    {"role": "user", "content": [{"type": "image", "image": image_path}, {"type": "text", "text": "What is the dome building in the picture? Think step by step."}]}
]
text = processor.apply_chat_template(messages, add_generation_prompt=True, return_tensors="pt")
outputs = llm.generate([{"prompt": text, "multi_modal_data": {"image": image}}], sampling_params = SamplingParams(max_tokens=512))
print("-" * 50)
for o in outputs:
    generated_text = o.outputs[0].text
    print(generated_text)
    print("-" * 50)
```
Serve Kimi-VL with `vllm serve` command:

```
# If you need a longer context window, you can set --max-model-len and --max-num-batched-tokens to 131072
# If you need more input images, you can set --limit-mm-per-prompt image=256 or 512
# kimi-vl-thinking-2506
vllm serve moonshotai/Kimi-VL-A3B-Thinking-2506 --served-model-name kimi-vl-thinking-2506 --trust-remote-code --tensor-parallel-size 1 --max-num-batched-tokens 32768 --max-model-len 32768 --limit-mm-per-prompt image=64
# kimi-vl-instruct
vllm serve moonshotai/Kimi-VL-A3B-Instruct --served-model-name kimi-vl --trust-remote-code --tensor-parallel-size 1 --max-num-batched-tokens 32768 --max-model-len 32768 --limit-mm-per-prompt image=64
```
Call the API

```
import base64
from PIL import Image
from io import BytesIO
from openai import OpenAI
client = OpenAI(
    base_url="http://localhost:8000/v1",
    api_key="token-abc123",
)
image_path = "./figures/demo.png"
image = Image.open(image_path).convert("RGB")
buffered = BytesIO()
image.save(buffered, format="JPEG")
img_b64_str = base64.b64encode(buffered.getvalue()).decode("utf-8")
base64_image_url = f"data:image/jpeg;base64,{img_b64_str}"
messages = [
    {"role": "user", "content": [{"type": "image_url", "image_url": {"url": base64_image_url}}, {"type": "text", "text": "What is the dome building in the picture? Think step by step."}]}
]
completion = client.chat.completions.create(
  model="kimi-vl-thinking-2506", # or kimi-vl
  messages=messages
)
print(completion.choices[0].message)
```
```
@misc{kimiteam2025kimivltechnicalreport,
      title={{Kimi-VL} Technical Report}, 
      author={Kimi Team and Angang Du and Bohong Yin and Bowei Xing and Bowen Qu and Bowen Wang and Cheng Chen and Chenlin Zhang and Chenzhuang Du and Chu Wei and Congcong Wang and Dehao Zhang and Dikang Du and Dongliang Wang and Enming Yuan and Enzhe Lu and Fang Li and Flood Sung and Guangda Wei and Guokun Lai and Han Zhu and Hao Ding and Hao Hu and Hao Yang and Hao Zhang and Haoning Wu and Haotian Yao and Haoyu Lu and Heng Wang and Hongcheng Gao and Huabin Zheng and Jiaming Li and Jianlin Su and Jianzhou Wang and Jiaqi Deng and Jiezhong Qiu and Jin Xie and Jinhong Wang and Jingyuan Liu and Junjie Yan and Kun Ouyang and Liang Chen and Lin Sui and Longhui Yu and Mengfan Dong and Mengnan Dong and Nuo Xu and Pengyu Cheng and Qizheng Gu and Runjie Zhou and Shaowei Liu and Sihan Cao and Tao Yu and Tianhui Song and Tongtong Bai and Wei Song and Weiran He and Weixiao Huang and Weixin Xu and Xiaokun Yuan and Xingcheng Yao and Xingzhe Wu and Xinxing Zu and Xinyu Zhou and Xinyuan Wang and Y. Charles and Yan Zhong and Yang Li and Yangyang Hu and Yanru Chen and Yejie Wang and Yibo Liu and Yibo Miao and Yidao Qin and Yimin Chen and Yiping Bao and Yiqin Wang and Yongsheng Kang and Yuanxin Liu and Yulun Du and Yuxin Wu and Yuzhi Wang and Yuzi Yan and Zaida Zhou and Zhaowei Li and Zhejun Jiang and Zheng Zhang and Zhilin Yang and Zhiqi Huang and Zihao Huang and Zijia Zhao and Ziwei Chen},
      year={2025},
      eprint={2504.07491},
      archivePrefix={arXiv},
      primaryClass={cs.CV},
      url={https://arxiv.org/abs/2504.07491}, 
}
```
