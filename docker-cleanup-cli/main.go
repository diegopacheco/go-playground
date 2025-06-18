package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type DockerImage struct {
	Repository string
	Tag        string
	ImageID    string
	Created    string
	Size       string
	SizeBytes  int64
}

func parseDockerImagesOutput(output string) []DockerImage {
	scanner := bufio.NewScanner(strings.NewReader(output))
	images := []DockerImage{}
	first := true

	for scanner.Scan() {
		line := scanner.Text()
		if first {
			first = false
			continue // skip header
		}
		fields := strings.Fields(line)
		if len(fields) < 7 {
			continue
		}
		created := strings.Join(fields[3:6], " ")
		image := DockerImage{
			Repository: fields[0],
			Tag:        fields[1],
			ImageID:    fields[2],
			Created:    created,
			Size:       fields[len(fields)-1],
			SizeBytes:  parseSize(fields[len(fields)-1]),
		}
		images = append(images, image)
	}

	sort.Slice(images, func(i, j int) bool {
		return images[i].SizeBytes > images[j].SizeBytes
	})
	return images
}

func parseSize(sizeStr string) int64 {
	unit := sizeStr[len(sizeStr)-2:]
	valueStr := sizeStr[:len(sizeStr)-2]
	value, err := strconv.ParseFloat(strings.TrimSpace(valueStr), 64)
	if err != nil {
		return 0
	}
	switch unit {
	case "GB":
		return int64(value * 1e9)
	case "MB":
		return int64(value * 1e6)
	case "KB":
		return int64(value * 1e3)
	default:
		return 0
	}
}

func main() {
	cmd := exec.Command("docker", "images")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run docker images: %v\n", err)
		os.Exit(1)
	}

	images := parseDockerImagesOutput(out.String())
	filteredImages := images

	app := tview.NewApplication()
	inputField := tview.NewInputField().
		SetLabel("Filter: ").
		SetFieldWidth(50)

	list := tview.NewList()

	inputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyDown || event.Key() == tcell.KeyUp {
			app.SetFocus(list)
			return event
		}
		return event
	})

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlF {
			app.SetFocus(inputField)
			return nil
		}
		return event
	})

	updateList := func(filter string) {
		list.Clear()
		filteredImages = []DockerImage{}

		for _, img := range images {
			if filter == "" || strings.Contains(strings.ToLower(img.Repository), strings.ToLower(filter)) ||
				strings.Contains(strings.ToLower(img.ImageID), strings.ToLower(filter)) {

				repo := img.Repository
				if len(repo) > 35 {
					repo = repo[:32] + "..."
				}

				imageID := img.ImageID
				if len(imageID) > 12 {
					imageID = imageID[:12]
				}

				created := img.Created
				if len(created) > 15 {
					created = created[:12] + "..."
				}

				display := fmt.Sprintf("%-8s | %-35s | %-12s | %-15s", img.Size, repo, imageID, created)
				list.AddItem(display, "", 0, nil)
				filteredImages = append(filteredImages, img)
			}
		}
	}
	updateList("")

	inputField.SetChangedFunc(func(text string) {
		updateList(text)
	})

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(inputField, 1, 0, true).
		AddItem(list, 0, 1, true)

	list.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		if index >= 0 && index < len(filteredImages) {
			imageID := filteredImages[index].ImageID
			repo := filteredImages[index].Repository
			cmd := exec.Command("docker", "rmi", imageID)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				errorModal := tview.NewModal().
					SetText(fmt.Sprintf("Failed to delete image %s: %v", imageID, err)).
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						app.SetRoot(flex, true)
					})
				app.SetRoot(errorModal, false)
			} else {
				for i, img := range images {
					if img.ImageID == imageID {
						images = append(images[:i], images[i+1:]...)
						break
					}
				}

				successModal := tview.NewModal().
					SetText(fmt.Sprintf("Image %s %s deleted", repo, imageID[:12])).
					SetBackgroundColor(tcell.ColorGreen)

				app.SetRoot(successModal, false)

				go func() {
					time.Sleep(2 * time.Second)
					app.QueueUpdateDraw(func() {
						app.SetRoot(flex, true)
						updateList(inputField.GetText())
					})
				}()
			}
		}
	})

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			app.Stop()
			return nil
		}
		if event.Key() == tcell.KeyTab {
			if app.GetFocus() == inputField {
				app.SetFocus(list)
			} else {
				app.SetFocus(inputField)
			}
			return nil
		}
		return event
	})

	app.SetFocus(inputField)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
