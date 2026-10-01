---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-23
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "MiniMax", "OpenAI", "Perplexity"]
dates: ["2026-09-10", "2026-09-15", "2026-09-24", "2026-09-27"]
keywords: ["apache", "arr", "gemini", "gpu", "lora", "perplexity"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [2415, 2501]
sha256: d4aabebd12eb4ad7ef0835228ff01480dc7af7e4d712b7f8a1805ebb6eb67425
---

# IA générative : image, vidéo, recherche

**Jalon 90 jours** : 1 doc illustrée livrée, 3+ clips en ligne, 12 digests de veille envoyés, coût total < 150 $. Si ces 4 cases sont cochées, l'outillage est rentré dans les mœurs — sinon, c'est le besoin pilote qui était mal choisi, pas les outils.

## 191. Cinq projets types budgétés de bout en bout

Tous les montants : tarifs vérifiés le 27/09/2026, hypothèse taux de rejet ×3 sur les brouillons.

**Projet A — Illustrer une procédure de 20 pages (15 visuels documentaires)**
- Brouillons : 45 images FLUX schnell via fal.ai ≈ 0,14 $
- Masters : 15 images flux-2-pro via BFL ≈ 0,45 $
- Relecture/rejetons PAO : 2 h humaines
- **Total : < 1 $ + 2 h.** Oui, moins d'un dollar.

**Projet B — Clip de formation de 2 minutes (8 clips de 10 s assemblés)**
- Brouillons : 24 clips Kling Standard ≈ 24 × 0,35 $ ≈ 8,40 $
- Masters : 8 clips Kling Pro 1080p ≈ 8 × 1,12 $ ≈ 9 $
- Montage + voix off : 1 journée humaine (DaVinci Resolve)
- **Total : ~18 $ + 1 jour.**

**Projet C — Veille hebdomadaire pendant 1 an (52 digests)**
- Sonar éco : 4 requêtes × 52 semaines ≈ 208 requêtes ≈ **~2-5 $/an**
- Envoi mail : 0 $ (script + SMTP existant)
- Relecture : 15 min/semaine
- **Total : < 5 $/an.** Le poste le moins cher de tout ce guide.

**Projet D — Catalogue produit : 30 visuels « identité verrouillée »**
- Option multi-références : 90 images Nano Banana Pro ≈ 90 × 0,134 $ ≈ 12 $
- Option LoRA local : 1 journée de setup GPU + électricité, puis 0 $/image
- **Seuil** : au-delà de ~30 visuels récurrents du même équipement, le LoRA local gagne.

**Projet E — Bilan annuel d'un service outillé (ordre de grandeur)**
- Perplexity Pro : 240 $/an · Kling Pro : 312 $/an · Veille Sonar : ~5 $/an
- Images API (500 masters/an) : ~15 $/an · 2-3 clips Veo premium : ~10 $/an
- **Total : ~580-670 $/an** — cohérent avec la section 72.

## 192. Dernière checklist : avant de fermer ce guide

- [ ] J'ai noté les **4-5 tarifs** que j'utiliserai vraiment (sections 19, 44, 46, 57, 66, 103, 110).
- [ ] J'ai choisi mon **résidence de données** par niveau de sensibilité (section 78).
- [ ] Mes clés API sont en **variables d'environnement**, 1 clé par usage, alertes activées (section 89).
- [ ] J'ai un **budget mensuel** et une règle d'arrêt (sections 73, 166).
- [ ] Je sais quoi faire quand un fournisseur **change ses prix ou ferme** (sections 88, 168-169).
- [ ] Mon premier besoin pilote est **défini et mesurable** (section 190, semaines 3-4).
- [ ] Je sais que les prix bougent : **re-vérification fin décembre 2026** inscrite à l'agenda.
- [ ] J'ai imprimé la **fiche réflexe** (section 185) pour l'afficher au bureau.

*Si les 8 cases sont cochées : tu es prêt. Sinon, retourne aux sections indiquées — c'est pour ça qu'elles existent.*

## 193. À venir — annonces vérifiées au 27/09/2026

Cette section est la synthèse « radar » du guide : tout ce qui est listé ci-dessous a été **vérifié le 27/09/2026** via recherche web. Le détail sourcé est en PARTIE P (sections 156-158) ; ici, l'essentiel actionnable.

**Annonces confirmées (impact direct sur tes choix)**

1. **Kling 3.0 est la version de référence** (sortie 15/09/2026) : la famille 2.x reste en API mais la doc et les nouveautés (audio natif, motion transfer, multi-prompt) portent sur la 3.0. Si tu démarres aujourd'hui, démarre en 3.0.
2. **FLUX.2 est la génération courante** (Pro/Max/Flex/Dev/Klein), avec Klein 4B sous Apache 2.0 pour le local gratuit. FLUX.1 reste pertinent en local léger (Schnell).
3. **Sora 2 fermé le 24/09/2026** (Videos API incluse) : aucun nouveau projet ne doit en dépendre ; les projets existants doivent migrer (Kling, Veo 3.1, Runway).
4. **Nano Banana Pro (Gemini 3 Pro Image)** : le nouveau haut de gamme image Google (~0,134 $/image), avec génération de texte dans l'image au niveau design.
5. **Jev / System One (TypeSafe AI)** : sorti de stealth le 15/09/2026 (`jev-1.13.0`) — décision sous la seconde à 0,042 $/1M, à intégrer comme **garde-fou** de tes pipelines (sections 119, 168, 183).
6. **Hailuo H3** : MiniMax a divisé ses prix par ~3 (0,70 $ les 6 s) — le brouillon vidéo le moins cher du marché au 27/09/2026.
7. **Runway Gen-4.5** (10/09/2026) : nouveau standard de la suite ; offre Entreprise/Éducation sans tarif public — devis uniquement.
8. **Veo 3.1** (août 2026) : Fast à 0,40 $/8 s pour les brouillons, Standard à 3,20 $/8 s pour le premium — la fourchette la plus lisible du marché vidéo.

**Rumeurs explicitement signalées comme telles (ne pas décider dessus)**

- **FLUX 3 vidéo** : signal faible uniquement, aucune annonce. Si confirmé, il rebattrait les cartes image+vidéo en local.
- **Midjourney V8.2** : attendu d'après le rythme historique (~6 mois après V7), non confirmé.
- **GPT Image 2 / DALL-E 4** : rien d'officiel ; le positionnement actuel d'OpenAI image reste l'édition dans l'écosystème.
- **Réouverture d'une API vidéo OpenAI** : rien n'indique un remplaçant de Sora 2 à court terme.

**Ce que ça change pour ta roadmap (section 190)**

- Verrouille tes choix 2026 sur : **Kling 3.0** (vidéo), **FLUX.2** (image), **Sonar** (veille), **Jev** (contrôle).
- Garde une **ligne budgétaire « prime au premium »** (~10-15 $/an) pour 2-3 visuels/clips Veo/Nano Banana Pro par an : c'est là que le haut de gamme se justifie.
- Re-vérification des tarifs : **fin décembre 2026** — c'est le seul rendez-vous calendaire que ce guide t'impose.

---

*Fin du guide — IA générative : image, vidéo, recherche (v1.0, 27/09/2026).*
*Prochaine re-vérification conseillée des tarifs : fin décembre 2026.*

---

## Note de version

- **v1.0 — 27/09/2026** : version initiale. 193 sections, 7 scripts Python prêts à adapter, 20 prompts commentés, 20 pièges, 20 questions FAQ, quiz 10 questions, 5 projets budgétés.
- Tous les tarifs et versions ont été vérifiés par recherche web le **27/09/2026** ; les points incertains sont marqués **« à vérifier »** dans le texte.
- Historique des modifications : à tenir ici à chaque re-vérification trimestrielle (date, ce qui a changé, sections impactées).
