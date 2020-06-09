package main

import (
	metrics "github.com/googleapis/gnostic/metrics"
)

/*
These variables were made globally because multiple
functions will be accessing and mutating them.
*/

func unpackageVocabulary(v *metrics.Vocabulary) {
	for _, s := range v.Schemas {
		schemas[s.Word] += int(s.Count)
	}
	for _, op := range v.Operations {
		operationID[op.Word] += int(op.Count)
	}
	for _, param := range v.Parameters {
		parameters[param.Word] += int(param.Count)
	}
	for _, prop := range v.Properties {
		properties[prop.Word] += int(prop.Count)
	}
}

func mapIntersection(v *metrics.Vocabulary) {
	schemastemp := make(map[string]int)
	operationIDTemp := make(map[string]int)
	parametersTemp := make(map[string]int)
	propertiesTemp := make(map[string]int)
	for _, s := range v.Schemas {
		value, ok := schemas[s.Word]
		if ok {
			schemastemp[s.Word] += (value + int(s.Count))
		}
	}
	for _, op := range v.Operations {
		value, ok := operationID[op.Word]
		if ok {
			operationIDTemp[op.Word] += (value + int(op.Count))
		}
	}
	for _, param := range v.Parameters {
		value, ok := parameters[param.Word]
		if ok {
			parametersTemp[param.Word] += (value + int(param.Count))
		}
	}
	for _, prop := range v.Properties {
		value, ok := properties[prop.Word]
		if ok {
			propertiesTemp[prop.Word] += (value + int(prop.Count))
		}
	}
	schemas = schemastemp
	operationID = operationIDTemp
	parameters = parametersTemp
	properties = propertiesTemp
}

func vocabularyIntersection(vocabSlices []*metrics.Vocabulary) *metrics.Vocabulary {
	schemas = make(map[string]int)
	operationID = make(map[string]int)
	parameters = make(map[string]int)
	properties = make(map[string]int)

	unpackageVocabulary(vocabSlices[0])
	for i := 1; i < len(vocabSlices); i++ {
		mapIntersection(vocabSlices[i])
	}

	v := &metrics.Vocabulary{
		Properties: FillProtoStructures(properties),
		Schemas:    FillProtoStructures(schemas),
		Operations: FillProtoStructures(operationID),
		Parameters: FillProtoStructures(parameters),
	}
	return v
}
