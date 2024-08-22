package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getAllWatchLists(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var allWatchLists []*WatchList
	encoder := json.NewEncoder(res)
	for index := range watchLists {
		allWatchLists = append(allWatchLists, watchLists[index])
	}
	encoder.Encode(allWatchLists)
}

func getWatchList(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")

	if watchLists[id] == nil {
		respondWithNotFound(res)
		return
	}
	writeJSONResponse(res, watchLists[id], http.StatusOK)
}

func createWatchList(res http.ResponseWriter, req *http.Request) {
	watchList := parseRequestAndValidate(res, req)
	if watchList == nil {
		return
	}

	watchList.ID = strconv.Itoa(watchListIdCounter)
	watchListIdCounter += 1
	watchLists[watchList.ID] = watchList

	writeJSONResponse(res, watchList, http.StatusCreated)
}

func setWatchList(res http.ResponseWriter, req *http.Request) {

	watchList := parseRequestAndValidate(res, req)
	if watchList == nil {
		return
	}

	id := req.PathValue("id")
	watchList.ID = id
	if watchLists[id] == nil {
		writeJSONResponse(res, watchList, http.StatusCreated)
	} else {
		writeJSONResponse(res, watchList, http.StatusOK)
	}
	watchLists[id] = watchList
}

func patchWatchList(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	watchList := parseRequestAndValidate(res, req)
	watchList.ID = id
	currentWatchList := watchLists[id]
	if watchList.Name != nil {
		currentWatchList.Name = watchList.Name
	}
	watchLists[id] = currentWatchList
	writeJSONResponse(res, currentWatchList, http.StatusOK)
}

func deleteWatchList(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	currentWatchList := watchLists[id]
	delete(watchLists, id)
	writeJSONResponse(res, currentWatchList, http.StatusOK)
}
