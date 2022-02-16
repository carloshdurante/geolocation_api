package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnrichmentServiceOk(t *testing.T) {
	expected := `{"summary":{"query":"08031450","queryType":"NON_NEAR","queryTime":14,"numResults":1,"offset":0,"totalResults":1,"fuzzyLevel":1},"results":[{"type":"Street","id":"BR/STR/p0/1585322","score":2.3055999279,"address":{"streetName":"Rua Parreira-Brava","municipalitySubdivision":"Vila Curuçá","municipality":"São Paulo","countrySubdivision":"São Paulo","postalCode":"08031","extendedPostalCode":"08031-450","countryCode":"BR","country":"Brasil","countryCodeISO3":"BRA","freeformAddress":"Rua Parreira-Brava, 08031-450, São Paulo","localName":"São Paulo"},"position":{"lat":-23.51206,"lon":-46.42461},"viewport":{"topLeftPoint":{"lat":-23.51075,"lon":-46.42575},"btmRightPoint":{"lat":-23.51353,"lon":-46.42352}}}]}`
	actual, _ := EnrichmentService("08031450")

	assert.Equal(t, expected, actual)
}

func TestEnrichmentServiceError(t *testing.T) {
	expected := `{"summary":{"query":"08031450","queryType":"NON_NEAR","queryTime":14,"numResults":0,"offset":0,"totalResults":0,"fuzzyLevel":1},"results":[]}`
	actual, _ := EnrichmentService("08031450")

	assert.NotEqual(t, expected, actual)
}
