---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/gerardoemr-unifi-portal-921d8ac1
title: "Edit .env — set UNIFI_HOST, UNIFI_PASSWORD, SECRET_KEY at minimum"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/gerardoemr-unifi-portal-921d8ac1.md
source_anchor: ""
source_lines: [1, 78]
sha256: 2b3c70e773746599d3c69a0281fcc87612ad8354d9e659525d1fa63e2c3c0ae1
---

# Edit .env — set UNIFI_HOST, UNIFI_PASSWORD, SECRET_KEY at minimum

An external captive portal for UniFi Wi-Fi networks.
Collects guest emails, stores them in MySQL, and authorizes devices via the UniFi Controller API.
- Backend: Flask 3 + SQLAlchemy
- Database: MySQL 8
- Frontend: Jinja2 templates · Tailwind CSS (CDN) · Alpine.js · HTMX
- Deployment: Docker Compose
cp .env.example .env
# Edit .env — set UNIFI_HOST, UNIFI_PASSWORD, SECRET_KEY at minimum
docker compose up --build
Portal: http://localhost:5000
Admin:  http://localhost:5000/admin
- In UniFi Network → Settings → Guest Control:
  - Enable Guest Portal
  - Set Portal Customization to External Portal Server
  - Set the portal URL to http://<your-server-ip>:5000
- The controller redirects connecting clients to http://<your-server-ip>:5000/?id=<MAC>&ap=<AP_MAC>&ssid=<SSID>&url=<original_url>
| Setup | UNIFI_CONTROLLER_TYPE | UNIFI_PORT | 
|---|---|---|
| UniFi Network Controller / CloudKey | classic | 8443 | 
| UDM-Pro / UDM-SE / UniFi OS | udm | 443 | 
| Variable | Default | Description | 
|---|---|---|
| SECRET_KEY | — | Flask secret key | 
| UNIFI_HOST | 192.168.1.1 | Controller IP or hostname | 
| UNIFI_PORT | 8443 | Controller port | 
| UNIFI_USERNAME | admin | Controller admin username | 
| UNIFI_PASSWORD | — | Controller admin password | 
| UNIFI_SITE | default | UniFi site name | 
| UNIFI_CONTROLLER_TYPE | classic | classic orudm | 
| UNIFI_VERIFY_SSL | false | Verify controller SSL cert | 
| GUEST_SESSION_MINUTES | 480 | Session duration (minutes) | 
| PORTAL_BRAND | Guest WiFi | Display name shown on portal | 
| PORTAL_LOGO_URL | — | Optional logo image URL | 
| MYSQL_PASSWORD | spotipo | MySQL user password | 
| MYSQL_ROOT_PASSWORD | rootpassword | MySQL root password | 
guest_sessions table:
| Column | Type | Notes | 
|---|---|---|
| id | INT PK | Auto-increment | 
| email | VARCHAR(255) | Guest email (indexed) | 
| mac_address | VARCHAR(17) | Client MAC ( aa:bb:cc:dd:ee:ff ) | 
| ap_mac | VARCHAR(17) | Access point MAC | 
| ssid | VARCHAR(255) | Network name | 
| original_url | VARCHAR(1024) | URL user was trying to reach | 
| authorized | BOOLEAN | Whether UniFi auth succeeded | 
| authorized_at | DATETIME | Timestamp of successful auth | 
| error_message | TEXT | Error details on failure | 
| created_at | DATETIME | Record creation time | 
User connects to WiFi
  → UniFi detects unauthenticated client
  → Redirects to GET /?id=<MAC>&ap=<AP>&ssid=<SSID>&url=<url>
  → User enters email on login page
  → POST /authenticate
      → Validate email
      → Store GuestSession in MySQL
      → Call UniFi Controller API to authorize MAC
      → Redirect to /success?url=<original_url>
  → Success page auto-redirects after 5 seconds
unifi-portal/
├── app/
│   ├── __init__.py      # App factory
│   ├── models.py        # SQLAlchemy models
│   ├── routes.py        # Route handlers
│   ├── unifi.py         # UniFi Controller API client
│   ├── i18n.py          # ES/EN translation strings
│   └── templates/
│       ├── base.html
│       ├── login.html
│       ├── success.html
│       ├── admin.html
│       └── admin_rows.html
├── static/css/
├── config.py
├── wsgi.py
├── requirements.txt
├── Dockerfile
├── docker-compose.yml
└── .env.example
