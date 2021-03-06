package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pmmp/CrashArchive/app/crashreport"
	"github.com/pmmp/CrashArchive/app/template"
	"log"
)

func DownloadGet(w http.ResponseWriter, r *http.Request) {
	reportID, err := strconv.Atoi(chi.URLParam(r, "reportID"))
	if err != nil {
		log.Println(err)
		template.ErrorTemplate(w, "", http.StatusBadRequest)
		return
	}

	_, jsonData, err := crashreport.ReadFile(int64(reportID))
	if err != nil {
		template.ErrorTemplate(w, "Report not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%d.log", reportID))
	w.Header().Set("Content-Length", strconv.Itoa(len(jsonData["report"].(string))))
	w.Write([]byte(jsonData["report"].(string)))
}
