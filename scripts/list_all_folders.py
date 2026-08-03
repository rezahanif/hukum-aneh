from gdrive_helpers import get_drive_service

def list_all_recursive(service, parent_id, prefix=""):
    folders = []
    page_token = None
    while True:
        resp = (
            service.files()
            .list(
                q=f"'{parent_id}' in parents",
                pageSize=100,
                fields="nextPageToken, files(id, name, mimeType)",
                pageToken=page_token,
            )
            .execute()
        )
        for f in resp.get("files", []):
            if f["mimeType"] == "application/vnd.google-apps.folder":
                folders.append(f)
                print(f"{prefix}{f['name']}  ->  {f['id']}")
                sub = list_all_recursive(service, f["id"], prefix + "  ")
                folders.extend(sub)
        page_token = resp.get("nextPageToken")
        if not page_token:
            break
    return folders

service = get_drive_service()
ROOT = "1vneHF9YxwgSnBh3ashORK0cYPo16vmQS"
list_all_recursive(service, ROOT)
