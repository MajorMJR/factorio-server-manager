import React from "react";

const Label = ({text, htmlFor, isInline = false}) => {
    return (
        <label
            className={"text-white text-sm font-bold mb-1 " + (isInline ? "inline-block mx-2" : "block")} htmlFor={htmlFor}>
            {text}
        </label>
    )
}

export default Label;