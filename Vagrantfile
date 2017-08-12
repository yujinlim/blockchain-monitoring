Vagrant.configure("2") do |config|
  config.vm.box = "geerlingguy/debian9"
  config.vm.provider "virtualbox" do |v|
    v.memory = 2048
  end

  config.vm.provision "docker"
end
