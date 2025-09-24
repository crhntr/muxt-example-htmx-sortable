package hypertext

//go:generate go run github.com/typelate/muxt/cmd/muxt generate --receiver-type-package=github.com/typelate/example-sortable/internal/domain --receiver-type=Service --routes-func Routes
//go:generate rm -rf internal/fake
//go:generate mkdir -p internal/fake
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o internal/fake/routes_receiver.go --fake-name RoutesReceiver . RoutesReceiver
