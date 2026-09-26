---
id: collect-240926-huggingface/huggingface/qwen-qwen-72b-hugging-face-2
title: "下方安装可选，安装可能比较缓慢。"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "Meta", "OpenAI"]
dates: []
keywords: ["attention", "benchmark", "benchmarks", "inference", "license", "qwen", "reasoning", "research", "training"]
source: docs/RAG/clean_en/huggingface/qwen-qwen-72b-hugging-face.md
source_anchor: ""
source_lines: [109, 185]
sha256: 02420024d439910e4cabd72d340c89c701e0b25b8b5a0c9f7fd6b2c7795016f0
---

# 下方安装可选，安装可能比较缓慢。

在分词器方面，相比目前主流开源模型以中英词表为主，Qwen-72B使用了超过15万token大小的词表。 该词表在GPT-4使用的BPE词表`cl100k_base`基础上，对中文、多语言进行了优化，在对中、英、代码数据的高效编解码的基础上，对部分多语言更加友好，方便用户在不扩展词表的情况下对部分语种进行能力增强。
词表对数字按单个数字位切分。调用较为高效的tiktoken分词库进行分词。

我们从部分语种各随机抽取100万个文档语料，以对比不同模型的编码压缩率（以支持100语种的XLM-R为基准值1，越低越好），具体性能见图。

可以看到Qwen-72B在保持中英代码高效解码的前提下，对部分使用人群较多的语种（泰语th、希伯来语he、阿拉伯语ar、韩语ko、越南语vi、日语ja、土耳其语tr、印尼语id、波兰语pl、俄语ru、荷兰语nl、葡萄牙语pt、意大利语it、德语de、西班牙语es、法语fr等）上也实现了较高的压缩率，使得模型在这些语种上也具备较强的可扩展性和较高的训练和推理效率。

在预训练数据方面，Qwen-72B模型一方面利用了部分开源通用语料， 另一方面也积累了海量全网语料以及高质量文本内容，去重及过滤后的语料超过3T tokens。 囊括全网文本、百科、书籍、代码、数学及各个领域垂类。



For position encoding, FFN activation function, and normalization methods, we adopt the prevalent practices, i.e., RoPE relative position encoding, SwiGLU for activation function, and RMSNorm for normalization (optional installation of flash-attention for acceleration).

For tokenization, compared to the current mainstream open-source models based on Chinese and English vocabularies, Qwen-72B uses a vocabulary of over 150K tokens. It first considers efficient encoding of Chinese, English, and code data, and is also more friendly to multilingual languages, enabling users to directly enhance the capability of some languages without expanding the vocabulary. It segments numbers by single digit, and calls the tiktoken tokenizer library for efficient tokenization.

We randomly selected 1 million document corpus of each language to test and compare the encoding compression rates of different models (with XLM-R, which supports 100 languages, as the base value 1). The specific performance is shown in the figure above.

As can be seen, while ensuring the efficient decoding of Chinese, English, and code, Qwen-72B also achieves a high compression rate for many other languages (such as th, he, ar, ko, vi, ja, tr, id, pl, ru, nl, pt, it, de, es, fr etc.), equipping the model with strong scalability as well as high training and inference efficiency in these languages.

For pre-training data, on the one hand, Qwen-72B uses part of the open-source generic corpus. On the other hand, it uses a massive amount of accumulated web corpus and high-quality text content. The scale of corpus reaches over 3T tokens after deduplication and filtration, encompassing web text, encyclopedias, books, code, mathematics, and various domain.

我们选取了MMLU，C-Eval，GSM8K, MATH, HumanEval, MBPP, BBH, CMMLU等目前较流行的benchmark，对模型的中英知识能力、翻译、数学推理、代码等能力进行综合评测。Qwen-72B模型在所有benchmark上均取得了开源模型中的最优表现。

We selected MMLU, C-Eval, GSM8K, MATH, HumanEval, MBPP, BBH, CMMLU, which are currently popular benchmarks, to test the model’s Chinese and English knowledge capabilities, translation, mathematical reasoning, coding and other capabilities. From the following comprehensive evaluation results, we can see that the Qwen model outperform the similarly sized open-source models on all tasks.

