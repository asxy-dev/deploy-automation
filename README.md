# Deploy Automation

A lightweight Go tool for automating deployment to multiple remote servers over SSH. Upload files, execute commands, and restart services across your infrastructure with a single command.

## Features

- 🚀 Deploy to multiple servers simultaneously
- 🔐 Secure SSH key-based authentication
- 📁 Automatic remote directory creation
- 🔄 Configurable restart commands per server
- 📊 Real-time deployment progress and output
- ⚡ Fast SFTP file transfers
- 🛡️ Graceful error handling per server

## Installation

### Prerequisites

- Go 1.16 or higher
- SSH access to target servers
- SSH private key for authentication

### Build from Source

```bash
git clone https://github.com/yourusername/deploy-automation.git
cd deploy-automation
go mod init deploy-automation
go get golang.org/x/crypto/ssh
go get github.com/pkg/sftp
go build
```

## Configuration

Edit `config/servers.json` to define your server infrastructure:

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
        "sudo systemctl stop myapp",
        "sudo systemctl start myapp",
        "sudo systemctl status myapp"
      ]
    }
  ]
}
```

### Configuration Fields

| Field | Description |
|-------|-------------|
| `name` | Friendly name for the server |
| `host` | Server IP address or hostname |
| `port` | SSH port (typically 22) |
| `user` | SSH username |
| `key_path` | Path to SSH private key |
| `deploy_path` | Remote path where file will be uploaded |
| `restart_commands` | Array of commands to execute after upload |

## Usage

Deploy a file to all configured servers:

```bash
./deploy-automation /path/to/your/binary
```

### Example

```bash
./deploy-automation ./myapp

=== Deploying to Production Server 1 (192.168.1.100) ===
File uploaded to /var/www/app/binary
Executing: sudo systemctl stop myapp
Output: 
Executing: sudo systemctl start myapp
Output: 
Executing: sudo systemctl status myapp
Output: ● myapp.service - My Application
   Active: active (running)
Deployment to Production Server 1 completed

=== All deployments completed ===
```

## Project Structure

```
deploy-automation/
├── main.go                 # Entry point and orchestration
├── deploy/
│   ├── ssh_client.go      # SSH connection management
│   └── deployer.go        # File upload and command execution
├── config/
│   └── servers.json       # Server configuration
└── README.md
```

## How It Works

1. **Configuration Loading**: Reads server definitions from `config/servers.json`
2. **SSH Connection**: Establishes secure SSH connection using private key
3. **File Upload**: Transfers file via SFTP to specified remote path
4. **Command Execution**: Runs restart commands in sequence
5. **Cleanup**: Closes connections and reports status

## Security Considerations

- Uses SSH key-based authentication (no passwords stored)
- Currently uses `InsecureIgnoreHostKey()` for host key verification
- For production, implement proper host key verification
- Ensure SSH keys have appropriate permissions (600)
- Use restricted SSH users with minimal sudo privileges

## Error Handling

- Per-server error isolation: failure on one server doesn't stop others
- Detailed error messages for troubleshooting
- Command output captured and displayed
- Connection failures logged with context

## Troubleshooting

### Permission Denied

Ensure your SSH key has correct permissions:
```bash
chmod 600 ~/.ssh/id_rsa
```

### Command Not Found

Verify commands are available on remote server and user has necessary permissions.

### Connection Timeout

Check firewall rules and ensure SSH service is running on target servers.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Contact

For questions or support, contact: contact.amish@yahoo.com

## Roadmap

- [ ] Password authentication support
- [ ] Parallel deployment option
- [ ] Rollback functionality
- [ ] Deployment hooks (pre/post)
- [ ] Health check verification
- [ ] Deployment history logging
- [ ] Docker container deployment
- [ ] Configuration validation
