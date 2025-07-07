# Cluster Portal

A full-stack application to manage and monitor server clusters. Built with **GoLang** for the backend and **React.js** for the frontend, with **Mysql** the primary database.

## Backend::

To insert DB into the database of mysql run below command from root directory

sh ./cmd/cluster-portal/run.sh


Go to Root Folder and execute below command to run go server or backend server

go mod tidy

go run ./cmd/cluster-portal/main.go --config ./config/local.yaml


## Frontend::

To start frontend server go to frontend directory and run below command

npm install

npm run dev



