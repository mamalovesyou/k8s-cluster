# Generating playbooks

### Playbook - initial.yml

This playbook will create a non-root user with sudo privileges on all servers so that you can SSH into them manually as an unprivileged user. This can be useful if, for example, you would like to see system information with commands such as top/htop, view a list of running containers, or change configuration files owned by root. These operations are routinely performed during the maintenance of a cluster, and using a non-root user for such tasks minimizes the risk of modifying or deleting important files or unintentionally performing other dangerous operations.

Here's a breakdown of what this playbook does:

Creates the non-root user ubuntu.
Configures the sudoers file to allow the ubuntu user to run sudo commands without a password prompt.
Adds the public key in your local machine (usually ~/.ssh/id_rsa.pub) to the remote ubuntu user's authorized key list. This will allow you to SSH into each server as the ubuntu user.


### Playbook - kube-dependencies.yml

This playbook will install the operating system level packages required by Kubernetes with Ubuntu's package manager.
The following will be installed:

	* Docker - a container runtime. It is the component that runs your containers.
    * apt-transport-https - allowing you to add external HTTPS sources to your APT sources list.
	* kubeadm- a CLI tool that will install and configure the various components of a cluster in a standard way.
	* kubelet- a system service/program that runs on all nodes and handles node-level operations.
	* kubectl- a CLI tool used for issuing commands to the cluster through its API Server.


### PLaybook - kube-cluster.yml


