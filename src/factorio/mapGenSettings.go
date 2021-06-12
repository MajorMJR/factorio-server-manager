package factorio

type MapResource struct {
	Frequency float32 `json:"frequency"`
	Size      float32 `json:"size"`
	Richness  float32 `json:"richness"`
}

type AutoPlaceControls struct {
	Coal       MapResource `json:"coal"`
	Stone      MapResource `json:"stone"`
	CopperOre  MapResource `json:"copper-ore"`
	IronOre    MapResource `json:"iron-ore"`
	UraniumOre MapResource `json:"uranium-ore"`
	CrudeOil   MapResource `json:"crude-oil"`
	Trees      MapResource `json:"trees"`
	EnemyBase  MapResource `json:"enemy-base"`
}

type CliffSettings struct {
	Name                   string `json:"name"`
	CliffElevation0        int    `json:"cliff_elevation_0"`
	CliffElevationInterval int    `json:"cliff_elevation_interval"`
	Richness               int    `json:"richness"`
}

type PropertyExpressionNames struct {
	Elevation                   string `json:"elevation"`
	MoistureFrequencyMultiplier string `json:"control-setting:moisture:frequency:multiplier"`
	MoistureBias                string `json:"control-setting:moisture:bias"`
	AuxFrequencyMultiplier      string `json:"control-setting:aux:frequency:multiplier"`
	AuxBias                     string `json:"control-setting:aux:bias"`
}

type StartingPoints []struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MapGenSettings struct {
	TerrainSegmentation int               `json:"terrain_segmentation"`
	Water               int               `json:"water"`
	Width               int               `json:"width"`
	Height              int               `json:"height"`
	StartingArea        float32           `json:"starting_area"`
	PeacefulMode        bool              `json:"peaceful_mode"`
	AutoPlaceControls   AutoPlaceControls `json:"autoplace_controls"`

	CliffSettings CliffSettings `json:"cliff_settings"`

	PropertyExpressionNames PropertyExpressionNames `json:"property_expression_names"`

	StartingPoints StartingPoints `json:"starting_points"`

	Seed *int `json:"seed"`
}

func DefaultMapGenSettings() MapGenSettings {
	return MapGenSettings{
		TerrainSegmentation: 1,
		Water:               1,
		Width:               0,
		Height:              0,
		StartingArea:        1,
		PeacefulMode:        false,
		AutoPlaceControls: AutoPlaceControls{
			Coal: MapResource{
				Frequency: 1,
				Size:      1,
				Richness:  1,
			},
			Stone: MapResource{
				Frequency: 1,
				Size:      1,
				Richness:  1,
			},
			CopperOre: MapResource{
				Frequency: 1,
				Size:      1,
				Richness:  1,
			},
			IronOre: MapResource{
				Frequency: 1,
				Size:      1,
				Richness:  1,
			},
			UraniumOre: MapResource{
				Frequency: 1,
				Size:      1,
				Richness:  1,
			},
			CrudeOil: MapResource{
				Frequency: 1,
				Size:      1,
				Richness:  1,
			},
			Trees: MapResource{
				Frequency: 1,
				Size:      1,
				Richness:  1,
			},
			EnemyBase: MapResource{
				Frequency: 1,
				Size:      1,
				Richness:  1,
			},
		},
		CliffSettings: CliffSettings{
			Name:                   string("cliff"),
			CliffElevation0:        10,
			CliffElevationInterval: 40,
			Richness:               1,
		},
		PropertyExpressionNames: PropertyExpressionNames{
			MoistureFrequencyMultiplier: "1",
			MoistureBias:                "0",
			AuxFrequencyMultiplier:      "1",
			AuxBias:                     "0",
		},
		StartingPoints: StartingPoints{
			{
				X: 0,
				Y: 0,
			},
		},
		Seed: nil,
	}
}
