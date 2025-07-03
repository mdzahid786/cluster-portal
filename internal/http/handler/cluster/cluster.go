package cluster

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/mdzahid786/cluster-portal/internal/storage"
	"github.com/mdzahid786/cluster-portal/internal/types"
	"github.com/mdzahid786/cluster-portal/internal/utils/response"
)

func New(storage storage.Storage)http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		slog.Info("Creating Cluster")
		var cluster types.Cluster
		err := json.NewDecoder(r.Body).Decode(&cluster)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		//w.Write([]byte("Welcome to the cluster api"))
		if err!=nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation
		err = validator.New().Struct(cluster)
		if err !=nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}

		lastId, err := storage.CreateCluster(cluster.Name, cluster.Servers)
		if err!=nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id":lastId})
	}
}

func GetByID(storage storage.Storage)http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		slog.Info("Fetching cluster information")
		id := r.PathValue("id")
		intId, err := strconv.ParseInt(id,10, 64)
		if err!=nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		cluster, err := storage.GetClusterByID(intId)
		if err!=nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, cluster)
	}
}

func GetClusters(storage storage.Storage)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request) {
		slog.Info("Fetching clusters list")
		clusters, err:=storage.GetClusters()
		if err!=nil{
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))	
		}
		response.WriteJson(w, http.StatusOK, clusters)
	}
}

func UpdateCluster(storage storage.Storage)http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Updating cluster")
		if r.Method != http.MethodPut {
			response.WriteJson(w, http.StatusMethodNotAllowed, response.GeneralError(errors.New("invalid method")))
			return
		}

		var cluster types.Cluster
		err := json.NewDecoder(r.Body).Decode(&cluster)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		id := r.PathValue("id")
		intId, err := strconv.ParseInt(id,10, 64)
		// if int64(cluster.Id) != intId {
		// 	response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("cluster id does not match")))
		// 	return
		// }
		
		if err!=nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation
		err = validator.New().Struct(cluster)
		if err !=nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}

		_, err = storage.UpdateCluster(int(intId), cluster.Servers)
		if err!=nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"message":"Updted successfully"})
	}
}