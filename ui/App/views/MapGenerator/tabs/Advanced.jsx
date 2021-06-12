import React from "react";
import Label from "../../../components/Label";
import Input from "../../../components/Input";
import Select from "../../../components/Select";
import Checkbox from "../../../components/Checkbox";
import copy from "../../../copy";

const Advanced = ({settings, setSettings}) => {
    return <div className="flex-1">
        <div className="flex justify-between border border-black rounded p-2 mb-1">
            <div className="text-dirty-white font-bold">
                Map
            </div>
            <div className="text-right">
                <div className="mb-1">
                    <Label isInline text="Height"/>
                    <Input isInline
                           value={settings?.height ? settings.height : 0}
                           onChange={event => {
                               let tmp = copy(settings)
                               tmp.height = event.target.value
                               setSettings(tmp)
                           }}
                    />
                </div>
                <div>
                    <Label isInline text="Width"/>
                    <Input isInline
                           value={settings?.width ? settings.width : 0}
                           onChange={event => {
                               let tmp = copy(settings)
                               tmp.width = event.target.value
                               setSettings(tmp)
                           }}
                    />
                </div>
            </div>
        </div>
        <div className="flex justify-between border border-black rounded p-2 mb-1">
            <div className="text-dirty-white font-bold">
                Recipes
            </div>
            <div className="text-right">
                <Label isInline text="Difficulty"/>
                <Select isInline
                        options={[
                            {value: 0, name: 'Normal'},
                            {value: 1, name: 'Expensive'}
                        ]}
                        onChange={event => {
                            let tmp = copy(settings)
                            tmp.difficulty_settings.recipe_difficulty = event.target.value
                            setSettings(tmp, false)
                        }}
                />
            </div>
        </div>
        <div className="flex justify-between border border-black rounded p-2 mb-1">
            <div className="text-dirty-white font-bold">
                Technology
            </div>
            <div className="text-right">
                <div className="mb-1">
                    <Label isInline text="Difficulty"/>
                    <Select isInline
                            options={[
                                {value: 0, name: 'Normal'},
                                {value: 1, name: 'Expensive'}
                            ]}
                            onChange={event => {
                                let tmp = copy(settings)
                                tmp.difficulty_settings.technology_difficulty = event.target.value
                                setSettings(tmp, false)
                            }}
                    />
                </div>
                <div className="mb-1">
                    <Label isInline text="Price multiplier"/>
                    <Input isInline
                           value={settings?.difficulty_settings?.technology_price_multiplier || 1}
                           onChange={event => {
                               let tmp = copy(settings)
                               tmp.difficulty_settings.technology_price_multiplier = parseInt(event.target.value)
                               setSettings(tmp)
                           }}
                    />
                </div>
                <div>
                    <Label isInline text="Research queue availability"/>
                    <Select isInline
                            options={[
                                {value: 'after-victory', name: 'After rocket launch'},
                                {value: 'always', name: 'Always'},
                                {value: 'never', name: 'Never'}
                            ]}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.difficulty_settings.research_queue_setting = event.target.value;
                                setSettings(settings, false);
                            }}
                    />
                </div>
            </div>
        </div>
        <div className="flex justify-between border border-black rounded p-2">
            <Checkbox checked={settings?.pollution?.enabled} onChange={event => {
                let tmp = copy(settings)
                tmp.pollution.enabled = event.target.checked
                setSettings(tmp)
            }} text="Pollution" className="text-dirty-white" textSize="md"/>
            <div className="text-right">
                <div className="mb-1">
                    <Label isInline text="Absorption modifier"/>
                    <Input isInline
                           disabled={!settings?.pollution?.enabled}
                           value={settings?.pollution?.absorption_modifier || 0}
                           onChange={event => {
                               let tmp = copy(settings)
                               tmp.pollution.absorption_modifier = parseFloat(event.target.value);
                               setSettings(tmp)
                           }}
                    />
                </div>
                <div className="mb-1">
                    <Label isInline text="Attack cost modifier"/>
                    <Input isInline
                           disabled={!settings?.pollution?.enabled}
                           value={settings?.pollution?.attack_cost_modifier || 0}
                           onChange={event => {
                               let tmp = copy(settings)
                               tmp.pollution.attack_cost_modifier = parseFloat(event.target.value);
                               setSettings(tmp)
                           }}
                    />
                </div>
                <div className="mb-1">
                    <Label isInline text="Minimum to damage trees"/>
                    <Input isInline
                           disabled={!settings?.pollution?.enabled}
                           value={settings?.pollution?.min_pollution_to_damage_trees || 0}
                           onChange={event => {
                               let tmp = copy(settings)
                               tmp.pollution.minimum_to_damage_trees = event.target.value;
                               setSettings(tmp)
                           }}
                    />
                </div>
                <div className="mb-1">
                    <Label isInline text="Absorbed per damaged tree"/>
                    <Input isInline
                           disabled={!settings?.pollution?.enabled}
                           value={settings?.pollution?.absorbed_per_damaged_tree || 0}
                           onChange={event => {
                               let tmp = copy(settings)
                               tmp.pollution.absorbed_per_damaged_tree = event.target.value;
                               setSettings(tmp)
                           }}
                    />
                </div>
                <div className="mb-1">
                    <Label isInline text="Diffusion ratio"/>
                    <Input isInline
                           disabled={!settings?.pollution?.enabled}
                           value={settings?.pollution?.diffusion_ratio || 0}
                           onChange={event => {
                               let tmp = copy(settings)
                               tmp.pollution.diffusion_ratio = event.target.value;
                               setSettings(tmp)
                           }}
                    />
                </div>
            </div>
        </div>
    </div>
}

export default Advanced;