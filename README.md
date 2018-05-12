# Go Http Docker

## Getting Project and Docker
- git clone https://github.com/Jagalite/GoHTTP.git
- sudo apt-get install docker.io

## AWS - Allow HTTP
- Go to the "Network & Security" -> Security Group settings in the left hand navigation
- Find the Security Group that your instance is apart of
- Click on Inbound Rules
- Use the drop down and add HTTP (port 80)
- Click Apply

## Running Docker
- sudo bash BuildAndRun.bash

## Hitting the URL
- http://<EC2_INSTANCE>:80

