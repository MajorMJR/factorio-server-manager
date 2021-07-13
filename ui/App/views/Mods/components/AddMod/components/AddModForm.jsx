import React, {useEffect, useState} from "react";
import modsResource from "../../../../../../api/resources/mods";
import Button from "../../../../../components/Button";
import Label from "../../../../../components/Label";
import {useForm} from "react-hook-form";
import Input from "../../../../../components/Input";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faExternalLinkAlt} from "@fortawesome/free-solid-svg-icons/faExternalLinkAlt";
import {faSpinner} from "@fortawesome/free-solid-svg-icons";
import SelectVersionForm from "./SelectVersionForm";

const LinkModPortal = () => {
    return <a href="https://mods.factorio.com" target="_blank" className="px-2 text-blue hover:text-blue-light">Mod
        Portal <FontAwesomeIcon icon={faExternalLinkAlt}/></a>
}

const AddModForm = ({setIsFactorioAuthenticated, fuse, refetchInstalledMods}) => {

    const {register, watch, setValue, handleSubmit} = useForm();
    const [suggestedMods, setSuggestedMods] = useState([]);
    const [selectedMod, setSelectedMod] = useState(null);
    const [hoveredMod, setHoveredMod] = useState(0)
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [releases, setReleases] = useState([]);

    const [autocomplete, setAutocomplete] = useState(NaN);
    const mod = watch('mod');

    const logout = () => {
        modsResource.portal.logout()
            .then(() => setIsFactorioAuthenticated(false));
    }

    const updateSuggestedMods = () => {
        if (typeof fuse != "undefined") {
            setHoveredMod(0)
            clearTimeout(autocomplete)
            setAutocomplete(setTimeout(() => setSuggestedMods(fuse.search(mod || '')), 200));
        }
    };

    const selectMod = mod => {
        clearTimeout(autocomplete);
        setValue('mod', mod.item.title); // triggers effect for mod
        setSuggestedMods([]);
        setSelectedMod(mod);
    }

    // get triggered if mod changed
    useEffect(() => {
        if (selectedMod === null) {
            updateSuggestedMods();
        } else if (selectedMod.item.title !== mod) {
            setSelectedMod(null);
            setReleases([]);
            updateSuggestedMods();
        }
    }, [mod]);

    const openSelectVersionModal = async data => {
        const mod = await modsResource.portal.info(selectedMod.item.name);
        setReleases(mod.releases || []);
        setIsModalOpen(true);
    }

    const install = async release => {
        return modsResource.portal
            .install(release.download_url, release.file_name, selectedMod.item.name)
            .then(refetchInstalledMods)
    }

    const incrementHoveredMod = () => {
        if (hoveredMod < suggestedMods.length - 1) {
            setHoveredMod(hoveredMod + 1)
        }
    }

    const decrementHoveredMod = () => {
        if (hoveredMod > 0) {
            setHoveredMod(hoveredMod - 1)
        }
    }

    const handleKeyDown = event => {
        switch (event.keyCode) {
            case 40: // down
                incrementHoveredMod()
                break;
            case 38: // up
                decrementHoveredMod()
                break;
            case 13: // enter
                selectMod(suggestedMods[hoveredMod])
                break;
            default:
                break;
        }
    }

    return (
        <form onSubmit={handleSubmit(openSelectVersionModal)}>
            <SelectVersionForm isOpen={isModalOpen} releases={releases} install={install} close={() => setIsModalOpen(false)}/>
            <div className="mb-4 relative" >
                <Label text="Mod" htmlFor="mod"/>
                { typeof fuse !== "undefined"
                    ? <Input inputRef={register({required: true})} hasAutoComplete={false} name="mod" onKeyDown={handleKeyDown}/>
                    : <div className="border border-gray-medium w-full py-2 px-3 text-white">
                        <FontAwesomeIcon icon={faSpinner} spin={true}/> Loading List of Mods from <LinkModPortal/>
                    </div>
                }
                {suggestedMods.length > 0 &&
                    <ul className="bg-white text-black h-64 overflow-y-scroll absolute bottom-0 left-0 w-full -mb-64">
                        {suggestedMods.map((mod, index) => <li className={"px-2 py-1 cursor-pointer" + (hoveredMod === index ? " bg-blue-light" : "")} onMouseEnter={() => setHoveredMod(index)} onClick={() => selectMod(mod)} key={index}>{mod.item.title}</li>)}
                    </ul>
                }
            </div>
            <Button isDisabled={selectedMod === null} isSubmit={true} onClick={() => setIsModalOpen(true)} className="mr-2">Install</Button>
            <Button onClick={logout} type="danger" className="mr-2">Logout</Button>
            <LinkModPortal/>
        </form>
    )
}

export default AddModForm;