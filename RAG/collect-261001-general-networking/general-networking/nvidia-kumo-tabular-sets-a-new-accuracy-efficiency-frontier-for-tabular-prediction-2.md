---
id: collect-261001-general-networking/general-networking/nvidia-kumo-tabular-sets-a-new-accuracy-efficiency-frontier-for-tabular-prediction-2
title: "Tensorize tabular data:"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["distribution", "gpu", "license", "nvidia", "training"]
source: docs/RAG/collect-261001-general-networking/nvidia-kumo-tabular-sets-a-new-accuracy-efficiency-frontier-for-tabular-prediction.md
source_anchor: ""
source_lines: [36, 59]
sha256: 8a543de2670b9092330051a06890d75aad177211a45d09935353c8ecb7acbddd
---

# Tensorize tabular data:

Kumo Tabular works on numerical and categorical columns only, while text, images, or timestamps can be turned into features via built-in pre-processing recipes. A single forward pass covers up to 10 classes, which the library extends to any number of classes with error-correcting output codes. Accuracy may degrade on tables far beyond the training ranges or when the query rows come from a different distribution than the context rows, so, as with any predictive model, validate accuracy and calibration on your own held-out data before deployment.

Kumo Tabular runs via NVIDIA's newly released GPU-native library for `structured-data-models`. The library downloads the weights from the Hub on first use and provides the preprocessing, ensembling, and many-class handling used in our evaluations. The code below is all it takes to go from a `pandas.DataFrame` to a prediction:

```
import sdm  # structured-data-models
# Tensorize tabular data:
table = sdm.TableTensor.from_pandas(pd.load_csv(...), device="cuda")
na_mask = table["target"].isnan()
model = sdm.models.KumoTabular(device="cuda")
pred = model(
    # In-context examples (features/targets):
    x_context=table[~na_mask].drop_columns("target"),
    y_context=table[~na_mask, "target"],
    # Prediction examples (features):
    x_query=table[na_mask].drop_column("target"),
)
```
Kumo Tabular is released under the OpenMDW License Agreement, version 1.1. NVIDIA believes Trustworthy AI is a shared responsibility, and we have established policies and practices to enable development for a wide array of AI applications. When downloaded or used in accordance with our terms of service, developers should work with their supporting model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse. Please report model quality, risk, security vulnerabilities, or NVIDIA AI concerns here.

- **Model Code:** https://github.com/NVIDIA/structured-data-models
- **Model Weights:** https://huggingface.co/nvidia/Kumo-Tabular

We thank David Holzmüller for contributing significant ideas and ablations to Kumo Tabular. We thank Vignesh Kothapalli for his help on Kumo Tabular during his internship.
