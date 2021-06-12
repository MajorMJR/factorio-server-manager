import React from "react";
import ResourceConfigurator from "./components/ResourceConfigurator";

const Resources = ({settings, setSettings}) => {
    return <table className="flex-1 border border-black rounded p-2">
        <thead>
            <tr className="border border-black rounded">
                <th className="border border-black"/>
                <th className="border border-black p-2 font-normal">Frequency</th>
                <th className="border border-black p-2 font-normal">Size</th>
                <th className="border border-black p-2 font-normal">Richness</th>
            </tr>
        </thead>
        <tbody>
            <ResourceConfigurator settings={settings} setSettings={setSettings} label="Iron Ore" name="iron-ore"/>
            <ResourceConfigurator settings={settings} setSettings={setSettings} label="Copper Ore" name="copper-ore"/>
            <ResourceConfigurator settings={settings} setSettings={setSettings} label="Stone Ore" name="stone"/>
            <ResourceConfigurator settings={settings} setSettings={setSettings} label="Coal" name="coal"/>
            <ResourceConfigurator settings={settings} setSettings={setSettings} label="Uranium Ore" name="uranium-ore"/>
            <ResourceConfigurator settings={settings} setSettings={setSettings} label="Crude Oil" name="crude-oil"/>
        </tbody>
    </table>
}

export default Resources;