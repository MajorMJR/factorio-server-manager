import React from "react";
import Button from "../../../components/Button";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faRandom} from "@fortawesome/free-solid-svg-icons";



const SeedInput = ({seed, updateSeed, generateRandomSeed}) => {

    return <div className="relative">
        <div className="w-32 inline-block">
            <input
                className="shadow border-t-2 border-gray-light bg-gray-light h-8 rounded-l-sm appearance-none w-full px-2 py-2 text-black"
                value={seed}
                onChange={event => {
                    const value = parseInt(event.target.value)
                    updateSeed(value)
                }}/>
        </div>

        <Button size="none" onClick={generateRandomSeed} className="text-center ml-1 rounded-r-sm w-8 h-8">
            <FontAwesomeIcon icon={faRandom}/>
        </Button>
    </div>
}

export default SeedInput;