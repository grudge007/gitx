## gitx

**gitx** is a lightweight local deployment helper written in Go.

It is designed to help engineers manage multi-node project deployments by storing node details, project paths, and execution preferences in a simple local configuration file. Unlike Git, gitx is not about version control. It is about **moving files and executing commands across servers in a controlled, repeatable way**.

gitx focuses on simplicity and transparency. It wraps existing system tools like SSH and rsync instead of hiding them behind heavy abstractions. Every action is explicit, predictable, and easy to debug.

### What gitx aims to solve

* Bootstrapping multi-node deployments with minimal setup
* Keeping deployment configuration local and human-readable
* Running commands remotely in a structured way
* Avoiding complex orchestration frameworks when they are unnecessary

### What gitx is not

* It is not a configuration management system
* It is not a replacement for Ansible, Terraform, or Kubernetes
* It is not a version control system

gitx is intentionally small. It is meant for engineers who want control over their deployment flow without introducing heavy tooling or complex learning curves.

### Philosophy

gitx follows a simple rule:

> Wrap the system. Don’t hide it.

The tool prioritizes clarity over abstraction, structure over magic, and reliability over convenience. If something fails, gitx shows exactly what failed and why.

### Target users

* Platform and infrastructure engineers
* Developers managing small to medium multi-node deployments
* Anyone who prefers simple, auditable tooling

gitx is built to stay small, understandable, and practical.
