---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-17
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Ant", "Anthropic", "DeepSeek", "Huawei", "LongCat", "Moonshot", "Nvidia", "OpenRouter", "SGLang", "Xiaomi", "Z.ai"]
dates: ["2026-08-02", "2026-09-27"]
keywords: ["apache", "ascend", "asic", "attention", "attribution", "benchmark", "benchmarks", "claude", "deepseek", "glm", "kimi", "kv cache"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [1602, 1741]
sha256: 49b18b5b5f377aba16d712f739fd4972d1570d335c52300421907179d1c188cf
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

**Huawei Ascend** (910B...) : les puces IA chinoises qui portent l'entraînement et l'inférence domestiques
(GLM-5 entraîné à 100 % sur Ascend, tutoriel Hy3 sur Ascend NPU). Avec T-Head, Kunlun, Enflame, Muxi
(supportées jour 1 par MiMo-V2.5) et la puce Zhenwu M890 d'Alibaba, elles forment le « silicium
domestique » : la pile matérielle qui rend la Chine indépendante de NVIDIA pour l'IA.

## 138. Glossaire — Vendor-reported

Chiffre de benchmark **fourni par le constructeur**, non vérifié indépendamment. Statut par défaut dans
ce volume sauf mention « indépendant » (AA, LMArena, Artificial Analysis). Les écarts peuvent être
massifs (ex. divergence GLM-5.3 SWE-bench : 57,8 % vs 82,4 %) : un chiffre vendor-reported isolé est une
indication, pas une preuve.

## 139. Quiz — question 1

Quel modèle détenait, à sa sortie en juillet 2026, le titre de « plus gros modèle open-weight de
l'histoire », et avec combien de paramètres totaux ?

## 140. Quiz — question 2

DeepSeek V4 abandonne le MLA pur de V3 au profit de quelle architecture d'attention ? Citez les trois
composantes.

## 141. Quiz — question 3

Quelle est la particularité tarifaire introduite par DeepSeek sur le V4.1-Flash, et à quelles heures
UTC s'applique le tarif « peak » ?

## 142. Quiz — question 4

Citez les trois séries de modèles d'Ant Group (inclusionAI) et ce que chacune désigne.

## 143. Quiz — question 5

Quelle brique d'attention linéaire Moonshot a été réutilisée par Ant Group dans le Ling-3.0-flash, et
dans quelle proportion (couches) ?

## 144. Quiz — question 6

Vrai ou faux : « Kimi 2.7 » existe comme modèle standalone. Justifiez.

## 145. Quiz — question 7

Qu'est-ce que l'« Ox Alpha » apparu en stealth sur OpenRouter le 20 août 2026, et que s'est-il passé
le 26–27 août ?

## 146. Quiz — question 8

Quelle est la restriction inédite de la licence `minimax-h3-community-license-agreement` (2026-08-02),
et quels territoires sont concernés ?

## 147. Quiz — question 9

Quel fait industriel inédit Xiaomi a-t-il réalisé du 15 au 20 septembre 2026 autour du MiMo-V2.6, et
quel était le coût RL publié ?

## 148. Quiz — question 10

Citez trois éléments prouvant que la pile IA chinoise se découple de NVIDIA (modèles ou infra de ce volume).

## 149. Quiz — question 11

Quelle divergence massive et non résolue concerne les benchmarks du GLM-5.3, et sur quel bench ?

## 150. Quiz — question 12

Pourquoi le Qwen3.8-Flash est-il présenté comme important pour l'avenir d'Alibaba, au-delà de ses
propres performances ?

## 151. Quiz — question 13

Quel modèle dense 27B open-weight a battu un flagship fermé occidental sur les benchmarks de code en
août 2026, et sur quels scores ?

## 152. Quiz — questions 14 et 15

14. Quelle clause distingue la « Modified MIT » de Moonshot d'une licence MIT classique, et à partir de
quels seuils se déclenche-t-elle ?
15. DeepSeek V4.1 Pro : sorti ou non au 27/09/2026 ? Que sait-on officiellement, et que dit la rumeur ?

## 153. Quiz — réponses

**1.** **Kimi K3** (Moonshot), **2,8T (2 800B) paramètres totaux**, 104B actifs par token (16 experts sur
896), poids publiés le 27 juillet 2026.
**2.** Une **attention hybride** : **CSA** (Compressed Sparse Attention, stride 4) + **HCA** (Heavily
Compressed Attention, stride 128) + **SWA** (sliding window ~128). Objectif : ~10 % du KV cache de V3.2.
**3.** Tarification **peak/off-peak** : prix **doublés en peak** ($0,15/$0,60 → $0,30/$1,20). Peak :
**01:00–04:00 et 06:00–10:00 UTC**, jours ouvrés hors jours fériés chinois.
**4.** **Ling** = modèles non-thinking ; **Ring** = modèles thinking (raisonnement) ; **Ming** = série
multimodale omni (texte, vision, audio, musique).
**5.** Le **KDA (Kimi Delta Attention)** : **35 couches KDA + 7 couches MLA** (gated MLA) dans le
Ling-3.0-flash — emprunt direct à l'architecture Moonshot.
**6.** **Faux.** Le modèle s'appelle **Kimi K2.7 Code** ; « Kimi 2.7 » seul n'existe pas, c'est
l'abréviation usuelle (post-training code du checkpoint K2.6, 12 juin 2026).
**7.** « Ox Alpha » (`stealth/ox-alpha`) était le **GLM-5.3-Flash de Zhipu en mode furtif** (gratuit une
semaine sur OpenRouter) ; **confirmé par Zhipu à Bloomberg le 26 août**, lancement officiel + poids
MIT le **27 août 2026** (320B/18B, 1M multimodal, servi sur 100 000 puces chinoises).
**8.** Elle **exclut le déploiement local sans autorisation individuelle** aux **États-Unis, Union
européenne, Royaume-Uni et Corée du Sud** — une première : des poids « ouverts » à accès territorial
restreint.
**9.** Xiaomi a **retransmis en direct le post-training RL** sur dashboard public (mimo.xiaomi.com/rl/) ;
coût RL publié : **~$3,47M** ($2,62M Pro + $854K Flash), 30 steps, ~750K trajectoires, <6 jours, 7 000+
environnements RL publiés.
**10.** Au choix : GLM-5 entraîné **100 % sur Huawei Ascend** (zéro NVIDIA) ; GLM-5.3-Flash servi depuis
**100 000 puces de fabrication chinoise** ; LongCat-2.0 entraîné sur **~50 000 ASIC domestiques** ;
MiMo-V2.5 adapté **jour 1 à 7 plateformes de chips chinoises** ; tutoriel SGLang Hy3 sur **Ascend 910B** ;
puce **Zhenwu M890** d'Alibaba lancée avec Qwen3.7-Max.
**11.** **SWE-bench Verified : 57,8 %** (codepick, « données officielles ») **vs 82,4 %** (matrice
ifnodoraemon) — écart massif non résolu au 27/09/2026.
**12.** C'est la **« preview explicite de l'architecture Qwen4 »** (Gated DeltaNet + Qwen Sparse Attention,
176B = 125B backbone + 51B table N-gram) : Alibaba y teaserait l'architecture du prochain flagship avant
de la décliner.
**13.** **Qwen3.8-27B** (dense, ouvert, Apache 2.0, 14 août 2026) : SWE-bench Pro **61,7** (vs Claude Opus
4.6 Max 53,4), LiveCodeBench v6 **90,3** (vs 88,8), IFBench **79,5** (vs 62,5).
**14.** Clause d'**attribution commerciale** : au-delà de **100M d'utilisateurs mensuels ou $20M de
revenus mensuels**, le produit doit afficher le nom du modèle (ex. « Kimi K2.6 ») sur son interface.
**15.** **Non sorti.** Officiellement : **annoncé le 10 septembre 2026** (WeChat DeepSeek), V4.1-Flash
« a comprehensively surpassed V4 Pro », aucune spec ni date publiées. Rumeur (orcarouter.ai, 25 sept.) :
fenêtre de sortie **28–30 septembre 2026**, non étayée par le changelog officiel.

---

*Fin du volume 1 — Encyclopédie des modèles IA : la Chine (février 2026 → 27 septembre 2026).*
*155 sections. Tous les faits sont issus de la recherche documentaire de septembre 2026 ; les points*
*marqués « non vérifié au 27/09/2026 » le restent jusqu'à confirmation par source primaire.*

## 154. DeepSeek : anatomie d'une dépréciation ordonnée

Le cas V4-Pro → V4.1-Flash (10–14 septembre 2026) est un cas d'école de gestion de fin de vie modèle.
Séquence : le 10 septembre, annonce WeChat — V4.1-Flash « a comprehensively surpassed V4 Pro » ; le même
jour, retrait du V4-Flash et redirection de ses routes vers V4.1-Flash ; le 14 septembre à 04:00 UTC,
bascule prévue des requêtes `deepseek-v4-pro` vers V4.1-Flash **au tarif Flash** (donc moins cher pour
l'utilisateur), « jusqu'au lancement de V4.1 Pro ». C'est une dépréciation « douce » : pas de coupure,
une période de transition tarifaire favorable, et un successeur nommé (V4.1 Pro) qui n'existe pas encore.
L'incertitude (benchlm.ai affirme le 26 septembre que la reroute du 14 a été annulée) montre la limite :
sans accès API direct, on ne sait pas quel modèle répond vraiment derrière un ID. Leçon opérationnelle :
**ne jamais coder un ID de modèle en dur sans stratégie de fallback** — en septembre 2026, un ID peut
changer de modèle sous vos pieds en 96 heures.

## 155. DeepSeek : le rapport technique fantôme

