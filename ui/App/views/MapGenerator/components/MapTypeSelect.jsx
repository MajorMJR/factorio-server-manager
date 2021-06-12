import React, {useState} from "react";
import saves from "../../../../api/resources/saves";

const MapTypeSelect = ({settings, setSettings}) => {

    const [value, setValue] = useState("default");

    const options = [
        {
            name: "Default",
            value: "default"
        },
        {
            name: "Rich resources",
            value: "rich-resources"
        },
        {
            name: "Marathon",
            value: "marathon"
        },
        {
            name: "Death world",
            value: "death-world"
        },
        {
            name: "Death world marathon",
            value: "death-world-marathon"
        },
        {
            name: "Rail world",
            value: "rail-world"
        },
        {
            name: "Ribbon world",
            value: "ribbon-world"
        },
        {
            name: "Island",
            value: "island"
        },
    ];

    const change = event => {
        event.persist();
        setValue(event.target.value);

        let tmpSettings = {};

        // reset all settings to default values
        Promise.all([
            saves.defaultMapGenSettings()
                .then(mapGenSettings => tmpSettings = Object.assign(tmpSettings, mapGenSettings)),
            saves.defaultMapSettings()
                .then(mapSettings => tmpSettings = Object.assign(tmpSettings, mapSettings)),
        ]).then(() => {

            // use the same seed as before
            tmpSettings.seed = settings.seed;

            let resourceSettings = {
                richness: 1,
                size: 1,
                frequency: 1
            }

            // adjust the settings based on the selected template
            // source: https://wiki.factorio.com/Map_generator#Map_generation_presets
            switch (event.target.value) {
                case 'rich-resources':
                    tmpSettings.autoplace_controls.coal.richness = 2;
                    tmpSettings.autoplace_controls["iron-ore"].richness = 2;
                    tmpSettings.autoplace_controls["copper-ore"].richness = 2;
                    tmpSettings.autoplace_controls["uranium-ore"].richness = 2;
                    tmpSettings.autoplace_controls.stone.richness = 2;
                    tmpSettings.autoplace_controls["crude-oil"].richness = 2;
                    break;
                case 'marathon':
                    tmpSettings.difficulty_settings.recipe_difficulty = 2;
                    tmpSettings.difficulty_settings.technology_difficulty = 2;
                    tmpSettings.difficulty_settings.technology_price_multiplier = 4;
                    break;
                case 'death-world':
                    tmpSettings.starting_area = 0.75;
                    tmpSettings.autoplace_controls["enemy-base"].frequency = 2;
                    tmpSettings.autoplace_controls["enemy-base"].size = 2;
                    tmpSettings.pollution.absorption_modifier = 0.5; // todo not found in map-(gen-)settings
                    tmpSettings.pollution.enemy_attack_pollution_consumption_modifier = 0.5;
                    break;
                case 'death-world-marathon':
                    // from death-world
                    tmpSettings.starting_area = 0.75;
                    tmpSettings.enemy_evolution.time_factor = 200;
                    tmpSettings.enemy_evolution.pollution_factor = 12;
                    tmpSettings.pollution.absorption_modifier = 0.5; // todo not found in map-(gen-)settings
                    tmpSettings.pollution.enemy_attack_pollution_consumption_modifier = 0.5;

                    // from marathon
                    tmpSettings.difficulty_settings.recipe_difficulty = 2;
                    tmpSettings.difficulty_settings.technology_difficulty = 2;
                    tmpSettings.difficulty_settings.technology_price_multiplier = 4;
                    break;
                case 'rail-world':
                    resourceSettings = {
                        richness: 1,
                        size: 3,
                        frequency: 0.33
                    }
                    tmpSettings.autoplace_controls.coal = resourceSettings;
                    tmpSettings.autoplace_controls["iron-ore"] = resourceSettings;
                    tmpSettings.autoplace_controls["copper-ore"] = resourceSettings;
                    tmpSettings.autoplace_controls["uranium-ore"] = resourceSettings;
                    tmpSettings.autoplace_controls.stone = resourceSettings;
                    tmpSettings.autoplace_controls["crude-oil"] = resourceSettings;
                    tmpSettings.enemy_expansion.enabled = false;
                    tmpSettings.enemy_evolution.time_factor = 20;
                    tmpSettings.water = 2;
                    // todo: Water is set to 200% scale and 150% coverage
                    break;
                case 'ribbon-world':
                    tmpSettings.height = 128;

                    resourceSettings = {
                        richness: 2,
                        size: 0.5,
                        frequency: 3
                    }
                    tmpSettings.autoplace_controls.coal = resourceSettings;
                    tmpSettings.autoplace_controls["iron-ore"] = resourceSettings;
                    tmpSettings.autoplace_controls["copper-ore"] = resourceSettings;
                    tmpSettings.autoplace_controls["uranium-ore"] = resourceSettings;
                    tmpSettings.autoplace_controls.stone = resourceSettings;
                    tmpSettings.autoplace_controls["crude-oil"] = resourceSettings;
                    tmpSettings.starting_area = 3;
                    // todo: Water is set to 25% coverage and size.
                    break;
                case 'island':
                    tmpSettings.property_expression_names.elevation = '0_17-island';
                    break;
                case 'default':
                default:
                    break;
            }

            setSettings(tmpSettings);
        });
    }

    return <div className="relative ">
        <select
            className="shadow appearance-none bg-gray-light w-full h-8 px-1 text-black"
            name={name}
            id={name}
            value={value}
            onChange={change}
        >
            {options.map(option => <option value={option.value} key={option.value}>{option.name}</option>)}
        </select>
        <div
            className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-black">
            <svg className="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg"
                 viewBox="0 0 20 20">
                <path
                    d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"/>
            </svg>
        </div>
    </div>
}

export default MapTypeSelect;