package factorio

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/OpenFactorioServerManager/factorio-server-manager/bootstrap"
)

type Save struct {
	Name    string    `json:"name"`
	LastMod time.Time `json:"last_mod"`
	Size    int64     `json:"size"`
}

func (s Save) String() string {
	return s.Name
}

// Lists save files in factorio/saves
func ListSaves(saveDir string) (saves []Save, err error) {
	saves = []Save{}
	err = filepath.Walk(saveDir, func(path string, info os.FileInfo, err error) error {
		if info == nil || (info.IsDir() && info.Name() == "saves") {
			return nil
		}
		saves = append(saves, Save{
			info.Name(),
			info.ModTime(),
			info.Size(),
		})
		return nil
	})
	return
}

func FindSave(name string) (*Save, error) {
	config := bootstrap.GetConfig()
	saves, err := ListSaves(config.FactorioSavesDir)
	if err != nil {
		return nil, fmt.Errorf("error listing saves: %v", err)
	}

	for _, save := range saves {
		if save.Name == name {
			return &save, nil
		}
	}

	return nil, errors.New("save not found")
}

func (s *Save) Remove() error {
	if s.Name == "" {
		return errors.New("save name cannot be blank")
	}
	config := bootstrap.GetConfig()
	return os.Remove(filepath.Join(config.FactorioSavesDir, s.Name))
}

// Create savefiles for Factorio
func CreateSave(filePath string, mapSettingsGenFilePath string, mapSettingsFilePath string) (string, error) {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		log.Printf("Error in creating Factorio save: %s", err)
		return "", err
	}

	args := []string{
		"--map-gen-settings",
		mapSettingsGenFilePath,
		"--map-settings",
		mapSettingsFilePath,
		"--create",
		filePath,
	}
	config := bootstrap.GetConfig()
	cmdOutput, err := exec.Command(config.FactorioBinary, args...).Output()
	if err != nil {
		log.Printf("Error in creating Factorio save: %s", err)
		log.Println(string(cmdOutput))
		return "", err
	}

	result := string(cmdOutput)

	return result, nil
}

func GenerateMapPreview(mapSettingsGenFilePath string, mapSettingsFilePath string) (string, error) {

	if _, err := os.Stat("./map_previews"); err != nil {
		err = os.Mkdir("./map_previews", 664)
		if err != nil {
			log.Printf("Failed to create folder %s with permission: %s\n", "./map_previews", err)
			return "", err
		}
	}

	previewFolder, _ := filepath.Abs("./map_previews")

	// adding slash will indicate for factorio it is a path and not an image file
	// source: https://wiki.factorio.com/Command_line_parameters
	previewFolder += string(filepath.Separator)

	args := []string{
		"--map-gen-settings",
		mapSettingsGenFilePath,
		"--generate-map-preview",
		previewFolder,
		"--map-settings",
		mapSettingsFilePath,
	}

	config := bootstrap.GetConfig()
	log.Println(filepath.Abs(config.FactorioBinary))
	cmdOutput, err := exec.Command(config.FactorioBinary, args...).Output()
	if err != nil {
		log.Printf("Error in creating Factorio save: %s", err)
		return "", err
	}

	result := string(cmdOutput)

	// extract image path from result
	startMark := "Wrote map preview image file: "
	endMark := ".png"

	startIndex := strings.Index(result, startMark)
	if startIndex == -1 {
		return "", errors.New("failed to detect start of the image path")
	}
	startIndex += len(startMark)
	endIndex := strings.Index(result[startIndex:], endMark)

	if endIndex == -1 {
		return "", errors.New("failed to detect end of the image path")
	}
	endIndex += len(endMark)
	imagePath := result[startIndex:][:endIndex]

	return imagePath, nil
}
