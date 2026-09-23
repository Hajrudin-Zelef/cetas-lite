# NOTES — corpus `etape7-phaseh-webproxy`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-webproxy` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways` — 751 lignes, 13 chunks.

- 1. Market landscape (September 2026)
- 2. Apache httpd
- 3. nginx
- 4. LiteSpeed (secondary focus)
- 5. Caddy
- 6. Traefik
- 7. HAProxy
- 8. Envoy
- 9. lighttpd (secondary)
- 10. Nginx Proxy Manager (homelab)
- 11. Reverse-proxy playbooks
- 12. SSO integration (Authelia, Authentik, OAuth2-Proxy, Keycloak)
- 13. API gateways
- 14. Performance benchmarks (with provenance)
- 15. Homelab vs enterprise patterns
- 16. Gaps, conflicts & unverified claims
- 17. Glossary
- 18. Source index (verbatim URLs)
- 19. Configuration references (copy-adapt snippets)
- 20. Performance tuning per proxy
- 21. Observability per proxy
- 22. High availability & edge patterns
- 23. OpenResty (Kong/APISIX substrate)
- 24. Kong plugin catalog (2026)
- 25. APISIX plugin catalog (2026)
- 26. Tyk OSS (2026)
- 27. cert-manager YAML reference
- 28. OAuth2-Proxy reference
- 29. Kubernetes: Ingress → Gateway API (2026)
- 30. Decision matrix (2026)
- 31. Migration notes
- 32. Troubleshooting checklist
- 33. nginx beyond HTTP: stream, mail, njs
- 34. Apache module catalog (selection, 2026)
- 35. HAProxy Community vs Enterprise (2026)
- 36. Traefik product line (2026)
- 37. Caddy module ecosystem (xcaddy, 2026)
- 38. API gateway comparison matrix (2026)
- 39. Protocol support matrix (2026)
- 40. TLS feature matrix (2026)
- 41. CVE & security roundup (2024–2026)
- 42. Deployment topologies
- 43. Hardening checklists (per proxy)
- 44. Community & support (2026)
- 45. Reference stacks (2026 recipes)
- 46. Version & EOL summary (2026-09-22)
- 47. Quick FAQ (2026)
- 48. Document control

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
