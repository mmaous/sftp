# Secure SFTP File Transfer Service

## Overview

This Python application provides a secure file transfer service using the SFTP (SSH File Transfer Protocol) protocol. It enables users to securely upload files to a remote server without storing any credentials locally. The application runs directly on the edge (i.e., your machine), ensuring that sensitive data remains under your control.

## Features

- **Secure File Transfer**: Utilizes SFTP for encrypted file transfer, ensuring data integrity and confidentiality.
- **Credential-Free Operation**: Does not store credentials locally, providing an added layer of security.
- **Edge Computing**: Runs directly on the edge (your machine), minimizing latency and ensuring data sovereignty.

## Usage

1. Clone the repository to your local machine.
2. Set the required environment variables (`LOCAL_FOLDER_PATH`, `REMOTE_FOLDER_PATH`, `HOST`, `PORT`, `USERNAME`, `PASSWORD`) to specify the local and remote folder paths, host, port, and credentials.
3. Run the application using Python.

```bash
python app.py
```

## Docker

> [Docker Repository](https://hub.docker.com/r/mmaous/sftp-uploader)

Alternatively, you can use Docker to run the application. Build the Docker image using the provided Dockerfile and then run the container with the appropriate environment variables.

```bash
docker build -t sftp-app .
docker run --rm -e LOCAL_FOLDER_PATH=/path/to/local/folder -e REMOTE_FOLDER_PATH=/path/to/remote/folder -e HOST=my_host -e PORT=my_port -e USERNAME=my_username -e PASSWORD=my_password sftp-app
```

## Requirements

- Python 3.x
- Paramiko library (`pip install paramiko`)

## Contributing

Contributions are welcome! If you have any suggestions or encounter any issues, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
