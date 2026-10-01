---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-25
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Hugging Face", "MiniMax", "Moonshot", "Xiaomi", "Z.ai"]
dates: ["2026-09-27"]
keywords: ["attention", "deepseek", "distillation", "glm", "kimi", "open-weight"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [2471, 2501]
sha256: a7466f6e7851c46bec245e4ab516dcab7d8c4b1a722b299230d12dbb92fcd566
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

Ce qui vieillira le plus vite, par ordre de péremption probable. **Prix API** (semaines) : la guerre des
Flash se joue au centime et les grilles bougent sans préavis — tout prix de ce volume est une
photographie de septembre 2026, à re-vérifier avant tout chiffrage. **Snapshots et IDs** (semaines à
mois) : Qwen3.8-Max-0902 sera remplacé, `deepseek-v4-pro` finira par basculer ou non, V4.1 Pro sortira
(ou non). **Statuts « annoncé »** (mois) : humain-m3 (open-weight planifié), Qwen4 (teasé), la licence
définitive du Kimi K3 — trois dossiers à rouvrir. **Licences non vérifiées** (à vérifier une fois,
puis stable) : MiniMax-M2/M3, GLM-5.3, Qwen3.8-Max/Flash — la première chose à faire avant toute mise
en production durable. **Controverses** (trimestre) : Anthropic–Xiaomi (distillation), la divergence
SWE-bench du GLM-5.3, le sort exact du V4-Pro — à trancher quand des sources primaires parleront. Ce qui
vieillira le moins vite : les **architectures** (GDN, KDA, DSA, attention hybride — les briques restent),
les **ordres de grandeur** (ratios total/actifs, coûts du 1M) et la **méthode** (ne croire qu'une source
primaire, douter des chiffres isolés). Dans six mois, ce volume sera un document d'histoire — c'est
exactement son rôle : fixer ce que l'on savait, quand on le savait, et ce que l'on ne savait pas.

---

*Fin du volume 1 — Encyclopédie des modèles IA : la Chine (février 2026 → 27 septembre 2026).*
*212 sections numérotées de 1 à 212 — sans trou ni doublon (vérifié par script).*

---

**Note de l'auteur.** Ce volume a été rédigé en septembre 2026 à partir de quatre dossiers de recherche
documentaire, sans qu'aucune spécification ne soit inventée : chaque fait renvoie à une source, chaque
doute est marqué « non vérifié au 27/09/2026 ». Si vous n'avez qu'une heure, lisez les sections 97
(grand tableau), 98 (timeline), 99–102 (que retenir) et 209 (10 faits en 30 secondes). Si vous devez
mettre un modèle en production, lisez 200 (check-list) puis vérifiez la licence sur la carte Hugging Face
du checkpoint exact — c'est le seul geste qui compte vraiment. Et si vous croisez un nom de modèle qui
n'est ni dans l'index (§210) ni dans la liste des noms non trouvés (§198) : méfiance, c'est une rumeur,
une erreur ou l'avenir — dans les trois cas, exigez une source primaire.

*Leo — pour Zelef, septembre 2026.*
