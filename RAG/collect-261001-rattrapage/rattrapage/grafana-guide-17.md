---
id: collect-261001-rattrapage/rattrapage/grafana-guide-17
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [3406, 3433]
sha256: 8e8a18a10f0ca52a4112fb50f0f1b38b051e1eb2d51d7fe07ea19bd26dff9371
---

# Guide Grafana — Dashboards, visualisation et alerting

> Après toute modification : `systemctl restart grafana-server`, puis
> `Administration → Settings` pour vérifier la configuration effective.

---

## 83. Annexe B — checklist de mise en production

- [ ] Installé depuis le dépôt APT officiel, version épinglée et notée.
- [ ] `grafana.ini` relu (Annexe A), `root_url` correcte, restart OK.
- [ ] `secret_key` générée, sauvegardée au coffre.
- [ ] Compte `admin` : mot de passe fort au coffre, non utilisé au quotidien.
- [ ] Comptes nominatifs créés, rôles attribués (Viewer par défaut).
- [ ] HTTPS via Nginx (HSTS, redirection 80→443), `cookie_secure = true`.
- [ ] SMTP configuré et testé (e-mail de test reçu).
- [ ] Datasources en mode Server, comptes lecture seule, Save & test au vert.
- [ ] Dashboards clés construits, variables fonctionnelles, provisioning en Git.
- [ ] Alerting : règles avec `for` et labels `severity`/`equipe`, contact
      points testés, arbre de routage relu, test mensuel planifié.
- [ ] Grafana supervisé : métriques internes + sonde externe sur `/api/health`.
- [ ] Sauvegarde quotidienne automatisée + copie hors site + test de
      restauration trimestriel planifié.
- [ ] Documentation : runbooks liés dans les dashboards, mots de passe au
      coffre, procédure de restauration écrite (section 65).
- [ ] Écran mural (si besoin) : orga/compte dédiés, playlist, kiosk mode.

---

*Fin du guide — bon dashboarding !*
