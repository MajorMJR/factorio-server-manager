import React from "react";

const Slider = ({min = 1, max = 12, step = 1}) => {
    return <input
        type="range"
        className="slider"
        defaultValue={6}
        min={min}
        max={max}
        step={step}/>
}

export default Slider;