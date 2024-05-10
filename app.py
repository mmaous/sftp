import os
import paramiko

# local folder path
local_folder_path = os.environ.get("LOCAL_FOLDER_PATH", ".")
remote_folder_path = os.environ.get("REMOTE_FOLDER_PATH", "/")

# remote server credentials
hostname =  os.environ.get("HOST", "sub.example.com")
port = os.environ.get("PORT", 22)  
username = os.environ.get("USERNAME")
password = os.environ.get("PASSWORD")

# the remote server
scp = paramiko.Transport((hostname, port))

scp.connect(username=username, password=password)

# SFTP session
sftp = paramiko.SFTPClient.from_transport(scp)

# the remote directory
sftp.chdir(remote_folder_path)

# count of transferred files
transferred_files_count = 0

# Iterate over files in the local folder
for filename in os.listdir(local_folder_path):
    local_file_path = os.path.join(local_folder_path, filename)
    if os.path.isfile(local_file_path):
        remote_file_path = os.path.join(remote_folder_path, filename)
        # Upload the file to the remote server
        sftp.put(local_file_path, remote_file_path)
        print(f"Uploaded {filename} to {remote_file_path}")
        transferred_files_count += 1

# Close the SFTP session and transport
sftp.close()
scp.close()

# Print count of transferred files
print(f"Transferred {transferred_files_count} files.")

print("All files uploaded successfully.")