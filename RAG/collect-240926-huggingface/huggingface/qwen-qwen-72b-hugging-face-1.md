---
id: collect-240926-huggingface/huggingface/qwen-qwen-72b-hugging-face-1
title: "下方安装可选，安装可能比较缓慢。"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "Hugging Face"]
dates: []
keywords: ["alignment", "attention", "decode", "distribution", "fine-tuning", "flash attention", "gpu", "int4", "memory", "qwen", "reasoning", "rotary"]
source: docs/RAG/clean_en/huggingface/qwen-qwen-72b-hugging-face.md
source_anchor: ""
source_lines: [1, 108]
sha256: 2904440182bebaafc5ba747f5d06b1d55fdd10ebecc0b7dfee5c104b42cb3289
---

# 下方安装可选，安装可能比较缓慢。

<!-- source: https://huggingface.co/Qwen/Qwen-72B -->

🤗 Hugging Face   |   🤖 ModelScope   |    📑 Paper    ｜   🖥️ Demo

WeChat (微信)   |   Discord   ｜    API 

**通义千问-72B**（**Qwen-72B**）是阿里云研发的通义千问大模型系列的720亿参数规模的模型。Qwen-72B是基于Transformer的大语言模型, 在超大规模的预训练数据上进行训练得到。预训练数据类型多样，覆盖广泛，包括大量网络文本、专业书籍、代码等。同时，在Qwen-72B的基础上，我们使用对齐机制打造了基于大语言模型的AI助手Qwen-72B-Chat。本仓库为Qwen-72B的仓库。

通义千问-72B（Qwen-72B）主要有以下特点：

1. **大规模高质量训练语料** ：使用超过3万亿tokens的数据进行预训练，包含高质量中、英、多语言、代码、数学等数据，涵盖通用及专业领域的训练语料。通过大量对比实验对预训练语料分布进行了优化。
2. **强大的性能** ：Qwen-72B在多个中英文下游评测任务上（涵盖常识推理、代码、数学、翻译等），效果显著超越现有的开源模型。具体评测结果请详见下文。
3. **覆盖更全面的词表** ：相比目前以中英词表为主的开源模型，Qwen-72B使用了约15万大小的词表。该词表对多语言更加友好，方便用户在不扩展词表的情况下对部分语种进行能力增强和扩展。
4. **较长的上下文支持** ：Qwen-72B支持32k的上下文长度。

如果您想了解更多关于通义千问72B开源模型的细节，我们建议您参阅GitHub代码库。

**Qwen-72B** is the 72B-parameter version of the large language model series, Qwen (abbr. Tongyi Qianwen), proposed by Alibaba Cloud. Qwen-72B is a Transformer-based large language model, which is pretrained on a large volume of data, including web texts, books, codes, etc. Additionally, based on the pretrained Qwen-72B, we release Qwen-72B-Chat, a large-model-based AI assistant, which is trained with alignment techniques. This repository is the one for Qwen-72B.

The features of Qwen-72B include:

1. **Large-scale high-quality training corpora** : It is pretrained on over 3 trillion tokens, including Chinese, English, multilingual texts, code, and mathematics, covering general and professional fields. The distribution of the pre-training corpus has been optimized through a large number of ablation experiments.
2. **Competitive performance** : It significantly surpasses existing open-source models on multiple Chinese and English downstream evaluation tasks (including commonsense, reasoning, code, mathematics, etc.). See below for specific evaluation results.
3. **More comprehensive vocabulary coverage** : Compared with other open-source models based on Chinese and English vocabularies, Qwen-72B uses a vocabulary of over 150K tokens. This vocabulary is more friendly to multiple languages, enabling users to directly further enhance the capability for certain languages without expanding the vocabulary.
4. **Longer context support** : Qwen-72B supports 32k context length.

For more details about the open-source model of Qwen-72B, please refer to the GitHub code repository.

- python 3.8及以上版本
- pytorch 1.12及以上版本，推荐2.0及以上版本
- 建议使用CUDA 11.4及以上（GPU用户、flash-attention用户等需考虑此选项）
- **运行BF16或FP16模型需要多卡至少144GB显存（例如2xA100-80G或5xV100-32G）；运行Int4模型至少需要48GB显存（例如1xA100-80G或2xV100-32G）。**
- python 3.8 and above
- pytorch 1.12 and above, 2.0 and above are recommended
- CUDA 11.4 and above are recommended (this is for GPU users, flash-attention users, etc.)
 **To run Qwen-72B-Chat in bf16/fp16, at least 144GB GPU memory is required (e.g., 2xA100-80G or 5xV100-32G). To run it in int4, at least 48GB GPU memory is requred (e.g., 1xA100-80G or 2xV100-32G).**

