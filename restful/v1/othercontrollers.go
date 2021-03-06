package v1

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/SmartMeshFoundation/SmartRaiden/log"
	"github.com/SmartMeshFoundation/SmartRaiden/network"
	"github.com/ant0ine/go-json-rest/rest"
)

/*
UpdateMeshNetworkNodes update nodes of this intranet
*/
func UpdateMeshNetworkNodes(w rest.ResponseWriter, r *rest.Request) {
	var nodes []*network.NodeInfo
	err := r.DecodeJsonPayload(&nodes)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = RaidenAPI.Raiden.Protocol.UpdateMeshNetworkNodes(nodes)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.(http.ResponseWriter).Write([]byte("ok"))
	if err != nil {
		log.Warn(fmt.Sprintf("writejson err %s", err))
	}
}

/*
SwitchNetwork  switch between mesh and internet
*/
func SwitchNetwork(w rest.ResponseWriter, r *rest.Request) {
	mesh := r.PathParam("mesh")
	isMesh, err := strconv.ParseBool(mesh)
	if err != nil {
		rest.Error(w, "arg error", http.StatusBadRequest)
		return
	}
	RaidenAPI.Raiden.Config.IsMeshNetwork = isMesh
	_, err = w.(http.ResponseWriter).Write([]byte("ok"))
	if err != nil {
		log.Warn(fmt.Sprintf("writejson err %s", err))
	}
}
