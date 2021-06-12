import React, {useEffect, useState} from "react";
import TabControl from "../../components/Tabs/TabControl";
import Tab from "../../components/Tabs/Tab";
import Button from "../../components/Button";
import Resources from "./tabs/resources/Resources";
import Terrain from "./tabs/Terrain";
import Enemy from "./tabs/Enemy";
import Advanced from "./tabs/Advanced";
import SeedInput from "./components/SeedInput";
import MapTypeSelect from "./components/MapTypeSelect";
import saves from "../../../api/resources/saves";
import MapPreviewImage from "./components/MapPreviewImage";
import copy from "../../copy";
import Input from "../../components/Input";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faExclamationTriangle} from "@fortawesome/free-solid-svg-icons";

let timeoutPreviewHandle = null;

const MapGenerator = ({serverStatus}) => {

    const isServerRunning = serverStatus.status === 'running';

    const [isPreviewDisplayed, setIsPreviewDisplayed] = useState(false)
    const [seed, setSeed] = useState(0);
    const [settings, setSettings] = useState({});
    const [previewImage, setPreviewImage] = useState(null);
    const [previewImageSeed, setPreviewImageSeed] = useState(null);
    const [isLoadingPreview, setIsLoadingPreview] = useState(false);
    const [saveFileName, setSaveFileName] = useState("");
    const [isGeneratingMap, setIsGeneratingMap] = useState(false);

    const updateSettings = (newSettings, shouldLoadPreview = true) => {
        setSettings(newSettings);
        if (shouldLoadPreview) {
            loadPreview(newSettings)
        }
    }

    const createSave = () => {
        setIsGeneratingMap(true);
        saves.create(saveFileName, settings)
            .then(() => flash(`Save "${saveFileName}" created`, "green"))
            .finally(() => setIsGeneratingMap(false));
    }

    const loadPreview = (settings, force = false) => {
        if (isLoadingPreview || (!isPreviewDisplayed && !force)) {
            return;
        }

        const previewHandler = () => {
            setIsLoadingPreview(true)
            setPreviewImageSeed(settings.seed)

            let tmpSettings = copy(settings);

            if (!tmpSettings.cliff_settings.enabled) {
                tmpSettings.cliff_settings.richness = 0;
            }

            if (!tmpSettings.water_enabled) {
                tmpSettings.water = 0;
            }

            saves.preview(tmpSettings)
                .then(imageData => setPreviewImage(imageData))
                .finally(() => setIsLoadingPreview(false))
        }

        if (timeoutPreviewHandle) {
            clearTimeout(timeoutPreviewHandle);
            timeoutPreviewHandle = null
        }
        timeoutPreviewHandle = setTimeout(previewHandler, 600);
    }

    const randomSeed = () => {
        const randomValue = Math.floor(Math.random() * 1000000000)
        updateSeed(randomValue)
    }

    const updateSeed = value => {
        setSeed(value);
        const newSettings = Object.assign(settings, {seed: value});
        setSettings(newSettings);
        loadPreview(newSettings);
    }

    useEffect(() => {
        Promise.all([
            saves.defaultMapGenSettings()
                .then(mapGenSettings => setSettings(Object.assign(settings, mapGenSettings))),
            saves.defaultMapSettings()
                .then(mapSettings => setSettings(Object.assign(settings, mapSettings))),
        ]).finally(() => {
            randomSeed()
        })
    }, []);

    return <TabControl
        actions={
            <div className="flex justify-between">
                <div>
                    <Input
                        size="sm"
                        isInline
                        placeholder="Save name"
                        value={saveFileName}
                        disabled={isGeneratingMap}
                        onChange={event => setSaveFileName(event.target.value)}
                    />
                    <Button
                        size="sm"
                        isDisabled={!saveFileName || isServerRunning}
                        isLoading={isGeneratingMap}
                        className="inline-block ml-1"
                        type="success"
                        onClick={createSave}
                    >Generate Map</Button>
                    { isServerRunning && <span className="ml-3 text-red">
                        <FontAwesomeIcon icon={faExclamationTriangle}/> Server mus be stopped before generating map.
                    </span>}
                </div>
                {isPreviewDisplayed
                    ? <Button size="sm" onClick={() => setIsPreviewDisplayed(false)}>Hide Preview</Button>
                    : <Button size="sm" onClick={() => {
                        setIsPreviewDisplayed(true)
                        loadPreview(settings,true)
                    }
                    }>Show Preview</Button>
                }
            </div>
        }
        title={
            <div className="flex justify-between my-1">
                <MapTypeSelect settings={settings} setSettings={updateSettings}/>
                <SeedInput updateSeed={updateSeed} seed={seed} generateRandomSeed={randomSeed}/>
            </div>
        }
    >
        <Tab title="Resources">
            <div className="flex space-x-8">
                <Resources settings={settings} setSettings={updateSettings}/>
                <MapPreviewImage imageData={previewImage} isLoading={isLoadingPreview} show={isPreviewDisplayed} seed={previewImageSeed}/>
            </div>
        </Tab>
        <Tab title="Terrain">
            <div className="flex space-x-8">
                <Terrain settings={settings} setSettings={updateSettings}/>
                <MapPreviewImage imageData={previewImage} isLoading={isLoadingPreview} show={isPreviewDisplayed} seed={previewImageSeed}/>
            </div>
        </Tab>
        <Tab title="Enemy">
            <div className="flex space-x-8">
                <Enemy settings={settings} setSettings={updateSettings}/>
                <MapPreviewImage imageData={previewImage} isLoading={isLoadingPreview} show={isPreviewDisplayed} seed={previewImageSeed}/>
            </div>
        </Tab>
        <Tab title="Advanced">
            <div className="flex space-x-8">
                <Advanced settings={settings} setSettings={updateSettings}/>
                <MapPreviewImage imageData={previewImage} isLoading={isLoadingPreview} show={isPreviewDisplayed} seed={previewImageSeed}/>
            </div>
        </Tab>
    </TabControl>
}

export default MapGenerator;