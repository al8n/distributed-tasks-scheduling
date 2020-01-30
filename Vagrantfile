Vagrant.require_version ">= 1.6.0"

boxes = [
    {
        :name => "alpine-linux",
        :eth1 => "192.168.10.10",
        :mem => "1024",
        :cpu => "2"
    },
]

Vagrant.configure(2) do |config|
  config.vm.box = "generic/alpine310"

  boxes.each do |opts|
      config.vm.define opts[:name] do |config|
        config.vm.hostname = opts[:name]
        config.ssh.username = "vagrant"
        config.ssh.password = "vagrant"
        config.vm.provider "vmware_fusion" do |v|
          v.vmx["memsize"] = opts[:mem]
          v.vmx["numvcpus"] = opts[:cpu]
        end
        config.vbguest.auto_update = false
        config.vm.provider "virtualbox" do |v|
          v.customize ["modifyvm", :id, "--memory", opts[:mem]]
          v.customize ["modifyvm", :id, "--cpus", opts[:cpu]]
        end
        config.vm.network "private_network", ip: "192.168.10.11"
        config.vm.network "private_network", ip: "192.168.10.12"
        config.vm.network "public_network", ip: opts[:eth1], bridge: "en1: Wi-Fi (AirPort)"
      end
  end

  config.vm.synced_folder "./share", "/home/vagrant/share"
  config.vm.provision "shell", privileged: true, path: "./setup.sh"
end
