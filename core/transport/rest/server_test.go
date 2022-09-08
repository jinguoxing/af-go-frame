package rest

import (
    "context"
    "testing"
    "time"
)

func TestServer(t *testing.T){

    ctx := context.Background()
    srv := NewServer()
    //srv.Handle("/index", newHandleFuncWrapper(h))
    //srv.HandleFunc("/index/{id:[0-9]+}", h)
    //srv.HandlePrefix("/test/prefix", newHandleFuncWrapper(h))
    //srv.HandleHeader("content-type", "application/grpc-web+json", func(w http.ResponseWriter, r *http.Request) {
    //    _ = json.NewEncoder(w).Encode(testData{Path: r.RequestURI})
    //})
    //srv.Route("/errors").GET("/cause", func(ctx Context) error {
    //    return errors.BadRequest("xxx", "zzz").
    //        WithMetadata(map[string]string{"foo": "bar"}).
    //        WithCause(fmt.Errorf("error cause"))
    //})
    //
    //if e, err := srv.Endpoint(); err != nil || e == nil || strings.HasSuffix(e.Host, ":0") {
    //    t.Fatal(e, err)
    //}


    go func() {
        if err := srv.Start(ctx); err != nil {
            panic(err)
        }
    }()
    time.Sleep(time.Second)

    time.Sleep(time.Second)
    if srv.Stop(ctx) != nil {
        t.Errorf("expected nil got %v", srv.Stop(ctx))
    }

}
