---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-41
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmarks", "embeddings", "gpu", "kv cache", "nvidia", "open-weight", "reranker"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3153, 3273]
sha256: a2a29d9a3f56e85d63e856a693fe0167309c6a2a4e85d9636e780c811cd9014e
---

# IA — Le grand dossier

1. Les grands providers : hyperscalers (§1)
2. Les labs : qui crée les modèles (§2)
3. Top 50 des fournisseurs IA vérifiés (§3)
4. Les passerelles et le routage intelligent (§4)
5. L'IA locale en profondeur (§5)
6. Hardware chiffré : dimensionner une machine locale (§6)
7. Cas d'usage professionnels concrets (§7)
8. Comparatif local vs cloud : le tableau de décision (§8)
9. Glossaire — 56 termes (§9)
10. Quiz — 10 questions corrigées (§10)
11. Sécurité et conformité de l'IA locale (§11)
12. Évaluer un modèle local avant de le mettre en prod (§12)
13. Coûts cachés du cloud (§13)
14. Plan 90 jours : du zéro au hybride opérationnel (§14)
A. Fiches détaillées des 12 providers clés
B. Recettes d'exploitation (Docker Compose + client Python + supervision)
C. Prix comparés sur 3 modèles de référence
D. Les erreurs API et la stratégie de retry
E. Checklist de mise en production d'une gateway
F. 3 scénarios budgétaires annuels
G. Veille : 10 signaux à surveiller
H. Matrice : quel modèle pour quel usage
I. 5 fiches matérielles détaillées
J. Licences open-weight : le mémo juridique pratique
K. FAQ : 20 questions
L. Lexique express du hardware IA

---

## Annexe M — Prompts système prêts à l'emploi (à adapter)

### M.1. Technicien systèmes & énergies (usage interne)

```
Tu es l'assistant technique du service Systèmes & Énergies.
Règles impératives :
- Réponds en français, de façon concise et structurée (titres, listes).
- Pour toute procédure électrique : rappelle la consignation (VAT, verrouillage)
  avant toute intervention.
- Donne des valeurs avec leurs unités et, si c'est un ordre de grandeur,
  dis-le explicitement (« ~ », « à valider sur la fiche technique »).
- Si la question porte sur un modèle d'équipement précis, demande la référence
  exacte si elle manque au lieu d'inventer.
- Termine par les sources utilisées : [doc, page] quand elles sont fournies
  dans le contexte.
- Si tu ne sais pas : dis « je ne sais pas » et propose la prochaine étape
  (mesure à faire, doc à consulter, qui appeler).
```

### M.2. Extraction JSON stricte (tickets, pièces, relevés)

```
Tu es un extracteur de données. Tu réponds UNIQUEMENT avec un objet JSON valide,
sans texte avant ni après, sans markdown.
Schéma : {"categorie": "string", "urgence": 1|2|3|4|5, "equipement": "string|null",
"resume": "string (max 140 caractères)", "action_requise": "string|null"}
Si une information est absente ou ambiguë : valeur null, jamais d'invention.
Température : 0.
```

### M.3. Relecture de script avant mise en production

```
Tu es un relecteur de code senior (bash, Python, PowerShell).
Analyse le script fourni et réponds en 4 parties :
1. BUGS : erreurs certaines ou probables (ligne par ligne).
2. RISQUES : opérations destructrices, absence de garde-fous, secrets en clair.
3. ROBUSTESSE : set -euo pipefail manquant, gestion d'erreurs, idempotence.
4. DIFF PROPOSÉ : le script corrigé, minimal, commenté.
Ne propose jamais d'exécuter quoi que ce soit : tu analyses, l'humain décide.
```

### M.4. Synthèse d'alertes de supervision

```
Tu es l'opérateur de quart. Voici des alertes brutes (outil, horodatage, message).
Produis :
1. En 3 lignes : la situation (quoi, où, depuis quand).
2. Les alertes probablement liées regroupées par cause racine (avec ton niveau
   de confiance : fort/moyen/faible).
3. Les 3 premières actions recommandées, dans l'ordre, avec pour chacune :
   commande ou clic précis + risque si on ne le fait pas.
Ignore les alertes informatives sauf si elles confirment une cause.
```

