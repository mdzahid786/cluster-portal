<b>Backend::</b> <br><br>
To insert DB into the database of mysql run below command from root directory <br>
sh ./cmd/cluster-portal/run.sh<br>
<br>
Go to Root Folder and execute below command to run go server or backend server<br>
go run ./cmd/cluster-portal/main.go --config ./config/local.yaml<br>
<br><br>
<b>Frontend::</b> <br><br>
To start frontend server go to frontend directory and run below command<br>
npm install<br>
npm run dev<br>