| Model | Avg | MMLU | C-Eval | GSM8K | MATH | HumanEval | MBPP | BBH | AGIEval | GaokaoBench | CMMLU | 
|---|---|---|---|---|---|---|---|---|---|---|---|
|  |  | 5-shot | 5-shot | 8-shot | 4-shot | 0-shot | 3-shot | 3-shot | 0-shot | 0-shot | 5-shot | 
| LLaMA2-7B | 24.4 | 46.8 | 32.5 | 16.7 | 3.3 | 12.8 | 20.8 | 38.2 | 21.8 | 18.9 | 31.8 | 
| LLaMA2-13B | 31.3 | 55.0 | 41.4 | 29.6 | 5.0 | 18.9 | 30.3 | 45.6 | 30.9 | 18.2 | 38.4 | 
| LLaMA2-70B | 45.7 | 69.7 | 50.1 | 63.5 | 12.0 | 26.2 | 39.6 | 64.9 | 54.2 | 23.3 | 53.6 | 
| InternLM-20B | 47.2 | 62.1 | 58.8 | 52.6 | 7.9 | 25.6 | 35.6 | 52.5 | 59.0 | 59.0 | 59.0 | 
| Yi-34B | 58.0 | 76.3 | 81.8 | 67.9 | 15.9 | 26.2 | 38.2 | 66.4 | 56.5 | 68.3 | 82.6 | 
| XVERSE-65B | - | 70.8 | 68.6 | 60.3 | - | 26.3 | - | - | - | - | - | 
| **Qwen-7B** | 46.2 | 58.2 | 63.5 | 51.7 | 11.6 | 29.9 | 31.6 | 45.0 | 45.3 | 62.5 | 62.2 | 
| **Qwen-14B** | 52.7 | 66.3 | 72.1 | 61.3 | 24.8 | 32.3 | 40.8 | 53.4 | 51.9 | 52.7 | 71.0 | 
| **Qwen-72B** | **66.4** | **77.4** | **83.3** | **78.9** | **35.2** | **35.4** | **52.2** | **67.7** | **62.5** | **87.6** | **83.6** | 

Qwen-72B采用扩展RoPE base的训练方法，支持32k的外推长度，我们使用arXiv数据进行语言建模评测，PPL（越低越好）结果如下：

Qwen-72B uses the method of extending RoPE base and supports the extrapolation length of 32k. We use arXiv data for language modeling evaluation. The PPL (lower is better) results are as follows:

| Model | Sequence Length |  |  |  |  |  | 
|---|---|---|---|---|---|---|
|  | 8192 | 16384 | 32768 |  |  |  | 
| Qwen-72B | 2.828 | 2.734 | 2.717 |  |  |  | 

我们提供了评测脚本，方便大家复现模型效果，详见链接。提示：由于硬件和框架造成的舍入误差，复现结果如有小幅波动属于正常现象。

We have provided evaluation scripts to reproduce the performance of our model, details as link.

如遇到问题，敬请查阅FAQ以及issue区，如仍无法解决再提交issue。

If you meet problems, please refer to FAQ and the issues first to search a solution before you launch a new issue.

如果你觉得我们的工作对你有帮助，欢迎引用！

If you find our work helpful, feel free to give us a cite.

```
@article{qwen,
  title={Qwen Technical Report},
  author={Jinze Bai and Shuai Bai and Yunfei Chu and Zeyu Cui and Kai Dang and Xiaodong Deng and Yang Fan and Wenbin Ge and Yu Han and Fei Huang and Binyuan Hui and Luo Ji and Mei Li and Junyang Lin and Runji Lin and Dayiheng Liu and Gao Liu and Chengqiang Lu and Keming Lu and Jianxin Ma and Rui Men and Xingzhang Ren and Xuancheng Ren and Chuanqi Tan and Sinan Tan and Jianhong Tu and Peng Wang and Shijie Wang and Wei Wang and Shengguang Wu and Benfeng Xu and Jin Xu and An Yang and Hao Yang and Jian Yang and Shusheng Yang and Yang Yao and Bowen Yu and Hongyi Yuan and Zheng Yuan and Jianwei Zhang and Xingxuan Zhang and Yichang Zhang and Zhenru Zhang and Chang Zhou and Jingren Zhou and Xiaohuan Zhou and Tianhang Zhu},
  journal={arXiv preprint arXiv:2309.16609},
  year={2023}
}
```
我们的代码和模型权重对学术研究完全开放，并支持商用。请查看LICENSE了解具体的开源协议细节。如需商用，请填写问卷申请。

Our code and checkpoints are open to research purpose, and they are allowed for commercial purposes. Check LICENSE for more details about the license. If you have requirements for commercial use, please fill out the form to apply.

如果你想给我们的研发团队和产品团队留言，欢迎加入我们的微信群、钉钉群以及Discord！同时，也欢迎通过邮件（qianwen_opensource@alibabacloud.com）联系我们。

If you are interested to leave a message to either our research team or product team, join our Discord or WeChat groups! Also, feel free to send an email to qianwen_opensource@alibabacloud.com.

- Downloads last month
- 4,027,721
