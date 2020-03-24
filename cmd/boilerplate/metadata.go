package boilerplate

var MetadataFileTemplate = `problem_id: {{.ProblemId}}`

type MetadataFileParam struct {
	ProblemId string
}
