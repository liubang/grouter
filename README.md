# Grouter

![License: MIT](https://img.shields.io/github/license/liubang/grouter?style=flat-square)

```go
router := grouter.NewRouter()
router.Post("/job/save", api.HandleJobSave)
router.Delete("/job/delete/@name:([^/]+)", api.HandleJobDel)
router.Get("/job/list", api.HandleJobLists)
router.Post("/job/kill/@name:([^/]+)", api.HandleKillJob)

http.ListenAndServe("127.0.0.1:8080", router)
......
```

