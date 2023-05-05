## jmeter installation

## run jmeter gui on EC2
```cmd
brew install xquartz
reboot
export DISPLAY=:0
ssh -Y -i "key.pem"  ec2-user@ip
./apache-jmeter-5.5/bin/jmeter
```
