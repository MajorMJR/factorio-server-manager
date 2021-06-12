import defaultMapGenSettings from "./default-map-gen-settings";

export default {
    ...defaultMapGenSettings,
    "autoplace_controls":
        {
            "coal": {"frequency": 1, "size": 1, "richness": 2},
            "stone": {"frequency": 1, "size": 1, "richness": 2},
            "copper-ore": {"frequency": 1, "size": 1,"richness": 2},
            "iron-ore": {"frequency": 1, "size": 1, "richness": 2},
            "uranium-ore": {"frequency": 1, "size": 1, "richness": 2},
            "crude-oil": {"frequency": 1, "size": 1, "richness": 2},
            "trees": {"frequency": 1, "size": 1, "richness": 2},
            "enemy-base": {"frequency": 1, "size": 1, "richness": 2}
        },
}