# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|

  config.vm.box = "generic/ubuntu2204"

  if Vagrant.has_plugin?("vagrant-vbguest")
    config.vbguest.auto_update = false
    config.vbguest.no_install = true
  end

  config.vm.define "ceph-controller" do |ceph_ctrl|
    ceph_ctrl.vm.network "private_network", ip: "192.168.19.10"
    ceph_ctrl.vm.hostname = "ceph-controller"
  
    ceph_ctrl.vm.provider "virtualbox" do |vb|
      vb.memory = "2048"
      vb.cpus = "2"
      vb.name = "ceph-controller"
    end
  
    ceph_ctrl.vm.provision "shell", inline: <<-SHELL
      apt update
      apt install python3-pip -y
      python3 -m pip install --user ansible
    SHELL
  end

  config.vm.define "ceph-1" do |ceph_1|
    ceph_1.vm.network "private_network", ip: "192.168.19.11"
    ceph_1.vm.hostname = "ceph-1"
  
    ceph_1.vm.provider "virtualbox" do |vb|
      vb.memory = "2048"
      vb.cpus = "2"
      vb.name = "ceph-1"
    end
  end

end
