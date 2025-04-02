package model

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

const (
	ErrExportModelCode                  = "mesheryctl-1127"
	ErrTemplateFileNotPresentCode       = "mesheryctl-1131"
	ErrModelUnsupportedOutputFormatCode = "mesheryctl-1146"
	ErrModelInitCode                    = "mesheryctl-1148" // TODO is it a correct code for this error?
	ErrModelUnsupportedVersionCode      = "mesheryctl-1149" // TODO is it a correct code for this error?

)

func ErrExportModel(err error, name string) error {
	return errors.New(ErrExportModelCode, errors.Fatal, []string{"Error exporting model"}, []string{fmt.Sprintf("Given model with name: %s could not be exported: %s", name, err)}, []string{"Model may not be present in the registry"}, []string{"Ensure that there are no typos in the model name"})
}

func ErrTemplateFileNotPresent() error {
	return errors.New(ErrTemplateFileNotPresentCode, errors.Fatal, []string{"error no template file provided"}, []string{"no template file is provided while using url for importing a model "}, []string{"template file not present"}, []string{"ensure that the template file is present in the given path"})
}

func ErrModelUnsupportedOutputFormat(message string) error {
	return errors.New(ErrModelUnsupportedOutputFormatCode, errors.Fatal, []string{"Error viewing model"}, []string{message}, []string{"Output format not supported"}, []string{"Ensure giving a valid format"})
}

func ErrModelUnsupportedVersion(message string) error {
	return errors.New(ErrModelUnsupportedVersionCode, errors.Fatal, []string{"Error in model version format"}, []string{message}, []string{"Version format not supported"}, []string{"Ensure giving a semver version format"})
}

func ErrModelInitFromString(message string) error {
	return errors.New(ErrModelUnsupportedVersionCode, errors.Fatal, []string{"Error model init"}, []string{message}, []string{"Error during run of model init command"}, []string{"Ensure passing all params according to the command description"})
}

func ErrModelInit(err error) error {
	return ErrModelInitFromString(err.Error())
}
