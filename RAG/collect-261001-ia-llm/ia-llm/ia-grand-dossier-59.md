---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-59
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: ["2025-02-02", "2026-09-27", "2027-12-02"]
keywords: ["agent", "chatgpt", "claude", "embedding", "gpu", "reranker", "valuation"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [4941, 5051]
sha256: a43f8b02ccb1e9494a692967bf8d1b11abc15cfb43ce94b0d019451f57dd25b2
---

# IA — Le grand dossier

1. **Les transitoires** : le passage idle → pleine charge d'un run d'entraînement est quasi
   instantané à l'échelle électrique. Vérifie la tenue des onduleurs aux **sauts de charge**
   (certains onduleurs « intelligents » en mode éco décrochent).
2. **L'harmonique** : les alims à découpage des serveurs GPU polluent — filtre actif ou
   transformateur dédié si le site est sensible.
3. **Le refroidissement liquide** : c'est un **deuxième réseau** (eau + électricité) avec ses
   propres risques (fuite + électricité = ton cauchemar). Détection de fuite par câble
   périmétrique, vannes d'isolement automatiques, eau traitée (pas d'eau du robinet !).
4. **La chaleur fatale** : 14 kW de chaleur, c'est du chauffage gratuit pour des bureaux en
   hiver — échangeur sur le retour CDU. L'IA rend le *heat reuse* rentable même à petite échelle.

### 11.3. L'IA comme outil du service énergies (le retour)

L'IA n'est pas qu'une charge : c'est un instrument de ton métier.

| Usage | Comment | Maturité |
|---|---|---|
| Maintenance prédictive onduleurs/batteries | Séries temporelles (tension, température, impédance) → détection de dérive avant panne | Éprouvé (cf. guide onduleurs du dossier) |
| Optimisation PUE | Le modèle apprend les corrélations météo/charge/consigne → ajuste les consignes de froid | Google l'a fait dès 2016 sur ses DC (DeepMind, −40 % de la facture de froid — « à vérifier » sur le chiffre exact) |
| Analyse de factures et contrats | RAG sur tes contrats + extraction automatique des dépassements de puissance | Faisable cette semaine avec ton app |
| Thermographie augmentée | Coupler caméra thermique + LLM pour rédiger les rapports d'inspection | Prototype rapide |
| Dimensionnement | L'exemple 11.2 + un LLM qui génère la note de calcul à partir de tes hypothèses | Faisable — mais **revérifié par un humain** (voir 1.6 : délégation ≠ compétence) |
| Formation équipe | Quiz générés à partir de tes propres procédures (comme le quiz de la section 8) | Immédiat |

### 11.4. Le message à porter en COMEX

« L'IA va multiplier par 5 à 10 la densité électrique de nos salles techniques. C'est un
investissement réseau + froid + supervision, pas juste "des serveurs". En contrepartie,
l'IA nous donne les outils pour optimiser notre propre consommation. Je propose : (1) un
audit de capacité électrique de nos salles, (2) un pilote de maintenance prédictive sur
les onduleurs, (3) une consigne de marquage énergétique de tout projet IA (kWh/mois). »
— Chiffres d'appui : section 1.5 et 10.1 de ce dossier.

---

## 12. Annexe D — Conformité AI Act express pour une PME qui déploie de l'IA

> Pas un avis juridique. Une check-list de bon sens pour savoir *où tu te situes* au 27/09/2026.
> Pour un avis contraignant : juriste spécialisé.

### 12.1. De quel côté es-tu ? (3 questions)

**Q1. Mets-tu sur le marché ou déploies-tu un système d'IA dans l'UE ?**
Un chatbot interne, un RAG documentaire, un outil de tri de CV, une caméra « intelligente » :
oui, tout ça compte comme « déploiement » (deployer).

**Q2. Ton système fait-il partie des pratiques INTERDITES (art. 5) ?**
Scoring social, manipulation subliminale, exploitation des vulnérabilités, identification
biométrique à distance en temps réel par la police (sauf exceptions), notation des citoyens
par les autorités : **interdit depuis le 02/02/2025**. Pour une PME classique : normalement non.

**Q3. Ton système est-il à HAUT RISQUE (annexes II/III) ?**
Exemples : recrutement / RH (tri de CV, évaluation), éducation (examens), crédit, assurance,
justice, migration, infrastructures critiques, sécurité des produits. **Si oui** : obligations
lourdes (gestion des risques, données, documentation, supervision humaine, évaluation de
conformité) — **applicables au 02/12/2027** (annexe III) grâce à l'Omnibus. D'ici là : préparer.

