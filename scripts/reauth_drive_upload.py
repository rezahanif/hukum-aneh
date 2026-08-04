Generate a Google OAuth re-authorization URL with drive.file write scope.
Run this, open the URL in browser, authorize, paste the code back.
"""
import json, sys, webbrowser
from urllib.parse import urlencode

CLIENT_SECRET_PATH = "/home/z/my-project/upload/client_secret_858768057989-7ql7keonq637uoejbmh8ptl7fl9cc5h3.apps.googleusercontent.com.json"
TOKEN_PATH = "/home/z/my-project/upload/google_token.json"

with open(CLIENT_SECRET_PATH) as f:
    client = json.load(f)["installed"]

# We need drive.file scope (enough to create/upload files)
SCOPES = ["https://www.googleapis.com/auth/drive.file"]

params = {
    "client_id": client["client_id"],
    "redirect_uri": client["redirect_uris"][0],
    "response_type": "code",
    "access_type": "offline",
    "prompt": "consent",
    "scope": " ".join(SCOPES),
}

auth_url = f"{client['auth_uri']}?{urlencode(params)}"
print("=== Google Drive Upload Authorization ===")
print()
print("Open this URL in your browser:")
print(auth_url)
print()
print("After authorizing, you'll be redirected to localhost with a ?code= parameter.")
print("Copy the FULL redirect URL or just the code value and paste it below.")
print()

code = input("Paste the authorization code here: ").strip()

if not code:
    print("No code provided. Exiting.")
    sys.exit(1)

# If user pasted full URL, extract code
if "code=" in code:
    from urllib.parse import urlparse, parse_qs
    qs = parse_qs(urlparse(code).query)
    code = qs["code"][0]

# Exchange code for tokens
import requests
token_resp = requests.post(
    client["token_uri"],
    data={
        "code": code,
        "client_id": client["client_id"],
        "client_secret": client["client_secret"],
        "redirect_uri": client["redirect_uris"][0],
        "grant_type": "authorization_code",
    },
)
if token_resp.status_code != 200:
    print(f"Token exchange failed: {token_resp.text}")
    sys.exit(1)

token_data = token_resp.json()
# Merge with existing token (keep old refresh if new one missing)
try:
    with open(TOKEN_PATH) as f:
        old = json.load(f)
    token_data.setdefault("refresh_token", old.get("refresh_token"))
except:
    pass

token_data["scopes"] = SCOPES
token_data["token_uri"] = client["token_uri"]
token_data["client_id"] = client["client_id"]
token_data["client_secret"] = client["client_secret"]

with open(TOKEN_PATH, "w") as f:
    json.dump(token_data, f, indent=2)

print(f"Token saved with scopes: {token_data.get('scopes', [])}")
print("You can now run the upload script again.")