---

## Annexe N — 25 erreurs d'exploitation et leurs remèdes

| # | Erreur | Symptôme | Remède |
|---|---|---|---|
| 1 | Contexte trop grand pour la VRAM | OOM au chargement | Baisser `-c`/`max-model-len` ; quantifier le KV cache ; Q4 au lieu de Q8 |
| 2 | Offload CPU involontaire | 2 tok/s au lieu de 100 | `--n-gpu-layers 99` ; vérifier `nvidia-smi` pendant l'inférence |
| 3 | Mauvais chat template | Réponses incohérentes, répétitions | Laisser l'outil appliquer le template natif ; ne pas le forcer à la main |
| 4 | Température trop haute en extraction | JSON invalide 1 fois sur 5 | `temperature: 0` + `response_format: json_object` |
| 5 | Prompt trop long non tronqué | 400 context_length_exceeded | Tronquer/résumer en amont ; routeur `triage-long` en fallback |
| 6 | Clé API en clair dans git | Compromission | Vault + rotation immédiate + audit des spend logs |
| 7 | Pas de timeout client | Threads bloqués pendant les pannes | `timeout: 120` partout + circuit breaker |
| 8 | Retry sans backoff | Facture ×10 pendant une panne | Backoff exponentiel + jitter, max 3 tentatives |
| 9 | Un seul provider par cas critique | Panne totale le jour J | 2 providers minimum + fallbacks testés |
| 10 | Modèle non épinglé | Comportement qui change sans prévenir | Versions datées / tags / révisions HF notés |
| 11 | Logs avec contenu des prompts | Fuite de données sensibles | Ne logger que métadonnées ; chiffrer le reste |
| 12 | Budget sans alerte | Découverte de la facture à 5 000 $ | Alertes à 50/80/100 % + hard stop par clé virtuelle |
| 13 | Embeddings non versionnés | RAG qui se dégrade après changement de modèle | Réindexer tout le corpus à chaque changement d'embedder |
| 14 | Chunks trop gros / trop petits | Précision RAG médiocre | 500–1 000 caractères + overlap 10–20 % ; mesurer précision@k |
| 15 | Pas de reranker | Bruit dans le top-k | Ajouter un reranker (local ou API) : +5–15 pts typiques |
| 16 | Éval uniquement sur benchmarks publics | Surprise en production | 50–100 questions métier + juge ; rejouer à chaque changement |
| 17 | Fine-tune sans eval avant/après | Régression invisible | Baseline gelée ; comparer sur le même jeu |
| 18 | Exposition d'Ollama sans auth | N'importe qui utilise ton GPU | Bind 127.0.0.1 + reverse proxy + clé ; jamais 0.0.0.0 nu |
| 19 | Docker `latest` en prod | Mise à jour surprise un lundi 8h | Tags épinglés ; montées de version en pré-prod |
| 20 | Pas de sauvegarde Postgres LiteLLM | Perte des clés/budgets/logs | `pg_dump` quotidien + test de restauration |
| 21 | Egress oublié | Surprise sur la facture cloud | Artefacts lourds chez le même provider ; egress 0 $ si possible |
| 22 | Shadow AI (clés perso) | Spend non piloté + fuite de données | Clés virtuelles centralisées + charte d'usage |
| 23 | Confiance aveugle au routage coût | Qualité qui chute sur les cas durs | Routage par compétence (§4.7), pas seulement par prix |
| 24 | Oubli de la licence du modèle | Risque juridique (seuils MAU/revenus) | Registre modèle/licence/date à chaque ajout |
| 25 | Aucun runbook de panne | Panique le jour J | §Annexe E + exercice de bascule trimestriel |

---

## Annexe O — Anti-sèche une page (à imprimer)

