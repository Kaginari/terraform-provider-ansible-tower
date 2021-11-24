package tower


type CredentialInput struct {
	Fields []Fields `json:"fields"`
}
type Fields struct {
	Id   string 	`json:"id"`
	Type string 	`json:"type"`
	Label string 	`json:"label"`
	Secret bool 	`json:"secret,omitempty"`
	Multiline bool 	`json:"multiline,omitempty"`
	Help string 	`json:"help_text,omitempty"`
	Format string 	`json:"format,omitempty"`
}
func CreateCredentialInputs(f []Fields) CredentialInput{
	var result CredentialInput
	for _, element := range f {
		result.Fields = append(result.Fields,element)
	}
	return result
}