Backend:: <br>
To insert DB into the database of mysql run below command from root directory
sh ./cmd/cluster-portal/run.sh

Go to Root Folder and execute below command to run go server or backend server
go run ./cmd/cluster-portal/main.go --config ./config/local.yaml

Frontend::
To start frontend server go to frontend directory and run below command
npm run dev
