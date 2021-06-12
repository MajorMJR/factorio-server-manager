import React from "react";
import Select from "../../../components/Select";
import Checkbox from "../../../components/Checkbox";
import Input from "../../../components/Input";
import Label from "../../../components/Label";
import copy from "../../../copy";

const Terrain = ({settings, setSettings}) => {
    return <div className="flex-1">
        <div className="flex justify-between border border-black rounded p-2 mb-1">
            <div className="self-center">
                <Label text="Maptype"/>
            </div>
            <Select isInline
                    value={settings?.property_expression_names?.elevation || 'normal'}
                    options={[
                        {value: 'normal', name: 'Normal'},
                        {value: '0_17-island', name: 'Island'},
                        {value: '0_16-elevation', name: 'Normal 0.16'},
                    ]}
                    onChange={event => {
                        let tmp = copy(settings);

                        switch (event.target.value) {
                            case '0_17-island':
                                tmp.property_expression_names.elevation = '0_17-island';
                                break;
                            case '0_16-elevation':
                                tmp.property_expression_names.elevation = '0_16-elevation';
                                break;
                            case 'normal':
                            default:
                                tmp.property_expression_names.elevation = '';
                                break;
                        }

                        setSettings(tmp);
                    }}
            />
        </div>
        <div className="border border-black rounded p-2 mb-1">
            <table className="w-full table-fixed">
                <thead>
                <tr>
                    <th className="w-3/5"/>
                    <th className="w-1/5">Scale</th>
                    <th className="w-1/5">Coverage</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                    <td>
                        <Checkbox text="Water"
                                  checked={settings?.water_enabled || true}
                                  onChange={event => {
                                      let tmp = copy(settings);
                                      tmp.water_enabled = event.target.checked;
                                      setSettings(tmp);
                                  }}
                        />
                    </td>
                    <td>
                        <Input
                            value={1 / (settings?.water || 1)}
                            disabled={!(typeof settings?.water_enabled !== 'undefined' ? settings?.water_enabled : true)}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.water = 1 / parseInt(event.target.value);
                                setSettings(tmp);
                            }}
                        />
                    </td>
                    <td>
                        <Input
                            value={settings?.water || 1}
                            disabled={!(typeof settings?.water_enabled !== 'undefined' ? settings?.water_enabled : true)}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.water = parseInt(event.target.value);
                                setSettings(tmp);
                            }}
                        />
                    </td>
                </tr>
                <tr>
                    <td>
                        <Checkbox text="Trees"
                                  checked={settings.autoplace_controls.trees.richness === 1}
                                  onChange={event => {
                                      let tmp = copy(settings);
                                      tmp.autoplace_controls.trees.richness = event.target.checked ? 1 : 0;
                                      setSettings(tmp)
                                  }}
                        />
                    </td>
                    <td>
                        <Input
                            disabled={!settings?.autoplace_controls?.trees?.richness}
                            value={settings?.autoplace_controls?.trees?.size || 1}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.autoplace_controls.trees.size = parseFloat(event.target.value);
                                setSettings(tmp)
                            }}
                        />
                    </td>
                    <td>
                        <Input
                            disabled={!settings?.autoplace_controls?.trees?.richness}
                            value={settings?.autoplace_controls?.trees?.frequency || 1}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.autoplace_controls.trees.frequency = parseFloat(event.target.value);
                                setSettings(tmp)
                            }}
                        />
                    </td>
                </tr>
                </tbody>
            </table>
        </div>
        <div className="border border-black rounded p-2 mb-1">
            <table className="w-full table-fixed">
                <thead>
                <tr>
                    <th className="w-3/5"/>
                    <th className="w-1/5">Frequency</th>
                    <th className="w-1/5">Continuity</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                    <td>
                        <Checkbox text="Cliffs"
                                  checked={settings?.cliff_settings?.enabled || true}
                                  onChange={event => {
                                      let tmp = copy(settings);
                                      tmp.cliff_settings.enabled = event.target.checked;
                                      setSettings(tmp);
                                  }}
                        />
                    </td>
                    <td>
                        <Input
                            value={40 / (settings?.cliff_settings?.cliff_elevation_interval || 40)}
                            disabled={!(typeof settings?.cliff_settings?.enabled !== 'undefined' ? settings.cliff_settings.enabled : true)}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.cliff_settings.cliff_elevation_interval = 40 / parseInt(event.target.value);
                                setSettings(tmp);
                            }}
                        />
                    </td>
                    <td>
                        <Input
                            value={settings?.cliff_settings?.richness || 1}
                            disabled={!(typeof settings?.cliff_settings?.enabled !== 'undefined' ? settings.cliff_settings.enabled : true)}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.cliff_settings.richness = parseInt(event.target.value);
                                setSettings(tmp);
                            }}
                        />
                    </td>
                </tr>
                </tbody>
            </table>
        </div>
        <div className="border border-black rounded p-2">
            <table className="w-full table-fixed">
                <thead>
                <tr>
                    <th className="w-3/5"/>
                    <th className="w-1/5">Scale</th>
                    <th className="w-1/5">Bias</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                    <td>Moisture</td>
                    <td>
                        <Input
                            value={settings?.property_expression_names?.['control-setting:moisture:frequency:multiplier'] || "1"}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.property_expression_names['control-setting:moisture:frequency:multiplier'] = event.target.value;
                                setSettings(tmp);
                            }}
                        />
                    </td>
                    <td>
                        <Input
                            value={settings?.property_expression_names?.['control-setting:moisture:bias'] || "0"}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.property_expression_names['control-setting:moisture:bias'] = event.target.value;
                                setSettings(tmp);
                            }}
                        />
                    </td>
                </tr>
                <tr>
                    <td>Terrain Typ</td>
                    <td>
                        <Input
                            value={1 / parseInt(settings?.property_expression_names?.['control-setting:aux:frequency:multiplier'] || "1")}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.property_expression_names['control-setting:aux:frequency:multiplier'] = `${1 / event.target.value}`;
                                setSettings(tmp);
                            }}
                        />
                    </td>
                    <td>
                        <Input
                            value={settings?.property_expression_names?.['control-setting:aux:bias'] || "0"}
                            onChange={event => {
                                let tmp = copy(settings);
                                tmp.property_expression_names['control-setting:aux:bias'] = event.target.value;
                                setSettings(tmp);
                            }}
                        />
                    </td>
                </tr>
                </tbody>
            </table>
        </div>
    </div>
}

export default Terrain;