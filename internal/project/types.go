package project

type Type string

const (
	TypeUnknown Type = "Unknown"
	TypeNodeJS  Type = "Node.js"
	TypeGo      Type = "Go"
	TypeReact   Type = "React"
	TypeNextJS  Type = "Next.js"
	TypeExpress Type = "Express"
)

type Info struct {
	Type           Type
	PackageManager string
	Files          []string
}
