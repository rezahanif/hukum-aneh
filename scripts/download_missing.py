"""Download 1 sample from each folder that we're missing."""
import os, sys
sys.path.insert(0, "/home/z/my-project/scripts")
from gdrive_helpers import get_drive_service, list_folder, download_file

SAMPLES_DIR = "/home/z/my-project/download/samples"

# Folder IDs from listing
FOLDERS = {
    "uu": "1QcBg7duRC7AT_qZZlyXQixLAOekA3I3m",
    "tap_mpr": "1PeUKX7CvxhgxdttQQMYBvLkZl1VDbOyL",
    "inpres": "1KnQsVfO06J6IGudhBbpsX47dC26QeW3q",
    "keppres": "14i_b17iwUGN7tDB6AhIN2nbxZZjAZxRi",
    "pp": "1v0dzbXt5BYx4cP8l1rbQcFAqG-hVW8E0",
    "perppu": "1dGFc-Anax10EyeMKujkSitlwK6fEmZav",
    "uud-1945": "1gMTVbgvAx_Kwvq-JfY6zJXOwSGnlbQmF",
    "perpres": "1RBhkXEH750LSjipVjAFrC8C9rcu0fTCw",
    "perda": "1aYuZPB07M9eINHfprNd22LyLt9kxleGM",
    "Putusan-MK": "18r2ZE9_R5zg72HToTBXwe3qH7Qr9Bidj",
    "JDIH_Kemnaker": "1-Q4-ITbhuvqlME3atWsESMKe1zhafKx0",
    "JDIH_Kemenkeu": "1fY5m3VJBR08Fv1HREUfk1_SuhKyrSqb0",
    "JDIH_Kemendag": "1KytH2-EORzsclM-1Z3rN4XwfVMhCeaK7",
    "JDIH_Komdigi": "1o_VyLh8oPToi-yEv42PNWVFOnRZuGlwE",
    "JDIH_KPU": "14YjUGuu5sS42De5VqqadpGXgZz-6Z-iE",
    # duplicate folders - might contain different content
    "inpres_2": "1ek40u5DxOBef76gHRhSBe1pDC58Bbcff",
    "keppres_2": "1kCVCRTeD0gOnQrRggxTAxkSiJPmAAOx0",
    "JDIH_Kemnaker_2": "10nad3dLQLMptwzsyJqq9Rwoq9Bvim1qG",
}

# Folders we already have samples for
HAVE = set()
for folder in os.listdir(SAMPLES_DIR):
    path = os.path.join(SAMPLES_DIR, folder)
    if os.path.isdir(path):
        HAVE.add(folder)

service = get_drive_service()
for name, fid in FOLDERS.items():
    # Check if we already have a sample (map folder name to our local dir)
    local_name = name.replace("_2", "")
    local_dir = os.path.join(SAMPLES_DIR, local_name)
    if os.path.isdir(local_dir):
        files = [f for f in os.listdir(local_dir) if f.endswith(".pdf")]
        if files:
            print(f"SKIP {name} - already have {files[0]}")
            continue
    
    print(f"\nFetching from {name} ({fid})...")
    try:
        pdfs = list_folder(service, fid, max_files=1)
        if not pdfs:
            print(f"  No PDFs found!")
            continue
        f = pdfs[0]
        os.makedirs(local_dir, exist_ok=True)
        # sanitize filename
        safe_name = f["name"].replace("/", "_").replace("\\", "_")
        dest = os.path.join(local_dir, safe_name)
        print(f"  Downloading: {f['name']}")
        download_file(service, f["id"], dest)
        print(f"  -> {dest}")
    except Exception as e:
        print(f"  ERROR: {e}")

print("\nDone!")
