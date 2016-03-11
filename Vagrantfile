# -*- mode: ruby -*-
# vi: set ft=ruby :

GO_ARCHIVE = "go1.6.linux-amd64.tar.gz"

Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    sudo apt-get update
    sudo apt-get install -y mercurial git

    echo "Downloading #{GO_ARCHIVE}..."
    wget --quiet https://storage.googleapis.com/golang/#{GO_ARCHIVE}
    echo "Done!"

    echo "Unpacking #{GO_ARCHIVE}"
    sudo tar -C /usr/local -xzf #{GO_ARCHIVE}
    echo "Done!"

    echo "Setting up environment..."

    mkdir /home/vagrant/go
    mkdir /home/vagrant/go/src
    mkdir /home/vagrant/go/pkg
    mkdir /home/vagrant/go/bin

    mkdir -p /home/vagrant/go/src/github.com/applikatoni
    ln -s /vagrant /home/vagrant/go/src/github.com/applikatoni/toni

    echo "Directory structure created."

    echo 'export GOPATH=/home/vagrant/go' >> /home/vagrant/.profile
    echo 'export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin' >> /home/vagrant/.profile

    source /home/vagrant/.profile

    echo "Environment variables set up."

    echo "Installing toni main repo..."
    cd /home/vagrant/go/src/github.com/applikatoni/toni
    go get ./...
    go install ./...
    echo "Installed."

    echo "Running tests..."
    go test -v ./...
  SHELL
end
