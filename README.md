# gitz

**gitz** is a lightweight local deployment helper written in Go.

It is designed to help engineers manage multi-node project deployments by storing node details, project paths, and execution preferences in a simple local configuration file. Unlike Git, gitz is not about version control. It is about **moving files and executing commands across servers in a controlled, repeatable way**.

### 🚀 Key Features

* **Initialization:** Quickly bootstrap a project with a `.gitz/gitz.conf` file.
* **File Deployment (Push):** Securely transfer files to multiple remote nodes via SFTP.
* **Remote Execution (Run):** Execute shell commands across your entire inventory in parallel via SSH.

### 🛠 How to Build from Source

Since **gitz** is written in Go, you can compile it into a single binary for your system.

1. **Clone the repository:**
```bash
git clone https://github.com/yourusername/gitz.git
cd gitz

```


2. **Download dependencies:**
```bash
go mod tidy

```


3. **Build the binary:**
```bash
go build -o gitz main.go

```


4. *(Optional)* **Move to your path:**
```bash
sudo mv gitz /usr/local/bin/

```



### 📖 Quick Start Guide

1. **Initialize your project:**
```bash
gitz init

```


2. **Configure your nodes:**
Open `.gitz/gitz.conf` and add your server IPs, usernames, and passwords.
3. **Deploy a file:**
Place a file in your directory and run:
```bash
gitz push

```


4. **Run a command:**
```bash
gitz run "uptime"

```



### What gitz aims to solve

* Bootstrapping multi-node deployments with minimal setup.
* Keeping deployment configuration local and human-readable.
* Running commands remotely in a structured way.
* Avoiding complex orchestration frameworks when they are unnecessary.

### What gitz is not

* It is **not** a configuration management system (like Ansible).
* It is **not** a replacement for Terraform or Kubernetes.
* It is **not** a version control system.

### Philosophy

> Wrap the system. Don’t hide it.

The tool prioritizes clarity over abstraction, structure over magic, and reliability over convenience. If something fails, gitz shows exactly what failed and why.
