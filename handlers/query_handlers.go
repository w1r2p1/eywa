package handlers

import (
	"github.com/vivowares/eywa/Godeps/_workspace/src/github.com/zenazn/goji/web"
	. "github.com/vivowares/eywa/models"
	. "github.com/vivowares/eywa/utils"
	"net/http"
	"os"
)

func QueryValue(c web.C, w http.ResponseWriter, r *http.Request) {
	ch, found := findCachedChannel(c, "id")
	if !found {
		Render.JSON(w, http.StatusNotFound, map[string]string{"error": "channel not found"})
		return
	}

	q := &ValueQuery{Channel: ch}
	err := q.Parse(QueryToMap(r.URL.Query()))
	if err != nil {
		Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	} else {
		value, err := q.QueryES()
		if err != nil {
			Render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		} else {
			Render.JSON(w, http.StatusOK, value)
		}
	}
}

func QuerySeries(c web.C, w http.ResponseWriter, r *http.Request) {
	ch, found := findCachedChannel(c, "id")
	if !found {
		Render.JSON(w, http.StatusNotFound, map[string]string{"error": "channel not found"})
		return
	}

	q := &SeriesQuery{Channel: ch}
	err := q.Parse(QueryToMap(r.URL.Query()))
	if err != nil {
		Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	} else {
		series, err := q.QueryES()
		if err != nil {
			Render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		} else {
			Render.JSON(w, http.StatusOK, series)
		}
	}
}

func QueryRaw(c web.C, w http.ResponseWriter, r *http.Request) {
	ch, found := findCachedChannel(c, "id")
	if !found {
		Render.JSON(w, http.StatusNotFound, map[string]string{"error": "channel not found"})
		return
	}

	q := &RawQuery{Channel: ch}
	err := q.Parse(QueryToMap(r.URL.Query()))
	if err != nil {
		Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	} else {
		res, err := q.QueryES()
		if err != nil {
			Render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		} else {
			if f, ok := res.(map[string]interface{})["file"]; ok {
				http.ServeFile(w, r, f.(string))
				os.Remove(f.(string))
			} else {
				Render.JSON(w, http.StatusOK, res)
			}
		}
	}
}
