# Gitz

**Gitz** is a high-performance, concurrent deployment and remote execution orchestrator. Designed for engineers managing proprietary cloud infrastructures, Gitz provides a streamlined way to synchronize code and execute commands across multiple nodes simultaneously using SSH and SFTP.

## Installation

Gitz is distributed via a custom Debian repository. You can install it on any compatible system using the following commands:

```bash
# Add the GPG key
curl -fsSL http://me.iamgrudge.online/gitz-repo.gpg \
| sudo gpg --dearmor -o /usr/share/keyrings/gitz.gpg

# Add the repository to your sources
echo "deb [signed-by=/usr/share/keyrings/gitz.gpg] http://me.iamgrudge.online stable main" \
| sudo tee /etc/apt/sources.list.d/gitz.list

# Update and install
sudo apt update && sudo apt install gitz

```

## Core Features

* **Concurrency by Default:** Leverages Go goroutines to handle multiple node connections in parallel.
* **Remote Execution:** A built-in engine (`runz`) for executing shell commands across the cluster.
* **Infrastructure as Code:** Simple YAML-based configuration for managing node inventories.
* **Flexible Filtering:** Uses `.gitzignore` patterns to exclude files from synchronization.
* **Secure:** Built-in support for SSH private key authentication.

## Usage

### 1. Initialize

Set up a new project configuration:

```bash
gitz init

```

### 2. Configure

Edit `.gitz/gitz.yaml` to define your target environment:

```yaml
project_name: MyCloudApp
nodes:
  - ip: 192.168.1.10
    user: root
    path: /var/www/app

```

### 3. Deploy and Execute

Sync your files and run remote commands:

```bash
gitz push
gitz run "ls -la"

```

## Contributing

The project is hosted at [github.com/grudge007/gitz](https://www.google.com/search?q=https://github.com/grudge007/gitz). Contributions, bug reports, and feature requests are welcome.

## License

This project is licensed under the **GNU General Public License v3.0 (GPL-3.0)**. See the `LICENSE` file for full details.
