---
id: vague2-datacamp/datacamp/kubernetes-certification
title: "Guide de certification Kubernetes : Examens, conseils et ressources d'étude"
domain: datacamp
role: reference
task: article
actors: []
dates: ["2026-09-23"]
keywords: ["cost", "parameters", "training"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/kubernetes-certification.md
source_anchor: ""
source_lines: [1, 60]
sha256: 65d3b45879b12749fc0803c6d637838d6247539fd5a056fb4359db712d60e03d
---

# Guide de certification Kubernetes : Examens, conseils et ressources d'étude

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/kubernetes-certification
- **Site** : DataCamp
- **Type** : Guide / Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp guide by Patrick Brus covers Kubernetes certifications, exam details, study resources, and preparation strategies. Kubernetes certifications are industry-recognized credentials offered by the Cloud Native Computing Foundation (CNCF) in collaboration with the Linux Foundation, validating expertise in managing, deploying, and securing containerized applications. The author passed the CKAD certification before landing a machine learning engineer role.

The guide profiles five CNCF certifications:

1. **CKA (Certified Kubernetes Administrator)** — for professionals managing/maintaining/troubleshooting production clusters (sysadmins, DevOps, MLOps, cloud engineers). Covers networking/security/storage, workloads/scheduling, cluster management, and troubleshooting. Performance-based, 2 hours, 66% pass, $445 (includes one free retake), recertification every 2 years.
2. **CKAD (Certified Kubernetes Application Developer)** — for developers building/deploying apps (backend devs, cloud-native engineers, DevOps, MLOps). Covers cloud-native app design/deployment, ConfigMaps/Secrets/PersistentVolumes, Deployments/Services/Ingress, multi-container Pods, debugging/optimization. Performance-based, 2 hours, 66% pass, $445, recertification every 2 years.
3. **CKS (Certified Kubernetes Security Specialist)** — most advanced; requires CKA. For security engineers, DevOps, MLOps, cloud architects. Covers security best practices (RBAC, network policies, Pod security), securing images/runtimes, security policies/compliance, threat monitoring/detection. Performance-based, 2 hours, 67% pass, $445, recertification every 2 years.
4. **KCNA (Kubernetes and Cloud-Native Associate)** — entry-level, multiple-choice, 90 minutes, 75% pass, $250, recertification every 2 years. Covers Kubernetes/cloud-native fundamentals, architecture/components, containerization basics (Docker, OCI), DevOps practices.
5. **KCSA (Kubernetes and Cloud-Native Security Associate)** — newest, entry-level security credential, multiple-choice, 90 minutes, 75% pass, $250, recertification every 2 years. Covers cloud-native security basics, Kubernetes security fundamentals, threat model, compliance frameworks.

The guide then explains how to get certified: understand exam format, choose learning resources (official Kubernetes documentation, DataCamp courses, Linux Foundation training, books), practice hands-on (Minikube or cloud clusters; deploy pods/deployments/services, configure ConfigMaps/Secrets/PersistentVolumes, RBAC, network policies), register on the CNCF site (schedule within 12 months), and exam-day tips (weight questions, use `--dry-run=client -o yaml`, aliases like `k=kubectl`, bookmark docs). Results arrive in 24–36 hours; free retake if failed. It closes with best practices and study groups/forums (Reddit r/kubernetes, Kubernetes Slack).

## Key points

- Five CNCF/Linux Foundation Kubernetes certifications: CKA, CKAD, CKS, KCNA, KCSA.
- CKA, CKAD, CKS are performance-based, 2-hour online proctored exams; KCNA and KCSA are multiple-choice (90 minutes).
- CKS requires a valid CKA as prerequisite; it is the most advanced certification.
- Exam costs: CKA/CKAD/CKS $445 (incl. one free retake); KCNA/KCSA $250 (incl. one free retake).
- All certifications require recertification every 2 years.
- Pass scores: CKA/CKAD 66%, CKS 67%, KCNA/KCSA 75%.
- Official Kubernetes documentation is allowed during practical exams; efficient navigation is critical.
- Recommended practice: Minikube/cloud clusters, deploying real apps, ConfigMaps/Secrets/PersistentVolumes, RBAC, network policies.
- Exam-day tips: use `--dry-run=client -o yaml`, command aliases (`k=kubectl`), manage time, bookmark docs.
- Mock exams (Killer.sh, Killercoda) are strongly recommended; aim for 85%+ before the real exam.
- Results are available within 24–36 hours; a free retake is included.

## Technical data / figures

| Certification | Level | Format | Duration | Pass | Cost | Prerequisite | Recertification |
|---------------|-------|--------|----------|------|------|--------------|-----------------|
| CKA | Professional | Performance-based | 2 h | 66% | $445 | None (experience recommended) | Every 2 years |
| CKAD | Professional | Performance-based | 2 h | 66% | $445 | None (experience recommended) | Every 2 years |
| CKS | Advanced | Performance-based | 2 h | 67% | $445 | CKA required | Every 2 years |
| KCNA | Entry | Multiple-choice | 90 min | 75% | $250 | None | Every 2 years |
| KCSA | Entry | Multiple-choice | 90 min | 75% | $250 | None | Every 2 years |

| Resource | Purpose |
|----------|---------|
| Kubernetes official docs | Allowed in practical exams; bookmark key sections |
| Killer.sh | Official CNCF exam simulator (included with purchase) |
| Killercoda | Interactive browser-based practice environments |
| Minikube | Local Kubernetes cluster for practice |
| Reddit r/kubernetes, Kubernetes Slack | Study groups and community |

## Why this source matters for the RAG

This guide provides a complete, structured reference on Kubernetes certification paths with precise exam parameters (format, duration, pass scores, costs, prerequisites), making it highly retrievable for career, training, and certification queries. Its practical study strategies and command tips add actionable technical value to the knowledge base.
