# GoDeploy

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![SSH](https://img.shields.io/badge/SSH-Secure-blue?style=for-the-badge&logo=gnubash)

A lightweight Go tool for automating deployment to multiple remote servers over SSH.

</div>

---

## Features

- Deploy to multiple servers simultaneously
- Secure SSH key-based authentication
- Configurable restart commands per server
- Real-time deployment progress
- Fast SFTP file transfers

## Installation

```bash
git clone https://github.com/yourusername/godeploy.git
cd godeploy
go mod init godeploy
go get golang.org/x/crypto/ssh
go get github.com/pkg/sftp
go build
```

## Configuration

Edit `config/servers.json`:

```json
{
  "servers": [
    {
      "name": "Production Server 1",
      "host": "192.168.1.100",
      "port": 22,
      "user": "deploy",
      "key_path": "/home/user/.ssh/id_rsa",
      "deploy_path": "/var/www/app/binary",
      "restart_commands": [
        "sudo systemctl restart myapp"
      ]
    }
  ]
}
```

## Usage

```bash
./godeploy /path/to/your/binary
```

## Security

- Uses SSH key-based authentication
- Ensure SSH keys have correct permissions: `chmod 600 ~/.ssh/id_rsa`
- Use restricted SSH users with minimal sudo privileges

## License

MIT License

## Contact

Email: contact.amish@yahoo.com
