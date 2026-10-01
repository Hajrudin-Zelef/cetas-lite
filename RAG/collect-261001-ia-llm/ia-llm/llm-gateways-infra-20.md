---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-20
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Meta", "OpenAI", "xAI"]
dates: ["2026-09-09", "2026-09-14", "2026-09-22", "2026-09-25", "2026-09-27", "2026-10-27"]
keywords: ["astra", "benchmark", "claude", "deepseek", "fable 5", "gemini", "gemini 3.8", "gemini 4", "gpt-5.6", "gpt-6", "grok", "grok 4"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [2399, 2497]
sha256: 86626485fedf459af923513533f4fab49911358ca20cbc7b336e86065c1a3fb3
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

| Lab | Mécanisme de cloisonnement (annoncé) |
|---|---|
| OpenAI | Capacités cybersécurité d'Astra ayant franchi le seuil « Critical » du Preparedness Framework → réservées aux entreprises vérifiées via le programme **Daybreak** (sur dossier) |
| Anthropic | **Mythos 5.1** séparé de Fable 5.1 → accès aux institutions vérifiées via programmes d'accès de confiance (cybersécurité, sciences de la vie) |
| Google | Variante cybersécurité de Gemini 3.8 Flash → programme **Fairwind** (gouvernements et opérateurs d'infrastructures de confiance) |
| Meta | Plus haut niveau de raisonnement **retenu** jusqu'à la fin des tests de sécurité supplémentaires |

**Conséquence pratique pour ton infra :** « le modèle est sorti » ne veut plus dire « tu y as accès ». Quand tu évalues un nouveau modèle, vérifie trois choses dans l'ordre : (1) l'annonce, (2) ton éligibilité au programme d'accès, (3) le prix. Un benchmark sur un modèle auquel tu n'as pas accès est du temps perdu.

### 159.16. Template de fiche RAG « à venir » (à copier tel quel)

Chaque modèle annoncé mais non sorti (ou sorti mais non vérifié par toi) mérite une fiche datée dans ton RAG. Modèle à copier :

```markdown
# FICHE MODÈLE — <nom>
- Statut : ANNONCÉ / SORTI / RUMEUR (rayer les mentions inutiles)
- Date d'annonce : <JJ/MM/AAAA> — source : <URL ou nom de l'éditeur>
- ID d'API : <id exact> ou « aucun publié au <date> »
- Prix vérifiés le : <date> — <montants> ou « aucun publié »
- Endpoint testé : oui/non — date du test : <date>
- Éligibilité : accès libre / programme <nom> / non éligible
- Expiration de la fiche : <date + 30 jours> → à recontrôler
- Décision : en prod / en éval / en attente / écarté — motif : <une ligne>
```

Exemple rempli au 27/09/2026 :

```markdown
# FICHE MODÈLE — DeepSeek V4.1 Pro
- Statut : ANNONCÉ (09/09/2026, membre de l'équipe DeepSeek @tianyi)
- ID d'API : aucun publié au 27/09/2026
- Prix vérifiés le : aucun publié au 27/09/2026
- Endpoint testé : non
- Éligibilité : inconnue
- Expiration de la fiche : 27/10/2026 → à recontrôler
- Décision : en attente — motif : aucun endpoint ni prix, ne pas coder en dur
```

### 159.17. Journal des modifications de la section 159 (à tenir à jour)

| Date de relecture | Changements constatés | Action |
|---|---|---|
| 27/09/2026 | Création de la section. V4.1 Pro annoncé non sorti ; Sonnet/Haiku 5.5 annoncés « dans les prochaines semaines » ; Opus 5.5, GPT-6 Sol/Luna, Grok 4.7 sortis. | Référence initiale |
| __/__/____ | (à compléter à chaque relecture mensuelle) | |

### 159.18. Plan d'action « lundi matin » — que faire de cette section concrètement

1. **Épingle tes modèles aujourd'hui** (section 142). Septembre 2026 a prouvé qu'un alias vendeur peut changer de modèle sous-jacent en une nuit (cas `deepseek-v4-pro` le 14/09, plan retiré in extremis). Tes configs LiteLLM/FreeLLMAPI/scripts n'utilisent que des IDs datés.
2. **Teste Opus 5.5 et GPT-6 Sol sur tes 30 questions d'éval** (section 127) : ce sont les deux sorties « prix cassés » du 22/09/2026 (Opus 5.5 : $4/$20/M contre $5/$25 ; Sol : moitié prix de la série 5.6). Si ton RAG tourne sur Opus 5 ou GPT-5.6, tu paies potentiellement 20 à 50 % trop cher pour la même qualité.
3. **Vérifie ton éligibilité aux programmes d'accès** (159.15) avant de benchmarker : Daybreak (OpenAI), trusted-access (Anthropic/Mythos), Fairwind (Google). Un score sur un modèle auquel tu n'as pas accès ne sert à rien.
4. **Crée les fiches « à venir »** (template 159.16) pour : DeepSeek V4.1 Pro, Sonnet 5.5, Haiku 5.5. Statut : en attente. Expiration : 27/10/2026.
5. **Programme la relecture mensuelle** : un cron le 1er de chaque mois qui t'affiche le journal 159.17 et te force à re-vérifier chaque ligne du tableau 159.10. Exemple :
   ```bash
   # cron mensuel — relecture section 159
   0 9 1 * * grep -A 40 '159.10. Tableau récapitulatif' ~/workspace/user/files/llm_gateways_infra.md | head -45
   ```
6. **Règle d'or** : aucun nom de cette section n'entre dans une config de prod tant que les trois cases ne sont pas cochées — annonce officielle de l'éditeur, ID d'API ou endpoint existant, prix publié (checklist 159.12).

### 159.19. Prix cités dans cette section — tableau consolidé (tous vérifiés le 27/09/2026)

Tous les montants ci-dessous viennent des recherches web du 27/09/2026 citées dans les sous-sections. Ils expirent vite : re-vérifie sur les pages pricing officielles avant toute décision budgétaire.

| Modèle | Entrée /M | Sortie /M | Cache (entrée) /M | Notes |
|---|---|---|---|---|
| Claude Opus 5.5 | $4 | $20 | $0.20 (lectures) | Baisse vs Opus 5 ($5/$25, cache $0.50) ; +30 % de vitesse annoncée |
| GPT-6 Sol | non publié dans l'annonce | non publié | — | Annoncé « moitié prix de la série 5.6 » par OpenAI (22/09/2026) |
| GPT-6 Luna | non publié dans l'annonce | non publié | — | Idem, moitié prix de la série 5.6 |
| Grok 4.7 (< 200K prompt) | $2 | $6 | $0.50 | Identique à Grok 4.6 sur cette tranche |
| Grok 4.7 (> 200K prompt) | $4 | $12 | $1 | Nouvel étage tarifaire ; contexte 500K inchangé |
| DeepSeek V4.1 Flash (hors pic) | $0.15 | $0.60 | — | Pic : $0.30/$1.20 |
| DeepSeek V4 Pro (hors pic) | $0.66 | $1.98 | — | Pic : $1.32/$3.96 ; service maintenu après le 14/09/2026 |
| Doubao Seed 2.1 Pro (Volcano) | ¥6 | ¥30 | ¥1.20 | Contexte 256K |
| Doubao Seed 2.1 Turbo (Volcano) | ¥3 | ¥15 | — | Moitié prix de Pro |
| Doubao Seed 2.1 Pro (ofox.ai, USD) | $0.884 | $4.42 | $0.177 | Agrégateur tiers, sans compte Volcano |
| Doubao Seed 2.1 Turbo (ofox.ai, USD) | $0.442 | $2.212 | $0.085 | Agrégateur tiers |
| Doubao Seed 2.0 Pro (Volcano) | ¥3.20 | ¥16 | — | Entrées ≤ 32K ; remplacé en flagship par Seed 2.1 |

Ligne de conduite : quand un nouveau modèle annoncé arrive avec un prix « moitié moins cher » (Sol, Luna, Opus 5.5), ne migre pas à l'aveugle — repasse tes 30 questions d'éval (section 127) sur l'ancien et le nouveau, compare qualité ET coût réel par tâche, puis bascule le routeur. Un prix divisé par deux avec 5 % de qualité en moins sur tes cas critiques n'est pas une économie.

### 159.20. Mini-FAQ — les questions que cette section va te poser

**Q : DeepSeek V4.1 Pro est-il disponible ?**
R : Non au 27/09/2026. Annoncé le 09/09/2026 par un membre de l'équipe DeepSeek comme successeur à venir, mais aucun endpoint, aucun prix, aucun poids publié. Le plan de transition du 14/09/2026 a été annulé.

**Q : Puis-je utiliser « GPT-6 Terra » ou « Gemini 4 » dans mon routeur ?**
R : Non. Rien d'officialisé au 27/09/2026 pour ces deux noms — ce sont des rumeurs ou des extrapolations. Les coder en dur, c'est garantir une erreur 404 le jour J.

**Q : Sonnet 5.5 / Haiku 5.5 sont-ils sortis ?**
R : Non au 27/09/2026. Anthropic les annonce « dans les prochaines semaines » (semaine du 19–25/09/2026). Fiche « en attente », à recontrôler.

**Q : Que faire si un nom de cette section sort officiellement demain ?**
R : Appliquer la checklist 159.12 (annonce éditeur + ID d'API + prix publié), créer la fiche RAG datée (template 159.16), passer les 30 questions d'éval, puis seulement envisager la prod.

**Q : Cette section sera-t-elle encore valable dans 3 mois ?**
R : Non — et c'est assumé. C'est la seule section du guide avec une date de péremption explicite : relis-la à chaque début de mois via le journal 159.17.

**Q : Où signaler une annonce que cette section aurait ratée ?**
R : Dans ton INFRA.md (section 157), rubrique « veille » : note la date, la source officielle et l'ID d'API. À la relecture mensuelle, la fiche rejoint le tableau 159.10 si les trois cases de la checklist 159.12 sont cochées.

