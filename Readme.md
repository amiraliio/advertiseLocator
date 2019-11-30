# Advertise Locator

### Config Files

    docker.env
    config.yaml
    nginx.conf

### Run Application

    - docker-compose build --no-cache
    - docker rmi -f $(docker images -f "dangling=true" -q)
    - docker-compose up -d --force-recreate

### Stop Application

    - docker-compose down


