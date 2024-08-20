package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Function to calculate the difference between two dates
func dateDifference(start, end time.Time) (years, months, days int) {
	if start.After(end) {
		start, end = end, start
	}

	// Calculate the year difference
	years = end.Year() - start.Year()

	// Calculate the month difference
	months = int(end.Month() - start.Month())
	if months < 0 {
		years--
		months += 12
	}

	// Calculate the day difference
	days = end.Day() - start.Day()
	if days < 0 {
		months--
		prevMonth := end.AddDate(0, -1, 0)
		days += time.Date(prevMonth.Year(), prevMonth.Month(), 0, 0, 0, 0, 0, prevMonth.Location()).Day()
	}

	return
}

func main() {
	// Create a new application
	a := app.New()
	w := a.NewWindow("Kalkulator wieku")
	w.Resize(fyne.NewSize(400, 200))

	// Create date entry widgets
	startDateEntry := widget.NewEntry()
	startDateEntry.SetPlaceHolder("RRRR-MM-DD")
	endDateEntry := widget.NewEntry()
	endDateEntry.SetPlaceHolder("RRRR-MM-DD")

	// Label to display the result
	resultLabel := widget.NewLabel("Różnica: ")

	// Button to calculate the difference
	calculateBtn := widget.NewButtonWithIcon("Oblicz", theme.ConfirmIcon(), func() {
		startDate, err1 := time.Parse("2006-01-02", startDateEntry.Text)
		endDate, err2 := time.Parse("2006-01-02", endDateEntry.Text)

		if err1 != nil || err2 != nil {
			dialog.ShowError(fmt.Errorf("Wprowadź daty w formacie: RRRR-MM-DD"), w)
			return
		}

		years, months, days := dateDifference(startDate, endDate)
		labelYears := ""
		labelMonths := ""
		labelDays := ""
		if years == 0 {
			labelYears = "lat"
		} else if years == 1 {
			labelYears = "rok"
		} else {
			labelYears = "lat"
		}

		if months == 1 {
			labelMonths = "miesiąc"
		} else if months > 1 && months < 5 {
			labelMonths = "miesiące"
		} else {
			labelMonths = "miesięcy"
		}

		if days == 1 {
			labelDays = "dzień"
		} else {
			labelDays = "dni"
		}
		resultLabel.SetText(fmt.Sprintf("Różnica: %d %s, %d %s i %d %s.", years, labelYears, months, labelMonths, days, labelDays))
	})

	// Layout the widgets in a form
	form := container.NewVBox(
		widget.NewForm(
			widget.NewFormItem("Data urodzenia", startDateEntry),
			widget.NewFormItem("Data obliczenia", endDateEntry),
		),
		calculateBtn,
		resultLabel,
	)

	w.SetContent(form)

	// Show the window and run the app
	w.ShowAndRun()
}
