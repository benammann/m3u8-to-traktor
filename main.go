package main

import (
	"fmt"
	"github.com/benammann/m3u8-to-traktor/converter"
	"github.com/sqweek/dialog"
	"os"
)

const (
	DIALOG_TITLE = "M3U8 to Traktor"
)

func main() {

	startConverter := dialog.Message("Would you like to convert .m3u8 files to .m3u ?").Title(DIALOG_TITLE).YesNo()

	if startConverter {

		converterClient := converter.NewConverter()

		enteredFileSelectDialog := false
		for {
			if enteredFileSelectDialog {
				selectMore := dialog.Message("File Added. Would you like to add another .m3u8 file ?").Title(DIALOG_TITLE).YesNo()
				if !selectMore {
					break
				}
			}

			filename, err := dialog.File().Filter("m3u8 playlist file", "m3u8").Load()

			if err == nil {

				err = converterClient.AddInputFile(filename)

				if err != nil {
					dialog.Message(err.Error()).Error()
				}

			} else {
				if err == dialog.Cancelled && !enteredFileSelectDialog {
					dialog.Message("No file selected. Bye :)").Info()
					os.Exit(0)
					return
				}

				dialog.Message(fmt.Sprintf("Error while selecting file: %s", err.Error())).Error()
			}

			enteredFileSelectDialog = true
		}

		dialog.Message("Please select the output directory").Title(DIALOG_TITLE).Info()

		outDir, err := dialog.Directory().Title("Output Directory").Browse()

		if err == nil {

			err = converterClient.SetOutputDirectory(outDir)

			if err != nil {
				dialog.Message(err.Error()).Error()
				os.Exit(1)
			}

			errWhileConverting := converterClient.Convert()

			if errWhileConverting != nil {
				dialog.Message(err.Error()).Error()
				os.Exit(1)
			} else {
				dialog.Message("Your files have been converted! Bye :)").Info()
				os.Exit(0)
			}

		} else {

			if err != dialog.Cancelled {
				dialog.Message(fmt.Sprintf("Error while selecting output directory: %s", err.Error())).Error()
			} else {
				dialog.Message("Operation Cancelled, leaving. Bye :)").Info()
			}

		}


	}

}