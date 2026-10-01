---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-47
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI", "United States"]
dates: []
keywords: ["agent", "agents", "benchmarks", "chatgpt", "claude", "cyber", "distribution", "jailbreak", "mai", "mcp"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3675, 3732]
sha256: a4bcd3f2b644a858c5181ef98cda410b7c76d0852ce8c4ca455293be7cbb8c3b
---

# IA — Le grand dossier

- Les jailbreaks « artisanaux » sont un jeu du chat et de la souris permanent : chaque rustine en
  inspire de nouveaux. Aucun labo ne prétend avoir résolu le problème — Anthropic elle-même
  documente des taux de succès non nuls sur ses propres modèles en red-team interne.
- Tendance 2025-2026 : l'automatisation des jailbreaks (agents qui en génèrent) et les attaques
  **multimodales** (image + texte) et **multi-tours** (diluer la demande interdite sur 20 échanges
  anodins). « À vérifier » : les taux précis varient selon les benchmarks (HarmBench, etc.).
- Point important pour ton RAG : un jailbreak réussi contre *ton* assistant expose *tes* documents.
  La parade n'est pas un meilleur modèle, c'est une **architecture** : moindre privilège des outils,
  validation des sorties, cloisonnement des données sensibles hors du contexte.

#### 2.1.3. Exfiltration et confidentialité

- **Memorization** : les LLM mémorisent des fragments de leurs données d'entraînement (noms, e-mails,
  clés, code). Des travaux académiques (Carlini et al., 2021-2023, puis 2024-2025) montrent
  l'extraction ciblée de données d'entraînement par requêtes adversariales. Conséquence : ne jamais
  entraîner/fine-tuner sur des secrets sans anonymisation — et même avec, rester prudent.
- **Fuite par le contexte** : en RAG d'entreprise, le risque n°1 n'est pas le modèle, c'est
  l'**absence de contrôle d'accès au niveau du retrieval** : si ton vector DB renvoie des documents
  RH confidentiels à un utilisateur non autorisé, le modèle les résumera poliment. La sécurité
  d'un RAG = la sécurité de son index (ACL par document, filtrage pré-génération).
- **Shadow AI** : les employés collent du code source, des contrats, des données clients dans des
  chatbots publics. Réponse classique : instance d'entreprise (ex. : Claude/ChatGPT Team/Enterprise
  avec exclusion d'entraînement), DLP, et surtout formation — l'interdiction pure ne marche pas.

#### 2.1.4. Check-list sécurité pour ton RAG / tes agents (opérationnel)

1. Cartographier le « lethal trifecta » de chaque agent : quelles données précieuses ? quel contenu non fiable ? quelles sorties externes ?
2. ACL par document dans l'index vectoriel ; jamais de filtrage « après génération ».
3. Épingler (hash) les définitions d'outils MCP approuvés ; alerter sur tout changement (anti-rug-pull).
4. Journaliser prompts, outils appelés, sorties — l'OWASP le classe dans son Top 10 (journalisation insuffisante).
5. Principe du moindre privilège pour les outils (un agent qui lit des e-mails n'a pas besoin d'un shell).
6. Validation des sorties avant action irréversible (humain dans la boucle sur : envoi d'e-mail, suppression, paiement, publication).
7. Scanner les dépendances (skills, serveurs MCP, extensions) comme du code : la campagne FakeGit prouve que l'écosystème agent est une supply chain comme une autre.
8. Red-team régulière : prompt injection indirecte via tes propres documents (mets un document piégé dans ton corpus et vérifie que l'agent ne l'exécute pas).

### 2.2. Mésusages : deepfakes, désinformation, cyber

#### 2.2.1. Deepfakes et fraude

- **État des lieux 2026** : la génération d'image/vidéo/voix est banale (cf. volumes IA : Sora 2 et consorts) ; la détection reste en retard d'une génération. Les cas documentés : arnaques au « faux PDG » en visioconférence (premier cas majeur à Hong Kong en 2024, 25 M$ — « à vérifier » sur le montant exact), sextorsion par deepfakes intimes non consentis, usurpation de voix pour contourner l'authentification vocale bancaire.
- **Réponses réglementaires** : le **TAKE IT DOWN Act** US (P.L. 119-12, signé le **19 mai 2025**) criminalise la publication non consentie d'images intimes, deepfakes inclus (jusqu'à 2 ans de prison, 3 ans si mineur) et impose aux plateformes un retrait sous 48 h (systèmes de signalement obligatoires depuis le 19 mai 2026). **47 États US** ont légiféré sur les deepfakes (46 sur les deepfakes sexuels, 28 sur les deepfakes politiques) — source : recordinglaw.com, 2026. Côté UE : l'**AI Omnibus (règlement UE 2026/1744, entré en vigueur le 27 juillet 2026**) ajoute à l'article 5 de l'AI Act une **interdiction des images intimes non consenties générées par IA** (applicable au **2 décembre 2026**). Côté Chine : étiquetage obligatoire des contenus synthétiques depuis le 1er septembre 2025 (voir 2.4.3).
- **Technique** : les labels « machine-readable » (métadonnées C2PA, filigranes) sont contournés par simple capture d'écran/recadrage — la traçabilité parfaite n'existe pas ; la parade réaliste combine provenance cryptographique (C2PA), détection statistique et *processus* (vérification par un second canal).

#### 2.2.2. Désinformation à l'échelle

- Le coût marginal de production de propagande ciblée est tombé à ~zéro : campagnes multilingues, faux sites d'info locaux, astroturfing sur les réseaux. Les élections 2024-2026 dans plusieurs pays ont vu des usages documentés (« à vérifier » au cas par cas — les attributions d'ingérence restent un sujet de renseignement, pas de certitude publique).
- **Point de méthode** : l'IA n'a pas inventé la désinformation ; elle a industrialisé sa *production*. La *distribution* reste le goulot (algorithmes des plateformes). D'où l'importance des règles de transparence (article 50 de l'AI Act : marquage des contenus synthétiques, en vigueur depuis le 2 août 2026).

#### 2.2.3. Cyber-offense

- Double usage classique : les mêmes modèles qui trouvent des vulnérabilités pour les corriger (Claude Code fait de l'analyse sémantique au-delà du pattern matching) peuvent les exploiter. Les rapports 2026 documentent des agents capables de chaînes d'attaque semi-autonomes sur des CTF.
- **Position des labs** (Anthropic, OpenAI) : seuils de refus sur les *cyber-capacités offensives* (développement d'exploits, assistance à l'intrusion), avec des « capability thresholds » publiés dans leurs frameworks de sécurité (Anthropic's Responsible Scaling Policy / Frontier Safety Framework — « à vérifier » sur les intitulés exacts en vigueur).
- **Pour un sysadmin défenseur** : l'IA est déjà un multiplicateur défensif (triage d'alertes, rédaction de playbooks, analyse de logs) — cf. les guides Wazuh/Fail2ban du dossier. L'asymétrie attaquant/défenseur reste débattue : les attaquants n'ont pas de change management.

### 2.3. Les voix critiques : qui dit quoi (attribué)

Tableau des positions — *ce sont leurs mots, pas ceux du dossier* :