运行Qwen-72B，请确保满足上述要求，再执行以下pip命令安装依赖库

To run Qwen-72B, please make sure you meet the above requirements, and then execute the following pip commands to install the dependent libraries.

```
pip install transformers==4.32.0 accelerate tiktoken einops scipy transformers_stream_generator==0.0.4 peft deepspeed
```
另外，推荐安装`flash-attention`库（**当前已支持flash attention 2**），以实现更高的效率和更低的显存占用。

In addition, it is recommended to install the `flash-attention` library (**we support flash attention 2 now.**) for higher efficiency and lower memory usage.

```
git clone https://github.com/Dao-AILab/flash-attention
cd flash-attention && pip install .
# 下方安装可选，安装可能比较缓慢。
# Below are optional. Installing them might be slow.
# pip install csrc/layer_norm
# 如果你的flash-attn版本高于2.1.1，下方不需要安装。
# If the version of flash-attn is higher than 2.1.1, the following is not needed.
# pip install csrc/rotary
```
您可以通过以下代码轻松调用：

You can easily call the model with the following code:

```
from transformers import AutoModelForCausalLM, AutoTokenizer
from transformers.generation import GenerationConfig
# Note: The default behavior now has injection attack prevention off.
tokenizer = AutoTokenizer.from_pretrained("Qwen/Qwen-72B", trust_remote_code=True)
# use bf16
# model = AutoModelForCausalLM.from_pretrained("Qwen/Qwen-72B", device_map="auto", trust_remote_code=True, bf16=True).eval()
# use fp16
# model = AutoModelForCausalLM.from_pretrained("Qwen/Qwen-72B", device_map="auto", trust_remote_code=True, fp16=True).eval()
# use cpu only
# model = AutoModelForCausalLM.from_pretrained("Qwen/Qwen-72B", device_map="cpu", trust_remote_code=True).eval()
# use auto mode, automatically select precision based on the device.
model = AutoModelForCausalLM.from_pretrained("Qwen/Qwen-72B", device_map="auto", trust_remote_code=True).eval()
# Specify hyperparameters for generation. But if you use transformers>=4.32.0, there is no need to do this.
# model.generation_config = GenerationConfig.from_pretrained("Qwen/Qwen-72B", trust_remote_code=True)
inputs = tokenizer('蒙古国的首都是乌兰巴托（Ulaanbaatar）\n冰岛的首都是雷克雅未克（Reykjavik）\n埃塞俄比亚的首都是', return_tensors='pt')
inputs = inputs.to(model.device)
pred = model.generate(**inputs)
print(tokenizer.decode(pred.cpu()[0], skip_special_tokens=True))
# 蒙古国的首都是乌兰巴托（Ulaanbaatar）\n冰岛的首都是雷克雅未克（Reykjavik）\n埃塞俄比亚的首都是亚的斯亚贝巴（Addis Ababa）...
```
关于更多的使用说明，请参考我们的GitHub repo获取更多信息。

For more information, please refer to our GitHub repo for more information.

注：作为术语的“tokenization”在中文中尚无共识的概念对应，本文档采用英文表达以利说明。


基于tiktoken的分词器有别于其他分词器，比如sentencepiece分词器。尤其在微调阶段，需要特别注意特殊token的使用。关于tokenizer的更多信息，以及微调时涉及的相关使用，请参阅文档。

Our tokenizer based on tiktoken is different from other tokenizers, e.g., sentencepiece tokenizer. You need to pay attention to special tokens, especially in finetuning. For more detailed information on the tokenizer and related use in fine-tuning, please refer to the documentation.

Qwen-72B模型规模基本情况如下所示：

The details of the model architecture of Qwen-72B are listed as follows:

| Hyperparameter | Value | 
|---|---|
| n_layers | 80 | 
| n_heads | 64 | 
| d_model | 8192 | 
| vocab size | 151851 | 
| sequence length | 32768 | 

在位置编码、FFN激活函数和normalization的实现方式上，我们也采用了目前最流行的做法， 即RoPE相对位置编码、SwiGLU激活函数、RMSNorm（可选安装flash-attention加速）。

