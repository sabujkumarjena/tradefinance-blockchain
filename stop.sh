docker-compose -f docker-compose.yml down

docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker rmi $(docker images| grep dev | awk '{print $3}')
echo y| docker volume prune
echo y| docker network prune
docker ps -a



