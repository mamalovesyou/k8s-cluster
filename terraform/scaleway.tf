provider "scaleway" {
  organization = "${var.organisation_id}"
  token        = "${var.secret_token}"
  region       = "${var.region}"
}

resource "scaleway_ip" "ip_test" {
  count = 3
}

data "scaleway_image" "ubuntu" {
  architecture = "x86_64"
  name         = "Ubuntu Mini Xenial 25G"
}

resource "scaleway_server" "load_balancer" {
  name           = "load_balancer"
  tags           = ["load-balancer"]
  image          = "${data.scaleway_image.ubuntu.id}"
  type           = "START1-XS"
  state          = "running"
  security_group = "${scaleway_security_group.http.id}"
  public_ip      = "${scaleway_ip.ip_test.0.ip}"
}
resource "scaleway_server" "node_1" {
  name           = "node_1"
  tags           = ["kube-master"]
  image          = "${data.scaleway_image.ubuntu.id}"
  type           = "START1-XS"
  state          = "running"
  security_group = "${scaleway_security_group.http.id}"
  public_ip      = "${scaleway_ip.ip_test.1.ip}"
}

resource "scaleway_server" "node_2" {
  name           = "node-2"
  tags           = ["kube-node"]
  image          = "${data.scaleway_image.ubuntu.id}"
  type           = "START1-XS"
  state          = "running"
  security_group = "${scaleway_security_group.http.id}"
  public_ip      = "${scaleway_ip.ip_test.2.ip}"
}

resource "scaleway_security_group" "http" {
  name                    = "http"
  description             = "allow HTTP and HTTPS traffic"
  enable_default_security = true
}

resource "scaleway_security_group_rule" "http_accept" {
  security_group = "${scaleway_security_group.http.id}"

  action    = "accept"
  direction = "inbound"
  ip_range  = "0.0.0.0/0"
  protocol  = "TCP"
  port      = 80
}

resource "scaleway_security_group_rule" "https_accept" {
  security_group = "${scaleway_security_group.http.id}"

  action    = "accept"
  direction = "inbound"
  ip_range  = "0.0.0.0/0"
  protocol  = "TCP"
  port      = 443
}

output "load_balancer_public_ip" {
  value = "${scaleway_server.load_balancer.public_ip}"
}
output "load_balancer_private_ip" {
  value = "${scaleway_server.load_balancer.private_ip}"
}
output "node_1_public_ip" {
  value = "${scaleway_server.node_1.public_ip}"
}
output "node_1_private_ip" {
  value = "${scaleway_server.node_1.private_ip}"
}
output "node_2_public_ip" {
  value = "${scaleway_server.node_2.public_ip}"
}
output "node_2_private_ip" {
  value = "${scaleway_server.node_2.private_ip}"
}
