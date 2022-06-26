# README

This example lets the user take two tests conisting of a few simple math questions. If the user answers all the questions in a particular test correctly, then the Wails events system (runtime.EventsEmit and runtime.EventsOn) is used to display an animated gif.

If you are running this example from a Windows computer and/or are using anti-virus software, you might need to run some extra steps instead of simply running `wails dev`. These extra steps are as follows:

While in the frontend directory

	`npm install`

	`npm run build`
	
	remove the two .gif files from the frontend/dist/assets directory

While in the project directory

	`go build -tags dev -gcflags "all=-N -l"`

	`wails dev`
	
