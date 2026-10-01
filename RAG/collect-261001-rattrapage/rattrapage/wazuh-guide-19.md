---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-19
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [3391, 3450]
sha256: 35cc7b1940b97ce33e10285365d0c72e3b1046ccd41aabc5d3f341fd63acd849
---

# Guide Wazuh — SIEM & XDR Open Source en production

**Q3. Dans quel ordre sont évalués decoders et règles, et pourquoi cet ordre compte ?**
> R : **Decoder d'abord** (parse le log brut en champs), **règles ensuite** (conditions sur les champs). Un log non décodé ne déclenchera jamais une règle fine.

**Q4. Quelle est la plage d'IDs réservée aux règles locales, et pourquoi ?**
> R : **100000–199999**. Elle évite toute collision avec les IDs du ruleset officiel lors des mises à jour.

**Q5. Citez les 3 modes de FIM et quand utiliser chacun.**
> R : **Planifié** (`frequency` : simple, détection différée), **temps réel** (`realtime` : immédiat, pour binaires/confs critiques), **who-data** (via auditd Linux : identifie *qui* a modifié, pour les confs sensibles). Jamais de temps réel sur `/var/log` ou `/tmp`.

**Q6. Avant d'activer l'active response `firewall-drop`, quelle configuration est obligatoire ?**
> R : La **allowlist** (`<white_list>` : LAN admin, VPN, supervision, IP publiques du SOC). Sans elle, on risque de se bannir soi-même ou de bloquer un équipement légitime.

**Q7. L'indexer affiche `status: red`. Que faites-vous en premier ?**
> R : C'est la **priorité absolue** : identifier les shards non assignés (`_cat/shards`, `_cat/allocation`), vérifier l'espace disque (flood stage à 95 % → lever le bloc `read_only_allow_delete` après avoir libéré de la place, section 69). Sans indexer, le SIEM est aveugle.

**Q8. Le dashboard est vide alors que les agents sont verts. Citez 3 maillons à vérifier dans l'ordre.**
> R : 1) `alerts.json` se remplit-il sur le manager ? 2) **Filebeat** tourne-t-il et envoie-t-il (TLS, 401) ? 3) Le *doc count* des index augmente-t-il côté indexer ? (+ vérifier la plage de temps et le pattern d'index dans le dashboard.)

**Q9. Quelle est la règle d'or des mises à jour Wazuh (ordre et précautions) ?**
> R : **Sauvegarde + changelog + fenêtre de maintenance** d'abord ; ordre : **indexer → manager → dashboard** ; manager toujours avant les agents ; `apt-mark hold` en temps normal ; mise à jour des agents **par vagues** (WPK), jamais tout le parc d'un coup.

**Q10. En tant que petite équipe, quel rituel hebdomadaire est le plus important et pourquoi ?**
> R : Le **tuning** (top 10 des règles bruyantes → 1–2 actions) couplé à la **revue des CVE critiques**. Sans tuning, le volume d'alertes sature l'équipe et les vraies attaques passent inaperçues ; sans revue CVE, le patch management dérive. Objectif : < 50 alertes/jour/analyste à qualifier.

---

## 87. Pour aller plus loin

### Documentation officielle (à mettre en favori)

- Documentation Wazuh 4.x : `https://documentation.wazuh.com/current/` — la référence, à jour à chaque version.
- Référence des règles : `https://documentation.wazuh.com/current/user-manual/ruleset/rules/` — descriptif de chaque règle officielle.
- API reference : `https://documentation.wazuh.com/current/user-manual/api/reference.html` — tous les endpoints.

### Communauté

- Slack Wazuh (`wazuh.slack.com`) : entraide rapide, canal `#general` et `#questions`.
- GitHub `wazuh/wazuh` : issues, propositions de règles, contributions.
- Forum et blog Wazuh : retours d'expérience, cas d'usage.

### Sujets d'approfondissement (par ordre de valeur pour une petite équipe)

1. **Règles de corrélation avancées** (`if_matched_sid`, `frequency`, `timeframe`, `same_source_ip`) : passer de la détection unitaire à la détection de campagnes.
2. **Threat intelligence** : intégrer des flux (MISP, listes d'IOC) via CDB lists dynamiques.
3. **Déploiement distribué + HA** (section 12) quand le parc dépasse ~150 agents.
4. **SOAR avec Shuffle** (section 55) : automatiser l'enrichissement et la réponse de niveau 1.
5. **Surveillance des conteneurs** : agent Wazuh + inventaire Docker/Kubernetes (wodle `docker-listener`).
6. **Conformité avancée** : écrire des politiques SCA métier complètes (section 43) et les présenter en audit.
7. **Sauvegarde 3-2-1 complète** : intégrer les snapshots Wazuh dans votre PRA existant (voir guide Proxmox : PBS, réplication).
8. **Exercices d'équipe** : rejouez les cas pratiques 1–6 (sections 75–80) chaque semestre ; mesurez le temps de détection et de réponse, et améliorez.

### Derniers conseils

- **Commencez petit** : 5 agents pilotes, 2 semaines de tuning, puis généralisez. Un déploiement « big bang » sur 200 agents sans tuning = 3 mois de bruit.
- **Le SIEM est un organisme vivant** : règles, tuning, mises à jour, exercices. Prévoyez ~0,5 ETP pour le faire vivre sur un parc de 100–200 machines.
- **Mesurez** : temps moyen de détection, temps moyen de qualification, % de faux positifs, score CIS. Ce qui se mesure s'améliore — et se défend en budget.

---

*Fin du guide. Bon SOC !*
