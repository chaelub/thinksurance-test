#!/bin/bash
set -e
dir=$(dirname $0)
cd "${dir}"
cwd=$(pwd)
cd "${cwd}"

echo "Removing all services...";
  docker-compose --project-name small-project down -v --remove-orphans

  echo "Build for MacOS/Windows"
	docker-compose --project-name small-project build

	echo "Recreating all services...";
  docker-compose --project-name small-project up -d

echo "Build complete."
