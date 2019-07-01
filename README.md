# Grouter

```go
router := grouter.NewRouter()
router.Post("/job/save", api.HandleJobSave)
router.Delete("/job/delete/@name:([^/]+)", api.HandleJobDel)
router.Get("/job/list", api.HandleJobLists)
router.Post("/job/kill/@name:([^/]+)", api.HandleKillJob)

httpServer := &http.Server{
    Handler:      router,
}

......
```