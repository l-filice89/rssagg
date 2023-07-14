package main

import (
	"github.com/go-chi/chi"
	"github.com/l-filice89/rssagg/internal/database"
	"net/http"
	"strconv"
)

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := chi.URLParam(r, "limit")
	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid limit")
		return
	}

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting posts")
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