### 12.2. Ce qui s'applique À TOUS dès maintenant (au 27/09/2026)

| Obligation | Base | Action concrète |
|---|---|---|
| **Transparence chatbot** | Art. 50 | Tout assistant conversationnel doit informer qu'il interagit avec une IA — un bandeau « assistant virtuel » suffit |
| **Marquage des contenus synthétiques** | Art. 50 | Tout contenu généré par IA et diffusé (images, textes marketing, vidéos) doit être identifiable comme tel |
| **AI literacy** | Art. 4 | Former les personnes qui exploitent/supervisent les systèmes IA (proportionné à ton usage — 1h de sensibilisation documentée, c'est déjà ça) |
| **GPAI** (si tu *fournis* un modèle, pas si tu l'utilises) | Art. 51-56 | Ne concerne que les fournisseurs de modèles — pas ton cas si tu utilises Claude/GPT via API |

### 12.3. Check-list PME « IA responsable » (12 points)

1. ☐ Inventaire : liste tous les usages d'IA dans l'entreprise (y compris le shadow AI — sondage interne).
2. ☐ Pour chaque usage : est-ce interdit (art. 5) ? à haut risque (annexe III) ? transparent (art. 50) ?
3. ☐ Chatbots et contenus générés : marquage en place (art. 50).
4. ☐ AI literacy : session de sensibilisation pour les utilisateurs + les superviseurs (tracer la date).
5. ☐ Données : d'où viennent les données d'entraînement/RAG ? (licences, consentements — la leçon du 1,5 Md$ d'Anthropic).
6. ☐ Contrats : que disent les CGU de tes fournisseurs IA sur l'usage de tes données pour l'entraînement ? (opt-out si possible).
7. ☐ Sécurité : threat model « lethal trifecta » par agent (voir 2.1.1), logs conservés.
8. ☐ Biais : si IA en RH/recrutement/évaluation → tester les biais *avant* déploiement (et c'est du haut risque → 02/12/2027).
9. ☐ Sous-traitance : clauses IA dans les contrats (qui est responsable en cas de sortie nocive ?).
10. ☐ Plan haut risque : si concerné, calendrier de mise en conformité d'ici décembre 2027 (ne pas attendre 2027).
11. ☐ Veille : point semestriel sur l'évolution du droit (ce dossier vieillit — revérifier).
12. ☐ Documentation : un dossier par système (finalité, données, modèle, risques, mesures) — c'est l'embryon du dossier de conformité.

### 12.4. Les erreurs classiques des PME

- « On utilise juste ChatGPT, l'AI Act ne nous concerne pas » → **faux** : l'art. 50 (transparence)
  s'applique aux déployeurs, pas seulement aux fournisseurs.
- « On attendra 2027 pour le haut risque » → le 02/12/2027 arrive vite pour une évaluation de
  conformité sérieuse ; et les achats/appels d'offres intègrent déjà ces exigences.
- « La sécurité, c'est le fournisseur qui gère » → **faux** : la prompt injection via *tes*
  documents et *tes* outils, c'est *ta* responsabilité d'architecture (voir scénarios 1-10).
- « On a interdit l'IA, on est tranquilles » → le shadow AI continue sans toi, sans garde-fous.
  Mieux vaut un cadre que l'aveuglement.

---

## 13. Annexe E — Mémos imprimables (une page chacun)

### Mémo 1 — Le RAG en 10 commandements (à afficher près du labo)

1. Chunker par section, 500-1000 tokens, 10-20 % d'overlap.
2. Même modèle d'embedding à l'indexation et à la requête. Toujours.
3. Hybride vectoriel + BM25 (RRF) — surtout en français technique.
4. Reranker sur le top-20 → top-5.
5. Température 0,1-0,3 pour le factuel.
6. Citer les sources, ou dire « je ne sais pas ».
7. ACL par document — jamais de filtrage après génération.
8. 50 questions de référence, rejouées à chaque changement.
9. Sanitizer à l'ingestion (texte invisible, HTML, métadonnées).
10. Journaliser : requêtes, chunks récupérés, coûts.

### Mémo 2 — L'agent en 8 règles (à afficher près de la prod)

